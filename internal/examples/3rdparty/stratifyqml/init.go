// +build ignore

package main

import (
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

	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		println("failed to get PWD", pwdErr.Error())
		os.Exit(1)
	}

	if _, err := os.Stat("StratifyQML"); err == nil {
		println("StratifyQML already cloned")
	} else {
		if out, err := exec.Command("git", "clone", "--depth=1", "https://github.com/StratifyLabs/StratifyQML.git").CombinedOutput(); err != nil {
			println("failed to clone StratifyQML", err.Error())
			println(string(out))
			os.Exit(1)
		}
		println("cloned StratifyQML")
	}

	for _, target := range []string{runtime.GOOS, "android"} {
		os.MkdirAll(filepath.Join(pwd, "StratifyQML", target), 0755)

		if target == "android" {
			if _, ok := os.LookupEnv("ANDROID_NDK_DIR"); !ok {
				println("please export ANDROID_NDK_DIR")
				os.Exit(1)
			}
		}

		var qmake string
		switch target {
		case "darwin":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.9.5", "clang_64", "bin", "qmake")

		case "linux":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.9.5", "gcc_64", "bin", "qmake")

		case "android":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.9.5", "android_armv7", "bin", "qmake")
		}

		ndkPATH, ndkOK := os.LookupEnv("ANDROID_NDK_DIR")

		qCmd := exec.Command(qmake, "../StratifyLabs/StratifyLabsUI/StratifyLabsUI.pro", "CONFIG-=android_install")
		qCmd.Dir = filepath.Join(pwd, "StratifyQML", target)
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
		iCmd.Dir = filepath.Join(pwd, "StratifyQML", target)
		if ndkOK {
			iCmd.Env = append(iCmd.Env, "ANDROID_NDK_ROOT="+ndkPATH)
		}
		if out, err := iCmd.CombinedOutput(); err != nil {
			println("failed to install StratifyQML for", target)
			println(string(out))
			os.Exit(1)
		}
		println("installed StratifyQML for", target)
	}
}
