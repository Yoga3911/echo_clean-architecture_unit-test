package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	err := IsNumber("123")
	if err != nil {
		t.Error("Not a number")
		t.Fail()
	}

	assert.Nil(t, err)
}
