package main

import "testing"

// Testing function
func TestHello(t *testing.T) {
	// Add a test case
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gabriel")
		want := "Hello Gabriel"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	// Add a test case
	t.Run("Say Hello World!", func(t *testing.T) {
		got := Hello("World!")
		want := "Hello World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add two positive numbers", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Add a positve and a negative number", func(t *testing.T) {
		got := Add(2, -5)
		want := -3

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestMultiply(t *testing.T) {
	t.Run("Multiply two positive numbers", func(t *testing.T) {
		got := Multiply(2, 2)
		want := 4

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Multiply a positve and a negative number", func(t *testing.T) {
		got := Multiply(2, -5)
		want := -10

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
