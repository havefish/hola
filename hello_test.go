package hello

import "testing"

func TestHello(t *testing.T) {
	want := ""
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
