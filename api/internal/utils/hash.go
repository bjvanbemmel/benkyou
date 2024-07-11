package utils

import (
	"crypto/sha256"
	"fmt"
)

func Hash[T any](t T) string {
	raw := []byte(fmt.Sprintf("%v", t))

	return fmt.Sprintf("%x", sha256.Sum256(raw))
}
