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

	configu := configurator.ConfigBuilder(configs.APACHE_INPUT, informations)

	if configu.GetName() != configs.NAME_FILE_TEST {
		t.Errorf("The name of the config is %q and need to be %q", configu.GetName(), configs.NAME_TEST)
	}
}

func TestInstanciationNginx(t *testing.T) {
	informations := map[string]string{
		"name":         configs.NAME_TEST,
		"serverName":   configs.SERVER_NAME_TEST,
		"documentRoot": configs.DOCUMENT_ROOT_TEST,
	}

	configu := configurator.ConfigBuilder("nginx", informations)

	if configu.GetName() != configs.NAME_TEST {
		t.Errorf("The name of the config is %q and need to be %q", configu.GetName(), configs.NAME_TEST)
	}
}

func TestCustomiseApacheConfigContent(t *testing.T) {
	apache := configurator.Apache{
		Name:         configs.NAME_TEST,
		ServerName:   configs.SERVER_NAME_TEST,
		DocumentRoot: configs.DOCUMENT_ROOT_TEST,
		Path:         configs.APACHE_PATH_REPOSITORY,
		FileContent:  configs.APACHE_CONFIG_CONTENT,
	}

	content := apache.CustomiseConfigContent()

	if content != configs.APACHE_CONFIG_CONTENT_TEST {
		t.Errorf("Error on fileContent struct apache Attendu: %s, Obtenu: %s", configs.APACHE_CONFIG_CONTENT_TEST, apache.FileContent)
	}
}

func TestCustomiseNginxConfigContent(t *testing.T) {
	nginx := configurator.Nginx{
		Name:         configs.NAME_TEST,
		ServerName:   configs.SERVER_NAME_TEST,
		DocumentRoot: configs.DOCUMENT_ROOT_TEST,
		Path:         configs.NGINX_PATH_REPOSITORY,
		FileContent:  configs.NGINX_CONFIG_CONTENT,
	}

	content := nginx.CustomiseConfigContent()

	if content != configs.NGINX_CONFIG_CONTENT_TEST {
		t.Errorf("Error on fileContent struct NGINX Attendu: %s, Obtenu: %s", configs.NGINX_CONFIG_CONTENT_TEST, nginx.FileContent)
	}
}
