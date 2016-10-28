package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var (
		buildTarget = func() string {
			if len(os.Args) > 1 {
				return os.Args[1]
			}
			return "desktop"
		}()
	)
	utils.Log.Infof("running setup/check.go %v", buildTarget)

	switch buildTarget {
	case "all":
		{
			for _, target := range []string{"desktop", "android", "ios", "ios-simulator",
				"sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "windows",
				"linux-docker", "windows-docker", "android-docker"} {
				utils.RunCmd(exec.Command(utils.GoQtPkgPath(func() string {
					if runtime.GOOS == "windows" {
						return "setup.bat"
					}
					return "setup.sh"
				}()), target), fmt.Sprintf("run setup %v", target))
			}
			os.Exit(0)
		}

	case "desktop", "android", "ios", "ios-simulator",
		"sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "windows",
		"linux-docker", "windows-docker", "android-docker":
		{

		}

	default:
		{
			utils.Log.Panicf("failed to recognize build target %v", buildTarget)
		}
	}

	utils.Log.Infoln("GOROOT:", runtime.GOROOT())
	utils.Log.Infoln("GOPATH:", utils.MustGoPath())
	utils.Log.Infoln("QT_DIR:", utils.QT_DIR())
	utils.Log.Infoln("QT_STUB:", utils.QT_STUB())

	switch buildTarget {
	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					utils.Log.Infoln("QT_HOMEBREW:", os.Getenv("QT_HOMEBREW"))
					utils.Log.Infoln("IsHomebrew:", utils.IsHomeBrewQtDir())
					utils.Log.Infoln("UseHomebrew:", utils.UseHomeBrew())
					utils.Log.Infoln("QT_DARWIN_DIR:", utils.QT_DARWIN_DIR())
					utils.Log.Infoln("XCODE_DIR:", utils.XCODE_DIR())
					utils.Log.Infoln("MACOS_SDK_DIR:", utils.MACOS_SDK_DIR())
					utils.Log.Infoln("IPHONEOS_SDK_DIR:", utils.IPHONEOS_SDK_DIR())
					utils.Log.Infoln("IPHONESIMULATOR_SDK_DIR:", utils.IPHONESIMULATOR_SDK_DIR())
				}

			case "linux":
				{
					utils.Log.Infoln("Distro:", utils.LinuxDistro())
					utils.Log.Infoln("QT_PKG_CONFIG:", os.Getenv("QT_PKG_CONFIG"))
					utils.Log.Infoln("UsePkgConfig:", utils.UsePkgConfig())
					if utils.UsePkgConfig() {
						utils.Log.Infoln("QT_DOC_DIR:", utils.QT_DOC_DIR())
						utils.Log.Infoln("QT_MISC_DIR:", utils.QT_MISC_DIR())
					}
				}
			}
		}

	case "android":
		{
			utils.Log.Infoln("JDK_DIR:", utils.JDK_DIR())
			utils.Log.Infoln("JAVA_HOME:", os.Getenv("JAVA_HOME"))
			utils.Log.Infoln("ANDROID_SDK_DIR:", utils.ANDROID_SDK_DIR())
			utils.Log.Infoln("ANDROID_NDK_DIR:", utils.ANDROID_NDK_DIR())
		}

	case "sailfish", "sailfish-emulator":
		{
			utils.Log.Infoln("VIRTUALBOX_DIR:", utils.VIRTUALBOX_DIR())
			utils.Log.Infoln("SAILFISH_DIR:", utils.SAILFISH_DIR())
		}

	case "rpi1", "rpi2", "rpi3":
		{
			utils.Log.Infoln("RPI_TOOLS_DIR:", utils.RPI_TOOLS_DIR())
			utils.Log.Infoln("RPI1_SYSROOT_DIR:", utils.RPI1_SYSROOT_DIR())
			utils.Log.Infoln("RPI2_SYSROOT_DIR:", utils.RPI2_SYSROOT_DIR())
			utils.Log.Infoln("RPI3_SYSROOT_DIR:", utils.RPI3_SYSROOT_DIR())
		}

	case "windows-docker", "linux-docker", "android-docker":
		{
			//TODO:
		}
	}

	if _, err := ioutil.ReadDir(utils.QT_DIR()); err != nil && !utils.UsePkgConfig() {
		utils.Log.WithError(err).Panic("failed to find Qt dir, did you export QT_DIR?")
	}

	if !strings.HasSuffix(buildTarget, "-docker") {
		switch runtime.GOOS {
		case "darwin":
			{
				if _, err := exec.LookPath("clang++"); err != nil {
					utils.Log.WithError(err).Panic("failed to find clang++, did you install Xcode?")
				}
			}

		case "linux":
			{
				if _, err := exec.LookPath("g++"); err != nil {
					utils.Log.WithError(err).Panic("failed to find g++, did you install g++?")
				}
			}

		case "windows":
			{
				if _, err := exec.LookPath("g++"); err != nil {
					utils.Log.WithError(err).Panic("failed to find g++, did you add the directory that contains g++ to your PATH?")
				}
			}
		}
	}

	switch buildTarget {
	case "android":
		{
			if _, err := ioutil.ReadDir(utils.ANDROID_SDK_DIR()); err != nil {
				utils.Log.WithError(err).Panic("failed to find android-sdk dir, did you export ANDROID_SDK_DIR?")
			}

			if _, err := ioutil.ReadDir(utils.ANDROID_NDK_DIR()); err != nil {
				utils.Log.WithError(err).Panic("failed to find android-ndk dir, did you export ANDROID_NDK_DIR?")
			}

			if _, err := ioutil.ReadDir(utils.JDK_DIR()); err != nil {
				utils.Log.WithError(err).Panic("failed to find jdk dir, did you export JDK_DIR?")
			}
		}

	case "sailfish", "sailfish-emulator":
		{
			if _, err := ioutil.ReadDir(utils.VIRTUALBOX_DIR()); err != nil {
				utils.Log.WithError(err).Panic("failed to find virtualbox dir, did you export VIRTUALBOX_DIR?")
			}

			if _, err := ioutil.ReadDir(utils.SAILFISH_DIR()); err != nil {
				utils.Log.WithError(err).Panic("failed to find sailfish-sdk dir, did you export SAILFISH_DIR?")
			}
		}

	case "rpi1", "rpi2", "rpi3":
		{
			//TODO:
		}

	case "windows-docker", "linux-docker", "android-docker":
		{
			if _, err := exec.LookPath("docker"); err != nil {
				utils.Log.WithError(err).Panic("failed to find docker, did you install docker?")
			}
		}
	}
}
