package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "12d134"
	// CREATE A HASH BASED A STRING
	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))
	// Decrypt the hash to compare with password
	result := bcrypt.CompareHashAndPassword(sb, []byte("12d134"))
	fmt.Println(result)
}
