package fn

import (
	"encoding/gob"
	"encoding/json"
	"io"
	"os"
)

// ReadJson reads a Go Object from io.Reader
func ReadJson(r io.Reader, i any) error {
	return json.NewDecoder(r).Decode(i)
}

// ReadJsonFile reads a Go Object from a file
func ReadJsonFile(path string, i any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return ReadJson(file, i)
}

// ReadGob reads a Go Object from io.Reader
func ReadGob(r io.Reader, i any) error {
	return gob.NewDecoder(r).Decode(i)
}

// ReadGobFile reads a Go Object from a file
func ReadGobFile(path string, i any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return ReadGob(file, i)
}

// WriteJson writes a Go Object to io.Writer
func WriteJson(w io.Writer, obj any) error {
	return json.NewEncoder(w).Encode(obj)
}

// WriteJsonFile writes a Go Object to a file
func WriteJsonFile(path string, obj any) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return WriteJson(file, obj)
}

// WriteGob writes a Go Object to io.Writer
func WriteGob(w io.Writer, obj any) error {
	return gob.NewEncoder(w).Encode(obj)
}

// WriteGobFile writes a Go Object to a file
func WriteGobFile(path string, obj any) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return WriteGob(file, obj)
}
