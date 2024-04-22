package utils

import (
	"config-forge/configurator"
	"os"
)

func createFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func EditFile(config configurator.Config) error {
	file, err := createFile(config.GetName())

	if err != nil {
		return err
	}

	_, err = file.WriteString(config.GetFileContent())
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
