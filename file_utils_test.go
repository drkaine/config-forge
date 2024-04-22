package main

import (
	"config-forge/configs"
	"config-forge/configurator"
	"config-forge/utils"
	"os"
	"testing"
)

func TestEditFile(t *testing.T) {
	informations := map[string]string{
		"name":         configs.NAME_TEST,
		"serverName":   configs.SERVER_NAME_TEST,
		"documentRoot": configs.DOCUMENT_ROOT_TEST,
	}

	config := configurator.ConfigBuilder(configs.APACHE_INPUT, informations)

	err := utils.EditFile(config)
	if err != nil {
		t.Errorf("Error on the edition of the file : %v", err)
		return
	}

	if _, err := os.Stat(configs.NAME_FILE_TEST); os.IsNotExist(err) {
		t.Errorf("The file don't exist : %v", err)
	}

	content, err := os.ReadFile(config.GetName())
	if err != nil {
		t.Errorf("Error for read the file : %v", err)
		return
	}

	if string(content) != config.CustomiseConfigContent() {
		t.Errorf("File content is incorrect. Wanted: %s, Obtained : %s", config.CustomiseConfigContent(), string(content))
	}

	defer func() {
		if err := os.Remove(config.GetName()); err != nil {
			t.Errorf("Error on the deletion of the file: %v", err)
		}
	}()
}
