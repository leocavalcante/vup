package vup_test

import (
	"strconv"
	"testing"

	"github.com/leocavalcante/vup/internal/vup"
	"github.com/stretchr/testify/assert"
)

func TestMajor(t *testing.T) {
	t.Run("no prefix", func(t *testing.T) {
		m, err := vup.NewMajor("1")
		assert.NoError(t, err)
		assert.Equal(t, "1", m.String())
	})

	t.Run("with prefix", func(t *testing.T) {
		m, err := vup.NewMajor("v1")
		assert.NoError(t, err)
		assert.Equal(t, "v1", m.String())
	})

	t.Run("invalid", func(t *testing.T) {
		_, err := vup.NewMajor("a")
		assert.ErrorIs(t, err, strconv.ErrSyntax)
	})
}

func TestMajor_Inc(t *testing.T) {
	m, _ := vup.NewMajor("1")
	m.Inc(1)
	assert.Equal(t, "2", m.String())
}

func TestMajor_Dec(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		m, _ := vup.NewMajor("2")
		err := m.Dec(1)
		assert.NoError(t, err)
		assert.Equal(t, "1", m.String())
	})

	t.Run("error", func(t *testing.T) {
		m, _ := vup.NewMajor("0")
		err := m.Dec(1)
		assert.ErrorIs(t, err, vup.ErrMajorLessThanZero)
	})
}
