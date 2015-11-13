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

	//TODO: check android env

	if os.Getenv("GOPATH") == "" {
		fmt.Print("\nerror:\nGOPATH NOT SET\n\n")
		os.Exit(1)
	}

	if strings.Contains(os.Getenv("GOPATH"), runtime.GOROOT()) || strings.Contains(runtime.GOROOT(), os.Getenv("GOPATH")) {
		fmt.Print("\nerror:\nGOPATH IS GOROOT\n\n")
		os.Exit(1)
	}

	switch runtime.GOOS {
	case "darwin", "linux":
		{
			var (
				qtPath = filepath.Join("/usr", "local", "Qt5.5.1")
				_, err = ioutil.ReadDir(qtPath)
			)
			if err != nil {
				fmt.Printf("\nerror: Qt not found\nsolution: install Qt in: \"%v\"\n\n", qtPath)
				os.Exit(1)
			}
		}
	case "windows":
		{
			var (
				qtPath = filepath.Join("C:", "Qt", "Qt5.5.1")
				_, err = ioutil.ReadDir(qtPath)
			)
			if err != nil {
				fmt.Printf("\nerror: Qt not found\nsolution: install Qt in: \"%v\"\n\n", qtPath)
				os.Exit(1)
			}
		}
	}

	switch runtime.GOOS {
	case "darwin":
		{
			var _, err = exec.LookPath("clang++")
			if err != nil {
				fmt.Printf("\nerror: clang++ not found\nsolution: install Xcode\n\n")
				os.Exit(1)
			}
		}
	case "linux":
		{
			var _, err = exec.LookPath("g++")
			if err != nil {
				fmt.Printf("\nerror: g++ not found\nsolution: sudo apt-get install g++\n\n")
				os.Exit(1)
			}
		}
	case "windows":
		{
			var _, err = exec.LookPath("gcc.exe")
			if err != nil {
				fmt.Printf("\nerror: gcc.exe not found\nsolution: add the directory that contains \"gcc\" to your PATH\n\n")
				os.Exit(1)
			}
		}
	}
}
