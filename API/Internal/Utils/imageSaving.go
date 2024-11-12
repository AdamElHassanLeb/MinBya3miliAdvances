package Utils

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// SaveFile saves the uploaded file to the given directory with a new name and returns the full path or an error.
func SaveFile(file io.Reader, dir, newFileName string) (string, error) {
	// Ensure the directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}
	}

	// Create the full path for the new file
	filePath := filepath.Join(dir, newFileName)

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Copy the content from the uploaded file to the new file
	if _, err := io.Copy(out, file); err != nil {
		return "", err
	}

	return filePath, nil
}

// RespondJSON sends a JSON response with the specified status code and data.
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Set the status code
	w.WriteHeader(statusCode)
	// Encode and send the data as JSON
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// If there is an error encoding the data, send an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
