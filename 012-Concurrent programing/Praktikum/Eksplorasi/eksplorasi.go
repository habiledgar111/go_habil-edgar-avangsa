package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type object []struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}

func main() {
	ch := make(chan object)
	var wg sync.WaitGroup

	wg.Add(1)
	go getData(ch, &wg)

	wg.Wait()
	// response, err := http.Get("https://fakestoreapi.com/products")

	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	// data, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var obj object
	// json.Unmarshal(data, &obj)
	for _, item := range <-ch {
		fmt.Println("======")
		fmt.Println(item.Title)
		fmt.Println(item.Price)
		fmt.Println(item.Price)
		fmt.Println("======")
	}

}

func getData(ch chan object, wg *sync.WaitGroup) {
	response, err := http.Get("https://fakestoreapi.com/products")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var obj object
	json.Unmarshal(data, &obj)
	ch <- obj
	close(ch)
	wg.Done()

}
