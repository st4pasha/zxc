package main

import (
	"errors"
	"fmt"
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
	Products, err := Products("1,молоко,150,Простоквашино")
	if err != nil {
		fmt.Println("Error !!! -", err)
		return
	}

	sliceProducts = append(sliceProducts, Products)
	fmt.Println(sliceProducts)
}

func Products(str string) (Product, error) {

	sliceStr := strings.Split(str, ",")
	if len(sliceStr) < 4 {
		return Product{}, errors.New("Недостаточно данных")
	}

	id, err := strconv.Atoi(sliceStr[0])
	if err != nil {
		return Product{}, err
	}

	price, err := strconv.ParseFloat(sliceStr[2], 64)
	if err != nil {
		return Product{}, err
	}

	product, err := CreateProduct(id, sliceStr[1], price, sliceStr[3])

	if err != nil {
		fmt.Println("Произошла ошибка при инициализации продукта.")
		return Product{}, err
	}

	return product, nil
}

func CreateProduct(id int, name string, price float64, description string) (Product, error) {
	if id == 0 {
		return Product{}, errors.New("Неправильно переданный id")
	}

	if name == "" {
		return Product{}, errors.New("Неправильно переданное имя")
	}

	if price <= 0 {
		return Product{}, errors.New("Неправильно переданная цена за товар.")
	}

	if description == "" {
		return Product{}, errors.New("Неправильно переданое описание товара.")
	}

	return Product{id, name, price, description}, nil
}
