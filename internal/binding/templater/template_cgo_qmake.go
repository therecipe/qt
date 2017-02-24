package templater

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func QmakeCgoTemplate(module, mocPath, buildTarget string) {
	if buildTarget == "" || buildTarget == "desktop" {
		buildTarget = runtime.GOOS
	}
	qmakeCgo(module, mocPath, buildTarget)
}

func qmakeCgo(module, mocPath, buildTarget string) {

	switch buildTarget {
	case "darwin":
		{
			if runtime.GOOS != buildTarget {
				return
			}
		}

	case "windows":
		{
			if runtime.GOOS != buildTarget {
				//can be windows,linux,darwin
			}
		}

	case "linux":
		{
			if runtime.GOOS != buildTarget {
				return
			}
		}

	case "android":
		{
			switch module {
			case "DBus", "WebEngine", "Designer":
				{
					return
				}

			default:
				{
					if strings.HasSuffix(module, "Extras") && module != "AndroidExtras" {
						return
					}
				}
			}
		}

	case "ios", "ios-simulator":
		{
			switch module {
			case "DBus", "WebEngine", "Designer", "SerialPort", "SerialBus":
				{
					return
				}

			default:
				{
					if strings.HasSuffix(module, "Extras") {
						return
					}
				}
			}
		}
	}

	switch module {
	case "AndroidExtras":
		{
			if buildTarget != "android" {
				return
			}
		}

	case "Sailfish":
		{
			if !(buildTarget == "sailfish" || buildTarget == "sailfish-emulator") {
				return
			}
		}
	}

	var path = utils.GoQtPkgPath(strings.ToLower(module))
	if mocPath != "" {
		path = mocPath
	}

	createProject(module, mocPath, buildTarget)
	createMakefile(module, mocPath, buildTarget)

	switch buildTarget {
	case "darwin":
		{
			qmakeToCgoDarwin(module, mocPath, buildTarget)
		}

	case "windows":
		{
			qmakeToCgoWindows(module, mocPath, buildTarget)
		}

	case "linux":
		{
			qmakeToCgoLinux(module, mocPath, buildTarget)
		}

	case "ios":
		{
			qmakeToCgoIos(module, mocPath, buildTarget)
		}

	case "ios-simulator":
		{
			qmakeToCgoIos(module, mocPath, buildTarget)
		}

	case "android":
		{
			qmakeToCgoAndroid(module, mocPath, buildTarget)
		}

	case "sailfish", "sailfish-emulator":
		{
			qmakeToCgoSailfish(module, mocPath, buildTarget)
		}
	}

	utils.RemoveAll(filepath.Join(path, "Makefile"))
	utils.RemoveAll(filepath.Join(path, "Makefile.Release"))
}

func cleanLibsQmake(module string) []string {

	if parser.State.Rcc {
		return []string{"Core"}
	}

	if parser.State.Moc || module == "build_ios" {
		return parser.LibDeps[module]
	}

	var out = append([]string{module}, parser.LibDeps[module]...)
	for i, v := range out {
		if v == "Speech" {
			out[i] = "TextToSpeech"
		}
	}

	return out
}

func createProject(module, mocPath, buildTarget string) {
	var path = utils.GoQtPkgPath(strings.ToLower(module))
	if mocPath != "" {
		path = mocPath
	}

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	bb.WriteString("QT +=")

	for _, l := range cleanLibsQmake(module) {
		fmt.Fprintf(bb, " %v", strings.ToLower(l))
	}
	utils.SaveBytes(filepath.Join(path, "..", fmt.Sprintf("%v.pro", strings.ToLower(module))), bb.Bytes())
}

func createMakefile(module, mocPath, buildTarget string) {
	var path = utils.GoQtPkgPath(strings.ToLower(module))
	if mocPath != "" {
		path = mocPath
	}

	utils.RunCmd(func() *exec.Cmd {
		var cmd = exec.Command(func() string {
			switch buildTarget {
			case "darwin":
				{
					return filepath.Join(utils.QT_DARWIN_DIR(), "bin", "qmake")
				}

			case "windows":
				{
					if runtime.GOOS != buildTarget {
						var prefix = "i686"
						if utils.QT_MXE_ARCH() == "amd64" {
							prefix = "x86_64"
						}
						var suffix = "shared"
						if utils.QT_MXE_STATIC() {
							suffix = "static"
						}
						return filepath.Join("/usr", "lib", "mxe", "usr", fmt.Sprintf("%v-w64-mingw32.%v", prefix, suffix), "qt5", "bin", "qmake")
					}
					if utils.UseMsys2() {
						return filepath.Join(utils.QT_MSYS2_DIR(), "bin", "qmake")
					}
					return filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "mingw53_32", "bin", "qmake")
				}

			case "linux":
				{
					if utils.UsePkgConfig() {
						return filepath.Join(utils.QT_MISC_DIR(), "bin", "qmake")
					} else {
						return filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "gcc_64", "bin", "qmake")
					}
				}

			case "ios", "ios-simulator":
				{
					return filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "ios", "bin", "qmake")
				}

			case "android":
				{
					return filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "android_armv7", "bin", "qmake")
				}

			case "sailfish":
				{
					return filepath.Join(os.Getenv("HOME"), ".config", "SailfishOS-SDK", "mer-sdk-tools", "MerSDK", "SailfishOS-armv7hl", "qmake")
				}

			case "sailfish-emulator":
				{
					return filepath.Join(os.Getenv("HOME"), ".config", "SailfishOS-SDK", "mer-sdk-tools", "MerSDK", "SailfishOS-i486", "qmake")
				}
			}
			return ""
		}(), filepath.Join(path, "..", fmt.Sprintf("%v.pro", strings.ToLower(module))))
		cmd.Dir = path
		cmd.Args = append(cmd.Args, func() []string {
			switch buildTarget {
			case "darwin":
				{
					return []string{"-spec", "macx-clang", "CONFIG+=x86_64"}
				}

			case "windows":
				{
					return []string{"-spec", "win32-g++"}
				}

			case "linux":
				{
					return []string{"-spec", "linux-g++"}
				}

			case "ios":
				{
					return []string{"-spec", "macx-ios-clang", "CONFIG+=release", "CONFIG+=iphoneos", "CONFIG+=device"}
				}

			case "ios-simulator":
				{
					return []string{"-spec", "macx-ios-clang", "CONFIG+=release", "CONFIG+=iphonesimulator", "CONFIG+=simulator"}
				}

			case "android":
				{
					return []string{"-spec", "android-g++"}
				}

			case "sailfish":
				{

				}

			case "sailfish-emulator":
				{

				}
			}
			return make([]string, 0)
		}()...)
		cmd.Env = append(cmd.Env, func() []string {
			switch buildTarget {
			case "android":
				{
					return []string{fmt.Sprintf("ANDROID_NDK_ROOT=%v", utils.ANDROID_NDK_DIR())}
				}

			case "sailfish", "sailfish-emulator":
				{
					return []string{
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
				}
			}
			return make([]string, 0)
		}()...)
		return cmd
	}(), fmt.Sprintf("run qmake for %v on %v", buildTarget, runtime.GOOS))

	utils.RemoveAll(filepath.Join(path, "..", fmt.Sprintf("%v.pro", strings.ToLower(module))))
	utils.RemoveAll(filepath.Join(path, ".qmake.stash"))

	switch buildTarget {
	case "darwin":
		{

		}

	case "windows":
		{
			for _, suf := range []string{"_plugin_import", "_qml_plugin_import"} {
				var pluginPath = filepath.Join(path, fmt.Sprintf("%v%v.cpp", strings.ToLower(module), suf))
				if utils.QT_MXE_STATIC() {
					if utils.ExistsFile(pluginPath) {
						var content = utils.Load(pluginPath)
						if !strings.Contains(content, "+build windows") {
							utils.Save(pluginPath, "// +build windows\n"+content)
						}
					}
				}
				if mocPath != "" {
					utils.RemoveAll(pluginPath)
				}
			}

			utils.RemoveAll(filepath.Join(path, "Makefile"))
			utils.RemoveAll(filepath.Join(path, "Makefile.Debug"))
			utils.RemoveAll(filepath.Join(path, "release"))
			utils.RemoveAll(filepath.Join(path, "debug"))
		}

	case "linux":
		{

		}

	case "ios", "ios-simulator":
		{
			for _, suf := range []string{"_plugin_import", "_qml_plugin_import"} {
				var pluginPath = filepath.Join(path, fmt.Sprintf("%v%v.cpp", strings.ToLower(module), suf))
				/*TODO: keep plugins here in 5.9 and link against c-shared with qtdeploy
				if utils.ExistsFile(pluginPath) {
				utils.Save(pluginPath, "// +build ios\n"+utils.Load(pluginPath))
				}
				if mocPath != "" {
				*/
				utils.RemoveAll(pluginPath)
				//}
			}

			utils.RemoveAll(filepath.Join(path, fmt.Sprintf("%v.xcodeproj", strings.ToLower(module))))
			utils.RemoveAll(filepath.Join(path, "Info.plist"))
			utils.RemoveAll(filepath.Join(path, "qt.conf"))
		}

	case "android":
		{
			utils.RemoveAll(filepath.Join(path, fmt.Sprintf("android-lib%v.so-deployment-settings.json", strings.ToLower(module))))
		}

	case "sailfish", "sailfish-emulator":
		{

		}
	}
}

func qmakeToCgoDarwin(module, mocPath, buildTarget string) {
	var path = utils.GoQtPkgPath(strings.ToLower(module))
	if mocPath != "" {
		path = mocPath
	}

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build !ios%v\n\n", func() string {
		if parser.State.Minimal {
			return ",minimal"
		}
		if parser.State.Moc || parser.State.Rcc {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if parser.State.Moc {
			return parser.State.Module
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	for _, l := range strings.Split(utils.Load(filepath.Join(path, "Makefile")), "\n") {
		switch {
		case
			strings.HasPrefix(l, "DEFINES"):
			{
				if strings.Contains(l, "-DQT_TESTCASE_BUILDDIR") {
					l = strings.Split(l, "-DQT_TESTCASE_BUILDDIR")[0]
				}

				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CFLAGS"):
			{
				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CXXFLAGS"),
			strings.HasPrefix(l, "INCPATH"):
			{
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "LFLAGS"),
			strings.HasPrefix(l, "LIBS"):
			{

				fmt.Fprintf(bb, "#cgo LDFLAGS: %v\n", strings.Split(l, " = ")[1])
			}
		}
	}

	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = bb.String()

	for _, bl := range []string{"DEFINES", "SUBLIBS", "EXPORT_QMAKE_XARCH_CFLAGS", "EXPORT_QMAKE_XARCH_LFLAGS"} {
		tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", bl), "", -1)
	}

	bb.Reset()
	bb.WriteString(tmp)

	switch {
	case parser.State.Rcc:
		{
			utils.SaveBytes(filepath.Join(path, "rcc_cgo_desktop_darwin_amd64.go"), bb.Bytes())
		}

	case parser.State.Moc:
		{
			utils.SaveBytes(filepath.Join(path, "moc_cgo_desktop_darwin_amd64.go"), bb.Bytes())
		}

	case parser.State.Minimal:
		{
			utils.SaveBytes(filepath.Join(path, "minimal_cgo_desktop_darwin_amd64.go"), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(filepath.Join(path, "cgo_desktop_darwin_amd64.go"), bb.Bytes())
		}
	}
}

func qmakeToCgoIos(module, mocPath, buildTarget string) string {
	var path = utils.GoQtPkgPath(strings.ToLower(module))
	if mocPath != "" {
		path = mocPath
	}

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build ${BUILDTARGET}%v\n\n", func() string {
		if parser.State.Minimal {
			return ",minimal"
		}
		if parser.State.Moc || parser.State.Rcc {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if parser.State.Moc {
			return parser.State.Module
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	for _, l := range strings.Split(utils.Load(filepath.Join(path, "Makefile")), "\n") {
		switch {
		case
			strings.HasPrefix(l, "DEFINES"):
			{
				if strings.Contains(l, "-DQT_TESTCASE_BUILDDIR") {
					l = strings.Split(l, "-DQT_TESTCASE_BUILDDIR")[0]
				}

				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CFLAGS"):
			{
				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CXXFLAGS"),
			strings.HasPrefix(l, "INCPATH"):
			{
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "LFLAGS"),
			strings.HasPrefix(l, "LIBS"):
			{

				fmt.Fprintf(bb, "#cgo LDFLAGS: %v\n", strings.Split(l, " = ")[1])
			}
		}
	}

	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = bb.String()
	for _, bl := range []string{"DEFINES", "SUBLIBS", "EXPORT_QMAKE_XARCH_CFLAGS", "EXPORT_QMAKE_XARCH_LFLAGS", "EXPORT_ARCH_ARGS", "-fvisibility=hidden", "-fembed-bitcode"} {
		switch bl {
		case "EXPORT_ARCH_ARGS":
			{
				tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", bl), "-arch x86_64", -1)
			}

		case "EXPORT_QMAKE_XARCH_LFLAGS":
			{
				var todo = "-Xarch_x86_64 -mios-simulator-version-min=7.0 -Xarch_x86_64 -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator10.2.sdk"
				tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", bl), todo, -1)
			}

		case "EXPORT_QMAKE_XARCH_CFLAGS":
			{
				var todo = "-Xarch_x86_64 -mios-simulator-version-min=7.0 -Xarch_x86_64 -isysroot/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator10.2.sdk"
				tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", bl), todo, -1)
			}
		}
		tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", bl), "", -1)
		tmp = strings.Replace(tmp, fmt.Sprintf("%v", bl), "", -1)
	}
	bb.Reset()
	bb.WriteString(tmp)

	if module == "build_ios" {
		return bb.String()
	}

	var prefix string
	switch {
	case parser.State.Rcc:
		{
			prefix = "rcc_"
		}

	case parser.State.Moc:
		{
			prefix = "moc_"
		}

	case parser.State.Minimal:
		{
			prefix = "minimal_"
		}
	}

	tmp = strings.Replace(bb.String(), "${BUILDTARGET}", "ios", -1)

	//amd64
	utils.Save(filepath.Join(path, prefix+"cgo_ios_simulator_darwin_amd64.go"), tmp)

	//386
	utils.Save(filepath.Join(path, prefix+"cgo_ios_simulator_darwin_386.go"), strings.Replace(tmp, "x86_64", "i386", -1))

	tmp = strings.Replace(bb.String(), "${BUILDTARGET}", "ios", -1)
	tmp = strings.Replace(tmp, "iPhoneSimulator", "iPhoneOS", -1)
	tmp = strings.Replace(tmp, "ios-simulator", "iphoneos", -1)

	//arm64
	utils.Save(filepath.Join(path, prefix+"cgo_ios_darwin_arm64.go"), strings.Replace(tmp, "x86_64", "arm64", -1))

	//armv7
	utils.Save(filepath.Join(path, prefix+"cgo_ios_darwin_arm.go"), strings.Replace(tmp, "x86_64", "armv7", -1))

	return ""
}

func qmakeToCgoAndroid(module, mocPath, buildTarget string) {
	var path = utils.GoQtPkgPath(strings.ToLower(module))
	if mocPath != "" {
		path = mocPath
	}

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build android%v\n\n", func() string {
		if parser.State.Minimal {
			return ",minimal"
		}
		if parser.State.Moc || parser.State.Rcc {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if parser.State.Moc {
			return parser.State.Module
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	for _, l := range strings.Split(utils.Load(filepath.Join(path, "Makefile")), "\n") {
		switch {
		case
			strings.HasPrefix(l, "DEFINES"):
			{
				if strings.Contains(l, "-DQT_TESTCASE_BUILDDIR") {
					l = strings.Split(l, "-DQT_TESTCASE_BUILDDIR")[0]
				}

				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CFLAGS"):
			{
				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CXXFLAGS"),
			strings.HasPrefix(l, "INCPATH"):
			{
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "LFLAGS"),
			strings.HasPrefix(l, "LIBS"):
			{

				fmt.Fprintf(bb, "#cgo LDFLAGS: %v\n", strings.Split(l, " = ")[1])
			}
		}
	}

	fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,--allow-shlib-undefined\n")

	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = strings.Replace(bb.String(), "\\", "/", -1)

	for _, bl := range []string{"DEFINES", "SUBLIBS"} {
		tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", bl), "", -1)
	}

	bb.Reset()
	bb.WriteString(tmp)

	switch {
	case parser.State.Rcc:
		{
			utils.SaveBytes(filepath.Join(mocPath, "rcc_cgo_android_linux_arm.go"), bb.Bytes())
		}

	case parser.State.Moc:
		{
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_android_linux_arm.go"), bb.Bytes())
		}

	case parser.State.Minimal:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_android_linux_arm.go"), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_android_linux_arm.go"), bb.Bytes())
		}
	}
}

func qmakeToCgoWindows(module, mocPath, buildTarget string) {
	var path = utils.GoQtPkgPath(strings.ToLower(module))
	if mocPath != "" {
		path = mocPath
	}

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if parser.State.Minimal {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if parser.State.Moc || parser.State.Rcc {
	} else {
		fmt.Fprint(bb, "// +build !minimal\n\n")
	}

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if parser.State.Moc {
			return parser.State.Module
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	for _, l := range strings.Split(utils.Load(filepath.Join(path, "Makefile.Release")), "\n") {
		switch {
		case
			strings.HasPrefix(l, "DEFINES"):
			{
				if strings.Contains(l, "-DQT_TESTCASE_BUILDDIR") {
					l = strings.Split(l, "-DQT_TESTCASE_BUILDDIR")[0]
				}

				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CFLAGS"):
			{
				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CXXFLAGS"),
			strings.HasPrefix(l, "INCPATH"):
			{
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "LFLAGS"),
			strings.HasPrefix(l, "LIBS"):
			{
				if !utils.QT_MXE_STATIC() {
					for _, path := range []string{
						filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "mingw53_32"),
						filepath.Join("/usr", "lib", "mxe", "usr", "i686-w64-mingw32.shared", "qt5"),
						filepath.Join("/usr", "lib", "mxe", "usr", "x86_64-w64-mingw32.shared", "qt5"),
						utils.QT_MSYS2_DIR(),
					} {
						path = strings.Replace(filepath.Join(path, "lib", "lib"), "\\", "/", -1)
						if strings.Contains(l, path) {
							var cleaned []string
							for _, s := range strings.Split(l, " ") {
								if strings.HasPrefix(s, path) && (strings.HasSuffix(s, ".a") || strings.HasSuffix(s, ".dll")) {
									s = strings.Replace(s, path, "-l", -1)
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
	}

	fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,--allow-multiple-definition\n")

	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = strings.Replace(bb.String(), "\\", "/", -1)

	for _, bl := range []string{"DEFINES"} {
		tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", bl), "", -1)
	}

	if utils.UseMsys2() && utils.QT_MSYS2_ARCH() == "amd64" {
		tmp = strings.Replace(tmp, " -Wa,-mbig-obj ", " ", -1)
	}

	if (utils.UseMsys2() && utils.QT_MSYS2_ARCH() == "amd64") || utils.QT_MXE_ARCH() == "amd64" {
		tmp = strings.Replace(tmp, " -Wl,-s ", " ", -1)
	}

	bb.Reset()
	bb.WriteString(tmp)

	var suffix = "386"
	if utils.QT_MXE_ARCH() == "amd64" || (utils.UseMsys2() && utils.QT_MSYS2_ARCH() == "amd64") {
		suffix = "amd64"
	}

	switch {
	case parser.State.Rcc:
		{
			utils.SaveBytes(filepath.Join(mocPath, fmt.Sprintf("rcc_cgo_desktop_windows_%v.go", suffix)), bb.Bytes())
		}

	case parser.State.Moc:
		{
			utils.SaveBytes(filepath.Join(mocPath, fmt.Sprintf("moc_cgo_desktop_windows_%v.go", suffix)), bb.Bytes())
		}

	case parser.State.Minimal:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), fmt.Sprintf("minimal_cgo_desktop_windows_%v.go", suffix)), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), fmt.Sprintf("cgo_desktop_windows_%v.go", suffix)), bb.Bytes())
		}
	}
}

func qmakeToCgoLinux(module, mocPath, buildTarget string) {
	var path = utils.GoQtPkgPath(strings.ToLower(module))
	if mocPath != "" {
		path = mocPath
	}

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if parser.State.Minimal {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if parser.State.Moc || parser.State.Rcc {
	} else {
		fmt.Fprint(bb, "// +build !minimal\n\n")
	}

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if parser.State.Moc {
			return parser.State.Module
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	for _, l := range strings.Split(utils.Load(filepath.Join(path, "Makefile")), "\n") {
		switch {
		case
			strings.HasPrefix(l, "DEFINES"):
			{
				if strings.Contains(l, "-DQT_TESTCASE_BUILDDIR") {
					l = strings.Split(l, "-DQT_TESTCASE_BUILDDIR")[0]
				}

				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CFLAGS"):
			{
				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CXXFLAGS"),
			strings.HasPrefix(l, "INCPATH"):
			{
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "LFLAGS"),
			strings.HasPrefix(l, "LIBS"):
			{

				fmt.Fprintf(bb, "#cgo LDFLAGS: %v\n", strings.Split(l, " = ")[1])
			}
		}
	}

	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = strings.Replace(bb.String(), "\\", "/", -1)

	for _, bl := range []string{"DEFINES", "SUBLIBS"} {
		tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", bl), "", -1)
	}

	bb.Reset()
	bb.WriteString(tmp)

	switch {
	case parser.State.Rcc:
		{
			utils.SaveBytes(filepath.Join(mocPath, "rcc_cgo_desktop_linux_amd64.go"), bb.Bytes())
		}

	case parser.State.Moc:
		{
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_desktop_linux_amd64.go"), bb.Bytes())
		}

	case parser.State.Minimal:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_desktop_linux_amd64.go"), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_desktop_linux_amd64.go"), bb.Bytes())
		}
	}
}

func qmakeToCgoSailfish(module, mocPath, buildTarget string) {
	var path = utils.GoQtPkgPath(strings.ToLower(module))
	if mocPath != "" {
		path = mocPath
	}

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build ${BUILDTARGET}%v\n\n", func() string {
		if parser.State.Minimal {
			return ",minimal"
		}
		if parser.State.Moc {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if parser.State.Moc {
			return parser.State.Module
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	for _, l := range strings.Split(utils.Load(filepath.Join(path, "Makefile")), "\n") {
		switch {
		case
			strings.HasPrefix(l, "DEFINES"):
			{
				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CFLAGS"):
			{
				fmt.Fprintf(bb, "#cgo CFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "CXXFLAGS"),
			strings.HasPrefix(l, "INCPATH"):
			{
				fmt.Fprintf(bb, "#cgo CXXFLAGS: %v\n", strings.Split(l, " = ")[1])
			}

		case
			strings.HasPrefix(l, "LFLAGS"),
			strings.HasPrefix(l, "LIBS"):
			{

				fmt.Fprintf(bb, "#cgo LDFLAGS: %v\n", strings.Split(l, " = ")[1])
			}
		}
	}

	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = strings.Replace(bb.String(), "\\", "/", -1)

	for _, bl := range []string{"DEFINES", "SUBLIBS"} {
		tmp = strings.Replace(tmp, fmt.Sprintf("$(%v)", bl), "", -1)
	}

	bb.Reset()
	bb.WriteString(tmp)

	tmp = strings.Replace(bb.String(), "${BUILDTARGET}", "sailfish_emulator", -1)

	switch {
	case parser.State.Moc:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_sailfish_emulator_linux_386.go"), tmp)
		}

	case parser.State.Minimal:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_sailfish_emulator_linux_386.go"), tmp)
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_sailfish_emulator_linux_386.go"), tmp)
		}
	}

	tmp = strings.Replace(bb.String(), "${BUILDTARGET}", "sailfish", -1)
	tmp = strings.Replace(tmp, "-m32 -msse -msse2 -march=i686 -mfpmath=sse -mtune=generic -fno-omit-frame-pointer -fasynchronous-unwind-tables", "-fmessage-length=0 -march=armv7-a -mfloat-abi=hard -mfpu=neon -mthumb -Wno-psabi", -1)
	tmp = strings.Replace(tmp, "i486", "armv7hl", -1)

	switch {
	case parser.State.Moc:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_sailfish_linux_arm.go"), tmp)
		}

	case parser.State.Minimal:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_sailfish_linux_arm.go"), tmp)
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_sailfish_linux_arm.go"), tmp)
		}
	}
}
