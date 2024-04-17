package main

import (
	"bytes"
	"testing"
)

const (
	WrongChoiceError = "Wrong choice"
	GoodChoiceError  = "Good choice"
	ApacheInput      = "apache"
	BadInput         = "bad"
	Name             = "default-test.conf"
	ServerName       = "test.net"
	DocumentRoot     = "/var/www/monsite.com/public_html"
)

var configAccepted = []string{
	ApacheInput,
	"ngnix\n",
}

func TestWrongAnalyseResponse(t *testing.T) {
	function := AnalyseResponse(BadInput)

	if function != WrongChoice {
		t.Errorf(GoodChoiceError)
	}
}

func TestGoodAnalyseResponse(t *testing.T) {
	function := AnalyseResponse(ApacheInput)

	if function != PrepareConfiguration {
		t.Errorf(WrongChoiceError)
	}
}

func TestListeningResponse(t *testing.T) {
	reader := bytes.NewReader([]byte(ApacheInput))

	response := ListeningResponse(reader)

	if response != ApacheInput {
		t.Errorf("The response %q isn't the same of the input %q", ApacheInput, response)
	}
}

func TestInArray(t *testing.T) {
	if !InArray(ApacheInput, configAccepted) {
		t.Errorf("Not in array")
	}
}

func TestNotInArray(t *testing.T) {
	if InArray(BadInput, configAccepted) {
		t.Errorf("In array")
	}
}

func TestStructureConfiguration(t *testing.T) {
	config := apache{
		name:         Name,
		serverName:   ServerName,
		documentRoot: DocumentRoot,
	}

	if config.name != Name {
		t.Errorf("Error on name struct apache")
	}

	if config.serverName != ServerName {
		t.Errorf("Error on serverName struct apache")
	}

	if config.documentRoot != DocumentRoot {
		t.Errorf("Error on documentRoot struct apache")
	}
}

func TestInstanciationConfig(t *testing.T) {
	config := InstanciationConfig(Name, ServerName, DocumentRoot)

	if config.name != Name {
		t.Errorf("The name of the config is %q and need to be %q", config.name, Name)
	}
}
