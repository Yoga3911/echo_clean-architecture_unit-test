package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber_Success(t *testing.T) {
	err := IsNumber("123")

	assert.Nil(t, err)
}

func TestIsNumber_Failure(t *testing.T) {
	err := IsNumber("12dwq3")

	assert.NotNil(t, err)
}
