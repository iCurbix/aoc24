package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(r io.Reader) {
	var s string
	sum := 0
	scanner := bufio.NewScanner(r)
	var left, right []int
	for scanner.Scan() {
		s = scanner.Text()
		ss := strings.Split(s, "   ")
		l, _ := strconv.Atoi(ss[0])
		r, _ := strconv.Atoi(ss[1])
		left = append(left, l)
		right = append(right, r)
	}
	sort.Ints(left)
	sort.Ints(right)

	var dif int
	for i := 0; i < len(left); i++ {
		dif = left[i] - right[i]
		if dif < 0 {
			dif *= -1
		}
		sum += dif
	}
	fmt.Println(sum)
}

func part2(r io.Reader) {
	var s string
	sum := 0
	scanner := bufio.NewScanner(r)
	var left, right []int
	for scanner.Scan() {
		s = scanner.Text()
		ss := strings.Split(s, "   ")
		l, _ := strconv.Atoi(ss[0])
		r, _ := strconv.Atoi(ss[1])
		left = append(left, l)
		right = append(right, r)
	}
	counts := make(map[int]int)
	for _, x := range right {
		counts[x]++
	}
	for _, x := range left {
		sum += counts[x] * x
	}
	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
