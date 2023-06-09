package main

import (
	"bytes"
	"fmt"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	//data yang dimasukan
	data := []byte(`{
		"UserID" : 1,
		"Title" : "Ini title",
		"Body" : "Ini Body"
	}`)

	response, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(response)
}
