package model

type Item struct {
	Title  string `json:"title"`
	Owner  string `json:"owner"`
	Rating int    `json:"rating"`
}
