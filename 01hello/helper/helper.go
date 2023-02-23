package helper

import (
	"strings"
)

func Validation(par1 string, par2 string) (bool, bool) {
	validUsername := len(par1) >= 2
	validEmail := strings.Contains(par2, "@")
	return validUsername, validEmail
}
