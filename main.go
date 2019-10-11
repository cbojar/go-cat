package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileNames := os.Args[1:]

	if len(fileNames) == 0 {
		fileNames = []string{"-"}
	}

	for _, fileName := range fileNames {
		var err error

		if fileName == "-" {
			err = catReader(os.Stdin, "<stdin>")
		} else {
			err = catFile(fileName)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "cat error: %v\n", err)
			os.Exit(1)
		}
	}
}

func catFile(fileName string) error {
	file, err := os.Open(fileName)

	if err != nil {
		return fmt.Errorf("could not open file \"%s\": %v", fileName, err)
	}

	defer file.Close()

	return catReader(file, fileName)
}

func catReader(reader io.Reader, name string) error {
	buffer := make([]byte, 4096)

	for {
		read, err := reader.Read(buffer)
		if read == 0 && err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			return fmt.Errorf("could not read file \"%s\": %v", name, err)
		}

		_, err = os.Stdout.Write(buffer[:read])

		if err != nil {
			return fmt.Errorf("could not write to stdout: %v", err)
		}
	}

	return nil
}
