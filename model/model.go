package model

type BasicResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	IsUsingSF bool   `json:"is_using_sf"`
	Value     int64  `json:"value,omitempty"`
}

type BasicConcurrentDataResponse struct {
	ID      string   `json:"id"`
	Notes   []string `json:"notes"`
	Details []string `json:"details"`
}
