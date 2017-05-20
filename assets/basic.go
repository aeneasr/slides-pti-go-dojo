package main

import (
	"fmt"
	"errors"
)

func main() {
	var foo int
	foo = 5
	fmt.Printf("foo=%d / 5\n", foo * 5)

	keeper, err := bar(false)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	keeper, _ = bar(true)
	fmt.Printf("bar%s\n", keeper)
}

func bar(ok bool) (string, error) {
	if !ok {
		return "", errors.New("an error occurred")
	}

	return "keeper", nil
}
