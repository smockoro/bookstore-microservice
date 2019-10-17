$(eval ROOT_DIRECTORY := $(shell pwd))

test:
	cd ${ROOT_DIRECTORY}/tool/isbn-book-data-maker && go test ./...
	cd ${ROOT_DIRECTORY}/bookshelf && go test ./...
	cd ${ROOT_DIRECTORY}/bookpub && go test ./...
	cd ${ROOT_DIRECTORY}/booksub && go test ./...
	cd ${ROOT_DIRECTORY}/bookranking && go test ./...
	cd ${ROOT_DIRECTORY}/booksearch && go test ./...
	cd ${ROOT_DIRECTORY}/searchpub && go test ./...
	cd ${ROOT_DIRECTORY}/searchsub && go test ./...