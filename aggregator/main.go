package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/andrei-kozel/toll-calculator/types"
)

func main() {
	var (
		store      = NewMemoryStore()
		svc        = NewInvoiceAggregator(store)
		listenAddr = flag.String("listenAddr", ":3000", "the listenAddr of the http server")
	)
	flag.Parse()
	makeHTTPTransport(*listenAddr, svc)
}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Println("HTTP transport running on port", listenAddr)
	http.HandleFunc("/agg", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
