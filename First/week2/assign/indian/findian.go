package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsFound(s string) bool {
	isFound := false
	str := strings.ToLower(s)

	if strings.HasSuffix(str, "\n") {
		str = strings.TrimSuffix(str, "\n")
	}

	// if exist, remove trailing Carriage Return (added on Windows)
	if strings.HasSuffix(str, "\r") {
		str = strings.TrimSuffix(str, "\r")
	}

	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		isFound = true
	}

	return isFound
}

func main() {
	fmt.Println("Please enter a string: ")
	stdinReader := bufio.NewReader(os.Stdin)

	if inputString, err := stdinReader.ReadString('\n'); err != nil {
		fmt.Println("Invalid input. Error: ", err)
	} else {
		if IsFound(inputString) {
			fmt.Println("Found !")
		} else {
			fmt.Println("Not found!")
		}

	}
}
