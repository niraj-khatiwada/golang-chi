package utils

import (
	nanoid "github.com/matoous/go-nanoid/v2"
	"strings"
)

func GenerateRandomUUID() (string, error) {
	return nanoid.New()
}

type AlphanumericUUID struct {
	length int
}

func (a *AlphanumericUUID) Generate() (string, error) {
	alphabets := ""
	for i := 'a'; i <= 'z'; i++ {
		alphabets += string(i)
	}

	var length int
	if a.length == 0 {
		length = 21
	}

	return nanoid.Generate(alphabets+strings.ToUpper(alphabets)+"0123456789", length)
}
