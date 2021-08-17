package fn

import (
	"encoding/gob"
	"io"
	"os"
)

// ReadGob reads a Go Object from io.Reader
// Unexported properties on structs will be lost
func ReadGob(r io.Reader, i interface{}) error {
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(i)
	return err
}

// ReadGobFile reads a Go Object from a file
// Unexported properties on structs will be lost
func ReadGobFile(path string, i interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return ReadGob(file, i)
}

// WriteGob writes a Go Object to io.Writer
// Unexported properties on structs will be lost
func WriteGob(w io.Writer, obj interface{}) error {
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(obj)
	return err
}

// WriteGobFile writes a Go Object to a file
// Unexported properties on structs will be lost
func WriteGobFile(path string, obj interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return WriteGob(file, obj)
}
