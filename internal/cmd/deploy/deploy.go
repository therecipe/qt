package deploy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"
	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/minimal"
	"github.com/therecipe/qt/internal/cmd/moc"
	"github.com/therecipe/qt/internal/cmd/rcc"
	"github.com/therecipe/qt/internal/utils"
)

var (
	appPath, appName       string
	depPath                string
	buildMode, buildTarget string
	ending                 string
	buildDocker            bool
	gLdFlags               string
)

type State struct {
	AppPath                string
	BuildMode, BuildTarget string
	BuildDocker            bool
	LdFlags                string
}

func Deploy(s *State) {
	buildMode = s.BuildMode
	buildTarget = s.BuildTarget
	appPath = filepath.Clean(s.AppPath)
	buildDocker = s.BuildDocker
	gLdFlags = s.LdFlags

	if !filepath.IsAbs(appPath) {
		appPath, _ = utils.Abs(appPath)
	}
	appName = filepath.Base(appPath)

	switch buildTarget {
	case "android", "ios", "ios-simulator", "sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "windows", "darwin", "linux":
		{
			if buildTarget == runtime.GOOS && !buildDocker {
				depPath = filepath.Join(appPath, "deploy", runtime.GOOS)
			}
			depPath = filepath.Join(appPath, "deploy", buildTarget)
		}

	case "desktop":
		{
			depPath = filepath.Join(appPath, "deploy", runtime.GOOS)
		}
	}

	switch buildMode {
	case "build", "test":
		{
			utils.RemoveAll(depPath)
			utils.MkdirAll(depPath)
		}
	}

	if runtime.GOOS == "windows" && buildTarget == "desktop" || buildTarget == "windows" {
		ending = ".exe"
	}

	switch buildMode {
	case "build", "test":
		{
			if buildDocker {
				if gLdFlags != "" {
					cmd.Docker([]string{"qtdeploy", "-debug", fmt.Sprintf("-ldflags=%v", gLdFlags), "build", buildTarget}, buildTarget, appPath)
				} else {
					cmd.Docker([]string{"qtdeploy", "-debug", "build", buildTarget}, buildTarget, appPath)
				}
			} else {

				//rcc
				utils.Log.Debug("qtrcc - start")
				var qtrcc_cwd string = appPath
				if env_cwd := os.Getenv("QTRCC_CWD"); env_cwd != "" {
					qtrcc_cwd = env_cwd
				}
				var qtrcc_output *string
				if env_output_dir := os.Getenv("QTRCC_OUTPUT_DIR"); env_output_dir != "" {
					qtrcc_output = &env_output_dir
				}
				rcc.Rcc(qtrcc_cwd, buildTarget, qtrcc_output)
				utils.Log.Debug("qtrcc - done")

				//moc
				utils.Log.Debug("qtmoc - start")
				moc.MocTree(appPath, buildTarget)

				utils.Log.Debug("qtmoc - done")

				//minimal
				utils.Log.Debug("qtminimal - start")
				minimal.Minimal(appPath, buildTarget)

				parser.State.Minimal = false

				for _, c := range parser.State.ClassMap {
					c.Export = false
					for _, f := range c.Functions {
						f.Export = false
					}
				}
				utils.Log.Debug("qtminimal - done")

				utils.Log.Debug("build - start")
				build()
				utils.Log.Debug("build - done")

				utils.Log.Debug("deploy - start")
				predeploy()
				deployInternal()
				pastdeploy()
				utils.Log.Debug("deploy - done")

				cleanup()
			}
		}
	}

	switch buildMode {
	case "run", "test":
		{
			run()
		}
	}
}

func build() {

	var (
		ldFlags    []string
		tagFlags   = []string{"minimal"}
		outputFile string
		env        map[string]string
	)

	if gLdFlags != "" {
		ldFlags = strings.Split(gLdFlags, " ")
	}

	if utils.QT_DOCKER() {
		tagFlags = append(tagFlags, "docker")
	}

	switch buildTarget {
	case "android":
		{
			ldFlags = append(ldFlags, []string{"-s", "-w"}...)
			tagFlags = append(tagFlags, buildTarget)
			outputFile = filepath.Join(depPath, "libgo_base.so")

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
						"GOPATH": os.Getenv("GOPATH"),
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

			utils.Save(filepath.Join(appPath, "cgo_main_wrapper.go"), "package main\nimport \"C\"\n//export go_main_wrapper\nfunc go_main_wrapper() { main() }")
		}

	case "ios", "ios-simulator":
		{
			ldFlags = append(ldFlags, "-w")
			tagFlags = append(tagFlags, "ios")

			outputFile = filepath.Join(depPath, "libgo.a")

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
				"GOPATH": os.Getenv("GOPATH"),
				"GOROOT": runtime.GOROOT(),

				"GOOS":   runtime.GOOS,
				"GOARCH": GOARCH,

				"CGO_ENABLED":  "1",
				"CGO_CPPFLAGS": fmt.Sprintf("-isysroot %v/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v -m%v-version-min=7.0 -arch %v", utils.XCODE_DIR(), ClangDir, ClangPlatform, ClangFlag, ClangArch),
				"CGO_LDFLAGS":  fmt.Sprintf("-isysroot %v/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v -m%v-version-min=7.0 -arch %v", utils.XCODE_DIR(), ClangDir, ClangPlatform, ClangFlag, ClangArch),
			}

			utils.Save(filepath.Join(appPath, "cgo_main_wrapper.go"), "package main\nimport \"C\"\n//export go_main_wrapper\nfunc go_main_wrapper() { main() }")
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					ldFlags = append(ldFlags, []string{"-w", fmt.Sprintf("-r=%v", filepath.Join(utils.QT_DARWIN_DIR(), "lib"))}...)
					outputFile = filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName))
				}

			case "linux":
				{
					ldFlags = append(ldFlags, []string{"-s", "-w"}...)
					outputFile = filepath.Join(depPath, appName)
				}

			case "windows":
				{
					ldFlags = append(ldFlags, []string{"-s", "-w", "-H=windowsgui"}...)
					outputFile = filepath.Join(depPath, appName)
					env = map[string]string{
						"PATH":   os.Getenv("PATH"),
						"GOPATH": os.Getenv("GOPATH"),
						"GOROOT": runtime.GOROOT(),

						"TMP":  os.Getenv("TMP"),
						"TEMP": os.Getenv("TEMP"),

						"GOOS":   runtime.GOOS,
						"GOARCH": "386",

						"CGO_ENABLED": "1",
					}

					if utils.UseMsys2() {
						env["GOARCH"] = utils.QT_MSYS2_ARCH()
					}
				}
			}
		}

	case "sailfish", "sailfish-emulator":
		{
			//TODO: ldflags support

			if !strings.Contains(appPath, utils.MustGoPath()) {
				utils.Log.Panicln("Project needs to be inside GOPATH", appPath, utils.MustGoPath())
			}

			utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "registervm", filepath.Join(utils.SAILFISH_DIR(), "mersdk", "MerSDK", "MerSDK.vbox")), fmt.Sprintf("register mersdk for %v on %v", buildTarget, runtime.GOOS))
			utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "sharedfolder", "add", "MerSDK", "--name", "GOROOT", "--hostpath", runtime.GOROOT(), "--automount"), fmt.Sprintf("share GOROOT dir for %v on %v", buildTarget, runtime.GOOS))
			utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "sharedfolder", "add", "MerSDK", "--name", "GOPATH", "--hostpath", utils.MustGoPath(), "--automount"), fmt.Sprintf("share GOPATH dir for %v on %v", buildTarget, runtime.GOOS))

			if runtime.GOOS == "windows" {
				utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "startvm", "--type", "headless", "MerSDK"), fmt.Sprintf("start vbox mersdk for %v on %v", buildTarget, runtime.GOOS))
			} else {
				utils.RunCmdOptional(exec.Command("nohup", filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "startvm", "--type", "headless", "MerSDK"), fmt.Sprintf("start vbox mersdk for %v on %v", buildTarget, runtime.GOOS))
			}

			time.Sleep(10 * time.Second)

			if buildTarget == "sailfish-emulator" {

				sshCommand("2222", "root", "ln", "-s", "/opt/cross/bin/i486-meego-linux-gnu-as", "/opt/cross/libexec/gcc/i486-meego-linux-gnu/4.8.3/as")
				sshCommand("2222", "root", "ln", "-s", "/opt/cross/bin/i486-meego-linux-gnu-ld", "/opt/cross/libexec/gcc/i486-meego-linux-gnu/4.8.3/ld")

				var errBuild = sshCommand("2222", "root", "cd", strings.Replace(strings.Replace(appPath, utils.MustGoPath(), "/media/sf_GOPATH", -1), "\\", "/", -1), "&&", "GOROOT=/media/sf_GOROOT", "GOPATH=/media/sf_GOPATH", "PATH=$PATH:$GOROOT/bin/linux_386", "GOOS=linux", "GOARCH=386", "CGO_ENABLED=1", "CC=/opt/cross/bin/i486-meego-linux-gnu-gcc", "CXX=/opt/cross/bin/i486-meego-linux-gnu-g++", "CPATH=/srv/mer/targets/SailfishOS-i486/usr/include", "LIBRARY_PATH=/srv/mer/targets/SailfishOS-i486/usr/lib:/srv/mer/targets/SailfishOS-i486/lib", "CGO_LDFLAGS=--sysroot=/srv/mer/targets/SailfishOS-i486/", "go", "build", "-ldflags=\"-s -w\"", "-tags=\"minimal sailfish_emulator\"", fmt.Sprintf("-installsuffix=%v", strings.Replace(buildTarget, "-", "_", -1)), "-o", "deploy/"+buildTarget+"/harbour-"+appName)
				if errBuild != nil {
					cleanup()
					utils.Log.WithError(errBuild).Panicf("failed to build for %v on %v", buildTarget, runtime.GOOS)
				}
			} else {

				sshCommand("2222", "root", "ln", "-s", "/opt/cross/bin/armv7hl-meego-linux-gnueabi-as", "/opt/cross/libexec/gcc/armv7hl-meego-linux-gnueabi/4.8.3/as")
				sshCommand("2222", "root", "ln", "-s", "/opt/cross/bin/armv7hl-meego-linux-gnueabi-ld", "/opt/cross/libexec/gcc/armv7hl-meego-linux-gnueabi/4.8.3/ld")

				var errBuild = sshCommand("2222", "root", "cd", strings.Replace(strings.Replace(appPath, utils.MustGoPath(), "/media/sf_GOPATH", -1), "\\", "/", -1), "&&", "GOROOT=/media/sf_GOROOT", "GOPATH=/media/sf_GOPATH", "PATH=$PATH:$GOROOT/bin/linux_386", "GOOS=linux", "GOARCH=arm", "GOARM=7", "CGO_ENABLED=1", "CC=/opt/cross/bin/armv7hl-meego-linux-gnueabi-gcc", "CXX=/opt/cross/bin/armv7hl-meego-linux-gnueabi-g++", "CPATH=/srv/mer/targets/SailfishOS-armv7hl/usr/include", "LIBRARY_PATH=/srv/mer/targets/SailfishOS-armv7hl/usr/lib:/srv/mer/targets/SailfishOS-armv7hl/lib", "CGO_LDFLAGS=--sysroot=/srv/mer/targets/SailfishOS-armv7hl/", "go", "build", "-ldflags=\"-s -w\"", "-tags=\"minimal sailfish\"", fmt.Sprintf("-installsuffix=%v", buildTarget), "-o", "deploy/"+buildTarget+"/harbour-"+appName)
				if errBuild != nil {
					cleanup()
					utils.Log.WithError(errBuild).Panicf("failed to build for %v on %v", buildTarget, runtime.GOOS)
				}
			}

			return
		}

	case "rpi1", "rpi2", "rpi3":
		{
			ldFlags = append(ldFlags, []string{"-s", "-w"}...)
			tagFlags = append(tagFlags, buildTarget)
			outputFile = filepath.Join(depPath, appName)

			env = map[string]string{
				"PATH":   os.Getenv("PATH"),
				"GOPATH": os.Getenv("GOPATH"),
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
			ldFlags = append(ldFlags, []string{"-s", "-w", "-H=windowsgui"}...)
			tagFlags = append(tagFlags, "windows")
			outputFile = filepath.Join(depPath, appName)

			env = map[string]string{
				"PATH":   os.Getenv("PATH"),
				"GOPATH": os.Getenv("GOPATH"),
				"GOROOT": runtime.GOROOT(),

				"GOOS":   "windows",
				"GOARCH": utils.QT_MXE_ARCH(),

				"CGO_ENABLED": "1",
				"CC":          utils.QT_MXE_BIN("gcc"),
				"CXX":         utils.QT_MXE_BIN("g++"),
			}
		}
	}

	var cmd = exec.Command("go", "build", "-p", strconv.Itoa(runtime.GOMAXPROCS(0)), "-v", fmt.Sprintf("-ldflags=\"%v\"", strings.Join(ldFlags, "\" \"")), "-o", outputFile+ending)
	cmd.Dir = appPath
	cmd.Args = append(cmd.Args, fmt.Sprintf("-tags=\"%v\"", strings.Join(tagFlags, "\" \"")))
	if buildTarget != "desktop" {
		cmd.Args = append(cmd.Args, []string{"-i", "-installsuffix", strings.Replace(buildTarget, "-", "_", -1)}...)
		cmd.Args = append(cmd.Args, []string{"-pkgdir", filepath.Join(utils.MustGoPath(), "pkg", fmt.Sprintf("%v_%v_%v", env["GOOS"], env["GOARCH"], strings.Replace(buildTarget, "-", "_", -1)))}...)
	}

	if buildTarget != "desktop" || (runtime.GOOS == "windows" && buildTarget == "desktop") {
		if buildTarget == "android" {
			cmd.Args = append(cmd.Args, "-buildmode", "c-shared")
		}
		if buildTarget == "ios" || buildTarget == "ios-simulator" {
			cmd.Args = append(cmd.Args, "-buildmode", "c-archive")
		}
		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}
	}
	utils.RunCmd(cmd, fmt.Sprintf("build for %v on %v", buildTarget, runtime.GOOS))

	if runtime.GOOS == "darwin" && buildTarget == "desktop" {
		var strip = exec.Command("strip", outputFile)
		strip.Dir = appPath
		utils.RunCmd(strip, fmt.Sprintf("strip binary for %v on %v", buildTarget, runtime.GOOS))
	}

	//armv7
	if buildTarget == "ios" {
		var cmdiOS = exec.Command("go", "build", "-p", strconv.Itoa(runtime.GOMAXPROCS(0)), "-v", fmt.Sprintf("-ldflags=\"%v\"", strings.Join(ldFlags, "\" \"")), "-o", strings.Replace(outputFile, "libgo.a", "libgo_armv7.a", -1))
		cmdiOS.Dir = appPath
		cmdiOS.Args = append(cmdiOS.Args, fmt.Sprintf("-tags=\"%v\"", strings.Join(tagFlags, "\" \"")))
		if buildTarget != "desktop" {
			cmdiOS.Args = append(cmdiOS.Args, []string{"-i", "-installsuffix", strings.Replace(buildTarget, "-", "_", -1)}...)
			cmdiOS.Args = append(cmdiOS.Args, []string{"-pkgdir", filepath.Join(utils.MustGoPath(), "pkg", fmt.Sprintf("%v_%v_%v", env["GOOS"], env["GOARCH"], strings.Replace(buildTarget, "-", "_", -1)))}...)
		}
		cmdiOS.Args = append(cmdiOS.Args, "-buildmode", "c-archive")
		var tmp = strings.Replace(strings.Join(cmd.Env, "|"), "-arch arm64", "-arch armv7", -1)
		tmp = strings.Replace(tmp, "arm64", "arm", -1)
		cmdiOS.Env = append(strings.Split(tmp, "|"), "GOARM=7")
		utils.RunCmd(cmdiOS, fmt.Sprintf("build for %v on %v", buildTarget, runtime.GOOS))
	}
}

func predeploy() {

	var copyCmd = func() string {
		if runtime.GOOS == "windows" {
			return "xcopy"
		}
		return "cp"
	}()

	switch buildTarget {
	case "android":
		{
			utils.MkdirAll(filepath.Join(appPath, "android"))

			var libPath = filepath.Join(depPath, "build", "libs", "armeabi-v7a")
			utils.MkdirAll(libPath)

			var compiler = filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-g++")

			//add c_main_wrappers
			utils.Save(filepath.Join(depPath, "c_main_wrapper.cpp"), "#include \"libgo_base.h\"\nint main(int argc, char *argv[]) { go_main_wrapper(); }")

			var cmd = exec.Command(compiler, "c_main_wrapper.cpp", "-o", filepath.Join(depPath, "libgo.so"), "-I../..", "-L.", "-lgo_base", fmt.Sprintf("--sysroot=%v", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-9", "arch-arm")), "-shared")
			cmd.Dir = depPath
			utils.RunCmd(cmd, fmt.Sprintf("compile for %v on %v", buildTarget, runtime.GOOS))

			var strip = exec.Command(filepath.Join(filepath.Dir(compiler), "arm-linux-androideabi-strip"), "libgo.so")
			strip.Dir = depPath
			utils.RunCmd(strip, fmt.Sprintf("strip binary for %v on %v", buildTarget, runtime.GOOS))

			utils.RunCmd(exec.Command(copyCmd, filepath.Join(depPath, "libgo_base.so"), libPath), fmt.Sprintf("copy libgo_base for %v on %v", buildTarget, runtime.GOOS))
			utils.RunCmd(exec.Command(copyCmd, filepath.Join(depPath, "libgo.so"), libPath), fmt.Sprintf("copy libgo for %v on %v", buildTarget, runtime.GOOS))

			//trick androiddeployqt into checking dependencies from libgo_base.so
			utils.RemoveAll(filepath.Join(depPath, "libgo.so"))

			if runtime.GOOS == "windows" {
				utils.RunCmd(exec.Command(copyCmd, filepath.Join(depPath, "libgo_base.so"), filepath.Join(depPath, "libgo.so*")), "")
			} else {
				utils.RunCmd(exec.Command(copyCmd, filepath.Join(depPath, "libgo_base.so"), filepath.Join(depPath, "libgo.so")), "")
			}

			var jsonStruct = &struct {
				Qt                            string `json:"qt"`
				Sdk                           string `json:"sdk"`
				SdkBuildToolsRevision         string `json:"sdkBuildToolsRevision"`
				Ndk                           string `json:"ndk"`
				Toolchainprefix               string `json:"toolchain-prefix"`
				Toolprefix                    string `json:"tool-prefix"`
				Toolchainversion              string `json:"toolchain-version"`
				Ndkhost                       string `json:"ndk-host"`
				Targetarchitecture            string `json:"target-architecture"`
				AndroidExtraLibs              string `json:"android-extra-libs"`
				AndroidPackageSourceDirectory string `json:"android-package-source-directory"`
				Qmlrootpath                   string `json:"qml-root-path"`
				Applicationbinary             string `json:"application-binary"`
			}{
				Qt:  filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "android_armv7"),
				Sdk: utils.ANDROID_SDK_DIR(),
				SdkBuildToolsRevision: "25.0.2",
				Ndk:                           utils.ANDROID_NDK_DIR(),
				Toolchainprefix:               "arm-linux-androideabi",
				Toolprefix:                    "arm-linux-androideabi",
				Toolchainversion:              "4.9",
				Ndkhost:                       runtime.GOOS + "-x86_64",
				Targetarchitecture:            "armeabi-v7a",
				AndroidExtraLibs:              filepath.Join(depPath, "libgo_base.so"),
				AndroidPackageSourceDirectory: filepath.Join(appPath, "android"),
				Qmlrootpath:                   appPath,
				Applicationbinary:             filepath.Join(depPath, "libgo.so"),
			}

			if utils.QT_DOCKER() {
				jsonStruct.AndroidExtraLibs += "," + filepath.Join(os.Getenv("HOME"), "openssl-1.0.2k", "libcrypto.so") + "," + filepath.Join(os.Getenv("HOME"), "openssl-1.0.2k", "libssl.so")
			}

			var out, err = json.Marshal(jsonStruct)
			if err != nil {
				utils.Log.WithError(err).Panicf("failed to create json-config file for androiddeployqt on %v", runtime.GOOS)
			}

			utils.Save(filepath.Join(depPath, "android-libgo.so-deployment-settings.json"), strings.Replace(string(out), `\\`, `/`, -1))
		}

	case "ios", "ios-simulator":
		{
			utils.MkdirAll(filepath.Join(appPath, buildTarget))

			var buildPath = filepath.Join(depPath, "build")
			utils.MkdirAll(filepath.Join(buildPath, "project.xcodeproj"))
			utils.MkdirAll(filepath.Join(buildPath, "Images.xcassets", "AppIcon.appiconset"))

			//add c_main_wrappers
			utils.Save(filepath.Join(depPath, "c_main_wrapper.cpp"), "#include \"libgo.h\"\nint main(int argc, char *argv[]) { go_main_wrapper(); }")
			if buildTarget == "ios" {
				utils.Save(filepath.Join(depPath, "c_main_wrapper_armv7.cpp"), "#include \"libgo_armv7.h\"\nint main(int argc, char *argv[]) { go_main_wrapper(); }")
			}

			utils.Save(filepath.Join(depPath, "gallery_plugin_import.cpp"), iosGalleryPluginImport)
			utils.Save(filepath.Join(depPath, "gallery_qml_plugin_import.cpp"), iosGalleryQmlPluginImport)

			utils.Save(filepath.Join(depPath, "qt.conf"), iosQtConf)

			//build arm64/amd64
			var cmd = exec.Command("xcrun", "clang++", "c_main_wrapper.cpp", "gallery_plugin_import.cpp", "gallery_qml_plugin_import.cpp", "-o", "build/main", "-u", "_qt_registerPlatformPlugin", "-Wl,-e,_qt_main_wrapper", "-I../..", "-L.", "-lgo")
			cmd.Args = append(cmd.Args, templater.GetiOSClang(buildTarget, "")...)
			cmd.Dir = depPath
			utils.RunCmd(cmd, fmt.Sprintf("compile for %v on %v", buildTarget, runtime.GOOS))

			var strip = exec.Command("strip", "main")
			strip.Dir = filepath.Join(depPath, "build")
			utils.RunCmd(strip, fmt.Sprintf("strip binary for %v on %v", buildTarget, runtime.GOOS))

			if buildTarget == "ios" {
				//build armv7
				cmd = exec.Command("xcrun", "clang++", "c_main_wrapper_armv7.cpp", "gallery_plugin_import.cpp", "gallery_qml_plugin_import.cpp", "-o", "build/main_armv7", "-u", "_qt_registerPlatformPlugin", "-Wl,-e,_qt_main_wrapper", "-I../..", "-L.", "-lgo_armv7")
				cmd.Args = append(cmd.Args, templater.GetiOSClang(buildTarget, "armv7")...)
				cmd.Dir = depPath
				utils.RunCmdOptional(cmd, fmt.Sprintf("compile for %v on %v", buildTarget, runtime.GOOS))

				strip = exec.Command("strip", "main_armv7")
				strip.Dir = filepath.Join(depPath, "build")
				utils.RunCmdOptional(strip, fmt.Sprintf("strip binary for %v on %v", buildTarget, runtime.GOOS))

				//binary size limits	=> https://developer.apple.com/library/ios/documentation/LanguagesUtilities/Conceptual/iTunesConnect_Guide/Chapters/SubmittingTheApp.html
				var lipo = exec.Command("xcrun", "lipo", "-create", "-arch", "arm64", "main", "-arch", "armv7", "main_armv7", "-output", "main")
				lipo.Dir = filepath.Join(depPath, "build")
				utils.RunCmdOptional(lipo, fmt.Sprintf("create fat binary for %v on %v", buildTarget, runtime.GOOS))
			}

			//create default assets
			utils.Save(filepath.Join(buildPath, "Info.plist"), iosPLIST())
			utils.Save(filepath.Join(buildPath, "Images.xcassets", "AppIcon.appiconset", "Contents.json"), iosAppIcon)
			utils.Save(filepath.Join(buildPath, "LaunchScreen.xib"), iosLaunchScreen())
			utils.Save(filepath.Join(buildPath, "project.xcodeproj", "project.pbxproj"), iosProject())

			utils.RunCmd(exec.Command(copyCmd, fmt.Sprintf("%v/%v/ios/mkspecs/macx-ios-clang/Default-568h@2x.png", utils.QT_DIR(), utils.QT_VERSION_MAJOR()), buildPath), fmt.Sprintf("copy icon for %v on %v", buildTarget, runtime.GOOS))

			//copy assets from buildTarget folder
			utils.RunCmd(exec.Command(copyCmd, "-R", fmt.Sprintf("%v/%v/", appPath, buildTarget), buildPath), fmt.Sprintf("copy assets for %v on %v", buildTarget, runtime.GOOS))
		}

	case "desktop", "rpi1", "rpi2", "rpi3":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					utils.MkdirAll(filepath.Join(appPath, runtime.GOOS))
					utils.Save(filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_sh", appName, appName)), darwinSH())
					utils.Save(filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/Info.plist", appName)), darwinPLIST())
					utils.RunCmd(exec.Command(copyCmd, "-R", fmt.Sprintf("%v/%v/", appPath, runtime.GOOS), filepath.Join(depPath, fmt.Sprintf("%v.app/", appName))), fmt.Sprintf("copy assets for %v on %v", buildTarget, runtime.GOOS))
				}

			case "linux":
				{
					utils.MkdirAll(filepath.Join(appPath, runtime.GOOS))
					utils.Save(filepath.Join(depPath, fmt.Sprintf("%v.sh", appName)), linuxSH())
					utils.RunCmd(exec.Command(copyCmd, "-R", fmt.Sprintf("%v/%v/", appPath, runtime.GOOS), depPath), fmt.Sprintf("copy assets for %v on %v", buildTarget, runtime.GOOS))
				}

			case "windows":
				{
					//TODO: icon windres
				}
			}
		}

	case "sailfish", "sailfish-emulator":
		{
			utils.MkdirAll(filepath.Join(appPath, buildTarget))

			//create default assets
			utils.MkdirAll(filepath.Join(depPath, "rpm"))
			utils.Save(filepath.Join(depPath, "rpm", appName+".spec"), sailfishSpec())
			utils.Save(filepath.Join(depPath, appName+".desktop"), sailfishDesktop())

			//copy assets from buildTarget folder
			if runtime.GOOS == "windows" {
				utils.RunCmd(exec.Command(copyCmd, filepath.Join(utils.SAILFISH_DIR(), "tutorials", "stocqt", "stocqt.png"), filepath.Join(depPath, fmt.Sprintf("harbour-%v.png*", appName))), fmt.Sprintf("copy icon for %v on %v", buildTarget, runtime.GOOS))

				var cmd = exec.Command(copyCmd, buildTarget, depPath)
				cmd.Dir = appPath
				utils.RunCmd(cmd, fmt.Sprintf("copy assets for %v on %v", buildTarget, runtime.GOOS))
			} else {
				utils.RunCmd(exec.Command(copyCmd, filepath.Join(utils.SAILFISH_DIR(), "tutorials", "stocqt", "stocqt.png"), filepath.Join(depPath, fmt.Sprintf("harbour-%v.png", appName))), fmt.Sprintf("copy icon for %v on %v", buildTarget, runtime.GOOS))

				utils.RunCmd(exec.Command(copyCmd, "-R", fmt.Sprintf("%v/%v/", appPath, buildTarget), depPath), fmt.Sprintf("copy assets for %v on %v", buildTarget, runtime.GOOS))
			}

			var errClean = sshCommand("2222", "mersdk", "cd", "/home/mersdk", "&&", "rm", "-R", buildTarget)
			if errClean != nil {
				utils.Log.WithError(errClean).Errorf("failed to cleanup for %v on %v", buildTarget, runtime.GOOS)
			}

			var errCopy = sshCommand("2222", "mersdk", "cd", strings.Replace(strings.Replace(appPath, utils.MustGoPath(), "/media/sf_GOPATH", -1)+"/deploy", "\\", "/", -1), "&&", "cp", "-R", buildTarget, "/home/mersdk")
			if errCopy != nil {
				cleanup()
				utils.Log.WithError(errClean).Panicf("failed to copy project for %v on %v", buildTarget, runtime.GOOS)
			}
		}
	}
}

func deployInternal() {

	switch buildTarget {
	case "android":
		{

			var deploy = exec.Command(filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "android_armv7", "bin", "androiddeployqt"))
			deploy.Args = append(deploy.Args,
				"--input", filepath.Join(depPath, "android-libgo.so-deployment-settings.json"),
				"--output", filepath.Join(depPath, "build"),
				"--deployment", "bundled",
				"--android-platform", "android-25",
				"--jdk", utils.JDK_DIR(),
				"--gradle",
				"--verbose",
			)

			if utils.ExistsFile(filepath.Join(appPath, "android", appName+".keystore")) {
				deploy.Args = append(deploy.Args,
					"--sign", filepath.Join(appPath, "android", appName+".keystore"),
					strings.TrimSpace(utils.Load(filepath.Join(appPath, "android", "alias.txt"))),
					"--storepass", strings.TrimSpace(utils.Load(filepath.Join(appPath, "android", "password.txt"))),
				)
			}

			deploy.Dir = filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "android_armv7", "bin")
			deploy.Env = append(deploy.Env, "JAVA_HOME="+utils.JDK_DIR())

			if runtime.GOOS == "windows" {
				utils.Save(filepath.Join(depPath, "build.bat"), fmt.Sprintf("set JAVA_HOME=%v\r\n%v", utils.JDK_DIR(), strings.Join(deploy.Args, " ")))
				utils.RunCmd(exec.Command(filepath.Join(depPath, "build.bat")), fmt.Sprintf("deploy %v on %v", buildTarget, runtime.GOOS))
				utils.RemoveAll(filepath.Join(depPath, "build.bat"))
			} else {
				utils.RunCmd(deploy, fmt.Sprintf("deploy %v on %v", buildTarget, runtime.GOOS))
			}
		}

	case "ios", "ios-simulator":
		{
			utils.RunCmd(exec.Command("xcrun", "xcodebuild", "clean", "build", "CODE_SIGN_IDENTITY=", "CODE_SIGNING_REQUIRED=NO", "CONFIGURATION_BUILD_DIR="+depPath, "-configuration", "Release", "-project", filepath.Join(depPath, "build", "project.xcodeproj")), fmt.Sprintf("deploy %v on %v", buildTarget, runtime.GOOS))
		}

	case "desktop", "rpi1", "rpi2", "rpi3", "windows":
		{
			switch {
			case buildTarget == "windows":
				{
					if utils.QT_MXE_STATIC() {
						return
					}

					var libraryPath = filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "bin")
					for _, d := range []string{"libbz2", "libfreetype-6", "libglib-2.0-0", "libharfbuzz-0", "libiconv-2", "libintl-8", "libpcre-1", "libpcre16-0", "libpng16-16", "libstdc++-6", "libwinpthread-1", "zlib1", "libeay32", "ssleay32"} {
						utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, buildTarget, runtime.GOOS))
					}
					for _, d := range []string{"libjasper-1", "libjpeg-9", "libmng-2", "libtiff-5", "libwebp-5", "liblcms2-2", "liblzma-5", "libwebpdemux-1"} {
						utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, buildTarget, runtime.GOOS))
					}

					var gccDep = "libgcc_s_sjlj-1"
					if utils.QT_MXE_ARCH() == "amd64" {
						gccDep = "libgcc_s_seh-1"
					}
					utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", gccDep)), depPath), fmt.Sprintf("copy %v for %v on %v", gccDep, buildTarget, runtime.GOOS))

					libraryPath = filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "qt5")
					utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "qml/")+"/.", depPath), fmt.Sprintf("copy qml dir for %v on %v", buildTarget, runtime.GOOS))
					utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "plugins/")+"/.", depPath), fmt.Sprintf("copy plugins dir for %v on %v", buildTarget, runtime.GOOS))

					libraryPath = filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "qt5", "bin")
					var output = utils.RunCmd(exec.Command(utils.QT_MXE_BIN("objdump"), "-x", filepath.Join(depPath, appName+ending)), fmt.Sprintf("objdump binary for %v on %v", buildTarget, runtime.GOOS))
					for lib, deps := range parser.LibDeps {
						if strings.Contains(output, lib) && lib != parser.MOC {
							for _, lib := range append(deps, lib) {
								utils.RunCmd(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("Qt5%v.dll", lib)), depPath), fmt.Sprintf("copy %v for %v on %v", lib, buildTarget, runtime.GOOS))
							}
						}
					}
					for _, d := range []string{"Qt5OpenGL", "Qt5Quick", "Qt5QuickControls2", "Qt5QuickTemplates2"} {
						utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, buildTarget, runtime.GOOS))
					}
				}

			case utils.UseMsys2():
				{
					var copyCmd = "xcopy"
					if os.Getenv("MSYSTEM") != "" {
						copyCmd = "cp"
					}

					var deploy = exec.Command(filepath.Join(utils.QT_MSYS2_DIR(), "bin", "windeployqt"))
					deploy.Args = append(deploy.Args,
						filepath.Join(depPath, appName+ending),
						fmt.Sprintf("-qmldir=%v", appPath),
						"-force")
					utils.RunCmd(deploy, fmt.Sprintf("depoy %v on %v", buildTarget, runtime.GOOS))

					var libraryPath = filepath.Join(utils.QT_MSYS2_DIR(), "bin")
					for _, d := range []string{"libbz2-1", "libfreetype-6", "libglib-2.0-0", "libharfbuzz-0", "libiconv-2", "libintl-8", "libpcre-1", "libpcre16-0", "libpng16-16", "libstdc++-6", "libwinpthread-1", "zlib1", "libgraphite2", "libicudt57", "libicuin57", "libicuuc57"} {
						utils.RunCmdOptional(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, buildTarget, runtime.GOOS))
					}

					var gccDep = "libgcc_s_dw2-1"
					if utils.QT_MSYS2_ARCH() == "amd64" {
						gccDep = "libgcc_s_seh-1"
					}

					utils.RunCmdOptional(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("%v.dll", gccDep)), depPath), fmt.Sprintf("copy %v for %v on %v", gccDep, buildTarget, runtime.GOOS))

					libraryPath = filepath.Join(utils.QT_MSYS2_DIR(), "share", "qt5")
					if os.Getenv("MSYSTEM") != "" {
						utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "qml/")+"/.", depPath), fmt.Sprintf("copy qml dir for %v on %v", buildTarget, runtime.GOOS))
						utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "plugins/")+"/.", depPath), fmt.Sprintf("copy plugins dir for %v on %v", buildTarget, runtime.GOOS))
					} else {
						utils.RunCmd(exec.Command("xcopy", "/S", "/Y", filepath.Join(libraryPath, "qml/"), depPath), fmt.Sprintf("copy qml dir for %v on %v", buildTarget, runtime.GOOS))
						utils.RunCmd(exec.Command("xcopy", "/S", "/Y", filepath.Join(libraryPath, "plugins/"), depPath), fmt.Sprintf("copy plugins dir for %v on %v", buildTarget, runtime.GOOS))
					}

					libraryPath = filepath.Join(utils.QT_MSYS2_DIR(), "bin")
					var output = utils.RunCmd(exec.Command(filepath.Join(utils.QT_MSYS2_DIR(), "bin", "objdump"), "-x", filepath.Join(depPath, appName+ending)), fmt.Sprintf("objdump binary for %v on %v", buildTarget, runtime.GOOS))
					for lib, deps := range parser.LibDeps {
						if strings.Contains(output, lib) && lib != parser.MOC {
							for _, lib := range append(deps, lib) {
								utils.RunCmd(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("Qt5%v.dll", lib)), depPath), fmt.Sprintf("copy %v for %v on %v", lib, buildTarget, runtime.GOOS))
							}
						}
					}
					utils.RunCmd(exec.Command(copyCmd, filepath.Join(libraryPath, "Qt5OpenGL.dll"), depPath), fmt.Sprintf("copy OpenGL for %v on %v", buildTarget, runtime.GOOS))

					var walkFn = func(path string, info os.FileInfo, err error) error {
						if strings.HasSuffix(info.Name(), "d.dll") {
							utils.RemoveAll(path)
						}
						return nil
					}
					filepath.Walk(depPath, walkFn)
				}

			case runtime.GOOS == "windows":
				{
					var deploy = exec.Command(filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "mingw53_32", "bin", "windeployqt"))
					deploy.Args = append(deploy.Args,
						filepath.Join(depPath, appName+ending),
						fmt.Sprintf("-qmldir=%v", appPath),
						"-force")
					utils.RunCmd(deploy, fmt.Sprintf("depoy %v on %v", buildTarget, runtime.GOOS))
				}

			case runtime.GOOS == "darwin":
				{
					if utils.UseHomeBrew() {
						return
					}

					var deploy = exec.Command(fmt.Sprintf("%v/bin/macdeployqt", utils.QT_DARWIN_DIR()))
					deploy.Args = append(deploy.Args,
						filepath.Join(depPath, fmt.Sprintf("%v.app/", appName)),
						fmt.Sprintf("-qmldir=%v", appPath))
					deploy.Dir = fmt.Sprintf("%v/bin/", utils.QT_DARWIN_DIR())
					utils.RunCmd(deploy, fmt.Sprintf("deploy %v on %v", buildTarget, runtime.GOOS))
				}

			case runtime.GOOS == "linux":
				{
					if utils.UsePkgConfig() {
						return
					}

					utils.MkdirAll(filepath.Join(depPath, "lib"))

					var (
						libraryPath   string
						lddPath       = "ldd"
						lddExtra      string
						lddOutput     string
						usesWebEngine bool
					)

					if strings.HasPrefix(buildTarget, "rpi") {
						libraryPath = fmt.Sprintf("%v/%v/%v/lib/", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), buildTarget)
						lddPath = fmt.Sprintf("%v/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf-ldd", utils.RPI_TOOLS_DIR())
						lddExtra = "--root=/"
						lddOutput = utils.RunCmd(exec.Command(lddPath, lddExtra, filepath.Join(depPath, appName)), fmt.Sprintf("ldd binary for %v on %v", buildTarget, runtime.GOOS))
					} else {
						lddOutput = utils.RunCmd(exec.Command(lddPath, filepath.Join(depPath, appName)), fmt.Sprintf("ldd binary for %v on %v", buildTarget, runtime.GOOS))
					}

					for _, dep := range strings.Split(lddOutput, "\n") {
						if strings.Contains(dep, "libQt5") || strings.Contains(dep, "libicu") {
							var libName string

							if strings.HasPrefix(buildTarget, "rpi") {
								libName = strings.TrimSpace(strings.Replace(strings.Split(dep, "=>")[0], "not found", "", -1))
							} else {
								if libraryPath == "" {
									libraryPath, libName = filepath.Split(strings.Split(dep, " ")[2])
								} else {
									_, libName = filepath.Split(strings.Split(dep, " ")[2])
								}
							}

							if utils.ExistsFile(filepath.Join(libraryPath, libName)) {
								utils.RunCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, libName), filepath.Join(depPath, "lib", libName)), fmt.Sprintf("copy %v for %v on %v", libName, buildTarget, runtime.GOOS))
							}

							if strings.Contains(dep, "WebEngine") {
								usesWebEngine = true
							}
						}
					}

					for _, libName := range []string{"DBus", "XcbQpa", "Quick", "Widgets", "EglDeviceIntegration", "EglFsKmsSupport", "OpenGL", "WaylandClient", "WaylandCompositor", "QuickControls2", "QuickTemplates2", "QuickWidgets", "QuickParticles", "CLucene", "Concurrent"} {
						if utils.ExistsFile(filepath.Join(libraryPath, fmt.Sprintf("libQt5%v.so.5", libName))) {
							utils.RunCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, fmt.Sprintf("libQt5%v.so.5", libName)), filepath.Join(depPath, "lib", fmt.Sprintf("libQt5%v.so.5", libName))), fmt.Sprintf("copy %v for %v on %v", libName, buildTarget, runtime.GOOS))
						}
					}
					if utils.ExistsFile(filepath.Join(libraryPath, "libqgsttools_p.so.1.0.0")) {
						utils.RunCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, "libqgsttools_p.so.1.0.0"), filepath.Join(depPath, "lib", "libqgsttools_p.so.1")), fmt.Sprintf("copy libqgsttools_p.so.1 for %v on %v", buildTarget, runtime.GOOS))
					}

					libraryPath = strings.TrimSuffix(libraryPath, "lib/")
					utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "qml/"), depPath), fmt.Sprintf("copy qml dir for %v on %v", buildTarget, runtime.GOOS))
					utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "plugins/"), depPath), fmt.Sprintf("copy plugins dir for %v on %v", buildTarget, runtime.GOOS))

					if usesWebEngine {
						utils.RunCmd(exec.Command("cp", filepath.Join(libraryPath, "libexec", "QtWebEngineProcess"), depPath), fmt.Sprintf("copy QtWebEngineProcess for %v on %v", buildTarget, runtime.GOOS))
						var fileList, err = ioutil.ReadDir(filepath.Join(libraryPath, "resources"))
						if err != nil {
							utils.Log.WithError(err).Error("failed to read resource folder")
						}
						for _, file := range fileList {
							utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "resources", file.Name()), depPath), fmt.Sprintf("copy resource %v for %v on %v", file.Name(), buildTarget, runtime.GOOS))
						}
						utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "translations/qtwebengine_locales/"), depPath), fmt.Sprintf("copy qtwebengine_locales dir for %v on %v", buildTarget, runtime.GOOS))
					}
				}
			}
		}

	case "sailfish", "sailfish-emulator":
		{
			if buildTarget == "sailfish-emulator" {
				var errDeploy = sshCommand("2222", "mersdk", "cd", "/home/mersdk/"+buildTarget, "&&", "mb2", "-t", "SailfishOS-i486", "build")
				if errDeploy != nil {
					cleanup()
					utils.Log.WithError(errDeploy).Errorf("failed to deploy for %v on %v", buildTarget, runtime.GOOS)
				}
			} else {
				var errDeploy = sshCommand("2222", "mersdk", "cd", "/home/mersdk/"+buildTarget, "&&", "mb2", "-t", "SailfishOS-armv7hl", "build")
				if errDeploy != nil {
					cleanup()
					utils.Log.WithError(errDeploy).Errorf("failed to deploy for %v on %v", buildTarget, runtime.GOOS)
				}
			}
		}
	}
}

func pastdeploy() {

	switch buildTarget {
	case "android":
		{
			var (
				copyCmd   string
				apkEnding string
			)

			switch runtime.GOOS {
			case "darwin", "linux":
				{
					copyCmd = "cp"
					apkEnding = "apk"
				}

			case "windows":
				{
					copyCmd = "xcopy"
					apkEnding = "apk*"
				}
			}

			if utils.ExistsFile(filepath.Join(appPath, "android", appName+".keystore")) {
				utils.RunCmd(exec.Command(copyCmd, filepath.Join(depPath, "build", "build", "outputs", "apk", "build-release-signed.apk"), filepath.Join(depPath, fmt.Sprintf("%v.%v", appName, apkEnding))), fmt.Sprintf("copy release apk for %v on %v", buildTarget, runtime.GOOS))
			} else {
				utils.RunCmd(exec.Command(copyCmd, filepath.Join(depPath, "build", "build", "outputs", "apk", "build-debug.apk"), filepath.Join(depPath, fmt.Sprintf("%v.%v", appName, apkEnding))), fmt.Sprintf("copy debug apk for %v on %v", buildTarget, runtime.GOOS))
			}

			//TODO: copy manifest to android folder and change mindSdkVersion >= 16
		}

	case "ios", "ios-simulator":
		{
			//TODO:
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					utils.RunCmd(exec.Command("mv", filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName)), filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_app", appName, appName))), fmt.Sprintf("move binary for %v on %v", buildTarget, runtime.GOOS))
					utils.RunCmd(exec.Command("mv", filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_sh", appName, appName)), filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName))), fmt.Sprintf("move script for %v on %v", buildTarget, runtime.GOOS))
				}
			}
		}

	case "sailfish", "sailfish-emulator":
		{
			var errPastDeploy = sshCommand("2222", "mersdk", "cd", "/home/mersdk/"+buildTarget+"/RPMS", "&&", "cp", "*", strings.Replace(strings.Replace(depPath, utils.MustGoPath(), "/media/sf_GOPATH", -1), "\\", "/", -1))
			if errPastDeploy != nil {
				cleanup()
				utils.Log.WithError(errPastDeploy).Panicf("failed to receive project for %v on %v", buildTarget, runtime.GOOS)
			}
		}
	}
}

func cleanup() error {
	utils.RemoveAll(filepath.Join(appPath, "cgo_main_wrapper.go"))

	if utils.QT_QMAKE_CGO() {
		rcc.QmakeCleanPath(appPath)
		moc.QmakeCleanPath(appPath)
		return nil
	}

	rcc.CleanPath(appPath)
	return moc.CleanPath(appPath)
}

func run() {

	switch buildTarget {
	case "android":
		{
			if runtime.GOOS != "windows" {
				utils.RunCmdOptional(exec.Command("killall", "adb"), fmt.Sprintf("run \"killall adb\" for %v on %v", buildTarget, runtime.GOOS))
			}
			exec.Command(filepath.Join(utils.ANDROID_SDK_DIR(), "platform-tools", "adb"), "install", "-r", filepath.Join(depPath, fmt.Sprintf("%v.apk", appName))).Start()
		}

	case /*"ios",*/ "ios-simulator":
		{
			utils.RunCmdOptional(exec.Command("xcrun", "instruments", "-w", "iPhone 7 Plus (10.2)#"), fmt.Sprintf("start simulator for %v on %v", buildTarget, runtime.GOOS))
			utils.RunCmd(exec.Command("xcrun", "simctl", "uninstall", "booted", filepath.Join(depPath, "main.app")), fmt.Sprintf("uninstall app for %v on %v", buildTarget, runtime.GOOS))
			utils.RunCmd(exec.Command("xcrun", "simctl", "install", "booted", filepath.Join(depPath, "main.app")), fmt.Sprintf("install app for %v on %v", buildTarget, runtime.GOOS))
			utils.RunCmd(exec.Command("xcrun", "simctl", "launch", "booted", fmt.Sprintf("com.identifier.%v", appName)), fmt.Sprintf("launch app for %v on %v", buildTarget, runtime.GOOS))
		}

	case "windows":
		{
			exec.Command("wine", filepath.Join(depPath, appName+ending)).Start()
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					utils.RunCmdOptional(exec.Command("open", filepath.Join(depPath, fmt.Sprintf("%v.app/", appName))), fmt.Sprintf("run binary for %v on %v", buildTarget, runtime.GOOS))
				}

			case "linux":
				{
					exec.Command(filepath.Join(depPath, fmt.Sprintf("%v.sh", appName))).Start()
				}

			case "windows":
				{
					exec.Command(filepath.Join(depPath, appName+ending)).Start()
				}
			}
		}

	case /*"sailfish",*/ "sailfish-emulator":
		{
			utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "registervm", filepath.Join(utils.SAILFISH_DIR(), "emulator", "SailfishOS Emulator", "SailfishOS Emulator.vbox")), fmt.Sprintf("register sailfish-emulator for %v on %v", buildTarget, runtime.GOOS))
			utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "sharedfolder", "add", "SailfishOS Emulator", "--name", "GOPATH", "--hostpath", utils.MustGoPath(), "--automount"), fmt.Sprintf("share GOPATH dir for %v on %v", buildTarget, runtime.GOOS))

			if runtime.GOOS == "windows" {
				utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "startvm", "SailfishOS Emulator"), fmt.Sprintf("start vbox sailfish-emulator for %v on %v", buildTarget, runtime.GOOS))
			} else {
				utils.RunCmdOptional(exec.Command("nohup", filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "startvm", "SailfishOS Emulator"), fmt.Sprintf("start vbox sailfish-emulator for %v on %v", buildTarget, runtime.GOOS))
			}

			time.Sleep(10 * time.Second)

			var errInstall = sshCommand("2223", "nemo", "sudo", "rpm", "-i", "--force", strings.Replace(strings.Replace(depPath, utils.MustGoPath(), "/media/sf_GOPATH", -1)+"/*.rpm", "\\", "/", -1))
			if errInstall != nil {
				cleanup()
				utils.Log.WithError(errInstall).Errorf("failed to install %v for %v", appName, buildTarget)
			}

			var errRun = sshCommand("2223", "nemo", "nohup", "/usr/bin/harbour-"+appName, ">", "/dev/null", "2>&1", "&")
			if errRun != nil {
				cleanup()
				utils.Log.WithError(errRun).Errorf("failed to run %v for %v", appName, buildTarget)
			}
		}
	}
}

//darwin
func darwinSH() string {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	fmt.Fprint(bb, "#!/bin/bash\n")
	fmt.Fprint(bb, "cd \"${0%/*}\"\n")
	fmt.Fprintf(bb, "./%v_app \"$@\"", appName)

	return bb.String()
}

func darwinPLIST() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>English</string>

	<key>CFBundleExecutable</key>
	<string>%v</string>

	<key>CFBundleIconFile</key>
	<string>%v.icns</string>

	<key>CFBundleIdentifier</key>
	<string>com.identifier.%v</string>

	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>

	<key>CFBundleName</key>
	<string>%v</string>

	<key>CFBundlePackageType</key>
	<string>APPL</string>

	<key>CFBundleShortVersionString</key>
	<string>1.0.0.0</string>

	<key>CFBundleVersion</key>
	<string>1</string>

	<key>NSPrincipalClass</key>
	<string>NSApplication</string>

	<key>NSHighResolutionCapable</key>
	<string>True</string>
</dict>
</plist>
`, appName, appName, appName, appName)
}

//linux
func linuxSH() string {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	fmt.Fprint(bb, "#!/bin/bash\n")
	fmt.Fprint(bb, "appname=`basename $0 | sed s,\\.sh$,,`\n\n")
	fmt.Fprint(bb, "dirname=`dirname $0`\n")
	fmt.Fprint(bb, "tmp=\"${dirname#?}\"\n\n")
	fmt.Fprint(bb, "if [ \"${dirname%$tmp}\" != \"/\" ]; then\n")
	fmt.Fprint(bb, "dirname=$PWD/$dirname\n")
	fmt.Fprint(bb, "fi\n")

	if strings.HasPrefix(buildTarget, "rpi") {
		fmt.Fprint(bb, "export DISPLAY=:0\n")
		fmt.Fprint(bb, "export LD_PRELOAD=\"/opt/vc/lib/libGLESv2.so /opt/vc/lib/libEGL.so\"\n")
	}

	if utils.UsePkgConfig() {
		var (
			libDir  = strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=libdir", "Qt5Core"), fmt.Sprintf("get lib dir for %v on %v", buildTarget, runtime.GOOS)))
			miscDir = utils.QT_MISC_DIR()
		)
		fmt.Fprintf(bb, "export LD_LIBRARY_PATH=%v\n", libDir)
		fmt.Fprintf(bb, "export QT_PLUGIN_PATH=$%v\n", filepath.Join(miscDir, "plugins"))
		fmt.Fprintf(bb, "export QML_IMPORT_PATH=%v\n", filepath.Join(miscDir, "qml"))
		fmt.Fprintf(bb, "export QML2_IMPORT_PATH=%v\n", filepath.Join(miscDir, "qml"))
	} else {
		fmt.Fprint(bb, "export LD_LIBRARY_PATH=$dirname/lib\n")
		fmt.Fprint(bb, "export QT_PLUGIN_PATH=$dirname/plugins\n")
		fmt.Fprint(bb, "export QML_IMPORT_PATH=$dirname/qml\n")
		fmt.Fprint(bb, "export QML2_IMPORT_PATH=$dirname/qml\n")
	}
	fmt.Fprint(bb, "$dirname/$appname \"$@\"\n")

	return bb.String()
}

//ios
func iosPLIST() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>en</string>
	<key>CFBundleExecutable</key>
	<string>main</string>
	<key>CFBundleIdentifier</key>
	<string>com.identifier.%v</string>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundleName</key>
	<string>%v</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleShortVersionString</key>
	<string>1.0</string>
	<key>CFBundleSignature</key>
	<string>????</string>
	<key>CFBundleVersion</key>
	<string>1.0</string>
	<key>LSRequiresIPhoneOS</key>
	<true/>
	<key>UILaunchStoryboardName</key>
	<string>LaunchScreen</string>%v
	<key>UISupportedInterfaceOrientations</key>
	<array>
		<string>UIInterfaceOrientationPortrait</string>
		<string>UIInterfaceOrientationPortraitUpsideDown</string>
		<string>UIInterfaceOrientationLandscapeLeft</string>
		<string>UIInterfaceOrientationLandscapeRight</string>
	</array>
	<key>UISupportedInterfaceOrientations~ipad</key>
	<array>
		<string>UIInterfaceOrientationPortrait</string>
		<string>UIInterfaceOrientationPortraitUpsideDown</string>
		<string>UIInterfaceOrientationLandscapeLeft</string>
		<string>UIInterfaceOrientationLandscapeRight</string>
	</array>
	<key>QtRunLoopIntegrationDisableSeparateStack</key>
	<true/>
	<key>NSAppTransportSecurity</key>
	<dict>
		<key>NSAllowsArbitraryLoads</key>
		<true/>
	</dict>
</dict>
</plist>
`, appName, appName, func() string {
		if strings.HasPrefix(runtime.Version(), "go1.7") || strings.HasPrefix(runtime.Version(), "devel") {
			return ""
		}
		return `
	<key>UIRequiredDeviceCapabilities</key>
	<array>
		<string>arm64</string>
	</array>`
	}())
}

func iosLaunchScreen() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<document type="com.apple.InterfaceBuilder3.CocoaTouch.XIB" version="3.0" toolsVersion="10117" systemVersion="15E65" targetRuntime="iOS.CocoaTouch" propertyAccessControl="none" useAutolayout="YES" launchScreen="YES" useTraitCollections="YES">
	    <dependencies>
	        <deployment identifier="iOS"/>
	        <plugIn identifier="com.apple.InterfaceBuilder.IBCocoaTouchPlugin" version="10085"/>
	        <capability name="Constraints with non-1.0 multipliers" minToolsVersion="5.1"/>
	    </dependencies>
	    <objects>
	        <placeholder placeholderIdentifier="IBFilesOwner" id="-1" userLabel="File's Owner"/>
	        <placeholder placeholderIdentifier="IBFirstResponder" id="-2" customClass="UIResponder"/>
	        <view contentMode="scaleToFill" id="iN0-l3-epB">
	            <rect key="frame" x="0.0" y="0.0" width="480" height="480"/>
	            <autoresizingMask key="autoresizingMask" widthSizable="YES" heightSizable="YES"/>
	            <subviews>
	                <label opaque="NO" clipsSubviews="YES" userInteractionEnabled="NO" contentMode="left" horizontalHuggingPriority="251" verticalHuggingPriority="251" text="%v" textAlignment="center" lineBreakMode="middleTruncation" baselineAdjustment="alignBaselines" minimumFontSize="18" translatesAutoresizingMaskIntoConstraints="NO" id="kId-c2-rCX">
	                    <rect key="frame" x="20" y="140" width="441" height="43"/>
	                    <fontDescription key="fontDescription" type="boldSystem" pointSize="36"/>
	                    <color key="textColor" red="0.0" green="0.0" blue="0.0" alpha="1" colorSpace="calibratedRGB"/>
	                    <nil key="highlightedColor"/>
	                </label>
	            </subviews>
	            <color key="backgroundColor" white="1" alpha="1" colorSpace="custom" customColorSpace="calibratedWhite"/>
	            <constraints>
	                <constraint firstItem="kId-c2-rCX" firstAttribute="centerY" secondItem="iN0-l3-epB" secondAttribute="bottom" multiplier="1/3" constant="1" id="Kid-kn-2rF"/>
	                <constraint firstAttribute="centerX" secondItem="kId-c2-rCX" secondAttribute="centerX" id="Koa-jz-hwk"/>
	                <constraint firstItem="kId-c2-rCX" firstAttribute="leading" secondItem="iN0-l3-epB" secondAttribute="leading" constant="20" symbolic="YES" id="fvb-Df-36g"/>
	            </constraints>
	            <nil key="simulatedStatusBarMetrics"/>
	            <freeformSimulatedSizeMetrics key="simulatedDestinationMetrics"/>
	            <point key="canvasLocation" x="404" y="445"/>
	        </view>
	    </objects>
	</document>
	`, appName)
}

const iosAppIcon = `{
  "images" : [
    {
      "idiom" : "iphone",
      "size" : "29x29",
      "scale" : "2x"
    },
    {
      "idiom" : "iphone",
      "size" : "29x29",
      "scale" : "3x"
    },
    {
      "idiom" : "iphone",
      "size" : "40x40",
      "scale" : "2x"
    },
    {
      "idiom" : "iphone",
      "size" : "40x40",
      "scale" : "3x"
    },
    {
      "idiom" : "iphone",
      "size" : "60x60",
      "scale" : "2x"
    },
    {
      "idiom" : "iphone",
      "size" : "60x60",
      "scale" : "3x"
    },
    {
      "idiom" : "ipad",
      "size" : "29x29",
      "scale" : "1x"
    },
    {
      "idiom" : "ipad",
      "size" : "29x29",
      "scale" : "2x"
    },
    {
      "idiom" : "ipad",
      "size" : "40x40",
      "scale" : "1x"
    },
    {
      "idiom" : "ipad",
      "size" : "40x40",
      "scale" : "2x"
    },
    {
      "idiom" : "ipad",
      "size" : "76x76",
      "scale" : "1x"
    },
    {
      "idiom" : "ipad",
      "size" : "76x76",
      "scale" : "2x"
    }
  ],
  "info" : {
    "version" : 1,
    "author" : "xcode"
  }
}
`

func iosProject() string {
	return fmt.Sprintf(`// !$*UTF8*$!
{
	archiveVersion = 1;
	classes = {
	};
	objectVersion = 46;
	objects = {

/* Begin PBXBuildFile section */
		254BB84F1B1FD08900C56DE9 /* Images.xcassets in Resources */ = {isa = PBXBuildFile; fileRef = 254BB84E1B1FD08900C56DE9 /* Images.xcassets */; };
		254BB8681B1FD16500C56DE9 /* main in Resources */ = {isa = PBXBuildFile; fileRef = 254BB8671B1FD16500C56DE9 /* main */; };
		25916F411CE65FF600695115 /* LaunchScreen.xib in Resources */ = {isa = PBXBuildFile; fileRef = 25916F401CE65FF600695115 /* LaunchScreen.xib */; };
		25F26AED1CE6675E0045FFBA /* Default-568h@2x.png in Resources */ = {isa = PBXBuildFile; fileRef = 25F26AEC1CE6675E0045FFBA /* Default-568h@2x.png */; };
/* End PBXBuildFile section */

/* Begin PBXFileReference section */
		254BB83E1B1FD08900C56DE9 /* main.app */ = {isa = PBXFileReference; explicitFileType = wrapper.application; includeInIndex = 0; path = main.app; sourceTree = BUILT_PRODUCTS_DIR; };
		254BB8421B1FD08900C56DE9 /* Info.plist */ = {isa = PBXFileReference; lastKnownFileType = text.plist.xml; path = Info.plist; sourceTree = "<group>"; };
		254BB84E1B1FD08900C56DE9 /* Images.xcassets */ = {isa = PBXFileReference; lastKnownFileType = folder.assetcatalog; path = Images.xcassets; sourceTree = "<group>"; };
		254BB8671B1FD16500C56DE9 /* main */ = {isa = PBXFileReference; lastKnownFileType = "compiled.mach-o.executable"; path = main; sourceTree = "<group>"; };
		25916F401CE65FF600695115 /* LaunchScreen.xib */ = {isa = PBXFileReference; fileEncoding = 4; lastKnownFileType = file.xib; path = LaunchScreen.xib; sourceTree = "<group>"; };
		25F26AEC1CE6675E0045FFBA /* Default-568h@2x.png */ = {isa = PBXFileReference; lastKnownFileType = image.png; path = "Default-568h@2x.png"; sourceTree = "<group>"; };
/* End PBXFileReference section */

/* Begin PBXGroup section */
		254BB8351B1FD08900C56DE9 = {
			isa = PBXGroup;
			children = (
				254BB8671B1FD16500C56DE9 /* main */,
				254BB8421B1FD08900C56DE9 /* Info.plist */,
				254BB84E1B1FD08900C56DE9 /* Images.xcassets */,
				25916F401CE65FF600695115 /* LaunchScreen.xib */,
				25F26AEC1CE6675E0045FFBA /* Default-568h@2x.png */,
				254BB83F1B1FD08900C56DE9 /* products */,
			);
			sourceTree = "<group>";
			usesTabs = 0;
		};
		254BB83F1B1FD08900C56DE9 /* products */ = {
			isa = PBXGroup;
			children = (
				254BB83E1B1FD08900C56DE9 /* main.app */,
			);
			name = products;
			sourceTree = "<group>";
		};
/* End PBXGroup section */

/* Begin PBXNativeTarget section */
		254BB83D1B1FD08900C56DE9 /* main */ = {
			isa = PBXNativeTarget;
			buildConfigurationList = 254BB8611B1FD08900C56DE9 /* Build configuration list for PBXNativeTarget "main" */;
			buildPhases = (
				254BB83C1B1FD08900C56DE9 /* Resources */,
				259BC5361CE6BA19005B5A05 /* ShellScript */,
			);
			buildRules = (
			);
			dependencies = (
			);
			name = main;
			productName = main;
			productReference = 254BB83E1B1FD08900C56DE9 /* main.app */;
			productType = "com.apple.product-type.application";
		};
/* End PBXNativeTarget section */

/* Begin PBXProject section */
		254BB8361B1FD08900C56DE9 /* Project object */ = {
			isa = PBXProject;
			attributes = {
				LastUpgradeCheck = 0630;
				ORGANIZATIONNAME = Developer;
				TargetAttributes = {
					254BB83D1B1FD08900C56DE9 = {
						CreatedOnToolsVersion = 6.3.1;
					};
				};
			};
			buildConfigurationList = 254BB8391B1FD08900C56DE9 /* Build configuration list for PBXProject "project" */;
			compatibilityVersion = "Xcode 3.2";
			developmentRegion = English;
			hasScannedForEncodings = 0;
			knownRegions = (
				en,
				Base,
			);
			mainGroup = 254BB8351B1FD08900C56DE9;
			productRefGroup = 254BB83F1B1FD08900C56DE9 /* products */;
			projectDirPath = "";
			projectRoot = "";
			targets = (
				254BB83D1B1FD08900C56DE9 /* main */,
			);
		};
/* End PBXProject section */

/* Begin PBXResourcesBuildPhase section */
		254BB83C1B1FD08900C56DE9 /* Resources */ = {
			isa = PBXResourcesBuildPhase;
			buildActionMask = 2147483647;
			files = (
				254BB8681B1FD16500C56DE9 /* main in Resources */,
				25F26AED1CE6675E0045FFBA /* Default-568h@2x.png in Resources */,
				25916F411CE65FF600695115 /* LaunchScreen.xib in Resources */,
				254BB84F1B1FD08900C56DE9 /* Images.xcassets in Resources */,
			);
			runOnlyForDeploymentPostprocessing = 0;
		};
/* End PBXResourcesBuildPhase section */

/* Begin PBXShellScriptBuildPhase section */
		259BC5361CE6BA19005B5A05 /* ShellScript */ = {
			isa = PBXShellScriptBuildPhase;
			buildActionMask = 2147483647;
			files = (
			);
			inputPaths = (
				"$(TARGET_BUILD_DIR)/$(EXECUTABLE_PATH)",
			);
			outputPaths = (
			);
			runOnlyForDeploymentPostprocessing = 0;
			shellPath = /bin/sh;
			shellScript = "cp %v/qt.conf $CODESIGNING_FOLDER_PATH/qt.conf;  test -d $CODESIGNING_FOLDER_PATH/qt_qml && rm -r $CODESIGNING_FOLDER_PATH/qt_qml;  mkdir -p $CODESIGNING_FOLDER_PATH/qt_qml &&  for p in %v/%v/ios/qml; do rsync -r --exclude='*.a' --exclude='*.prl' --exclude='*.qmltypes'  $p/ $CODESIGNING_FOLDER_PATH/qt_qml; done";
			showEnvVarsInLog = 0;
		};
/* End PBXShellScriptBuildPhase section */

/* Begin XCBuildConfiguration section */
		254BB8601B1FD08900C56DE9 /* Release */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
				ALWAYS_SEARCH_USER_PATHS = NO;
				CLANG_CXX_LANGUAGE_STANDARD = "gnu++0x";
				CLANG_CXX_LIBRARY = "libc++";
				CLANG_ENABLE_MODULES = YES;
				CLANG_ENABLE_OBJC_ARC = YES;
				CLANG_WARN_BOOL_CONVERSION = YES;
				CLANG_WARN_CONSTANT_CONVERSION = YES;
				CLANG_WARN_DIRECT_OBJC_ISA_USAGE = YES_ERROR;
				CLANG_WARN_EMPTY_BODY = YES;
				CLANG_WARN_ENUM_CONVERSION = YES;
				CLANG_WARN_INT_CONVERSION = YES;
				CLANG_WARN_OBJC_ROOT_CLASS = YES_ERROR;
				CLANG_WARN_UNREACHABLE_CODE = YES;
				CLANG_WARN__DUPLICATE_METHOD_MATCH = YES;
				"CODE_SIGN_IDENTITY[sdk=iphoneos*]" = "iPhone Developer";
				COPY_PHASE_STRIP = NO;
				DEBUG_INFORMATION_FORMAT = "dwarf-with-dsym";
				ENABLE_NS_ASSERTIONS = NO;
				ENABLE_STRICT_OBJC_MSGSEND = YES;
				GCC_C_LANGUAGE_STANDARD = gnu99;
				GCC_NO_COMMON_BLOCKS = YES;
				GCC_WARN_64_TO_32_BIT_CONVERSION = YES;
				GCC_WARN_ABOUT_RETURN_TYPE = YES_ERROR;
				GCC_WARN_UNDECLARED_SELECTOR = YES;
				GCC_WARN_UNINITIALIZED_AUTOS = YES_AGGRESSIVE;
				GCC_WARN_UNUSED_FUNCTION = YES;
				GCC_WARN_UNUSED_VARIABLE = YES;
				IPHONEOS_DEPLOYMENT_TARGET = 10.0;
				MTL_ENABLE_DEBUG_INFO = NO;
				SDKROOT = iphoneos;
				TARGETED_DEVICE_FAMILY = "1,2";
				VALIDATE_PRODUCT = YES;
			};
			name = Release;
		};
		254BB8631B1FD08900C56DE9 /* Release */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
				ASSETCATALOG_COMPILER_APPICON_NAME = AppIcon;
				INFOPLIST_FILE = Info.plist;
				LD_RUNPATH_SEARCH_PATHS = "$(inherited) @executable_path/Frameworks";
				PRODUCT_NAME = "$(TARGET_NAME)";
			};
			name = Release;
		};
/* End XCBuildConfiguration section */

/* Begin XCConfigurationList section */
		254BB8391B1FD08900C56DE9 /* Build configuration list for PBXProject "project" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				254BB8601B1FD08900C56DE9 /* Release */,
			);
			defaultConfigurationIsVisible = 0;
			defaultConfigurationName = Release;
		};
		254BB8611B1FD08900C56DE9 /* Build configuration list for PBXNativeTarget "main" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				254BB8631B1FD08900C56DE9 /* Release */,
			);
			defaultConfigurationIsVisible = 0;
			defaultConfigurationName = Release;
		};
/* End XCConfigurationList section */
	};
	rootObject = 254BB8361B1FD08900C56DE9 /* Project object */;
}
`, depPath, utils.QT_DIR(), utils.QT_VERSION_MAJOR())
}

const (
	iosGalleryPluginImport = `#include <QtPlugin>
Q_IMPORT_PLUGIN(QIOSIntegrationPlugin)
Q_IMPORT_PLUGIN(AVFServicePlugin)
Q_IMPORT_PLUGIN(AVFMediaPlayerServicePlugin)
Q_IMPORT_PLUGIN(AudioCaptureServicePlugin)
Q_IMPORT_PLUGIN(CoreAudioPlugin)
Q_IMPORT_PLUGIN(QM3uPlaylistPlugin)
//Q_IMPORT_PLUGIN(ContextPlugin)
//Q_IMPORT_PLUGIN(QDDSPlugin)
Q_IMPORT_PLUGIN(QGifPlugin)
Q_IMPORT_PLUGIN(QICNSPlugin)
Q_IMPORT_PLUGIN(QICOPlugin)
Q_IMPORT_PLUGIN(QJpegPlugin)
Q_IMPORT_PLUGIN(QMacJp2Plugin)
Q_IMPORT_PLUGIN(QTgaPlugin)
Q_IMPORT_PLUGIN(QTiffPlugin)
Q_IMPORT_PLUGIN(QWbmpPlugin)
Q_IMPORT_PLUGIN(QWebpPlugin)
Q_IMPORT_PLUGIN(QQmlDebuggerServiceFactory)
Q_IMPORT_PLUGIN(QQmlInspectorServiceFactory)
Q_IMPORT_PLUGIN(QLocalClientConnectionFactory)
Q_IMPORT_PLUGIN(QQmlNativeDebugConnectorFactory)
Q_IMPORT_PLUGIN(QQuickProfilerAdapterFactory)
Q_IMPORT_PLUGIN(QQmlProfilerServiceFactory)
Q_IMPORT_PLUGIN(QQmlDebugServerFactory)
Q_IMPORT_PLUGIN(QTcpServerConnectionFactory)
//Q_IMPORT_PLUGIN(QGenericEnginePlugin)
`

	iosGalleryQmlPluginImport = `#include <QtPlugin>
Q_IMPORT_PLUGIN(QtQuick2Plugin)
Q_IMPORT_PLUGIN(QMultimediaDeclarativeModule)
Q_IMPORT_PLUGIN(QtQuickLayoutsPlugin)
Q_IMPORT_PLUGIN(QtQuickControls2Plugin)
Q_IMPORT_PLUGIN(QtQuickControls2MaterialStylePlugin)
Q_IMPORT_PLUGIN(QtQuickControls2UniversalStylePlugin)
Q_IMPORT_PLUGIN(QtQuick2DialogsPlugin)
Q_IMPORT_PLUGIN(QtQuickControls1Plugin)
Q_IMPORT_PLUGIN(QmlFolderListModelPlugin)
Q_IMPORT_PLUGIN(QmlSettingsPlugin)
Q_IMPORT_PLUGIN(QtQuickTemplates2Plugin)
Q_IMPORT_PLUGIN(QtQuick2DialogsPrivatePlugin)
Q_IMPORT_PLUGIN(QtQuick2WindowPlugin)
Q_IMPORT_PLUGIN(QtQmlModelsPlugin)
Q_IMPORT_PLUGIN(QtQuickExtrasPlugin)
Q_IMPORT_PLUGIN(QtGraphicalEffectsPlugin)
Q_IMPORT_PLUGIN(QtGraphicalEffectsPrivatePlugin)
Q_IMPORT_PLUGIN(QWebViewModule)
`

	iosQtConf = `[Paths]
Imports = qt_qml
Qml2Imports = qt_qml
`
)

//sailfish
const (
	sailfishSpecTemplate = `#
# Do NOT Edit the Auto-generated Part!
# Generated by: spectacle version 0.27
#

Name:       harbour-${NAME}

# >> macros
# << macros

Summary:    Put your summary here
Version:    0.1
Release:    1
Group:      Qt/Qt
License:    MIT
Source0:    %{name}-%{version}.tar.bz2
#Requires:  mapplauncherd-booster-silica-qt5
#Requires:  nemo-qml-plugin-thumbnailer-qt5
Requires:   sailfishsilica-qt5
#Requires:  qt5-qtdocgallery
BuildRequires:  pkgconfig(sailfishapp)
BuildRequires:  pkgconfig(Qt5Quick)
BuildRequires:  pkgconfig(Qt5Qml)
BuildRequires:  pkgconfig(Qt5Core)
#BuildRequires: pkgconfig(qdeclarative5-boostable)
BuildRequires:  desktop-file-utils

%description
Put your description here


%prep
%setup -q -n %{name}-%{version}

# >> setup
# << setup

%build
# >> build pre
# << build pre

# >> build post
# << build post

%install
rm -rf %{buildroot}
# >> install pre
# << install pre
install -d %{buildroot}%{_bindir}
install -p -m 0755 %(pwd)/%{name} %{buildroot}%{_bindir}/%{name}
install -d %{buildroot}%{_datadir}/applications
install -d %{buildroot}%{_datadir}/%{name}
install -d %{buildroot}%{_datadir}/icons/hicolor/86x86/apps
install -m 0444 -t %{buildroot}%{_datadir}/icons/hicolor/86x86/apps %{name}.png
install -p %(pwd)/${NAME}.desktop %{buildroot}%{_datadir}/applications/%{name}.desktop

# >> install post
# << install post

desktop-file-install --delete-original       \
  --dir %{buildroot}%{_datadir}/applications             \
   %{buildroot}%{_datadir}/applications/*.desktop

%files
%defattr(-,root,root,-)
%{_bindir}
%{_datadir}/%{name}
%{_datadir}/icons/hicolor/86x86/apps
%{_datadir}/applications/%{name}.desktop

# >> files
# << files`
)

func sailfishDesktop() string {
	return fmt.Sprintf(`[Desktop Entry]
Encoding=UTF-8
Version=1.0
Type=Application
X-Nemo-Application-Type=generic
Comment=Put your comment here
Name=%v
Icon=harbour-%v
Exec=harbour-%v`, appName, appName, appName)
}

func sailfishSpec() string {
	return strings.Replace(sailfishSpecTemplate, "${NAME}", appName, -1)
}

func sshCommand(port, login string, cmd ...string) error {

	var emuType = func() string {
		if port == "2222" {
			return "engine"
		}
		return "SailfishOS_Emulator"
	}()

	var signer, errPrivKey = ssh.ParsePrivateKey([]byte(utils.Load(filepath.Join(utils.SAILFISH_DIR(), "vmshare", "ssh", "private_keys", emuType, login))))
	if errPrivKey != nil {
		return errPrivKey
	}

	var client, errDial = ssh.Dial("tcp", "localhost:"+port, &ssh.ClientConfig{User: login, Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)}})
	if errDial != nil {
		return errDial
	}
	defer client.Close()

	var sess, errSess = client.NewSession()
	if errSess != nil {
		return errSess
	}

	var output, errCmd = sess.CombinedOutput(strings.Join(cmd, " "))
	if errCmd != nil {
		return errors.New(string(output))
	}

	return nil
}
