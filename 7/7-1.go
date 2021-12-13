package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "regexp"
)

var wires = make(map[string]int)
var instructions []string

func main() {
    // Read file, each line set as element in 'instructions'
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
        instructions = append(instructions, s)
    }
    
    for ind, val := range instructions {
        checkOperator(ind, val)

    }

        
}

func add(num ...int) int {
    sum := 0
    for j := range num {
        sum += j
    }
    return sum
}
  
func main() {
    fmt.Println("Sum =", add(1, 2, 3, 4, 5, 7, 8, 6, 5, 4))
}

// func parseInstructions(str string) {
//     argsMatch,_ := regexp.Compile(`(\w*) ->`)
//     args := argsMatch.FindStringSubmatch(str)[1]

// }


func checkOperator(ind int, val string) {
    notMatch, _     := regexp.Compile(`NOT [a-z]*`)
    orMatch, _      := regexp.Compile(`[a-z]* OR [a-z]*`)
    andMatch, _     := regexp.Compile(`[a-z]* AND [a-z]*`)
    lShiftMatch, _  := regexp.Compile(`[a-z]* LSHIFT`)
    rShiftMatch, _  := regexp.Compile(`[a-z]* RSHIFT`)
    valMatch, _     := regexp.Compile(`\d.* ->*`)
    
    switch {
        case notMatch.MatchString(val): 
            not(ind, val)
        case orMatch.MatchString(val): 
        case andMatch.MatchString(val): 
        case lShiftMatch.MatchString(val): 
        case rShiftMatch.MatchString(val): 
        case valMatch.MatchString(val):
        }
}

func not(val string) {
    target_regex, _ := regexp.Compile(`-> (\w*)`)
    source_regex, _ := regexp.Compile(`NOT (\w*) ->`)
    targetWire := target_regex.FindStringSubmatch(val)[1]
    sourceWire := source_regex.FindStringSubmatch(val)[1]
    fmt.Println(sourceWire, targetWire)

    if sourceWire, ok := wires[sourceWire]; ok {
        wires[targetWire] = ^sourceWire
    }
    
}

// func or()

// func and()

// func lShift()

// func rShift()

// func setVal()




