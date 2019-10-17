$(eval ROOT_DIRECTORY := $(shell pwd))

test:
	cd ${ROOT_DIRECTORY}/tool/isbn-book-data-maker && go test ./... -covermode=count -coverprofile=profile.cov
	cd ${ROOT_DIRECTORY}/bookshelf && go test ./... -covermode=count -coverprofile=profile.cov
	cd ${ROOT_DIRECTORY}/bookpub && go test ./... -covermode=count -coverprofile=profile.cov
	cd ${ROOT_DIRECTORY}/booksub && go test ./... -covermode=count -coverprofile=profile.cov
	cd ${ROOT_DIRECTORY}/bookranking && go test ./... -covermode=count -coverprofile=profile.cov
	cd ${ROOT_DIRECTORY}/booksearch && go test ./... -covermode=count -coverprofile=profile.cov
	cd ${ROOT_DIRECTORY}/searchpub && go test ./... -covermode=count -coverprofile=profile.cov
	cd ${ROOT_DIRECTORY}/searchsub && go test ./... -covermode=count -coverprofile=profile.cov
