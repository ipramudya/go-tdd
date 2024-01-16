package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("sayin hello to people", func(t *testing.T) {
		got := Hello("John")
		want := "Hello, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string default to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot %q \nwant %q", got, want)
	}
}
