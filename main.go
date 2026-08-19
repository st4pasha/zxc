package main

import (
	"fmt"
	errors1 "gostudy/errors"
	"strconv"
	"strings"
)

type Product struct {
	id          int     // айди
	name        string  // тип товара
	price       float64 // цена товара
	description string  // название товара
}

var sliceProducts []Product

func main() {
	products, err := parseProduct("1,молоко,150,Простоквашино")
	if err != nil {
		fmt.Println("Error !!! -", err)
		return
	}
	sliceProducts = append(sliceProducts, products)
	fmt.Println(sliceProducts)
}

func parseProduct(str string) (Product, error) {

	sliceStr := strings.Split(str, ",")
	if len(sliceStr) < 4 {
		return Product{}, errors1.ErrMissData
	}

	id, err := strconv.Atoi(sliceStr[0])
	if err != nil {
		return Product{}, err
	}

	price, err := strconv.ParseFloat(sliceStr[2], 64)
	if err != nil {
		return Product{}, err
	}

	if err := validProduct(id, sliceStr[1], price, sliceStr[3]); err != nil {
		return Product{}, err
	}

	return Product{id, sliceStr[1], price, sliceStr[3]}, nil
}

func validProduct(id int, name string, price float64, description string) error {
	if id == 0 {
		return errors1.ErrNotFound
	}

	if name == "" {
		return errors1.ErrNotFound
	}

	if price <= 0 {
		return errors1.ErrNotFound
	}
	if description == "" {
		return errors1.ErrNotFound
	}

	return nil
}
