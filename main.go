package main

import "fmt"
import "net/http"
import "encoding/json"

type Log struct{
	Ip	string
}

// Accept log record and extract IP address from it
func acceptJson(w http.ResponseWriter, r *http.Request) {
	var log Log

	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, log.Ip)
}

// Provide Prometheus metrics 
func returnMetric(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {
	acceptJsonServerMux := http.NewServeMux()
    acceptJsonServerMux.HandleFunc("/logs", acceptJson)

    returnMetricServer := http.NewServeMux()
    returnMetricServer.HandleFunc("/metrics", returnMetric)

    go func() {
        http.ListenAndServe("localhost:5000", acceptJsonServerMux)
    }()

    http.ListenAndServe("localhost:9092", returnMetricServer)
}