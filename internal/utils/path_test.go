package utils

import (
	"os"
	"testing"
)

func TestGetQtPath(t *testing.T) {
	os.Setenv("GOPATH", "/Users/HOME/golang/")
	var check = "/Users/HOME/golang/src/github.com/therecipe/qt"

	if r := GetQtPkgPath(); r != check {
		t.Errorf("%v != %v", r, check)
	}
}
