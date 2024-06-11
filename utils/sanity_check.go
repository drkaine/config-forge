package utils

import (
	"config-forge/configs"
)

func CheckArguments(arguments []string) string {
	if !InArray(arguments[2], configs.ConfigAccepted) {
		return configs.WRONG_CHOICE_RETURN
	}

	if len(arguments) < 6 {
		return configs.MINIMUM_ARGUMENTS_RETURN
	}

	if len(arguments) > 6 {
		return configs.MAXIMUM_ARGUMENTS_RETURN
	}

	if len(arguments[3]) < 2 {
		return configs.MINIMUM_FILE_NAME_RETURN
	}

	if len(arguments[4]) < 4 {
		return configs.MINIMUM_SERVER_NAME_RETURN
	}

	if len(arguments[5]) < 2 {
		return configs.MINIMUM_DOCUMENT_ROOT_RETURN
	}

	return configs.SANITY_CHECK_PASSED
}
