package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func JDK_DIR() string {
	if dir, ok := os.LookupEnv("JDK_DIR"); ok {
		return filepath.Clean(dir)
	}
	if dir, ok := os.LookupEnv("JAVA_HOME"); ok {
		return filepath.Clean(dir)
	}

	for _, l := range strings.Split(RunCmd(exec.Command("java", "-XshowSettings:properties", "-version"), "deploy.jdk"), "\n") {
		l = strings.TrimSpace(l)
		if strings.HasPrefix(l, "java.home") {
			return strings.Split(l, " = ")[1]
		}
	}

	return filepath.Join(os.Getenv("HOME"), "jdk")
}

func ANDROID_SDK_DIR() string {
	if dir, ok := os.LookupEnv("ANDROID_SDK_DIR"); ok {
		return filepath.Clean(dir)
	}
	if dir, ok := os.LookupEnv("ANDROID_SDK_ROOT"); ok {
		return filepath.Clean(dir)
	}
	switch runtime.GOOS {
	case "windows":
		return windowsSystemDrive() + "\\android-sdk-windows"
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "android-sdk-macosx")
	default:
		return filepath.Join(os.Getenv("HOME"), "android-sdk-linux")
	}
}

func ANDROID_NDK_DIR() string {
	if dir, ok := os.LookupEnv("ANDROID_NDK_DIR"); ok {
		return filepath.Clean(dir)
	}
	if dir, ok := os.LookupEnv("ANDROID_NDK_ROOT"); ok {
		return filepath.Clean(dir)
	}
	if runtime.GOOS == "windows" {
		return windowsSystemDrive() + "\\android-ndk-r18b"
	}
	return filepath.Join(os.Getenv("HOME"), "android-ndk-r18b")
}

func ANDROID_NDK_MAJOR_VERSION() int {

	rawVersion := strings.Split(RunCmd(exec.Command("sed", "-En", "-e", `s/^Pkg.Revision\s*=\s*([0-9a-f]+)/r\1/p`, fmt.Sprintf("%s/source.properties", ANDROID_NDK_DIR())), "Android NDK version parsing"), ".")
	// discards the prefix in r<Number> major version
	major, err := strconv.Atoi(rawVersion[0][1:])
	if err != nil {
		fields := logrus.Fields{"_func": "ANDROID_NDK_MAJOR_VERSION", "major_version": rawVersion[0]}
		Log.WithError(err).WithFields(fields).Error("Failed to parse major version of Android NDK")
		os.Exit(1)
	}
	return major
}

// ANDROID_NDK_REQUIRE_NOSTDLIBPP_LDFLAG returns true if the -nostdlib++ flag is required by Android NDK and false otherwise
func ANDROID_NDK_REQUIRE_NOSTDLIBPP_LDFLAG() bool {
	if ANDROID_NDK_MAJOR_VERSION() > 19 {
		return true
	}
	return false
}

// ANDROID_NDK_NOSTDLIBPP_LDFLAG return -nostdlib++ or empty string if that flag is no required by Android NDK
func ANDROID_NDK_NOSTDLIBPP_LDFLAG() string {
	if ANDROID_NDK_MAJOR_VERSION() > 19 {
		return "-nostdlib++"
	}
	return ""
}
