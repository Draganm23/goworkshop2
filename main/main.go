package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"goworkshop/importer"
	"goworkshop/model"
)

type Animal struct {

	//ca sa ignore `jsno: "-"`
	NoOfLegs int `json:"noOfLegs"`
	Name string `json:"name"`
}

func (a Animal) String() string {
	return fmt.Sprintf("Animal {noOfLegs:%d, name:%s}", a.NoOfLegs, a.Name)
}

func (a Animal) CanTalk() bool {
	return false
}

type Talker interface {
	CanTalk() bool
}

//pot sa ignor ceea ce nu folosesc si pun in locul acelui lucru "_"

func main() {
	var creature Talker
	creature = Animal {
		NoOfLegs:4,
		Name:"John",
	}
	fmt.Println(creature)

	//cautarea se face din path directory
	fileContent, err := ioutil.ReadFile("main/animal.json")
	if (err != nil) {
		fmt.Println("Unable to open the file")
		panic(err)
	}

	fmt.Println(string(fileContent))
	var animals []Animal
	//deserializeaza ce e in fileContent si fa un obiect si memoreaza-l in animal
	err = json.Unmarshal(fileContent,&animals)

	if(err != nil) {
		fmt.Println("cannot deserialize Animal")
		panic(err)
	}

	//check the values
	fmt.Println(animals)

	if serializedAnimals, err := json.Marshal(animals); err != nil {
		fmt.Println("unable to serialize")
		panic(err)
	} else {

		fmt.Println(string(serializedAnimals))
	}

	importer.ImportAuthors()
	importer.ImportBooks()

	fmt.Println("Books:")
	fmt.Println(model.Books)
	fmt.Println("Authors")
	fmt.Println(model.Authors)
}
