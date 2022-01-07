package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputSent string
	var removed string

	var cond1 bool
	var cond2 bool
	var cond3 bool

	prefix := "I"
	suffix := "N"
	middle := "A"

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter sentence :")
	if scanner.Scan() {
		inputSent = scanner.Text()
	}

	inputSent = strings.ToUpper(inputSent)
	inputSent = strings.TrimSpace(inputSent)
	removed = inputSent[1 : len(inputSent)-1]

	cond1 = strings.HasPrefix(inputSent, prefix)
	cond2 = strings.HasSuffix(inputSent, suffix)
	cond3 = strings.Contains(removed, middle)

	if cond1 && cond2 && cond3 {
		fmt.Println("Found !")
	} else {
		fmt.Println("Not Found !")
	}
}
