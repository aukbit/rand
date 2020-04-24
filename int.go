// Package rand contains helper random generator functions to be used around
package rand

import (
	"math/rand"
	"strconv"
	"time"
)

const numberBytes = "123456789"
const (
	numberIdxBits = 6                    // 6 bits to represent a letter index
	numberIdxMask = 1<<numberIdxBits - 1 // All 1-bits, as many as letterIdxBits
	numberIdxMax  = 63 / numberIdxBits   // # of letter indices fitting in 63 bits
)

// Int returns a random int with n digits
func Int(digits int) int {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, digits)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := digits-1, src.Int63(), numberIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), numberIdxMax
		}
		if idx := int(cache & numberIdxMask); idx < len(numberBytes) {
			b[i] = numberBytes[idx]
			i--
		}
		cache >>= numberIdxBits
		remain--
	}
	i, _ := strconv.Atoi(string(b))
	return i
}
