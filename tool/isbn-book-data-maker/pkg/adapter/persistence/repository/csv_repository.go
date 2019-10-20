package repository

import (
	"context"
	"encoding/csv"
	"fmt"

	repo "github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/domain/external/repository"
	"github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/domain/model"
)

func NewCSVRepository(path string, writer *csv.Writer) repo.CSVRepository {
	return &csvRepository{
		path:   path,
		writer: writer,
	}
}

type csvRepository struct {
	path   string
	writer *csv.Writer
}

func (cr *csvRepository) Save(ctx context.Context, book *model.Book) error {
	isbn := ""
	for _, v := range book.VolumeInfo.IndustryIdentifiers {
		if v.Type == "ISBN_13" {
			isbn = v.Identifier
		}
	}

	author := ""
	if len(book.VolumeInfo.Authors) > 0 {
		author = book.VolumeInfo.Authors[0]
	}

	record := []string{
		isbn,
		book.VolumeInfo.Title,
		author,
		book.VolumeInfo.PublishedDate,
		book.VolumeInfo.Description,
		fmt.Sprintf("%d", book.VolumeInfo.PageCount),
		fmt.Sprintf("%f", book.VolumeInfo.AverageRating),
		fmt.Sprintf("%d", book.VolumeInfo.RatingsCount),
	}

	fmt.Println(record)
	cr.writer.Write(record)
	cr.writer.Flush()
	return nil
}
