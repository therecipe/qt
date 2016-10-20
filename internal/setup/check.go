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
	var buildTarget = "desktop"
	if len(os.Args) > 1 {
		buildTarget = os.Args[1]
	}

	gopaths := os.Getenv("GOPATH")
	if len(gopaths) == 0 {
		fmt.Print("\nerror:\nGOPATH not set\nsolution: define GOPATH\n\n")
		os.Exit(1)
	}

	for _, path :=  range strings.Split(gopaths, ":") {
		if _, err := ioutil.ReadDir(path); err != nil {
			fmt.Printf("\nerror:\nInvalid GOPATH %q\nsolution: specify valid GOPATH\n\n", path)
			os.Exit(1)
		}

		if strings.Contains(path, runtime.GOROOT()) {
			fmt.Printf("\nerror:\nGOPATH %q is or contains GOROOT\n\n", path)
			os.Exit(1)
		}

		if strings.Contains(runtime.GOROOT(), path) {
			fmt.Printf("\nerror:\nGOROOT is or contains GOPATH %q\n\n", path)
			os.Exit(1)
		}
	}


	if _, err := ioutil.ReadDir(utils.QT_DIR()); err != nil && !utils.UsePkgConfig() {
		fmt.Printf("\nerror: Qt not found\nsolution: install Qt in \"%v\" or define QT_DIR\n\n", utils.QT_DIR())
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
			if _, err := exec.LookPath("g++"); err != nil {
				fmt.Printf("\nerror: g++ not found\nsolution: add the directory that contains \"g++\" to your PATH\n\n")
				os.Exit(1)
			}
		}
	}

	switch buildTarget {
	case "android":
		{
			if _, err := ioutil.ReadDir(utils.ANDROID_SDK_DIR()); err != nil {
				fmt.Printf("\nerror: android-sdk not found\nsolution: install the android-sdk in \"%v\" or define ANDROID_SDK_DIR\n\n", utils.ANDROID_SDK_DIR())
				os.Exit(1)
			}

			if _, err := ioutil.ReadDir(utils.ANDROID_NDK_DIR()); err != nil {
				fmt.Printf("\nerror: android-ndk not found\nsolution: install the android-ndk in \"%v\" or define ANDROID_NDK_DIR\n\n", utils.ANDROID_NDK_DIR())
				os.Exit(1)
			}

			if _, err := ioutil.ReadDir(utils.JDK_DIR()); err != nil {
				fmt.Printf("\nerror: jdk not found\nsolution: install jdk in \"%v\" or define JDK_DIR\n\n", utils.JDK_DIR())
				os.Exit(1)
			}
		}

	case "sailfish", "sailfish-emulator":
		{
			if _, err := ioutil.ReadDir(utils.VIRTUALBOX_DIR()); err != nil {
				fmt.Printf("\nerror: virtualbox not found\nsolution: install virtualbox in \"%v\" or define VIRTUALBOX_DIR\n\n", utils.VIRTUALBOX_DIR())
				os.Exit(1)
			}

			if _, err := ioutil.ReadDir(utils.SAILFISH_DIR()); err != nil {
				fmt.Printf("\nerror: sailfish-sdk not found\nsolution: install the sailfish-sdk in \"%v\" or define SAILFISH_DIR\n\n", utils.SAILFISH_DIR())
				os.Exit(1)
			}
		}

	case "rpi1", "rpi2", "rpi3":
		{
			//TODO:
		}
	}
}
