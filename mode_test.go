package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewModeList(t *testing.T) {
	assert.Equal(t, []Mode{3, 2, 1}, NewModeList(12301, 3))
	assert.Equal(t, []Mode{}, NewModeList(12301, 0))
	assert.Equal(t, []Mode{3, 2, 1, 0}, NewModeList(12301, 4))
}
