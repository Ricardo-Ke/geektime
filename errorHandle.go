package main

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return "", err
	}
	response = fmt.Sprintf("request: %s", request)
	return response, nil
}

func main() {
	for _, req := range []string{"", "hello!"} {
		res, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		} else {
			fmt.Printf("response: %s\n", res)
		}

	}
}
