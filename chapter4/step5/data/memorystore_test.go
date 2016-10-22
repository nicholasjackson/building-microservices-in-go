package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturns1KittenWhenSearchGarfield(t *testing.T) {
	store := MemoryStore{}
	kittens := store.Search("Garfield")

	assert.Equal(t, 1, len(kittens))
}

func TestReturns0KittenWhenSearchTom(t *testing.T) {
	store := MemoryStore{}
	kittens := store.Search("Tom")

	assert.Equal(t, 0, len(kittens))
}
