package main

import (
	"strings"
)

type Config interface {
	ImplementConfigContent()
}

type Apache struct {
	name         string
	serverName   string
	documentRoot string
	path         string
	fileContent  string
}

func (apache *Apache) ImplementConfigContent() {
	apache.fileContent = strings.Replace(apache.fileContent, "ServerNameInput", apache.serverName, 1)
	apache.fileContent = strings.Replace(apache.fileContent, "ServerAliasInput", "www."+apache.serverName, 1)
	apache.fileContent = strings.Replace(apache.fileContent, "DocumentRootInput", apache.documentRoot, 2)
}

func InstanciationConfig(name string, serverName string, documentRoot string) Apache {
	config := Apache{
		name:         name + ".conf",
		serverName:   serverName,
		documentRoot: documentRoot,
		path:         ApachePath,
		fileContent:  ApacheConfig,
	}

	return config
}
