package main

import "fmt"
import "net/http"
import "encoding/json"
import "github.com/prometheus/client_golang/prometheus"
import "github.com/prometheus/client_golang/prometheus/promhttp"

type Log struct{
	Ip	string
}

// provide access to channel that will be used for IPs sharing
type Server struct{
	ipsChan chan<-string
}

// check if element is in list
func isArrayContainValue(array []string, value string) bool{
	for _, element := range array{
		if element == value {
			return true
		}
	}
	return false
}

// check if IP is unique
// return channel, that should be used for IPs values sending
func startIPsCounter(ipsCounter prometheus.Counter) chan <- string{

	// tree will be used as storage for all unique IPs
	tree := Node{' ', []*Node{}}
	// all unique IPs will be splitted in symbols, each symbol will be stored
	// as node in tree

	ipsChan := make(chan string)

	// run commands handling
	go func(){
		for newIp := range ipsChan{
			// split IP in symbols for further processing
			newIpSplitted := []rune(newIp)
			unique := !tree.isBranchExists(newIpSplitted)

			// if IP unique - add it to tree, increment counter
			if unique {
				tree.insertBranch(newIpSplitted)
				ipsCounter.Inc()
			}
		}
	}()

	return ipsChan
}

// Accept log record and extract IP address from it
func (server *Server) acceptJson(w http.ResponseWriter, r *http.Request) {
	var log Log

	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send command with new IP, wait for response channel
	server.ipsChan <- log.Ip

	fmt.Fprintf(w, "ok")
}

func main() {
	ipsCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "unique_ip_addresses",
		})
	prometheus.MustRegister(ipsCounter)

	ipsChan := startIPsCounter(ipsCounter)
	server := Server{
		ipsChan: ipsChan,
	}

	acceptJsonServerMux := http.NewServeMux()
    acceptJsonServerMux.HandleFunc("/logs", server.acceptJson)

    go func() {
        http.ListenAndServe("localhost:5000", acceptJsonServerMux)
    }()

    http.ListenAndServe("localhost:9102", promhttp.Handler())
}