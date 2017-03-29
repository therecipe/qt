package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/therecipe/qt/internal/utils"
)

func ParseFlags() bool {
	var (
		debug      = flag.Bool("debug", false, "print debug logs")
		help       = flag.Bool("help", false, "print help")
		p          = flag.Int("p", runtime.NumCPU(), "specify the number of cpu's to be used")
		qt_dir     = flag.String("qt_dir", utils.QT_DIR(), "export QT_DIR")
		qt_version = flag.String("qt_version", utils.QT_VERSION(), "export QT_VERSION")
	)
	flag.Parse()

	if *debug {
		utils.Log.Level = logrus.DebugLevel
	}

	runtime.GOMAXPROCS(*p)

	if dir := *qt_dir; dir != utils.QT_DIR() {
		os.Setenv("QT_DIR", dir)
	}

	if version := *qt_version; version != utils.QT_VERSION() {
		os.Setenv("QT_VERSION", version)
	}

	return help != nil && *help
}

func Docker(arg []string, target, path string, writeCacheToHost bool) {
	arg = append(arg, target)

	var image string
	switch target {
	case "windows":
		if utils.QT_MXE_ARCH() == "386" {
			if utils.QT_MXE_STATIC() {
				image = "windows_32_static"
			} else {
				image = "windows_32_shared"
			}
		} else {
			if utils.QT_MXE_STATIC() {
				image = "windows_64_static"
			} else {
				image = "windows_64_shared"
			}
		}

	case "linux", "android", "rpi1", "rpi2", "rpi3":
		image = target

	default:
		utils.Log.Fatalf("%v is currently not supported", target)
	}

	args := []string{"run", "--rm"}
	paths := make([]string, 0)

	for i, gp := range strings.Split(utils.GOPATH(), string(filepath.ListSeparator)) {
		args = append(args, []string{"-v", fmt.Sprintf("%v:/media/sf_GOPATH%v", gp, i)}...)
		paths = append(paths, fmt.Sprintf("/media/sf_GOPATH%v", i))
	}

	gpath := strings.Join(paths, ":")
	if writeCacheToHost {
		gpath += ":/home/user/work"
		args = append(args, []string{"-e", "QT_STUB=true"}...)
	} else {
		gpath = "/home/user/work:" + gpath
	}

	args = append(args, []string{"-e", "GOPATH=" + gpath}...)

	args = append(args, []string{"-i", fmt.Sprintf("therecipe/qt:%v", image)}...)

	args = append(args, arg...)

	var found bool
	for i, gp := range strings.Split(utils.GOPATH(), string(filepath.ListSeparator)) {
		gp = filepath.Clean(gp)
		if strings.HasPrefix(path, gp) {
			args = append(args, strings.Replace(strings.Replace(path, gp, fmt.Sprintf("/media/sf_GOPATH%v", i), -1), "\\", "/", -1))
			found = true
			break
		}
	}

	if !found && path != "" {
		utils.Log.Panicln("Project needs to be inside GOPATH", path, utils.GOPATH())
	}

	utils.RunCmd(exec.Command("docker", args...), fmt.Sprintf("deploy binary for %v on %v with docker", target, runtime.GOOS))
}

func BuildEnv(target, name, depPath string) (map[string]string, []string, []string, string) {

	var tags []string
	var ldFlags []string
	var out string
	var env map[string]string

	switch target {
	case "android":
		tags = []string{target}
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, "libgo_base.so")
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "android",
			"GOARCH": "arm",
			"GOARM":  "7",

			"CGO_ENABLED":  "1",
			"CC":           filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-gcc"),
			"CXX":          filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-g++"),
			"CGO_CPPFLAGS": fmt.Sprintf("-isystem %v", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-16", "arch-arm", "usr", "include")),
			"CGO_LDFLAGS":  fmt.Sprintf("--sysroot=%v -llog", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-16", "arch-arm")),
		}
		if runtime.GOOS == "windows" {
			env["TMP"] = os.Getenv("TMP")
			env["TEMP"] = os.Getenv("TEMP")
		}

	case "ios", "ios-simulator":
		tags = []string{"ios"}
		ldFlags = []string{"-w"}
		out = filepath.Join(depPath, "libgo.a")

		GOARCH := "amd64"
		CLANGDIR := "iPhoneSimulator"
		CLANGPLAT := utils.IPHONESIMULATOR_SDK_DIR()
		CLANGFLAG := "ios-simulator"
		CLANGARCH := "x86_64"
		if target == "ios" {
			GOARCH = "arm64"
			CLANGDIR = "iPhoneOS"
			CLANGPLAT = utils.IPHONEOS_SDK_DIR()
			CLANGFLAG = "iphoneos"
			CLANGARCH = "arm64"
		}

		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   runtime.GOOS,
			"GOARCH": GOARCH,

			"CGO_ENABLED":  "1",
			"CGO_CPPFLAGS": fmt.Sprintf("-isysroot %v/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v -m%v-version-min=7.0 -arch %v", utils.XCODE_DIR(), CLANGDIR, CLANGPLAT, CLANGFLAG, CLANGARCH),
			"CGO_LDFLAGS":  fmt.Sprintf("-isysroot %v/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v -m%v-version-min=7.0 -arch %v", utils.XCODE_DIR(), CLANGDIR, CLANGPLAT, CLANGFLAG, CLANGARCH),
		}

	case "darwin":
		ldFlags = []string{"-w", fmt.Sprintf("-r=%v", filepath.Join(utils.QT_DARWIN_DIR(), "lib"))}
		out = filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", name, name))

	case "windows":
		ldFlags = []string{"-s", "-w", "-H=windowsgui"}
		if runtime.GOOS != target {
			tags = []string{"windows"}
		}
		out = filepath.Join(depPath, name)
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"TMP":  os.Getenv("TMP"),
			"TEMP": os.Getenv("TEMP"),

			"GOOS":   "windows",
			"GOARCH": "386",

			"CGO_ENABLED": "1",
		}

		if runtime.GOOS == target {
			if utils.QT_MSYS2() {
				env["GOARCH"] = utils.QT_MSYS2_ARCH()
			}
		} else {
			delete(env, "TMP")
			delete(env, "TEMP")

			env["GOARCH"] = utils.QT_MXE_ARCH()
			env["CC"] = utils.QT_MXE_BIN("gcc")
			env["CXX"] = utils.QT_MXE_BIN("g++")
		}

	case "linux":
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, name)

	case "rpi1", "rpi2", "rpi3":
		tags = []string{target}
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, name)

		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "linux",
			"GOARCH": "arm",
			"GOARM":  "7",

			"CGO_ENABLED": "1",
			"CC":          fmt.Sprintf("%v/arm-bcm2708/%v/bin/arm-linux-gnueabihf-gcc", utils.RPI_TOOLS_DIR(), utils.RPI_COMPILER()),
			"CXX":         fmt.Sprintf("%v/arm-bcm2708/%v/bin/arm-linux-gnueabihf-g++", utils.RPI_TOOLS_DIR(), utils.RPI_COMPILER()),
		}

		if target == "rpi1" {
			env["GOARM"] = "6"
		}
	}
	return env, tags, ldFlags, out
}
