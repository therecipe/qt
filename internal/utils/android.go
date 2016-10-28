package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func JDK_DIR() string {
	if dir := os.Getenv("JDK_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	if dir := os.Getenv("JAVA_HOME"); dir != "" {
		return filepath.Clean(dir)
	}

	switch runtime.GOOS {
	case "windows":
		{
			return fmt.Sprintf("C:\\Program Files\\Java\\jdk%v", strings.Split(RunCmd(exec.Command("java", "-version"), "deploy.jdk"), "\"")[1])
		}

	case "darwin":
		{
			return fmt.Sprintf("/Library/Java/JavaVirtualMachines/jdk%v.jdk/Contents/Home", strings.Split(RunCmd(exec.Command("java", "-version"), "deploy.jdk"), "\"")[1])
		}

	default:
		{
			return filepath.Join(os.Getenv("HOME"), "jdk")
		}
	}
}

func ANDROID_SDK_DIR() string {
	if dir := os.Getenv("ANDROID_SDK_DIR"); dir != "" {
		return filepath.Clean(dir)
	}

	switch runtime.GOOS {
	case "windows":
		{
			return "C:\\android-sdk-windows"
		}

	case "darwin":
		{
			return filepath.Join(os.Getenv("HOME"), "android-sdk-macosx")
		}

	default:
		{
			return filepath.Join(os.Getenv("HOME"), "android-sdk-linux")
		}
	}
}

func ANDROID_NDK_DIR() string {
	if dir := os.Getenv("ANDROID_NDK_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	if runtime.GOOS == "windows" {
		return "C:\\android-ndk-r12b"
	}
	return filepath.Join(os.Getenv("HOME"), "android-ndk-r12b")
}
