package importer

import (
	"io/ioutil"
	"goworkshop/model"
	"fmt"
	"encoding/json"
)

func ImportAuthors() map[string]model.AuthorDto{
	fileContent, err := ioutil.ReadFile("importer/authors.json")
	if (err != nil) {
		fmt.Println("Unable to open the file")
		panic(err)
	}
	var authors map[string]model.AuthorDto
	authors = make(map[string]model.AuthorDto)
	var list []model.AuthorDto

	err = json.Unmarshal(fileContent, &list)
	for _, author := range list {
		authors[author.UUID] = author
	}
	return authors
 }

func ImportBooks() map[string]model.BookDto{
	fileContent, err := ioutil.ReadFile("importer/books.json")
	if (err != nil) {
		fmt.Println("Unable to open the file")
		panic(err)
	}

	var books []model.BookDto
	var booksMap map[string]model.BookDto

	booksMap = make(map[string]model.BookDto)

	for _, book := range books {
		booksMap[book.UUID] = book
	}

	err = json.Unmarshal(fileContent, &books)
	return booksMap
}