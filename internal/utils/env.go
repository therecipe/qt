package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var (
	qtVersionCache      string
	qtVersionCacheMutex = new(sync.Mutex)
)

func QT_VERSION() string {
	if version, ok := os.LookupEnv("QT_VERSION"); ok {
		return version
	}

	switch {
	case QT_MXE():
		return "5.13.0"

	case os.Getenv("QT_HOMEBREW") == "true":
		return "5.13.0"

	case QT_MACPORTS(), QT_NIX():
		return "5.11.1"

	case QT_MSYS2(), QT_FELGO():
		return "5.12.0"

	case QT_UBPORTS_VERSION() == "xenial":
		return "5.9.0"

	case QT_SAILFISH():
		return "5.6.3"

	case QT_RPI():
		return "5.7.0"

	case QT_PKG_CONFIG():
		var r string
		qtVersionCacheMutex.Lock()
		if qtVersionCache == "" {
			qtVersionCache = strings.TrimSpace(RunCmd(exec.Command("pkg-config", "--modversion", "Qt5Core"), "utils.PKG_CONFIG_modVersion"))
		}
		r = qtVersionCache
		qtVersionCacheMutex.Unlock()
		return r

	default:
		return "5.13.0"
	}
}

func QT_VERSION_NUM() int {
	version := QT_VERSION()
	vmaj, _ := strconv.Atoi(string(version[0]))
	vmin, _ := strconv.Atoi(strings.Replace(version[1:], ".", "", -1))
	return vmaj*1e3 + vmin
}

func QT_VERSION_MAJOR() string {
	if QT_FELGO() {
		return "Felgo"
	}
	if version, ok := os.LookupEnv("QT_VERSION_MAJOR"); ok {
		return version
	}
	if QT_VERSION_NUM() >= 5091 {
		return QT_VERSION()
	}
	return strings.Join(strings.Split(QT_VERSION(), ".")[:2], ".")
}

func QT_API(def string) string {
	if api, ok := os.LookupEnv("QT_API"); ok {
		return api
	}
	return def
}

func QT_API_NUM(def string) int {
	version := QT_API(def)
	vmaj, _ := strconv.Atoi(string(version[0]))
	vmin, _ := strconv.Atoi(strings.Replace(version[1:], ".", "", -1))
	return vmaj*1e3 + vmin
}

func QT_DIR() string {
	path := qT_DIR()
	if ExistsDir(path) {
		return path
	}
	return strings.Replace(path, QT_VERSION_MAJOR(), QT_VERSION(), -1)
}

var (
	qtInstallPrefixCache      = make(map[string]string)
	qtInstallPrefixCacheMutex = new(sync.Mutex)
)

func QT_INSTALL_PREFIX(target string) (r string) {
	qtInstallPrefixCacheMutex.Lock()
	if _, ok := qtInstallPrefixCache[target]; !ok {
		qtInstallPrefixCache[target] = strings.TrimSpace(RunCmd(exec.Command(ToolPath("qmake", target), "-query", "QT_INSTALL_PREFIX"), fmt.Sprintf("query install prefix path for %v on %v", target, runtime.GOOS)))
	}
	r = qtInstallPrefixCache[target]
	qtInstallPrefixCacheMutex.Unlock()
	return
}

func qT_DIR() string {
	if dir, ok := os.LookupEnv("QT_DIR"); ok {
		return filepath.Clean(dir)
	}

	prefix := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		prefix = windowsSystemDrive() + "\\"
	}

	if QT_FELGO() {
		return filepath.Join(prefix, "Felgo")
	}

	if dir := filepath.Join(prefix, "Qt", "Qt"+QT_VERSION()); ExistsDir(dir) {
		return dir
	}

	if dir := filepath.Join(prefix, "Qt"+QT_VERSION()); ExistsDir(dir) {
		return dir
	}

	return filepath.Join(prefix, "Qt")
}

func QT_FAT() bool {
	return os.Getenv("QT_FAT") == "true"
}

func QT_STUB() bool {
	return os.Getenv("QT_STUB") == "true" && !QT_FAT()
}

func QT_DEBUG() bool {
	return os.Getenv("QT_DEBUG") == "true"
}

func QT_DEBUG_CPP() bool {
	return os.Getenv("QT_DEBUG_CPP") == "true"
}

func QT_DEBUG_QML() bool {
	return os.Getenv("QT_DEBUG_QML") == "true"
}

func QT_DEBUG_CONSOLE() bool {
	return os.Getenv("QT_DEBUG_CONSOLE") == "true"
}

func CheckBuildTarget(buildTarget string, docker bool) {
	switch buildTarget {
	case "android", "android-emulator",
		"ios", "ios-simulator",
		"sailfish", "sailfish-emulator", "asteroid",
		"rpi1", "rpi2", "rpi3",
		"windows", "darwin", "linux",
		"homebrew", "ubports",
		"js", "wasm", "freebsd": //TODO: pkg_config ?
	default:
		if !strings.Contains(buildTarget, "_") {
			Log.Panicf("failed to recognize build target %v", buildTarget)
		}
	}
	if !docker && buildTarget != runtime.GOOS && !strings.Contains(buildTarget, "_") {
		switch {
		case QT_MSYS2():
			Log.Fatalf("%v is not supported as a deploy target on %v with MSYS2 -> install the official Qt version instead and try again", buildTarget, runtime.GOOS)
		case QT_HOMEBREW(), QT_MACPORTS(), QT_NIX():
			Log.Fatalf("%v is not supported as a deploy target on %v with HomeBrew/MacPorts/Nix -> install the official Qt version instead and try again", buildTarget, runtime.GOOS)
		case QT_PKG_CONFIG() && !QT_UBPORTS():
			Log.Fatalf("%v is not supported as a deploy target on %v with PkgConfig -> install the official Qt version instead and try again", buildTarget, runtime.GOOS)
		}
	}
}

func CI() bool {
	return os.Getenv("CI") == "true"
}

func QT_QMAKE_DIR() string {
	if dir, ok := os.LookupEnv("QT_QMAKE_DIR"); ok {
		return filepath.Clean(dir)
	}
	return ""
}

func QT_DOCKER() bool {
	return os.Getenv("QT_DOCKER") == "true"
}

func QT_VAGRANT() bool {
	return os.Getenv("QT_VAGRANT") == "true"
}

//TODO: use qmake props
func ToolPath(tool, target string) string {
	if dir := QT_QMAKE_DIR(); dir != "" {
		return filepath.Join(dir, tool)
	}

	if strings.HasPrefix(target, "sailfish") && !QT_SAILFISH() {
		target = runtime.GOOS
	}

	switch target {
	case "darwin":
		if QT_PKG_CONFIG() {
			return filepath.Join(strings.TrimSpace(RunCmd(exec.Command("pkg-config", "--variable=host_bins", "Qt5Core"), "utils.PKG_CONFIG_hostBins")), tool)
		}
		if QT_NIX() {
			path, _ := exec.LookPath(tool)
			path, _ = filepath.Abs(path)
			return path
		}
		return filepath.Join(QT_DARWIN_DIR(), "bin", tool)
	case "windows":
		if runtime.GOOS == target {
			if QT_MSYS2() {
				if QT_MSYS2_STATIC() {
					return filepath.Join(QT_MSYS2_DIR(), "qt5-static", "bin", tool)
				}
				return filepath.Join(QT_MSYS2_DIR(), "bin", tool)
			}
			path := filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), MINGWDIR(), "bin", tool)
			if !ExistsDir(filepath.Join(QT_DIR(), QT_VERSION_MAJOR())) {
				path = filepath.Join(QT_DIR(), QT_VERSION(), MINGWDIR(), "bin", tool)
			}
			if !ExistsFile(path + ".exe") {
				path = strings.Replace(path, MINGWDIR(), "mingw53_32", -1)
			}
			if !ExistsFile(path + ".exe") {
				path = strings.Replace(path, "mingw53_32", "mingw49_32", -1)
			}
			return path
		}
		return filepath.Join(QT_MXE_DIR(), "usr", QT_MXE_TRIPLET(), "qt5", "bin", tool)
	case "linux", "ubports", "freebsd":
		if QT_PKG_CONFIG() {
			return filepath.Join(strings.TrimSpace(RunCmd(exec.Command("pkg-config", "--variable=host_bins", "Qt5Core"), "utils.PKG_CONFIG_hostBins")), tool)
		}
		path := filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "gcc_64", "bin", tool)
		if !ExistsDir(filepath.Join(QT_DIR(), QT_VERSION_MAJOR())) {
			path = filepath.Join(QT_DIR(), QT_VERSION(), "gcc_64", "bin", tool)
		}
		return path
	case "ios", "ios-simulator":
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "ios", "bin", tool)
	case "android":
		if QT_VERSION_NUM() >= 5140 {
			return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "android", "bin", tool)
		}
		if GOARCH() == "arm64" {
			return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "android_arm64_v8a", "bin", tool)
		}
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "android_armv7", "bin", tool)
	case "android-emulator":
		if QT_VERSION_NUM() >= 5140 {
			return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "android", "bin", tool)
		}
		//TODO: x86_64
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "android_x86", "bin", tool)
	case "sailfish", "sailfish-emulator":
		return filepath.Join("/srv/mer/targets/SailfishOS-"+QT_SAILFISH_VERSION()+"-i486/usr/lib/qt5/bin/", tool)
		//TODO support indirect access on desktop: return filepath.Join(os.Getenv("HOME"), ".config", "SailfishOS-SDK", "mer-sdk-tools", "MerSDK", "SailfishOS-i486", tool)
	case "asteroid":
		//TODO:
	case "rp1", "rpi2", "rpi3":
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), target, "bin", tool)
	case "wasm", "js":
		return filepath.Join(QT_DIR(), QT_VERSION_MAJOR(), "wasm_32", "bin", tool)
	}
	return ""
}

//TODO: detect webkit support automatically
func QT_WEBKIT() bool {
	return os.Getenv("QT_WEBKIT") == "true"
}

func CGO_CFLAGS_ALLOW() string {
	if allowed, ok := os.LookupEnv("CGO_CFLAGS_ALLOW"); ok {
		return allowed
	}
	return ".*"
}

func CGO_CXXFLAGS_ALLOW() string {
	if allowed, ok := os.LookupEnv("CGO_CXXFLAGS_ALLOW"); ok {
		return allowed
	}
	return ".*"
}

func CGO_LDFLAGS_ALLOW() string {
	if allowed, ok := os.LookupEnv("CGO_LDFLAGS_ALLOW"); ok {
		return allowed
	}
	return ".*"
}

func CGO_MSCFLAGS_ALLOW() string {
	if allowed, ok := os.LookupEnv("CGO_MSCFLAGS_ALLOW"); ok {
		return allowed
	}
	return ".*"
}

func CGO_MSCXXFLAGS_ALLOW() string {
	if allowed, ok := os.LookupEnv("CGO_MSCXXFLAGS_ALLOW"); ok {
		return allowed
	}
	return ".*"
}

func CGO_MSLDFLAGS_ALLOW() string {
	if allowed, ok := os.LookupEnv("CGO_MSLDFLAGS_ALLOW"); ok {
		return allowed
	}
	return ".*"
}

func GOARCH() string {
	if arch, ok := os.LookupEnv("GOARCH"); ok {
		return arch
	}
	return runtime.GOARCH
}

func GOARM() (string, bool) {
	if arm, ok := os.LookupEnv("GOARM"); ok {
		return arm, true
	}
	if arm := strings.TrimSpace(RunCmd(exec.Command("go", "env", "GOARM"), "get GOARM")); arm != "" {
		return arm, true
	}
	return "", false
}

func QT_DYNAMIC_SETUP() bool {
	return os.Getenv("QT_DYNAMIC_SETUP") == "true"
}

func GOFLAGS() string {
	if flags, ok := os.LookupEnv("GOFLAGS"); ok {
		return flags
	}
	if UseGOMOD("") {
		return "-mod=vendor"
	}
	return ""
}

func GOMOD(path string) string {
	if useGOMODPath != "" {
		return useGOMODPath
	}
	if mod, ok := os.LookupEnv("GOMOD"); ok {
		return mod
	}
	cmd := exec.Command("go", "env", "GOMOD")
	if path != "" {
		cmd.Dir = path
	}
	return strings.TrimSpace(RunCmd(cmd, "GOMOD"))
}

var (
	useGOMOD      int
	useGOMODPath  string
	useGOMODMutex = new(sync.Mutex)
)

func UseGOMOD(path string) (r bool) {
	if strings.Contains(strings.Replace(path, "\\", "/", -1), "src/"+PackageName) {
		os.Setenv("GO111MODULE", "off")
		return
	}
	if wd, _ := os.Getwd(); path == "" && strings.Contains(strings.Replace(wd, "\\", "/", -1), "src/"+PackageName) {
		os.Setenv("GO111MODULE", "off")
		return
	}
	useGOMODMutex.Lock()
	if useGOMOD == 0 {
		if gm := GOMOD(path); len(gm) != 0 {
			useGOMOD = 1
			useGOMODPath = gm
		} else {
			if path != "" {
				useGOMOD = -1
			}
		}
	}
	r = useGOMOD == 1
	useGOMODMutex.Unlock()
	return
}

func QT_GEN_OPENGL() bool {
	return os.Getenv("QT_GEN_OPENGL") == "true"
}

func QT_GEN_QUICK_EXTRAS() bool {
	return os.Getenv("QT_GEN_QUICK_EXTRAS") == "true"
}

func GoList(args ...string) *exec.Cmd {
	cmd := exec.Command("go", "list")
	if UseGOMOD("") {
		if !(strings.Contains(strings.Join(args, "|"), "github.com/bluszcz/cutego/internal") && !strings.Contains(strings.Join(args, "|"), "github.com/bluszcz/cutego/internal/binding/runtime")) {
			cmd.Args = append(cmd.Args, GOFLAGS())
		} else if GOVERSION_NUM() >= 114 {
			cmd.Args = append(cmd.Args, "-mod=mod")
		}
		cmd.Dir = filepath.Dir(GOMOD(""))
	}

	for i := len(args) - 1; i >= 0; i-- {
		a := args[i]
		if strings.HasPrefix(a, "-") {
			args = append(args[:i], args[i+1:]...)
			if GOVERSION_NUM() > 110 {
				cmd.Args = append(cmd.Args, a)
			}
		}
	}

	cmd.Args = append(cmd.Args, "-e", "-f")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = append(os.Environ(), []string{"CGO_ENABLED=0"}...)
	return cmd
}

var (
	goListCache      = make(map[string]string)
	goListCacheMutex = new(sync.Mutex)
)

func GoListOptional(args ...string) (r string) {
	goListCacheMutex.Lock()
	if _, ok := goListCache[strings.Join(args, "|")]; !ok {
		goListCache[strings.Join(args, "|")] = strings.TrimSpace(RunCmdOptional(GoList(args[:len(args)-1]...), args[len(args)-1]))
	}
	r = goListCache[strings.Join(args, "|")]
	goListCacheMutex.Unlock()
	return r
}

func QT_STATIC() bool {
	return os.Getenv("QT_STATIC") == "true"
}

func QT_NOT_CACHED() bool {
	return os.Getenv("QT_NOT_CACHED") == "true"
}

var (
	goVersionCache      string
	goVersionCacheMutex = new(sync.Mutex)
)

func GOVERSION() (r string) {
	if v, ok := os.LookupEnv("GOVERSION"); ok {
		return v
	}
	if QT_MSVC() {
		return "go1.10"
	}
	goVersionCacheMutex.Lock()
	var version string
	if goVersionCache == "" {
		version = strings.TrimSpace(RunCmd(exec.Command("go", "version"), "get go version"))
		goVersionCache = strings.Split(version, " ")[2]
	}
	if strings.Contains(version, "devel") {
		if strings.Contains(version, "+6741b7009d") { //go build that supports the msvc
			goVersionCache = "go1.10"
		} else {
			goVersionCache = "go1.13"
		}
		Log.Warnln("go dev build detected; setting GOVERSION to:", goVersionCache)
	}
	r = goVersionCache
	goVersionCacheMutex.Unlock()
	return r
}

func GOVERSION_NUM() int {
	version := strings.TrimPrefix(GOVERSION(), "go")
	vmaj, _ := strconv.Atoi(string(version[0]))
	vmin, _ := strconv.Atoi(strings.Replace(version[1:], ".", "", -1))
	if strings.Count(version, ".") == 2 {
		vmin, _ = strconv.Atoi(strings.Split(version, ".")[1])
	}
	return vmaj*1e2 + vmin
}

func QT_RESOURCES_BIG() bool {
	return os.Getenv("QT_RESOURCES_BIG") == "true"
}

func QT_GEN_TSD() bool {
	return os.Getenv("QT_GEN_TSD") == "true"
}

func QT_GEN_DART() bool {
	return os.Getenv("QT_GEN_DART") == "true"
}

func QT_GEN_HAXE() bool {
	return os.Getenv("QT_GEN_HAXE") == "true"
}

func QT_GEN_SWIFT() bool {
	return os.Getenv("QT_GEN_SWIFT") == "true"
}

func QT_GEN_GO_WRAPPER() bool {
	return os.Getenv("QT_GEN_GO") == "true"
}
