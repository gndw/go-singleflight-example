package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gndw/go-singleflight-example/model"
)

func OkWithSingleflightPointer(w http.ResponseWriter, r *http.Request) {

	// do external call database using singleflight
	v, err, shared := requestGroup.Do("key", func() (interface{}, error) {
		return DoExternalCallToDatabaseWithPointer()
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// cast singleflight response to basic response
	response := v.(*model.BasicResponse)

	// modify response if shared
	response.IsUsingSF = shared

	// do mutation to response (pointer)
	response.Value++

	// write response to http
	responseInByte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	w.Write(responseInByte)

}
