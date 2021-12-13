package main

import (
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	counter := 0

	dat, err := os.ReadFile("input.txt")
	check(err)
	for _, c := range dat {
		if string(c) == "(" {
			counter += 1
		} else if string(c) == ")" {
			counter -= 1
		}
	}

	fmt.Println(counter)
	
}



