package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if go := Hello(); got != want {
		t.Errorf("Hellow() = %q, want %q", got, want)
	}
}
