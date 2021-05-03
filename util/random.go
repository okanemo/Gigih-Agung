package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	// if a constant seed is set, it will output the same number every time
	// we need to make it a variable seed which changes after each call
	// using time is a way to do it, hence the code below:
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomString(n int) string {

	const alphabet = "abcdefghijknlmnopqrstuvwxyz"
	k := len(alphabet)

	var sb strings.Builder

	for i := 0; i < n; i++ {
		char := alphabet[rand.Intn(k)]
		sb.WriteByte(char)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(5)
}

func RandomLineID() string {
	return RandomString(7)
}
