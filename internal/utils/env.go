package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func QT_VERSION() string {
	if version := os.Getenv("QT_VERSION"); version != "" {
		return version
	}

	if QT_PKG_CONFIG() {
		return strings.TrimSpace(RunCmd(exec.Command("pkg-config", "--modversion", "Qt5Core"), "cgo.LinuxPkgConfig_modVersion"))
	}

	return "5.8.0"
}

func QT_VERSION_NUM() int {
	vmaj, _ := strconv.Atoi(string(QT_VERSION()[0]))
	vmaj *= 1000
	vmin, _ := strconv.Atoi(strings.Replace(QT_VERSION()[1:], ".", "", -1))
	return vmaj + vmin
}

func QT_VERSION_MAJOR() string {
	if version := os.Getenv("QT_VERSION_MAJOR"); version != "" {
		return version
	}
	if QT_VERSION_NUM() >= 5091 {
		return QT_VERSION()
	}
	return strings.Join(strings.Split(QT_VERSION(), ".")[:2], ".")
}

func QT_DIR() string {
	if dir := os.Getenv("QT_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	if runtime.GOOS == "windows" {
		return "C:\\Qt\\Qt" + QT_VERSION()
	}
	return filepath.Join(os.Getenv("HOME"), "Qt"+QT_VERSION())
}

func QT_FAT() bool {
	return strings.ToLower(os.Getenv("QT_FAT")) == "true"
}

func QT_STUB() bool {
	return (strings.ToLower(os.Getenv("QT_STUB")) == "true" || QT_MSYS2()) && !QT_FAT()
}

func QT_DEBUG() bool {
	return strings.ToLower(os.Getenv("QT_DEBUG")) == "true"
}

func QT_DEBUG_QML() bool {
	return strings.ToLower(os.Getenv("QT_DEBUG_QML")) == "true"
}

func CheckBuildTarget(buildTarget string) {
	switch buildTarget {
	case "android", "android-emulator",
		"ios", "ios-simulator",
		"sailfish", "sailfish-emulator", "asteroid",
		"rpi1", "rpi2", "rpi3",
		"windows", "darwin", "linux",
		"homebrew":

	default:
		if !strings.Contains(buildTarget, "_") {
			Log.Panicf("failed to recognize build target %v", buildTarget)
		}
	}

	if buildTarget != runtime.GOOS && !strings.Contains(buildTarget, "_") {
		switch {
		case QT_MSYS2():
			Log.Fatalf("%v is not supported as a deploy target on %v with MSYS2 -> install the official Qt version instead and try again", buildTarget, runtime.GOOS)

		case QT_HOMEBREW():
			Log.Fatalf("%v is not supported as a deploy target on %v with HomeBrew -> install the official Qt version instead and try again", buildTarget, runtime.GOOS)

		case QT_PKG_CONFIG():
			Log.Fatalf("%v is not supported as a deploy target on %v with PkgConfig -> install the official Qt version instead and try again", buildTarget, runtime.GOOS)
		}
	}
}

func CI() bool {
	return strings.ToLower(os.Getenv("CI")) == "true"
}

func QT_QMAKE_DIR() string {
	if dir := os.Getenv("QT_QMAKE_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return ""
}

func QT_DOCKER() bool {
	return strings.ToLower(os.Getenv("QT_DOCKER")) == "true"
}

func QT_VAGRANT() bool {
	return strings.ToLower(os.Getenv("QT_VAGRANT")) == "true"
}

//TODO: use qmake props
func ToolPath(tool, target string) string {
	if target == "sailfish" || target == "sailfish-emulator" {
		target = runtime.GOOS
	}

	if dir := QT_QMAKE_DIR(); dir != "" {
		return filepath.Join(dir, tool)
	}

	switch target {
	case "darwin":
		return filepath.Join(QT_DARWIN_DIR(), "bin", tool)
	case "windows":
		if runtime.GOOS == target {
			if QT_MSYS2() {
				if QT_MSYS2_STATIC() {
					return filepath.Join(QT_MSYS2_DIR(), "qt5-static", "bin", tool)
				}
				return filepath.Join(QT_MSYS2_DIR(), "bin", tool)
			}
			return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "mingw53_32", "bin", tool)
		}
		return filepath.Join(QT_MXE_DIR(), "usr", QT_MXE_TRIPLET(), "qt5", "bin", tool)
	case "linux":
		if QT_PKG_CONFIG() {
			return filepath.Join(strings.TrimSpace(RunCmd(exec.Command("pkg-config", "--variable=host_bins", "Qt5Core"), "cgo.LinuxPkgConfig_hostBins")), tool)
		}
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "gcc_64", "bin", tool)
	case "ios", "ios-simulator":
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "ios", "bin", tool)
	case "android":
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "android_armv7", "bin", tool)
	case "android-emulator":
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "android_x86", "bin", tool)
		/*
			case "sailfish":
				return filepath.Join(os.Getenv("HOME"), ".config", "SailfishOS-SDK", "mer-sdk-tools", "MerSDK", "SailfishOS-armv7hl", tool)
			case "sailfish-emulator":
				return filepath.Join(os.Getenv("HOME"), ".config", "SailfishOS-SDK", "mer-sdk-tools", "MerSDK", "SailfishOS-i486", tool)
		*/
	case "asteroid":
	case "rp1", "rpi2", "rpi3":
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), target, "bin", tool)
	}
	return ""
}

func QT_WEBKIT() bool {
	return strings.ToLower(os.Getenv("QT_WEBKIT")) == "true"
}
