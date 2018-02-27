package importer

import (
	"io/ioutil"
	"goworkshop/model"
	"fmt"
	"encoding/json"
)

func ImportAuthors() {
	fileContent, err := ioutil.ReadFile("importer/authors.json")
	if (err != nil) {
		fmt.Println("Unable to open the file")
		panic(err)
	}

	err = json.Unmarshal(fileContent, &model.Authors)

 }

func ImportBooks() {
	fileContent, err := ioutil.ReadFile("importer/books.json")
	if (err != nil) {
		fmt.Println("Unable to open the file")
		panic(err)
	}

	err = json.Unmarshal(fileContent, &model.Books)

}