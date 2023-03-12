package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func P71A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var (
		n, l int
		s string
	)
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		var b strings.Builder
		Fscan(in, &s)
		l = len(s)
		if l > 10 {
			b.WriteByte(s[0])
			b.WriteString(strconv.Itoa(l-2))
			b.WriteByte(s[l-1])
		} else {
			b.WriteString(s)
		}
		Fprintln(out, b.String())
	}
}

func main() { 
	P71A(os.Stdin, os.Stdout)
}