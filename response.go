package main

import (
	"bufio"
	"io"
	"strings"
)

func ListeningResponse(reader io.Reader) string {
	bufReader := bufio.NewReader(reader)
	response, _ := bufReader.ReadString('\n')
	return strings.TrimSpace(response)
}

func AnalyseResponse(choice string) string {
	if InArray(choice, configAccepted) {
		return PrepareConfiguration
	}
	return WrongChoice
}
