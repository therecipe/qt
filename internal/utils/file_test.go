package utils

import (
	"io/ioutil"
	"testing"
)

func TestMakeFolder(t *testing.T) {
	var dir = "TestFolder/"
	defer RemoveAll(dir)

	MakeFolder(dir)

	if _, err := ioutil.ReadDir(dir); err != nil {
		t.Error(err)
	}
}

func TestSaveLoad(t *testing.T) {
	var (
		check = "TEST STRING FOR TEST FILE"
		file  = "TestFile.txt"
	)
	defer RemoveAll(file)

	Save(file, check)

	if output := Load(file); output != check {
		t.Errorf("%v != %v", output, check)
	}
}
