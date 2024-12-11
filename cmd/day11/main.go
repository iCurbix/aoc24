package main

import (
	"advent_of_code/pkg/util"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	a, b int
}

var pepega = map[pair]int{}

func do2(stone int, rounds int) int {
	if rounds == 0 {
		return 1
	}
	if res, ok := pepega[pair{stone, rounds}]; ok {
		return res
	}
	var strStone string
	var a, b int
	strStone = strconv.Itoa(stone)
	if stone == 0 {
		a = do2(1, rounds-1)
		pepega[pair{1, rounds - 1}] = a
		return a
	}
	if len(strStone)%2 == 0 {
		a, _ = strconv.Atoi(strStone[:len(strStone)/2])
		b, _ = strconv.Atoi(strStone[len(strStone)/2:])
		aa := do2(a, rounds-1)
		bb := do2(b, rounds-1)
		pepega[pair{a, rounds - 1}] = aa
		pepega[pair{b, rounds - 1}] = bb
		return aa + bb
	}
	a = do2(stone*2024, rounds-1)
	pepega[pair{stone * 2024, rounds - 1}] = a
	return a
}

func part1(r io.Reader) {
	sum := 0
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s := scanner.Text()
	numScanner := bufio.NewScanner(strings.NewReader(s))
	numScanner.Split(util.ScanNumbers)
	for numScanner.Scan() {
		num, _ := strconv.Atoi(numScanner.Text())
		sum += do2(num, 25)
	}
	fmt.Println(sum)
}

func part2(r io.Reader) {
	sum := 0
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s := scanner.Text()
	numScanner := bufio.NewScanner(strings.NewReader(s))
	numScanner.Split(util.ScanNumbers)
	for numScanner.Scan() {
		num, _ := strconv.Atoi(numScanner.Text())
		for i := 0; i < 74; i++ {
			do2(num, i)
		}
		sum += do2(num, 75)
	}
	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day11.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
