package main

import (
	"encoding/json"
	"fmt"
)

// `json:"userId"` is called by tags
type Informacao struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Slice of bytes is equal to simple json
	sliceOfBytes := []byte(`{"userId": 1,"id": 1,"title": "delectus aut autem","completed": false}`)
	fmt.Println(sliceOfBytes)
	var info Informacao
	// We pass the memory address as second parameter
	err := json.Unmarshal(sliceOfBytes, &info)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

}

// {
// 	"userId": 1,
// 	"id": 1,
// 	"title": "delectus aut autem",
// 	"completed": false
//   }
