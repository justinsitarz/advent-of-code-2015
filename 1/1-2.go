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
	for i, c := range dat {
		if string(c) == "(" {
			counter += 1
		} else if string(c) == ")" {
			counter -= 1
		}

		if counter < 0 {
			// From instructions: 'The first character in the instructions has position 1', so
			// add 1 to index to get the position number
			fmt.Println("Entered the basement at position: ", i + 1)
			break;
		}
	}

	fmt.Println(counter)
	
}



