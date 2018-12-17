package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	fileName := "chronal_calibration.txt"
	f, err := os.Open(fileName)
	var result int64

	if err != nil {
		log.Fatal("Failed to open")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		numString := scanner.Text()
		num, _ := strconv.ParseInt(numString, 0, 0)
		result = result + num
	}

	println(result)
}
