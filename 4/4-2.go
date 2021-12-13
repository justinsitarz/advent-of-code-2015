package main

import (
    "crypto/md5"
    "fmt"
    "strconv"
)

func main() {
	seed := "bgvyzdsv"
	for i := 1; i <= 1000000000; i++ {
	    data := []byte(seed + strconv.Itoa(i))
	    str := fmt.Sprintf("%x", md5.Sum(data))
	    first6 :=  str[0:6]
	    if first6 == "000000" {
	    	fmt.Println("i = ", i)
	    }

	}
}