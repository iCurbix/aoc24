package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
)

func do(maze [][]bool, scores [][]int, score, x, y int) {
	sc := scores[y][x]
	if sc <= score {
		return
	}
	scores[y][x] = score
	nx, ny := x+1, y
	if !maze[ny][nx] {
		do(maze, scores, score+1, nx, ny)
	}
	nx, ny = x-1, y
	if !maze[ny][nx] {
		do(maze, scores, score+1, nx, ny)
	}
	nx, ny = x, y+1
	if !maze[ny][nx] {
		do(maze, scores, score+1, nx, ny)
	}
	nx, ny = x, y-1
	if !maze[ny][nx] {
		do(maze, scores, score+1, nx, ny)
	}
}

func part1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var s string

	var maze [][]bool
	var stepsS [][]int
	var stepsE [][]int
	var sx, sy, ex, ey int

	y := 0
	for scanner.Scan() {
		s = scanner.Text()
		mazeRow := make([]bool, len(s))
		stepsS = append(stepsS, make([]int, len(s)))
		stepsE = append(stepsE, make([]int, len(s)))
		for x, ch := range s {
			mazeRow[x] = ch == '#'
			if ch == 'S' {
				sx = x
				sy = y
			}
			if ch == 'E' {
				ex = x
				ey = y
			}
			stepsS[y][x] = math.MaxInt
			stepsE[y][x] = math.MaxInt
		}
		maze = append(maze, mazeRow)
		y++
	}

	do(maze, stepsS, 0, sx, sy)
	do(maze, stepsE, 0, ex, ey)

	ly := len(stepsS)
	lx := len(stepsS[0])

	target := 100
	saved := 0
	sum := 0

	for y = 0; y < ly; y++ {
		for x := 0; x < lx; x++ {
			if stepsS[y][x] < math.MaxInt {
				nx, ny := x+2, y
				if nx >= 0 && ny >= 0 && nx < lx && ny < ly && stepsE[ny][nx] < math.MaxInt {
					saved = stepsS[ey][ex] - stepsS[y][x] - stepsE[ny][nx] - 2
					if saved >= target {
						sum++
					}
				}
				nx, ny = x-2, y
				if nx >= 0 && ny >= 0 && nx < lx && ny < ly && stepsE[ny][nx] < math.MaxInt {
					saved = stepsS[ey][ex] - stepsS[y][x] - stepsE[ny][nx] - 2
					//fmt.Println(saved)
					if saved >= target {
						sum++
					}
				}
				nx, ny = x, y+2
				if nx >= 0 && ny >= 0 && nx < lx && ny < ly && stepsE[ny][nx] < math.MaxInt {
					saved = stepsS[ey][ex] - stepsS[y][x] - stepsE[ny][nx] - 2
					//fmt.Println(saved)
					if saved >= target {
						sum++
					}
				}
				nx, ny = x, y-2
				if nx >= 0 && ny >= 0 && nx < lx && ny < ly && stepsE[ny][nx] < math.MaxInt {
					saved = stepsS[ey][ex] - stepsS[y][x] - stepsE[ny][nx] - 2
					//fmt.Println(saved)
					if saved >= target {
						sum++
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func part2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var s string

	var maze [][]bool
	var stepsS [][]int
	var stepsE [][]int
	var sx, sy, ex, ey int

	y := 0
	for scanner.Scan() {
		s = scanner.Text()
		mazeRow := make([]bool, len(s))
		stepsS = append(stepsS, make([]int, len(s)))
		stepsE = append(stepsE, make([]int, len(s)))
		for x, ch := range s {
			mazeRow[x] = ch == '#'
			if ch == 'S' {
				sx = x
				sy = y
			}
			if ch == 'E' {
				ex = x
				ey = y
			}
			stepsS[y][x] = math.MaxInt
			stepsE[y][x] = math.MaxInt
		}
		maze = append(maze, mazeRow)
		y++
	}

	do(maze, stepsS, 0, sx, sy)
	do(maze, stepsE, 0, ex, ey)

	ly := len(stepsS)
	lx := len(stepsS[0])

	target := 100
	saved := 0
	sum := 0
	aaa := make(map[int]int)

	for y = 0; y < ly; y++ {
		for x := 0; x < lx; x++ {
			if stepsS[y][x] < math.MaxInt {
				for dy := 0; dy <= 20; dy++ {
					for dx := 0; dx <= 20-dy; dx++ {
						nx, ny := x+dx, y+dy
						if nx >= 0 && ny >= 0 && nx < lx && ny < ly && stepsE[ny][nx] < math.MaxInt {
							saved = stepsE[y][x] - stepsE[ny][nx] - dx - dy
							aaa[saved]++
							if saved >= target {
								sum++
							}
						}
						if dx != 0 {
							nx, ny = x-dx, y+dy
							if nx >= 0 && ny >= 0 && nx < lx && ny < ly && stepsE[ny][nx] < math.MaxInt {
								saved = stepsE[y][x] - stepsE[ny][nx] - dx - dy
								aaa[saved]++
								if saved >= target {
									sum++
								}
							}
						}
						if dy != 0 {
							nx, ny = x+dx, y-dy
							if nx >= 0 && ny >= 0 && nx < lx && ny < ly && stepsE[ny][nx] < math.MaxInt {
								saved = stepsE[y][x] - stepsE[ny][nx] - dx - dy
								aaa[saved]++
								if saved >= target {
									sum++
								}
							}
						}
						if dx != 0 && dy != 0 {
							nx, ny = x-dx, y-dy
							if nx >= 0 && ny >= 0 && nx < lx && ny < ly && stepsE[ny][nx] < math.MaxInt {
								saved = stepsE[y][x] - stepsE[ny][nx] - dx - dy
								aaa[saved]++
								if saved >= target {
									sum++
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(aaa)
	fmt.Println(">=", target, sum)
}
func main() {
	f, err := os.Open("inputs/day20.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
