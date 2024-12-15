package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type dir int

const (
	up dir = iota
	right
	down
	left
)

type thing int

const (
	empty thing = iota
	robot
	box
	box2
	wall
)

func move(x, y int, di dir, grid [][]thing) bool {
	var nx, ny int
	switch di {
	case up:
		nx, ny = x, y-1
	case right:
		nx, ny = x+1, y
	case down:
		nx, ny = x, y+1
	case left:
		nx, ny = x-1, y
	}
	if grid[ny][nx] == wall {
		return false
	}
	if grid[ny][nx] == empty {
		grid[ny][nx] = grid[y][x]
		grid[y][x] = empty
		return true
	}
	if move(nx, ny, di, grid) {
		grid[ny][nx] = grid[y][x]
		grid[y][x] = empty
		return true
	}
	return false
}

func part1(r io.Reader) {
	var s string
	var grid [][]thing
	scanner := bufio.NewScanner(r)
	var x, y int
	yy := 0
	for scanner.Scan() {
		s = scanner.Text()
		if s == "" {
			break
		}
		row := make([]thing, len(s))
		for i := range row {
			switch s[i] {
			case '#':
				row[i] = wall
			case '.':
				row[i] = empty
			case 'O':
				row[i] = box
			case '@':
				row[i] = robot
				x = i
				y = yy
			default:
				panic("what the heck")
			}
		}
		grid = append(grid, row)
		yy++
	}

	for scanner.Scan() {
		s = scanner.Text()
		for _, ch := range s {
			switch ch {
			case '^':
				if move(x, y, up, grid) {
					y--
				}
			case '>':
				if move(x, y, right, grid) {
					x++
				}
			case 'v':
				if move(x, y, down, grid) {
					y++
				}
			case '<':
				if move(x, y, left, grid) {
					x--
				}
			}
		}
	}

	sum := 0
	for i, row := range grid {
		for j, t := range row {
			if t == box {
				sum += 100*i + j
			}
		}
	}
	fmt.Println(sum)
}

func check(x, y int, di dir, grid [][]thing) bool {
	var nx, ny int
	switch di {
	case up:
		nx, ny = x, y-1
	case right:
		nx, ny = x+1, y
	case down:
		nx, ny = x, y+1
	case left:
		nx, ny = x-1, y
	}
	if grid[ny][nx] == wall {
		return false
	}
	if grid[ny][nx] == empty {
		return true
	}
	if grid[ny][nx] == box {
		switch di {
		case up, down:
			return check(nx, ny, di, grid) && check(nx+1, ny, di, grid)
		case right:
			return check(nx+1, ny, di, grid)
		case left:
			panic("left???")
		}
	}
	switch di {
	case up, down:
		return check(nx, ny, di, grid) && check(nx-1, ny, di, grid)
	case right:
		panic("right???")
	case left:
		return check(nx-1, ny, di, grid)
	}
	return false
}

func move2(x, y int, di dir, grid [][]thing) {
	var nx, ny int
	switch di {
	case up:
		nx, ny = x, y-1
	case right:
		nx, ny = x+1, y
	case down:
		nx, ny = x, y+1
	case left:
		nx, ny = x-1, y
	}
	if grid[ny][nx] == empty {
		grid[ny][nx] = grid[y][x]
		grid[y][x] = empty
		return
	}
	if grid[ny][nx] == box {
		move2(nx, ny, di, grid)
		if di == up || di == down {
			move2(nx+1, ny, di, grid)
		}
		grid[ny][nx] = grid[y][x]
		grid[y][x] = empty
		return
	}
	if grid[ny][nx] == box2 {
		move2(nx, ny, di, grid)
		if di == up || di == down {
			move2(nx-1, ny, di, grid)
		}
		grid[ny][nx] = grid[y][x]
		grid[y][x] = empty
		return
	}
	panic("ruuun!!!")
}

func part2(r io.Reader) {
	var s string
	var grid [][]thing
	scanner := bufio.NewScanner(r)
	var x, y int
	yy := 0
	for scanner.Scan() {
		s = scanner.Text()
		if s == "" {
			break
		}
		row := make([]thing, len(s)*2)
		for i := range s {
			switch s[i] {
			case '#':
				row[i*2] = wall
				row[i*2+1] = wall
			case '.':
				row[i*2] = empty
				row[i*2+1] = empty
			case 'O':
				row[i*2] = box
				row[i*2+1] = box2
			case '@':
				row[i*2] = robot
				row[i*2+1] = empty
				x = i * 2
				y = yy
			default:
				panic("what the heck")
			}
		}
		grid = append(grid, row)
		yy++
	}

	for scanner.Scan() {
		s = scanner.Text()
		for _, ch := range s {
			//fmt.Println()
			//fmt.Println(ch)
			//for _, row := range grid {
			//	for _, t := range row {
			//		switch t {
			//		case wall:
			//			fmt.Print("#")
			//		case empty:
			//			fmt.Print(".")
			//		case robot:
			//			fmt.Print("@")
			//		case box:
			//			fmt.Print("[")
			//		case box2:
			//			fmt.Print("]")
			//		}
			//	}
			//	fmt.Println()
			//}
			switch ch {
			case '^':
				if check(x, y, up, grid) {
					move2(x, y, up, grid)
					y--
				}
			case '>':
				if check(x, y, right, grid) {
					move2(x, y, right, grid)
					x++
				}
			case 'v':
				if check(x, y, down, grid) {
					move2(x, y, down, grid)
					y++
				}
			case '<':
				if check(x, y, left, grid) {
					move2(x, y, left, grid)
					x--
				}
			}
		}
	}

	sum := 0
	for i, row := range grid {
		for j, t := range row {
			if t == box {
				sum += 100*i + j
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day15.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
