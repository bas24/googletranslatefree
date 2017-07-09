package translategooglefree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeURI(t *testing.T) {
	test := `Just test string.`
	must := `Just%20test%20string.`
	result, _ := encodeURI(test)
	assert.Equal(t, must, result, "Must be equal.")
}
