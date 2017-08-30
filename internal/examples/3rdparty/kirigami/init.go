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
	if _, err := ioutil.ReadDir("kirigami"); err == nil {
		println("kirigami already cloned")
		os.Exit(1)
	}

	if out, err := exec.Command("git", "clone", "https://github.com/KDE/kirigami.git").CombinedOutput(); err != nil {
		println("failed to clone kirigami", err.Error())
		println(string(out))
		os.Exit(1)
	}
	println("cloned kirigami")

	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		println("failed to get PWD", pwdErr.Error())
		os.Exit(1)
	}

	for _, target := range []string{runtime.GOOS, "android"} {
		os.MkdirAll(filepath.Join(pwd, "kirigami", target), 0755)

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

		qCmd := exec.Command(qmake, "../kirigami.pro", "CONFIG-=android_install")
		qCmd.Dir = filepath.Join(pwd, "kirigami", target)
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
		iCmd.Dir = filepath.Join(pwd, "kirigami", target)
		if ndkOK {
			iCmd.Env = append(iCmd.Env, "ANDROID_NDK_ROOT="+ndkPATH)
		}
		if out, err := iCmd.CombinedOutput(); err != nil {
			println("failed to install kirigami for", target)
			println(string(out))
			os.Exit(1)
		}
		println("installed kirigami for", target)
	}
}
