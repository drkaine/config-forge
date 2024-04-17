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
	Name             = "default"
	NameInput        = "default-test.conf"
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
		name: Name,
	}

	if config.name != Name {
		t.Errorf("Error on name struct")
	}
}

func TestInstanciationConfigurator(t *testing.T) {
	configurator := InstanciationConfigurator(NameInput)

	if configurator.name != NameInput {
		t.Errorf("The name of the configurator is %q and need to be %q", configurator.name, NameInput)
	}
}
