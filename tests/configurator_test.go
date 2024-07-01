package tests

import (
	"config-forge/configs"
	"config-forge/configurator"
	"config-forge/templates"
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
		FileContent:  templates.APACHE_CONFIG,
	}

	content := apache.CustomiseConfigContent()

	if content != templates.APACHE_CONFIG_TEST {
		t.Errorf("Error on fileContent struct apache Attendu: %s, Obtenu: %s", templates.APACHE_CONFIG_TEST, apache.FileContent)
	}
}

func TestCustomiseNginxConfigContent(t *testing.T) {
	nginx := configurator.Nginx{
		Name:         configs.NAME_TEST,
		ServerName:   configs.SERVER_NAME_TEST,
		DocumentRoot: configs.DOCUMENT_ROOT_TEST,
		Path:         configs.NGINX_PATH_REPOSITORY,
		FileContent:  templates.NGINX_CONFIG,
	}

	content := nginx.CustomiseConfigContent()

	if content != templates.NGINX_CONFIG_TEST {
		t.Errorf("Error on fileContent struct NGINX Attendu: %s, Obtenu: %s", templates.NGINX_CONFIG_TEST, nginx.FileContent)
	}
}
