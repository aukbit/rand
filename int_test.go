package rand_test

import (
	"strconv"
	"testing"

	"github.com/aukbit/rand"
)

func TestInt(t *testing.T) {

	a := rand.Int(5)
	if len(strconv.Itoa(a)) < 5 {
		t.Fatalf("total number of digits is not correct")
	}

}
