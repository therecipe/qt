package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func MakeFolder(dir string) {
	var err = os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println("file.MakeFolder", err)
	}
}

func RemoveAll(name string) {
	var err = os.RemoveAll(name)
	if err != nil {
		fmt.Println("file.RemoveAll", err)
	}
}

func Save(name, data string) {
	var err = ioutil.WriteFile(name, []byte(data), 0777)
	if err != nil {
		fmt.Println("file.Save", err)
	}
}

func SaveBytes(name string, data []byte) {
	var err = ioutil.WriteFile(name, data, 0777)
	if err != nil {
		fmt.Println("file.SaveBytes", err)
	}
}

func Load(name string) string {
	var b, err = ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("file.Load", err)
	}
	return string(b)
}

func GetAbsPath(appPath string) string {
	var wd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Clean(filepath.Join(wd, appPath))
}

func GoQtPkgPath(s ...string) string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "therecipe", "qt", filepath.Join(s...))
}

func Exists(name string) bool {
	var _, err = ioutil.ReadFile(name)
	return err == nil
}

func QT_DIR() string {
	if dir := os.Getenv("QT_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	//TODO: remove with 5.8
	if dir := os.Getenv("QT_INSTALL_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	//
	if runtime.GOOS == "windows" {
		return filepath.Join("C:\\", "Qt", "Qt5.7.0")
	}
	return filepath.Join(os.Getenv("HOME"), "Qt5.7.0")
}

func ANDROID_SDK_DIR() string {
	if dir := os.Getenv("ANDROID_SDK_DIR"); dir != "" {
		return filepath.Clean(dir)
	}

	switch runtime.GOOS {
	case "windows":
		{
			return filepath.Join("C:\\", "android-sdk-windows")
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
		return filepath.Join("C:\\", "android-ndk-r12b")
	}
	return filepath.Join(os.Getenv("HOME"), "android-ndk-r12b")
}

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
			var version = strings.Split(RunCmd(exec.Command("java", "-version"), "deploy.jdk"), "\"")[1]
			return filepath.Join("C:\\", "Program Files", "Java", fmt.Sprintf("jdk%v", version))
		}

	case "darwin":
		{
			var version = strings.Split(RunCmd(exec.Command("java", "-version"), "deploy.jdk"), "\"")[1]
			return filepath.Join("/Library", "Java", "JavaVirtualMachines", fmt.Sprintf("jdk%v.jdk", version), "Contents", "Home")
		}

	case "linux":
		{
			return filepath.Join(os.Getenv("HOME"), "jdk")
		}

	default:
		{
			return ""
		}
	}
}

func VIRTUALBOX_DIR() string {
	if dir := os.Getenv("VIRTUALBOX_DIR"); dir != "" {
		return filepath.Clean(dir)
	}

	if runtime.GOOS == "windows" {
		return filepath.Join("C:\\", "Program Files", "Oracle", "VirtualBox")
	}

	var path, err = exec.LookPath("vboxmanage")
	if err != nil {
		fmt.Println("VIRTUALBOX_DIR:", "vboxmanage not found")
	}
	path, err = filepath.Abs(filepath.Dir(path))
	if err != nil {
		fmt.Println("VIRTUALBOX_DIR:", "could not create absolute path")
	}

	return path
}

func SAILFISH_DIR() string {
	if dir := os.Getenv("SAILFISH_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	if runtime.GOOS == "windows" {
		return filepath.Join("C:\\", "SailfishOS")
	}
	return filepath.Join(os.Getenv("HOME"), "SailfishOS")
}

func RPI_TOOLS_DIR() string {
	if dir := os.Getenv("RPI_TOOLS_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join(os.Getenv("HOME"), "raspi", "tools")
}

func RPI1_SYSROOT_DIR() string {
	if dir := os.Getenv("RPI1_SYSROOT_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join(os.Getenv("HOME"), "raspi", "sysroot")
}

func RPI2_SYSROOT_DIR() string {
	if dir := os.Getenv("RPI2_SYSROOT_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join(os.Getenv("HOME"), "raspi", "sysroot")
}

func RPI3_SYSROOT_DIR() string {
	if dir := os.Getenv("RPI3_SYSROOT_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join(os.Getenv("HOME"), "raspi", "sysroot")
}

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
			return filepath.Join("/usr", "share", "doc", "qt")
		}

	case "fedora":
		{
			return filepath.Join("/usr", "share", "doc", "qt5")
		}

	case "suse":
		{
			return filepath.Join("/usr", "share", "doc", "packages", "qt5")
		}

	case "ubuntu":
		{
			return filepath.Join("/usr", "share", "qt5", "doc")
		}

	default:
		{
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
		{
			return filepath.Join(strings.TrimSpace(RunCmd(exec.Command("pkg-config", "--variable=libdir", "Qt5Core"), "cgo.LinuxPkgConfig_libDir")), "qt")
		}

	case "fedora", "suse", "ubuntu":
		{
			return strings.TrimSuffix(strings.TrimSpace(RunCmd(exec.Command("pkg-config", "--variable=host_bins", "Qt5Core"), "cgo.LinuxPkgConfig_hostBins")), "/bin")
		}

	default:
		{
			return ""
		}
	}
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

	fmt.Println("couldn't detect distro")
	os.Exit(1)
	return ""
}

func XCODE_DIR() string {
	if dir := os.Getenv("XCODE_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join("/Applications", "Xcode.app")
}

func MACOS_SDK_DIR() string {

	if runtime.GOOS != "darwin" {
		return "MacOSX.sdk"
	}

	var basePath = filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs")
	for _, i := range []string{"12", "11", "10"} {
		if _, err := ioutil.ReadDir(filepath.Join(basePath, fmt.Sprintf("MacOSX10.%v.sdk", i))); err == nil {
			return fmt.Sprintf("MacOSX10.%v.sdk", i)
		}
	}

	fmt.Println("couldn't find MacOSX sdk")
	os.Exit(1)
	return ""
}

func IPHONEOS_SDK_DIR() string {

	if runtime.GOOS != "darwin" {
		return "iPhoneOS.sdk"
	}

	var basePath = filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs")
	for _, i := range []string{"10.1", "10.0", "9.3", "9.2", "9.1", "9.0"} {
		if _, err := ioutil.ReadDir(filepath.Join(basePath, fmt.Sprintf("iPhoneOS%v.sdk", i))); err == nil {
			return fmt.Sprintf("iPhoneOS%v.sdk", i)
		}
	}

	fmt.Println("couldn't find iPhoneOS sdk")
	os.Exit(1)
	return ""
}

func IPHONESIMULATOR_SDK_DIR() string {

	if runtime.GOOS != "darwin" {
		return "iPhoneSimulator.sdk"
	}

	var basePath = filepath.Join(XCODE_DIR(), "Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs")
	for _, i := range []string{"10.1", "10.0", "9.3", "9.2", "9.1", "9.0"} {
		if _, err := ioutil.ReadDir(filepath.Join(basePath, fmt.Sprintf("iPhoneSimulator%v.sdk", i))); err == nil {
			return fmt.Sprintf("iPhoneSimulator%v.sdk", i)
		}
	}

	fmt.Println("couldn't find iPhoneSimulator sdk")
	os.Exit(1)
	return ""
}

func RunCmd(cmd *exec.Cmd, name string) string {
	var out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", name, out, err)
		os.Exit(1)
	}
	return string(out)
}

func RunCmdOptional(cmd *exec.Cmd, name string) string {
	var out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", name, out, err)
	}
	return string(out)
}
