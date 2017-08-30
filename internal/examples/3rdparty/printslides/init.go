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
	if _, err := ioutil.ReadDir("qml-presentation-system"); err == nil {
		println("qml-presentation-system already cloned")
		os.Exit(1)
	}

	if out, err := exec.Command("git", "clone", "https://github.com/qt-labs/qml-presentation-system.git").CombinedOutput(); err != nil {
		println("failed to clone qml-presentation-system", err.Error())
		println(string(out))
		os.Exit(1)
	}
	println("cloned qml-presentation-system")

	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		println("failed to get PWD", pwdErr.Error())
		os.Exit(1)
	}

	for _, target := range []string{runtime.GOOS, "android"} {
		os.MkdirAll(filepath.Join(pwd, "qml-presentation-system", target), 0755)

		var qmake string
		switch target {
		case "darwin":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.9.1", "clang_64", "bin", "qmake")

		case "linux":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.9.1", "gcc_64", "bin", "qmake")

		case "android":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.9.1", "android_armv7", "bin", "qmake")
		}

		ndkPATH, ndkOK := os.LookupEnv("ANDROID_NDK_DIR")

		qCmd := exec.Command(qmake, "../presentation.pro")
		qCmd.Dir = filepath.Join(pwd, "qml-presentation-system", target)
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
		iCmd.Dir = filepath.Join(pwd, "qml-presentation-system", target)
		if ndkOK {
			iCmd.Env = append(iCmd.Env, "ANDROID_NDK_ROOT="+ndkPATH)
		}
		if out, err := iCmd.CombinedOutput(); err != nil {
			println("failed to install qml-presentation-system for", target)
			println(string(out))
			os.Exit(1)
		}
		println("installed qml-presentation-system for", target)
	}
}
