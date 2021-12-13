package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "math"
)


func main() {

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		pairs := checkForPairs(s)
		repeat 	:= checkForRepeat(s)

		if (pairs == true && repeat == true) {
			count ++
		}
	}
	fmt.Println(count)
	
}

func checkForPairs(s string) bool {
	pairs := map[int]string{}
	for i := 0; i < 15; i ++ {
		pair := string(s[i]) + string(s[i+1])
		pairs[i] = pair
	}
	for key1, value1 := range pairs {
		for key2, value2 := range pairs {
			if (value1 == value2 && math.Abs(float64(key1 - key2)) > 1) {
				fmt.Println(value1, key1, value2, key2)
				return true
			}
		}
	}
	return false
}
// contains at least one letter which repeats with exactly one letter between them, like 'xyx'
func checkForRepeat(s string) bool {
	var pairs [14]string
	for i := 0; i < 14; i ++ {
		pair := string(s[i]) + string(s[i+1]) + string(s[i+2])
		pairs[i] = pair
	}
	for _, element := range pairs {
		if (element[0] == element[2]) {
			return true
		}
	}
	return false

}

// func checkForVowels(s string) bool  {
// 	count := 0
// 	for _, char := range s {
// 		switch string(char) {
// 		case
// 			"a","e","i","o","u":
// 			count += 1
// 		}
// 	}
// 	if count > 2 {
// 		return true
// 	}
// 	return false
// }
	

// func checkDoubleLetters(s string) bool {
// 	last := ""
// 	for _, char := range s {
// 		if string(char) == last {
// 			return true
// 		}
// 		last = string(char)
// 	}
// 	return false
// }

// func checkForbiddenStrings(s string) bool {
// 	// forbiddenStrings := ["ab", "cd", "pq", "xy"]
// 	regex := regexp.MustCompile(`.*ab|cd|pq|xy.*`)
// 	val := regex.MatchString(s)
// 	return val
// }


