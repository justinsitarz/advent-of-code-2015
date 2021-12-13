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

	pos1 := [2]int{0, 0}
	pos2 := [2]int{0, 0}

	set := make(map[[2]int]bool)
	set[pos1] = true

	dat, err := os.ReadFile("input.txt")
	check(err)
	for i, c := range dat {
		if i % 2 == 0 {

			if string(c) == "<" {
				pos1[0] --
			} else if string(c) == "^" {
				pos1[1] ++
			} else if string(c) == ">" {
				pos1[0] ++
			} else if string(c) == "v" {
				pos1[1] --
			}
			set[pos1] = true
		}

		if i % 2 == 1 {
			if string(c) == "<" {
				pos2[0] --
			} else if string(c) == "^" {
				pos2[1] ++
			} else if string(c) == ">" {
				pos2[0] ++
			} else if string(c) == "v" {
				pos2[1] --
			}
			set[pos2] = true
			
		}

	}
	fmt.Println(set)
	fmt.Println(len(set))
	
}

