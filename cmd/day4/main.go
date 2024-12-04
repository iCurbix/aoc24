package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func part1(r io.Reader) {
	var s string
	var arr []string
	sum := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		arr = append(arr, s)
	}

	myi := len(arr) - 1
	mxi := len(arr[0]) - 1
	for y := 0; y <= myi; y++ {
		for x := 0; x <= mxi; x++ {
			if x+3 <= mxi {
				s = (string)([]uint8{arr[y][x], arr[y][x+1], arr[y][x+2], arr[y][x+3]})
				if s == "XMAS" || s == "SAMX" {
					sum += 1
				}
			}
			if y+3 <= myi && x+3 <= mxi {
				s = (string)([]uint8{arr[y][x], arr[y+1][x+1], arr[y+2][x+2], arr[y+3][x+3]})
				if s == "XMAS" || s == "SAMX" {
					sum += 1
				}
			}
			if y+3 <= myi {
				s = (string)([]uint8{arr[y][x], arr[y+1][x], arr[y+2][x], arr[y+3][x]})
				if s == "XMAS" || s == "SAMX" {
					sum += 1
				}
			}
			if x-3 >= 0 && y+3 <= myi {
				s = (string)([]uint8{arr[y][x], arr[y+1][x-1], arr[y+2][x-2], arr[y+3][x-3]})
				if s == "XMAS" || s == "SAMX" {
					sum += 1
				}
			}
		}
	}

	fmt.Println(sum)
}

func part2(r io.Reader) {
	var s string
	var arr []string
	sum := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		arr = append(arr, s)
	}

	myi := len(arr) - 1
	mxi := len(arr[0]) - 1
	for y := 1; y < myi; y++ {
		for x := 1; x < mxi; x++ {
			if arr[y][x] != 'A' {
				continue
			}
			dos := 0
			if arr[y-1][x-1] == 'M' && arr[y+1][x+1] == 'S' {
				dos += 1
			}
			if arr[y-1][x-1] == 'S' && arr[y+1][x+1] == 'M' {
				dos += 1
			}
			if arr[y-1][x+1] == 'M' && arr[y+1][x-1] == 'S' {
				dos += 1
			}
			if arr[y-1][x+1] == 'S' && arr[y+1][x-1] == 'M' {
				dos += 1
			}
			if dos == 2 {
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
