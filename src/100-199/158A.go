package main

import (
	"bufio"
	. "fmt"
	"io"
)

func P158A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, a, b int
	Fscan(in, &n, &k)
	for i := 0; i < k; i++ {
		Fscan(in, &a)
		if a == 0 {
			Fprint(out, i)
			return
		}
	}
	for i := k; i < n; i++ {
		Fscan(in, &b)
		if a != b {
			Fprint(out, i)
			return
		}
	}
	Fprint(out, n)
}

/*func main() {
	P158A(os.Stdin, os.Stdout)
}*/