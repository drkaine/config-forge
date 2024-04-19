package main

import (
	"testing"
)

func TestInstanciationConfig(t *testing.T) {
	config := InstanciationConfig(NAME_TEST, SERVER_NAME_TEST, DOCUMENT_ROOT_TEST)

	if config.name != NAME_FILE_TEST {
		t.Errorf("The name of the config is %q and need to be %q", config.name, NAME_TEST)
	}

	if config.serverName != SERVER_NAME_TEST {
		t.Errorf("Error on serverName struct apache")
	}

	if config.documentRoot != DOCUMENT_ROOT_TEST {
		t.Errorf("Error on documentRoot struct apache")
	}

	if config.path != APACHE_PATH_REPOSITORY {
		t.Errorf("Error on documentRoot struct apache")
	}

	if config.fileContent != APACHE_CONFIG_CONTENT {
		t.Errorf("Error on documentRoot struct apache")
	}
}

func TestImplementConfigContent(t *testing.T) {
	config := Apache{
		name:         NAME_TEST,
		serverName:   SERVER_NAME_TEST,
		documentRoot: DOCUMENT_ROOT_TEST,
		path:         APACHE_PATH_REPOSITORY,
		fileContent:  APACHE_CONFIG_CONTENT,
	}

	config.ImplementConfigContent()

	if config.fileContent != APACHE_CONFIG_CONTENT_TEST {
		t.Errorf("Error on fileContent struct apache Attendu: %s, Obtenu: %s", APACHE_CONFIG_CONTENT_TEST, config.fileContent)
	}
}
