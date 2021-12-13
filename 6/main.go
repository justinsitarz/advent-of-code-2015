package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
	)

type Vertex struct {
    X, Y int
}

func main() {

	str := "turn on 489,959 through 759,964"
	checkBool, _ := regexp.Compile(`turn (\w*) `)
	ver1, _ := regexp.Compile(` (\w*,\w*) through`)
	ver2, _ := regexp.Compile(`through (\w*,\w*)`)
	a := checkBool.FindStringSubmatch(str)
	b := ver1.FindStringSubmatch(str)
	c := ver2.FindStringSubmatch(str)
	fmt.Println(a[1])
	v1a, _ := strconv.Atoi(strings.Split(b[1], ",")[0])
	v1b, _ := strconv.Atoi(strings.Split(b[1], ",")[1])
	v1 := &Vertex{v1a, v1b}
	fmt.Println(v1.X, v1.Y)
	fmt.Println(c[1])
}

