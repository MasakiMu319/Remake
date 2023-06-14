package main

import (
	"errors"
	"fmt"
)

type myError struct {
	error
}

func (e *myError) Error() string {
	return "this is my error"
}

func main() {
	anotherErr := errors.New("this is another error")
	newErr := fmt.Errorf("warp error: %w", anotherErr)
	fmt.Println(newErr)
	fmt.Println(errors.Unwrap(newErr))
	fmt.Println(errors.Is(newErr, newErr))
}
