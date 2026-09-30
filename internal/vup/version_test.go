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

	t.Run("rc0 is kept", func(t *testing.T) {
		v, err := vup.NewVersion("1.2.3-rc0")
		assert.NoError(t, err)
		assert.Equal(t, "1.2.3-rc0", v.String())
	})

	t.Run("rejects extra suffix", func(t *testing.T) {
		_, err := vup.NewVersion("1.2.3-rc1-extra")
		assert.Error(t, err)
	})

	t.Run("rejects empty suffix", func(t *testing.T) {
		_, err := vup.NewVersion("1.2.3-")
		assert.Error(t, err)
	})

	t.Run("rejects non-canonical rc", func(t *testing.T) {
		_, err := vup.NewVersion("1.2.3-rc01")
		assert.Error(t, err)
	})

	t.Run("rejects extra core component", func(t *testing.T) {
		_, err := vup.NewVersion("1.2.3.4-rc1")
		assert.Error(t, err)
	})
}
