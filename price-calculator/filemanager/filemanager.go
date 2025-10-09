package filemanager

import (
	"bufio" //gives us utility functions for i/o data for files?
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("an error occured when opening the file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("an error occured whil scanning this file")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJson(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("something went wront when creating the file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("an error occured")
	}
	file.Close()
	return nil
}

func New(inputPath string, resultPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: resultPath,
	}
}
