package utils

import "math/rand"

// Generates a string of random characters.
// Includes letters, numbers and an array of symbols.
func RandomString() string {
	limit := 12
	chars := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#%^&*()-=\\_+|[]{};':\",./<>?")

	key := make([]rune, limit)
	for i := range key {
		key[i] = chars[rand.Intn(len(chars))]
	}

	return string(key)
}
