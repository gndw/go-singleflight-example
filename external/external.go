package external

import (
	"log"
	"sync"
	"time"

	"github.com/gndw/go-singleflight-example/model"
)

type Service struct {
	ConcurrencyCount                           int
	MapOfExternalCallInMillisecondsByRequestID map[string][]int64
	Mutex                                      *sync.Mutex
}

func New(mutex *sync.Mutex) *Service {
	return &Service{
		ConcurrencyCount: 0,
		MapOfExternalCallInMillisecondsByRequestID: map[string][]int64{},
		Mutex: mutex,
	}
}

func (s *Service) DoExternalCallToDatabase(requestID string) (model.BasicResponse, error) {

	// begin
	log.Println("external call to database...")
	timeStart := time.Now()

	// increase simulated concurrency
	s.ConcurrencyCount++

	// defer ending
	defer func() {
		// reduce simulated concurrency
		s.ConcurrencyCount--
		// calculating processing time & storing the data
		processingTimeInMilliseconds := time.Since(timeStart).Milliseconds()
		s.Mutex.Lock()
		s.MapOfExternalCallInMillisecondsByRequestID[requestID] = append(s.MapOfExternalCallInMillisecondsByRequestID[requestID], processingTimeInMilliseconds)
		s.Mutex.Unlock()
		log.Printf("external call to database ending in %v ms", processingTimeInMilliseconds)
	}()

	// simulating external call
	time.Sleep(1 * time.Second)
	// simulating additional wait time due to concurrency
	time.Sleep(time.Duration(s.ConcurrencyCount*100) * time.Millisecond)

	return model.BasicResponse{
		Status:  "OK",
		Message: "Simulated data from database",
	}, nil

}

func (s *Service) DoExternalCallToDatabaseWithPointerResponse() (*model.BasicResponse, error) {

	log.Printf("external call to database...")
	time.Sleep(1 * time.Second)

	return &model.BasicResponse{
		Status:  "OK",
		Message: "Simulated data from database",
		Value:   1,
	}, nil

}
