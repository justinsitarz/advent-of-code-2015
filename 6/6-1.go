package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "regexp"
    "strconv"
    "strings"
    "math"

)

type Vertex struct {
    a, b, x, y int
    value string
}


func main() {

    var matrix [1000][1000]int
    value := 0
    count := 0

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
        v := parseInstructions(s)

        if v.value == "on" || v.value == "off" {
            if v.value == "off" {
                value = 0
            } else {
                value = 1
            }

            for i := v.a; i <= v.x; i ++ {
                for j := v.b; j <= v.y; j ++ {
                    matrix[i][j] = value
                }
            }
        } else {
            for i := v.a; i <= v.x; i ++ {
                for j := v.b; j <= v.y; j ++ {
                    tmp := math.Abs(float64(matrix[i][j] - 1))
                    matrix[i][j] = int(tmp)
                }
            }
        }
    }
    for i := 0; i < 1000; i ++ {
        for _, element := range matrix[i] {
            if element == 1 {
                count ++
            }
        }		
    }
    fmt.Println(count)

}

func parseInstructions(str string) (Vertex) {
    checkBool, _ := regexp.Compile(`(\w*) \d`)
    ver1, _ := regexp.Compile(` (\w*,\w*) through`)
    ver2, _ := regexp.Compile(`through (\w*,\w*)`)

    a := checkBool.FindStringSubmatch(str)[1]
    b := ver1.FindStringSubmatch(str)
    c := ver2.FindStringSubmatch(str)
    fmt.Println(a)
    

    v1a, _ := strconv.Atoi(strings.Split(b[1], ",")[0])
    v1b, _ := strconv.Atoi(strings.Split(b[1], ",")[1])
    v2a, _ := strconv.Atoi(strings.Split(c[1], ",")[0])
    v2b, _ := strconv.Atoi(strings.Split(c[1], ",")[1])

    v := Vertex{v1a, v1b, v2a, v2b, a}
    fmt.Println(v)

    return v

}


	
// }
// // example: turn on 489,959 through 759,964
// func parseLine(s string) (Vertex, Vertex, bool) {
//     var (
//         v1 = Vertex{}
//         v2 = Vertex{}
//     )


// }




