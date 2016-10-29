package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/templater"
	"github.com/therecipe/qt/internal/utils"
)

func install(buildTarget string) {
	utils.Log.Infof("running setup/install %v", buildTarget)

	if strings.HasSuffix(buildTarget, "-docker") {
		utils.Log.Debugf("build target is %v -> skipping installation", buildTarget)
		return
	}

	var env, tagFlags = getEnvAndTagflags(buildTarget)

	if buildTarget != "desktop" || (runtime.GOOS == "windows" && buildTarget == "desktop") {
		if buildTarget == "sailfish" {
			var _, err = ioutil.ReadDir(filepath.Join(runtime.GOROOT(), "bin", "linux_386"))
			if err != nil {
				var build = exec.Command(filepath.Join(runtime.GOROOT(), "src", func() string {
					if runtime.GOOS == "windows" {
						return "run.bat"
					}
					return "run.bash"
				}()))
				for key, value := range env {
					build.Env = append(build.Env, fmt.Sprintf("%v=%v", key, value))
				}
				build.Run()
			}
		}

		utils.RunCmd(installPkgCmd(buildTarget, tagFlags, "std", env), "install.std")
		//also install armv7 std
		if buildTarget == "ios" {
			var cmd = installPkgCmd(buildTarget, tagFlags, "std", env)
			cmd.Env = append(strings.Split(strings.Replace(strings.Replace(strings.Join(cmd.Env, "|"), "-arch arm64", "-arch armv7", -1), "arm64", "arm", -1), "|"), "GOARM=7")
			utils.RunCmd(cmd, "install.std")
		}
	}

	if buildTarget != "desktop" {
		utils.Log.Debugf("build target is %v -> skipping installation of modules", buildTarget)
		//return
	}

	if strings.ToLower(os.Getenv("CI")) == "true" && !utils.QT_STUB() {
		utils.Log.Debug("running on CI without QT_STUB -> skipping installation of modules")
		return
	}

	utils.Log.Infof("starting to install modules (~%v5min)", func() string {
		if !utils.QT_STUB() {
			return "1"
		}
		return ""
	}())

	for _, module := range templater.GetLibs() {
		if templater.ShouldBuild(module) {

			if !(buildTarget == "android" && (module == "DBus" || module == "WebEngine" || module == "Designer" || (strings.HasSuffix(module, "Extras") && module != "AndroidExtras"))) &&
				!(strings.HasPrefix(buildTarget, "ios") && (module == "SerialPort" || module == "SerialBus" || module == "WebEngine" || module == "PrintSupport" || module == "Designer" || strings.HasSuffix(module, "Extras"))) && //TODO: support for PrintSupport
				!(strings.HasPrefix(buildTarget, "rpi") && (module == "WebEngine" || module == "Designer" || strings.HasSuffix(module, "Extras"))) { //TODO: support for WebEngine (rpi2 + rpi3)

				utils.Log.Infof("installing%v qt/%v", func() string {
					if utils.QT_STUB() {
						return " stub"
					}
					return " full"
				}(), strings.ToLower(module))

				utils.RunCmdOptional(installPkgCmd(buildTarget, tagFlags, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(module)), env), fmt.Sprintf("install.%v", strings.ToLower(module)))
				//also install armv7 modules
				if buildTarget == "ios" {
					var cmd = installPkgCmd(buildTarget, tagFlags, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(module)), env)
					cmd.Env = append(strings.Split(strings.Replace(strings.Replace(strings.Join(cmd.Env, "|"), "-arch arm64", "-arch armv7", -1), "arm64", "arm", -1), "|"), "GOARM=7")
					utils.RunCmdOptional(cmd, fmt.Sprintf("install.%v", strings.ToLower(module)))
				}
			}
		}
	}
}

func installPkgCmd(buildTarget, tagFlags, pkg string, env map[string]string) *exec.Cmd {
	var cmd = exec.Command("go", "install")
	if tagFlags != "" {
		cmd.Args = append(cmd.Args, tagFlags)
	}
	if buildTarget != "desktop" {
		cmd.Args = append(cmd.Args, fmt.Sprintf("-installsuffix=%v", strings.Replace(buildTarget, "-", "_", -1)))
	}
	cmd.Args = append(cmd.Args, pkg)
	for key, value := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
	}
	return cmd
}

func getEnvAndTagflags(buildTarget string) (env map[string]string, tagFlags string) {
	switch buildTarget {
	case "desktop", "android", "ios", "ios-simulator",
		"sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "windows", "darwin", "linux",
		"linux-docker", "windows-docker", "android-docker":
		{
			var buildDocker = strings.HasSuffix(buildTarget, "-docker")
			switch buildTarget {
			case "windows":
				{
					if runtime.GOOS == "windows" && !buildDocker {
					} else if runtime.GOOS == "linux" || buildDocker {
					} else {
						utils.Log.Fatalf("%v is currently not supported as a deploy target on %v", buildTarget, runtime.GOOS)
					}
				}

			case "darwin", "ios", "ios-simulator":
				{
					if runtime.GOOS == "darwin" && !buildDocker {
					} else {
						utils.Log.Fatalf("%v is currently not supported as a deploy target on %v (not even with docker)", buildTarget, runtime.GOOS)
					}
				}

			case "linux":
				{
					if runtime.GOOS == "linux" && !buildDocker {
					} else if buildDocker {
					} else {
						utils.Log.Fatalf("%v is currently not supported as a deploy target on %v", buildTarget, runtime.GOOS)
					}
				}
			}
		}

	default:
		{
			utils.Log.Panicf("failed to recognize build target %v", buildTarget)
		}
	}

	switch buildTarget {
	case
		"android":
		{
			tagFlags = fmt.Sprintf("-tags=\"%v\"", buildTarget)

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
						"CGO_CPPFLAGS": fmt.Sprintf("-isystem %v", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-9", "arch-arm", "usr", "include")),
						"CGO_LDFLAGS":  fmt.Sprintf("--sysroot=%v -llog", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-9", "arch-arm")),
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
						"CGO_CPPFLAGS": fmt.Sprintf("-isystem %v", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-9", "arch-arm", "usr", "include")),
						"CGO_LDFLAGS":  fmt.Sprintf("--sysroot=%v -llog", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-9", "arch-arm")),
					}
				}
			}
		}

	case
		"ios", "ios-simulator":
		{
			tagFlags = "-tags=\"ios\""

			var ClangDir, ClangPlatform, ClangFlag, ClangArch, GoArch = func() (string, string, string, string, string) {
				if buildTarget == "ios" {
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
			}
		}

	case "sailfish", "sailfish-emulator":
		{
			tagFlags = fmt.Sprintf("-tags=\"%v\"", strings.Replace(buildTarget, "-", "_", -1))

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

			if buildTarget == "sailfish" {
				env["GOARCH"] = "arm"
				env["GOARM"] = "7"
			}
		}

	case "rpi1", "rpi2", "rpi3":
		{
			if runtime.GOOS == "linux" {
				tagFlags = fmt.Sprintf("-tags=\"%v\"", buildTarget)

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

				if buildTarget == "rpi1" {
					env["GOARM"] = "6"
				}
			} else {
				utils.Log.Panicf("failed to install %v on %v", buildTarget, runtime.GOOS)
			}
		}

	case "windows":
		{
			if runtime.GOOS == "linux" {
				tagFlags = fmt.Sprintf("-tags=\"%v\"", buildTarget)

				env = map[string]string{
					"PATH":   os.Getenv("PATH"),
					"GOPATH": utils.MustGoPath(),
					"GOROOT": runtime.GOROOT(),

					"GOOS":   "windows",
					"GOARCH": "386",

					"CGO_ENABLED": "1",
					"CC":          "/usr/lib/mxe/usr/bin/i686-w64-mingw32.shared-gcc",
					"CXX":         "/usr/lib/mxe/usr/bin/i686-w64-mingw32.shared-g++",
				}
			} else {
				utils.Log.Panicf("failed to install %v on %v", buildTarget, runtime.GOOS)
			}
		}

	default:
		{
			utils.Log.Panicf("failed to install %v on %v", buildTarget, runtime.GOOS)
		}
	}
	return
}
