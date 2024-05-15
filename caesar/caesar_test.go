package caesar

import (
	"bytes"
	"strings"
	"testing"
)

// TODO: Write tests for larger keys

func TestEncrypt(t *testing.T) {
	tests := []struct {
		input string
		key   int
		want  string
	}{
		{input: "Hello, World!", key: 3, want: "Khoor/#Zruog$"},
		{input: "can you crack the code?", key: 1, want: "dbo!zpv!dsbdl!uif!dpef@"},
	}

	for _, tc := range tests {
		in := strings.NewReader(tc.input)
		out := new(bytes.Buffer)
		c := New(in, out, tc.key)
		err := c.Encrypt()
		if err != nil {
			t.Errorf("Encrypt() error: %v", err)
		}
		output := out.String()
		if output != tc.want {
			t.Errorf("Encrypt() = %v, want %v", output, tc.want)
			t.Errorf("Encrypt:\n\tEncrypting:%v\n\tGot:%v\nWant:%v", in, output, tc.want)
		}
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		input string
		key   int
		want  string
	}{
		{input: "Khoor/#Zruog$", key: 3, want: "Hello, World!"},
		{input: "dbo!zpv!dsbdl!uif!dpef@", key: 1, want: "can you crack the code?"},
	}

	for _, tc := range tests {
		in := strings.NewReader(tc.input)
		out := new(bytes.Buffer)

		c := New(in, out, tc.key)

		err := c.Decrypt()
		if err != nil {
			t.Errorf("Decrypt() error: %v", err)
		}
		output := out.String()
		if output != tc.want {
			t.Errorf("Decrypt() = %v, want %v", output, tc.want)
		}
	}
}
