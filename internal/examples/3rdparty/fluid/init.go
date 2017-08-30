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
	if _, err := ioutil.ReadDir("fluid"); err == nil {
		println("fluid already cloned")
		os.Exit(1)
	}

	if out, err := exec.Command("git", "clone", "https://github.com/lirios/fluid.git").CombinedOutput(); err != nil {
		println("failed to clone fluid", err.Error())
		println(string(out))
		os.Exit(1)
	}
	println("cloned fluid")

	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		println("failed to get PWD", pwdErr.Error())
		os.Exit(1)
	}

	sCmd := exec.Command("git", "submodule", "update", "--init")
	sCmd.Dir = filepath.Join(pwd, "fluid")
	if out, err := sCmd.CombinedOutput(); err != nil {
		println("failed to init fluid submodules", err.Error())
		println(string(out))
		os.Exit(1)
	}
	println("inited fluid submodules")

	for _, target := range []string{runtime.GOOS, "android"} {
		os.MkdirAll(filepath.Join(pwd, "fluid", target), 0755)

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

		qCmd := exec.Command(qmake, "../fluid.pro", "CONFIG+=install_under_qt")
		qCmd.Dir = filepath.Join(pwd, "fluid", target)
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
		iCmd.Dir = filepath.Join(pwd, "fluid", target)
		if ndkOK {
			iCmd.Env = append(iCmd.Env, "ANDROID_NDK_ROOT="+ndkPATH)
		}
		if out, err := iCmd.CombinedOutput(); err != nil {
			println("failed to install fluid for", target)
			println(string(out))
			os.Exit(1)
		}
		println("installed fluid for", target)
	}
}
