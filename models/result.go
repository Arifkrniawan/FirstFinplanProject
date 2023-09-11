package models

type Result struct {
	Id     uint8   `json:"id"`
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Status string  `json:"status"`
	Hobby  []Hobby `json:"hobbies"`
}
