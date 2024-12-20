package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var kasza = make(map[string]bool)

func possible(towels map[uint8]map[int][]string, pattern string) bool {
	if jazda, ok := kasza[pattern]; ok {
		if jazda {
			return true
		}
		return false
	}
	if pattern == "" {
		return true
	}
	for i := len(pattern); i > 0; i-- {
		for _, t := range towels[pattern[0]][i] {
			if t == pattern[:i] && possible(towels, pattern[i:]) {
				kasza[pattern] = true
				return true
			}
		}
	}
	kasza[pattern] = false
	return false
}

var kasza2 = make(map[string]int)

func nums(towels map[uint8]map[int][]string, pattern string) int {
	//fmt.Println(pattern)
	if jazda, ok := kasza[pattern]; ok {
		if jazda {
			if ile, ok := kasza2[pattern]; ok {
				return ile
			}
		} else {
			return 0
		}
	}
	if pattern == "" {
		return 1
	}
	posnum := 0
	for i := len(pattern); i > 0; i-- {
		for _, t := range towels[pattern[0]][i] {
			//fmt.Println(pattern, d, i, j)
			if t == pattern[:i] {
				posnum += nums(towels, pattern[i:])
			}
		}
	}
	kasza2[pattern] = posnum
	return posnum
}

func part1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var s string

	scanner.Scan()
	s = scanner.Text()
	ss := strings.Split(s, ", ")
	towels := make(map[uint8]map[int][]string)
	for _, sss := range ss {
		if towels[sss[0]] == nil {
			towels[sss[0]] = make(map[int][]string)
		}
		towels[sss[0]][len(sss)] = append(towels[sss[0]][len(sss)], sss)
	}
	scanner.Scan()
	sum := 0
	i := 0
	for scanner.Scan() {
		i++
		s = scanner.Text()
		for q := len(s) - 1; q > 1; q-- {
			possible(towels, s[q:])
		}
		if possible(towels, s) {
			sum++
		}
	}
	fmt.Println(sum)
}

func part2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var s string

	scanner.Scan()
	s = scanner.Text()
	ss := strings.Split(s, ", ")
	towels := make(map[uint8]map[int][]string)
	for _, sss := range ss {
		if towels[sss[0]] == nil {
			towels[sss[0]] = make(map[int][]string)
		}
		towels[sss[0]][len(sss)] = append(towels[sss[0]][len(sss)], sss)
	}
	scanner.Scan()
	sum := 0
	i := 0
	for scanner.Scan() {
		fmt.Println(i)
		i++
		s = scanner.Text()
		//fmt.Println(s)
		for q := len(s) - 1; q > 1; q-- {
			//fmt.Println(q)
			nums(towels, s[q:])
		}
		fmt.Println(nums(towels, s))
		sum += nums(towels, s)
	}
	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day19.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
