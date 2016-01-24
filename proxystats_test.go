package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProxyStats(t *testing.T) {
	pr := NewProxyStats(10, 1)
	t.Logf("Pr: %#v", pr)
	pr.Track()

	assert.True(t, true)
}
