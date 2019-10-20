package model

type Book struct {
	ISBN          string
	Title         string
	Author        string
	PublishedDate string
	Description   string
	PageCount     int64
	AverageRating float64
	RatingsCount  int64
}
