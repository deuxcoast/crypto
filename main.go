package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
)

type App struct {
	in  io.Reader
	out io.Writer
	key int
}

func main() {
	var scheme string
	var output string
	var decrypt bool
	var key int

	flag.StringVar(&scheme, "scheme", CAESAR, "The crypto scheme to be used for encryption")
	flag.StringVar(&output, "output", "", "Where output will be directed")
	flag.BoolVar(&decrypt, "decrypt", false, "Decrypt the input using specified scheme and key")
	flag.IntVar(&key, "key", 0, "Key for encryption/decryption")
	flag.Parse()

	var in io.Reader
	in = os.Stdin

	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file: %v", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f

	}
	var out io.Writer
	out = os.Stdout

	if output != "" {
		path := path.Clean(output)
		f, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating output file: %v", err)
			os.Exit(1)
		}
		defer f.Close()

		out = f
	}

	app := &App{
		in:  in,
		out: out,
	}
	cipher, err := parseScheme(scheme, *app)
	if err != nil {
		fmt.Printf("Error parsing scheme: %v", err)
		os.Exit(1)
	}

	if decrypt {
		cipher.Decrypt()
	} else {
		cipher.Encrypt()
	}
}
