package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetModName(t *testing.T) {
	name := getModName(".")
	assert.Equal(t, "tidb-wrapper", name)
}
