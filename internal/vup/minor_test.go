package vup_test

import (
	"strconv"
	"testing"

	"github.com/leocavalcante/vup/internal/vup"
	"github.com/stretchr/testify/assert"
)

func TestMinor(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m, err := vup.NewMinor("1")
		assert.NoError(t, err)
		assert.Equal(t, "1", m.String())
	})

	t.Run("invalid", func(t *testing.T) {
		_, err := vup.NewMinor("a")
		assert.ErrorIs(t, err, strconv.ErrSyntax)
	})
}

func TestMinor_Inc(t *testing.T) {
	m, _ := vup.NewMinor("1")
	m.Inc(1)
	assert.Equal(t, "2", m.String())
}

func TestMinor_Dec(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		m, _ := vup.NewMinor("2")
		err := m.Dec(1)
		assert.NoError(t, err)
		assert.Equal(t, "1", m.String())
	})

	t.Run("error", func(t *testing.T) {
		m, _ := vup.NewMinor("1")
		err := m.Dec(1)
		assert.ErrorIs(t, err, vup.ErrMinorLessThanOne)
	})
}
