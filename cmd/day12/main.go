package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func part1(r io.Reader) {
	var inpt []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		inpt = append(inpt, scanner.Text())
	}
	ly := len(inpt)
	lx := len(inpt[0])
	ids := make([][]int, ly)
	areas := make(map[int]int)
	perimeters := make(map[int]int)
	idSeq := 1
	for y := 0; y < ly; y++ {
		ids[y] = make([]int, lx)
		for x := 0; x < lx; x++ {
			neigh := 0
			a := true
			if x > 0 && inpt[y][x-1] == inpt[y][x] {
				neigh++
			}
			if y > 0 && inpt[y-1][x] == inpt[y][x] {
				neigh++
				a = false
			}
			switch neigh {
			case 2:
				if ids[y][x-1] == ids[y-1][x] {
					ids[y][x] = ids[y][x-1]
					areas[ids[y][x-1]]++
				} else {
					sigmaId := ids[y-1][x]
					betaId := ids[y][x-1]
					areas[sigmaId] += areas[betaId] + 1
					perimeters[sigmaId] += perimeters[betaId]
					ids[y][x] = sigmaId
					delete(areas, betaId)
					delete(perimeters, betaId)
					for yy := 0; yy <= y; yy++ {
						for xx := 0; xx <= x; xx++ {
							if ids[yy][xx] == betaId {
								ids[yy][xx] = sigmaId
							}
						}
					}

				}
			case 1:
				neighId := 0
				if a {
					neighId = ids[y][x-1]
				} else {
					neighId = ids[y-1][x]
				}
				ids[y][x] = neighId
				areas[neighId]++
				perimeters[neighId] += 2
			case 0:
				ids[y][x] = idSeq
				areas[idSeq] = 1
				perimeters[idSeq] = 4
				idSeq++
			}
		}
	}

	sum := 0
	for k, v := range areas {
		sum += v * perimeters[k]
	}

	fmt.Println(sum)

}

func part2(r io.Reader) {
	var inpt []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		inpt = append(inpt, scanner.Text())
	}
	ly := len(inpt)
	lx := len(inpt[0])
	ids := make([][]int, ly)
	areas := make(map[int]int)
	sides := make(map[int]int)
	idSeq := 1
	for y := 0; y < ly; y++ {
		ids[y] = make([]int, lx)
		for x := 0; x < lx; x++ {
			neigh := 0
			a := true
			if x > 0 && inpt[y][x-1] == inpt[y][x] {
				neigh++
			}
			if y > 0 && inpt[y-1][x] == inpt[y][x] {
				neigh++
				a = false
			}
			switch neigh {
			case 2:
				if ids[y][x-1] == ids[y-1][x] {
					ids[y][x] = ids[y][x-1]
					areas[ids[y][x-1]]++
					if x == lx-1 || ids[y-1][x] != ids[y-1][x+1] {
						sides[ids[y][x-1]] -= 2
					}
				} else {
					sigmaId := ids[y-1][x]
					betaId := ids[y][x-1]
					areas[sigmaId] += areas[betaId] + 1
					sides[sigmaId] += sides[betaId]
					if x == lx-1 || ids[y-1][x] != ids[y-1][x+1] {
						sides[sigmaId] -= 2
					}
					ids[y][x] = sigmaId
					delete(areas, betaId)
					delete(sides, betaId)
					for yy := 0; yy <= y; yy++ {
						for xx := 0; xx <= x; xx++ {
							if ids[yy][xx] == betaId {
								ids[yy][xx] = sigmaId
							}
						}
					}

				}
			case 1:
				neighId := 0
				if a {
					neighId = ids[y][x-1]
					if y > 0 && ids[y-1][x-1] == neighId {
						sides[neighId] += 2
					}
				} else {
					neighId = ids[y-1][x]
					if x < lx-1 && ids[y-1][x+1] == neighId {
						sides[neighId] += 2
					}
					if x > 0 && ids[y-1][x-1] == neighId {
						sides[neighId] += 2
					}
				}
				ids[y][x] = neighId
				areas[neighId]++
			case 0:
				ids[y][x] = idSeq
				areas[idSeq] = 1
				sides[idSeq] = 4
				idSeq++
			}
		}
	}

	sum := 0
	for k, v := range areas {
		sum += v * sides[k]
	}

	fmt.Println(sum)

}

func main() {
	f, err := os.Open("inputs/day12.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
