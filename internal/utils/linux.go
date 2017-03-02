package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func UsePkgConfig() bool {
	return strings.ToLower(os.Getenv("QT_PKG_CONFIG")) == "true"
}

func QT_DOC_DIR() string {
	if dir := os.Getenv("QT_DOC_DIR"); dir != "" {
		return filepath.Clean(dir)
	}

	switch LinuxDistro() {
	case "arch":
		{
			return "/usr/share/doc/qt"
		}

	case "fedora":
		{
			return "/usr/share/doc/qt5"
		}

	case "suse":
		{
			return "/usr/share/doc/packages/qt5"
		}

	case "ubuntu":
		{
			return "/usr/share/qt5/doc"
		}

	default:
		{
			Log.Error("failed to detect the Linux distro")
			return ""
		}
	}
}

func QT_MISC_DIR() string {
	if dir := os.Getenv("QT_MISC_DIR"); dir != "" {
		return filepath.Clean(dir)
	}

	switch LinuxDistro() {
	case "arch":
		return filepath.Join(strings.TrimSpace(RunCmd(exec.Command("pkg-config", "--variable=libdir", "Qt5Core"), "cgo.LinuxPkgConfig_libDir")), "qt")
	case "fedora", "suse", "ubuntu":
		return strings.TrimSuffix(strings.TrimSpace(RunCmd(exec.Command("pkg-config", "--variable=host_bins", "Qt5Core"), "cgo.LinuxPkgConfig_hostBins")), "/bin")
	}
	Log.Error("failed to detect the Linux distro")
	return ""
}

func LinuxDistro() string {

	if _, err := exec.LookPath("pacman"); err == nil {
		return "arch"
	}

	if _, err := exec.LookPath("yum"); err == nil {
		return "fedora"
	}

	if _, err := exec.LookPath("zypper"); err == nil {
		return "suse"
	}

	if _, err := exec.LookPath("apt-get"); err == nil {
		return "ubuntu"
	}

	Log.Error("failed to detect the Linux distro")
	return ""
}

func QT_MXE_ARCH() string {
	if arch := os.Getenv("QT_MXE_ARCH"); arch == "386" || arch == "amd64" {
		return arch
	}
	return "386"
}

func QT_MXE_STATIC() bool {
	return strings.ToLower(os.Getenv("QT_MXE_STATIC")) == "true"
}

func QT_MXE_TRIPLET() string {
	var prefix = "i686"
	if QT_MXE_ARCH() == "amd64" {
		prefix = "x86_64"
	}
	var suffix = "shared"
	if QT_MXE_STATIC() {
		suffix = "static"
	}
	return fmt.Sprintf("%v-w64-mingw32.%v", prefix, suffix)
}

func QT_MXE_DIR() string {
	if dir := os.Getenv("QT_MXE_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join("/usr", "lib", "mxe")
}

func QT_MXE_BIN(tool string) string {
	return filepath.Join(QT_MXE_DIR(), "usr", "bin", fmt.Sprintf("%v-%v", QT_MXE_TRIPLET(), tool))
}
