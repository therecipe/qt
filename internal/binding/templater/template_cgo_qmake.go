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

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/utils"
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
	utils.Log.WithField("module", module).WithField("path", path).WithField("target", target).WithField("mode", mode).WithField("pkg", ipkg)

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
	if !(target == "sailfish" || target == "sailfish-emulator" || target == "js") {
		if !parser.ShouldBuildForTarget(module, target) ||
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
	for _, file := range cgoFileNames(module, path, target, mode) {
		file = filepath.Join(path, file)
		if utils.ExistsFile(file) {
			file = utils.Load(file)

			for _, dep := range libs {
				if !strings.Contains(strings.ToLower(file), strings.ToLower(dep)) {
					utils.Log.Debugln("cgo does not contain:", strings.ToLower(dep))
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

			if !strings.Contains(file, utils.QT_VERSION()) {
				utils.Log.Debugln("wrong cgo file qt version, re-creating ...")
				return false
			}

			switch target {
			case "darwin", "linux", "windows", "ubports":
				//TODO: msys pkg-config mxe brew
				switch {
				case utils.QT_HOMEBREW():
					return strings.Contains(file, utils.QT_DARWIN_DIR())
				case utils.QT_MSYS2():
					return strings.Contains(file, utils.QT_MSYS2_DIR())
				default:
					return strings.Contains(file, utils.QT_DIR()) || strings.Contains(file, utils.QT_MXE_TRIPLET())
				}
			case "android", "android-emulator":
				return strings.Contains(file, utils.QT_DIR()) && strings.Contains(file, utils.ANDROID_NDK_DIR())
			case "ios", "ios-simulator":
				return strings.Contains(file, utils.QT_DIR()) || strings.Contains(file, utils.QT_DARWIN_DIR())
			case "sailfish", "sailfish-emulator", "asteroid":
			case "rpi1", "rpi2", "rpi3":
				return strings.Contains(file, strings.TrimSpace(utils.RunCmd(exec.Command(utils.ToolPath("qmake", target), "-query", "QT_INSTALL_LIBS"), fmt.Sprintf("query lib path for %v on %v", target, runtime.GOOS))))
			case "js":
			}
		}
	}
	return false
}

func createProject(module, path, target string, mode int, libs []string) {
	var out []string

	switch {
	case mode == RCC:
		out = []string{"Core"}
	case mode == MOC, module == "build_ios":
		out = libs
	case mode == MINIMAL, mode == NONE:
		out = append([]string{module}, libs...)
	}

	if strings.HasPrefix(target, "rpi") && module == "Core" {
		out = append([]string{module}, []string{"Widgets", "Gui"}...)
	}

	for i, v := range out {
		if v == "Speech" {
			out[i] = "TextToSpeech"
		}
		if v == "Multimedia" {
			out = append(out, "multimedia-private")
		}
		out[i] = strings.ToLower(out[i])
	}

	proPath := filepath.Join(path, "..", fmt.Sprintf("%v.pro", filepath.Base(path)))
	if module == "build_ios" {
		proPath = filepath.Join(path, "..", "..", fmt.Sprintf("%v.pro", filepath.Base(path)))
	}

	if utils.QT_UBPORTS() {
		proPath = strings.Replace(proPath, "/../", "/", -1)
		proPath = strings.Replace(proPath, "/", "_", -1)
		proPath = filepath.Join("/home", "user", proPath)
	}

	utils.Save(proPath, fmt.Sprintf("QT += %v", strings.Join(out, " ")))
}

func createMakefile(module, path, target string, mode int) {
	proPath := filepath.Join(path, "..", fmt.Sprintf("%v.pro", filepath.Base(path)))
	if module == "build_ios" {
		proPath = filepath.Join(path, "..", "..", fmt.Sprintf("%v.pro", filepath.Base(path)))
	}

	mPath := "Mfile"
	if utils.QT_UBPORTS() {
		proPath = strings.Replace(proPath, "/../", "/", -1)
		proPath = strings.Replace(proPath, "/", "_", -1)
		proPath = filepath.Join("/home", "user", proPath)
		mPath = proPath + mPath
	}

	relProPath, err := filepath.Rel(path, proPath)
	if err != nil || utils.QT_UBPORTS() {
		relProPath = proPath
	}
	env, _, _, _ := cmd.BuildEnv(target, "", "")
	cmd := exec.Command(utils.ToolPath("qmake", target), "-o", mPath, relProPath)
	cmd.Dir = path
	switch target {
	case "darwin":
		cmd.Args = append(cmd.Args, []string{"-spec", "macx-clang", "CONFIG+=x86_64"}...)
	case "windows":
		cmd.Args = append(cmd.Args, []string{"-spec", "win32-g++", "CONFIG+=windows"}...)
	case "linux":
		cmd.Args = append(cmd.Args, []string{"-spec", "linux-g++"}...)
	case "ios":
		cmd.Args = append(cmd.Args, []string{"-spec", "macx-ios-clang", "CONFIG+=iphoneos", "CONFIG+=device"}...)
	case "ios-simulator":
		cmd.Args = append(cmd.Args, []string{"-spec", "macx-ios-clang", "CONFIG+=iphonesimulator", "CONFIG+=simulator"}...)
	case "android", "android-emulator":
		cmd.Args = append(cmd.Args, []string{"-spec", "android-g++"}...)
		cmd.Env = []string{fmt.Sprintf("ANDROID_NDK_ROOT=%v", utils.ANDROID_NDK_DIR())}
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
	case "js":
		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}
	}

	if utils.QT_DEBUG_QML() {
		cmd.Args = append(cmd.Args, []string{"CONFIG+=debug", "CONFIG+=declarative_debug", "CONFIG+=qml_debug"}...)
	} else {
		cmd.Args = append(cmd.Args, "CONFIG+=release")
	}

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
	}

	utils.RemoveAll(proPath)
	utils.RemoveAll(filepath.Join(path, ".qmake.stash"))
	switch target {
	case "darwin":
	case "windows":
		for _, suf := range []string{"_plugin_import", "_qml_plugin_import"} {
			pPath := filepath.Join(path, fmt.Sprintf("%v%v.cpp", filepath.Base(path), suf))
			if (utils.QT_MXE_STATIC() || utils.QT_MSYS2_STATIC()) && utils.ExistsFile(pPath) {
				if content := utils.Load(pPath); !strings.Contains(content, "+build windows") {
					utils.Save(pPath, "// +build windows\r\n"+content)
				}
			}
			if mode == MOC || mode == RCC || !(utils.QT_MXE_STATIC() || utils.QT_MSYS2_STATIC()) || (!strings.HasPrefix(module, "Q") && strings.Contains(pPath, "_qml_")) {
				utils.RemoveAll(pPath)
			}
		}
		for _, n := range []string{"Mfile", "Mfile.Debug", "release", "debug"} {
			utils.RemoveAll(filepath.Join(path, n))
		}
	case "linux":
	case "ios", "ios-simulator":
		for _, suf := range []string{"_plugin_import", "_qml_plugin_import"} {
			pPath := filepath.Join(path, fmt.Sprintf("%v%v.cpp", filepath.Base(path), suf))
			/* TODO:
			if utils.QT_VERSION_MAJOR() == "5.9" && utils.ExistsFile(pPath) {
				if content := utils.Load(pPath); !strings.Contains(content, "+build ios,!darwin") {
					utils.Save(pPath, "// +build ios,!darwin\n"+utils.Load(pPath))
				}
			}
			*/
			if module != "build_ios" /*TODO: utils.QT_VERSION_MAJOR() != "5.9"*/ || mode == MOC || mode == RCC {
				utils.RemoveAll(pPath)
			}
		}
		for _, n := range []string{"Info.plist", "qt.conf"} {
			utils.RemoveAll(filepath.Join(path, n))
		}
		utils.RemoveAll(filepath.Join(path, fmt.Sprintf("%v.xcodeproj", filepath.Base(path))))
	case "android", "android-emulator":
		utils.RemoveAll(filepath.Join(path, fmt.Sprintf("android-lib%v.so-deployment-settings.json", filepath.Base(path))))
	case "sailfish", "sailfish-emulator":
	case "asteroid":
	case "rpi1", "rpi2", "rpi3":
	case "ubports":
	case "js":
		for _, suf := range []string{".js_plugin_import", ".js_qml_plugin_import"} {
			pPath := filepath.Join(path, fmt.Sprintf("%v%v.cpp", filepath.Base(path), suf))
			if module != "build_ios" || mode == MOC || mode == RCC {
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
	case "js":
		guards += "ignore"
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
	if mode == MOC || pkg == "build_ios" {
		pkg = ipkg
	}
	fmt.Fprintf(bb, "package %v\n\n/*\n", pkg)

	//

	file := "Mfile"
	if target == "windows" {
		file += ".Release"
	}
	content := utils.Load(filepath.Join(path, file))

	for _, l := range strings.Split(content, "\n") {
		switch {
		case strings.HasPrefix(l, "CFLAGS"):
			fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
		case strings.HasPrefix(l, "CXXFLAGS"), strings.HasPrefix(l, "INCPATH"):
			fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
		case strings.HasPrefix(l, "LFLAGS"), strings.HasPrefix(l, "LIBS"):
			if target == "windows" && !(utils.QT_MXE_STATIC() || utils.QT_MSYS2_STATIC()) {
				pFix := []string{
					filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "mingw53_32"),
					filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "qt5"),
					utils.QT_MSYS2_DIR(),
				}
				for _, pFix := range pFix {
					pFix = strings.Replace(filepath.Join(pFix, "lib", "lib"), "\\", "/", -1)
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
			}
			fmt.Fprintf(bb, "#cgo LDFLAGS: %v\n", strings.Split(l, " = ")[1])
		}
	}

	switch target {
	case "android", "android-emulator":
		fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,--allow-shlib-undefined\n")
	case "windows":
		fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,--allow-multiple-definition\n")
	case "ios":
		fmt.Fprintf(bb, "#cgo CXXFLAGS: -isysroot %v/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/%v -miphoneos-version-min=10.0\n", utils.XCODE_DIR(), utils.IPHONEOS_SDK_DIR())
		fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-syslibroot,%v/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/%v -miphoneos-version-min=10.0\n", utils.XCODE_DIR(), utils.IPHONEOS_SDK_DIR())
	case "ios-simulator":
		fmt.Fprintf(bb, "#cgo CXXFLAGS: -isysroot %v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/%v -mios-simulator-version-min=10.0\n", utils.XCODE_DIR(), utils.IPHONESIMULATOR_SDK_DIR())
		fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-syslibroot,%v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/%v -mios-simulator-version-min=10.0\n", utils.XCODE_DIR(), utils.IPHONESIMULATOR_SDK_DIR())
	}

	fmt.Fprintf(bb, "#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type\n")
	fmt.Fprintf(bb, "#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type\n")

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
	case "android", "android-emulator":
		tmp = strings.Replace(tmp, fmt.Sprintf("-Wl,-soname,lib%v.so", filepath.Base(path)), "", -1)
		tmp = strings.Replace(tmp, "-shared", "", -1)
	case "js":
		tmp = strings.Replace(tmp, "\"", "", -1)
		if utils.QT_DEBUG() {
			tmp = strings.Replace(tmp, "-s USE_FREETYPE=1", "-s USE_FREETYPE=1 -s ASSERTIONS=1", -1)
		}
		tmp = strings.Replace(tmp, "-s NO_EXIT_RUNTIME=0", "-s NO_EXIT_RUNTIME=1", -1) //TODO: block main instead
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

	if module == "build_ios" {
		return tmp
	}

	for _, file := range cgoFileNames(module, path, target, mode) {
		switch target {
		case "darwin":
			for _, lib := range []string{"WebKitWidgets", "WebKit"} {
				tmp = strings.Replace(tmp, "-lQt5"+lib, "-framework Qt"+lib, -1)
			}
		case "windows":
			if utils.QT_MSYS2() {
				tmp = strings.Replace(tmp, ",--relax,--gc-sections", "", -1)
			}
			if utils.QT_MSYS2() && utils.QT_MSYS2_ARCH() == "amd64" {
				tmp = strings.Replace(tmp, " -Wa,-mbig-obj ", " ", -1)
			}
			if (utils.QT_MSYS2() && utils.QT_MSYS2_ARCH() == "amd64") || utils.QT_MXE_ARCH() == "amd64" {
				tmp = strings.Replace(tmp, " -Wl,-s ", " ", -1)
			}
		case "ios":
			if strings.HasSuffix(file, "darwin_arm.go") {
				tmp = strings.Replace(tmp, "arm64", "armv7", -1)
			}
		case "ios-simulator":
			if strings.HasSuffix(file, "darwin_386.go") {
				tmp = strings.Replace(tmp, "x86_64", "i386", -1)
			}
		case "js":
			if mode == RCC {
				utils.Save(filepath.Join(path, strings.Replace(file, "_cgo_", "_stub_", -1)), "package "+pkg+"\n")
			}
		}
		utils.Save(filepath.Join(path, file), tmp)
	}

	return ""
}

func cgoFileNames(module, path, target string, mode int) []string {
	var pFix string
	switch mode {
	case RCC:
		pFix = "rcc_"
	case MOC:
		pFix = "moc_"
	case MINIMAL:
		pFix = "minimal_"
	}

	var sFixes []string
	switch target {
	case "darwin":
		sFixes = []string{"darwin_amd64"}
	case "linux":
		sFixes = []string{"linux_" + utils.GOARCH()}
	case "windows":
		if utils.QT_MXE_ARCH() == "amd64" || (utils.QT_MSYS2() && utils.QT_MSYS2_ARCH() == "amd64") {
			sFixes = []string{"windows_amd64"}
		} else {
			sFixes = []string{"windows_386"}
		}
	case "android":
		sFixes = []string{"linux_arm"}
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
	}

	var o []string
	for _, sFix := range sFixes {
		o = append(o, fmt.Sprintf("%vcgo_%v_%v.go", pFix, strings.Replace(target, "-", "_", -1), sFix))
	}
	return o
}

func ParseCgo(module, target string) (string, string) {
	utils.Log.WithField("module", module).WithField("target", target).Debug("parse cgo for shared lib")

	tmp := utils.LoadOptional(utils.GoQtPkgPath(module, cgoFileNames(module, "", target, NONE)[0]))
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
		case "js":
			env, _, _, _ := cmd.BuildEnv(target, "", "")
			return filepath.Join(env["EMSCRIPTEN"], "em++"), fmt.Sprintf("%v -o %[2]v.o %[2]v.js_plugin_import.cpp %[2]v.cpp", tmp, module)
		}
	}

	return "", tmp
}

func ReplaceCgo(module, target string) {
	utils.Log.WithField("module", module).WithField("target", target).Debug("replace cgo for shared lib")

	if target == "js" {
		//TODO: cleanup ?
		//utils.RemoveAll(utils.GoQtPkgPath(module, cgoFileNames(module, "", target, NONE)[0]))
		return
	}

	tmp := utils.LoadOptional(utils.GoQtPkgPath(module, cgoFileNames(module, "", target, NONE)[0]))
	if tmp != "" {
		pre := strings.Split(tmp, "/*")[0]
		past := strings.Split(tmp, "*/")[1]
		utils.Save(utils.GoQtPkgPath(module, cgoFileNames(module, "", target, NONE)[0]), fmt.Sprintf("%v/*\n#cgo CFLAGS: -I.\n#cgo LDFLAGS: -L. -l%v -Wl,-rpath,%v\n*/%v", pre, module, utils.GoQtPkgPath(), past))
	}
}
