package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	values := readExpenseReport("input")
	fmt.Printf("Num values: %v", len(values))
}

func readExpenseReport(filename string) (report []int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if v, err := strconv.Atoi(sc.Text()); err == nil {
			report = append(report, v)
		} else {
			log.Fatalf("invalid input value: %v", err)
		}
	}
	return
}
