package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func part1(r io.Reader) {
	sum := 0
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s := scanner.Text()
	inptBackIx := len(s) - 1
	fileBackIx := len(s) / 2
	fileFrontIx := 1
	compressedFrontIx := int(s[0] - 48)
	re := int(s[inptBackIx] - 48)
LUP:
	for inptFrontIx := 1; inptFrontIx < len(s); inptFrontIx += 2 {
		empty := int(s[inptFrontIx] - 48)
		for j := 0; j < empty; j++ {
			sum += fileBackIx * compressedFrontIx
			compressedFrontIx++
			re--
			if re == 0 {
				fileBackIx--
				if fileBackIx == fileFrontIx-1 {
					break LUP
				}
				inptBackIx -= 2
				re = int(s[inptBackIx] - 48)
			}
		}
		if fileBackIx == fileFrontIx {
			x := re
			for j := 0; j < x; j++ {
				sum += fileFrontIx * compressedFrontIx
				compressedFrontIx++
			}
			break
		}
		x := int(s[inptFrontIx+1] - 48)
		for j := 0; j < x; j++ {
			sum += fileFrontIx * compressedFrontIx
			compressedFrontIx++
		}
		fileFrontIx++
	}

	fmt.Println(sum)
}

type File struct {
	ix int
	l  int
}

type Empty struct {
	remaining int
	files     []File
}

func part2(r io.Reader) {
	sum := 0
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s := scanner.Text()
	files := make([]File, 0, len(s)/2+1)
	emptys := make([]Empty, 0, len(s)/2)
	for i, chr := range s {
		if i%2 == 1 {
			emptys = append(emptys, Empty{remaining: int(chr - 48), files: nil})
			continue
		}
		files = append(files, File{i / 2, int(chr - 48)})
	}

FLUP:
	for i := len(files) - 1; i >= 0; i-- {
		for j := range emptys {
			if j >= i {
				break
			}
			if emptys[j].remaining >= files[i].l {
				emptys[j].remaining -= files[i].l
				emptys[j].files = append(emptys[j].files, files[i])
				files[i].ix = 0
				continue FLUP
			}
		}
	}

	aaa := 0
	for i := 0; i < len(files)-1; i++ {
		for a := 0; a < files[i].l; a++ {
			sum += files[i].ix * aaa
			aaa++
		}
		for _, file := range emptys[i].files {
			for a := 0; a < file.l; a++ {
				sum += file.ix * aaa
				aaa++
			}
		}
		aaa += emptys[i].remaining
	}

	fmt.Println(sum)
}

func main() {
	f, err := os.Open("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	tr := io.TeeReader(f, buf)
	part1(tr)
	part2(buf)
}
