package calc

import "testing"

func TestAdd(t *testing.T) {
	a := 3
	b := 2
	want := a + b
	if got := Add(a, b); got != want {
		t.Errorf("Add(%d, %d) = %q, want %q", a, b, got, want)
	}
}
