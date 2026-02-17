package main

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestDiv(t *testing.T) {
	a := gofakeit.Float64Range(1, 5)
	b := gofakeit.Float64Range(1, 5)

	require.Equal(t, a/b, Div(a, b))
	// assert.Equal(t, a/b, Div(a, b))

	t.Run("10/5", func(t *testing.T) {
		a := 10.0
		b := 5.0

		expected := 2.0

		require.Equal(t, expected, Div(a, b))
	})
	t.Run("10/0", func(t *testing.T) {
		a := 10.0
		b := 0.0

		expected := 0.0

		require.Equal(t, expected, Div(a, b))
	})
	// if a/b != Div(a, b) {
	// 	t.Errorf("Div(%f, %f) = %f; want %f", a, b, Div(a, b), a/b)
	// }
}
