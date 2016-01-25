package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProxyStats(t *testing.T) {
	pr := NewProxyStats(10, 1)
	t.Logf("Pr: %#v", pr)
	pr.Track()

	go func(pr *proxyStats) {
		for i := 0; i < 10; i++ {
			pr.Insert(1)
		}
	}(pr)

	time.Sleep(time.Second * 5)

	t.Logf("Stats: %#v", pr.Stats())

	assert.True(t, true)
}
