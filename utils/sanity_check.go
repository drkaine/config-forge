package utils

import (
	"config-forge/configs"
)

func CheckArguments(arguments []string) bool {
	if !InArray(arguments[1], configs.ConfigAccepted) {
		return false
	}

	if len(arguments[2]) == 0 {
		return false
	}

	if len(arguments[3]) == 0 {
		return false
	}

	if len(arguments[4]) == 0 {
		return false
	}

	return true
}
