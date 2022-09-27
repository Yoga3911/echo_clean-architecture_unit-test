package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	responseModel := ResponseModel{
		Data: "OKE",
		Message: "Coba",
		Status: true,
	}

	assert.Equal(t, "OKE", responseModel.Data)
	assert.Equal(t, "Coba", responseModel.Message)
	assert.Equal(t, true, responseModel.Status)
}
