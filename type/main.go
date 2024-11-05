package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"year"`
	Price         int    `json:"price"`
	Copies        int    `json:"copies"`
}

func main() {

	myBooks := [1]Book{
		{
			Title:         "Rise",
			Author:        "Pawan",
			PublishedYear: 2023,
			Price:         99,
			Copies:        700,
		},
	}
	v, err := json.Marshal(myBooks)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(v))
}
