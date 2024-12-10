package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type Point struct {
	x, y int
}

func part1(r io.Reader) {
	var inpt [][]int
	var s string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		aaa := make([]int, len(s))
		for i := range s {
			aaa[i] = int(s[i] - '0')
		}
		inpt = append(inpt, aaa)
	}
	ly := len(inpt)
	lx := len(inpt[0])
	qqq := make([][]map[Point]struct{}, ly)
	for i := 0; i < ly; i++ {
		qqq[i] = make([]map[Point]struct{}, lx)
		for j := 0; j < lx; j++ {
			qqq[i][j] = make(map[Point]struct{})
		}
	}

	for y, row := range inpt {
		for x, h := range row {
			if h == 9 {
				qqq[y][x][Point{x, y}] = struct{}{}
			}
		}
	}

	sum := 0

	for i := 8; i >= 0; i-- {
		for y, row := range inpt {
			for x, h := range row {
				if h == i {
					if y > 0 && inpt[y-1][x] == i+1 {
						for k := range qqq[y-1][x] {
							qqq[y][x][k] = struct{}{}
						}
					}
					if x < lx-1 && inpt[y][x+1] == i+1 {
						for k := range qqq[y][x+1] {
							qqq[y][x][k] = struct{}{}
						}
					}
					if y < ly-1 && inpt[y+1][x] == i+1 {
						for k := range qqq[y+1][x] {
							qqq[y][x][k] = struct{}{}
						}
					}
					if x > 0 && inpt[y][x-1] == i+1 {
						for k := range qqq[y][x-1] {
							qqq[y][x][k] = struct{}{}
						}
					}
					if i == 0 {
						sum += len(qqq[y][x])
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func part2(r io.Reader) {
	var inpt [][]int
	var s string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		aaa := make([]int, len(s))
		for i := range s {
			aaa[i] = int(s[i] - '0')
		}
		inpt = append(inpt, aaa)
	}
	ly := len(inpt)
	lx := len(inpt[0])
	qqq := make([][]int, ly)
	for i := 0; i < ly; i++ {
		qqq[i] = make([]int, lx)
	}

	for y, row := range inpt {
		for x, h := range row {
			if h == 9 {
				qqq[y][x] = 1
			}
		}
	}

	sum := 0

	for i := 8; i >= 0; i-- {
		for y, row := range inpt {
			for x, h := range row {
				if h == i {
					if y > 0 && inpt[y-1][x] == i+1 {
						qqq[y][x] += qqq[y-1][x]
					}
					if x < lx-1 && inpt[y][x+1] == i+1 {
						qqq[y][x] += qqq[y][x+1]
					}
					if y < ly-1 && inpt[y+1][x] == i+1 {
						qqq[y][x] += qqq[y+1][x]
					}
					if x > 0 && inpt[y][x-1] == i+1 {
						qqq[y][x] += qqq[y][x-1]
					}
					if i == 0 {
						sum += qqq[y][x]
					}
				}
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
