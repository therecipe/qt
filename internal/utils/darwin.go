package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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
		for _, i := range []string{"15", "14", "13", "12", "11", "10"} {
			if ExistsDir(filepath.Join(basePath, fmt.Sprintf("MacOSX10.%v.sdk", i))) {
				return fmt.Sprintf("MacOSX10.%v.sdk", i)
			}
		}
		if ExistsDir(filepath.Join(basePath, "MacOSX.sdk")) {
			return "MacOSX.sdk"
		}
		Log.Errorf("failed to find MacOSX sdk in %v", basePath)
	}
	return ""
}

func IPHONEOS_SDK_DIR() string {
	if runtime.GOOS == "darwin" {
		basePath := filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs")
		for _, i := range []string{"13.0", "12.0", "11.0", "10.1", "10.0", "9.3", "9.2", "9.1", "9.0"} {
			if ExistsDir(filepath.Join(basePath, fmt.Sprintf("iPhoneOS%v.sdk", i))) {
				return fmt.Sprintf("iPhoneOS%v.sdk", i)
			}
		}
		if ExistsDir(filepath.Join(basePath, "iPhoneOS.sdk")) {
			return "iPhoneOS.sdk"
		}
		Log.Errorf("failed to find iPhoneOS sdk in %v", basePath)
	}
	return ""
}

func IPHONESIMULATOR_SDK_DIR() string {
	if runtime.GOOS == "darwin" {
		basePath := filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs")
		for _, i := range []string{"13.0", "12.0", "11.0", "10.1", "10.0", "9.3", "9.2", "9.1", "9.0"} {
			if ExistsDir(filepath.Join(basePath, fmt.Sprintf("iPhoneSimulator%v.sdk", i))) {
				return fmt.Sprintf("iPhoneSimulator%v.sdk", i)
			}
		}
		if ExistsDir(filepath.Join(basePath, "iPhoneSimulator.sdk")) {
			return "iPhoneSimulator.sdk"
		}
		Log.Errorf("failed to find iPhoneSimulator sdk in %v", basePath)
	}
	return ""
}

func QT_HOMEBREW() bool {
	return os.Getenv("QT_HOMEBREW") == "true" || isHomeBrewQtDir()
}

func QT_MACPORTS() bool {
	return os.Getenv("QT_MACPORTS") == "true"
}

func QT_NIX() bool {
	_, ok := os.LookupEnv("NIX_STORE")
	return ok
}

func isHomeBrewQtDir() bool {
	return ExistsFile(filepath.Join(QT_DIR(), "INSTALL_RECEIPT.json"))
}

func QT_DARWIN_DIR() string {
	path := qT_DARWIN_DIR()
	if ExistsDir(path) {
		return path
	}
	return strings.Replace(path, QT_VERSION_MAJOR(), QT_VERSION(), -1)
}

var qt_darwin_dir_nix string

func qT_DARWIN_DIR() string {
	if QT_HOMEBREW() {
		if isHomeBrewQtDir() {
			return QT_DIR()
		}
		return "/usr/local/opt/qt5"
	}
	if QT_MACPORTS() {
		return "/opt/local/libexec/qt5"
	}
	if QT_NIX() {
		if len(qt_darwin_dir_nix) == 0 {
			qt_darwin_dir_nix = strings.TrimSpace(RunCmd(exec.Command(ToolPath("qmake", "darwin"), "-query", "QT_INSTALL_PREFIX"), "nix qt dir"))
		}
		return qt_darwin_dir_nix
	}
	return filepath.Join(QT_DIR(), fmt.Sprintf("%v/clang_64", QT_VERSION_MAJOR()))
}
