package models

//
type Empty struct{}

//
type Set struct {
	S map[string]Empty `json:"hostname"`
}
