package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

var (
	appPath, appName       string
	depPath                string
	buildMode, buildTarget string
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

	if !path.IsAbs(appPath) {
		appPath = utils.GetAbsPath(appPath)
	}
	appName = path.Base(appPath)

	switch buildTarget {
	case "android":
		{
			depPath = path.Join(appPath, "deploy", buildTarget)
		}

	case "desktop":
		{
			depPath = path.Join(appPath, "deploy", runtime.GOOS)
		}
	}

	switch buildMode {
	case "build", "test":
		{
			utils.RemoveAll(depPath)
			utils.MakeFolder(depPath)
		}
	}
}

func qrc() {

	utils.MakeFolder(path.Join(appPath, "qml"))

	var (
		rccPath string
		qmlGo   = path.Join(appPath, "qtQrc.go")
		qmlQrc  = path.Join(appPath, "qtQrc.qrc")
		qmlCpp  = path.Join(appPath, "qtQrc.cpp")
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
					rccPath = "rcc" //TODO:
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
					rccPath = "rcc"
				}
			}
		}
	}

	utils.Save(qmlGo, `package main

  /*
  #cgo +build android,arm LDFLAGS: -L/usr/local/Qt5.5.1/5.5/android_armv7/lib -lQt5Core

  #cgo +build darwin,amd64 LDFLAGS: -F/usr/local/Qt5.5.1/5.5/clang_64/lib -framework QtCore

  #cgo +build linux,amd64 LDFLAGS: -Wl,-rpath,/usr/local/Qt5.5.1/5.5/gcc_64/lib
  #cgo +build linux,amd64 LDFLAGS: -L/usr/local/Qt5.5.1/5.5/gcc_64/lib -lQt5Core

	#cgo +build linux,386 LDFLAGS: -Wl,-rpath,/usr/local/Qt5.5.1/5.5/gcc/lib
  #cgo +build linux,386 LDFLAGS: -L/usr/local/Qt5.5.1/5.5/gcc/lib -lQt5Core

  #cgo +build windows,386 LDFLAGS: -LC:/Qt/Qt5.5.1/5.5/mingw492_32/bin -lQt5Core
  */
  import "C"`)

	var rcc = exec.Command(rccPath)
	rcc.Args = append(rcc.Args,
		"-project",
		"-o", qmlQrc)
	rcc.Dir = path.Join(appPath, "qml")
	runCmd(rcc, "qrc.qrc")

	utils.Save(qmlQrc, strings.Replace(utils.Load(qmlQrc), "<file>./", "<file>qml/", -1))

	rcc = exec.Command(rccPath)
	rcc.Args = append(rcc.Args,
		"-name", appName,
		"-o", qmlCpp,
		qmlQrc)
	runCmd(rcc, "qrc.cpp")
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
			outputFile = path.Join(depPath, fmt.Sprintf("lib%v.so", appName))

			switch runtime.GOOS {
			case "darwin", "linux":
				{
					env = map[string]string{
						"GOPATH":       os.Getenv("GOPATH"),
						"GOOS":         "android",
						"GOARCH":       "arm",
						"GOARM":        "7",
						"CC":           path.Join("/opt", "android-ndk", "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-gcc"),
						"CXX":          path.Join("/opt", "android-ndk", "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-g++"),
						"CGO_ENABLED":  "1",
						"CGO_CPPFLAGS": "-isystem /opt/android-ndk/platforms/android-9/arch-arm/usr/include",
						"CGO_LDFLAGS":  "--sysroot=/opt/android-ndk/platforms/android-9/arch-arm/ -llog",
					}
				}

			case "windows":
				{
					//TODO:
				}
			}
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					ldFlags = "-ldflags=\"-w\" \"-r=/usr/local/Qt5.5.1/5.5/clang_64/lib\""
					outputFile = path.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName))
				}

			case "linux":
				{
					ldFlags = "-ldflags=\"-s\" \"-w\""
					outputFile = path.Join(depPath, appName)
				}

			case "windows":
				{
					ldFlags = "-ldflags=\"-s\" \"-w\" \"-H=windowsgui\""
					outputFile = path.Join(depPath, appName)
				}
			}
		}
	}

	var cmd = exec.Command("go")
	cmd.Args = append(cmd.Args,
		"build",
		"-p", strconv.Itoa(runtime.NumCPU()),
		ldFlags,
		"-o", outputFile)

	cmd.Dir = appPath

	if buildTarget != "desktop" {
		cmd.Args = append(cmd.Args, "-buildmode", "c-shared")
		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}
	}
	runCmd(cmd, "build")
}

func predeploy() {

	switch buildTarget {
	case "android":
		{
			utils.MakeFolder(path.Join(appPath, "android"))

			switch runtime.GOOS {
			case "darwin", "linux":
				{

					for _, dir := range []string{"drawable-hdpi", "drawable-ldpi", "drawable-mdpi"} {
						utils.MakeFolder(path.Join(depPath, "build", "res", dir))
						runCmdOptional(exec.Command("cp", path.Join(appPath, "android", "icon.png"), path.Join(depPath, "build", "res", dir, "icon.png")), "predeploy.cpicon")
					}

					var libPath = path.Join(depPath, "build", "libs", "armeabi-v7a")
					utils.MakeFolder(libPath)
					runCmd(exec.Command("cp", path.Join(depPath, fmt.Sprintf("lib%v.so", appName)), libPath), "predeploy.cplib")

					runCmd(exec.Command("cp", "/usr/local/Qt5.5.1/5.5/android_armv7/jar/QtAndroid-bundled.jar", path.Join(depPath, "build", "libs")), "predeploy.cpqtandroid")

					runCmd(exec.Command("zip", "-d", path.Join(depPath, "build", "libs", "QtAndroid-bundled.jar"), "org/qtproject/qt5/android/QtNative.class"), "predeploy.patchqtandroid_main")
					for i := 1; i < 19; i++ {
						runCmd(exec.Command("zip", "-d", path.Join(depPath, "build", "libs", "QtAndroid-bundled.jar"), fmt.Sprintf("org/qtproject/qt5/android/QtNative$%v.class", i)), fmt.Sprintf("predeploy.patchqtandroid_%v", i))
					}

					utils.MakeFolder(path.Join(depPath, "build", "src", "org", "qtproject", "qt5", "android"))
					runCmd(exec.Command("cp", utils.GetQtPkgPath("internal", "android", "patch", "QtNative.java"), path.Join(depPath, "build", "src", "org", "qtproject", "qt5", "android")), "predeploy.cpqtandroidpatched")

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
						Qt:  "/usr/local/Qt5.5.1/5.5/android_armv7",
						Sdk: "/opt/android-sdk",
						SdkBuildToolsRevision: "23.0.1",
						Ndk:                           "/opt/android-ndk",
						Toolchainprefix:               "arm-linux-androideabi",
						Toolprefix:                    "arm-linux-androideabi",
						Toolchainversion:              "4.9",
						Ndkhost:                       runtime.GOOS + "-x86_64",
						Targetarchitecture:            "armeabi-v7a",
						AndroidPackageSourceDirectory: path.Join(appPath, "android"),
						Qmlrootpath:                   path.Join(appPath, "qml"),
						Applicationbinary:             path.Join(depPath, fmt.Sprintf("lib%v.so", appName)),
					})
					if err != nil {
						fmt.Println("predeploy.json", string(out), err)
						os.Exit(1)
					}

					utils.Save(path.Join(depPath, fmt.Sprintf("android-lib%v.so-deployment-settings.json", appName)), string(out))
				}

			case "windows":
				{
					//TODO:
				}
			}
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					utils.Save(path.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_sh", appName, appName)), darwinSH())
					utils.Save(path.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/Info.plist", appName)), darwinPLIST())
					//TODO: icon
				}

			case "linux":
				{
					utils.Save(path.Join(depPath, fmt.Sprintf("%v.sh", appName)), linuxSH())
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
			switch runtime.GOOS {
			case "darwin", "linux":
				{

					var jdkLib string
					switch runtime.GOOS {
					case "darwin":
						jdkLib = "/Library/Java/JavaVirtualMachines/jdk1.8.0_66.jdk/Contents/Home" //TODO: Get on runtime (java -version)
					case "linux":
						jdkLib = "/opt/jdk"
					}

					var deploy = exec.Command("/usr/local/Qt5.5.1/5.5/android_armv7/bin/androiddeployqt")
					deploy.Args = append(deploy.Args,
						"--input", path.Join(depPath, fmt.Sprintf("android-lib%v.so-deployment-settings.json", appName)),
						"--output", path.Join(depPath, "build"),
						"--deployment", "bundled",
						"--android-platform", "android-23",
						"--jdk", jdkLib,
						"--ant", "/opt/apache-ant/bin/ant")

					if ks := utils.Load(path.Join(appPath, "android", appName+".keystore")); ks != "" {
						deploy.Args = append(deploy.Args,
							"--sign", path.Join(appPath, "android", appName+".keystore"),
							strings.TrimSpace(utils.Load(path.Join(appPath, "android", "alias.txt"))),
							"--storepass", strings.TrimSpace(utils.Load(path.Join(appPath, "android", "password.txt"))))
					}

					deploy.Dir = "/usr/local/Qt5.5.1/5.5/android_armv7/bin/"
					if runtime.GOOS == "linux" {
						deploy.Env = append(deploy.Env, "JAVA_HOME=/opt/jdk")
					}

					runCmd(deploy, "deploy")
				}

			case "windows":
				{
					//TODO:
				}
			}
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					var deploy = exec.Command("/usr/local/Qt5.5.1/5.5/clang_64/bin/macdeployqt")
					deploy.Args = append(deploy.Args,
						path.Join(depPath, fmt.Sprintf("%v.app/", appName)),
						fmt.Sprintf("-qmldir=%v", path.Join(appPath, "qml")),
						"-always-overwrite")
					deploy.Dir = "/usr/local/Qt5.5.1/5.5/clang_64/bin/"
					runCmd(deploy, "deploy")
				}

			case "linux":
				{
					var libraryPath string

					for _, dep := range strings.Split(runCmd(exec.Command("ldd", path.Join(depPath, appName)), "deploy.ldd"), "\n") {
						if strings.Contains(dep, "libQt5") || strings.Contains(dep, "libicu") {
							var libraryP, libName = path.Split(strings.Split(dep, " ")[2])
							libraryPath = libraryP
							runCmd(exec.Command("cp", "-L", path.Join(libraryPath, libName), path.Join(depPath, libName)), fmt.Sprintf("deploy.%v", libName))
						}
					}

					for _, libName := range []string{"DBus", "XcbQpa", "Quick", "Widgets"} {
						runCmd(exec.Command("cp", "-L", path.Join(libraryPath, fmt.Sprintf("libQt5%v.so.5", libName)), path.Join(depPath, fmt.Sprintf("libQt5%v.so.5", libName))), fmt.Sprintf("deploy.%v", libName))
					}

					libraryPath = strings.TrimSuffix(libraryPath, "lib/")

					for _, libDir := range []string{"platforms", "platformthemes", "xcbglintegrations"} {
						utils.MakeFolder(path.Join(depPath, libDir))
					}

					runCmd(exec.Command("cp", "-R", path.Join(libraryPath, "qml/"), depPath), "deploy.qml")
					runCmd(exec.Command("cp", "-L", path.Join(libraryPath, "plugins", "platforms", "libqxcb.so"), path.Join(depPath, "platforms", "libqxcb.so")), "deploy.qxcb")
					runCmd(exec.Command("cp", "-L", path.Join(libraryPath, "plugins", "platformthemes", "libqgtk2.so"), path.Join(depPath, "platformthemes", "libqgtk2.so")), "deploy.qgtk2")
					runCmd(exec.Command("cp", "-L", path.Join(libraryPath, "plugins", "xcbglintegrations", "libqxcb-glx-integration.so"), path.Join(depPath, "xcbglintegrations", "libqxcb-glx-integration.so")), "deploy.qxcb-glx-integration")
				}

			case "windows":
				{
					var deploy = exec.Command("windeployqt")
					deploy.Args = append(deploy.Args,
						path.Join(depPath, appName),
						fmt.Sprintf("-qmldir=%v", path.Join(appPath, "qml")),
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
			switch runtime.GOOS {
			case "darwin", "linux":
				{
					if ks := utils.Load(path.Join(appPath, "android", appName+".keystore")); ks != "" {
						runCmdOptional(exec.Command("mv", path.Join(depPath, "build", "bin", "QtApp-release-signed.apk"), path.Join(depPath, fmt.Sprintf("%v.apk", appName))), "pastdeploy.release")
					} else {
						runCmdOptional(exec.Command("mv", path.Join(depPath, "build", "bin", "QtApp-debug.apk"), path.Join(depPath, fmt.Sprintf("%v.apk", appName))), "pastdeploy.debug")
					}
				}

			case "windows":
				{
					//TODO:
				}
			}
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					runCmd(exec.Command("mv", path.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName)), path.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_app", appName, appName))), "pastdeploy.moveApp")
					runCmd(exec.Command("mv", path.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_sh", appName, appName)), path.Join(depPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName))), "pastdeploy.moveSh")
				}
			}
		}
	}
}

func cleanup() {
	var (
		qmlGo  = path.Join(appPath, "qtQrc.go")
		qmlQrc = path.Join(appPath, "qtQrc.qrc")
		qmlCpp = path.Join(appPath, "qtQrc.cpp")
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
				exec.Command("/opt/android-sdk/platform-tools/adb", "install", "-r", path.Join(depPath, fmt.Sprintf("%v.apk", appName))).Start() //TODO:

			case "windows":
				{
					//TODO:
				}
			}
		}

	case "desktop":
		{
			switch runtime.GOOS {
			case "darwin":
				{
					runCmdOptional(exec.Command("open", path.Join(depPath, fmt.Sprintf("%v.app/", appName))), "run")
				}

			case "linux":
				{
					exec.Command(path.Join(depPath, fmt.Sprintf("%v.sh", appName))).Start() //TODO:
				}

			case "windows":
				{
					runCmdOptional(exec.Command(path.Join(depPath, appName)), "run")
				}
			}
		}
	}
}

func runCmd(cmd *exec.Cmd, n string) string {
	var out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("=>%v<=\noutput:%s\nerror:%s\n\n", n, out, err)
		cleanup()
		os.Exit(1)
	}
	return string(out)
}

func runCmdOptional(cmd *exec.Cmd, n string) string {
	var out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("=>%v<=\noutput:%s\n\n", n, out)
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
