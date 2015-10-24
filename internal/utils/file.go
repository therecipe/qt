package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func MakeFolder(dir string) {
	var err = os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println("file.MakeFolder", err)
	}
}

func RemoveAll(name string) {
	var err = os.RemoveAll(name)
	if err != nil {
		fmt.Println("file.RemoveAll", err)
	}
}

func Save(name, data string) {
	var err = ioutil.WriteFile(name, []byte(data), 0777)
	if err != nil {
		fmt.Println("file.Save", err)
	}
}

func Load(name string) string {
	var b, err = ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("file.Load", err)
	}
	return string(b)
}
