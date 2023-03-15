package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func P50A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var m, n, t int
	Fscan(in, &m, &n)
	t = m * n / 2	
	Fprint(out, t)
}

func main() {
	P50A(os.Stdin, os.Stdout)
}
