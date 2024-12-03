package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func part1(r io.Reader) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	var s string
	var sum, a, b int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		muls := re.FindAllStringSubmatch(s, -1)
		for _, mul := range muls {
			a, _ = strconv.Atoi(mul[1])
			b, _ = strconv.Atoi(mul[2])
			sum += a * b
		}
	}

	fmt.Println(sum)
}

func part2(r io.Reader) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	var s string
	var sum, a, b int
	enabled := true
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		muls := re.FindAllStringSubmatch(s, -1)
		for _, mul := range muls {
			switch mul[0][:3] {
			case "mul":
				if enabled {
					a, _ = strconv.Atoi(mul[1])
					b, _ = strconv.Atoi(mul[2])
					sum += a * b
				}
			case "do(":
				enabled = true
			case "don":
				enabled = false
			default:
				panic(mul[0])
			}

		}
	}

	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
