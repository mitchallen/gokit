package gokit

import "testing"

func TestInfo(t *testing.T) {
	want := "gokit info"
	if got := Info(); got != want {
		t.Errorf("Info() = %q, want %q", got, want)
	}
}
