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
	part1(values)
	part2(values)
}

func part1(values []int) {
	for i, v1 := range values {
		for j, v2 := range values {
			if i == j {
				continue
			}
			if v1+v2 == 2020 {
				fmt.Printf("Part1: %v + %v = 2020, %v * %v = %v\n",
					v1, v2, v1, v2, v1*v2)
				return
			}
		}
	}
}

func part2(values []int) {
	for i, v1 := range values {
		for j, v2 := range values {
			for k, v3 := range values {
				if i == j || i == k || j == k {
					continue
				}
				if v1+v2+v3 == 2020 {
					fmt.Printf("Part2: %v + %v + %v = 2020, %v * %v * %v = %v\n",
						v1, v2, v3, v1, v2, v3, v1*v2*v3)
					return
				}
			}
		}
	}
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
