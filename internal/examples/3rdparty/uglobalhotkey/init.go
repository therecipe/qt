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

	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		println("failed to get PWD", pwdErr.Error())
		os.Exit(1)
	}

	if _, err := os.Stat("UGlobalHotkey"); err == nil {
		println("UGlobalHotkey already cloned")
	} else {
		if out, err := exec.Command("git", "clone", "--depth=1", "https://github.com/falceeffect/UGlobalHotkey.git").CombinedOutput(); err != nil {
			println("failed to clone UGlobalHotkey", err.Error())
			println(string(out))
			os.Exit(1)
		}
		println("cloned UGlobalHotkey")
	}

	for _, target := range []string{runtime.GOOS} {

		var qmake string
		switch target {
		case "darwin":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.13.0", "clang_64", "bin", "qmake")

		case "linux":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.13.0", "gcc_64", "bin", "qmake")

		case "windows":
			if _, ok := os.LookupEnv("QT_MSVC"); ok {
				qmake = filepath.Join(os.Getenv("QT_DIR"), "5.13.0", "msvc2017_64", "bin", "qmake")
			} else {
				qmake = filepath.Join(os.Getenv("QT_DIR"), "5.13.0", "mingw73_64", "bin", "qmake")
			}
		}

		qCmd := exec.Command(qmake, filepath.Join(pwd, "UGlobalHotkey", "uglobalhotkey.pri"))
		qCmd.Dir = filepath.Join(pwd, "UGlobalHotkey")
		if out, err := qCmd.CombinedOutput(); err != nil {
			println("failed to generate makefile for", target, err.Error())
			println(string(out))
			os.Exit(1)
		}
		println("generated makefile for", target)

		iCmd := exec.Command("make", "mocables")
		iCmd.Dir = filepath.Join(pwd, "UGlobalHotkey")
		if out, err := iCmd.CombinedOutput(); err != nil {
			println("failed to run moc for", target)
			println(string(out))
			os.Exit(1)
		}
		println("ran moc for", target)

		files, err := ioutil.ReadDir(filepath.Join(pwd, "patch"))
		if err != nil {
			println("failed to read patch dir", err.Error())
			os.Exit(1)
		}
		for _, file := range files {
			exec.Command("cp", filepath.Join(pwd, "patch", file.Name()), filepath.Join(pwd, "UGlobalHotkey", file.Name())).Run()
		}
		println("copied patch files into UGlobalHotkey for", target)

		exec.Command("qtmoc", target, filepath.Join(pwd, "UGlobalHotkey")).Run()
		exec.Command("go", "install", filepath.Join(pwd, "UGlobalHotkey")).Run()
	}
}
