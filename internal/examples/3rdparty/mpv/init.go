// +build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
)

func main() {
	if _, ok := os.LookupEnv("QT_PKG_CONFIG"); !ok {
		println("please export QT_PKG_CONFIG")
		os.Exit(1)
	}

	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		println("failed to get PWD", pwdErr.Error())
		os.Exit(1)
	}

	if _, err := os.Stat("mpv-examples"); err == nil {
		println("mpv-examples already cloned")
	} else {
		if out, err := exec.Command("git", "clone", "--depth=1", "https://github.com/mpv-player/mpv-examples.git").CombinedOutput(); err != nil {
			println("failed to clone mpv-examples", err.Error())
			println(string(out))
			os.Exit(1)
		}
		println("cloned mpv-examples")
	}

	//patch *.pro to build a static lib
	exec.Command("bash", "-c", fmt.Sprintf("echo \"TEMPLATE = lib\" >> %v", filepath.Join(pwd, "mpv-examples", "libmpv", "qml", "mpvtest.pro"))).CombinedOutput()
	exec.Command("bash", "-c", fmt.Sprintf("echo \"CONFIG += staticlib\" >> %v", filepath.Join(pwd, "mpv-examples", "libmpv", "qml", "mpvtest.pro"))).CombinedOutput()

	//patch main.cpp
	exec.Command("bash", "-c", fmt.Sprintf("head -n -17 %[1]v > %[2]v && mv %[2]v %[1]v", filepath.Join(pwd, "mpv-examples", "libmpv", "qml", "main.cpp"), filepath.Join(pwd, "mpv-examples", "libmpv", "qml", "copy_main.cpp"))).CombinedOutput()
	exec.Command("bash", "-c", fmt.Sprintf("echo 'extern \"C\" void initMpv() { std::setlocale(LC_NUMERIC, \"C\"); qmlRegisterType<MpvObject>(\"mpvtest\", 1, 0, \"MpvObject\");	}' >> %[1]v", filepath.Join(pwd, "mpv-examples", "libmpv", "qml", "main.cpp"))).CombinedOutput()

	//only tested on arch, since the libmpv version is quite old on ubuntu 16.04 and 18.04
	for _, target := range []string{runtime.GOOS} {

		qCmd := exec.Command("qmake", "mpvtest.pro")
		qCmd.Dir = filepath.Join(pwd, "mpv-examples", "libmpv", "qml")
		if out, err := qCmd.CombinedOutput(); err != nil {
			println("failed to generate makefile for", target, err.Error())
			println(string(out))
			os.Exit(1)
		}
		println("generated makefile for", target)

		iCmd := exec.Command("make", "-j", strconv.Itoa(runtime.NumCPU()))
		iCmd.Dir = filepath.Join(pwd, "mpv-examples", "libmpv", "qml")
		if out, err := iCmd.CombinedOutput(); err != nil {
			println("failed to make mpv-examples for", target)
			println(string(out))
			os.Exit(1)
		}
		println("built mpv-examples for", target)
	}
}
