package fn

import (
	. "github.com/zytekaron/gotil/v2/fn"
	"os"
	"testing"
)

type Gob struct {
	A, B, C, D int
}

var (
	gob  = Gob{1, 2, 3, 4}
	path = "C:\\Users\\Zytekaron\\AppData\\Local\\Temp\\gob.dat"
)

func TestWrite(t *testing.T) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = WriteGob(file, &gob)
	if err != nil {
		t.Error(err)
	}
}

func TestRead(t *testing.T) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var res Gob
	err = ReadGob(file, &res)
	if err != nil {
		t.Error(err)
	}
	if res != gob {
		// unexported properties will be lost,
		// but none were used here
		t.Error("gobs are not the same")
	}
}

func TestWriteFile(t *testing.T) {
	err := WriteGobFile(path, &gob)
	if err != nil {
		t.Error(err)
	}
}

func TestReadFile(t *testing.T) {
	var res Gob
	err := ReadGobFile(path, &res)
	if err != nil {
		t.Error(err)
	}
	if res != gob {
		// unexported properties will be lost,
		// but none were used here
		t.Error("gobs are not the same")
	}
}
