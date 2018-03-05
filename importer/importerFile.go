package importer

import (
	"io/ioutil"
	"goworkshop/model"
	"fmt"
	"encoding/json"
)

func ImportAuthors() map[string]model.Author {
	fileContent, err := ioutil.ReadFile("importer/authors.json")
	if (err != nil) {
		fmt.Println("Unable to open the file")
		panic(err)
	}
	var authors map[string]model.Author
	authors = make(map[string]model.Author)
	var list []model.Author

	err = json.Unmarshal(fileContent, &list)
	for _, author := range list {
		authors[author.UUID] = author
	}
	return authors
 }

func ImportBooks() map[string]model.Book {
	fileContent, err := ioutil.ReadFile("importer/books.json")
	if (err != nil) {
		fmt.Println("Unable to open the file")
		panic(err)
	}

	var books []model.Book
	var booksMap map[string]model.Book

	booksMap = make(map[string]model.Book)

	for _, book := range books {
		booksMap[book.UUID] = book
	}

	err = json.Unmarshal(fileContent, &books)
	return booksMap
}