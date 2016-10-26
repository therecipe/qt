package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/therecipe/qt/internal/binding/templater"
	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var (
		buildTarget = "desktop"
		tagFlags    string
		env         map[string]string
	)
	if len(os.Args) > 1 {
		buildTarget = os.Args[1]
	}
	if strings.Contains(buildTarget, "docker") {
		return
	}

	if buildTarget != "desktop" || (runtime.GOOS == "windows" && buildTarget == "desktop") {

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

				var (
					GOARCH = func() string {
						if buildTarget == "ios" {
							return "arm64"
						}
						return "amd64"
					}()

					ClangDir, ClangPlatform, ClangFlag, ClangArch = func() (string, string, string, string) {
						if buildTarget == "ios" {
							return "iPhoneOS", utils.IPHONEOS_SDK_DIR(), "iphoneos", "arm64"
						}
						return "iPhoneSimulator", utils.IPHONESIMULATOR_SDK_DIR(), "ios-simulator", "x86_64"
					}()
				)

				env = map[string]string{
					"PATH":   os.Getenv("PATH"),
					"GOPATH": utils.MustGoPath(),
					"GOROOT": runtime.GOROOT(),

					"GOOS":   runtime.GOOS,
					"GOARCH": GOARCH,

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

				//build linux,386 tools
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

				if buildTarget == "sailfish" {
					env["GOARCH"] = "arm"
					env["GOARM"] = "7"
				}
			}

		case "rpi1", "rpi2", "rpi3":
			{
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
			}

		case "windows":
			{
				tagFlags = fmt.Sprintf("-tags=\"%v\"", buildTarget)

				if runtime.GOOS == "linux" {
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
				}
			}
		}

		var cmd = exec.Command("go", "install")
		if tagFlags != "" {
			cmd.Args = append(cmd.Args, tagFlags)
		}
		if buildTarget != "desktop" {
			cmd.Args = append(cmd.Args, fmt.Sprintf("-installsuffix=%v", strings.Replace(buildTarget, "-", "_", -1)))
		}
		cmd.Args = append(cmd.Args, "std")
		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}
		utils.RunCmd(cmd, "install.std_1")

		//armv7
		if buildTarget == "ios" && (strings.HasPrefix(runtime.Version(), "go1.7") || strings.HasPrefix(runtime.Version(), "devel")) {
			var cmdiOS = exec.Command("go", "install")
			if tagFlags != "" {
				cmdiOS.Args = append(cmdiOS.Args, tagFlags)
			}
			if buildTarget != "desktop" {
				cmdiOS.Args = append(cmdiOS.Args, fmt.Sprintf("-installsuffix=%v", buildTarget))
			}
			cmdiOS.Args = append(cmdiOS.Args, "std")
			var tmp = strings.Replace(strings.Join(cmd.Env, "|"), "-arch arm64", "-arch armv7", -1)
			tmp = strings.Replace(tmp, "arm64", "arm", -1)
			cmdiOS.Env = append(strings.Split(tmp, "|"), "GOARM=7")
			utils.RunCmd(cmdiOS, "install.std_2")
		}
	}

	if buildTarget != "desktop" || strings.ToLower(os.Getenv("CI")) == "true" {
		return
	}

	fmt.Println("------------------------install-------------------------")

	for _, m := range templater.GetLibs() {

		if !(buildTarget == "android" && (m == "DBus" || m == "WebEngine" || m == "Designer" || (strings.HasSuffix(m, "Extras") && m != "AndroidExtras"))) &&
			!(strings.HasPrefix(buildTarget, "ios") && (m == "SerialPort" || m == "SerialBus" || m == "WebEngine" || m == "PrintSupport" || m == "Designer" || strings.HasSuffix(m, "Extras"))) && //TODO: support for PrintSupport
			!(strings.HasPrefix(buildTarget, "rpi") && (m == "WebEngine" || m == "Designer" || strings.HasSuffix(m, "Extras"))) { //TODO: support for WebEngine (rpi2 + rpi3)

			var before = time.Now()

			fmt.Print(strings.ToLower(m))

			if templater.ShouldBuild(m) {
				var cmd = exec.Command("go", "install")
				if tagFlags != "" {
					cmd.Args = append(cmd.Args, tagFlags)
				}
				if buildTarget != "desktop" {
					cmd.Args = append(cmd.Args, fmt.Sprintf("-installsuffix=%v", strings.Replace(buildTarget, "-", "_", -1)))
				}
				cmd.Args = append(cmd.Args, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(m)))

				if buildTarget != "desktop" || (runtime.GOOS == "windows" && buildTarget == "desktop") {
					for key, value := range env {
						cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
					}
				}
				utils.RunCmdOptional(cmd, fmt.Sprintf("install.%v_1", m))

				//armv7
				if buildTarget == "ios" && (strings.HasPrefix(runtime.Version(), "go1.7") || strings.HasPrefix(runtime.Version(), "devel")) {
					var cmdiOS = exec.Command("go", "install")
					if tagFlags != "" {
						cmdiOS.Args = append(cmdiOS.Args, tagFlags)
					}
					if buildTarget != "desktop" {
						cmdiOS.Args = append(cmdiOS.Args, fmt.Sprintf("-installsuffix=%v", buildTarget))
					}
					cmdiOS.Args = append(cmdiOS.Args, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(m)))
					var tmp = strings.Replace(strings.Join(cmd.Env, "|"), "-arch arm64", "-arch armv7", -1)
					tmp = strings.Replace(tmp, "arm64", "arm", -1)
					cmdiOS.Env = append(strings.Split(tmp, "|"), "GOARM=7")
					utils.RunCmdOptional(cmdiOS, fmt.Sprintf("install.%v_2", m))
				}
			}

			fmt.Println(strings.Repeat(" ", 45-len(m)), time.Since(before)/time.Second*time.Second)
		}
	}
}
