package vup_test

import (
	"testing"

	"github.com/leocavalcante/vup/internal/vup"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v, err := vup.NewVersion("v1.2.3")
		assert.NoError(t, err)
		assert.Equal(t, "v1.2.3", v.String())
	})

	t.Run("invalid major", func(t *testing.T) {
		_, err := vup.NewVersion("a.2.3")
		assert.Error(t, err)
	})

	t.Run("invalid minor", func(t *testing.T) {
		_, err := vup.NewVersion("1.a.3")
		assert.Error(t, err)
	})

	t.Run("invalid patch", func(t *testing.T) {
		_, err := vup.NewVersion("1.2.a")
		assert.Error(t, err)
	})
}
