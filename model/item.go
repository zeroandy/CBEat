package model

type Item struct {
	Name  string `json:"title"`
	Desc  string `json:"desc"`
	Price int    `json:"price"`
}
