// +build ignore

package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	if _, ok := os.LookupEnv("QT_DIR"); !ok {
		println("please export QT_DIR")
		os.Exit(1)
	}

	if _, err := ioutil.ReadDir("qml-material"); err == nil {
		println("qml-material already cloned")
		os.Exit(1)
	}

	if out, err := exec.Command("git", "clone", "https://github.com/papyros/qml-material.git").CombinedOutput(); err != nil {
		println("failed to clone qml-material", err.Error())
		println(string(out))
		os.Exit(1)
	}
	println("cloned qml-material")

	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		println("failed to get PWD", pwdErr.Error())
		os.Exit(1)
	}

	for _, target := range []string{runtime.GOOS, "android"} {
		os.MkdirAll(filepath.Join(pwd, "qml-material", target), 0755)

		var qmake string
		switch target {
		case "darwin":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.12.0", "clang_64", "bin", "qmake")

		case "linux":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.12.0", "gcc_64", "bin", "qmake")

		case "android":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.12.0", "android_armv7", "bin", "qmake")
		}

		ndkPATH, ndkOK := os.LookupEnv("ANDROID_NDK_DIR")

		qCmd := exec.Command(qmake, "../qml-material.pro", "CONFIG-=android_install")
		qCmd.Dir = filepath.Join(pwd, "qml-material", target)
		if ndkOK {
			qCmd.Env = append(qCmd.Env, "ANDROID_NDK_ROOT="+ndkPATH)
		}
		if out, err := qCmd.CombinedOutput(); err != nil {
			println("failed to generate makefile for", target, err.Error())
			println(string(out))
			os.Exit(1)
		}
		println("generated makefile for", target)

		iCmd := exec.Command("make", "install")
		iCmd.Dir = filepath.Join(pwd, "qml-material", target)
		if ndkOK {
			iCmd.Env = append(iCmd.Env, "ANDROID_NDK_ROOT="+ndkPATH)
		}
		if out, err := iCmd.CombinedOutput(); err != nil {
			println("failed to install qml-material for", target)
			println(string(out))
			os.Exit(1)
		}
		println("installed qml-material for", target)
	}
}
