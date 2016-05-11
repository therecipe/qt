package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/therecipe/qt/internal/binding/templater"
)

func main() {
	var (
		buildTarget = "desktop"
		env         map[string]string
	)

	if len(os.Args) > 1 {
		buildTarget = os.Args[1]
	}

	if buildTarget == "android" {
		switch runtime.GOOS {
		case "darwin", "linux":
			{
				env = map[string]string{
					"GOPATH":       os.Getenv("GOPATH"),
					"GOOS":         "android",
					"GOARCH":       "arm",
					"GOARM":        "7",
					"CC":           filepath.Join("/opt", "android-ndk", "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-gcc"),
					"CXX":          filepath.Join("/opt", "android-ndk", "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-g++"),
					"CGO_ENABLED":  "1",
					"CGO_CPPFLAGS": "-isystem /opt/android-ndk/platforms/android-9/arch-arm/usr/include",
					"CGO_LDFLAGS":  "--sysroot=/opt/android-ndk/platforms/android-9/arch-arm/ -llog",
				}
			}

		case "windows":
			{
				env = map[string]string{
					"GOPATH":       os.Getenv("GOPATH"),
					"GOOS":         "android",
					"GOARCH":       "arm",
					"GOARM":        "7",
					"CC":           "C:\\android\\android-ndk\\toolchains\\arm-linux-androideabi-4.9\\prebuilt\\windows\\bin\\arm-linux-androideabi-gcc.exe",
					"CXX":          "C:\\android\\android-ndk\\toolchains\\arm-linux-androideabi-4.9\\prebuilt\\windows\\bin\\arm-linux-androideabi-g++.exe",
					"CGO_ENABLED":  "1",
					"CGO_CPPFLAGS": "-isystem C:\\android\\android-ndk\\platforms\\android-9\\arch-arm\\usr\\include",
					"CGO_LDFLAGS":  "--sysroot=C:\\android\\android-ndk\\platforms\\android-9\\arch-arm\\ -llog",
				}
			}
		}
	} else {
		switch runtime.GOOS {
		case "windows":
			{
				env = map[string]string{
					"PATH":        os.Getenv("PATH"),
					"GOPATH":      os.Getenv("GOPATH"),
					"GOOS":        runtime.GOOS,
					"GOARCH":      "386",
					"CGO_ENABLED": "1",
				}
			}
		}
	}

	var cmd = exec.Command("go")
	cmd.Args = append(cmd.Args,
		"install",
		"-p", strconv.Itoa(runtime.NumCPU()),
		"std")

	if buildTarget != "desktop" || runtime.GOOS == "windows" {
		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}
	}
	runCmd(cmd, "install.std")

	if runtime.GOOS == "windows" {
		setLargeAddressAware(filepath.Join("C:\\", "Qt", "Qt5.6.0", "Tools", "mingw492_32", "libexec", "gcc", "i686-w64-mingw32", "4.9.2", "cc1plus.exe"))
		if runtime.GOARCH == "386" {
			setLargeAddressAware(filepath.Join(runtime.GOROOT(), "pkg", "tool", "windows_386", "compile.exe"))
		}
	}

	fmt.Printf("________________________Install-%v________________________\n", buildTarget)

	for _, m := range templater.GetLibs() {

		if !(buildTarget == "android" && (m == "DBus" || m == "SerialPort" || m == "MacExtras" || m == "WebEngine" || m == "WebKit" || m == "Designer")) {

			var before = time.Now()

			fmt.Print(strings.ToLower(m))

			if templater.ShouldBuild(m) {
				var cmd = exec.Command("go")
				cmd.Args = append(cmd.Args,
					"install",
					"-p", strconv.Itoa(runtime.NumCPU()),
					fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(m)))

				if buildTarget != "desktop" || runtime.GOOS == "windows" {
					for key, value := range env {
						cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
					}
				}

				runCmd(cmd, fmt.Sprintf("install.%v", m))
			}

			fmt.Println(strings.Repeat(" ", 30-len(m)), time.Since(before)/time.Second*time.Second)
		}
	}
}

func runCmd(cmd *exec.Cmd, n string) string {
	var out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", n, out, err)
		os.Exit(1)
	}
	return string(out)
}

func setLargeAddressAware(fp string) {

	var (
		b, err = ioutil.ReadFile(fp)
		br     = bytes.NewReader(b)
	)

	if err != nil {
		println("setLargeAddressAware error: ReadFile")
	} else {
		//check if mz header is valid
		var bb = make([]byte, 2)
		br.Read(bb)
		if !bytes.Equal(bb, []byte{'M', 'Z'}) {
			println("setLargeAddressAware error: MZ header is invalid")
			return
		}

		//seek pe header offset
		br.Seek(60, 0)

		bb = make([]byte, 1)
		br.Read(bb)
		var peHeaderLoc = int64(bb[0])

		//seek pe header location
		br.Seek(peHeaderLoc, 0)

		//check if pe header is valid
		bb = make([]byte, 2)
		br.Read(bb)
		if !bytes.Equal(bb, []byte{'P', 'E'}) {
			println("setLargeAddressAware error: PE header is invalid")
			return
		}

		//seek characteristics
		br.Seek(peHeaderLoc+4+18, 0)

		//check if large address aware is set
		bb = make([]byte, 2)
		br.Read(bb)
		if !bytes.Equal(bb, []byte{47, 3}) {
			ioutil.WriteFile(strings.Replace(fp, ".exe", "_backup.exe", -1), b, 0777)
			ioutil.WriteFile(fp, bytes.Replace(b, bb, []byte{47, 3}, 1), 0777)
		}
	}
}
