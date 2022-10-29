package models

type Product struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
