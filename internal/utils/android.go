package utils

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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

	var (
		propsRegex = regexp.MustCompile(`[#].*\\n|\\s+\\n|\\S+[=]|.*\n`)
		pairRegex  = regexp.MustCompile(`\s+=\s+`)
	)

	source := filepath.Join(ANDROID_NDK_DIR(), "source.properties")
	props, err := ioutil.ReadFile(source)
	if err != nil {
		fields := logrus.Fields{"_func": "ANDROID_NDK_MAJOR_VERSION", "source": source}
		Log.WithError(err).WithFields(fields).Error("Failed to read source.properties file of Android NDK")
		os.Exit(1)
	}

	var version_string string
	pairs := propsRegex.FindAll(props, -1)
	for _, pair := range pairs {
		str_pair := strings.TrimSpace(string(pair))
		if strings.HasPrefix(str_pair, "Pkg.Revision") {
			if aux := pairRegex.Split(str_pair, -1); len(aux) == 2 {
				version_string = aux[1]
				break
			} else {
				fields := logrus.Fields{"_func": "ANDROID_NDK_MAJOR_VERSION", "property_pair": str_pair}
				Log.WithError(err).WithFields(fields).Error("Failed to proccess source.properties file of Android NDK")
				os.Exit(1)
			}
		}
	}

	rawVersion := strings.Split(version_string, ".")
	major, err := strconv.Atoi(rawVersion[0])
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

// ANDROID_NDK_PLATFORM returns the Android API level to use in the building process
func ANDROID_NDK_PLATFORM() string {

	// default value as set in cmd.go BuildEnv function
	if QT_FELGO() {
		return "android-16"
	}
	// if env var exists and is not empty use env var of the system
	if value, exist := os.LookupEnv("ANDROID_NDK_PLATFORM"); exist && value != "" {
		return value
	}

	// default value as set in cmd.go BuildEnv function
	return "android-21"
}
