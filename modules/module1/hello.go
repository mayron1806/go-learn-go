package module1

import "errors"

func Soma(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("numeros não podem ser zero")
	}
	return a + b, nil
}
