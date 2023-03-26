package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Posts []struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var obj Posts
	json.Unmarshal(responseData, &obj)

	for _, item := range obj {
		fmt.Println("======")
		fmt.Println("userID = ", item.UserID)
		fmt.Println("ID = ", item.ID)
		fmt.Println("Title = ", item.Title)
		fmt.Println("Body = ", item.Body)
		fmt.Println("======")
	}

}
