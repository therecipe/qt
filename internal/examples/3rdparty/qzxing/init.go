// +build ignore

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
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

	if _, err := os.Stat("qzxing"); err == nil {
		println("qzxing already cloned")
	} else {
		if out, err := exec.Command("git", "clone", "--depth=1", "https://github.com/ftylitak/qzxing.git").CombinedOutput(); err != nil {
			println("failed to clone qzxing", err.Error())
			println(string(out))
			os.Exit(1)
		}
		println("cloned qzxing")
	}

	//patch *.pro to build a static lib
	exec.Command("sed", "-i", "''", "-e", "s/# CONFIG/CONFIG/g", filepath.Join(pwd, "qzxing", "src", "QZXing.pro")).CombinedOutput()

	for _, target := range []string{runtime.GOOS, "android", "android_emulator", "ios"} {
		os.MkdirAll(filepath.Join(pwd, "qzxing", "src", target), 0755)

		if target == "android" {
			if _, ok := os.LookupEnv("ANDROID_NDK_DIR"); !ok {
				println("please export ANDROID_NDK_DIR")
				os.Exit(1)
			}
		}

		var qmake string
		switch target {
		case "darwin":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.13.0", "clang_64", "bin", "qmake")

		case "linux":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.13.0", "gcc_64", "bin", "qmake")

		case "android":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.13.0", "android_armv7", "bin", "qmake")

		case "android_emulator":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.13.0", "android_x86", "bin", "qmake")

		case "ios":
			if runtime.GOOS != "darwin" {
				return
			}
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.13.0", "ios", "bin", "qmake")
		}

		ndkPATH, ndkOK := os.LookupEnv("ANDROID_NDK_DIR")

		qCmd := exec.Command(qmake, "../QZXing.pro", "CONFIG+=qzxing_multimedia")
		qCmd.Dir = filepath.Join(pwd, "qzxing", "src", target)
		if ndkOK {
			qCmd.Env = append(qCmd.Env, "ANDROID_NDK_ROOT="+ndkPATH)
		}
		if out, err := qCmd.CombinedOutput(); err != nil {
			println("failed to generate makefile for", target, err.Error())
			println(string(out))
			os.Exit(1)
		}
		println("generated makefile for", target)

		iCmd := exec.Command("make", "-j", strconv.Itoa(runtime.NumCPU()))
		iCmd.Dir = filepath.Join(pwd, "qzxing", "src", target)
		if ndkOK {
			iCmd.Env = append(iCmd.Env, "ANDROID_NDK_ROOT="+ndkPATH)
		}
		if out, err := iCmd.CombinedOutput(); err != nil {
			println("failed to make qzxing for", target)
			println(string(out))
			os.Exit(1)
		}
		println("built qzxing for", target)
	}
}
