package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "regexp"

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
		vowels 		  := checkForVowels(s)
		dubLetters    := checkDoubleLetters(s)
		forbiddenStrs := checkForbiddenStrings(s)

		if (vowels == true && dubLetters == true && forbiddenStrs == false) {
			count ++
		}
	}
	fmt.Println(count)	

	
}

func checkForVowels(s string) bool  {
	count := 0
	for _, char := range s {
		switch string(char) {
		case
			"a","e","i","o","u":
			count += 1
		}
	}
	if count > 2 {
		return true
	}
	return false
}
	

func checkDoubleLetters(s string) bool {
	last := ""
	for _, char := range s {
		if string(char) == last {
			return true
		}
		last = string(char)
	}
	return false
}

func checkForbiddenStrings(s string) bool {
	// forbiddenStrings := ["ab", "cd", "pq", "xy"]
	regex := regexp.MustCompile(`.*ab|cd|pq|xy.*`)
	val := regex.MatchString(s)
	return val
}


