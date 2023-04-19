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

func (r rot13Reader) Read(b []byte) (int, error) {
	in := make([]byte, cap(b))
	c, err := r.r.Read(in)
	for i, bt := range in {
		ibt := int32(bt)
		if ibt >= int32('a') && ibt <= int32('z') {
			ch := ibt + 13
			if ch > int32('z') {
				ch = int32('a') + (ch - int32('z'))
			}
			b[i] = byte(ch)
		} else if ibt >= int32('A') && ibt <= int32('Z') {
			ch := ibt + 13
			if ch > int32('Z') {
				ch = int32('A') + (ch - int32('Z'))
			}
			b[i] = byte(ch)
		} else {
			b[i] = bt
		}
	}
	if err != nil {
		return c, err
	}
	return c, nil
}

func mainTest() {
	i := int32(' ')
	fmt.Println(i)
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
