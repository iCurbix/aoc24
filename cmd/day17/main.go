package main

import (
	"advent_of_code/pkg/util"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type computr struct {
	A, B, C, ptr, outIx int
	instr               []int
}

func (c *computr) getCombo() int {
	a := c.instr[c.ptr+1]
	switch a {
	case 0, 1, 2, 3:
		return a
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		panic("panik")
	}
}

func (c *computr) getLit() int {
	return c.instr[c.ptr+1]
}

func (c *computr) goo() {
	for {
		//fmt.Println(c.ptr, len(c.instr))
		if c.ptr >= len(c.instr) {
			break
		}
		switch c.instr[c.ptr] {
		case 0:
			c.A = c.A / int(math.Pow(2, float64(c.getCombo())))
			c.ptr += 2
		case 1:
			c.B = c.B ^ c.getLit()
			c.ptr += 2
		case 2:
			c.B = c.getCombo() % 8
			c.ptr += 2
		case 3:
			if c.A != 0 {
				c.ptr = c.getLit()
			} else {
				c.ptr += 2
			}
		case 4:
			c.B = c.B ^ c.C
			c.ptr += 2
		case 5:
			fmt.Print(c.getCombo() % 8)
			fmt.Print(",")
			c.ptr += 2
		case 6:
			c.B = c.A / int(math.Pow(2, float64(c.getCombo())))
			c.ptr += 2
		case 7:
			c.C = c.A / int(math.Pow(2, float64(c.getCombo())))
			c.ptr += 2
		default:
			panic("panik!!!")
		}
	}
}

func part1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var s string
	var A, B, C int
	var inpt []int
	scanner.Scan()
	s = scanner.Text()
	numScanner := bufio.NewScanner(strings.NewReader(s))
	numScanner.Split(util.ScanNumbers)
	numScanner.Scan()
	A, _ = strconv.Atoi(numScanner.Text())
	scanner.Scan()
	s = scanner.Text()
	numScanner = bufio.NewScanner(strings.NewReader(s))
	numScanner.Split(util.ScanNumbers)
	numScanner.Scan()
	B, _ = strconv.Atoi(numScanner.Text())
	scanner.Scan()
	s = scanner.Text()
	numScanner = bufio.NewScanner(strings.NewReader(s))
	numScanner.Split(util.ScanNumbers)
	numScanner.Scan()
	C, _ = strconv.Atoi(numScanner.Text())
	scanner.Scan()
	scanner.Scan()
	s = scanner.Text()
	numScanner = bufio.NewScanner(strings.NewReader(s))
	numScanner.Split(util.ScanNumbers)
	for numScanner.Scan() {
		q, _ := strconv.Atoi(numScanner.Text())
		inpt = append(inpt, q)
	}

	c := &computr{A: A, B: B, C: C, ptr: 0, instr: inpt}
	fmt.Println(c)
	c.goo()
	fmt.Println()
}

func main() {
	f, err := os.Open("inputs/day17.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	//part2(buf)

}
