package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func checkLoop(realObstacles, realObstacles2 [][]int, y, x, d int) bool {
	//sy := y
	//sx := x
	//sd := d
	oy := 0
	ox := 0
	switch d {
	case 0:
		oy = y - 1
		ox = x
	case 1:
		oy = y
		ox = x + 1
	case 2:
		oy = y + 1
		ox = x
	case 3:
		oy = y
		ox = x - 1
	}

	obstacles := make([][]int, len(realObstacles))
	obstacles2 := make([][]int, len(realObstacles2))
	copy(obstacles, realObstacles)
	copy(obstacles2, realObstacles2)
	obstacles[oy] = make([]int, len(realObstacles[oy]), len(realObstacles[oy])+1)
	obstacles2[ox] = make([]int, len(realObstacles2[ox]), len(realObstacles2[ox])+1)
	copy(obstacles[oy], realObstacles[oy])
	copy(obstacles2[ox], realObstacles2[ox])
	obstacles[oy] = append(obstacles[oy], ox)
	obstacles2[ox] = append(obstacles2[ox], oy)
	slices.Sort(obstacles[oy])
	slices.Sort(obstacles2[ox])

	d = (d + 1) % 4
	lx := len(obstacles2)
	ly := len(obstacles)
	vis := make([][]int, ly)
	for i := 0; i < ly; i++ {
		vis[i] = make([]int, lx)
	}
	for {
		switch d {
		case 0:
			newy := -1
			for _, yy := range obstacles2[x] {
				if yy > y {
					break
				}
				newy = yy
				//if x == ox && oy > yy && oy < y {
				//	newy = oy
				//}
			}
			if newy == -1 {
				//if sd == d && x == sx {
				//	return true
				//}
				return false
			}
			for yy := y - 1; yy > newy; yy-- {
				if vis[yy][x] == d+1 {
					return true
				}
				vis[yy][x] = d + 1
			}
			y = newy + 1
			//if sd == d && x == sx && y < sy {
			//	return true
			//}
		case 1:
			newx := -1
			for _, xx := range obstacles[y] {
				if xx < x {
					continue
				}
				//if y == oy && ox < xx && ox > x {
				//	newx = ox
				//	break
				//}
				newx = xx
				break
			}
			if newx == -1 {
				//if sd == d && y == sy {
				//	return true
				//}
				return false
			}
			for xx := x + 1; xx < newx; xx++ {
				if vis[y][xx] == d+1 {
					return true
				}
				vis[y][xx] = d + 1
			}
			x = newx - 1
			//if sd == d && y == sy && x > sx {
			//	return true
			//}
		case 2:
			newy := -1
			for _, yy := range obstacles2[x] {
				if yy < y {
					continue
				}
				//if x == ox && oy < yy && oy > y {
				//	newy = oy
				//	break
				//}
				newy = yy
				break
			}
			if newy == -1 {
				//if sd == d && x == sx {
				//	return true
				//}
				return false
			}
			for yy := y + 1; yy < newy; yy++ {
				if vis[yy][x] == d+1 {
					return true
				}
				vis[yy][x] = d + 1
			}
			y = newy - 1
			//if sd == d && x == sx && y > sy {
			//	return true
			//}
		case 3:
			newx := -1
			for _, xx := range obstacles[y] {
				if xx > x {
					break
				}
				newx = xx
				//if y == oy && ox > xx && ox < x {
				//	newx = ox
				//}
			}
			if newx == -1 {
				//if sd == d && y == sy {
				//	return true
				//}
				return false
			}
			for xx := x - 1; xx > newx; xx-- {
				if vis[y][xx] == d+1 {
					return true
				}
				vis[y][xx] = d + 1
			}
			x = newx + 1
			//if sd == d && y == sy && x < sx {
			//	return true
			//}
		}
		d = (d + 1) % 4
	}
}

func part1(r io.Reader) {
	sum := 1
	var s string
	scanner := bufio.NewScanner(r)
	var obstacles, obstacles2 [][]int
	var x, y, d, lx, ly int
	for scanner.Scan() {
		s = scanner.Text()
		if i := strings.Index(s, "^"); i != -1 {
			x = i
			y = len(obstacles)
		}
		rowObstacles := make([]int, 0)
		i := strings.Index(s, "#")
		for {
			if i == -1 {
				break
			}
			rowObstacles = append(rowObstacles, i)
			//s = s[i+1:]
			s = strings.Replace(s, "#", ".", 1)
			i = strings.Index(s, "#")
		}
		obstacles = append(obstacles, rowObstacles)
	}
	obstacles2 = make([][]int, len(s))
	for i, row := range obstacles {
		for _, el := range row {
			obstacles2[el] = append(obstacles2[el], i)
		}
	}
	lx = len(obstacles2)
	ly = len(obstacles)
	vis := make([][]int, ly)
	for i := 0; i < ly; i++ {
		vis[i] = make([]int, lx)
	}
	vis[y][x] = 1

	sum2 := 0

LUP:
	for {
		fmt.Println(sum)
		switch d {
		case 0:
			newy := -1
			for _, yy := range obstacles2[x] {
				if yy > y {
					break
				}
				newy = yy
			}
			if newy == -1 {
				sum += y
				break LUP
			}
			sum += y - newy - 1
			for yy := y - 1; yy > newy; yy-- {
				if vis[yy][x] == 0 {
					if checkLoop(obstacles, obstacles2, yy+1, x, d) {
						sum2 += 1
					}
				}
				vis[yy][x] = d + 1
			}
			y = newy + 1
		case 1:
			newx := -1
			for _, xx := range obstacles[y] {
				if xx < x {
					continue
				}
				newx = xx
				break
			}
			if newx == -1 {
				sum += lx - x - 1
				break LUP
			}
			sum += newx - x - 1
			for xx := x + 1; xx < newx; xx++ {
				if vis[y][xx] == 0 {
					if checkLoop(obstacles, obstacles2, y, xx-1, d) {
						sum2 += 1
					}
				}
				vis[y][xx] = d + 1
			}
			x = newx - 1
		case 2:
			newy := -1
			for _, yy := range obstacles2[x] {
				if yy < y {
					continue
				}
				newy = yy
				break
			}
			if newy == -1 {
				sum += ly - y - 1
				break LUP
			}
			sum += newy - y - 1
			for yy := y + 1; yy < newy; yy++ {
				if vis[yy][x] == 0 {
					if checkLoop(obstacles, obstacles2, yy-1, x, d) {
						sum2 += 1
					}
				}
				vis[yy][x] = d + 1
			}
			y = newy - 1
		case 3:
			newx := -1
			for _, xx := range obstacles[y] {
				if xx > x {
					break
				}
				newx = xx
			}
			if newx == -1 {
				sum += x
				break LUP
			}
			sum += x - newx - 1
			for xx := x - 1; xx > newx; xx-- {
				if vis[y][xx] == 0 {
					if sum == 493 {
						fmt.Println(xx)
					}
					if checkLoop(obstacles, obstacles2, y, xx+1, d) {
						sum2 += 1
					}
				}
				vis[y][xx] = d + 1
			}
			x = newx + 1
		}
		d = (d + 1) % 4
	}
	switch d {
	case 0:
		for yy := y - 1; yy >= 0; yy-- {
			if vis[yy][x] == 0 {
				if checkLoop(obstacles, obstacles2, yy+1, x, d) {
					sum2 += 1
				}
			}
			vis[yy][x] = d + 1
		}
	case 1:
		for xx := x + 1; xx < lx; xx++ {
			if vis[y][xx] == 0 {
				if checkLoop(obstacles, obstacles2, y, xx-1, d) {
					sum2 += 1
				}
			}
			vis[y][xx] = d + 1
		}
	case 2:
		for yy := y + 1; yy < ly; yy++ {
			if vis[yy][x] == 0 {
				if checkLoop(obstacles, obstacles2, yy-1, x, d) {
					sum2 += 1
				}
			}
			vis[yy][x] = d + 1
		}
	case 3:
		for xx := x - 1; xx >= 0; xx-- {
			if vis[y][xx] == 0 {
				if checkLoop(obstacles, obstacles2, y, xx+1, d) {
					sum2 += 1
				}
			}
			vis[y][xx] = d + 1
		}
	}
	sum = 0
	for _, row := range vis {
		fmt.Println()
		for _, v := range row {
			if v != 0 {
				fmt.Print("X")
				sum++
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println(sum, sum2)
}

func part2(r io.Reader) {

	//fmt.Println(mini)
}

func main() {
	f, err := os.Open("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
