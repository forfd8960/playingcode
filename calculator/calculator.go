package calculator

import "fmt"

// Calculator ...
type Calculator struct {
	s      string
	result interface{}
}

// Exec run the experssion
func (cal *Calculator) Exec() error {
	return nil
}

// Result get the result of the calculation
func (cal *Calculator) Result() interface{} {
	return cal.result
}
