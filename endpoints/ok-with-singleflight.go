package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gndw/go-singleflight-example/model"
)

var counter = 0

func OkWithSingleflight(w http.ResponseWriter, r *http.Request) {

	response := model.BasicResponse{}

	counter++
	id := counter
	log.Printf("request %v is coming...", id)
	defer func() {
		log.Printf("request %v is ending with shared: %v", id, response.IsUsingSF)
	}()

	// do external call database using singleflight
	v, err, shared := requestGroup.Do("key", func() (interface{}, error) {
		return DoExternalCallToDatabase()
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// cast singleflight response to basic response
	response = v.(model.BasicResponse)

	// modify response if shared
	response.IsUsingSF = shared

	// write response to http
	responseInByte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseInByte)

}
