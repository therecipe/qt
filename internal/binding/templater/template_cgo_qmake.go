package templater

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bluszcz/cutego/internal/binding/parser"
	"github.com/bluszcz/cutego/internal/cmd"
	"github.com/bluszcz/cutego/internal/utils"
)

const (
	NONE = iota
	MOC
	MINIMAL
	RCC
)

func CgoTemplate(module, path, target string, mode int, ipkg, tags string) (o string) {
	return cgoTemplate(module, path, target, mode, ipkg, tags, parser.LibDeps[module])
}

func CgoTemplateSafe(module, path, target string, mode int, ipkg, tags string, libs []string) (o string) {
	return cgoTemplate(module, path, target, mode, ipkg, tags, libs)
}

func cgoTemplate(module, path, target string, mode int, ipkg, tags string, libs []string) (o string) {
	utils.Log.WithField("module", module).WithField("path", path).WithField("target", target).WithField("mode", mode).WithField("pkg", ipkg).Debug("running cgoTemplate")

	if utils.UseGOMOD(utils.GoQtPkgPath("core")) {
		utils.MkdirAll(utils.GoQtPkgPath(strings.ToLower(module)))
	}

	switch module {
	case "AndroidExtras":
		if !(target == "android" || target == "android-emulator") {
			return
		}
	case "Sailfish":
		if !strings.HasPrefix(target, "sailfish") {
			return
		}
	}

	if path == "" {
		path = utils.GoQtPkgPath(strings.ToLower(module))
	}

	//TODO: differentiate between docker and virtual-box build for sailfish targets
	if !(target == "sailfish" || target == "sailfish-emulator" || target == "js" || target == "wasm") {
		if !(parser.ShouldBuildForTarget(module, target) || mode == MOC || mode == RCC) ||
			isAlreadyCached(module, path, target, mode, libs) {
			utils.Log.Debugf("skipping cgo generation")
			return
		}
	}

	switch target {
	case "sailfish", "sailfish-emulator":
		cgoSailfish(module, path, mode, ipkg, libs) //TODO:
	case "asteroid":
		cgoAsteroid(module, path, mode, ipkg) //TODO:
	default:
		createProject(module, path, target, mode, libs)
		createMakefile(module, path, target, mode)
		o = createCgo(module, path, target, mode, ipkg, tags)
	}

	utils.RemoveAll(filepath.Join(path, "Mfile"))
	utils.RemoveAll(filepath.Join(path, "Mfile.Release"))

	return
}

//TODO: use qmake props ?
func isAlreadyCached(module, path, target string, mode int, libs []string) bool {
	if utils.QT_NOT_CACHED() {
		return false
	}

	for _, file := range CgoFileNames(path, target, mode) {
		file = filepath.Join(path, file)
		if utils.ExistsFile(file) {
			file = utils.Load(file)

			for _, dep := range libs {
				if !strings.Contains(strings.ToLower(file), "_"+strings.ToLower(dep)+"_") {
					utils.Log.Debugln("cgo does not contain:", strings.ToLower(dep))
					return false
				}
			}

			allLibs := parser.GetLibs()
			parser.LibDepsMutex.Lock()
			for i := len(allLibs) - 1; i >= 0; i-- {
				for _, dep := range append(libs, module) {
					var broke bool
					for _, lib := range append(parser.LibDeps[dep], dep) {
						if allLibs[i] == lib {
							allLibs = append(allLibs[:i], allLibs[i+1:]...)
							broke = true
							break
						}
					}
					if broke {
						break
					}
				}
			}
			parser.LibDepsMutex.Unlock()

			for _, dep := range allLibs {
				if strings.Contains(strings.ToLower(file), "_"+strings.ToLower(dep)+"_") {
					utils.Log.Debugln("cgo does contain extra:", strings.ToLower(dep))
					return false
				}
			}

			if utils.QT_DEBUG_QML() {
				if strings.Contains(file, "-DQT_NO_DEBUG") {
					utils.Log.Debugln("non debug cgo file, re-creating ...")
					return false
				}
			} else {
				if strings.Contains(file, "-DQT_QML_DEBUG") || strings.Contains(file, "-DQT_DECLARATIVE_DEBUG") {
					utils.Log.Debugln("non release cgo file, re-creating ...")
					return false
				}
			}

			if !strings.Contains(file, utils.QT_VERSION()) && strings.Contains(file, "5.") {
				utils.Log.Debugln("wrong cgo file qt version, re-creating ...")
				return false
			}

			switch target {
			case "windows":
				if utils.QT_DEBUG_CONSOLE() {
					if strings.Contains(file, "subsystem,windows") || strings.Contains(file, "/SUBSYSTEM:WINDOWS") {
						utils.Log.Debugln("wrong subsystem: have windows and want console, re-creating ...")
						return false
					}
				} else {
					if strings.Contains(file, "subsystem,console") || strings.Contains(file, "/SUBSYSTEM:CONSOLE") {
						utils.Log.Debugln("wrong subsystem: have console and want windows, re-creating ...")
						return false
					}
				}
			case "darwin":
				if !strings.Contains(file, utils.MACOS_SDK_DIR()) {
					utils.Log.Debugln("wrong MACOS_SDK_DIR, re-creating ...")
					return false
				}
			case "ios":
				if !strings.Contains(file, utils.IPHONEOS_SDK_DIR()) {
					utils.Log.Debugln("wrong IPHONEOS_SDK_DIR, re-creating ...")
					return false
				}
			case "ios-simulator":
				if !strings.Contains(file, utils.IPHONESIMULATOR_SDK_DIR()) {
					utils.Log.Debugln("wrong IPHONESIMULATOR_SDK_DIR, re-creating ...")
					return false
				}
			}

			containsPath := func(file, path string) bool {
				r := strings.Contains(strings.Replace(strings.Replace(file, "\\", "", -1), "/", "", -1), strings.Replace(strings.Replace(strings.TrimPrefix(path, filepath.VolumeName(path)), "\\", "", -1), "/", "", -1))
				if !r {
					utils.Log.Debugln("wrong qt path, re-creating ...")
				}
				return r
			}

			switch target {
			case "darwin", "linux", "windows", "ubports", "freebsd":
				//TODO: msys pkg-config mxe brew
				switch {
				case utils.QT_HOMEBREW(), utils.QT_MACPORTS(), utils.QT_NIX(), utils.QT_FELGO(), utils.QT_MSYS2():
					return containsPath(file, utils.QT_INSTALL_PREFIX(target))
				default:
					return containsPath(file, utils.QT_INSTALL_PREFIX(target)) || strings.Contains(file, utils.QT_MXE_TRIPLET())
				}
			case "android", "android-emulator":
				return containsPath(file, utils.QT_INSTALL_PREFIX(target)) && strings.Contains(file, utils.ANDROID_NDK_DIR())
			case "ios", "ios-simulator":
				return containsPath(file, utils.QT_INSTALL_PREFIX(target))
			case "sailfish", "sailfish-emulator", "asteroid":
			case "rpi1", "rpi2", "rpi3":
				return containsPath(file, utils.QT_INSTALL_PREFIX(target))
			case "js", "wasm":
				return containsPath(file, utils.QT_INSTALL_PREFIX(target))
			}
		}
	}
	return false
}

func createProject(module, path, target string, mode int, libs []string) {
	var out []string

	switch {
	case mode == RCC, mode == MOC, module == "build_static":
		out = libs
	case mode == MINIMAL, mode == NONE:
		out = append([]string{module}, libs...)
	}

	var hasVirtualKeyboard bool
	for i, v := range out {
		if v == "VirtualKeyboard" {
			hasVirtualKeyboard = true
		}

		if v == "Speech" {
			out[i] = "TextToSpeech"
		}
		out[i] = strings.ToLower(out[i])
	}

	proPath := filepath.Join(path, "..", fmt.Sprintf("%v.pro", filepath.Base(path)))
	if module == "build_static" {
		proPath = filepath.Join(path, "..", "..", fmt.Sprintf("%v.pro", filepath.Base(path)))
	}

	if utils.QT_UBPORTS() {
		proPath = strings.Replace(proPath, "/../", "/", -1)
		proPath = strings.Replace(proPath, "/", "_", -1)
		proPath = filepath.Join("/home", "user", proPath)
	}

	if utils.UseGOMOD(path) && (mode == NONE || mode == MINIMAL) {
		//TODO: use project root for MINIMAL regardless of go module usage?
		proPath = filepath.Join(filepath.Dir(utils.GOMOD(path)), fmt.Sprintf("%v.pro", filepath.Base(path)))
	}

	bb := new(bytes.Buffer)
	defer bb.Reset()
	for _, o := range out {
		if !cmd.ImportsQmlOrQuick() && (o == "qml" || o == "quick") && module == "build_static" {
			continue
		}
		fmt.Fprintf(bb, "qtHaveModule(%[1]v) { QT+=%[1]v }\n", o)
		if o == "quickcontrols2" && utils.QT_GEN_QUICK_EXTRAS() {
			fmt.Fprintf(bb, "qtHaveModule(qml-private) { QT+=qml-private }\n", o)
			fmt.Fprintf(bb, "qtHaveModule(quick-private) { QT+=quick-private }\n", o)
			fmt.Fprintf(bb, "qtHaveModule(quicktemplates2-private) { QT+=quicktemplates2-private }\n", o)
			fmt.Fprintf(bb, "qtHaveModule(quickcontrols2-private) { QT+=quickcontrols2-private }\n", o)
		}
	}
	if hasVirtualKeyboard && (utils.QT_STATIC() || target == "js" || target == "wasm") {
		fmt.Fprintf(bb, "QTPLUGIN += qtvirtualkeyboardplugin\n")
	}
	if utils.QT_STATIC() && strings.ToLower(module) == "core" && target == "linux" {
		fmt.Fprintf(bb, "QTPLUGIN += ibusplatforminputcontextplugin\n")
	}
	utils.SaveBytes(proPath, bb.Bytes())
}

func createMakefile(module, path, target string, mode int) {

	for _, suf := range []string{"_plugin_import", "_qml_plugin_import"} {
		pPath := filepath.Join(path, fmt.Sprintf("%v%v.cpp", filepath.Base(path), suf))
		if utils.ExistsFile(pPath) {
			utils.RemoveAll(pPath)
		}
	}

	mPath := "Mfile"
	proPath := filepath.Join(path, "..", fmt.Sprintf("%v.pro", filepath.Base(path)))

	if module == "build_static" {
		proPath = filepath.Join(path, "..", "..", fmt.Sprintf("%v.pro", filepath.Base(path)))
	}

	if utils.QT_UBPORTS() {
		proPath = strings.Replace(proPath, "/../", "/", -1)
		proPath = strings.Replace(proPath, "/", "_", -1)
		proPath = filepath.Join("/home", "user", proPath)
		mPath = proPath + mPath
	}

	if utils.UseGOMOD(path) && (mode == NONE || mode == MINIMAL) {
		//TODO: use project root for MINIMAL regardless of go module usage?
		proPath = filepath.Join(filepath.Dir(utils.GOMOD(path)), fmt.Sprintf("%v.pro", filepath.Base(path)))
	}

	relProPath, err := filepath.Rel(path, proPath)
	if err != nil || utils.QT_UBPORTS() {
		relProPath = proPath
	}
	env, _, _, _ := cmd.BuildEnv(target, "", "")
	cmd := exec.Command(utils.ToolPath("qmake", target), "-nocache", "-nodepend", "-nomoc", "-norecursive", "-o", mPath, relProPath)
	cmd.Dir = path
	switch target {
	case "darwin":
		cmd.Args = append(cmd.Args, []string{"-spec", "macx-clang", "CONFIG+=x86_64"}...)
	case "windows":
		subsystem := "windows"
		if utils.QT_DEBUG_CONSOLE() {
			subsystem = "console"
		}
		spec := "win32-g++"
		if utils.QT_MSVC() {
			spec = "win32-msvc"
		}
		cmd.Args = append(cmd.Args, []string{"-spec", spec, "CONFIG+=" + subsystem}...)
	case "linux":
		cmd.Args = append(cmd.Args, []string{"-spec", "linux-g++"}...)
	case "freebsd":
		cmd.Args = append(cmd.Args, []string{"-spec", "freebsd-clang"}...)
	case "ios":
		cmd.Args = append(cmd.Args, []string{"-spec", "macx-ios-clang", "CONFIG+=iphoneos", "CONFIG+=device"}...)
	case "ios-simulator":
		cmd.Args = append(cmd.Args, []string{"-spec", "macx-ios-clang", "CONFIG+=iphonesimulator", "CONFIG+=simulator"}...)
	case "android", "android-emulator":
		cmd.Args = append(cmd.Args, []string{"-spec", "android-clang"}...)
		if utils.QT_VERSION_NUM() >= 5140 {
			var arch string
			if target == "android" {
				if utils.GOARCH() == "arm64" {
					arch = "arm64-v8a"
				} else {
					arch = "armeabi-v7a"
				}
			} else {
				arch = "x86"
				//TODO: x86_64
			}
			cmd.Args = append(cmd.Args, "ANDROID_ABIS="+arch)
		}
		cmd.Env = []string{fmt.Sprintf("ANDROID_NDK_ROOT=%v", utils.ANDROID_NDK_DIR()), fmt.Sprintf("ANDROID_NDK_PLATFORM=%v", utils.ANDROID_NDK_PLATFORM())}
	case "sailfish", "sailfish-emulator":
		cmd.Args = append(cmd.Args, []string{"-spec", "linux-g++"}...)
		cmd.Env = []string{
			"MER_SSH_PORT=2222",
			fmt.Sprintf("MER_SSH_PRIVATE_KEY=%v", filepath.Join(utils.SAILFISH_DIR(), "vmshare", "ssh", "private_keys", "engine", "mersdk")),
			fmt.Sprintf("MER_SSH_PROJECT_PATH=%v", cmd.Dir),
			fmt.Sprintf("MER_SSH_SDK_TOOLS=%v/.config/SailfishOS-SDK/mer-sdk-tools/MerSDK/SailfishOS-armv7hl", os.Getenv("HOME")),
			fmt.Sprintf("MER_SSH_SHARED_HOME=%v", os.Getenv("HOME")),
			fmt.Sprintf("MER_SSH_SHARED_SRC=%v", utils.MustGoPath()),
			"MER_SSH_SHARED_TARGET=/opt/SailfishOS/mersdk/targets",
			"MER_SSH_TARGET_NAME=SailfishOS-armv7hl",
			"MER_SSH_USERNAME=mersdk",
		}
	case "asteroid":
	case "rpi1":
		cmd.Args = append(cmd.Args, []string{"-spec", "devices/linux-rasp-pi-g++"}...)
	case "rpi2":
		cmd.Args = append(cmd.Args, []string{"-spec", "devices/linux-rasp-pi2-g++"}...)
	case "rpi3":
		cmd.Args = append(cmd.Args, []string{"-spec", "devices/linux-rpi3-g++"}...)
	case "ubports":
		if utils.QT_UBPORTS_ARCH() == "arm" {
			if utils.QT_UBPORTS_VERSION() == "vivid" {
				cmd.Args = append(cmd.Args, []string{"-spec", "ubuntu-arm-gnueabihf-g++"}...)
			} else {
				cmd.Args = append(cmd.Args, []string{"-spec", "linux-g++"}...)
			}
		} else {
			if utils.QT_UBPORTS_VERSION() == "vivid" {
				cmd.Args = append(cmd.Args, []string{"-spec", "linux-g++-64"}...)
			} else {
				cmd.Args = append(cmd.Args, []string{"-spec", "linux-g++"}...)
			}
		}
	case "js", "wasm":
		cmd.Args = append(cmd.Args, []string{"-spec", "wasm-emscripten"}...)
		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}
	}

	if utils.QT_DEBUG_QML() {
		cmd.Args = append(cmd.Args, []string{"CONFIG+=debug", "CONFIG+=declarative_debug", "CONFIG+=qml_debug"}...)
	} else {
		cmd.Args = append(cmd.Args, "CONFIG+=release", "CONFIG-=declarative_debug", "CONFIG-=qml_debug")
	}

	if utils.QT_FELGO() && (module == "Felgo" || module == "build_static" || mode == MOC) {
		cmd.Args = append(cmd.Args, []string{"CONFIG+=felgo"}...)
		if utils.QT_FELGO_LIVE() || mode == NONE {
			cmd.Args = append(cmd.Args, []string{"CONFIG+=felgo-live"}...)
		}
	}

	if utils.QT_RESOURCES_BIG() && mode == RCC {
		cmd.Args = append(cmd.Args, "CONFIG+=resources_big")
	}

	if utils.QT_STATIC() {
		cmd.Args = append(cmd.Args, "CONFIG+=static")
	}

	//needed for QMAKE_MSC_VER bug: https://bugreports.qt.io/browse/QTBUG-59718
	utils.RemoveAll(filepath.Join(cmd.Dir, ".qmake.stash"))

	if (target == "android" || target == "android-emulator") && runtime.GOOS == "windows" {
		//TODO: use os.Setenv instead? -->
		utils.SaveExec(filepath.Join(cmd.Dir, "qmake.bat"), fmt.Sprintf("set ANDROID_NDK_ROOT=%v\r\nset ANDROID_NDK_HOST=windows-x86_64\r\n%v", utils.ANDROID_NDK_DIR(), strings.Join(cmd.Args, " ")))
		cmd = exec.Command(".\\qmake.bat")
		cmd.Dir = path
		utils.RunCmdOptional(cmd, fmt.Sprintf("run qmake for %v on %v", target, runtime.GOOS))
		utils.RemoveAll(filepath.Join(cmd.Dir, "qmake.bat"))
		//<--
	} else {
		utils.RunCmdOptional(cmd, fmt.Sprintf("run qmake for %v on %v", target, runtime.GOOS))
	}

	if utils.QT_UBPORTS() {
		utils.Save(filepath.Join(path, "Mfile"), utils.Load(mPath))
		utils.RemoveAll(mPath)
	}

	if utils.QT_FELGO() {
		utils.RemoveAll(filepath.Join(filepath.Dir(proPath), "FelgoLive"))
		utils.RemoveAll(filepath.Join(path, "android-build"))
		utils.RemoveAll(filepath.Join(path, "gradle.properties"))
	}

	utils.RemoveAll(proPath)
	utils.RemoveAll(filepath.Join(path, ".qmake.stash"))
	switch target {
	case "darwin":
		for _, suf := range []string{"_plugin_import", "_qml_plugin_import"} {
			pPath := filepath.Join(path, fmt.Sprintf("%v%v.cpp", filepath.Base(path), suf))
			if utils.QT_STATIC() && utils.ExistsFile(pPath) {
				if content := utils.Load(pPath); !strings.Contains(content, "+build darwin") {
					utils.Save(pPath, "// +build darwin\r\n"+content)
				}
			}
			if mode == MOC || mode == RCC || !utils.QT_STATIC() || (module != "Qml" && strings.Contains(pPath, "_qml_")) {
				utils.RemoveAll(pPath)
			}
		}
	case "windows":
		for _, suf := range []string{"_plugin_import", "_qml_plugin_import"} {
			pPath := filepath.Join(path, fmt.Sprintf("%v%v.cpp", filepath.Base(path), suf))
			if (utils.QT_MXE_STATIC() || utils.QT_MSYS2_STATIC()) && utils.ExistsFile(pPath) {
				if content := utils.Load(pPath); !strings.Contains(content, "+build windows") {
					utils.Save(pPath, "// +build windows\r\n"+content)
				}
			}
			if mode == MOC || mode == RCC || !(utils.QT_MXE_STATIC() || utils.QT_MSYS2_STATIC()) || (module != "Qml" && strings.Contains(pPath, "_qml_")) {
				utils.RemoveAll(pPath)
			}
		}
		for _, n := range []string{"Mfile", "Mfile.Debug", "release", "debug"} {
			utils.RemoveAll(filepath.Join(path, n))
		}
	case "linux":
		for _, suf := range []string{"_plugin_import", "_qml_plugin_import"} {
			pPath := filepath.Join(path, fmt.Sprintf("%v%v.cpp", filepath.Base(path), suf))
			if utils.QT_STATIC() && utils.ExistsFile(pPath) {
				if content := utils.Load(pPath); !strings.Contains(content, "+build linux") {
					if strings.ToLower(module) == "core" && suf == "_plugin_import" && utils.ExistsFile(filepath.Join(utils.QT_INSTALL_PREFIX(target), "plugins", "platforminputcontexts", "libfcitxplatforminputcontextplugin.a")) {
						utils.Save(pPath, "// +build linux\r\n"+content+"Q_IMPORT_PLUGIN(QFcitxPlatformInputContextPlugin)\r\n")
					} else {
						utils.Save(pPath, "// +build linux\r\n"+content)
					}
				}
			}
			if mode == MOC || mode == RCC || !utils.QT_STATIC() || (module != "Qml" && strings.Contains(pPath, "_qml_")) {
				utils.RemoveAll(pPath)
			}
		}
	case "ios", "ios-simulator":
		for _, suf := range []string{"_plugin_import", "_qml_plugin_import"} {
			pPath := filepath.Join(path, fmt.Sprintf("%v%v.cpp", filepath.Base(path), suf))
			/* TODO when shared builds are available:
			if utils.QT_VERSION_MAJOR() == "5.9" && utils.ExistsFile(pPath) {
				if content := utils.Load(pPath); !strings.Contains(content, "+build ios,!darwin") {
					utils.Save(pPath, "// +build ios,!darwin\n"+utils.Load(pPath))
				}
			}
			*/
			if module != "build_static" || mode == MOC || mode == RCC /*TODO when shared builds are available: utils.QT_VERSION_MAJOR() != "5.9"*/ {
				utils.RemoveAll(pPath)
			}
		}
		for _, n := range []string{"Info.plist", "qt.conf"} {
			utils.RemoveAll(filepath.Join(path, n))
		}
		utils.RemoveAll(filepath.Join(path, fmt.Sprintf("%v.xcodeproj", filepath.Base(path))))
	case "android", "android-emulator":
		utils.RemoveAll(filepath.Join(path, fmt.Sprintf("android-lib%v.so-deployment-settings.json", filepath.Base(path))))
		utils.RemoveAll(filepath.Join(path, fmt.Sprintf("android-%v-deployment-settings.json", filepath.Base(path))))

		var arch string
		if target == "android" {
			if utils.GOARCH() == "arm64" {
				arch = "arm64-v8a"
			} else {
				arch = "armeabi-v7a"
			}
		} else {
			arch = "x86"
			//TODO: x86_64
		}
		utils.RemoveAll(filepath.Join(path, arch))
	case "sailfish", "sailfish-emulator":
	case "asteroid":
	case "rpi1", "rpi2", "rpi3":
	case "ubports":
	case "freebsd":
	case "js", "wasm":
		for _, suf := range []string{".js_plugin_import", ".js_qml_plugin_import"} {
			pPath := filepath.Join(path, fmt.Sprintf("%v%v.cpp", filepath.Base(path), suf))
			if module != "build_static" || mode == MOC || mode == RCC {
				utils.RemoveAll(pPath)
			}
		}
	}
}

func createCgo(module, path, target string, mode int, ipkg, tags string) string {
	bb := new(bytes.Buffer)
	defer bb.Reset()

	if mode == MOC && tags != "" {
		bb.WriteString("// +build " + tags + "\n")
	}

	guards := "// +build "
	switch target {
	case "darwin":
		guards += "!ios"
	case "android", "android-emulator":
		guards += strings.Replace(target, "-", "_", -1)
	case "ios", "ios-simulator":
		guards += "ios"
	case "sailfish", "sailfish-emulator":
		guards += strings.Replace(target, "-", "_", -1)
	case "asteroid":
		guards += target
	case "rpi1", "rpi2", "rpi3":
		guards += target
	case "js", "wasm":
		guards += "ignore"
	case "linux":
		guards += "!ubports"
	}
	//TODO: move "minimal" build tag in separate line -->
	switch mode {
	case NONE:
		if len(guards) > 10 {
			guards += ","
		}
		guards += "!minimal"
	case MINIMAL:
		if len(guards) > 10 {
			guards += ","
		}
		guards += "minimal"
	}
	if len(guards) > 10 {
		bb.WriteString(guards + "\n\n")
	}
	//<--

	pkg := strings.ToLower(module)
	if mode == MOC || pkg == "build_static" {
		pkg = ipkg
	}
	fmt.Fprintf(bb, "package %v\n\n/*\n", pkg)

	//

	switch target {
	case "linux":
		if f := filepath.Join(utils.QT_INSTALL_PREFIX(target), "plugins", "platforminputcontexts", "libfcitxplatforminputcontextplugin.a"); utils.QT_STATIC() && strings.ToLower(module) == "core" && utils.ExistsFile(f) {
			fmt.Fprintf(bb, "#cgo LDFLAGS: %v\n", f)
		}
	}

	file := "Mfile"
	if target == "windows" {
		file += ".Release"
	}
	var content string
	if utils.ExistsFile(filepath.Join(path, file)) {
		content = utils.Load(filepath.Join(path, file))

		for _, l := range strings.Split(content, "\n") {
			switch {
			case strings.HasPrefix(l, "CFLAGS"):
				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
			case strings.HasPrefix(l, "CXXFLAGS"), strings.HasPrefix(l, "INCPATH"):
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			case strings.HasPrefix(l, "LFLAGS"), strings.HasPrefix(l, "LIBS"):
				if target == "windows" && !(utils.QT_MXE_STATIC() || utils.QT_MSYS2_STATIC()) {
					pFix := strings.Replace(filepath.Join(utils.QT_INSTALL_PREFIX(target), "lib", "lib"), "\\", "/", -1)
					if strings.Contains(l, pFix) {
						var cleaned []string
						for _, s := range strings.Split(l, " ") {
							if strings.HasPrefix(s, pFix) && (strings.HasSuffix(s, ".a") || strings.HasSuffix(s, ".dll")) {
								s = strings.Replace(s, pFix, "-l", -1)
								s = strings.TrimSuffix(s, ".a")
								s = strings.TrimSuffix(s, ".dll")
							}
							cleaned = append(cleaned, s)
						}
						l = strings.Join(cleaned, " ")
					}
				}
				fmt.Fprintf(bb, "#cgo LDFLAGS: %v\n", strings.Split(l, " = ")[1])
			}
		}
	}

	switch target {
	case "darwin":
		if utils.QT_STATIC() && utils.QT_DOCKER() {
			fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,-U,___isPlatformVersionAtLeast -Wl,-U,___isOSVersionAtLeast\n")
		}
	case "android", "android-emulator":
		fmt.Fprintf(bb, "#cgo LDFLAGS: %s -Wl,--allow-shlib-undefined\n", utils.ANDROID_NDK_NOSTDLIBPP_LDFLAG())
	case "windows":
		if !utils.QT_MSVC() {
			fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,--allow-multiple-definition\n")
		} else {
			fmt.Fprint(bb, "#cgo LDFLAGS: \"notelemetry.obj\"\n")
		}
	case "ios":
		fmt.Fprintf(bb, "#cgo CXXFLAGS: -isysroot %v/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/%v -miphoneos-version-min=11.0\n", utils.XCODE_DIR(), utils.IPHONEOS_SDK_DIR())
		fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-syslibroot,%v/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/%v -miphoneos-version-min=11.0\n", utils.XCODE_DIR(), utils.IPHONEOS_SDK_DIR())
	case "ios-simulator":
		fmt.Fprintf(bb, "#cgo CXXFLAGS: -isysroot %v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/%v -mios-simulator-version-min=11.0\n", utils.XCODE_DIR(), utils.IPHONESIMULATOR_SDK_DIR())
		fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-syslibroot,%v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/%v -mios-simulator-version-min=11.0\n", utils.XCODE_DIR(), utils.IPHONESIMULATOR_SDK_DIR())
	case "js", "wasm":
		fmt.Fprint(bb, "#cgo CFLAGS: -s EXTRA_EXPORTED_RUNTIME_METHODS=['getValue','setValue']\n")
	}

	if !utils.QT_MSVC() {
		fmt.Fprint(bb, "#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type\n")
		fmt.Fprint(bb, "#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type\n")
	}

	fmt.Fprint(bb, "*/\nimport \"C\"\n")

	out, err := format.Source(bb.Bytes())
	if err != nil {
		utils.Log.WithError(err).Panicln("failed to format:", module)
	}

	tmp := string(out)

	switch target {
	case "darwin":
		tmp = strings.Replace(tmp, "$(EXPORT_ARCH_ARGS)", "-arch x86_64", -1)
	case "ios":
		tmp = strings.Replace(tmp, "$(EXPORT_ARCH_ARGS)", "-arch arm64", -1)
		tmp = strings.Replace(tmp, "$(EXPORT_QMAKE_XARCH_CFLAGS)", "", -1)
		tmp = strings.Replace(tmp, "$(EXPORT_QMAKE_XARCH_LFLAGS)", "", -1)
	case "ios-simulator":
		tmp = strings.Replace(tmp, "$(EXPORT_ARCH_ARGS)", "-arch x86_64", -1)
		tmp = strings.Replace(tmp, "$(EXPORT_QMAKE_XARCH_CFLAGS)", "", -1)
		tmp = strings.Replace(tmp, "$(EXPORT_QMAKE_XARCH_LFLAGS)", "", -1)
	case "android", "android-emulator": //TODO:
		libName := filepath.Base(path)
		if utils.QT_VERSION_NUM() >= 5140 {
			libName += "_"
			// Calculate arch suffix
			if target == "android-emulator" {
				// Android emulator is x86
				libName += "x86"
			} else if utils.GOARCH() == "arm64" {
				// AARCH64 on android
				libName += "arm64-v8a"
			} else {
				// ARMv7 on android
				libName += "armeabi-v7a"
			}
		}
		utils.Log.WithField("module", module).WithField("path", path).WithField("target", target).WithField("mode", mode).WithField("pkg", ipkg).Debugf("cgoTemplate libName: %s", libName)

		tmp = strings.Replace(tmp, fmt.Sprintf("-Wl,-soname,lib%v.so", libName), "-Wl,-soname,libgo_base.so", -1)
		tmp = strings.Replace(tmp, "-shared", "", -1)
	case "js", "wasm":
		tmp = strings.Replace(tmp, "\"", "", -1)
		if utils.QT_DEBUG() {
			tmp = strings.Replace(tmp, "-s USE_FREETYPE=1", "-s USE_FREETYPE=1 -s ASSERTIONS=1", -1)
		}
	}

	for _, variable := range []string{"DEFINES", "SUBLIBS", "EXPORT_QMAKE_XARCH_CFLAGS", "EXPORT_QMAKE_XARCH_LFLAGS", "EXPORT_ARCH_ARGS", "-fvisibility=hidden", "-fembed-bitcode"} {
		for _, l := range strings.Split(content, "\n") {
			if strings.HasPrefix(l, variable+" ") {
				if strings.Contains(l, "-DQT_TESTCASE_BUILDDIR") {
					l = strings.Split(l, "-DQT_TESTCASE_BUILDDIR")[0]
				}
				tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", variable), strings.Split(l, " = ")[1], -1)
			}
		}
		tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", variable), "", -1)
		tmp = strings.Replace(tmp, variable, "", -1)
	}
	tmp = strings.Replace(tmp, "\\", "/", -1)

	if module == "build_static" {
		return tmp
	}

	for _, file := range CgoFileNames(path, target, mode) {
		switch target {
		case "android", "android-emulator":
			tmp = strings.Replace(tmp, "/opt/android/"+filepath.Base(utils.ANDROID_NDK_DIR()), utils.ANDROID_NDK_DIR(), -1)
		case "darwin":
			for _, lib := range []string{"WebKitWidgets", "WebKit"} {
				tmp = strings.Replace(tmp, "-lQt5"+lib, "-framework Qt"+lib, -1)
			}
			tmp = strings.Replace(tmp, "-Wl,-rpath,@executable_path/Frameworks", "", -1)
			tmp = strings.Replace(tmp, "-Wl,-rpath,@executable_path/../Frameworks", "", -1)
			tmp = strings.Replace(tmp, "-weak_framework XCTest", "", -1)
			tmp = strings.Replace(tmp, "-ffunction-sections", "", -1)
			tmp = strings.Replace(tmp, "-fdata-sections", "", -1)
			tmp = strings.Replace(tmp, "-Wl,-dead_strip", "", -1)
		case "windows":
			if utils.QT_MSVC() {
				tmp = strings.Replace(tmp, "-DUNICODE", "-D_ALLOW_KEYWORD_MACROS -DUNICODE", -1)
				tmp = strings.Replace(tmp, "-wd4467", "-wd4467 -wd4716", -1)
				tmp = strings.Replace(tmp, `"/MANIFESTDEPENDENCY:type='win32' name='Microsoft.Windows.Common-Controls' version='6.0.0.0' publicKeyToken='6595b64144ccf1df' language='*' processorArchitecture='*'"`, "", -1) //TODO:

				tmp = strings.Replace(tmp, " CFLAGS:", " MSCFLAGS:", -1)
				tmp = strings.Replace(tmp, " CXXFLAGS:", " MSCXXFLAGS:", -1)
				tmp = strings.Replace(tmp, " LDFLAGS:", " MSLDFLAGS:", -1)
			}
			if utils.QT_MSYS2() {
				tmp = strings.Replace(tmp, ",--relax,--gc-sections", "", -1)
				if utils.QT_MSYS2_STATIC() {
					tmp = strings.Replace(tmp, "-ffunction-sections", "", -1)
					tmp = strings.Replace(tmp, "-fdata-sections", "", -1)
					tmp = strings.Replace(tmp, "-Wl,--gc-sections", "", -1)
					tmp = strings.Replace(tmp, filepath.Join(utils.QT_MSYS2_DIR(), "qt5-static", "lib", "libqmldbg_profiler.a"), "", -1)
				}
				tmp = strings.Replace(tmp, " -lwinmm ", " ", -1)
				tmp = strings.Replace(tmp, " -lkernel32 ", " -lwinmm -lkernel32 ", -1)
			}
			if utils.QT_MSYS2() && utils.QT_MSYS2_ARCH() == "amd64" {
				tmp = strings.Replace(tmp, " -Wa,-mbig-obj ", " ", -1)
			}
			if (utils.QT_MSYS2() && utils.QT_MSYS2_ARCH() == "amd64") || utils.QT_MXE_ARCH() == "amd64" ||
				(!utils.QT_MXE() && !utils.QT_MSYS2() && utils.QT_VERSION_NUM() >= 5120) {
				tmp = strings.Replace(tmp, " -Wl,-s ", " ", -1)
			}
			if utils.QT_DEBUG_CONSOLE() { //TODO: necessary at all?
				tmp = strings.Replace(tmp, "subsystem,windows", "subsystem,console", -1)
				tmp = strings.Replace(tmp, "/SUBSYSTEM:WINDOWS", "/SUBSYSTEM:CONSOLE", -1)
			} else {
				tmp = strings.Replace(tmp, "subsystem,console", "subsystem,windows", -1)
				tmp = strings.Replace(tmp, "/SUBSYSTEM:CONSOLE", "/SUBSYSTEM:WINDOWS", -1)
			}
			if utils.QT_MXE() {
				if li := fmt.Sprintf("-L %v", filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "qt5", "lib")); !strings.Contains(tmp, li) {
					for _, s := range []string{",console -mthreads", ",windows -mthreads"} {
						tmp = strings.Replace(tmp, s, s+" "+li, -1)
					}
				}
			} else if !utils.QT_MSYS2() {
				if li := fmt.Sprintf("-L %v", strings.Replace(filepath.Join(utils.QT_INSTALL_PREFIX(target), "lib"), "\\", "/", -1)); !strings.Contains(tmp, li) {
					for _, s := range []string{",console -mthreads", ",windows -mthreads"} {
						tmp = strings.Replace(tmp, s, s+" "+li, -1)
					}
				}
			}
			if mode == MOC || mode == RCC {
				tmp = strings.Replace(tmp, "-Wl,-subsystem,windows", "", -1)
				tmp = strings.Replace(tmp, "/SUBSYSTEM:WINDOWS", "", -1)
				tmp = strings.Replace(tmp, "-Wl,-subsystem,console", "", -1)
				tmp = strings.Replace(tmp, "/SUBSYSTEM:CONSOLE", "", -1)
			}
		case "ios":
			if strings.HasSuffix(file, "darwin_arm.go") {
				tmp = strings.Replace(tmp, "arm64", "armv7", -1)
			}
		case "ios-simulator":
			if strings.HasSuffix(file, "darwin_386.go") {
				tmp = strings.Replace(tmp, "x86_64", "i386", -1)
			}
		case "js", "wasm":
			if mode == RCC {
				utils.Save(filepath.Join(path, strings.Replace(file, "_cgo_", "_stub_", -1)), "package "+pkg+"\n")
			}
		case "linux", "freebsd":
			tmp = strings.Replace(tmp, "-Wl,-O1", "-O1", -1)
			tmp = strings.Replace(tmp, "-ffunction-sections", "", -1)
			tmp = strings.Replace(tmp, "-fdata-sections", "", -1)
			tmp = strings.Replace(tmp, "-Wl,--gc-sections", "", -1)
		}
		utils.Save(filepath.Join(path, file), tmp)
	}

	return ""
}

func CgoFileNames(path string, target string, mode int) []string {
	return cgoFileNamesRecursive(path, target, mode, true)
}

func cgoFileNamesRecursive(path string, target string, mode int, recursive bool) []string {
	var pFix string
	switch mode {
	case RCC:
		pFix = "rcc_"
		if utils.QT_STATIC() && target == "linux" && recursive { //TODO: fix wrong order issue in LDFLAGS instead
			if mocName := filepath.Join(path, cgoFileNamesRecursive(path, target, MOC, false)[0]); utils.ExistsFile(mocName) {
				pFix = "omit"
			}
		}
	case MOC:
		pFix = "moc_"
		if utils.QT_STATIC() && target == "linux" && recursive { //TODO: fix wrong order issue in LDFLAGS instead
			if rccName := filepath.Join(path, cgoFileNamesRecursive(path, target, RCC, false)[0]); utils.ExistsFile(rccName) {
				utils.RemoveAll(rccName)
			}
		}
	case MINIMAL:
		pFix = "minimal_"
	}

	var sFixes []string
	switch target {
	case "darwin":
		sFixes = []string{"darwin_amd64"}
	case "linux":
		sFixes = []string{"linux_" + utils.GOARCH()}
	case "freebsd":
		sFixes = []string{"freebsd_" + utils.GOARCH()}
	case "windows":
		if utils.QT_MXE_ARCH() == "amd64" || (utils.QT_MSYS2() && utils.QT_MSYS2_ARCH() == "amd64") ||
			(!utils.QT_MXE() && !utils.QT_MSYS2() && utils.QT_VERSION_NUM() >= 5120) {
			if utils.QT_VERSION_NUM() >= 5122 {
				sFixes = []string{"windows_" + utils.GOARCH()}
			} else {
				sFixes = []string{"windows_amd64"}
			}
		} else {
			sFixes = []string{"windows_386"}
		}
	case "android":
		if utils.GOARCH() == "arm64" {
			sFixes = []string{"linux_arm64"}
		} else {
			sFixes = []string{"linux_arm"}
		}
	case "android-emulator":
		sFixes = []string{"linux_386"}
	case "ios":
		sFixes = []string{"darwin_arm64"}
	case "ios-simulator":
		sFixes = []string{"darwin_amd64"}
	case "sailfish":
		sFixes = []string{"linux_arm"}
	case "sailfish-emulator":
		sFixes = []string{"linux_386"}
	case "asteroid":
		sFixes = []string{"linux_arm"}
	case "rpi1", "rpi2", "rpi3":
		sFixes = []string{"linux_arm"}
	case "ubports":
		sFixes = []string{"linux_" + utils.QT_UBPORTS_ARCH()}
	case "js":
		sFixes = []string{"js"}
	case "wasm":
		sFixes = []string{"wasm"}
	}

	var o []string
	for _, sFix := range sFixes {
		if pFix == "omit" {
			continue
		}
		o = append(o, fmt.Sprintf("%vcgo_%v_%v.go", pFix, strings.Replace(target, "-", "_", -1), sFix))
	}
	return o
}

func ParseCgo(module, target string) (string, string) {
	utils.Log.WithField("module", module).WithField("target", target).Debug("parse cgo for shared lib")

	tmp := utils.LoadOptional(utils.GoQtPkgPath(module, CgoFileNames("", target, NONE)[0]))
	if tmp != "" {
		tmp = strings.Split(tmp, "/*")[1]
		tmp = strings.Split(tmp, "*/")[0]

		tmp = strings.Replace(tmp, "#cgo CFLAGS: ", "", -1)
		tmp = strings.Replace(tmp, "#cgo CXXFLAGS: ", "", -1)
		tmp = strings.Replace(tmp, "#cgo LDFLAGS: ", "", -1)
		tmp = strings.Replace(tmp, "\n", " ", -1)

		switch target {
		case "darwin":
			return "clang++", fmt.Sprintf("%v -Wl,-S -Wl,-x -install_name @rpath/%[2]v/lib%[2]v.so -undefined dynamic_lookup -shared -o lib%[2]v.so %[2]v.cpp", tmp, module)
		case "js", "wasm":
			env, _, _, _ := cmd.BuildEnv(target, "", "")
			return filepath.Join(env["EMSCRIPTEN"], "em++"), fmt.Sprintf("%v -o %[2]v.o %[2]v.js_plugin_import.cpp %[2]v.cpp", tmp, module)
		}
	}

	return "", tmp
}

func ReplaceCgo(module, target string) {
	utils.Log.WithField("module", module).WithField("target", target).Debug("replace cgo for shared lib")

	tmp := utils.LoadOptional(utils.GoQtPkgPath(module, CgoFileNames("", target, NONE)[0]))
	if tmp != "" {
		pre := strings.Split(tmp, "/*")[0]
		past := strings.Split(tmp, "*/")[1]
		utils.Save(utils.GoQtPkgPath(module, CgoFileNames("", target, NONE)[0]), fmt.Sprintf("%v/*\n#cgo CFLAGS: -I.\n#cgo LDFLAGS: -L. -l%v -Wl,-rpath,%v\n*/%v", pre, module, utils.GoQtPkgPath(), past))
	}
}
