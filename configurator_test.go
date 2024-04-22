package main

import (
	"config-forge/configs"
	"config-forge/configurator"
	"testing"
)

func TestInstanciationApache(t *testing.T) {
	informations := map[string]string{
		"name":         configs.NAME_TEST,
		"serverName":   configs.SERVER_NAME_TEST,
		"documentRoot": configs.DOCUMENT_ROOT_TEST,
	}

	configu := configurator.InstanciationConfig(configs.APACHE_INPUT, informations)

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

func TestCustomiseApacheConfigContent(t *testing.T) {
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

func TestCustomiseNgnixConfigContent(t *testing.T) {
	configu := configurator.Ngnix{
		Name:         configs.NAME_TEST,
		ServerName:   configs.SERVER_NAME_TEST,
		DocumentRoot: configs.DOCUMENT_ROOT_TEST,
		Path:         configs.NGNIX_PATH_REPOSITORY,
		FileContent:  configs.NGNIX_CONFIG_CONTENT,
	}

	configu.CustomiseConfigContent()

	if configu.FileContent != configs.NGNIX_CONFIG_CONTENT_TEST {
		t.Errorf("Error on fileContent struct NGNIX Attendu: %s, Obtenu: %s", configs.NGNIX_CONFIG_CONTENT_TEST, configu.FileContent)
	}
}
