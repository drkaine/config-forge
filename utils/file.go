package utils

import (
	"config-forge/configurator"
	"os"
)

func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func EditFile(apache configurator.Apache) error {
	file, err := CreateFile(apache.Name)

	if err != nil {
		return err
	}

	_, err = file.WriteString(apache.FileContent)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
