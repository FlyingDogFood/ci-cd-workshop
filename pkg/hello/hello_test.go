package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}

func TestGreetingPrefix(t *testing.T) {
	t.Run("test spanish", func(t *testing.T) {
		got := greetingPrefix(spanish)
		want := spanishHelloPrefix
		assertCorrectMessage(t, got, want)
	})

	t.Run("test french", func(t *testing.T) {
		got := greetingPrefix(french)
		want := frenchHelloPrefix
		assertCorrectMessage(t, got, want)
	})

	t.Run("test default", func(t *testing.T) {
		got := greetingPrefix("french")
		want := englishHelloPrefix
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
