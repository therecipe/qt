package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {

	//TODO: check android / ios env

	if os.Getenv("GOPATH") == "" {
		fmt.Print("\nerror:\nGOPATH NOT SET\n\n")
		os.Exit(1)
	}

	if strings.Contains(os.Getenv("GOPATH"), runtime.GOROOT()) || strings.Contains(runtime.GOROOT(), os.Getenv("GOPATH")) {
		fmt.Print("\nerror:\nGOPATH IS GOROOT\n\n")
		os.Exit(1)
	}

	var qtPath = filepath.Join("/usr", "local", "Qt5.6.0")
	if runtime.GOOS == "windows" {
		qtPath = filepath.Join("C:\\", "Qt", "Qt5.6.0")
	}
	if _, err := ioutil.ReadDir(qtPath); err != nil {
		fmt.Printf("\nerror: Qt not found\nsolution: install Qt in: \"%v\"\n\n", qtPath)
		os.Exit(1)
	}

	switch runtime.GOOS {
	case "darwin":
		{
			if _, err := exec.LookPath("clang++"); err != nil {
				fmt.Printf("\nerror: clang++ not found\nsolution: install Xcode\n\n")
				os.Exit(1)
			}
		}

	case "linux":
		{
			if _, err := exec.LookPath("g++"); err != nil {
				fmt.Printf("\nerror: g++ not found\nsolution: sudo apt-get install g++\n\n")
				os.Exit(1)
			}
		}

	case "windows":
		{
			if _, err := exec.LookPath("g++.exe"); err != nil {
				fmt.Printf("\nerror: g++.exe not found\nsolution: add the directory that contains \"g++.exe\" to your PATH\n\n")
				os.Exit(1)
			}
		}
	}
}
