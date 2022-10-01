package endpoints

import (
	"encoding/json"
	"net/http"
)

func OkWithData(w http.ResponseWriter, r *http.Request) {

	response, err := DoExternalCallToDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	responseInByte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	w.Write(responseInByte)
}
