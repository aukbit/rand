// Package rand contains helper random generator functions to be used around
package rand

import (
	"math/rand"
	"time"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// String returns a random string with n length
// source http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func String(length int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// StringPrefix returns random string with chosen prefix
func StringPrefix(length int, prefix string) string {
	return prefix + String(length)
}

// StringSuffix returns random string with chosen sufix
func StringSuffix(length int, suffix string) string {
	return String(length) + suffix
}
