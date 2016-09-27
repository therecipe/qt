package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func main() {

	//TODO: check android / ios / sailfish / raspberry pi env

	if os.Getenv("GOPATH") == "" {
		fmt.Print("\nerror:\nGOPATH not set\n\n")
		os.Exit(1)
	}

	if strings.Contains(os.Getenv("GOPATH"), runtime.GOROOT()) || strings.Contains(runtime.GOROOT(), os.Getenv("GOPATH")) {
		fmt.Print("\nerror:\nGOPATH is GOROOT\n\n")
		os.Exit(1)
	}

	if _, err := ioutil.ReadDir(utils.QtInstallDir()); err != nil {
		fmt.Printf("\nerror: Qt not found\nsolution: install Qt in: \"%v\"\n\n", utils.QtInstallDir())
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
				fmt.Printf("\nerror: g++ not found\nsolution: sudo apt-get -y install build-essential\n\n")
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
