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
	shift     int
	modulo    int
	startChar int
	endChar   int
}

func New(in io.Reader, out io.Writer, shift int) *Caesar {
	return &Caesar{
		in:    in,
		out:   out,
		shift: shift,
		// range of ascii codes including all [A-Za-z] letters and major
		// special characters, ranging from <space> to ~
		modulo:    95,
		startChar: 32,
		endChar:   126,
	}
}

func (c *Caesar) Encrypt() error {
	scanner := bufio.NewScanner(c.in)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		letterCode := scanner.Bytes()[0]
		encryptedLetter := c.encryptCode(letterCode)
		fmt.Fprintf(c.out, encryptedLetter)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
	return nil
}

func (c *Caesar) encryptCode(code byte) string {
	codeVal := code - byte(c.startChar)
	encryptedCode := ((codeVal + byte(c.shift)) % byte(c.modulo)) + byte(c.startChar)
	return string(encryptedCode)
}

func (c *Caesar) Decrypt() error {
	scanner := bufio.NewScanner(c.in)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		letterCode := scanner.Bytes()[0]
		decryptedLetter := c.decryptCode(letterCode)
		fmt.Fprintf(c.out, decryptedLetter)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
	return nil
}

func (c *Caesar) decryptCode(code byte) string {
	codeVal := code - byte(c.startChar)
	decryptedCode := (codeVal-byte(c.shift))%byte(c.modulo) + byte(c.startChar)
	return string(decryptedCode)
}
