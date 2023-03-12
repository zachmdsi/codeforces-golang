package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func P4A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var w, t int
	Fscan(in, &w)
	if w % 2 == 0 {
		t = 0
		exit:
		for i := w/2; i <= w-2; i++ {
			for j := w/2; j >= 2; j-- {
				if i%2 == 0 && j%2 == 0 {
					t = 1
					break exit 
				}
			}
		}
	}
	if t == 1 {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

func main() {
	P4A(os.Stdin, os.Stdout)
}
