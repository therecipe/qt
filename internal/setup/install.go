package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/therecipe/qt/internal/binding/templater"
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

	if buildTarget != "desktop" || runtime.GOOS == "windows" {

		switch buildTarget {
		case
			"android":
			{
				switch runtime.GOOS {
				case "darwin", "linux":
					{
						env = map[string]string{
							"PATH":   os.Getenv("PATH"),
							"GOPATH": os.Getenv("GOPATH"),
							"GOROOT": runtime.GOROOT(),

							"GOOS":   "android",
							"GOARCH": "arm",
							"GOARM":  "7",

							"CGO_ENABLED":  "1",
							"CC":           filepath.Join("/opt", "android-ndk", "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-gcc"),
							"CXX":          filepath.Join("/opt", "android-ndk", "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-g++"),
							"CGO_CPPFLAGS": "-isystem /opt/android-ndk/platforms/android-9/arch-arm/usr/include",
							"CGO_LDFLAGS":  "--sysroot=/opt/android-ndk/platforms/android-9/arch-arm/ -llog",
						}
					}

				case "windows":
					{
						env = map[string]string{
							"PATH":   os.Getenv("PATH"),
							"GOPATH": os.Getenv("GOPATH"),
							"GOROOT": runtime.GOROOT(),

							"GOOS":   "android",
							"GOARCH": "arm",
							"GOARM":  "7",

							"CGO_ENABLED":  "1",
							"CC":           "C:\\android\\android-ndk\\toolchains\\arm-linux-androideabi-4.9\\prebuilt\\windows\\bin\\arm-linux-androideabi-gcc.exe",
							"CXX":          "C:\\android\\android-ndk\\toolchains\\arm-linux-androideabi-4.9\\prebuilt\\windows\\bin\\arm-linux-androideabi-g++.exe",
							"CGO_CPPFLAGS": "-isystem C:\\android\\android-ndk\\platforms\\android-9\\arch-arm\\usr\\include",
							"CGO_LDFLAGS":  "--sysroot=C:\\android\\android-ndk\\platforms\\android-9\\arch-arm\\ -llog",
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
						return "386"
					}()

					CLANGARCH, CLANGDIR, CLANGFLAG = func() (string, string, string) {
						if buildTarget == "ios" {
							return "arm64", "iPhoneOS", "iphoneos"
						}
						return "i386", "iPhoneSimulator", "ios-simulator"
					}()
				)

				env = map[string]string{
					"PATH":   os.Getenv("PATH"),
					"GOPATH": os.Getenv("GOPATH"),
					"GOROOT": runtime.GOROOT(),

					"GOOS":   runtime.GOOS,
					"GOARCH": GOARCH,

					"CGO_ENABLED":  "1",
					"CGO_CPPFLAGS": fmt.Sprintf("-isysroot /Applications/Xcode.app/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v9.3.sdk -m%v-version-min=6.1 -arch %v", CLANGDIR, CLANGDIR, CLANGFLAG, CLANGARCH),
					"CGO_LDFLAGS":  fmt.Sprintf("-isysroot /Applications/Xcode.app/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v9.3.sdk -m%v-version-min=6.1 -arch %v", CLANGDIR, CLANGDIR, CLANGFLAG, CLANGARCH),
				}
			}

		case
			"desktop":
			{
				if runtime.GOOS == "windows" {
					env = map[string]string{
						"PATH":   os.Getenv("PATH"),
						"GOPATH": os.Getenv("GOPATH"),
						"GOROOT": runtime.GOROOT(),

						"GOOS":   runtime.GOOS,
						"GOARCH": "386",

						"CGO_ENABLED": "1",
					}
				}
			}
		}

		var cmd = exec.Command("go", "install")
		if tagFlags != "" {
			cmd.Args = append(cmd.Args, tagFlags)
		}
		cmd.Args = append(cmd.Args, "std")

		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}
		runCmd(cmd, "install.std_1")

		//armv7
		/*
			if buildTarget == "ios" {
				var cmdiOS = exec.Command("go", "install")
				if tagFlags != "" {
					cmdiOS.Args = append(cmdiOS.Args, tagFlags)
				}
				cmdiOS.Args = append(cmdiOS.Args, "std")
				var tmp = strings.Replace(strings.Join(cmd.Env, "|"), "-arch arm64", "-arch armv7", -1)
				tmp = strings.Replace(tmp, "arm64", "arm", -1)
				cmdiOS.Env = append(strings.Split(tmp, "|"), "GOARM=7")
				runCmd(cmdiOS, "install.std_2")
			}
		*/
	}

	fmt.Println("------------------------install-------------------------")

	for _, m := range templater.GetLibs() {

		if !(buildTarget == "android" && (m == "DBus" || m == "MacExtras" || m == "WebEngine" || m == "Designer")) &&
			!((buildTarget == "ios" || buildTarget == "ios-simulator") && (m == "SerialPort" || m == "MacExtras" || m == "WebEngine" || m == "PrintSupport" || m == "Designer")) { //TODO: support for PrintSupport

			var before = time.Now()

			fmt.Print(strings.ToLower(m))

			if templater.ShouldBuild(m) {
				var cmd = exec.Command("go", "install")
				if tagFlags != "" {
					cmd.Args = append(cmd.Args, tagFlags)
				}
				cmd.Args = append(cmd.Args, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(m)))

				if buildTarget != "desktop" || runtime.GOOS == "windows" {
					for key, value := range env {
						cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
					}
				}
				runCmd(cmd, fmt.Sprintf("install.%v_1", m))

				//armv7
				/*
					if buildTarget == "ios" {
						var cmdiOS = exec.Command("go", "install")
						if tagFlags != "" {
							cmdiOS.Args = append(cmdiOS.Args, tagFlags)
						}
						cmdiOS.Args = append(cmdiOS.Args, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(m)))
						var tmp = strings.Replace(strings.Join(cmd.Env, "|"), "-arch arm64", "-arch armv7", -1)
						tmp = strings.Replace(tmp, "arm64", "arm", -1)
						cmdiOS.Env = append(strings.Split(tmp, "|"), "GOARM=7")
						runCmd(cmdiOS, fmt.Sprintf("install.%v_2", m))
					}
				*/
			}

			fmt.Println(strings.Repeat(" ", 30-len(m)), time.Since(before)/time.Second*time.Second)
		}
	}
}

func runCmd(cmd *exec.Cmd, name string) {
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", name, out, err)
		os.Exit(1)
	}
}
