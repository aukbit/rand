package rand_test

import (
	"strings"
	"testing"

	"github.com/aukbit/rand"
)

func TestString(t *testing.T) {

	a := rand.String(10)
	if len(a) != 10 {
		t.Fatalf("string lenght is not correct")
	}

	b := rand.StringPrefix(10, "ok")
	if len(b) != 12 {
		t.Fatalf("string lenght is not correct")
	}
	if !strings.HasPrefix(b, "ok") {
		t.Fatalf("string does not have the prefix defined")
	}

	c := rand.StringSuffix(10, "ok")
	if len(b) != 12 {
		t.Fatalf("string lenght is not correct")
	}
	if !strings.HasSuffix(c, "ok") {
		t.Fatalf("string does not have the suffix defined")
	}

}
