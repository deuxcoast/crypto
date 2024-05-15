package caesar

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Caesar struct {
	in  io.Reader
	out io.Writer
}

func New(in io.Reader, out io.Writer) *Caesar {
	return &Caesar{
		in:  in,
		out: out,
	}
}

func (c *Caesar) Encrypt() error {
	scanner := bufio.NewScanner(c.in)
	for scanner.Scan() {
		fmt.Fprintln(c.out, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
	return nil
}
