package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"

	"github.com/therecipe/qt/internal/utils"
)

var buildVersion = "no build version"

func ParseFlags() bool {
	var (
		debug      = flag.Bool("debug", false, "print debug logs")
		help       = flag.Bool("help", false, "print help")
		p          = flag.Int("p", runtime.NumCPU(), "specify the number of cpu's to be used")
		qt_dir     = flag.String("qt_dir", utils.QT_DIR(), "export QT_DIR")
		qt_version = flag.String("qt_version", utils.QT_VERSION(), "export QT_VERSION")
		version    = flag.Bool("version", false, "print build version (if available)")
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

	if *version {
		println(buildVersion)
		os.Exit(0)
	}

	return help != nil && *help
}

func Docker(arg []string, target, path string, writeCacheToHost bool) {
	virtual(arg, target, path, writeCacheToHost, true, "")
}

func Vagrant(arg []string, target, path string, writeCacheToHost bool, system string) {
	virtual(arg, target, path, writeCacheToHost, false, system)
}

func virtual(arg []string, target, path string, writeCacheToHost bool, docker bool, system string) {
	arg = append(arg, target)
	if system == "" {
		system = "linux"
	}

	var image string
	switch target {
	case "windows":
		if utils.QT_MXE_ARCH() == "amd64" || utils.QT_MSYS2_ARCH() == "amd64" {
			if utils.QT_MXE_STATIC() || utils.QT_MSYS2_STATIC() {
				image = "windows_64_static"
			} else {
				image = "windows_64_shared"
			}
		} else {
			if utils.QT_MXE_STATIC() || utils.QT_MSYS2_STATIC() {
				image = "windows_32_static"
			} else {
				image = "windows_32_shared"
			}
		}
		if !docker {
			image = target
		}

	case "linux", "android", "rpi1", "rpi2", "rpi3":
		image = target

	default:
		switch system {
		case "darwin":
		case "linux":
			if strings.HasPrefix(target, "windows") && strings.Contains(target, "_") {
				image = target
			} else if docker && strings.Contains(target, "_") {
				utils.Log.Fatalf("%v is currently not supported", target)
			}
		case "windows":
		}
	}
	if image != "" {
		target = image
	}

	var args []string
	if docker {
		args = []string{"run", "--rm"}
	} else {
		if target == "pkg_config" {
			args = []string{"source /home/vagrant/.profile"}
		}
	}
	if runtime.GOOS == "linux" {
		u, err := user.Current()
		if err != nil {
			utils.Log.WithError(err).Error("failed to lookup current user")
		} else {
			args = append(args, "-e", fmt.Sprintf("IDUG=%v:%v", u.Uid, u.Gid))
		}
	}

	paths := make([]string, 0)

	for i, gp := range strings.Split(utils.GOPATH(), string(filepath.ListSeparator)) {
		if runtime.GOOS == "windows" {
			gp = "//" + strings.ToLower(gp[:1]) + gp[1:]
		}
		gp = strings.Replace(gp, "\\", "/", -1)
		gp = strings.Replace(gp, ":", "", -1)

		var pathprefix string
		if docker {
			args = append(args, []string{"-v", fmt.Sprintf("%v:/media/sf_GOPATH%v", gp, i)}...)
		} else if system == "windows" {
			pathprefix = "C:"
		}
		paths = append(paths, fmt.Sprintf(pathprefix+"/media/sf_GOPATH%v", i))
	}

	var gpfs string
	if docker {
		gpfs = "/home/user/work"
	} else {
		switch system {
		case "linux", "docker":
			gpfs = "/home/vagrant/gopath"
		case "darwin":
			gpfs = "/Users/vagrant/gopath"
		case "windows":
			gpfs = "C:/gopath"
		}
	}

	pathseperator := ":"
	if system == "windows" {
		pathseperator = ";"
	}
	gpath := strings.Join(paths, pathseperator)
	if writeCacheToHost {
		gpath += pathseperator + gpfs
		args = append(args, []string{"-e", "QT_STUB=true"}...)
	} else {
		gpath = gpfs + pathseperator + gpath
	}

	if utils.QT_DEBUG() {
		args = append(args, []string{"-e", "QT_DEBUG=true"}...)
	}

	if utils.QT_DEBUG_QML() {
		args = append(args, []string{"-e", "QT_DEBUG_QML=true"}...)
	}

	if utils.CI() {
		args = append(args, []string{"-e", "CI=true"}...)
	}

	if utils.QT_WEBKIT() {
		args = append(args, []string{"-e", "QT_WEBKIT=true"}...)
	}

	if docker {
		args = append(args, []string{"-e", "GOPATH=" + gpath}...)
	} else {
		args = append(args, []string{"-e", "QT_VAGRANT=true"}...)
		args = append(args, []string{"-e", "GOPATH='" + gpath + "'"}...)
	}

	if docker {
		args = append(args, []string{"-i", fmt.Sprintf("therecipe/qt:%v", image)}...)
	} else {
		for i, a := range args {
			if a == "-e" {
				args[i] = "&&"
				args[i+1] = "export " + args[i+1]
			}
		}
		if args[0] == "&&" {
			args = args[1:]
		}
		args = append(args, "&&")

		if system == "windows" {
			if strings.Contains(target, "_") {
				args = append(args, []string{"C:/msys64/usr/bin/bash -l -c"}...)
				arg[0] = "'C:/gopath/bin/" + arg[0]
			} else {
				args = append(args, []string{"C:/windows/system32/cmd /C"}...)
			}
		}
	}

	args = append(args, arg...)

	var found bool
	for i, gp := range strings.Split(utils.GOPATH(), string(filepath.ListSeparator)) {
		gp = filepath.Clean(gp)
		if strings.HasPrefix(path, gp) {
			if !docker && system == "windows" && strings.Contains(target, "_") {
				args = append(args, strings.Replace(strings.Replace(path, gp, fmt.Sprintf("C:/media/sf_GOPATH%v", i), -1), "\\", "/", -1))
			} else {
				args = append(args, strings.Replace(strings.Replace(path, gp, fmt.Sprintf("/media/sf_GOPATH%v", i), -1), "\\", "/", -1))
			}
			found = true
			break
		}
	}

	if !found && path != "" {
		utils.Log.Panicln("Project needs to be inside GOPATH", path, utils.GOPATH())
	}

	if docker {
		for i := range args {
			if strings.HasPrefix(args[i], "windows_") {
				args[i] = "windows"
			}
		}

		utils.RunCmd(exec.Command("docker", args...), fmt.Sprintf("deploy binary for %v on %v with docker", target, runtime.GOOS))
	} else {
		switch target {
		case "ios-simulator":
			target = "ios"
		case "android-emulator":
			target = "android" //TODO:
		}

		upCmd := exec.Command("vagrant", "up", "--no-provision", target)
		upCmd.Dir = utils.GoQtPkgPath("internal", "vagrant", system)
		utils.RunCmd(upCmd, fmt.Sprintf("vagrant up for %v/%v on %v", system, target, runtime.GOOS))

		command := strings.Join(args, " ")
		command = strings.Replace(command, " pkg_config ", " linux ", -1)
		command = strings.Replace(command, " homebrew ", " desktop ", -1)
		for _, v := range []string{"windows_32_shared", "windows_32_static", "windows_64_shared", "windows_64_static"} {
			command = strings.Replace(command, " "+v+" ", " windows ", -1)
		}
		if system == "windows" && strings.Contains(target, "_") {
			command += "'"
		}

		var cmd *exec.Cmd
		if system == "windows" {
			cmd = exec.Command("vagrant", []string{"ssh", target, "--", command}...)
		} else {
			cmd = exec.Command("vagrant", []string{"ssh", "-c", command, target}...)
		}

		cmd.Dir = utils.GoQtPkgPath("internal", "vagrant", system)
		utils.RunCmd(cmd, fmt.Sprintf("deploy binary for %v/%v on %v with vagrant", system, target, runtime.GOOS))

		haltCmd := exec.Command("vagrant", "halt", target)
		haltCmd.Dir = utils.GoQtPkgPath("internal", "vagrant", system)
		utils.RunCmd(haltCmd, fmt.Sprintf("vagrant halt for %v/%v on %v", system, target, runtime.GOOS))
	}
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

	case "android-emulator":
		tags = []string{strings.Replace(target, "-", "_", -1)}
		ldFlags = []string{"-s", "-w"}
		out = filepath.Join(depPath, "libgo_base.so")
		env = map[string]string{
			"PATH":   os.Getenv("PATH"),
			"GOPATH": utils.GOPATH(),
			"GOROOT": runtime.GOROOT(),

			"GOOS":   "android",
			"GOARCH": "386",

			"CGO_ENABLED":  "1",
			"CC":           filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "x86-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "i686-linux-android-gcc"),
			"CXX":          filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "x86-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "i686-linux-android-g++"),
			"CGO_CPPFLAGS": fmt.Sprintf("-isystem %v", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-16", "arch-x86", "usr", "include")),
			"CGO_LDFLAGS":  fmt.Sprintf("--sysroot=%v -llog", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-16", "arch-x86")),
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
