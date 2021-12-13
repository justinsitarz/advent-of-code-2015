package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
    "sort"
)


func main() {

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    totalRibbonLength := 0


    // Each line is dims in LxWxH
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "x")
		sNums := make([]int, len(s))
		for i, x := range s {
    		sNums[i], err = strconv.Atoi(x)
		}
		ribbonLength := calcRibbonLength(sNums)
		totalRibbonLength += ribbonLength
	}

	fmt.Println(totalRibbonLength)
	
}

func calcRibbonLength(s []int) int {
	sort.Ints(s)
	length := 2*s[0] + 2*s[1]
	bow := s[0] * s[1] * s[2]
	totalLength := length + bow
	return totalLength

}

