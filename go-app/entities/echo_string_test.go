package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEchoString(t *testing.T) {
	t.Run("empty init", func(t *testing.T) {
		_, got := NewEchoString("")
		want := "invalid parameters"

		assert.EqualError(t, got, want)

	})

	t.Run("get echo string", func(t *testing.T) {
		echo_string, err := NewEchoString("Hello, World")

		if err != nil {
			t.Fatal("expected an error")
		}

		got, err := echo_string.GetEchoString()
		want := "Hello, World"

		assert.Equal(t, got, want)

	})
}
