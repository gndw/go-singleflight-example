package endpoints

import (
	"github.com/gndw/go-singleflight-example/external"
	"golang.org/x/sync/singleflight"
)

type Service struct {
	requestGroup    *singleflight.Group
	externalService *external.Service
}

func New(requestGroup *singleflight.Group, externalService *external.Service) *Service {
	return &Service{
		requestGroup:    requestGroup,
		externalService: externalService,
	}
}
