package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gndw/go-singleflight-example/model"
)

func OkWithSingleflightKey(w http.ResponseWriter, r *http.Request) {

	// get parameters
	id := r.URL.Query().Get("id")

	// construct key
	key := fmt.Sprintf("sf:key:%v", id)
	log.Printf("key: %v", key)

	// do external call database using singleflight
	v, err, shared := requestGroup.Do(key, func() (interface{}, error) {
		return DoExternalCallToDatabase()
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// cast singleflight response to basic response
	response := v.(model.BasicResponse)

	// modify response if shared
	response.IsUsingSF = shared

	// write response to http
	responseInByte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	w.Write(responseInByte)

}
