package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, v := range os.Args[1:] {
		s, err := ASCII(v)
		if err != nil {
			panic(err)
		}
		fmt.Println(s)
	}
}

func ASCII(code string) (string, error) {
	i, err := strconv.Atoi(code)
	if err != nil {
		return "", err
	}
	return string(i), nil
}
