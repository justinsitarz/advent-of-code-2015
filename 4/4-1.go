package main

import (
    "crypto/md5"
    "fmt"
    "strconv"
)

func main() {
	seed := "bgvyzdsv"
	for i := 1; i <= 10000000; i++ {
	    data := []byte(seed + strconv.Itoa(i))
	    str := fmt.Sprintf("%x", md5.Sum(data))
	    first5 :=  str[0:5]
	    if first5 == "00000" {
	    	fmt.Println("i = ", i)
	    }

	}
}




