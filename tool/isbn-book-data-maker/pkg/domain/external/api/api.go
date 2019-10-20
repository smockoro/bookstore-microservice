package api

import "github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/domain/model"

// GoogleBookAPIClient : HTTP Cient
type GoogleBookAPIClient interface {
	Get(string) (*model.Book, error)
}

// ResponseJSON : Google Book API Response JSON
type ResponseJSON struct {
	Kind       string      `json:"kind"`
	TotalItems int         `json:"totalItems"`
	Items      model.Books `json:"items"`
}
