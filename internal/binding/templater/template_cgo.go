package templater

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func CgoTemplate(module, mocPath string) {

	if !(strings.Contains(module, "droid") || strings.Contains(module, "fish")) {
		cgoDarwin(module, mocPath)
		if runtime.GOOS == "windows" {
			if utils.UseMsys2() {
				cgoWindowsMsys2(module, mocPath)
			} else {
				cgoWindows(module, mocPath)
			}
		} else {
			cgoWindowsForLinux(module, mocPath)
		}
		if utils.UsePkgConfig() {
			cgoLinuxPkgConfig(module, mocPath)
		} else {
			cgoLinux(module, mocPath)
		}
		cgoSailfish(module, mocPath)
		cgoRaspberryPi1(module, mocPath)
		cgoRaspberryPi2(module, mocPath)
		cgoRaspberryPi3(module, mocPath)

		if runtime.GOOS == "darwin" {
			cgoIos(module, mocPath)
		}
	}

	if strings.Contains(module, "fish") {
		cgoSailfish(module, mocPath)
	} else {
		cgoAndroid(module, mocPath)
	}
}

func cgoDarwin(module, mocPath string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build !ios%v\n\n", func() string {
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

	fmt.Fprintf(bb, "#cgo CFLAGS: -pipe -O2 -isysroot %v/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/%v -mmacosx-version-min=10.8 -Wall -W -fPIC\n", utils.XCODE_DIR(), utils.MACOS_SDK_DIR())
	fmt.Fprintf(bb, "#cgo CXXFLAGS: -pipe -stdlib=libc++ -O2 -std=gnu++11 -isysroot %v/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/%v -mmacosx-version-min=10.8 -Wall -W -fPIC\n", utils.XCODE_DIR(), utils.MACOS_SDK_DIR())

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/mkspecs/macx-clang -F%v/lib\n", utils.QT_DARWIN_DIR(), utils.QT_DARWIN_DIR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			fmt.Fprintf(bb, " -I%v/include/Qt%v", utils.QT_DARWIN_DIR(), m)
		} else {
			fmt.Fprintf(bb, " -I%v/lib/Qt%v.framework/Headers", utils.QT_DARWIN_DIR(), m)
		}
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/%v/System/Library/Frameworks/OpenGL.framework/Headers -I%v/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/%v/System/Library/Frameworks/AGL.framework/Headers\n", utils.XCODE_DIR(), utils.MACOS_SDK_DIR(), utils.XCODE_DIR(), utils.MACOS_SDK_DIR())
	fmt.Fprintf(bb, "\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,%v/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/%v -mmacosx-version-min=10.8 -Wl,-rpath,%v/lib\n", utils.XCODE_DIR(), utils.MACOS_SDK_DIR(), utils.QT_DARWIN_DIR())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -F%v/lib", utils.QT_DARWIN_DIR())

	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			fmt.Fprintf(bb, " -L%v/lib -lQt5%v", utils.QT_DARWIN_DIR(), m)
		} else {
			if m != "UiPlugin" {
				fmt.Fprintf(bb, " -framework Qt%v", m)
			}
		}
	}

	fmt.Fprint(bb, " -framework DiskArbitration -framework IOKit -framework OpenGL -framework AGL\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	switch {
	case parser.State.Moc:
		{
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_desktop_darwin_amd64.go"), bb.Bytes())
		}

	case parser.State.Minimal:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_desktop_darwin_amd64.go"), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_desktop_darwin_amd64.go"), bb.Bytes())
		}
	}
}

func cgoWindows(module, mocPath string) {
	var (
		bb     = new(bytes.Buffer)
		libs   = cleanLibs(module)
		QT_DIR = func() string { return strings.Replace(utils.QT_DIR(), "\\", "/", -1) }
	)
	defer bb.Reset()

	if parser.State.Minimal {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if parser.State.Moc {
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

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -O2 -std=gnu++11 -frtti -Wall -Wextra -fexceptions -mthreads\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DUNICODE -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, " -DQT_NEEDS_QMAIN")
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/%v/mingw53_32/include -I%v/%v/mingw53_32/mkspecs/win32-g++\n", QT_DIR(), utils.QT_VERSION_MAJOR(), QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/%v/mingw53_32/include/Qt%v", QT_DIR(), utils.QT_VERSION_MAJOR(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v/%v/mingw53_32/lib", QT_DIR(), utils.QT_VERSION_MAJOR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lmingw32 -lqtmain -lshell32\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	switch {
	case parser.State.Moc:
		{
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_desktop_windows_386.go"), bb.Bytes())
		}

	case parser.State.Minimal:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_desktop_windows_386.go"), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_desktop_windows_386.go"), bb.Bytes())
		}
	}
}

func cgoWindowsForLinux(module, mocPath string) {
	var (
		bb     = new(bytes.Buffer)
		libs   = cleanLibs(module)
		QT_DIR = func() string { return "/usr/lib/mxe/usr/i686-w64-mingw32.shared/qt5" }
	)
	defer bb.Reset()

	if parser.State.Minimal {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if parser.State.Moc {
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

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -O2 -std=gnu++11 -frtti -Wall -Wextra -fexceptions -mthreads\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DUNICODE -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, " -DQT_NEEDS_QMAIN")
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/include -I%v/mkspecs/win32-g++\n", QT_DIR(), QT_DIR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/include/Qt%v", QT_DIR(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -lglu32 -lopengl32 -lgdi32 -luser32 -L%v/lib", QT_DIR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lmingw32 -lqtmain -lshell32\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp64 = strings.Replace(bb.String(), "i686", "x86_64", -1)
	tmp64 = strings.Replace(tmp64, " -Wl,-s ", " ", -1)

	switch {
	case parser.State.Moc:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_desktop_windows_amd64.go"), tmp64)
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_desktop_windows_386.go"), bb.Bytes())
		}

	case parser.State.Minimal:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_desktop_windows_amd64.go"), tmp64)
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_desktop_windows_386.go"), bb.Bytes())
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_desktop_windows_amd64.go"), tmp64)
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_desktop_windows_386.go"), bb.Bytes())
		}
	}
}

func cgoWindowsMsys2(module, mocPath string) {
	var (
		bb        = new(bytes.Buffer)
		libs      = cleanLibs(module)
		MSYS2_DIR = func() string { return strings.Replace(utils.QT_MSYS2_DIR(), "\\", "/", -1) }
	)
	defer bb.Reset()

	if parser.State.Minimal {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if parser.State.Moc {
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

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -march=i686 -mtune=core2 -Wa,-mbig-obj -O2 -Wall -Wextra\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -march=i686 -mtune=core2 -Wa,-mbig-obj -O2 -std=gnu++0x -frtti -Wall -Wextra -fexceptions -mthreads\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DUNICODE -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, " -DQT_NEEDS_QMAIN")
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/include -I%v/share/qt5/mkspecs/win32-g++\n", MSYS2_DIR(), MSYS2_DIR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/include/Qt%v", MSYS2_DIR(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v/lib", MSYS2_DIR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lglu32 -lopengl32 -lgdi32 -luser32 -lmingw32 -lqtmain\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp64 = strings.Replace(bb.String(), "-march=i686", "-march=nocona", -1)
	tmp64 = strings.Replace(tmp64, " -Wa,-mbig-obj ", " ", -1)
	tmp64 = strings.Replace(tmp64, " -Wl,-s ", " ", -1)

	switch {
	case parser.State.Moc:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_desktop_windows_amd64.go"), tmp64)
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_desktop_windows_386.go"), bb.Bytes())
		}

	case parser.State.Minimal:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_desktop_windows_amd64.go"), tmp64)
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_desktop_windows_386.go"), bb.Bytes())
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_desktop_windows_amd64.go"), tmp64)
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_desktop_windows_386.go"), bb.Bytes())
		}
	}
}

func cgoLinux(module, mocPath string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	if parser.State.Minimal {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if parser.State.Moc {
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

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -O2 -Wall -W -D_REENTRANT -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -O2 -std=gnu++11 -Wall -W -D_REENTRANT -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/%v/gcc_64/include -I%v/%v/gcc_64/mkspecs/linux-g++\n", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/%v/gcc_64/include/Qt%v", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,%v/%v/gcc_64 -Wl,-rpath,%v/%v/gcc_64/lib\n", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v/%v/gcc_64/lib -L/usr/lib64", utils.QT_DIR(), utils.QT_VERSION_MAJOR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	switch {
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

func cgoLinuxPkgConfig(module, mocPath string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	if parser.State.Minimal {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if parser.State.Moc {
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

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -O2 -Wall -W -D_REENTRANT -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -O2 -std=gnu++11 -Wall -W -D_REENTRANT -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	var (
		includeDir = strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=includedir", "Qt5Core"), "cgo.LinuxPkgConfig_includeDir"))
		libDir     = strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=libdir", "Qt5Core"), "cgo.LinuxPkgConfig_libDir"))
		miscDir    = utils.QT_MISC_DIR()
	)

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v -I%v/mkspecs/linux-g++\n", includeDir, miscDir)

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/Qt%v", includeDir, m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,%v -Wl,-rpath,%v/lib\n", libDir, libDir)

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v -L/usr/lib64", libDir)
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	switch {
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

func cgoAndroid(module, mocPath string) {
	var (
		bb              = new(bytes.Buffer)
		libs            = cleanLibs(module)
		QT_DIR          = func() string { return strings.Replace(utils.QT_DIR(), "\\", "/", -1) }
		ANDROID_NDK_DIR = func() string { return strings.Replace(utils.ANDROID_NDK_DIR(), "\\", "/", -1) }
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build android%v\n\n", func() string {
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

	fmt.Fprint(bb, "#cgo CFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -std=c++11 -O2 -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/%v/android_armv7/include -isystem %v/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem %v/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem %v/platforms/android-9/arch-arm/usr/include -I%v/%v/android_armv7/mkspecs/android-g++\n", QT_DIR(), utils.QT_VERSION_MAJOR(), ANDROID_NDK_DIR(), ANDROID_NDK_DIR(), ANDROID_NDK_DIR(), QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/%v/android_armv7/include/Qt%v", QT_DIR(), utils.QT_VERSION_MAJOR(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: --sysroot=%v/platforms/android-9/arch-arm -Wl,-rpath=%v/%v/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition -Wl,--allow-shlib-undefined\n", ANDROID_NDK_DIR(), QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L%v/platforms/android-9/arch-arm/usr/lib -L%v/%v/android_armv7/lib -L/opt/android/ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android/ndk/platforms/android-9/arch-arm/usr/lib", ANDROID_NDK_DIR(), ANDROID_NDK_DIR(), QT_DIR(), utils.QT_VERSION_MAJOR())

	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lgnustl_shared -llog -lz -lm -ldl -lc -lgcc\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	switch {
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

func cgoIos(module, mocPath string) string {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
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

	fmt.Fprintf(bb, "#cgo CFLAGS: -pipe -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -O2 -isysroot %v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/%v -mios-simulator-version-min=7.0 -arch x86_64 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC\n", utils.XCODE_DIR(), utils.IPHONESIMULATOR_SDK_DIR())
	fmt.Fprintf(bb, "#cgo CXXFLAGS: -pipe -stdlib=libc++ -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -fvisibility-inlines-hidden -Wno-non-virtual-dtor -Wno-overloaded-virtual -Wno-exit-time-destructors -O2 -std=gnu++11 -isysroot %v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/%v -mios-simulator-version-min=7.0 -arch x86_64 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC\n", utils.XCODE_DIR(), utils.IPHONESIMULATOR_SDK_DIR())

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DDARWIN_NO_CARBON -DQT_NO_PRINTER -DQT_NO_PRINTDIALOG -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/%v/ios/mkspecs/macx-ios-clang/ios -I%v/%v/ios/include -I%v/%v/ios/mkspecs/macx-ios-clang\n", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/%v/ios/include/Qt%v", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,%v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/%v -mios-simulator-version-min=7.0 -arch x86_64\n", utils.XCODE_DIR(), utils.IPHONESIMULATOR_SDK_DIR())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v/%v/ios/plugins/platforms -lqios_iphonesimulator -framework Foundation -framework UIKit -framework QuartzCore -framework AssetsLibrary -L%v/%v/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework CoreText -framework CoreGraphics -framework OpenGLES -lqtfreetype_iphonesimulator -framework Security -framework SystemConfiguration -framework CoreBluetooth -L%v/%v/ios/plugins/imageformats -lqdds_iphonesimulator -lqicns_iphonesimulator -lqico_iphonesimulator -lqtga_iphonesimulator -lqtiff_iphonesimulator -lqwbmp_iphonesimulator -lqwebp_iphonesimulator -lqtharfbuzzng_iphonesimulator -lz -lqtpcre_iphonesimulator -lm", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v_iphonesimulator", m)
		}
	}

	if !strings.Contains(bb.String(), "-lQt5Gui_iphonesimulator") {
		fmt.Fprint(bb, " -lQt5Gui_iphonesimulator")
	}

	switch module {
	case
		"Multimedia":
		{
			fmt.Fprint(bb, " -lQt5OpenGL_iphonesimulator")
		}

	case "TestLib":
		{
			fmt.Fprintf(bb, " -F%v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest", utils.XCODE_DIR())
		}

	case "Qml", "WebChannel", "Quick":
		{
			if module != "Quick" {
				fmt.Fprint(bb, " -lQt5Quick_iphonesimulator -lQt5QuickParticles_iphonesimulator -lQt5QuickTest_iphonesimulator -lQt5QuickWidgets_iphonesimulator")
			}
			fmt.Fprintf(bb, " -L%v/%v/ios/plugins/qmltooling -lqmldbg_debugger_iphonesimulator -lqmldbg_inspector_iphonesimulator -lqmldbg_local_iphonesimulator -lqmldbg_native_iphonesimulator -lqmldbg_profiler_iphonesimulator -lqmldbg_quickprofiler_iphonesimulator -lqmldbg_server_iphonesimulator -lQt5PacketProtocol_iphonesimulator -lqmldbg_tcp_iphonesimulator", utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		}

	case "Purchasing":
		{
			fmt.Fprint(bb, " -framework StoreKit")
		}
	}

	if module == "build_ios" {
		fmt.Fprint(bb, " -lQt5OpenGL_iphonesimulator")
		fmt.Fprintf(bb, " -F%v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest", utils.XCODE_DIR())
		fmt.Fprint(bb, " -lQt5Quick_iphonesimulator -lQt5QuickParticles_iphonesimulator -lQt5QuickTest_iphonesimulator -lQt5QuickWidgets_iphonesimulator")
		fmt.Fprintf(bb, " -L%v/%v/ios/plugins/qmltooling -lqmldbg_debugger_iphonesimulator -lqmldbg_inspector_iphonesimulator -lqmldbg_local_iphonesimulator -lqmldbg_native_iphonesimulator -lqmldbg_profiler_iphonesimulator -lqmldbg_quickprofiler_iphonesimulator -lqmldbg_server_iphonesimulator -lQt5PacketProtocol_iphonesimulator -lqmldbg_tcp_iphonesimulator", utils.QT_DIR(), utils.QT_VERSION_MAJOR())

		fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtQuick.2 -lqtquick2plugin_iphonesimulator -L%v/%v/ios/qml/QtQuick/Layouts -lqquicklayoutsplugin_iphonesimulator -L%v/%v/ios/qml/QtQuick/Dialogs -ldialogplugin_iphonesimulator -L%v/%v/ios/qml/QtQuick/Controls -lqtquickcontrolsplugin_iphonesimulator -L%v/%v/ios/qml/Qt/labs/folderlistmodel -lqmlfolderlistmodelplugin_iphonesimulator -L%v/%v/ios/qml/Qt/labs/settings -lqmlsettingsplugin_iphonesimulator -L%v/%v/ios/qml/QtQuick/Dialogs/Private -ldialogsprivateplugin_iphonesimulator -L%v/%v/ios/qml/QtQuick/Window.2 -lwindowplugin_iphonesimulator -L%v/%v/ios/qml/QtQml/Models.2 -lmodelsplugin_iphonesimulator -L%v/%v/ios/qml/QtQuick/Extras -lqtquickextrasplugin_iphonesimulator -L%v/%v/ios/qml/QtGraphicalEffects/private -lqtgraphicaleffectsprivate_iphonesimulator", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtQuick/Controls.2 -lqtquickcontrols2plugin_iphonesimulator -L%v/%v/ios/qml/QtQuick/Controls.2/Material -lqtquickcontrols2materialstyleplugin_iphonesimulator -L%v/%v/ios/qml/QtQuick/Controls.2/Universal -lqtquickcontrols2universalstyleplugin_iphonesimulator -lQt5QuickControls2_iphonesimulator", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtQuick/Templates.2 -lqtquicktemplates2plugin_iphonesimulator -lQt5QuickTemplates2_iphonesimulator -L%v/%v/ios/qml/QtGraphicalEffects -lqtgraphicaleffectsplugin_iphonesimulator", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())

		fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtMultimedia -ldeclarative_multimedia_iphonesimulator -lQt5MultimediaQuick_p_iphonesimulator", utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		fmt.Fprintf(bb, " -L%v/%v/ios/plugins/mediaservice -lqtmedia_audioengine_iphonesimulator -L%v/%v/ios/plugins/audio -lqtaudio_coreaudio_iphonesimulator -L%v/%v/ios/plugins/playlistformats -lqtmultimedia_m3u_iphonesimulator -lQt5Multimedia_iphonesimulator", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		fmt.Fprint(bb, " -lqavfcamera_iphonesimulator -framework AudioToolbox -framework CoreAudio -framework AVFoundation -framework CoreMedia -framework CoreVideo -lqavfmediaplayer_iphonesimulator -lQt5MultimediaWidgets_iphonesimulator")
	}

	fmt.Fprint(bb, " -lQt5PlatformSupport_iphonesimulator\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	if module == "build_ios" {
		return bb.String()
	}

	var path, prefix = func() (string, string) {
		switch {
		case parser.State.Moc:
			{
				return mocPath, "moc_"
			}

		case parser.State.Minimal:
			{
				return utils.GoQtPkgPath(strings.ToLower(module)), "minimal_"
			}

		default:
			{
				return utils.GoQtPkgPath(strings.ToLower(module)), ""
			}
		}
	}()

	var tmp = strings.Replace(bb.String(), "${BUILDTARGET}", "ios", -1)

	//amd64
	utils.Save(filepath.Join(path, prefix+"cgo_ios_simulator_darwin_amd64.go"), tmp)

	//386
	utils.Save(filepath.Join(path, prefix+"cgo_ios_simulator_darwin_386.go"), strings.Replace(tmp, "x86_64", "i386", -1))

	tmp = strings.Replace(bb.String(), "${BUILDTARGET}", "ios", -1)
	tmp = strings.Replace(tmp, "_iphonesimulator", "", -1)
	tmp = strings.Replace(tmp, "iPhoneSimulator", "iPhoneOS", -1)
	tmp = strings.Replace(tmp, "ios-simulator", "iphoneos", -1)

	//arm64
	utils.Save(filepath.Join(path, prefix+"cgo_ios_darwin_arm64.go"), strings.Replace(tmp, "x86_64", "arm64", -1))

	//armv7
	utils.Save(filepath.Join(path, prefix+"cgo_ios_darwin_arm.go"), strings.Replace(tmp, "x86_64", "armv7", -1))

	return ""
}

func cgoSailfish(module, mocPath string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
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

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -Wformat -Wformat-security -m32 -msse -msse2 -march=i686 -mfpmath=sse -mtune=generic -fno-omit-frame-pointer -fasynchronous-unwind-tables -fPIC -fvisibility=hidden -fvisibility-inlines-hidden -Wall -W -D_REENTRANT -fPIE\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -Wformat -Wformat-security -m32 -msse -msse2 -march=i686 -mfpmath=sse -mtune=generic -fno-omit-frame-pointer -fasynchronous-unwind-tables -fPIC -fvisibility=hidden -fvisibility-inlines-hidden -Wall -W -D_REENTRANT -fPIE\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -I/srv/mer/targets/SailfishOS-i486/usr/share/qt5/mkspecs/linux-g++ -I/srv/mer/targets/SailfishOS-i486/usr/include -I/srv/mer/targets/SailfishOS-i486/usr/include/sailfishapp -I/srv/mer/targets/SailfishOS-i486/usr/include/mdeclarativecache5 -I/srv/mer/targets/SailfishOS-i486/usr/include/qt5\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I/srv/mer/targets/SailfishOS-i486/usr/include/qt5/Qt%v", m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath-link,/srv/mer/targets/SailfishOS-i486/usr/lib -Wl,-rpath-link,/srv/mer/targets/SailfishOS-i486/lib\n")

	fmt.Fprint(bb, "#cgo LDFLAGS: -rdynamic -L/srv/mer/targets/SailfishOS-i486/usr/lib -L/srv/mer/targets/SailfishOS-i486/lib -lsailfishapp -lmdeclarativecache5")
	for _, m := range libs {
		if m != "UiPlugin" {
			if IsWhiteListedSailfishLib(m) {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = strings.Replace(bb.String(), "${BUILDTARGET}", "sailfish_emulator", -1)

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

func cgoRaspberryPi1(module, mocPath string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build rpi1%v\n\n", func() string {
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

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -marm -mfpu=vfp -mtune=arm1176jzf-s -march=armv6zk -mabi=aapcs-linux -mfloat-abi=hard --sysroot=${RPI1_SYSROOT_DIR} -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -marm -mfpu=vfp -mtune=arm1176jzf-s -march=armv6zk -mabi=aapcs-linux -mfloat-abi=hard --sysroot=${RPI1_SYSROOT_DIR} -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/%v/rpi1/include -I${RPI1_SYSROOT_DIR}/opt/vc/include -I${RPI1_SYSROOT_DIR}/opt/vc/include/interface/vcos/pthreads -I${RPI1_SYSROOT_DIR}/opt/vc/include/interface/vmcs_host/linux -I%v/%v/rpi1/mkspecs/devices/linux-rasp-pi-g++\n", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/%v/rpi1/include/Qt%v", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-rpath-link,${RPI1_SYSROOT_DIR}/opt/vc/lib -Wl,-rpath-link,${RPI1_SYSROOT_DIR}/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,${RPI1_SYSROOT_DIR}/lib/arm-linux-gnueabihf -Wl,-rpath-link,%v/%v/rpi1/lib -mfloat-abi=hard --sysroot=${RPI1_SYSROOT_DIR} -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin\n", utils.QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L${RPI1_SYSROOT_DIR}/opt/vc/lib -L%v/%v/rpi1/lib", utils.QT_DIR(), utils.QT_VERSION_MAJOR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = strings.Replace(bb.String(), "${RPI1_SYSROOT_DIR}", utils.RPI1_SYSROOT_DIR(), -1)

	switch {
	case parser.State.Moc:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_rpi1_linux_arm.go"), tmp)
		}

	case parser.State.Minimal:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_rpi1_linux_arm.go"), tmp)
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_rpi1_linux_arm.go"), tmp)
		}
	}
}

func cgoRaspberryPi2(module, mocPath string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build rpi2%v\n\n", func() string {
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

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -march=armv7-a -marm -mthumb-interwork -mfpu=neon-vfpv4 -mtune=cortex-a7 -mabi=aapcs-linux -mfloat-abi=hard --sysroot=${RPI2_SYSROOT_DIR} -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -march=armv7-a -marm -mthumb-interwork -mfpu=neon-vfpv4 -mtune=cortex-a7 -mabi=aapcs-linux -mfloat-abi=hard --sysroot=${RPI2_SYSROOT_DIR} -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/%v/rpi2/include -I${RPI2_SYSROOT_DIR}/opt/vc/include -I${RPI2_SYSROOT_DIR}/opt/vc/include/interface/vcos/pthreads -I${RPI2_SYSROOT_DIR}/opt/vc/include/interface/vmcs_host/linux -I%v/%v/rpi2/mkspecs/devices/linux-rasp-pi2-g++\n", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/%v/rpi2/include/Qt%v", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-rpath-link,${RPI2_SYSROOT_DIR}/opt/vc/lib -Wl,-rpath-link,${RPI2_SYSROOT_DIR}/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,${RPI2_SYSROOT_DIR}/lib/arm-linux-gnueabihf -Wl,-rpath-link,%v/%v/rpi2/lib -mfloat-abi=hard --sysroot=${RPI2_SYSROOT_DIR} -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin\n", utils.QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L${RPI2_SYSROOT_DIR}/opt/vc/lib -L%v/%v/rpi2/lib", utils.QT_DIR(), utils.QT_VERSION_MAJOR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = strings.Replace(bb.String(), "${RPI2_SYSROOT_DIR}", utils.RPI2_SYSROOT_DIR(), -1)

	switch {
	case parser.State.Moc:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_rpi2_linux_arm.go"), tmp)
		}

	case parser.State.Minimal:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_rpi2_linux_arm.go"), tmp)
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_rpi2_linux_arm.go"), tmp)
		}
	}
}

func cgoRaspberryPi3(module, mocPath string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build rpi3%v\n\n", func() string {
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

	fmt.Fprint(bb, "#cgo CFLAGS: -march=armv8-a+crc -mtune=cortex-a53 -mfpu=crypto-neon-fp-armv8 -pipe -Os -mthumb -mfloat-abi=hard --sysroot=${RPI3_SYSROOT_DIR} -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -march=armv8-a+crc -mtune=cortex-a53 -mfpu=crypto-neon-fp-armv8 -pipe -Os -mthumb -std=c++11 -mfloat-abi=hard --sysroot=${RPI3_SYSROOT_DIR} -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/%v/rpi3/include -I${RPI3_SYSROOT_DIR}/opt/vc/include -I${RPI3_SYSROOT_DIR}/opt/vc/include/interface/vcos/pthreads -I${RPI3_SYSROOT_DIR}/opt/vc/include/interface/vmcs_host/linux -I%v/%v/rpi3/mkspecs/devices/linux-rpi3-g++\n", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/%v/rpi3/include/Qt%v", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-rpath-link,${RPI3_SYSROOT_DIR}/opt/vc/lib -Wl,-rpath-link,${RPI3_SYSROOT_DIR}/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,${RPI3_SYSROOT_DIR}/lib/arm-linux-gnueabihf -Wl,-rpath-link,%v/%v/rpi3/lib -mfloat-abi=hard --sysroot=${RPI3_SYSROOT_DIR} -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin\n", utils.QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L${RPI3_SYSROOT_DIR}/opt/vc/lib -L%v/%v/rpi3/lib", utils.QT_DIR(), utils.QT_VERSION_MAJOR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = strings.Replace(bb.String(), "${RPI3_SYSROOT_DIR}", utils.RPI3_SYSROOT_DIR(), -1)

	switch {
	case parser.State.Moc:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_rpi3_linux_arm.go"), tmp)
		}

	case parser.State.Minimal:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_rpi3_linux_arm.go"), tmp)
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_rpi3_linux_arm.go"), tmp)
		}
	}
}

func cleanLibs(module string) []string {

	if parser.State.Moc || module == "build_ios" {
		return parser.LibDeps[module]
	}

	var out = append([]string{module}, parser.LibDeps[module]...)
	for i, v := range out {
		if v == "TestLib" {
			out[i] = "Test"
		}
		if v == "Speech" {
			out[i] = "TextToSpeech"
		}
	}
	return out
}

func GetiOSClang(buildTarget, buildARM string) []string {
	var tmp = cgoIos("build_ios", "")

	tmp = strings.Split(tmp, "/*")[1]
	tmp = strings.Split(tmp, "*/")[0]

	tmp = strings.Replace(tmp, "#cgo CFLAGS: ", "", -1)
	tmp = strings.Replace(tmp, "#cgo CXXFLAGS: ", "", -1)
	tmp = strings.Replace(tmp, "#cgo LDFLAGS: ", "", -1)
	tmp = strings.Replace(tmp, "\n", " ", -1)

	if buildTarget == "ios" {
		tmp = strings.Replace(tmp, "_iphonesimulator", "", -1)
		tmp = strings.Replace(tmp, "x86_64", "arm64", -1)
		tmp = strings.Replace(tmp, "iPhoneSimulator", "iPhoneOS", -1)
		tmp = strings.Replace(tmp, "ios-simulator", "iphoneos", -1)

		if buildARM == "armv7" {
			tmp = strings.Replace(tmp, "arm64", "armv7", -1)
		}
	}

	return strings.Split(tmp, " ")
}

func IsWhiteListedSailfishLib(name string) bool {
	switch name {
	case "Core", "Quick", "Qml", "Network", "Gui", "Concurrent", "Multimedia", "Sql", "Svg", "XmlPatterns", "Xml", "DBus", "WebKit", "Sensors", "Positioning":
		{
			return true
		}

	default:
		{
			return false
		}
	}
}
