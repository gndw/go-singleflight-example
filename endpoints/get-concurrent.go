package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gndw/go-singleflight-example/model"
)

func (s *Service) GetConcurrent(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	data := s.externalService.MapOfExternalCallInMillisecondsByRequestID[id]

	response := model.BasicConcurrentDataResponse{ID: id}
	response.Notes = append(response.Notes, fmt.Sprintf("number of database call: %v", len(data)))
	maxTime := int64(0)
	totalTime := int64(0)
	for _, d := range data {
		response.Details = append(response.Details, fmt.Sprintf("external request with %v ms", d))
		totalTime += d
		if d > int64(maxTime) {
			maxTime = d
		}
	}
	response.Notes = append(response.Notes, fmt.Sprintf("maximum processing time: %v ms", maxTime))
	response.Notes = append(response.Notes, fmt.Sprintf("average processing time: %v ms", totalTime/int64(len(data))))

	// construct response
	responseInByte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	w.Write(responseInByte)

}
