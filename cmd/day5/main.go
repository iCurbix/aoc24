package main

import (
	"advent_of_code/pkg/util"
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part12(r io.Reader) {
	var s string
	sum := 0
	sum2 := 0
	rules := make(map[int][]int)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		s = scanner.Text()
		if s == "" {
			break
		}
		numScanner := bufio.NewScanner(strings.NewReader(s))
		numScanner.Split(util.ScanNumbers)
		numScanner.Scan()
		a, _ := strconv.Atoi(numScanner.Text())
		numScanner.Scan()
		b, _ := strconv.Atoi(numScanner.Text())
		rules[a] = append(rules[a], b)
	}

	for scanner.Scan() {
		var inpt []int
		s = scanner.Text()
		numScanner := bufio.NewScanner(strings.NewReader(s))
		numScanner.Split(util.ScanNumbers)
		for numScanner.Scan() {
			a, _ := strconv.Atoi(numScanner.Text())
			inpt = append(inpt, a)
		}

		is := slices.IsSortedFunc(inpt, func(a, b int) int {
			q, ok := rules[a]
			if ok {
				if slices.Contains(q, b) {
					return -1
				}
			}
			q, ok = rules[b]
			if ok {
				if slices.Contains(q, a) {
					return 1
				}
			}
			return 0
		})

		if is {
			sum += inpt[len(inpt)/2]
		} else {
			slices.SortFunc(inpt, func(a, b int) int {
				q, ok := rules[a]
				if ok {
					if slices.Contains(q, b) {
						return -1
					}
				}
				q, ok = rules[b]
				if ok {
					if slices.Contains(q, a) {
						return 1
					}
				}
				return 0
			})
			sum2 += inpt[len(inpt)/2]
		}

	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

//func part2(r io.Reader) {
//}

func main() {
	f, err := os.Open("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//buf := &bytes.Buffer{}
	//tr := io.TeeReader(f, buf)
	part12(f)
	//part2(buf)
}
