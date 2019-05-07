package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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
