package models

//
type Message struct {
	Message string  `json:"message"`
	Time    string  `json:"time"`
	CPU     float64 `json:"cpu"`
	Mem     float64 `json:"mem"`
}
