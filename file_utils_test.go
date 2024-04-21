package main

import (
	"config-forge/configs"
	"config-forge/configurator"
	"config-forge/utils"
	"os"
	"testing"
)

func TestEditFile(t *testing.T) {
	config := configurator.InstanciationConfig(configs.NAME_TEST, configs.SERVER_NAME_TEST, configs.DOCUMENT_ROOT_TEST)

	err := utils.EditFile(config)
	if err != nil {
		t.Errorf("Error on the edition of the file : %v", err)
		return
	}

	if _, err := os.Stat(configs.NAME_FILE_TEST); os.IsNotExist(err) {
		t.Errorf("The file don't exist : %v", err)
	}

	content, err := os.ReadFile(config.Name)
	if err != nil {
		t.Errorf("Error for read the file : %v", err)
		return
	}

	if string(content) != config.FileContent {
		t.Errorf("File content is incorrect. Wanted: %s, Obtained : %s", config.FileContent, string(content))
	}

	defer func() {
		if err := os.Remove(config.Name); err != nil {
			t.Errorf("Error on the deletion of the file: %v", err)
		}
	}()
}
