package main

import (
	"bytes"
	"os"
	"testing"
)

const (
	WrongChoiceError    = "Wrong choice"
	GoodChoiceError     = "Good choice"
	ApacheInput         = "apache"
	BadInput            = "bad"
	Name                = "default-test"
	ServerName          = "test.net"
	DocumentRoot        = "/var/www/monsite.com/public_html"
	ApacheConfigContent = `
	<VirtualHost *:80>
	    ServerName test.net
	    ServerAlias www.test.net
	    DocumentRoot /var/www/monsite.com/public_html

	    <Directory /var/www/monsite.com/public_html>
	        Options Indexes FollowSymLinks
	        AllowOverride All
	        Require all granted
	    </Directory>

	    ErrorLog ${APACHE_LOG_DIR}/monsite_error.log
	    CustomLog ${APACHE_LOG_DIR}/monsite_access.log combined
	</VirtualHost>`
)

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
	if !InArray(ApacheInput, configAccepted) {
		t.Errorf("Not in array")
	}
}

func TestNotInArray(t *testing.T) {
	if InArray(BadInput, configAccepted) {
		t.Errorf("In array")
	}
}

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
