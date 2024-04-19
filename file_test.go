package main

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	tmpFileName := "testfile.txt"

	file, err := CreateFile(tmpFileName)

	if err != nil {
		t.Errorf("Error on the creation of the file : %v", err)
	}

	if file == nil {
		t.Errorf("Error on the creation of the file : %v", err)
	}

	if _, err := os.Stat(tmpFileName); os.IsNotExist(err) {
		t.Errorf("The file don't exist : %v", err)
	}

	defer file.Close()

	defer func() {
		if err := os.Remove(tmpFileName); err != nil {
			t.Errorf("Error on the delete of the file: %v", err)
		}
	}()
}

func TestEditFile(t *testing.T) {
	config := InstanciationConfig(Name, ServerName, DocumentRoot)

	err := EditFile(config)
	if err != nil {
		t.Errorf("Erreur lors de la lecture du fichier : %v", err)
		return
	}

	content, err := os.ReadFile(config.name)
	if err != nil {
		t.Errorf("Erreur lors de la lecture du fichier : %v", err)
		return
	}

	if string(content) != config.fileContent {
		t.Errorf("Contenu du fichier incorrect. Attendu: %s, Obtenu: %s", config.fileContent, string(content))
	}

	defer func() {
		if err := os.Remove(config.name); err != nil {
			t.Errorf("Error on the delete of the file: %v", err)
		}
	}()
}
