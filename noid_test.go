package noid

import (
	"testing"
)

func TestNoid(t *testing.T) {
	n, err := NewNoid(".r2dk")
	if err != nil {
		t.Errorf("Got error %v\n", err)
	}
	ids := []string{
		"00", "66", "11", "77", "22", "88", "33", "99", "44", "55",
	}
	for _, expected := range ids {
		z := n.Mint()
		if z != expected {
			t.Errorf("%v != %v\n", z, expected)
		}
	}
}
