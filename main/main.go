package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Animal struct {
	NoOfLegs int
	Name string
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
}
