package main

import "fmt"

// Map is a hash table basicly with key value
func main() {
	amigos := map[string]int{
		"Igor": 1231231,
		"João": 423423,
	}
	fmt.Println(amigos)

	amigos["allan"] = 45466
	fmt.Println(amigos)
	// It's possible know if the value 0 returned by the map is the real value or only that not exist
	// Using comma ok idiom to validate

	exist, ok := amigos["haha"]

	if ok {
		fmt.Println("The key exist", exist)
	} else {
		fmt.Println("The key dind't exit")
	}

	for key, _ := range amigos {
		println(key)
	}
	// Delete item from map
	delete(amigos, "allan")

	fmt.Println(amigos)
}
