package setup

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func Check(target string) {
	utils.Log.Infof("running: 'qtsetup check %v'", target)

	hash := "please install git"
	if _, err := exec.LookPath("git"); err == nil {
		hCmd := exec.Command("git", "rev-parse", "--verify", "HEAD")
		hCmd.Dir = utils.GoQtPkgPath()
		hash = strings.TrimSpace(utils.RunCmdOptional(hCmd, "get git hash"))
	}

	vars := [][]string{
		{"GOOS", runtime.GOOS},
		{"GOARCH", runtime.GOARCH},
		{"GOVERSION", runtime.Version()},
		{"GOROOT", runtime.GOROOT()},
		{"GOPATH", utils.MustGoPath()},
		{"GOBIN", utils.MustGoBin()},
		{"QT_HASH", hash},
		{"QT_VERSION", utils.QT_VERSION()},
		{"QT_VERSION_MAJOR", utils.QT_VERSION_MAJOR()},
		{"QT_DIR", utils.QT_DIR()},
		{"QT_STUB", fmt.Sprint(utils.QT_STUB())},
		{"QT_DEBUG", fmt.Sprint(utils.QT_DEBUG())},
	}

	if utils.CI() {
		vars = append(vars, [][]string{
			{"CI", fmt.Sprint(utils.CI())},
			{"QT_QMAKE_DIR", utils.QT_QMAKE_DIR()},
			{"QT_QMAKE_CGO", fmt.Sprint(utils.QT_QMAKE_CGO())},
		}...)
	}

	switch target {
	case "desktop", "ios", "ios-simulator":
		switch runtime.GOOS {
		case "darwin":
			vars = append(vars, [][]string{
				{"QT_HOMEBREW", fmt.Sprint(utils.QT_HOMEBREW())},
				{"XCODE_DIR", utils.XCODE_DIR()},
				//{"IPHONEOS_SDK_DIR", utils.IPHONEOS_SDK_DIR()},               //TODO: re-add after deploy is done; with absolute path
				//{"IPHONESIMULATOR_SDK_DIR", utils.IPHONESIMULATOR_SDK_DIR()}, //TODO: re-add after deploy is done; with absolute path
			}...)
		case "linux":
			vars = append(vars, [][]string{
				{"QT_DISTRO", utils.QT_DISTRO()},
				{"QT_PKG_CONFIG", fmt.Sprint(utils.QT_PKG_CONFIG())},
			}...)

			if utils.QT_PKG_CONFIG() {
				vars = append(vars, [][]string{
					{"QT_DOC_DIR", utils.QT_DOC_DIR()},
					{"QT_MISC_DIR", utils.QT_MISC_DIR()},
				}...)
			}
		case "windows":
			vars = append(vars, [][]string{
				{"QT_MSYS2", fmt.Sprint(utils.QT_MSYS2())},
				{"QT_MSYS2_DIR", utils.QT_MSYS2_DIR()},
				{"QT_MSYS2_ARCH", utils.QT_MSYS2_ARCH()},
			}...)
		}
	case "android":
		vars = append(vars, [][]string{
			{"JDK_DIR", utils.JDK_DIR()},
			{"ANDROID_SDK_DIR", utils.ANDROID_SDK_DIR()},
			{"ANDROID_NDK_DIR", utils.ANDROID_NDK_DIR()},
		}...)
	case "sailfish", "sailfish-emulator":
		vars = append(vars, [][]string{
			{"VIRTUALBOX_DIR", utils.VIRTUALBOX_DIR()},
			{"SAILFISH_DIR", utils.SAILFISH_DIR()},
		}...)
	case "rpi1", "rpi2", "rpi3":
		vars = append(vars, [][]string{
			{"RPI_TOOLS_DIR", utils.RPI_TOOLS_DIR()},
			{"RPI1_SYSROOT_DIR", utils.RPI1_SYSROOT_DIR()},
			{"RPI2_SYSROOT_DIR", utils.RPI2_SYSROOT_DIR()},
			{"RPI3_SYSROOT_DIR", utils.RPI3_SYSROOT_DIR()},
		}...)
	case "windows":
		vars = append(vars, [][]string{
			{"QT_MXE_DIR", utils.QT_MXE_DIR()},
			{"QT_MXE_ARCH", utils.QT_MXE_ARCH()},
			//{"QT_MXE_STATIC", fmt.Sprint(utils.QT_MXE_STATIC())},
		}...)
	}

	for _, v := range vars {
		var set string
		if _, ok := os.LookupEnv(v[0]); ok {
			set = "*"
		}
		utils.Log.Infof("%v:%v%v%v'%v'", v[0], strings.Repeat(" ", 25-len(v[0])), set, strings.Repeat(" ", 3-len(set)), v[1])
		if !strings.HasSuffix(v[0], "_DIR") {
			continue
		}
		if v[0] == "QT_DIR" && (utils.QT_HOMEBREW() || utils.QT_MSYS2() || utils.QT_PKG_CONFIG()) {
			continue
		}
		if _, err := ioutil.ReadDir(v[1]); err != nil && v[1] != "" {
			utils.Log.WithError(err).Panicf("failed to find %v (%v)", v[0], v[1])
		}
	}

	//TODO: combine -->
	if !strings.HasSuffix(target, "-docker") {
		switch runtime.GOOS {
		case "darwin":
			if _, err := exec.LookPath("clang++"); err != nil {
				utils.Log.WithError(err).Panic("failed to find clang++, did you install Xcode?")
			}
		case "linux":
			if _, err := exec.LookPath("g++"); err != nil {
				utils.Log.WithError(err).Panic("failed to find g++, did you install g++?")
			}
		case "windows":
			if _, err := exec.LookPath("g++"); err != nil && !utils.QT_MSYS2() {
				utils.Log.WithError(err).Panic("failed to find g++, did you add the directory that contains g++ to your PATH?")
			}
		}
	}
	switch target {
	case "windows-docker", "linux-docker", "android-docker":
		if _, err := exec.LookPath("docker"); err != nil {
			utils.Log.WithError(err).Panic("failed to find docker, did you install docker?")
		}
	}
	//<--
}
