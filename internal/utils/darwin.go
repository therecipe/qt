package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func XCODE_DIR() string {
	if dir, ok := os.LookupEnv("XCODE_DIR"); ok {
		return filepath.Clean(dir)
	}
	return filepath.Join("/Applications/Xcode.app")
}

func MACOS_SDK_DIR() string {
	if runtime.GOOS == "darwin" {
		basePath := filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs")
		for _, i := range []string{"13", "12", "11", "10"} {
			if ExistsDir(filepath.Join(basePath, fmt.Sprintf("MacOSX10.%v.sdk", i))) {
				return fmt.Sprintf("MacOSX10.%v.sdk", i)
			}
		}
		Log.Errorf("failed to find MacOSX sdk in %v", basePath)
	}
	return ""
}

func IPHONEOS_SDK_DIR() string {
	if runtime.GOOS == "darwin" {
		basePath := filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs")
		for _, i := range []string{"10.3", "10.2", "10.1", "10.0", "9.3", "9.2", "9.1", "9.0"} {
			if ExistsDir(filepath.Join(basePath, fmt.Sprintf("iPhoneOS%v.sdk", i))) {
				return fmt.Sprintf("iPhoneOS%v.sdk", i)
			}
		}
		Log.Errorf("failed to find iPhoneOS sdk in %v", basePath)
	}
	return ""
}

func IPHONESIMULATOR_SDK_DIR() string {
	if runtime.GOOS == "darwin" {
		basePath := filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs")
		for _, i := range []string{"10.3", "10.2", "10.1", "10.0", "9.3", "9.2", "9.1", "9.0"} {
			if ExistsDir(filepath.Join(basePath, fmt.Sprintf("iPhoneSimulator%v.sdk", i))) {
				return fmt.Sprintf("iPhoneSimulator%v.sdk", i)
			}
		}
		Log.Errorf("failed to find iPhoneSimulator sdk in %v", basePath)
	}
	return ""
}

func QT_HOMEBREW() bool {
	return os.Getenv("QT_HOMEBREW") == "true" || isHomeBrewQtDir()
}

func isHomeBrewQtDir() bool {
	return ExistsFile(filepath.Join(QT_DIR(), "INSTALL_RECEIPT.json"))
}

func QT_DARWIN_DIR() string {
	if QT_HOMEBREW() {
		if isHomeBrewQtDir() {
			return QT_DIR()
		}
		return "/usr/local/opt/qt5"
	}
	return filepath.Join(QT_DIR(), fmt.Sprintf("%v/clang_64", QT_VERSION_MAJOR()))
}
