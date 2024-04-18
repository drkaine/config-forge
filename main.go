package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var configAccepted = []string{
	"apache",
	"ngnix",
}

const (
	Hello                = "Hello, choose your tools at configure : "
	PrepareConfiguration = "Now go to prepare the configuration"
	WrongChoice          = "Uncorrect choice !"
	AskNameFile          = "How would you name the configuration file ? :"
	AskServerName        = "Wath is the name of the server ? :"
	AskDocumentRoot      = "What is the path to the document root of the project ? :"
	ApachePath           = "/etc/apache2/sites-available/"
)

func main() {
	Runner()
}

func Runner() {
	fmt.Println(Hello)
	fmt.Println(strings.Join(configAccepted, " "))
	answer := ListeningResponse(os.Stdin)
	analyseResponse := AnalyseResponse(answer)
	fmt.Println(analyseResponse)

	if analyseResponse == WrongChoice {
		Runner()
	}

	fmt.Println(AskNameFile)

	name := ListeningResponse(os.Stdin)

	fmt.Println(AskServerName)

	serverName := ListeningResponse(os.Stdin)

	fmt.Println(AskDocumentRoot)

	documentRoot := ListeningResponse(os.Stdin)

	config := InstanciationConfig(name, serverName, documentRoot)

	fmt.Println(config.name)

	er := EditFile(config)
	fmt.Println(er)
}

func ListeningResponse(reader io.Reader) string {
	bufReader := bufio.NewReader(reader)
	response, _ := bufReader.ReadString('\n')
	return response
}

func AnalyseResponse(choice string) string {
	if InArray(choice, configAccepted) {
		return PrepareConfiguration
	}
	return WrongChoice
}

func InArray(search string, target []string) bool {
	for _, value := range target {
		if value == search {
			return true
		}
	}
	return false
}

type Apache struct {
	name         string
	serverName   string
	documentRoot string
	path         string
	fileContent  string
}

func InstanciationConfig(name string, serverName string, documentRoot string) Apache {
	config := Apache{
		name:         name,
		serverName:   serverName,
		documentRoot: documentRoot,
		path:         ApachePath,
		fileContent:  "Hello world",
	}

	return config
}

func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func EditFile(apache Apache) error {
	file, err := CreateFile(apache.name)

	if err != nil {
		return err
	}

	_, err = file.WriteString(apache.fileContent)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
