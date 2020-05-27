package model

type Item struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Owner  string `json:"owner"`
	Rating int    `json:"rating"`
}
