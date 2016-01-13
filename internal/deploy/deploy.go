package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

var (
	appPath, appName       string
	depPath                string
	buildMode, buildTarget string
	ending                 string
)

func main() {

	args()

	switch buildMode {
	case "build", "test":
		{
			qrc()
			build()
			predeploy()
			deploy()
			pastdeploy()
			cleanup()
		}
	}

	switch buildMode {
	case "run", "test":
		{
			run()
		}
	}
}

func args() {

	switch len(os.Args) {
	case 1:
		buildMode = "test"
		buildTarget = "desktop"
		appPath, _ = os.Getwd()

	case 2:
		buildMode = os.Args[1]
		buildTarget = "desktop"
		appPath, _ = os.Getwd()

	case 3:
		buildMode = os.Args[1]
		buildTarget = os.Args[2]
		appPath, _ = os.Getwd()

	case 4:
		buildMode = os.Args[1]
		buildTarget = os.Args[2]
		appPath = os.Args[3]
	}

	switch buildMode {
	case "build", "run", "test":
		{
			switch buildTarget {
			case "desktop", "android":

			default:
				{
					fmt.Println("usage:", "{ build | run | test }", "({ desktop | android })", "( dir )")
					os.Exit(1)
				}
			}
		}
	default:
		{
			fmt.Println("usage:", "{ build | run | test }", "({ desktop | android })", "( dir )")
			os.Exit(1)
		}
	}

	if !filepath.IsAbs(appPath) {
		appPath = utils.GetAbsPath(appPath)
	}
	appName = filepath.Base(appPath)

	switch buildTarget {
	case "android":
		{
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
			utils.MakeFolder(depPath)
		}
	}

	if runtime.GOOS == "windows" && buildTarget != "android" {
		ending = ".exe"
	}
}

func qrc() {

	utils.MakeFolder(filepath.Join(appPath, "qml"))

	var (
		rccPath string
		qmlGo   = filepath.Join(appPath, "qtQrc.go")
		qmlQrc  = filepath.Join(appPath, "qtQrc.qrc")
		qmlCpp  = filepath.Join(appPath, "qtQrc.cpp")
	)

	switch buildTarget {
	case "android":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					rccPath = "/usr/local/Qt5.5.1/5.5/android_armv7/bin/rcc"
				}

			case "linux":
				{

					switch runtime.GOARCH {
					case "amd64":
						{
							rccPath = "/usr/local/Qt5.5.1/5.5/android_armv7/bin/rcc"
						}

					case "386":
						{
							rccPath = "/usr/local/Qt5.5.1/5.5/android_armv7/bin/rcc"
						}
					}
				}

			case "windows":
				{
					rccPath = "C:\\Qt\\Qt5.5.1\\5.5\\android_armv7\\bin\\rcc.exe"
				}
			}
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					rccPath = "/usr/local/Qt5.5.1/5.5/clang_64/bin/rcc"
				}

			case "linux":
				{

					switch runtime.GOARCH {
					case "amd64":
						{
							rccPath = "/usr/local/Qt5.5.1/5.5/gcc_64/bin/rcc"
						}

					case "386":
						{
							rccPath = "/usr/local/Qt5.5.1/5.5/gcc/bin/rcc"
						}
					}
				}

			case "windows":
				{
					rccPath = "C:\\Qt\\Qt5.5.1\\5.5\\mingw492_32\\bin\\rcc.exe"
				}
			}
		}
	}

	utils.Save(qmlGo, qmlHeader())

	var rcc = exec.Command(rccPath)
	rcc.Args = append(rcc.Args,
		"-project",
		"-o", qmlQrc)
	rcc.Dir = filepath.Join(appPath, "qml")
	runCmd(rcc, "qrc.qrc")

	utils.Save(qmlQrc, strings.Replace(utils.Load(qmlQrc), "<file>./", "<file>qml/", -1))

	rcc = exec.Command(rccPath)
	rcc.Args = append(rcc.Args,
		"-name", appName,
		"-o", qmlCpp,
		qmlQrc)
	runCmd(rcc, "qrc.cpp")
}

func qmlHeader() string {

	var hloc string

	switch runtime.GOOS {
	case "darwin", "linux":
		{
			hloc = "/usr/local"
		}
	case "windows":
		{
			hloc = "C:/Qt"
		}
	}

	return fmt.Sprintf(`package main

  	/*
  	#cgo +build android,arm LDFLAGS: -L%v/Qt5.5.1/5.5/android_armv7/lib -lQt5Core
		#cgo +build darwin,amd64 LDFLAGS: -F/usr/local/Qt5.5.1/5.5/clang_64/lib -framework QtCore

	  #cgo +build linux,amd64 LDFLAGS: -Wl,-rpath,/usr/local/Qt5.5.1/5.5/gcc_64/lib
	  #cgo +build linux,amd64 LDFLAGS: -L/usr/local/Qt5.5.1/5.5/gcc_64/lib -lQt5Core

	  #cgo +build linux,386 LDFLAGS: -Wl,-rpath,/usr/local/Qt5.5.1/5.5/gcc/lib
	  #cgo +build linux,386 LDFLAGS: -L/usr/local/Qt5.5.1/5.5/gcc/lib -lQt5Core

	  #cgo +build windows,386 LDFLAGS: -LC:/Qt/Qt5.5.1/5.5/mingw492_32/lib -lQt5Core
	  */
	  import "C"`, hloc)
}

func build() {

	var (
		ldFlags    string
		outputFile string
		env        map[string]string
	)

	switch buildTarget {
	case "android":
		{
			ldFlags = "-ldflags=\"-s\" \"-w\""
			outputFile = filepath.Join(depPath, "libgo.so")

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
		}
	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					ldFlags = "-ldflags=\"-w\" \"-r=/usr/local/Qt5.5.1/5.5/clang_64/lib\""
					outputFile = filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName))
				}

			case "linux":
				{
					ldFlags = "-ldflags=\"-s\" \"-w\""
					outputFile = filepath.Join(depPath, appName)
				}

			case "windows":
				{
					ldFlags = "-ldflags=\"-s\" \"-w\" \"-H=windowsgui\""
					outputFile = filepath.Join(depPath, appName)
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
	}

	var cmd = exec.Command("go")
	cmd.Args = append(cmd.Args,
		"build",
		"-p", strconv.Itoa(runtime.NumCPU()),
		ldFlags,
		"-o", outputFile+ending)

	cmd.Dir = appPath

	if buildTarget != "desktop" || runtime.GOOS == "windows" {
		if buildTarget != "desktop" {
			cmd.Args = append(cmd.Args, "-buildmode", "c-shared") //TODO: pie in go 1.6
		}
		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}
	}
	runCmd(cmd, "build")
}

func predeploy() {

	var copyCmd string

	switch runtime.GOOS {
	case "windows":
		copyCmd = "xcopy"
	default:
		copyCmd = "cp"
	}

	switch buildTarget {
	case "android":
		{
			utils.MakeFolder(filepath.Join(appPath, "android"))

			for _, dir := range []string{"drawable-hdpi", "drawable-ldpi", "drawable-mdpi"} {
				utils.MakeFolder(filepath.Join(depPath, "build", "res", dir))
				runCmdOptional(exec.Command(copyCmd, filepath.Join(appPath, "android", "icon.png"), filepath.Join(depPath, "build", "res", dir, "icon.png")), "predeploy.cpicon")
			}

			var libPath = filepath.Join(depPath, "build", "libs", "armeabi-v7a")
			utils.MakeFolder(libPath)
			runCmd(exec.Command(copyCmd, filepath.Join(depPath, "libgo.so"), libPath), "predeploy.cplib")

			var (
				qtPrefix      string
				androidPrefix string
				ndkhost       string
			)

			switch runtime.GOOS {
			case "darwin", "linux":
				{
					qtPrefix = "/usr/local"
					androidPrefix = "/opt"
					ndkhost = runtime.GOOS + "-x86_64"

					runCmd(exec.Command("cp", "/usr/local/Qt5.5.1/5.5/android_armv7/jar/QtAndroid-bundled.jar", filepath.Join(depPath, "build", "libs")), "predeploy.cpqtandroid")

					runCmd(exec.Command("zip", "-d", filepath.Join(depPath, "build", "libs", "QtAndroid-bundled.jar"), "org/qtproject/qt5/android/QtNative.class"), "predeploy.patchqtandroid_main")
					for i := 1; i < 19; i++ {
						runCmd(exec.Command("zip", "-d", filepath.Join(depPath, "build", "libs", "QtAndroid-bundled.jar"), fmt.Sprintf("org/qtproject/qt5/android/QtNative$%v.class", i)), fmt.Sprintf("predeploy.patchqtandroid_%v", i))
					}
				}

			case "windows":
				{
					qtPrefix = "C:\\Qt"
					androidPrefix = "C:\\android"
					ndkhost = runtime.GOOS

					var (
						jdkVersion = strings.Split(runCmd(exec.Command("java", "-version"), "predeploy.jdk"), "\"")[1]
						jarPath    = fmt.Sprintf("C:\\Program Files\\Java\\jdk%v\\bin\\jar.exe", jdkVersion)
					)

					runCmd(exec.Command("xcopy", "C:\\Qt\\Qt5.5.1\\5.5\\android_armv7\\jar\\QtAndroid-bundled.jar", filepath.Join(depPath, "build", "libs")), "predeploy.cpqtandroid")

					var jarUnzip = exec.Command(jarPath, "-xf", filepath.Join(depPath, "build", "libs", "QtAndroid-bundled.jar"))
					jarUnzip.Dir = filepath.Join(depPath, "build", "libs")
					runCmd(jarUnzip, "predeploy.unzipqtandroid")

					utils.RemoveAll(filepath.Join(depPath, "build", "libs", "QtAndroid-bundled.jar"))
					utils.RemoveAll(filepath.Join(depPath, "build", "libs", "org", "qtproject", "qt5", "android", "QtNative.class"))
					for i := 1; i < 19; i++ {
						utils.RemoveAll(filepath.Join(depPath, "build", "libs", "org", "qtproject", "qt5", "android", fmt.Sprintf("QtNative$%v.class", i)))
					}

					var zipQtAndroid = exec.Command(jarPath, "-cmf", filepath.Join(depPath, "build", "libs", "META-INF", "MANIFEST.MF"), filepath.Join(depPath, "build", "libs", "QtAndroid-bundled.jar"), "org")
					zipQtAndroid.Dir = filepath.Join(depPath, "build", "libs")
					runCmd(zipQtAndroid, "predeploy.zipqtandroid")

					utils.RemoveAll(filepath.Join(depPath, "build", "libs", "org"))
					utils.RemoveAll(filepath.Join(depPath, "build", "libs", "META-INF"))
				}
			}

			utils.MakeFolder(filepath.Join(depPath, "build", "src", "org", "qtproject", "qt5", "android"))
			runCmd(exec.Command(copyCmd, utils.GetQtPkgPath("internal", "android", "patch", "QtNative.java"), filepath.Join(depPath, "build", "src", "org", "qtproject", "qt5", "android")), "predeploy.cpqtandroidpatched")

			var out, err = json.Marshal(&struct {
				Qt                            string `json:"qt"`
				Sdk                           string `json:"sdk"`
				SdkBuildToolsRevision         string `json:"sdkBuildToolsRevision"`
				Ndk                           string `json:"ndk"`
				Toolchainprefix               string `json:"toolchain-prefix"`
				Toolprefix                    string `json:"tool-prefix"`
				Toolchainversion              string `json:"toolchain-version"`
				Ndkhost                       string `json:"ndk-host"`
				Targetarchitecture            string `json:"target-architecture"`
				AndroidPackageSourceDirectory string `json:"android-package-source-directory"`
				Qmlrootpath                   string `json:"qml-root-path"`
				Applicationbinary             string `json:"application-binary"`
			}{
				Qt:  filepath.Join(qtPrefix, "Qt5.5.1", "5.5", "android_armv7"),
				Sdk: filepath.Join(androidPrefix, "android-sdk"),
				SdkBuildToolsRevision: "23.0.2",
				Ndk:                           filepath.Join(androidPrefix, "android-ndk"),
				Toolchainprefix:               "arm-linux-androideabi",
				Toolprefix:                    "arm-linux-androideabi",
				Toolchainversion:              "4.9",
				Ndkhost:                       ndkhost,
				Targetarchitecture:            "armeabi-v7a",
				AndroidPackageSourceDirectory: filepath.Join(appPath, "android"),
				Qmlrootpath:                   filepath.Join(appPath, "qml"),
				Applicationbinary:             filepath.Join(depPath, "libgo.so"),
			})
			if err != nil {
				fmt.Println("predeploy.json", string(out), err)
				os.Exit(1)
			}

			utils.Save(filepath.Join(depPath, "android-libgo.so-deployment-settings.json"), string(out))
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					utils.Save(filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_sh", appName, appName)), darwinSH())
					utils.Save(filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/Info.plist", appName)), darwinPLIST())
					//TODO: icon + plist
				}

			case "linux":
				{
					utils.Save(filepath.Join(depPath, fmt.Sprintf("%v.sh", appName)), linuxSH())
				}

			case "windows":
				{
					//TODO: icon windres
				}
			}
		}
	}
}

func deploy() {

	switch buildTarget {
	case "android":
		{

			var (
				jdkLib        string
				qtPrefix      string
				androidPrefix string
				ending        string
			)

			switch runtime.GOOS {
			case "darwin", "linux":
				{

					if runtime.GOOS == "darwin" {
						var version = strings.Split(runCmd(exec.Command("java", "-version"), "deploy.jdk"), "\"")[1]
						jdkLib = fmt.Sprintf("/Library/Java/JavaVirtualMachines/jdk%v.jdk/Contents/Home", version)
					} else {
						jdkLib = "/opt/jdk"
					}

					qtPrefix = "/usr/local"
					androidPrefix = "/opt"
				}
			case "windows":
				{
					var version = strings.Split(runCmd(exec.Command("java", "-version"), "deploy.jdk"), "\"")[1]
					jdkLib = fmt.Sprintf("C:\\Program Files\\Java\\jdk%v", version)

					qtPrefix = "C:\\Qt"
					androidPrefix = "C:\\android"
					ending = ".exe"
				}
			}

			var deploy = exec.Command(filepath.Join(qtPrefix, "Qt5.5.1", "5.5", "android_armv7", "bin", "androiddeployqt"+ending))
			deploy.Args = append(deploy.Args,
				"--input", filepath.Join(depPath, "android-libgo.so-deployment-settings.json"),
				"--output", filepath.Join(depPath, "build"),
				"--deployment", "bundled",
				"--android-platform", "android-22",
				"--jdk", jdkLib,
				"--ant", filepath.Join(androidPrefix, "apache-ant", "bin", "ant"),
			)

			if ks := utils.Load(filepath.Join(appPath, "android", appName+".keystore")); ks != "" {
				deploy.Args = append(deploy.Args,
					"--sign", filepath.Join(appPath, "android", appName+".keystore"),
					strings.TrimSpace(utils.Load(filepath.Join(appPath, "android", "alias.txt"))),
					"--storepass", strings.TrimSpace(utils.Load(filepath.Join(appPath, "android", "password.txt"))),
				)
			}

			deploy.Dir = filepath.Join(qtPrefix, "Qt5.5.1", "5.5", "android_armv7", "bin")
			deploy.Env = append(deploy.Env, "JAVA_HOME="+jdkLib)

			if runtime.GOOS == "windows" {
				utils.Save(filepath.Join(depPath, "build.bat"), fmt.Sprintf("set JAVA_HOME=%v\r\n%v", jdkLib, strings.Join(deploy.Args, " ")))
				runCmd(exec.Command(filepath.Join(depPath, "build.bat")), "deploy")
				utils.RemoveAll(filepath.Join(depPath, "build.bat"))
			} else {
				runCmd(deploy, "deploy")
			}
		}
	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					var deploy = exec.Command("/usr/local/Qt5.5.1/5.5/clang_64/bin/macdeployqt")
					deploy.Args = append(deploy.Args,
						filepath.Join(depPath, fmt.Sprintf("%v.app/", appName)),
						fmt.Sprintf("-qmldir=%v", filepath.Join(appPath, "qml")),
						"-always-overwrite")
					deploy.Dir = "/usr/local/Qt5.5.1/5.5/clang_64/bin/"
					runCmd(deploy, "deploy")
				}

			case "linux":
				{
					var libraryPath string

					for _, dep := range strings.Split(runCmd(exec.Command("ldd", filepath.Join(depPath, appName)), "deploy.ldd"), "\n") {
						if strings.Contains(dep, "libQt5") || strings.Contains(dep, "libicu") {
							var libraryP, libName = filepath.Split(strings.Split(dep, " ")[2])
							libraryPath = libraryP
							runCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, libName), filepath.Join(depPath, libName)), fmt.Sprintf("deploy.%v", libName))
						}
					}

					for _, libName := range []string{"DBus", "XcbQpa", "Quick", "Widgets"} {
						runCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, fmt.Sprintf("libQt5%v.so.5", libName)), filepath.Join(depPath, fmt.Sprintf("libQt5%v.so.5", libName))), fmt.Sprintf("deploy.%v", libName))
					}

					libraryPath = strings.TrimSuffix(libraryPath, "lib/")

					for _, libDir := range []string{"platforms", "platformthemes", "xcbglintegrations"} {
						utils.MakeFolder(filepath.Join(depPath, libDir))
					}

					runCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "qml/"), depPath), "deploy.qml")
					runCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, "plugins", "platforms", "libqxcb.so"), filepath.Join(depPath, "platforms", "libqxcb.so")), "deploy.qxcb")
					runCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, "plugins", "platformthemes", "libqgtk2.so"), filepath.Join(depPath, "platformthemes", "libqgtk2.so")), "deploy.qgtk2")
					runCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, "plugins", "xcbglintegrations", "libqxcb-glx-integration.so"), filepath.Join(depPath, "xcbglintegrations", "libqxcb-glx-integration.so")), "deploy.qxcb-glx-integration")
				}

			case "windows":
				{
					var deploy = exec.Command("C:\\Qt\\Qt5.5.1\\5.5\\mingw492_32\\bin\\windeployqt.exe")
					deploy.Args = append(deploy.Args,
						filepath.Join(depPath, appName+ending),
						fmt.Sprintf("-qmldir=%v", filepath.Join(appPath, "qml")),
						"-force")
					runCmd(deploy, "deploy")
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

			if ks := utils.Load(filepath.Join(appPath, "android", appName+".keystore")); ks != "" {
				runCmd(exec.Command(copyCmd, filepath.Join(depPath, "build", "bin", "QtApp-release-signed.apk"), filepath.Join(depPath, fmt.Sprintf("%v.%v", appName, apkEnding))), "pastdeploy.release")
			} else {
				runCmd(exec.Command(copyCmd, filepath.Join(depPath, "build", "bin", "QtApp-debug.apk"), filepath.Join(depPath, fmt.Sprintf("%v.%v", appName, apkEnding))), "pastdeploy.debug")
			}

			//TODO: copy manifest to android folder and change mindSdkVersion >= 16

		}
	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					runCmd(exec.Command("mv", filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName)), filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_app", appName, appName))), "pastdeploy.moveApp")
					runCmd(exec.Command("mv", filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_sh", appName, appName)), filepath.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName))), "pastdeploy.moveSh")
				}
			}
		}
	}
}

func cleanup() {
	var (
		qmlGo  = filepath.Join(appPath, "qtQrc.go")
		qmlQrc = filepath.Join(appPath, "qtQrc.qrc")
		qmlCpp = filepath.Join(appPath, "qtQrc.cpp")
	)

	utils.RemoveAll(qmlGo)
	utils.RemoveAll(qmlQrc)
	utils.RemoveAll(qmlCpp)
}

func run() {

	switch buildTarget {
	case "android":
		{
			switch runtime.GOOS {
			case "darwin", "linux":
				runCmdOptional(exec.Command("killall", "adb"), "run.killadb")
				//runCmdOptional(exec.Command("/opt/android-sdk/platform-tools/adb", "logcat", "-c"), "run.adblogcat")
				exec.Command("/opt/android-sdk/platform-tools/adb", "install", "-r", filepath.Join(depPath, fmt.Sprintf("%v.apk", appName))).Start()

			case "windows":
				{
					exec.Command("C:\\android\\android-sdk\\platform-tools\\adb.exe", "install", "-r", filepath.Join(depPath, fmt.Sprintf("%v.apk", appName))).Start()
				}
			}
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					runCmdOptional(exec.Command("open", filepath.Join(depPath, fmt.Sprintf("%v.app/", appName))), "run")
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
	}
}

func runCmd(cmd *exec.Cmd, n string) string {
	var out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", n, out, err)
		cleanup()
		os.Exit(1)
	}
	return string(out)
}

func runCmdOptional(cmd *exec.Cmd, n string) string {
	var out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n\n%v\noutput:%s\n\n", n, out)
	}
	return string(out)
}

//darwin
func darwinSH() string {
	var o = "#!/bin/bash\n"
	o += "cd \"${0%/*}\"\n"
	o += fmt.Sprintf("./%v_app", appName)
	return o
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

	<key>NSHighResolutionCapable</key>
	<string>True</string>
</dict>
</plist>
`, appName, appName, appName, appName)
}

//linux
func linuxSH() string {

	var o = "#!/bin/sh\n"
	o += "appname=`basename $0 | sed s,\\.sh$,,`\n\n"
	o += "dirname=`dirname $0`\n"
	o += "tmp=\"${dirname#?}\"\n\n"
	o += "if [ \"${dirname%$tmp}\" != \"/\" ]; then\n"
	o += "dirname=$PWD/$dirname\n"
	o += "fi\n"

	o += "LD_LIBRARY_PATH=$dirname\n"
	o += "export LD_LIBRARY_PATH\n"

	o += "QML_IMPORT_PATH=$dirname/\"qml\"\n"
	o += "export QML_IMPORT_PATH\n"

	o += "QML2_IMPORT_PATH=$dirname/\"qml\"\n"
	o += "export QML2_IMPORT_PATH\n"

	o += "$dirname/$appname \"$@\"\n"

	return o
}
