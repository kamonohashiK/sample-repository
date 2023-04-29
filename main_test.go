package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello World!"
	actual := Hello()

	if actual != expected {
		t.Errorf("Expected: %s, but got: %s", expected, actual)
	}
}
