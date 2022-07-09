package fn

import (
	"os"
	"path"
	"testing"
)

type Gob struct {
	A, B, C, D int
}

var (
	obj      = Gob{1, 2, 3, 4}
	filePath = path.Join(os.TempDir(), "test_gob.dat")
)

func TestWrite(t *testing.T) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = WriteGob(file, &obj)
	if err != nil {
		t.Error(err)
	}
}

func TestRead(t *testing.T) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var res Gob
	err = ReadGob(file, &res)
	if err != nil {
		t.Error(err)
	}
	if res != obj {
		t.Error("gobs are not the same")
	}
}

func TestWriteFile(t *testing.T) {
	err := WriteGobFile(filePath, &obj)
	if err != nil {
		t.Error(err)
	}
}

func TestReadFile(t *testing.T) {
	var res Gob
	err := ReadGobFile(filePath, &res)
	if err != nil {
		t.Error(err)
	}
	if res != obj {
		t.Error("gobs are not the same")
	}
}
