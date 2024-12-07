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
	sum := 0
	var s string
	scanner := bufio.NewScanner(r)
	aaa := 0

	for scanner.Scan() {
		//fmt.Println(aaa)
		aaa += 1
		s = scanner.Text()
		numScanner := bufio.NewScanner(strings.NewReader(s))
		numScanner.Split(util.ScanNumbers)
		numScanner.Scan()
		num, _ := strconv.Atoi(numScanner.Text())
		var nums []int
		for numScanner.Scan() {
			numm, _ := strconv.Atoi(numScanner.Text())
			nums = append(nums, numm)
		}
		for i := 0; i < int(math.Pow(2, float64(len(nums)-1))); i++ {
			res := nums[0]
			for ix, x := range fmt.Sprintf(fmt.Sprintf("%%0%db", len(nums)-1), i) {
				if res > num {
					break
				}
				switch x {
				case '0':
					res *= nums[ix+1]
				case '1':
					res += nums[ix+1]
				}
			}
			if res == num {
				sum += num
				break
			}
		}
	}

	fmt.Println(sum)
}

func part2(r io.Reader) {
	sum := 0
	var s string
	scanner := bufio.NewScanner(r)
	aaa := 0

	for scanner.Scan() {
		fmt.Println(aaa)
		aaa += 1
		s = scanner.Text()
		numScanner := bufio.NewScanner(strings.NewReader(s))
		numScanner.Split(util.ScanNumbers)
		numScanner.Scan()
		num, _ := strconv.Atoi(numScanner.Text())
		var nums []int
		for numScanner.Scan() {
			numm, _ := strconv.Atoi(numScanner.Text())
			nums = append(nums, numm)
		}
		var i int64
		for i = 0; i < int64(math.Pow(3, float64(len(nums)-1))); i++ {
			res := nums[0]
			for ix, x := range fmt.Sprintf(fmt.Sprintf("%%0%ds", len(nums)-1), strconv.FormatInt(i, 3)) {
				if res > num {
					break
				}
				switch x {
				case '0':
					res *= nums[ix+1]
				case '1':
					res += nums[ix+1]
				case '2':
					res, _ = strconv.Atoi(fmt.Sprintf("%d%d", res, nums[ix+1]))
				}
			}
			if res == num {
				sum += num
				break
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
