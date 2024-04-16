package main

import (
	"bytes"
	"strings"
	"testing"
)

const (
	WrongChoiceError = "Wrong choice"
	GoodChoiceError  = "Good choice"
	ApacheInput      = "apache\n"
	BadInput         = "bad"
	Name             = "default"
)

var configAccepted = []string{
	ApacheInput,
	"ngnix\n",
}

func TestPresentChoices(t *testing.T) {
	function := PresentChoices()
	want := strings.Join(configAccepted, "")

	if function != want {
		t.Errorf(WrongChoiceError)
	}
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
	if !InArray(ApacheInput, choiceAccepted) {
		t.Errorf("Not in array")
	}
}

func TestNotInArray(t *testing.T) {
	if InArray(BadInput, choiceAccepted) {
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
