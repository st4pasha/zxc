package errors1

import "errors"

var (
	ErrNotFound = errors.New("Ресурс не найден !")
	ErrMissData = errors.New("Недостаточно данных !")
)
