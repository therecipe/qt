package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func check(buildTarget string) {
	utils.Log.Infof("running setup/check %v", buildTarget)

	switch buildTarget {
	case "desktop", "android", "ios", "ios-simulator",
		"sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "windows", "darwin", "linux",
		"linux-docker", "windows-docker", "android-docker":
		{
			var buildDocker = strings.HasSuffix(buildTarget, "-docker")
			switch buildTarget {
			case "windows":
				{
					if runtime.GOOS == "windows" && !buildDocker {
					} else if runtime.GOOS == "linux" || buildDocker {
					} else {
						utils.Log.Fatalf("%v is currently not supported as a deploy target on %v", buildTarget, runtime.GOOS)
					}
				}

			case "darwin", "ios", "ios-simulator":
				{
					if runtime.GOOS == "darwin" && !buildDocker {
					} else {
						utils.Log.Fatalf("%v is currently not supported as a deploy target on %v (not even with docker)", buildTarget, runtime.GOOS)
					}
				}

			case "linux":
				{
					if runtime.GOOS == "linux" && !buildDocker {
					} else if buildDocker {
					} else {
						utils.Log.Fatalf("%v is currently not supported as a deploy target on %v", buildTarget, runtime.GOOS)
					}
				}
			}
		}

	default:
		{
			utils.Log.Panicf("failed to recognize build target %v", buildTarget)
		}
	}

	utils.Log.Infoln("GOOS:", runtime.GOOS)
	utils.Log.Infoln("GOARCH:", runtime.GOARCH)
	utils.Log.Infoln("GOVERSION:", runtime.Version())
	utils.Log.Infoln("GOROOT:", runtime.GOROOT())
	utils.Log.Infoln("GOPATH:", utils.MustGoPath())

	var hash = exec.Command("git", "rev-parse", "--verify", "HEAD")
	hash.Dir = utils.GoQtPkgPath()
	utils.Log.Infoln("HASH:", strings.TrimSpace(utils.RunCmdOptional(hash, "get git hash")))

	utils.Log.Infoln("QT_DIR:", utils.QT_DIR())
	utils.Log.Infoln("QT_STUB:", utils.QT_STUB())

	switch buildTarget {
	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					utils.Log.Infoln("QT_HOMEBREW:", os.Getenv("QT_HOMEBREW"))
					utils.Log.Infoln("IsHomeBrewQtDir:", utils.IsHomeBrewQtDir())
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
					utils.Log.Infoln("QT_MXE_ARCH:", os.Getenv("QT_MXE_ARCH"))
				}

			case "windows":
				{
					utils.Log.Infoln("QT_MSYS2:", os.Getenv("QT_MSYS2"))
					utils.Log.Infoln("QT_MSYS2_DIR:", os.Getenv("QT_MSYS2_DIR"))
					utils.Log.Infoln("IsMsys2QtDir:", utils.IsMsys2QtDir())
					utils.Log.Infoln("UseMsys2:", utils.UseMsys2())
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

	if _, err := ioutil.ReadDir(utils.QT_DIR()); err != nil && !(utils.UsePkgConfig() || utils.UseHomeBrew() || utils.UseMsys2()) {
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
