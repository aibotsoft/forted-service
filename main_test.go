package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple(t *testing.T) {
	got := Simple()
	assert.Equal(t, 1, got)
}
