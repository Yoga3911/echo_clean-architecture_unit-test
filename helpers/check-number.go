package helpers

import (
	"fmt"
	"strconv"
)

func IsNumber(data string) error {
	_, err := strconv.Atoi(data)
	if err != nil {
		return fmt.Errorf("path parameter hanya menerima angka")
	}

	return nil
}