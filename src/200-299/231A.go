package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func P231A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, a, b, c, t, f int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &a, &b, &c)
		t = a + b + c
		if t > 1 {
			f++
		}
	}
	Fprint(out, f)
}

func main() {
	P231A(os.Stdin, os.Stdout)
}