package api

import (
	"encoding/json"
	"net/http"
)

// NewGoogleBookAPIClient : Constractor Injector for Google Book API
func NewGoogleBookAPIClient(client *http.Client) GoogleBookAPIClient {
	return &googleBookAPIClient{
		client:   client,
		endpoint: "https://www.googleapis.com/books/v1/volumes",
	}
}

// GoogleBookAPIClient : HTTP Cient
type GoogleBookAPIClient interface {
	Get(string) (*ResponseJSON, error)
}

type googleBookAPIClient struct {
	client   *http.Client
	endpoint string
}

func (gbac *googleBookAPIClient) Get(isbn string) (*ResponseJSON, error) {
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

	data := &ResponseJSON{}
	err = decodeBody(res, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

// ResponseJSON : Google Book API Response JSON
type ResponseJSON struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []Book `json:"items"`
}

// Book : Book data
type Book struct {
	Kind       string `json:"kind"`
	ID         string `json:"id"`
	Etag       string `json:"etag"`
	SelfLink   string `json:"selfLink"`
	VolumeInfo struct {
		Title               string   `json:"title"`
		Authors             []string `json:"authors"`
		PublishedDate       string   `json:"publishedDate"`
		Description         string   `json:"description"`
		IndustryIdentifiers []struct {
			Type       string `json:"type"`
			Identifier string `json:"identifier"`
		} `json:"industryIdentifiers"`
		ReadingModes struct {
			Text  bool `json:"text"`
			Image bool `json:"image"`
		} `json:"readingModes"`
		PageCount        int      `json:"pageCount"`
		PrintType        string   `json:"printType"`
		Categories       []string `json:"categories"`
		AverageRating    float64  `json:"averageRating"`
		RatingsCount     int      `json:"ratingsCount"`
		MaturityRating   string   `json:"maturityRating"`
		AllowAnonLogging bool     `json:"allowAnonLogging"`
		ContentVersion   string   `json:"contentVersion"`
		ImageLinks       struct {
			SmallThumbnail string `json:"smallThumbnail"`
			Thumbnail      string `json:"thumbnail"`
		} `json:"imageLinks"`
		Language            string `json:"language"`
		PreviewLink         string `json:"previewLink"`
		InfoLink            string `json:"infoLink"`
		CanonicalVolumeLink string `json:"canonicalVolumeLink"`
	} `json:"volumeInfo"`
	SaleInfo struct {
		Country     string `json:"country"`
		Saleability string `json:"saleability"`
		IsEbook     bool   `json:"isEbook"`
	} `json:"saleInfo"`
	AccessInfo struct {
		Country                string `json:"country"`
		Viewability            string `json:"viewability"`
		Embeddable             bool   `json:"embeddable"`
		PublicDomain           bool   `json:"publicDomain"`
		TextToSpeechPermission string `json:"textToSpeechPermission"`
		Epub                   struct {
			IsAvailable bool `json:"isAvailable"`
		} `json:"epub"`
		Pdf struct {
			IsAvailable bool `json:"isAvailable"`
		} `json:"pdf"`
		WebReaderLink       string `json:"webReaderLink"`
		AccessViewStatus    string `json:"accessViewStatus"`
		QuoteSharingAllowed bool   `json:"quoteSharingAllowed"`
	} `json:"accessInfo"`
	SearchInfo struct {
		TextSnippet string `json:"textSnippet"`
	} `json:"searchInfo"`
}
