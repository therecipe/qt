package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"

	"github.com/therecipe/qt/internal/utils"
)

func ParseFlags() bool {
	var (
		debug = flag.Bool("debug", false, "print debug logs")
		//TODO: docker
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

func Docker(arg []string, buildTarget, appPath string) {

	var dockerImage string

	switch buildTarget {
	case "desktop":
		{
			switch runtime.GOOS {
			case "windows":
				{
					dockerImage = "base_windows"
					if utils.QT_MXE_ARCH() == "amd64" {
						dockerImage = "base_windows_64"
					}
				}

			case "darwin":
				{
					utils.Log.Fatalf("%v is currently not supported as a deploy target by docker", runtime.GOOS)
				}

			case "linux":
				{
					dockerImage = "base"
				}
			}
		}

	case "windows":
		{
			dockerImage = "base_windows"
			if utils.QT_MXE_ARCH() == "amd64" {
				dockerImage = "base_windows_64"
			}
		}

	case "darwin":
		{
			utils.Log.Fatalf("%v is currently not supported as a deploy target by docker", runtime.GOOS)
		}

	case "linux":
		{
			dockerImage = "base"
		}

	case "android":
		{
			dockerImage = "base_android"
		}

	default:
		{
			utils.Log.Fatalf("%v is currently not supported as a deploy target by docker", runtime.GOOS)
		}
	}

	var (
		args = []string{"run", "--rm"}
		path = make([]string, 0)
	)

	for i, gp := range strings.Split(os.Getenv("GOPATH"), string(filepath.ListSeparator)) {
		args = append(args, []string{"-v", fmt.Sprintf("%v:/media/sf_GOPATH%v", gp, i)}...)
		path = append(path, fmt.Sprintf("/media/sf_GOPATH%v", i))
	}

	args = append(args, []string{"-e", "GOPATH=" + strings.Join(path, ":")}...)

	args = append(args, []string{"-i", fmt.Sprintf("therecipe/qt:%v", dockerImage)}...)

	args = append(args, arg...)

	var found bool
	for i, gp := range strings.Split(os.Getenv("GOPATH"), string(filepath.ListSeparator)) {
		gp = filepath.Clean(gp)
		if strings.HasPrefix(appPath, gp) {
			args = append(args, strings.Replace(strings.Replace(appPath, gp, fmt.Sprintf("/media/sf_GOPATH%v", i), -1), "\\", "/", -1))
			found = true
			break
		}
	}

	if !found {
		utils.Log.Panicln("Project needs to be inside GOPATH", appPath, os.Getenv("GOPATH"))
	}

	utils.RunCmd(exec.Command("docker", args...), fmt.Sprintf("deploy binary for %v on %v with docker", buildTarget, runtime.GOOS))
}

func EnvAndTagFlags(target string) (env map[string]string, tagFlags string) {

	switch target {
	case
		"android":
		{
			tagFlags = fmt.Sprintf("-tags=\"%v\"", target)

			switch runtime.GOOS {
			case "darwin", "linux":
				{
					env = map[string]string{
						"PATH":   os.Getenv("PATH"),
						"GOPATH": utils.MustGoPath(),
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
				}

			case "windows":
				{
					env = map[string]string{
						"PATH":   os.Getenv("PATH"),
						"GOPATH": utils.MustGoPath(),
						"GOROOT": runtime.GOROOT(),

						"TMP":  os.Getenv("TMP"),
						"TEMP": os.Getenv("TEMP"),

						"GOOS":   "android",
						"GOARCH": "arm",
						"GOARM":  "7",

						"CGO_ENABLED":  "1",
						"CC":           filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-gcc"),
						"CXX":          filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-g++"),
						"CGO_CPPFLAGS": fmt.Sprintf("-isystem %v", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-16", "arch-arm", "usr", "include")),
						"CGO_LDFLAGS":  fmt.Sprintf("--sysroot=%v -llog", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-16", "arch-arm")),
					}
				}
			}
		}

	case
		"ios", "ios-simulator":
		{
			tagFlags = "-tags=\"ios\""

			var ClangDir, ClangPlatform, ClangFlag, ClangArch, GoArch = func() (string, string, string, string, string) {
				if target == "ios" {
					return "iPhoneOS", utils.IPHONEOS_SDK_DIR(), "iphoneos", "arm64", "arm64"
				}
				return "iPhoneSimulator", utils.IPHONESIMULATOR_SDK_DIR(), "ios-simulator", "x86_64", "amd64"
			}()

			env = map[string]string{
				"PATH":   os.Getenv("PATH"),
				"GOPATH": utils.MustGoPath(),
				"GOROOT": runtime.GOROOT(),

				"GOOS":   runtime.GOOS,
				"GOARCH": GoArch,

				"CGO_ENABLED":  "1",
				"CGO_CPPFLAGS": fmt.Sprintf("-isysroot %v/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v -m%v-version-min=7.0 -arch %v", utils.XCODE_DIR(), ClangDir, ClangPlatform, ClangFlag, ClangArch),
				"CGO_LDFLAGS":  fmt.Sprintf("-isysroot %v/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v -m%v-version-min=7.0 -arch %v", utils.XCODE_DIR(), ClangDir, ClangPlatform, ClangFlag, ClangArch),
			}
		}

	case
		"desktop":
		{
			if runtime.GOOS == "windows" {
				env = map[string]string{
					"PATH":   os.Getenv("PATH"),
					"GOPATH": utils.MustGoPath(),
					"GOROOT": runtime.GOROOT(),

					"TMP":  os.Getenv("TMP"),
					"TEMP": os.Getenv("TEMP"),

					"GOOS":   runtime.GOOS,
					"GOARCH": "386",

					"CGO_ENABLED": "1",
				}
				if utils.QT_MSYS2() {
					env["GOARCH"] = utils.QT_MSYS2_ARCH()
				}
			}
		}

	case "sailfish", "sailfish-emulator":
		{
			tagFlags = fmt.Sprintf("-tags=\"%v\"", strings.Replace(target, "-", "_", -1))

			env = map[string]string{
				"PATH":   os.Getenv("PATH"),
				"GOPATH": utils.MustGoPath(),
				"GOROOT": runtime.GOROOT(),

				"GOOS":   "linux",
				"GOARCH": "386",
			}

			if runtime.GOOS == "windows" {
				env["TMP"] = os.Getenv("TMP")
				env["TEMP"] = os.Getenv("TEMP")
			}

			if target == "sailfish" {
				env["GOARCH"] = "arm"
				env["GOARM"] = "7"
			}
		}

	case "rpi1", "rpi2", "rpi3":
		{
			if runtime.GOOS == "linux" {
				tagFlags = fmt.Sprintf("-tags=\"%v\"", target)

				env = map[string]string{
					"PATH":   os.Getenv("PATH"),
					"GOPATH": utils.MustGoPath(),
					"GOROOT": runtime.GOROOT(),

					"GOOS":   "linux",
					"GOARCH": "arm",
					"GOARM":  "7",

					"CGO_ENABLED": "1",
					"CC":          fmt.Sprintf("%v/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf-gcc", utils.RPI_TOOLS_DIR()),
					"CXX":         fmt.Sprintf("%v/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf-g++", utils.RPI_TOOLS_DIR()),
				}

				if target == "rpi1" {
					env["GOARM"] = "6"
				}
			} else {
				utils.Log.Panicf("failed to install %v on %v", target, runtime.GOOS)
			}
		}

	case "windows":
		{
			tagFlags = fmt.Sprintf("-tags=\"%v\"", target)

			env = map[string]string{
				"PATH":   os.Getenv("PATH"),
				"GOPATH": utils.MustGoPath(),
				"GOROOT": runtime.GOROOT(),

				"GOOS":   "windows",
				"GOARCH": utils.QT_MXE_ARCH(),

				"CGO_ENABLED": "1",
				"CC":          utils.QT_MXE_BIN("gcc"),
				"CXX":         utils.QT_MXE_BIN("g++"),
			}
		}

	default:
		{
			utils.Log.Panicf("failed to install %v on %v", target, runtime.GOOS)
		}
	}
	return
}
