package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func QT_VERSION() string {
	if version := os.Getenv("QT_VERSION"); version != "" {
		return version
	}

	return "5.8.0"
}

func QT_VERSION_MAJOR() string {
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

func QT_STUB() bool {
	return strings.ToLower(os.Getenv("QT_STUB")) == "true" || UseMsys2()
}

func QT_DEBUG() bool {
	return strings.ToLower(os.Getenv("QT_DEBUG")) == "true"
}

func CheckBuildTarget(buildTarget string) {
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
						Log.Fatalf("%v is currently not supported as a deploy target on %v", buildTarget, runtime.GOOS)
					}
				}

			case "darwin", "ios", "ios-simulator":
				{
					if runtime.GOOS == "darwin" && !buildDocker {
					} else {
						Log.Fatalf("%v is currently not supported as a deploy target on %v (not even with docker)", buildTarget, runtime.GOOS)
					}
				}

			case "linux":
				{
					if runtime.GOOS == "linux" && !buildDocker {
					} else if buildDocker {
					} else {
						Log.Fatalf("%v is currently not supported as a deploy target on %v", buildTarget, runtime.GOOS)
					}
				}
			}
		}

	default:
		{
			Log.Panicf("failed to recognize build target %v", buildTarget)
		}
	}

	if buildTarget == "android" || strings.HasPrefix(buildTarget, "ios") {
		switch {
		case UseMsys2():
			{
				Log.Fatalf("%v is not supported as a deploy target on %v with MSYS2 -> install the official Qt version instead and try again", buildTarget, runtime.GOOS)
			}

		case UseHomeBrew():
			{
				Log.Fatalf("%v is not supported as a deploy target on %v with HomeBrew -> install the official Qt version instead and try again", buildTarget, runtime.GOOS)
			}

		case UsePkgConfig():
			{
				Log.Fatalf("%v is not supported as a deploy target on %v with PkgConfig -> install the official Qt version instead and try again", buildTarget, runtime.GOOS)
			}
		}
	}
}

func IsCI() bool {
	return strings.ToLower(os.Getenv("CI")) == "true"
}

func QT_QMAKE_DIR() string {
	if dir := os.Getenv("QT_QMAKE_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return ""
}

func QT_QMAKE_CGO() bool {
	return strings.ToLower(os.Getenv("QT_QMAKE_CGO")) == "true"
}

func QT_DOCKER() bool {
	return strings.ToLower(os.Getenv("QT_DOCKER")) == "true"
}

func ToolPath(tool, target string) string {
	//TODO: only temporary
	if target == "desktop" {
		target = runtime.GOOS
	}
	//

	switch tool {
	case "qmake":
		if dir := QT_QMAKE_DIR(); dir != "" {
			return filepath.Join(dir, tool)
		}
	}

	switch target {
	case "darwin":
		return filepath.Join(QT_DARWIN_DIR(), "bin", tool)
	case "windows":
		if runtime.GOOS == target {
			if UseMsys2() {
				return filepath.Join(QT_MSYS2_DIR(), "bin", tool)
			}
			return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "mingw53_32", "bin", tool)
		}
		return filepath.Join(QT_MXE_DIR(), "usr", QT_MXE_TRIPLET(), "qt5", "bin", tool)
	case "linux":
		if UsePkgConfig() {
			if QT_QMAKE_CGO() {
				return filepath.Join(strings.TrimSpace(RunCmd(exec.Command("pkg-config", "--variable=host_bins", "Qt5Core"), "cgo.LinuxPkgConfig_hostBins")), tool)
			}
			return filepath.Join(QT_MISC_DIR(), "bin", tool)
		}
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "gcc_64", "bin", tool)
	case "ios", "ios-simulator":
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "ios", "bin", tool)
	case "android":
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "android_armv7", "bin", tool)
	case "sailfish":
		return filepath.Join(os.Getenv("HOME"), ".config", "SailfishOS-SDK", "mer-sdk-tools", "MerSDK", "SailfishOS-armv7hl", tool)
	case "sailfish-emulator":
		return filepath.Join(os.Getenv("HOME"), ".config", "SailfishOS-SDK", "mer-sdk-tools", "MerSDK", "SailfishOS-i486", tool)
	case "asteroid":
	case "rp1", "rpi2", "rpi3":
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), target, "bin", tool)
	}
	return ""
}
