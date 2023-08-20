package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gndw/go-singleflight-example/endpoints"
	"github.com/gndw/go-singleflight-example/external"
	"golang.org/x/sync/singleflight"
)

func main() {

	externalService := external.New(&sync.Mutex{})
	endpointService := endpoints.New(&singleflight.Group{}, externalService)

	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/ok-with-data", endpointService.OkWithData)
	http.HandleFunc("/ok-with-sf", endpointService.OkWithSingleflight)
	http.HandleFunc("/get-concurrent", endpointService.GetConcurrent)
	http.HandleFunc("/ok-with-sf-pointer", endpointService.OkWithSingleflightPointer)

	log.Println("server running on port 4000...")
	http.ListenAndServe(":4000", nil)
}
