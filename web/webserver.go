package web

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"io/ioutil"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

type Route struct {
	route      string
	handler    func(http.ResponseWriter, *http.Request)
	httpMethod string
}

var routes = []Route{
	{
		route:      "/books",
		handler:    getAllBooks,
		httpMethod: "GET",
	},
	{
		route:      "/books/{uuid}",
		handler:    getBookByUuid,
		httpMethod: "GET",
	},
	{
		route:      "/authors",
		handler:    getAllAuthors,
		httpMethod: "GET",
	},
	{
		route:      "/authors/create",
		handler:    createAuthor,
		httpMethod: "POST",
	},
	{
		route:      "/authors/update/{uuid}",
		handler:    updateAuthor,
		httpMethod: "POST",
	},
	{
		route:      "/authors/delete/{uuid}",
		handler:    deleteAuthor,
		httpMethod: "GET",
	},
}

func StartServer() {
	mux := mux.NewRouter()
	for _, route := range routes {
		mux.HandleFunc(route.route, route.handler).Methods(route.httpMethod)
	}

	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}

func serializeData(data interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	if data, err := json.Marshal(data); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading!\"}")
		return err
	} else {
		fmt.Fprintln(w, string(data))
		return nil
	}
}

func getAllAuthors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	if err := serializeData(model.Authors, w); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading body!\"}")
		return
	}
	var author model.Author
	if err := json.Unmarshal(body, &author); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error unmarshling the body!\"}")
		return
	}
	erra := model.Authors[author.UUID]
	if erra == (model.Author{}) {
		model.Authors[author.UUID] = author
		return
	}


	fmt.Fprintln(w, "{\"message\":\"The author already exists with this UUID!\"}")
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading body!\"}")
		return
	}
	var authorUpdated model.Author
	if err := json.Unmarshal(body, &authorUpdated); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error unmarshling the body!\"}")
		return
	}
	uuid := mux.Vars(r)["uuid"]
	erra := model.Authors[uuid]

	if erra == (model.Author{}) {

		fmt.Fprintln(w, "{\"message\":\"The author does not exist!\"}")
		w.WriteHeader(http.StatusNotFound)
		return

	}
	delete(model.Authors, uuid)
	model.Authors[erra.UUID] = authorUpdated
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	w.Header().Set("Content-Type", "application/json")
	erra := model.Authors[uuid]
	if erra.UUID == "" {
		fmt.Fprintln(w, "{\"message\":\"The author does not exist!\"}")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(model.Authors,uuid)
}

func getBookAuthor(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	for _, book := range model.Books {
		if book.UUID == uuid {
			if err := serializeData(book.Author, w); err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}

	fmt.Fprintln(w, "{\"message\":\"The book does not exist!\"}")
	w.WriteHeader(http.StatusNotFound)
}

func getBookByUuid(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	w.Header().Set("Content-Type", "application/json")

	for _, book := range model.Books {
		if book.UUID == uuid {
			if data, err := json.Marshal(book); err != nil {
				fmt.Fprintln(w, "{\"message\":\"Error reading!\"}")
				return
			} else {
				fmt.Fprintln(w, string(data))
				return
			}
		}
	}

	fmt.Fprintln(w, "{\"message\":\"The book does not exist!\"}")
	w.WriteHeader(http.StatusNotFound)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	w.Header().Set("Content-Type", "application/json")

	if data, err := json.Marshal(model.Books); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading!\"}")
	} else {
		fmt.Fprintln(w, string(data))
	}

}

func log(funcHandler http.HandlerFunc) http.HandlerFunc{
	fmt.Println("Returning the function")
	return func (rw http.ResponseWriter, r *http.Request){
		fmt.Println("New REST request to url: "+r.URL.Path)
		funcHandler(rw, r)
		fmt.Println("Rest request ended")
	}
}

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}