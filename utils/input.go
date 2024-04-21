package utils

import (
	"bufio"
	"config-forge/configs"
	"io"
	"strings"
)

func ListeningResponse(reader io.Reader) string {
	bufReader := bufio.NewReader(reader)
	response, _ := bufReader.ReadString('\n')
	return strings.TrimSpace(response)
}

func AnalyseResponse(choice string) string {
	if InArray(choice, configs.ConfigAccepted) {
		return configs.PREPARE_CONFIGURATION_RETURN
	}
	return configs.WRONG_CHOICE_RETURN
}
