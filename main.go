package main

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"

// types of commands that's can be used through main channel
const (
	AddIp = iota
	GetCount
)

type CommandType int

// Structure of commands, that will be used for data exchange between endpoints
type Command struct{
	commandType CommandType
	payload string				// new IPs will be sent in this field
	responseChannel	chan int  	// channel, that will be used to retrieve new data and for 
								// /logs endpoint blocking in case if new IP should be added
}

type Log struct{
	Ip	string
}

// provide access to command channel for both endpoints
type Server struct{
	commandsChannel chan<- Command
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
// calculate count of unique IPs 
func startIPsCounter() chan <- Command{

	// as a temporary solution, all unique IPs will be stored in list
	// that will be used to check if received IP is unique or not
	var uniqueIpsList []string
	uniqueIpsCount := 0

	commandsChannel := make(chan Command)

	// run commands handling
	go func(){
		for command := range commandsChannel{
			switch command.commandType{
			// handle command for new IP adding
			case AddIp:
				newIp := command.payload
				unique := !isArrayContainValue(uniqueIpsList, newIp)

				// if IP unique - add it to list, increment counter
				if unique {
					uniqueIpsList = append(uniqueIpsList, newIp)
					uniqueIpsCount += 1
				}
				command.responseChannel <- 0
			// handle command for unique IPs count recieving
			case GetCount:
				command.responseChannel <- uniqueIpsCount
			}
		}
	}()

	return commandsChannel
}

// Accept log record and extract IP address from it
func (server *Server) acceptJson(w http.ResponseWriter, r *http.Request) {
	var log Log

	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseChannel := make(chan int)

	// send command with new IP, wait for response channel
	command := Command{AddIp, log.Ip, responseChannel}
	server.commandsChannel <- command
	<- responseChannel

	fmt.Fprintf(w, "ok")
}

// Provide Prometheus metrics 
func (server *Server) returnMetric(w http.ResponseWriter, r *http.Request) {
	responseChannel := make(chan int)

	// send command for unique IPs count retrieving, recieve value 
	// in response channel
	command := Command{GetCount, "", responseChannel}
	server.commandsChannel <- command
	count := <- responseChannel
	fmt.Fprintf(w, strconv.Itoa(count))
}

func main() {
	commandsChannel := startIPsCounter()
	server := Server{
		commandsChannel: commandsChannel,
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