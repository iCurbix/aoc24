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

func part1(r io.Reader) {
	var s string
	var ax, ay, bx, by, px, py, tok, mini int
	sum := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		numScanner := bufio.NewScanner(strings.NewReader(s))
		numScanner.Split(util.ScanNumbers)
		numScanner.Scan()
		ax, _ = strconv.Atoi(numScanner.Text())
		numScanner.Scan()
		ay, _ = strconv.Atoi(numScanner.Text())

		scanner.Scan()
		s = scanner.Text()
		numScanner = bufio.NewScanner(strings.NewReader(s))
		numScanner.Split(util.ScanNumbers)
		numScanner.Scan()
		bx, _ = strconv.Atoi(numScanner.Text())
		numScanner.Scan()
		by, _ = strconv.Atoi(numScanner.Text())

		scanner.Scan()
		s = scanner.Text()
		numScanner = bufio.NewScanner(strings.NewReader(s))
		numScanner.Split(util.ScanNumbers)
		numScanner.Scan()
		px, _ = strconv.Atoi(numScanner.Text())
		numScanner.Scan()
		py, _ = strconv.Atoi(numScanner.Text())

		scanner.Scan()
		mini = math.MaxInt

		for i := 0; i <= 100; i++ {
			for j := 0; j <= 100; j++ {
				if i*ax+j*bx == px && i*ay+j*by == py {
					tok = 3*i + j
					if tok < mini {
						mini = tok
					}
				}
			}
		}

		if mini != math.MaxInt {
			sum += mini
		}
	}

	fmt.Println(sum)
}

func part2(r io.Reader) {
	var s string
	var ax, ay, bx, by, px, py int
	sum := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		numScanner := bufio.NewScanner(strings.NewReader(s))
		numScanner.Split(util.ScanNumbers)
		numScanner.Scan()
		ax, _ = strconv.Atoi(numScanner.Text())
		numScanner.Scan()
		ay, _ = strconv.Atoi(numScanner.Text())

		scanner.Scan()
		s = scanner.Text()
		numScanner = bufio.NewScanner(strings.NewReader(s))
		numScanner.Split(util.ScanNumbers)
		numScanner.Scan()
		bx, _ = strconv.Atoi(numScanner.Text())
		numScanner.Scan()
		by, _ = strconv.Atoi(numScanner.Text())

		scanner.Scan()
		s = scanner.Text()
		numScanner = bufio.NewScanner(strings.NewReader(s))
		numScanner.Split(util.ScanNumbers)
		numScanner.Scan()
		px, _ = strconv.Atoi(numScanner.Text())
		numScanner.Scan()
		py, _ = strconv.Atoi(numScanner.Text())

		scanner.Scan()
		px += 10000000000000
		py += 10000000000000

		if ax*by == ay*bx {
			// really? no optimization needed after all?
			fmt.Println("AAAAAA")
		}

		q := bx * ay
		qq := px * ay
		qqq := by * ax
		qqqq := py * ax
		b := (qqqq - qq) / (qqq - q)
		if b > 0 && b*(qqq-q) == qqqq-qq {
			a := (px - b*bx) / ax
			if a*ax == px-b*bx {
				sum += b + a*3
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day13.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
