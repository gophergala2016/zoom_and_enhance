package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenID(t *testing.T) {
	id := genID()
	assert.IsType(t, "this is a string", id)

	id2 := genID()
	assert.NotEqual(t, id, id2)
}
