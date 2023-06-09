package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Posts struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	//menampilkan data dengan id == 3
	response, _ := http.Get("https://jsonplaceholder.typicode.com/posts/3")
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var obj Posts
	json.Unmarshal(responseData, &obj)

	// for _, item := range obj {
	fmt.Println("======")
	fmt.Println("userID = ", obj.UserID)
	fmt.Println("ID = ", obj.ID)
	fmt.Println("Title = ", obj.Title)
	fmt.Println("Body = ", obj.Body)
	fmt.Println("======")
	// }

}
