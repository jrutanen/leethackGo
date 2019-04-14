package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	binaryNumbers, err := os.Open("bin.txt")
	defer binaryNumbers.Close()
	fmt.Print(err)

	scanner := bufio.NewScanner(binaryNumbers)
	//Go through the file line by line
	for scanner.Scan() {
		//convert the string presenting a binary value to int
		n, _ := strconv.ParseUint(scanner.Text(), 2, 8)
		//print the character for the ascii code
		fmt.Printf("%c", n)
	}
}
