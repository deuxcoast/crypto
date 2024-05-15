package caesar

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TODO: Write tests for larger keys

func TestEncrypt(t *testing.T) {
	tests := []struct {
		input string
		key   int
		want  []byte
	}{
		{input: "Hello, World!", key: 3, want: []byte("Khoor/#Zruog$")},
		{input: "can you crack the code?", key: 1, want: []byte("dbo!zpv!dsbdl!uif!dpef@")},
		{input: "aaa", key: 55, want: []byte("999")},
	}

	for _, tc := range tests {
		in := strings.NewReader(tc.input)
		out := new(bytes.Buffer)
		c := New(in, out, tc.key)
		err := c.Encrypt()
		if err != nil {
			t.Errorf("Encrypt() error: %v", err)
		}
		output := out
		if !cmp.Equal(output.Bytes(), tc.want) {
			t.Errorf("Encrypt:\n\tEncrypting: %v\n\tGot:\t%v\n\tWant:\t%v", in, output.Bytes(), tc.want)
		}
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		input string
		key   int
		want  []byte
	}{
		{input: "Khoor/#Zruog$", key: 3, want: []byte("Hello, World!")},
		{input: "dbo!zpv!dsbdl!uif!dpef@", key: 1, want: []byte("can you crack the code?")},
	}

	for _, tc := range tests {
		in := strings.NewReader(tc.input)
		out := new(bytes.Buffer)

		c := New(in, out, tc.key)

		err := c.Decrypt()
		if err != nil {
			t.Errorf("Decrypt() error: %v", err)
		}
		output := out
		if !cmp.Equal(output.Bytes(), tc.want) {
			t.Errorf("Decrypt:\n\tDecrypting: %v\n\tGot:\t%v\n\tWant:\t%v", in, output.Bytes(), tc.want)
		}
	}
}
