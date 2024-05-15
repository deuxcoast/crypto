package main

import (
	"errors"
	"fmt"

	"github.com/duexcoast/crypto/caesar"
)

const (
	CAESAR = "caesar"
	AFFINE = "affine"
)

var cipherScheme = map[string]string{
	CAESAR: "caesar",
	AFFINE: "affine",
}

func parseScheme(cipher string, app App) (Scheme, error) {
	switch cipher {
	case CAESAR:
		return caesar.New(app.in, app.out), nil
	default:
		return nil, errors.New(fmt.Sprintf("Scheme is not implemented: %s", cipher))
	}
}

type Scheme interface {
	Encrypt() error
}
