package main

import "github.com/stretchr/testify/assert"
import "testing"

func alpha() bool {
	return true
}

func beta(s string, i int) {
	// do nothing
}

func delta() int {
	return 0
}

func TestItAllowsPatchingOfFunctions(t *testing.T) {
	Patch(alpha, func() bool {
		return false
	})

	assert.True(t, alpha())

	Patch(delta, func() int {
		return 10
	})

	assert.Equal(t, 10, delta())
}

func TestItAllowsUnpatchingOfFunctions(t *testing.T) {
	handle := Patch(alpha, func() bool {
		return false
	})

	assert.False(t, alpha())

	handle.Unpatch()

	assert.True(t, alpha())
}

func TestItPassesArgsToThePatch(t *testing.T) {
	var passedString string
	var passedInt int

	Patch(beta, func(s string, i int) {
		passedString = s
		passedInt = i
	})

	beta("asdf", 100)
	assert.Equal(t, "asdf", passedString)
	assert.Equal(t, 100, passedInt)
}
