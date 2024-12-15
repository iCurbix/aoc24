package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var lx = 101
var ly = 103

func part1(r io.Reader) {
	var s string
	var x, y, vx, vy int
	sum1 := 0
	sum2 := 0
	sum3 := 0
	sum4 := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		ss := strings.Split(s, " ")
		sss := strings.Split(ss[0], "=")
		sss = strings.Split(sss[1], ",")
		ssss := strings.Split(ss[1], "=")
		ssss = strings.Split(ssss[1], ",")
		x, _ = strconv.Atoi(sss[0])
		y, _ = strconv.Atoi(sss[1])
		vx, _ = strconv.Atoi(ssss[0])
		vy, _ = strconv.Atoi(ssss[1])

		x = (x + vx*100) % lx
		if x < 0 {
			x += lx
		}
		y = (y + vy*100) % ly
		if y < 0 {
			y += ly
		}

		if x < lx/2 {
			if y < ly/2 {
				sum1++
			}
			if y > ly/2 {
				sum2++
			}
		}
		if x > lx/2 {
			if y < ly/2 {
				sum3++
			}
			if y > ly/2 {
				sum4++
			}
		}
	}
	fmt.Println(sum1, sum2, sum3, sum4)
	fmt.Println(sum1 * sum2 * sum3 * sum4)
}

type vel struct {
	vx, vy int
}

func part2(r io.Reader) {
	var s string
	var x, y, vx, vy int
	grid := make([][][]*vel, ly)
	for i := 0; i < ly; i++ {
		grid[i] = make([][]*vel, lx)
		for j := 0; j < lx; j++ {
			grid[i][j] = make([]*vel, 0)
		}
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		ss := strings.Split(s, " ")
		sss := strings.Split(ss[0], "=")
		sss = strings.Split(sss[1], ",")
		ssss := strings.Split(ss[1], "=")
		ssss = strings.Split(ssss[1], ",")
		x, _ = strconv.Atoi(sss[0])
		y, _ = strconv.Atoi(sss[1])
		vx, _ = strconv.Atoi(ssss[0])
		vy, _ = strconv.Atoi(ssss[1])

		grid[y][x] = append(grid[y][x], &vel{vx, vy})
	}

	f, _ := os.Create("cmd/day14/out.txt")
	defer f.Close()

	for i := 0; i < 10000; i++ {
		grid2 := make([][][]*vel, ly)
		for i := 0; i < ly; i++ {
			grid2[i] = make([][]*vel, lx)
			for j := 0; j < lx; j++ {
				grid2[i][j] = make([]*vel, 0)
			}
		}
		fmt.Fprintln(f, i)
		fmt.Println(i)
		for y := 0; y < ly; y++ {
			for x := 0; x < lx; x++ {
				if len(grid[y][x]) != 0 {
					fmt.Fprint(f, "*")
					for _, v := range grid[y][x] {
						xx := (x + v.vx) % lx
						if xx < 0 {
							xx += lx
						}
						yy := (y + v.vy) % ly
						if yy < 0 {
							yy += ly
						}
						grid2[yy][xx] = append(grid2[yy][xx], v)
					}
				} else {
					fmt.Fprint(f, " ")
				}
			}
			fmt.Fprintln(f)
		}
		fmt.Fprintln(f)
		grid = grid2
	}
}

func main() {
	f, err := os.Open("inputs/day14.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
