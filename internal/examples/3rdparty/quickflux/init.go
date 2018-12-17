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

	if _, err := ioutil.ReadDir("quickflux"); err == nil {
		println("quickflux already cloned")
		os.Exit(1)
	}

	if out, err := exec.Command("git", "clone", "https://github.com/benlau/quickflux.git").CombinedOutput(); err != nil {
		println("failed to clone quickflux", err.Error())
		println(string(out))
		os.Exit(1)
	}
	println("cloned quickflux")

	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		println("failed to get PWD", pwdErr.Error())
		os.Exit(1)
	}

	pri, priErr := ioutil.ReadFile(filepath.Join(pwd, "quickflux", "quickflux.pri"))
	if priErr != nil {
		println("failed to patch quickflux.pri", priErr.Error())
		os.Exit(1)
	}
	ioutil.WriteFile(filepath.Join(pwd, "quickflux", "quickflux.pri"), []byte("QT += qml quick\n"+string(pri)), 0644)
	println("patched quickflux.pri")

	for _, target := range []string{runtime.GOOS} {

		var qmake string
		switch target {
		case "darwin":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.12.0", "clang_64", "bin", "qmake")

		case "linux":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.12.0", "gcc_64", "bin", "qmake")

		case "android":
			qmake = filepath.Join(os.Getenv("QT_DIR"), "5.12.0", "android_armv7", "bin", "qmake")
		}

		qCmd := exec.Command(qmake, filepath.Join(pwd, "quickflux", "quickflux.pri"))
		qCmd.Dir = filepath.Join(pwd, "quickflux")
		if out, err := qCmd.CombinedOutput(); err != nil {
			println("failed to generate makefile for", target, err.Error())
			println(string(out))
			os.Exit(1)
		}
		println("generated makefile for", target)

		iCmd := exec.Command("make", "mocables")
		iCmd.Dir = filepath.Join(pwd, "quickflux")
		if out, err := iCmd.CombinedOutput(); err != nil {
			println("failed to run moc for", target)
			println(string(out))
			os.Exit(1)
		}
		println("ran moc for", target)

		ioutil.WriteFile(filepath.Join(pwd, "quickflux", "cgo_stub.go"), []byte(quickflux_stub), 0644)
		println("saved cgo_stub in quickflux for", target)

		files, err := ioutil.ReadDir(filepath.Join(pwd, "quickflux", "priv"))
		if err != nil {
			println("failed to read priv dir", err.Error())
			os.Exit(1)
		}
		for _, file := range files {
			exec.Command("cp", filepath.Join(pwd, "quickflux", "priv", file.Name()), filepath.Join(pwd, "quickflux", file.Name())).Run()
		}
		println("copied priv *.h and *.cpp into quickflux for", target)

		files, err = ioutil.ReadDir(filepath.Join(pwd, "quickflux"))
		if err != nil {
			println("failed to read quickflux dir", err.Error())
			os.Exit(1)
		}
		for _, file := range files {
			if !file.IsDir() {
				exec.Command("sed", "-i", "''", "-e", "'s!./priv/!!g'", filepath.Join(pwd, "quickflux", file.Name())).Run()
				exec.Command("sed", "-i", "''", "-e", "'s!priv/!!g'", filepath.Join(pwd, "quickflux", file.Name())).Run()
			}
		}
		println("patched *.h and *.cpp in quickflux for", target)

		exec.Command("qtmoc", target, filepath.Join(pwd, "quickflux")).Run()
		exec.Command("go", "install", filepath.Join(pwd, "quickflux")).Run()
	}
}

const quickflux_stub = `package quickflux

import (
	"github.com/therecipe/qt/core"
	_ "github.com/therecipe/qt/quick"
)

type quickflux_stub struct {
	core.QObject
}`
