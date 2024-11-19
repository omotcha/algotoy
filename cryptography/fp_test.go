package cryptography

import (
	"fmt"
	"testing"
)

func TestFpInverse(t *testing.T) {
	x := 4
	p := 23
	y := fpInverse(x, p)
	fmt.Print(y)
}
