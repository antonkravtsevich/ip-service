package main

import "fmt"
import "net/http"
import "encoding/json"

type Log struct{
	Ip	string
}


type Server struct{
	ipsChan chan<- string
}

func isArrayContainValue(array []string, value string) bool{
	for _, element := range array{
		if element == value {
			return true
		}
	}
	return false
}

func isIPUnique() chan <- string{
	var uniqueIpsList []string
	uniqueIpsCount := 0

	ipsChan := make(chan string)

	go func(){

		for newIp := range ipsChan{
			unique := !isArrayContainValue(uniqueIpsList, newIp)
			if unique {
				uniqueIpsList = append(uniqueIpsList, newIp)
				uniqueIpsCount += 1
				fmt.Printf("new unique IP! %v\n", newIp)
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

	server.ipsChan <- log.Ip

	fmt.Fprintf(w, "ok")
}

// Provide Prometheus metrics 
func (server *Server) returnMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {
	ipsChannel := isIPUnique()
	server := Server{
		ipsChan: ipsChannel,
	}

	acceptJsonServerMux := http.NewServeMux()
    acceptJsonServerMux.HandleFunc("/logs", server.acceptJson)

    returnMetricServer := http.NewServeMux()
    returnMetricServer.HandleFunc("/metrics", server.returnMetric)

    go func() {
        http.ListenAndServe("localhost:5000", acceptJsonServerMux)
    }()

    http.ListenAndServe("localhost:9092", returnMetricServer)
}