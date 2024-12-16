package Utils

import "io/ioutil"

// ReadFileAsString reads the entire content of a file and returns it as a string
func ReadFileAsString(filename string) (string, error) {
	// Read the file content
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string and return
	return string(content), nil
}
