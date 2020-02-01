package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' and want '%q'.", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Vitor", "english")
		const want string = "Hello Vitor"
		assertMessage(t, got, want)
	})

	t.Run("Say 'Hello World' when an empty string is supplied.", func(t *testing.T) {
		got := Hello("", "english")
		const want string = "Hello World"
		assertMessage(t, got, want)
	})

	t.Run("Saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Vitor", "Spanish")
		const want string = "Hola Vitor"
		assertMessage(t, got, want)
	})

	t.Run("Saying hello to people in French", func(t *testing.T) {
		got := Hello("Vitor", "French")
		const want string = "Bonjour Vitor"
		assertMessage(t, got, want)
	})

}
