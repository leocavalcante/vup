package vup_test

import (
	"strconv"
	"testing"

	"github.com/leocavalcante/vup/internal/vup"
	"github.com/stretchr/testify/assert"
)

func TestPatch(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		p, err := vup.NewPatch("1")
		assert.NoError(t, err)
		assert.Equal(t, "1", p.String())
	})

	t.Run("invalid", func(t *testing.T) {
		_, err := vup.NewPatch("a")
		assert.ErrorIs(t, err, strconv.ErrSyntax)
	})
}

func TestPatch_Inc(t *testing.T) {
	p, _ := vup.NewPatch("1")
	p.Inc(1)
	assert.Equal(t, "2", p.String())
}

func TestPatch_Dec(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		p, _ := vup.NewPatch("2")
		err := p.Dec(1)
		assert.NoError(t, err)
		assert.Equal(t, "1", p.String())
	})

	t.Run("error", func(t *testing.T) {
		p, _ := vup.NewPatch("0")
		err := p.Dec(1)
		assert.ErrorIs(t, err, vup.ErrPatchLessThanZero)
	})
}
