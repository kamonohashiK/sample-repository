package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Actions書き換え完了！"
	actual := Hello()

	if actual != expected {
		t.Errorf("Expected: %s, but got: %s", expected, actual)
	}
}

func TestNewHello(t *testing.T) {
	expected := "Hello New World!"
	actual := NewHello()

	if actual != expected {
		t.Errorf("Expected: %s, but got: %s", expected, actual)
	}
}
