package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Product struct {
	id          int
	name        string
	price       int // usd
	description string
}

func main() {
	fmt.Println(returnProduct("1,молоко,100,Простоквашино"))
}

func returnProduct(str string) []Product {

	sliceStr := strings.Split(str, ",")
	id, _ := strconv.Atoi(sliceStr[0])
	price, _ := strconv.Atoi(sliceStr[2])

	product := Product{id, sliceStr[1], price, sliceStr[3]}

	return []Product{product}
}
