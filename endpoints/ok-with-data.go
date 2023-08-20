package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (s *Service) OkWithData(w http.ResponseWriter, r *http.Request) {

	// get parameters
	id := r.URL.Query().Get("id")    // id to grouping request, sent from client
	uniqueID := uuid.NewString()[:4] // id to differentiate between each request

	// logging
	log.Printf("request %v:%v is coming...", id, uniqueID)
	defer func() {
		log.Printf("request %v:%v is ending.", id, uniqueID)
	}()

	// simulating external call to database
	response, err := s.externalService.DoExternalCallToDatabase(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// construct response
	responseInByte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	w.Write(responseInByte)
}
