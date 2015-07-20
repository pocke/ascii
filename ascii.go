package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var flag_reverse bool

func main() {
	f := flag.NewFlagSet("ascii", flag.ExitOnError)
	f.BoolVar(&flag_reverse, "reverse", false, "String to ascii code")
	f.Parse(os.Args[1:])

	for _, v := range f.Args() {
		if !flag_reverse {
			s, err := ASCII(v)
			if err != nil {
				panic(err)
			}
			fmt.Println(s)
		} else {
			for _, ch := range []byte(v) {
				s, err := ASCIIR(string(ch))
				if err != nil {
					panic(err)
				}
				fmt.Println(s)
			}
		}
	}
}

// "65" => "A"
func ASCII(code string) (string, error) {
	i, err := strconv.Atoi(code)
	if err != nil {
		return "", err
	}
	return string(i), nil
}

// "A" => "65"
func ASCIIR(str string) (string, error) {
	if len(str) != 1 {
		return "", fmt.Errorf("")
	}
	return fmt.Sprintf("%d", int([]byte(str)[0])), nil
}
