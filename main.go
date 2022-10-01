package main

import (
	"log"
	"net/http"

	"github.com/gndw/go-singleflight-example/endpoints"
)

func main() {

	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/ok-with-data", endpoints.OkWithData)
	http.HandleFunc("/ok-with-sf", endpoints.OkWithSingleflight)
	http.HandleFunc("/ok-with-sf-key", endpoints.OkWithSingleflightKey)
	http.HandleFunc("/ok-with-sf-pointer", endpoints.OkWithSingleflightPointer)

	log.Println("server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
