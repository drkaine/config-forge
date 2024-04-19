package main

import (
	"testing"
)

func TestInstanciationConfig(t *testing.T) {
	config := InstanciationConfig(Name, ServerName, DocumentRoot)

	if config.name != Name+".conf" {
		t.Errorf("The name of the config is %q and need to be %q", config.name, Name)
	}

	if config.serverName != ServerName {
		t.Errorf("Error on serverName struct apache")
	}

	if config.documentRoot != DocumentRoot {
		t.Errorf("Error on documentRoot struct apache")
	}

	if config.path != ApachePath {
		t.Errorf("Error on documentRoot struct apache")
	}

	if config.fileContent != ApacheConfig {
		t.Errorf("Error on documentRoot struct apache")
	}
}

func TestImplementConfigContent(t *testing.T) {
	config := Apache{
		name:         Name,
		serverName:   ServerName,
		documentRoot: DocumentRoot,
		path:         ApachePath,
		fileContent:  ApacheConfig,
	}

	config.ImplementConfigContent()

	if config.fileContent != ApacheConfigContent {
		t.Errorf("Error on fileContent struct apache Attendu: %s, Obtenu: %s", ApacheConfigContent, config.fileContent)
	}
}
