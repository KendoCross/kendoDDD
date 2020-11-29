package helper

import (
	mrand "math/rand"
)

var letter = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// RandString returns a fixed length random string.
func RandString(n int) string {
	//mrand.Seed(time.Now().UTC().UnixNano())
	// time.Sleep(time.Nanosecond)

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[mrand.Intn(len(letter))]
	}

	return string(b)
}
