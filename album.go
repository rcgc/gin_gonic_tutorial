package main

// Album model info
// @Description album information
// @Description with album id, album title, arist and price
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}