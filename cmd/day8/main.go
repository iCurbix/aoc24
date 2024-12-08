package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type Point struct {
	y, x int
}

func part1(r io.Reader) {
	sum := 0
	var s string
	m := make(map[int32][]Point)
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		s = scanner.Text()
		for j, c := range s {
			if c != '.' {
				m[c] = append(m[c], Point{i, j})
			}
		}
		i++
	}
	ly := i
	lx := len(s)
	m2 := make(map[Point]struct{})
	for _, v := range m {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				vx := v[i].x - v[j].x
				vy := v[i].y - v[j].y
				m2[Point{v[i].x + vx, v[i].y + vy}] = struct{}{}
				m2[Point{v[j].x - vx, v[j].y - vy}] = struct{}{}
			}
		}
	}

	for p := range m2 {
		if p.x >= 0 && p.x < lx && p.y >= 0 && p.y < ly {
			//fmt.Println(p)
			sum++
		}
	}
	fmt.Println(sum)
}

func part2(r io.Reader) {
	sum := 0
	var s string
	m := make(map[int32][]Point)
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		s = scanner.Text()
		for j, c := range s {
			if c != '.' {
				m[c] = append(m[c], Point{i, j})
			}
		}
		i++
	}
	ly := i
	lx := len(s)
	m2 := make(map[Point]struct{})
	for _, v := range m {
		for i := 0; i < len(v); i++ {
			for j := 0; j < len(v); j++ {
				if i == j {
					continue
				}
				vx := v[i].x - v[j].x
				vy := v[i].y - v[j].y
				for y, x := v[i].y, v[i].x; x >= 0 && x < lx && y >= 0 && y < ly; y, x = y+vy, x+vx {
					m2[Point{y, x}] = struct{}{}
				}
			}
		}
	}

	for range m2 {
		sum++
	}
	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
