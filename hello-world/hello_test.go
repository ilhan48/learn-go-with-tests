package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ilhan")
	want := "Hello, Ilhan!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
