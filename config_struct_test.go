package main

import (
	"config-forge/configs"
	"config-forge/configurator"
	"testing"
)

func TestInstanciationConfig(t *testing.T) {
	configu := configurator.InstanciationConfig(configs.NAME_TEST, configs.SERVER_NAME_TEST, configs.DOCUMENT_ROOT_TEST)

	if configu.Name != configs.NAME_FILE_TEST {
		t.Errorf("The name of the config is %q and need to be %q", configu.Name, configs.NAME_TEST)
	}

	if configu.ServerName != configs.SERVER_NAME_TEST {
		t.Errorf("Error on serverName struct apache")
	}

	if configu.DocumentRoot != configs.DOCUMENT_ROOT_TEST {
		t.Errorf("Error on documentRoot struct apache")
	}

	if configu.Path != configs.APACHE_PATH_REPOSITORY {
		t.Errorf("Error on documentRoot struct apache")
	}

	if configu.FileContent != configs.APACHE_CONFIG_CONTENT {
		t.Errorf("Error on documentRoot struct apache")
	}
}

func TestCustomiseConfigContent(t *testing.T) {
	configu := configurator.Apache{
		Name:         configs.NAME_TEST,
		ServerName:   configs.SERVER_NAME_TEST,
		DocumentRoot: configs.DOCUMENT_ROOT_TEST,
		Path:         configs.APACHE_PATH_REPOSITORY,
		FileContent:  configs.APACHE_CONFIG_CONTENT,
	}

	configu.CustomiseConfigContent()

	if configu.FileContent != configs.APACHE_CONFIG_CONTENT_TEST {
		t.Errorf("Error on fileContent struct apache Attendu: %s, Obtenu: %s", configs.APACHE_CONFIG_CONTENT_TEST, configu.FileContent)
	}
}
