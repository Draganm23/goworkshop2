package model

import "fmt"

//BookDto - The DTO used to access books
type BookDto struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	NoPages     int       `json:"noPages"`
	ReleaseDate string    `json:"releaseDate"`
	Author      AuthorDto `json:"author"`
}

//AuthorDto - The DTO used to access authors
type AuthorDto struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

type AuthorsSlice []AuthorDto

func (p *AuthorsSlice) Remove(element AuthorDto) error {
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
var Books map[string]BookDto

// Authors - the list of available authors
var Authors map[string]AuthorDto
