package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world!"

	if got != want {
		t.Errorf("Hello: got %q want %q", got, want)
	}
}
