package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Payload struct {
	Numbers []float64 `json:"numbers"`
}

func QuickSortFloat(arr []float64) []float64 {

	if len(arr) < 2 {
		// fmt.Println("len:", len(arr))
		// fmt.Println("Conteúdo do Array:", arr)
		return arr
	} else {
		// Pego primeiro elemento como pivo para comparar
		pivo := arr[0]
		// fmt.Println("Pivo:", pivo)
		var maiores []float64
		var menores []float64
		for i := 1; i <= (len(arr) - 1); i++ {
			if pivo <= arr[i] {
				maiores = append(maiores, arr[i])
			}
			if arr[i] <= pivo {
				menores = append(menores, arr[i])
			}
		}

		var result []float64
		result = append(result, QuickSortFloat(menores)...)
		result = append(result, pivo)
		result = append(result, QuickSortFloat(maiores)...)
		return result
	}
}

func requetUrl(url string) []float64 {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	payload := Payload{}
	jsonErr := json.Unmarshal(body, &payload)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return payload.Numbers
}

func RecursiveMethod(array []float64, index int) []float64 {
	url := "http://challenge.dienekes.com.br/api/numbers?page="
	url = fmt.Sprintf("%s%d", url,
		index)
	fmt.Println("uRL:", url)
	res := requetUrl(url)
	if len(res) != 0 {
		array = append(array, res...)
		return RecursiveMethod(array, index+1)
	}
	fmt.Println(res)
	return array

}

func main() {
	var array []float64
	var i int
	for i = 1; i <= 1000; i++ {
		url := "http://challenge.dienekes.com.br/api/numbers?page="
		url = fmt.Sprintf("%s%d", url, i)
		fmt.Println(url)
		res := requetUrl(url)
		fmt.Println(res)
		if len(res) == 0 {
			println(len(res))
			println("Erro ao pegar os dados-------------------------:", i)
			fmt.Println(res)
			return
		} else {
			array = append(array, res...)

		}
	}

	// }
	// array = RecursiveMethod(array, 1)
	fmt.Println(array)

}
