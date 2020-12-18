package main

import (
	"fmt"
)

type ErrDivideByZero float32

func (e ErrDivideByZero) Error() string {
	return fmt.Sprintf("Cannot divide by zero: divisor = %v", float32(e))
}
func divider(x float32, y float32) (float32, error) {
	if y == 0 {
		return float32(0), ErrDivideByZero(0)
	}
	return x / y, nil
}

func main() {
	z, err := divider(10, 0)
	fmt.Printf("%v, %v", z, err)
}
