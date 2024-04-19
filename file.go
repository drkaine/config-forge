package main

import "os"

func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func EditFile(apache Apache) error {
	file, err := CreateFile(apache.name)

	if err != nil {
		return err
	}

	_, err = file.WriteString(apache.fileContent)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
