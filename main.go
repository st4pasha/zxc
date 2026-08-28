package main

import (
	"fmt"
	"gostudy/errors"
	"strconv"
	"strings"
	"sync"

	stderrors "github.com/pkg/errors"
)

type Product struct {
	id          int     // айди
	name        string  // тип товара
	price       float64 // цена товара
	description string  // название товара
}

var mtx = sync.Mutex{}
var productsStrSlice = []string{"1,молоко,150,Простоквашино", "2,вода,100,Дарида"}

func main() {
	sliceProducts, err := toProducts(productsStrSlice)
	if err != nil {
		fmt.Println("error,", err)
		return
	}

	fmt.Println(sliceProducts)
}

func parseProduct(str []string, index int) (Product, error) {

	sliceStr := strings.Split(str[index], ",")
	if len(sliceStr) < 4 {
		return Product{}, stderrors.Wrap(
			errors.ErrMissingData,
			"string must contain at least 4 comma-separated values",
		)

	}

	id, err := strconv.Atoi(sliceStr[0])
	if err != nil {
		return Product{}, stderrors.Wrap(err, "conversion error")
	}

	price, err := strconv.ParseFloat(sliceStr[2], 64)
	if err != nil {
		return Product{}, stderrors.Wrap(err, "conversion error")
	}

	if err := validProduct(id, sliceStr[1], price, sliceStr[3]); err != nil {
		return Product{}, stderrors.Wrap(err, "error with valid")
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

func toProducts(productsStrSlice []string) ([]Product, error) {
	mtx.Lock()
	sliceProduct := []Product{}
	for i := 0; i < len(productsStrSlice); i++ {
		product, err := parseProduct(productsStrSlice, i)
		if err != nil {
			return []Product{}, stderrors.Wrap(err, "error with parse")
		}
		sliceProduct = append(sliceProduct, product)
	}
	mtx.Unlock()

	return sliceProduct, nil
}
