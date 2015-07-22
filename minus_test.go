package minus

import (
	"testing"
)

func TestMinus(t *testing.T) {
	if Minus(2, 1) == 1 {
		t.Fatal("Minus test fail")
	} else {

		t.Log("Minus test passed")
	}
}
