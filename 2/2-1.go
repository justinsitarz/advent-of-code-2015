package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
)


func main() {

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    totalArea := 0


    // Each line is dims in LxWxH
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "x")
		sNums := make([]int, len(s))
		for i, x := range s {
    		sNums[i], err = strconv.Atoi(x)
		}
		surfaceArea := calcSurfaceArea(sNums)
		totalArea += surfaceArea
	}

	fmt.Println(totalArea)
	
}

func calcSurfaceArea(s []int) int {
	var surfaceAreas [3]int
	surfaceAreas[0] = s[0] * s[1]
	surfaceAreas[1] = s[1] * s[2]
	surfaceAreas[2] = s[0] * s[2]

	smallestSide := surfaceAreas[0]
	for _, element := range surfaceAreas {
		if element < smallestSide {
			smallestSide = element
		}

	}

	area := surfaceAreas[0] + surfaceAreas[1] + surfaceAreas[2]
	totalArea := area * 2 + smallestSide
	return totalArea

}

