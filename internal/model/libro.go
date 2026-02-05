package model

type Libro struct {
	ID     int    `json:"id"`
	Titulo string `json:"tittle"`
	Autor  string `json:"autor"`
}
