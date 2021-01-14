// +build unit

package main_test

import "testing"

func TestUnit(t *testing.T) {
	src := 2
	tgt := 2
	if src != tgt {
		t.Fatalf("expected %d, got %d", src, tgt)
	}
}
