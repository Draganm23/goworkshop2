package model

import "fmt"

//Book - The DTO used to access books
type Book struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	NoPages     int    `json:"noPages"`
	ReleaseDate string `json:"releaseDate"`
	Author      Author `json:"author"`
}

//Author - The DTO used to access authors
type Author struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

type AuthorsSlice []Author

func (p *AuthorsSlice) Remove(element Author) error {
	var err = fmt.Errorf("could not find element %s", element)
	var updatedSlice AuthorsSlice
	for _, value := range *p {
		if value == element {
			err = nil
		} else {
			updatedSlice = append(updatedSlice, value)
		}
	}
	if err == nil {
		*p = updatedSlice
	}
	return err
}


//Books - the list of available books
var Books map[string]Book

// Authors - the list of available authors
var Authors map[string]Author
