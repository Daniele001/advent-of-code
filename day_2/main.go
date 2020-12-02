package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


// This function find the number of correct passwords of part 1
func correctPasswords(data []string, c *int) {
	rang := strings.Split(data[0], "-")
	max, _ := strconv.Atoi(rang[1])
	min, _ := strconv.Atoi(rang[0])
	occurrences := strings.Count(data[2], string(data[1][0]))
	if occurrences >= min && occurrences <= max {
		*c++
	}
}

// This function find the number of correct passwords of part 2
func correctPasswords2(data []string, c *int) {
	rang := strings.Split(data[0], "-")
	pos1, _ := strconv.Atoi(rang[1])
	pos2, _ := strconv.Atoi(rang[0])
	charInPass1 := string(data[2][pos1-1])
	charInPass2 := string(data[2][pos2-1])
	charToFind := string(data[1][0])
	if (charInPass1 == charToFind || charInPass2 == charToFind) && !(charInPass1 == charToFind && charInPass2 == charToFind) {
		*c++
	}
}

func main() {
	buffer := bufio.NewScanner(os.Stdin)
	count := 0
	count2 := 0
	for buffer.Scan() {
		data := strings.Split(buffer.Text(), " ")
		correctPasswords(data, &count)
		correctPasswords2(data, &count2)
	}
	fmt.Println(count)
	fmt.Println(count2)
}
