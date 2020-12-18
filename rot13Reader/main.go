package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rd *rot13Reader) Read(b []byte) (read int, err error) {
	read, err = rd.r.Read(b)
	for i := 0; i < len(b); i++ {
		char := b[i]
		if (char >= 'a' && char <= 'm') || (char >= 'A' && char <= 'M') {
			b[i] += 13
		} else if (char >= 'n' && char <= 'z') || (char >= 'N' && char <= 'Z') {
			b[i] -= 13
		}
	}
	fmt.Println(read)
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
