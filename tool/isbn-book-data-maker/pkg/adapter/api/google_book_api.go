package api

import (
	"encoding/json"
	"net/http"

	Iapi "github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/domain/external/api"
	"github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/domain/model"
)

// NewGoogleBookAPIClient : Constractor Injector for Google Book API
func NewGoogleBookAPIClient(client *http.Client) Iapi.GoogleBookAPIClient {
	return &googleBookAPIClient{
		client:   client,
		endpoint: "https://www.googleapis.com/books/v1/volumes",
	}
}

type googleBookAPIClient struct {
	client   *http.Client
	endpoint string
}

func (gbac *googleBookAPIClient) Get(isbn string) (*model.Book, error) {
	query := "q=isbn:"
	requestURL := gbac.endpoint + "?" + query + isbn
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := gbac.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := &Iapi.ResponseJSON{}
	err = decodeBody(res, data)
	if err != nil {
		return nil, err
	}

	if data.TotalItems == 1 {
		return data.Items[0], nil
	}
	return nil, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
