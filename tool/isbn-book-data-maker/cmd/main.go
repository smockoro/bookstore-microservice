package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/adapter/api"
)

func main() {
	c := &http.Client{}
	client := api.NewGoogleBookAPIClient(c)

	data, err := client.Get("9784167741013")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
