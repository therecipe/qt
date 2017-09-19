package deploy

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/utils"
)

func bundle(mode, target, path, name, depPath string) {
	copy := func(src, dst string) {
		copy := "cp"

		if runtime.GOOS == "windows" {
			copy = "xcopy"

			//TODO: -->
			src = strings.TrimSuffix(src, "/")
			src = strings.TrimSuffix(src, "\\")

			dst = strings.TrimSuffix(dst, "/")
			dst = strings.TrimSuffix(dst, "\\")
			//<--
		}

		var args []string
		if _, err := ioutil.ReadDir(src); err == nil {
			if runtime.GOOS != "windows" {
				args = append(args, "-R")
			}
		}

		var suffix string
		if _, err := ioutil.ReadDir(dst); err != nil {
			if runtime.GOOS == "windows" {
				suffix = "*"
			}
		}

		utils.RunCmd(exec.Command(copy, append(args, src, dst+suffix)...), fmt.Sprintf("copy %v to %v for %v on %v", filepath.Base(src), filepath.Base(dst), target, runtime.GOOS))
	}

	switch target {
	case "darwin":

		//copy default assets
		utils.Save(filepath.Join(depPath, filepath.Join(name+".app", "Contents", "Info.plist")), darwin_plist(name))
		utils.Save(filepath.Join(depPath, filepath.Join(name+".app", "Contents", "PkgInfo")), darwin_pkginfo())
		utils.MkdirAll(filepath.Join(depPath, filepath.Join(name+".app", "Contents", "Resources")))
		utils.Save(filepath.Join(depPath, filepath.Join(name+".app", "Contents", "Resources", "empty.lproj")), "")

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/", filepath.Join(depPath, name+".app"))

		dep := exec.Command(filepath.Join(utils.QT_DARWIN_DIR(), "bin", "macdeployqt"))
		dep.Args = append(dep.Args, filepath.Join(depPath, name+".app"), "-qmldir="+path)
		dep.Dir = filepath.Join(utils.QT_DARWIN_DIR(), "bin")
		utils.RunCmd(dep, fmt.Sprintf("deploy for %v on %v", target, runtime.GOOS))

	case "linux", "rpi1", "rpi2", "rpi3":

		//copy default assets
		utils.SaveExec(filepath.Join(depPath, fmt.Sprintf("%v.sh", name)), linux_sh(target, name))

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/", depPath)

		//TODO: -->
		{
			if utils.QT_PKG_CONFIG() {
				return
			}

			libDir := "lib"
			if name == libDir {
				libDir = "libs"
			}
			utils.MkdirAll(filepath.Join(depPath, libDir))

			var (
				libraryPath   = strings.TrimSpace(utils.RunCmd(exec.Command(utils.ToolPath("qmake", target), "-query", "QT_INSTALL_LIBS"), fmt.Sprintf("query lib path for %v on %v", target, runtime.GOOS)))
				lddPath       = "ldd"
				lddExtra      string
				lddOutput     string
				usesWebEngine bool
				usesQml       bool
			)

			if strings.HasPrefix(target, "rpi") {
				//libraryPath = fmt.Sprintf("%v/%v/%v/lib/", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), target)
				lddPath = fmt.Sprintf("%v/arm-bcm2708/%v/bin/arm-linux-gnueabihf-ldd", utils.RPI_TOOLS_DIR(), utils.RPI_COMPILER())
				lddExtra = "--root=/"
				lddOutput = utils.RunCmd(exec.Command(lddPath, lddExtra, filepath.Join(depPath, name)), fmt.Sprintf("ldd binary for %v on %v", target, runtime.GOOS))
			} else {
				lddOutput = utils.RunCmd(exec.Command(lddPath, filepath.Join(depPath, name)), fmt.Sprintf("ldd binary for %v on %v", target, runtime.GOOS))
			}

			for _, dep := range strings.Split(lddOutput, "\n") {
				if strings.Contains(dep, "libQt5") || strings.Contains(dep, "libicu") {
					var libName string

					if strings.HasPrefix(target, "rpi") {
						libName = strings.TrimSpace(strings.Replace(strings.Split(dep, "=>")[0], "not found", "", -1))
					} else {
						if libraryPath == "" {
							libraryPath, libName = filepath.Split(strings.Split(dep, " ")[2])
						} else {
							_, libName = filepath.Split(strings.Split(dep, " ")[2])
						}
					}

					if utils.ExistsFile(filepath.Join(libraryPath, libName)) {
						utils.RunCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, libName), filepath.Join(depPath, libDir, libName)), fmt.Sprintf("copy %v for %v on %v", libName, target, runtime.GOOS))
					}

					if strings.Contains(dep, "WebEngine") {
						usesWebEngine = true
					}
					if strings.Contains(dep, "Quick") || strings.Contains(dep, "Qml") {
						usesQml = true
					}
				}
			}

			libs := []string{"DBus", "XcbQpa", "Quick", "Widgets", "EglDeviceIntegration", "EglFsKmsSupport", "OpenGL", "WaylandClient", "WaylandCompositor", "QuickControls2", "QuickTemplates2", "QuickWidgets", "QuickParticles", "CLucene", "Concurrent", "Svg"}
			if usesQml {
				libs = append(libs, []string{"3DCore", "3DExtras", "3DInput", "3DLogic", "3DQuick", "3DQuickExtras", "3DQuickInput", "3DQuickRender", "3DRender", "Gamepad"}...)
			}
			for _, libName := range libs {
				if utils.ExistsFile(filepath.Join(libraryPath, fmt.Sprintf("libQt5%v.so.5", libName))) {
					utils.RunCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, fmt.Sprintf("libQt5%v.so.5", libName)), filepath.Join(depPath, libDir, fmt.Sprintf("libQt5%v.so.5", libName))), fmt.Sprintf("copy %v for %v on %v", libName, target, runtime.GOOS))
				}
			}
			if utils.ExistsFile(filepath.Join(libraryPath, "libqgsttools_p.so.1.0.0")) {
				utils.RunCmd(exec.Command("cp", "-L", filepath.Join(libraryPath, "libqgsttools_p.so.1.0.0"), filepath.Join(depPath, libDir, "libqgsttools_p.so.1")), fmt.Sprintf("copy libqgsttools_p.so.1 for %v on %v", target, runtime.GOOS))
			}

			libraryPath = filepath.Dir(libraryPath)
			utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "qml/"), depPath), fmt.Sprintf("copy qml dir for %v on %v", target, runtime.GOOS))
			utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "plugins/"), depPath), fmt.Sprintf("copy plugins dir for %v on %v", target, runtime.GOOS))

			if usesWebEngine {
				utils.RunCmd(exec.Command("cp", filepath.Join(libraryPath, "libexec", "QtWebEngineProcess"), depPath), fmt.Sprintf("copy QtWebEngineProcess for %v on %v", target, runtime.GOOS))
				var fileList, err = ioutil.ReadDir(filepath.Join(libraryPath, "resources"))
				if err != nil {
					utils.Log.WithError(err).Error("failed to read resource folder")
				}
				for _, file := range fileList {
					utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "resources", file.Name()), depPath), fmt.Sprintf("copy resource %v for %v on %v", file.Name(), target, runtime.GOOS))
				}
				utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "translations/qtwebengine_locales/"), depPath), fmt.Sprintf("copy qtwebengine_locales dir for %v on %v", target, runtime.GOOS))
			}
		}
		//<--

	case "windows":

		//TODO: -->
		switch {
		case runtime.GOOS != target:
			if utils.QT_MXE_STATIC() {
				return
			}

			var libraryPath = filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "bin")
			for _, d := range []string{"libbz2", "libfreetype-6", "libglib-2.0-0", "libharfbuzz-0", "libiconv-2", "libintl-8", "libpcre-1", "libpcre16-0", "libpng16-16", "libstdc++-6", "libwinpthread-1", "zlib1", "libeay32", "ssleay32"} {
				utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, target, runtime.GOOS))
			}
			for _, d := range []string{"libjasper-1", "libjpeg-9", "libmng-2", "libtiff-5", "libwebp-5", "liblcms2-2", "liblzma-5", "libwebpdemux-1"} {
				utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, target, runtime.GOOS))
			}

			var gccDep = "libgcc_s_sjlj-1"
			if utils.QT_MXE_ARCH() == "amd64" {
				gccDep = "libgcc_s_seh-1"
			}
			utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", gccDep)), depPath), fmt.Sprintf("copy %v for %v on %v", gccDep, target, runtime.GOOS))

			libraryPath = filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "qt5")
			utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "qml/")+"/.", depPath), fmt.Sprintf("copy qml dir for %v on %v", target, runtime.GOOS))
			utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "plugins/")+"/.", depPath), fmt.Sprintf("copy plugins dir for %v on %v", target, runtime.GOOS))

			libraryPath = filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "qt5", "bin")
			var output = utils.RunCmd(exec.Command(utils.QT_MXE_BIN("objdump"), "-x", filepath.Join(depPath, name+".exe")), fmt.Sprintf("objdump binary for %v on %v", target, runtime.GOOS))
			for lib, deps := range parser.LibDeps {
				if strings.Contains(output, lib) && lib != parser.MOC {
					for _, lib := range append(deps, lib) {
						utils.RunCmd(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("Qt5%v.dll", lib)), depPath), fmt.Sprintf("copy %v for %v on %v", lib, target, runtime.GOOS))
					}
				}
			}
			for _, d := range []string{"Qt5OpenGL", "Qt5Quick", "Qt5QuickControls2", "Qt5QuickTemplates2"} {
				utils.RunCmdOptional(exec.Command("cp", filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, target, runtime.GOOS))
			}

		case utils.QT_MSYS2():
			if utils.QT_MSYS2_STATIC() {
				return
			}

			var copyCmd = "xcopy"
			if utils.MSYSTEM() != "" {
				copyCmd = "cp"
			}

			var deploy = exec.Command(filepath.Join(utils.QT_MSYS2_DIR(), "bin", "windeployqt"))
			deploy.Args = append(deploy.Args,
				filepath.Join(depPath, name+".exe"),
				fmt.Sprintf("-qmldir=%v", path),
				"-force")
			utils.RunCmd(deploy, fmt.Sprintf("depoy %v on %v", target, runtime.GOOS))

			var libraryPath = filepath.Join(utils.QT_MSYS2_DIR(), "bin")
			for _, d := range []string{"libbz2-1", "libfreetype-6", "libglib-2.0-0", "libharfbuzz-0", "libiconv-2", "libintl-8", "libpcre-1", "libpcre16-0", "libpng16-16", "libstdc++-6", "libwinpthread-1", "zlib1", "libgraphite2", "libicudt58", "libicuin58", "libicuuc58"} {
				utils.RunCmdOptional(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("%v.dll", d)), depPath), fmt.Sprintf("copy %v for %v on %v", d, target, runtime.GOOS))
			}

			var gccDep = "libgcc_s_dw2-1"
			if utils.QT_MSYS2_ARCH() == "amd64" {
				gccDep = "libgcc_s_seh-1"
			}

			utils.RunCmdOptional(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("%v.dll", gccDep)), depPath), fmt.Sprintf("copy %v for %v on %v", gccDep, target, runtime.GOOS))

			libraryPath = filepath.Join(utils.QT_MSYS2_DIR(), "share", "qt5")
			if utils.MSYSTEM() != "" {
				utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "qml/")+"/.", depPath), fmt.Sprintf("copy qml dir for %v on %v", target, runtime.GOOS))
				utils.RunCmd(exec.Command("cp", "-R", filepath.Join(libraryPath, "plugins/")+"/.", depPath), fmt.Sprintf("copy plugins dir for %v on %v", target, runtime.GOOS))
			} else {
				utils.RunCmd(exec.Command("xcopy", "/S", "/Y", filepath.Join(libraryPath, "qml/"), depPath), fmt.Sprintf("copy qml dir for %v on %v", target, runtime.GOOS))
				utils.RunCmd(exec.Command("xcopy", "/S", "/Y", filepath.Join(libraryPath, "plugins/"), depPath), fmt.Sprintf("copy plugins dir for %v on %v", target, runtime.GOOS))
			}

			libraryPath = filepath.Join(utils.QT_MSYS2_DIR(), "bin")
			var output = utils.RunCmd(exec.Command(filepath.Join(utils.QT_MSYS2_DIR(), "bin", "objdump"), "-x", filepath.Join(depPath, name+".exe")), fmt.Sprintf("objdump binary for %v on %v", target, runtime.GOOS))
			for lib, deps := range parser.LibDeps {
				if strings.Contains(output, lib) && lib != parser.MOC {
					for _, lib := range append(deps, lib) {
						if utils.ExistsFile(filepath.Join(libraryPath, fmt.Sprintf("Qt5%v.dll", lib))) {
							utils.RunCmd(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("Qt5%v.dll", lib)), depPath), fmt.Sprintf("copy %v for %v on %v", lib, target, runtime.GOOS))
						}
					}
				}
			}

			deps := []string{"Qt5OpenGL"}
			if utils.QT_WEBKIT() {
				deps = append(deps, []string{"libjpeg-8", "libsqlite3-0", "libwebp-7", "libxml2-2", "liblzma-5", "libxslt-1"}...)
			}
			for _, lib := range deps {
				utils.RunCmd(exec.Command(copyCmd, filepath.Join(libraryPath, fmt.Sprintf("%v.dll", lib)), depPath), fmt.Sprintf("copy %v for %v on %v", lib, target, runtime.GOOS))
			}

			var walkFn = func(path string, info os.FileInfo, err error) error {
				if strings.HasSuffix(info.Name(), "d.dll") {
					utils.RemoveAll(path)
				}
				return nil
			}
			filepath.Walk(depPath, walkFn)

		default:

			//copy default assets
			//TODO: windres icon

			//copy custom assets
			assets := filepath.Join(path, target)
			utils.MkdirAll(assets)
			copy(assets+"/", depPath)

			if utils.QT_WEBKIT() {
				libraryPath := filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "mingw53_32", "bin")
				output := utils.RunCmd(exec.Command(filepath.Join("objdump"), "-x", filepath.Join(depPath, name+".exe")), fmt.Sprintf("objdump binary for %v on %v", target, runtime.GOOS))
				for lib, deps := range parser.LibDeps {
					if strings.Contains(output, lib) && lib == "WebKit" {
						for _, lib := range append(deps, lib) {
							for _, pref := range []string{"lib", ""} {
								libName := filepath.Join(libraryPath, fmt.Sprintf("%vQt5%v.dll", pref, lib))
								if utils.ExistsFile(libName) {
									copy(libName, depPath)
								}
							}
						}
					}
				}

				for _, lib := range []string{"icudt57", "icuin57", "icuuc57", "libxml2-2", "libxslt-1", "Qt5MultimediaWidgets", "Qt5OpenGL", "Qt5PrintSupport"} {
					copy(filepath.Join(libraryPath, lib+".dll"), depPath)
				}
			}

			dep := exec.Command(filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "mingw53_32", "bin", "windeployqt"))
			dep.Args = append(dep.Args, filepath.Join(depPath, name+".exe"), "-qmldir="+path)
			utils.RunCmd(dep, fmt.Sprintf("deploy for %v on %v", target, runtime.GOOS))
		}
		//<--

	case "android", "android-emulator":

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/", filepath.Join(depPath, "build"))

		//create default gradle-wrapper.properties
		if target == "android-emulator" {
			path := filepath.Join(depPath, "build", "gradle", "wrapper")
			utils.MkdirAll(path)
			utils.Save(filepath.Join(path, "gradle-wrapper.properties"), android_gradle_wrapper())
		}

		//wrap exported go main inside c main
		env, _, _, _ := cmd.BuildEnv(target, name, depPath)
		compiler := env["CXX"]

		wrapper := filepath.Join(depPath, "c_main_wrapper.cpp")
		utils.Save(wrapper, "#include \"libgo_base.h\"\nint main(int argc, char *argv[]) { go_main_wrapper(); }")
		cmd := exec.Command(compiler, "c_main_wrapper.cpp", "-o", filepath.Join(depPath, "libgo.so"), "-I../..", "-L.", "-lgo_base", "--sysroot="+filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-16", "arch-arm"), "-shared")
		if target == "android-emulator" {
			cmd = exec.Command(compiler, "c_main_wrapper.cpp", "-o", filepath.Join(depPath, "libgo.so"), "-I../..", "-L.", "-lgo_base", "--sysroot="+filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-16", "arch-x86"), "-shared")
		}
		cmd.Dir = depPath
		utils.RunCmd(cmd, fmt.Sprintf("compile wrapper for %v on %v", target, runtime.GOOS))
		utils.RemoveAll(wrapper)

		strip := exec.Command(filepath.Join(filepath.Dir(compiler), "arm-linux-androideabi-strip"), "libgo.so")
		if target == "android-emulator" {
			strip = exec.Command(filepath.Join(filepath.Dir(compiler), "i686-linux-android-strip"), "libgo.so")
		}
		strip.Dir = depPath
		utils.RunCmd(strip, fmt.Sprintf("strip binary for %v on %v", target, runtime.GOOS))

		libPath := filepath.Join(depPath, "build", "libs", "armeabi-v7a")
		if target == "android-emulator" {
			libPath = filepath.Join(depPath, "build", "libs", "x86")
		}
		utils.MkdirAll(libPath)

		if utils.QT_VAGRANT() {
			libPath = strings.Replace(libPath, "C:\\media\\sf_GOPATH", "C:\\media\\UNC\\vboxsrv\\media_sf_GOPATH", -1)
			utils.RemoveAll(libPath)
			utils.MkdirAll(libPath)
			copy(filepath.Join(depPath, "libgo.so"), filepath.Join(libPath, "libgo.so"))
			copy(filepath.Join(depPath, "libgo_base.so"), filepath.Join(libPath, "libgo_base.so"))
			utils.RemoveAll(filepath.Join(depPath, "libgo.so"))
			utils.RemoveAll(filepath.Join(depPath, "libgo_base.so"))
		} else {
			os.Rename(filepath.Join(depPath, "libgo.so"), filepath.Join(libPath, "libgo.so"))
			os.Rename(filepath.Join(depPath, "libgo_base.so"), filepath.Join(libPath, "libgo_base.so"))
		}

		//trick androiddeployqt into checking dependencies from libgo_base.so
		copy(filepath.Join(libPath, "libgo_base.so"), depPath)
		copy(filepath.Join(libPath, "libgo_base.so"), filepath.Join(depPath, "libgo.so"))

		utils.Save(filepath.Join(depPath, "android-libgo.so-deployment-settings.json"), android_config(target, path, depPath))

		dep := exec.Command(filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "android_armv7", "bin", "androiddeployqt"))
		dep.Dir = filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "android_armv7", "bin")
		if target == "android-emulator" {
			dep := exec.Command(filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "android_x86", "bin", "androiddeployqt"))
			dep.Dir = filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "android_x86", "bin")
		}
		dep.Env = append(dep.Env, "JAVA_HOME="+utils.JDK_DIR())
		dep.Args = append(dep.Args,
			"--input", filepath.Join(depPath, "android-libgo.so-deployment-settings.json"),
			"--output", filepath.Join(depPath, "build"),
			"--deployment", "bundled",
			"--android-platform", "android-25",
			"--jdk", utils.JDK_DIR(),
			"--gradle",
			"--verbose",
		)
		if utils.ExistsFile(filepath.Join(path, target, name+".keystore")) {
			dep.Args = append(dep.Args,
				"--sign", filepath.Join(path, target, name+".keystore"),
				strings.TrimSpace(utils.Load(filepath.Join(path, target, "alias.txt"))),
				"--storepass", strings.TrimSpace(utils.Load(filepath.Join(path, target, "password.txt"))))
		}

		if runtime.GOOS == "windows" {
			//TODO: -->
			utils.SaveExec(filepath.Join(depPath, "build.bat"), fmt.Sprintf("set JAVA_HOME=%v\r\n%v", utils.JDK_DIR(), strings.Join(dep.Args, " ")))
			utils.RunCmd(exec.Command(filepath.Join(depPath, "build.bat")), fmt.Sprintf("deploy for %v on %v", target, runtime.GOOS))
			utils.RemoveAll(filepath.Join(depPath, "build.bat"))
			//<--
		} else {
			utils.RunCmd(dep, fmt.Sprintf("deploy for %v on %v", target, runtime.GOOS))
		}

		if utils.QT_VAGRANT() {
			depPathUNC := strings.Replace(depPath, "C:\\media\\sf_GOPATH", "C:\\media\\UNC\\vboxsrv\\media_sf_GOPATH", -1)
			if utils.ExistsFile(filepath.Join(path, target, name+".keystore")) {
				copy(filepath.Join(depPathUNC, "build", "build", "outputs", "apk", "build-release-signed.apk"), depPath)
			} else {
				copy(filepath.Join(depPathUNC, "build", "build", "outputs", "apk", "build-debug.apk"), depPath)
			}
		} else {
			if utils.ExistsFile(filepath.Join(path, target, name+".keystore")) {
				copy(filepath.Join(depPath, "build", "build", "outputs", "apk", "build-release-signed.apk"), depPath)
			} else {
				copy(filepath.Join(depPath, "build", "build", "outputs", "apk", "build-debug.apk"), depPath)
			}
		}

	case "ios", "ios-simulator":

		//copy default assets
		buildPath := filepath.Join(depPath, "build")
		utils.MkdirAll(filepath.Join(buildPath, "project.xcodeproj"))
		utils.MkdirAll(filepath.Join(buildPath, "Images.xcassets", "AppIcon.appiconset"))
		utils.Save(filepath.Join(buildPath, "Info.plist"), ios_plist(name))
		utils.Save(filepath.Join(buildPath, "Images.xcassets", "AppIcon.appiconset", "Contents.json"), ios_appicon())
		utils.Save(filepath.Join(buildPath, "LaunchScreen.xib"), ios_launchscreen(name))
		utils.Save(filepath.Join(buildPath, "project.xcodeproj", "project.pbxproj"), ios_xcodeproject(depPath))
		copy(filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "ios", "mkspecs", "macx-ios-clang", "Default-568h@2x.png"), buildPath)

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/", depPath)

		//add c_main_wrappers
		utils.Save(filepath.Join(depPath, "gallery_plugin_import.cpp"), ios_plugins())
		utils.Save(filepath.Join(depPath, "gallery_qml_plugin_import.cpp"), ios_qmlplugins())
		utils.Save(filepath.Join(depPath, "qt.conf"), ios_qtconf())

		var targets []string
		switch target {
		case "ios":
			targets = []string{"arm64", "armv7"}
		case "ios-simulator":
			targets = []string{"x86_64", "i386"}
		}

		for i, t := range targets {
			var lib string
			if i == 1 {
				if t == "i386" {
					lib += "_386"
				} else {
					lib += "_" + t
				}
			}
			utils.Save(filepath.Join(depPath, "c_main_wrapper_"+t+".cpp"), "#include \"libgo"+lib+".h\"\nint main(int argc, char *argv[]) { go_main_wrapper(); }")
			cmd := exec.Command("xcrun", "clang++", "c_main_wrapper_"+t+".cpp", "gallery_plugin_import.cpp", "gallery_qml_plugin_import.cpp", "-o", "build/main_"+t, "-u", "_qt_registerPlatformPlugin", "-Wl,-e,_qt_main_wrapper", "-I../..", "-L.", "-lgo"+lib)
			cmd.Args = append(cmd.Args, templater.GetiOSClang(target, t)...) //TODO: parse cached core instead
			cmd.Dir = depPath
			utils.RunCmd(cmd, fmt.Sprintf("compile wrapper for %v (%v) on %v", target, t, runtime.GOOS))

			strip := exec.Command("strip", "main_"+t)
			strip.Dir = filepath.Join(depPath, "build")
			utils.RunCmd(strip, fmt.Sprintf("strip binary for %v (%v) on %v", target, t, runtime.GOOS))
		}

		//binary size limits	=> https://developer.apple.com/library/ios/documentation/LanguagesUtilities/Conceptual/iTunesConnect_Guide/Chapters/SubmittingTheApp.html
		lipo := exec.Command("xcrun", "lipo", "-create", "-arch", targets[0], "main_"+targets[0], "-arch", targets[1], "main_"+targets[1], "-output", "main")
		lipo.Dir = filepath.Join(depPath, "build")
		utils.RunCmdOptional(lipo, fmt.Sprintf("create fat binary for %v on %v", target, runtime.GOOS))

		//run xcodebuild
		utils.RunCmd(exec.Command("xcrun", "xcodebuild", "clean", "build", "CODE_SIGN_IDENTITY=", "CODE_SIGNING_REQUIRED=NO", "CONFIGURATION_BUILD_DIR="+depPath, "-configuration", "Release", "-project", filepath.Join(depPath, "build", "project.xcodeproj")), fmt.Sprintf("deploy for %v on %v", target, runtime.GOOS))

	case "sailfish", "sailfish-emulator":

		//copy default assets
		utils.MkdirAll(filepath.Join(depPath, "rpm"))
		utils.Save(filepath.Join(depPath, "rpm", name+".spec"), sailfish_spec(name))
		utils.Save(filepath.Join(depPath, name+".desktop"), sailfish_desktop(name))
		copy(filepath.Join(utils.SAILFISH_DIR(), "tutorials", "stocqt", "stocqt.png"), filepath.Join(depPath, fmt.Sprintf("harbour-%v.png", name)))

		//copy custom assets
		assets := filepath.Join(path, target)
		utils.MkdirAll(assets)
		copy(assets+"/", depPath)

		err := sailfish_ssh("2222", "mersdk", "cd", "/home/mersdk", "&&", "rm", "-R", target)
		if err != nil {
			utils.Log.WithError(err).Warnf("failed to cleanup for %v on %v", target, runtime.GOOS)
		}

		err = sailfish_ssh("2222", "mersdk", "cd", strings.Replace(strings.Replace(path, utils.MustGoPath(), "/media/sf_GOPATH", -1)+"/deploy", "\\", "/", -1), "&&", "cp", "-R", target, "/home/mersdk")
		if err != nil {
			utils.Log.WithError(err).Panicf("failed to copy project for %v on %v", target, runtime.GOOS)
		}

		arch := "i486"
		if target == "sailfish" {
			arch = "armv7hl"
		}
		err = sailfish_ssh("2222", "mersdk", "cd", "/home/mersdk/"+target, "&&", "mb2", "-t", "SailfishOS-"+arch, "build")
		if err != nil {
			utils.Log.WithError(err).Errorf("failed to deploy for %v (%v) on %v", target, arch, runtime.GOOS)
		}

		err = sailfish_ssh("2222", "mersdk", "cd", "/home/mersdk/"+target+"/RPMS", "&&", "cp", "*", strings.Replace(strings.Replace(depPath, utils.MustGoPath(), "/media/sf_GOPATH", -1), "\\", "/", -1))
		if err != nil {
			utils.Log.WithError(err).Panicf("failed to receive project for %v on %v", target, runtime.GOOS)
		}
	}

	if utils.QT_DOCKER() {
		if idug, ok := os.LookupEnv("IDUG"); ok {
			utils.RunCmd(exec.Command("chown", "-R", idug, path), "chown files to user")
		}
	}
}
