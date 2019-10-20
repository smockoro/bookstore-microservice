package main

import (
	"context"
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/adapter/api"
	"github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/adapter/persistence/repository"
	"github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/app/usecase"
)

func main() {
	c := &http.Client{}
	client := api.NewGoogleBookAPIClient(c)
	file, err := os.OpenFile("./data.csv", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	err = file.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)

	csvRepo := repository.NewCSVRepository("./", writer)

	app := usecase.NewBookUsecase(csvRepo, client)
	ctx := context.Background()
	err = app.Make(ctx)

	//data, err := client.Get("9784167741013")
	if err != nil {
		log.Fatal(err)
	}

}
