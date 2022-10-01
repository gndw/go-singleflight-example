package endpoints

import "golang.org/x/sync/singleflight"

var ExternalCallToDatabaseConcurrency int = 0
var requestGroup singleflight.Group
