package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func print_out_type(x interface{}) string {
	if reflect.TypeOf(x).String() == "int64" {
		x = strconv.FormatInt(x, 10)
		fmt.Println("O retorno int")
	}
	if reflect.TypeOf(x).String() == "string" {
		fmt.Println("O retorno string")
	}
	return reflect.TypeOf(x).String()
}

func main() {
	var v int64 = 23
	fmt.Println(print_out_type(v))
	fmt.Println(print_out_type("foo"))
}
