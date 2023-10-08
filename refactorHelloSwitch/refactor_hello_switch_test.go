package main

import "testing"

func TestRefactorHelloSwitch(t *testing.T) {
	t.Run("saying hello to people english", func(t *testing.T) {
		got := Hello("hn", "")
		want := "Hello, hn"

		assertionMessage(t, got, want)
	})

	t.Run("saying hello to people in french", func(t *testing.T) {
		got := Hello("hn", "French")
		want := "Bonjour, hn"

		assertionMessage(t, got, want)
	})

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("hn", "Spanish")
		want := "Hola, hn"

		assertionMessage(t, got, want)
	})

	t.Run("with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertionMessage(t, got, want)
	})

	t.Run("no name", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"

		assertionMessage(t, got, want)
	})

	t.Run("no language", func(t *testing.T) {
		got := Hello("hn", "")
		want := "Hello, hn"

		assertionMessage(t, got, want)
	})
}

func assertionMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
