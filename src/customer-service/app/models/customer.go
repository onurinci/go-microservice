package models

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}
