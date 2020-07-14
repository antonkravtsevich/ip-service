package main

import "fmt"
import "net/http"

func acceptJson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

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