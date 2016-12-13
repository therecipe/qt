package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func XCODE_DIR() string {
	if dir := os.Getenv("XCODE_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join("/Applications/Xcode.app")
}

func MACOS_SDK_DIR() string {
	if runtime.GOOS == "darwin" {
		var basePath = filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs")
		for _, i := range []string{"13", "12", "11", "10"} {
			if _, err := ioutil.ReadDir(filepath.Join(basePath, fmt.Sprintf("MacOSX10.%v.sdk", i))); err == nil {
				return fmt.Sprintf("MacOSX10.%v.sdk", i)
			}
		}
		Log.Errorf("failed to find MacOSX sdk in %v", basePath)
	}
	return ""
}

func IPHONEOS_SDK_DIR() string {
	if runtime.GOOS == "darwin" {
		var basePath = filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs")
		for _, i := range []string{"10.3", "10.2", "10.1", "10.0", "9.3", "9.2", "9.1", "9.0"} {
			if _, err := ioutil.ReadDir(filepath.Join(basePath, fmt.Sprintf("iPhoneOS%v.sdk", i))); err == nil {
				return fmt.Sprintf("iPhoneOS%v.sdk", i)
			}
		}
		Log.Errorf("failed to find iPhoneOS sdk in %v", basePath)
	}
	return ""
}

func IPHONESIMULATOR_SDK_DIR() string {
	if runtime.GOOS == "darwin" {
		var basePath = filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs")
		for _, i := range []string{"10.3", "10.2", "10.1", "10.0", "9.3", "9.2", "9.1", "9.0"} {
			if _, err := ioutil.ReadDir(filepath.Join(basePath, fmt.Sprintf("iPhoneSimulator%v.sdk", i))); err == nil {
				return fmt.Sprintf("iPhoneSimulator%v.sdk", i)
			}
		}
		Log.Errorf("failed to find iPhoneSimulator sdk in %v", basePath)
	}
	return ""
}

func UseHomeBrew() bool {
	return strings.ToLower(os.Getenv("QT_HOMEBREW")) == "true" || IsHomeBrewQtDir()
}

func IsHomeBrewQtDir() bool {
	return ExistsFile(filepath.Join(QT_DIR(), "INSTALL_RECEIPT.json"))
}

func QT_DARWIN_DIR() string {
	if UseHomeBrew() {
		if IsHomeBrewQtDir() {
			return QT_DIR()
		}
		return "/usr/local/opt/qt5"
	}
	return filepath.Join(QT_DIR(), "5.7/clang_64")
}
