package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gndw/go-singleflight-example/model"
	"github.com/google/uuid"
)

var counter = 0

func (s *Service) OkWithSingleflight(w http.ResponseWriter, r *http.Request) {

	// construct response
	response := model.BasicResponse{}

	// get parameters
	id := r.URL.Query().Get("id")     // id to grouping request, sent from client
	requestID := uuid.NewString()[:4] // id to differentiate between each request

	// logging
	log.Printf("request %v:%v is coming...", id, requestID)
	defer func() {
		log.Printf("request %v:%v is ending with shared: %v", id, requestID, response.IsUsingSF)
	}()

	// do external call database using singleflight
	data, err, shared := s.requestGroup.Do(id, func() (interface{}, error) {
		return s.externalService.DoExternalCallToDatabase(id)
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// cast singleflight response to basic response
	response = data.(model.BasicResponse)

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
