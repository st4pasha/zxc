package main

import (
	"fmt"
	"gostudy/errors"
	"strconv"
	"strings"

	stderrors "github.com/pkg/errors"
)

type Product struct {
	id          int     // айди
	name        string  // тип товара
	price       float64 // цена товара
	description string  // название товара
}

func main() {
	products, err := parseProduct("1,молоко,150,Простоквашино")
	if err != nil {
		fmt.Println("Error !!! -", err)
		return
	}

	fmt.Println(products)
}

func parseProduct(str string) (Product, error) {

	sliceStr := strings.Split(str, ",")
	if len(sliceStr) < 4 {
		return Product{}, errors.ErrMissingData
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
		return stderrors.Wrap(errors.ErrParse, "id zero")
	}

	if name == "" {
		return stderrors.Wrap(errors.ErrParse, "invalid name")
	}

	if price <= 0 {
		return stderrors.Wrap(errors.ErrParse, "incorrect price")
	}
	if description == "" {
		return stderrors.Wrap(errors.ErrParse, "invalid description")
	}

	return nil
}
