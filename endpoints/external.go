package endpoints

import (
	"log"
	"time"

	"github.com/gndw/go-singleflight-example/model"
)

func DoExternalCallToDatabase() (model.BasicResponse, error) {

	// begin
	log.Println("external call to database...")
	timeStart := time.Now()
	ExternalCallToDatabaseConcurrency++
	additionalWait := ExternalCallToDatabaseConcurrency

	// defer ending
	defer func() {
		ExternalCallToDatabaseConcurrency--
		log.Printf("external call to database ending in %v ms", time.Since(timeStart).Milliseconds())
	}()

	// simulating external call
	time.Sleep(1 * time.Second)
	// simulating additional wait time due to concurrency
	time.Sleep(time.Duration(additionalWait*100) * time.Millisecond)

	return model.BasicResponse{
		Status:  "OK",
		Message: "Simulated data from database",
	}, nil

}

func DoExternalCallToDatabaseWithPointer() (*model.BasicResponse, error) {

	log.Printf("external call to database...")
	time.Sleep(1 * time.Second)

	return &model.BasicResponse{
		Status:  "OK",
		Message: "Simulated data from database",
		Value:   1,
	}, nil

}
