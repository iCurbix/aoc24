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

func isSafe(inpt []int) bool {
	var a, b, dif int
	dec := false
	b = inpt[0]
	a = inpt[1]
	dif = a - b
	if dif > 0 {
		dec = false
	} else {
		dec = true
	}
	if dif == 0 {
		return false
	}
	if dif > 3 || dif < -3 {
		return false
	}
	for _, b = range inpt[2:] {
		dif = b - a
		if dif > 0 {
			if dec {
				return false
			}
		} else {
			if !dec {
				return false
			}
		}
		if dif == 0 {
			return false
		}
		if dif > 3 || dif < -3 {
			return false
		}

		a = b
	}
	return true
}

func part1(r io.Reader) {
	//var s string
	sum := 0
	var a, b, dif int
	var safe, dec bool
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		//s = scanner.Text()
		ns := bufio.NewScanner(strings.NewReader(scanner.Text()))
		ns.Split(util.ScanNumbers)
		ns.Scan()
		safe = true
		b, _ = strconv.Atoi(ns.Text())
		ns.Scan()
		a, _ = strconv.Atoi(ns.Text())
		dif = a - b
		if dif > 0 {
			dec = false
		} else {
			dec = true
		}
		if dif == 0 {
			safe = false
		}
		if dif > 3 || dif < -3 {
			safe = false
		}
		for ns.Scan() {
			if !safe {
				continue
			}
			b, _ = strconv.Atoi(ns.Text())
			dif = b - a
			if dif > 0 {
				if dec {
					safe = false
				}
			} else {
				if !dec {
					safe = false
				}
			}
			if dif == 0 {
				safe = false
			}
			if dif > 3 || dif < -3 {
				safe = false
			}

			a = b
		}

		if safe {
			sum += 1
		}

	}
	fmt.Println(sum)
}

func part2(r io.Reader) {
	//var s string
	sum := 0
	var a int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		//s = scanner.Text()
		inpt := make([]int, 0, 16)
		ns := bufio.NewScanner(strings.NewReader(scanner.Text()))
		ns.Split(util.ScanNumbers)
		for ns.Scan() {
			a, _ = strconv.Atoi(ns.Text())
			inpt = append(inpt, a)
		}
		for i := 0; i < len(inpt); i++ {
			inpt2 := make([]int, i, len(inpt)-1)
			copy(inpt2, inpt[:i])
			inpt2 = append(inpt2, inpt[i+1:]...)
			if isSafe(inpt2) {
				sum += 1
				break
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
