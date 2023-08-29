package main

import "testing"

func TestRefactorHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John"

		assertionMessage(t, got, want)
	})
	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertionMessage(t, got, want)
	})

	t.Run("Hello in different language-Spanish", func(t *testing.T) {
		got := Hello("Heta", "Spanish")
		want := "Hola, Heta"

		assertionMessage(t, got, want)
	})

	t.Run("hello in different langaugae-French", func(t *testing.T) {
		got := Hello("Test", "French")
		want := "Bonjour, Test"

		assertionMessage(t, got, want)
	})

}

func assertionMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

//no name

//no string
