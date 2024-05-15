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

	// each letter will substituted  with the letter
	// in position = letterPos + shift % modulo
	shift  int
	modulo int
}

func New(in io.Reader, out io.Writer) *Caesar {
	return &Caesar{
		in:     in,
		out:    out,
		shift:  1,
		modulo: 95,
	}
}

func (c *Caesar) Encrypt() error {
	scanner := bufio.NewScanner(c.in)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		letterCode := scanner.Bytes()[0]
		encryptedLetter := encryptCode(letterCode, c.shift, c.modulo)
		fmt.Fprintf(c.out, encryptedLetter)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
	return nil
}

func encryptCode(code byte, shift, mod int) string {
	encryptedCode := code + byte(shift)%byte(mod)
	return string(encryptedCode)
}

func (c *Caesar) Decrypt() error {
	scanner := bufio.NewScanner(c.in)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		letterCode := scanner.Bytes()[0]
		decryptedLetter := decryptCode(letterCode, c.shift, c.modulo)
		fmt.Fprintf(c.out, decryptedLetter)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
	return nil
}

func decryptCode(code byte, shift, mod int) string {
	decryptedCode := code - byte(shift)%byte(mod)
	return string(decryptedCode)
}
