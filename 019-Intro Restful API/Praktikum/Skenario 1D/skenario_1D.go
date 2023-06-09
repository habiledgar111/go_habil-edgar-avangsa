package main

import (
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
	//menghapus id == 3
	response, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/3", nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(response)
}
