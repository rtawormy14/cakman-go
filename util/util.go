package util

import (
	"crypto/sha256"
)

// Hash is ...
func Hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return string(h.Sum(nil))
}
