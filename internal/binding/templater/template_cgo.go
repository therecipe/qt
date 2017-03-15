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

func CgoTemplate(module, mocPath, buildTarget string, mode int, pkg string) {
	if utils.QT_QMAKE_CGO() {
		QmakeCgoTemplate(module, mocPath, buildTarget, mode, pkg)
		return
	}

	if !(strings.Contains(module, "droid") || strings.Contains(module, "fish")) {
		cgoDarwin(module, mocPath, mode, pkg)
		if runtime.GOOS == "windows" {
			if utils.QT_MSYS2() {
				cgoWindowsMsys2(module, mocPath, mode, pkg) //TODO: docker
			} else {
				cgoWindows(module, mocPath, mode, pkg) //TODO: docker
			}
		} else {
			cgoWindowsForLinux(module, mocPath, mode, pkg)
		}
		if utils.QT_PKG_CONFIG() {
			cgoLinuxPkgConfig(module, mocPath, mode, pkg) //TODO: docker
		} else {
			cgoLinux(module, mocPath, mode, pkg)
		}
		cgoSailfish(module, mocPath, mode, pkg)
		cgoAsteroid(module, mocPath, mode, pkg)
		cgoRaspberryPi1(module, mocPath, mode, pkg) //TODO: 5.8
		cgoRaspberryPi2(module, mocPath, mode, pkg) //TODO: 5.8
		cgoRaspberryPi3(module, mocPath, mode, pkg) //TODO: 5.8

		if runtime.GOOS == "darwin" {
			cgoIos(module, mocPath, mode, pkg)
		}
	}

	if strings.Contains(module, "fish") {
		cgoSailfish(module, mocPath, mode, pkg)
	} else {
		cgoAndroid(module, mocPath, mode, pkg)
	}
}

func cgoDarwin(module, mocPath string, mode int, pkg string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module, mode)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build !ios%v\n\n", func() string {
		if mode == MINIMAL {
			return ",minimal"
		}
		if mode == MOC {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	fmt.Fprintf(bb, "#cgo CFLAGS: -pipe -O2 -isysroot %v/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/%v -mmacosx-version-min=10.9 -Wall -W -fPIC\n", utils.XCODE_DIR(), utils.MACOS_SDK_DIR())
	fmt.Fprintf(bb, "#cgo CXXFLAGS: -pipe -stdlib=libc++ -O2 -std=gnu++11 -isysroot %v/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/%v -mmacosx-version-min=10.9 -Wall -W -fPIC\n", utils.XCODE_DIR(), utils.MACOS_SDK_DIR())

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

	fmt.Fprintf(bb, "#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,%v/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/%v -mmacosx-version-min=10.9 -Wl,-rpath,%v/lib\n", utils.XCODE_DIR(), utils.MACOS_SDK_DIR(), utils.QT_DARWIN_DIR())

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
	case mode == MOC:
		{
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_darwin_darwin_amd64.go"), bb.Bytes())
		}

	case mode == MINIMAL:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_darwin_darwin_amd64.go"), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_darwin_darwin_amd64.go"), bb.Bytes())
		}
	}
}

func cgoWindows(module, mocPath string, mode int, pkg string) {
	var (
		bb     = new(bytes.Buffer)
		libs   = cleanLibs(module, mode)
		QT_DIR = func() string { return strings.Replace(utils.QT_DIR(), "\\", "/", -1) }
	)
	defer bb.Reset()

	if mode == MINIMAL {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if mode == MOC {
	} else {
		fmt.Fprint(bb, "// +build !minimal\n\n")
	}

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
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

	fmt.Fprintf(bb, "#cgo LDFLAGS: -lglu32 -lopengl32 -lgdi32 -luser32 -L%v/%v/mingw53_32/lib", QT_DIR(), utils.QT_VERSION_MAJOR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lmingw32 -lqtmain -lshell32\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	switch {
	case mode == MOC:
		{
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_windows_windows_386.go"), bb.Bytes())
		}

	case mode == MINIMAL:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_windows_windows_386.go"), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_windows_windows_386.go"), bb.Bytes())
		}
	}
}

func cgoWindowsForLinux(module, mocPath string, mode int, pkg string) {
	var (
		bb     = new(bytes.Buffer)
		libs   = cleanLibs(module, mode)
		QT_DIR = func() string { return filepath.Join(utils.QT_MXE_DIR(), "usr", utils.QT_MXE_TRIPLET(), "qt5") }
	)
	defer bb.Reset()

	if mode == MINIMAL {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if mode == MOC {
	} else {
		fmt.Fprint(bb, "// +build !minimal\n\n")
	}

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
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
	case mode == MOC:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_windows_windows_amd64.go"), tmp64)
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_windows_windows_386.go"), bb.Bytes())
		}

	case mode == MINIMAL:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_windows_windows_amd64.go"), tmp64)
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_windows_windows_386.go"), bb.Bytes())
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_windows_windows_amd64.go"), tmp64)
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_windows_windows_386.go"), bb.Bytes())
		}
	}
}

func cgoWindowsMsys2(module, mocPath string, mode int, pkg string) {
	var (
		bb        = new(bytes.Buffer)
		libs      = cleanLibs(module, mode)
		MSYS2_DIR = func() string { return strings.Replace(utils.QT_MSYS2_DIR(), "\\", "/", -1) }
	)
	defer bb.Reset()

	if mode == MINIMAL {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if mode == MOC {
	} else {
		fmt.Fprint(bb, "// +build !minimal\n\n")
	}

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
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

	fmt.Fprintf(bb, "#cgo LDFLAGS: -lglu32 -lopengl32 -lgdi32 -luser32 -L%v/lib", MSYS2_DIR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lmingw32 -lqtmain -lshell32\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp64 = strings.Replace(bb.String(), "-march=i686", "-march=nocona", -1)
	tmp64 = strings.Replace(tmp64, " -Wa,-mbig-obj ", " ", -1)
	tmp64 = strings.Replace(tmp64, " -Wl,-s ", " ", -1)

	switch {
	case mode == MOC:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_windows_windows_amd64.go"), tmp64)
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_windows_windows_386.go"), bb.Bytes())
		}

	case mode == MINIMAL:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_windows_windows_amd64.go"), tmp64)
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_windows_windows_386.go"), bb.Bytes())
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_windows_windows_amd64.go"), tmp64)
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_windows_windows_386.go"), bb.Bytes())
		}
	}
}

func cgoLinux(module, mocPath string, mode int, pkg string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module, mode)
	)
	defer bb.Reset()

	if mode == MINIMAL {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if mode == MOC {
	} else {
		fmt.Fprint(bb, "// +build !minimal\n\n")
	}

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
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

	for _, s := range []struct{ pre, dir string }{{"!", utils.QT_DIR()}, {"", "/opt/Qt5.8.0"}} {

		fmt.Fprintf(bb, "#cgo +build %vdocker CXXFLAGS: -I%v/%v/gcc_64/include -isystem /usr/include/libdrm -I%v/%v/gcc_64/mkspecs/linux-g++\n", s.pre, s.dir, utils.QT_VERSION_MAJOR(), s.dir, utils.QT_VERSION_MAJOR())

		fmt.Fprintf(bb, "#cgo +build %vdocker CXXFLAGS:", s.pre)
		for _, m := range libs {
			fmt.Fprintf(bb, " -I%v/%v/gcc_64/include/Qt%v", s.dir, utils.QT_VERSION_MAJOR(), m)
		}
		fmt.Fprint(bb, "\n\n")

		fmt.Fprintf(bb, "#cgo +build %vdocker LDFLAGS: -Wl,-O1 -Wl,-rpath,%v/%v/gcc_64/lib -Wl,-rpath-link,%v/%v/gcc_64/lib\n", s.pre, s.dir, utils.QT_VERSION_MAJOR(), s.dir, utils.QT_VERSION_MAJOR())

		fmt.Fprintf(bb, "#cgo +build %vdocker LDFLAGS: -L%v/%v/gcc_64/lib -L/usr/lib64", s.pre, s.dir, utils.QT_VERSION_MAJOR())
		for _, m := range libs {
			if m != "UiPlugin" {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
		fmt.Fprint(bb, " -lGL -lpthread\n")
	}

	fmt.Fprint(bb, "*/\n")
	fmt.Fprint(bb, "import \"C\"\n")

	switch {
	case mode == MOC:
		{
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_linux_linux_amd64.go"), bb.Bytes())
		}

	case mode == MINIMAL:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_linux_linux_amd64.go"), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_linux_linux_amd64.go"), bb.Bytes())
		}
	}
}

func cgoLinuxPkgConfig(module, mocPath string, mode int, pkg string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module, mode)
	)
	defer bb.Reset()

	if mode == MINIMAL {
		fmt.Fprint(bb, "// +build minimal\n\n")
	} else if mode == MOC {
	} else {
		fmt.Fprint(bb, "// +build !minimal\n\n")
	}

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
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

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v -isystem /usr/include/libdrm -I%v/mkspecs/linux-g++\n", includeDir, miscDir)

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/Qt%v", includeDir, m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,%v/lib -Wl,-rpath-link,%v/lib\n", libDir, libDir)

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v -L/usr/lib64", libDir)
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	fmt.Fprint(bb, " -lGL -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	switch {
	case mode == MOC:
		{
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_linux_linux_amd64.go"), bb.Bytes())
		}

	case mode == MINIMAL:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_linux_linux_amd64.go"), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_linux_linux_amd64.go"), bb.Bytes())
		}
	}
}

func cgoAndroid(module, mocPath string, mode int, pkg string) {
	var (
		bb              = new(bytes.Buffer)
		libs            = cleanLibs(module, mode)
		QT_DIR          = func() string { return strings.Replace(utils.QT_DIR(), "\\", "/", -1) }
		ANDROID_NDK_DIR = func() string { return strings.Replace(utils.ANDROID_NDK_DIR(), "\\", "/", -1) }
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build android%v\n\n", func() string {
		if mode == MINIMAL {
			return ",minimal"
		}
		if mode == MOC {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	fmt.Fprintf(bb, "#cgo CFLAGS: -fstack-protector-strong -DANDROID -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -fno-builtin-memmove --sysroot=%v/platforms/android-16/arch-arm/ -Os -mthumb -Wall -W -D_REENTRANT -fPIC\n", ANDROID_NDK_DIR())
	fmt.Fprintf(bb, "#cgo CXXFLAGS: -fstack-protector-strong -DANDROID -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -fno-builtin-memmove --sysroot=%v/platforms/android-16/arch-arm/ -O2 -Os -mthumb -std=gnu++11 -Wall -W -D_REENTRANT -fPIC\n", ANDROID_NDK_DIR())

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	for _, s := range []struct{ pre, dir, android_dir string }{{"!", QT_DIR(), ANDROID_NDK_DIR()}, {"", "/opt/Qt5.8.0", "/home/user/android-ndk-r13b"}} {

		fmt.Fprintf(bb, "#cgo +build %vdocker CXXFLAGS: -I%v/%v/android_armv7/include -isystem %v/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem %v/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem %v/platforms/android-16/arch-arm/usr/include -I%v/%v/android_armv7/mkspecs/android-g++\n", s.pre, s.dir, utils.QT_VERSION_MAJOR(), s.android_dir, s.android_dir, s.android_dir, s.dir, utils.QT_VERSION_MAJOR())

		fmt.Fprintf(bb, "#cgo +build %vdocker CXXFLAGS:", s.pre)
		for _, m := range libs {
			fmt.Fprintf(bb, " -I%v/%v/android_armv7/include/Qt%v", s.dir, utils.QT_VERSION_MAJOR(), m)
		}
		fmt.Fprint(bb, "\n\n")

		fmt.Fprintf(bb, "#cgo +build %vdocker LDFLAGS: --sysroot=%v/platforms/android-16/arch-arm/ -Wl,-rpath=%v/%v/android_armv7/lib -Wl,-rpath-link=%v/%v/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition -Wl,--allow-shlib-undefined\n", s.pre, s.android_dir, s.dir, utils.QT_VERSION_MAJOR(), s.dir, utils.QT_VERSION_MAJOR())

		fmt.Fprintf(bb, "#cgo +build %vdocker LDFLAGS: -L%v/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L%v/platforms/android-16/arch-arm//usr/lib -L%v/toolchains/arm-linux-androideabi-4.9/prebuilt/%v-x86_64/bin/../lib/gcc/arm-linux-androideabi/4.9.x -L%v/%v/android_armv7/lib", s.pre, s.android_dir, s.android_dir, s.android_dir, runtime.GOOS, s.dir, utils.QT_VERSION_MAJOR())

		for _, m := range libs {
			if m != "UiPlugin" {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}

		fmt.Fprint(bb, " -lGLESv2 -lgnustl_shared -llog -lz -lm -ldl -lc -lgcc\n")
	}

	fmt.Fprint(bb, "*/\n")
	fmt.Fprint(bb, "import \"C\"\n")

	switch {
	case mode == MOC:
		{
			utils.SaveBytes(filepath.Join(mocPath, "moc_cgo_android_linux_arm.go"), bb.Bytes())
		}

	case mode == MINIMAL:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_android_linux_arm.go"), bb.Bytes())
		}

	default:
		{
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(module), "cgo_android_linux_arm.go"), bb.Bytes())
		}
	}
}

func cgoIos(module, mocPath string, mode int, pkg string) string {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module, mode)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build ${BUILDTARGET}%v\n\n", func() string {
		if mode == MINIMAL {
			return ",minimal"
		}
		if mode == MOC {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	fmt.Fprintf(bb, "#cgo CFLAGS: -pipe -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -O2 -isysroot %v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/%v -mios-simulator-version-min=7.0 -arch x86_64 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC\n", utils.XCODE_DIR(), utils.IPHONESIMULATOR_SDK_DIR())
	fmt.Fprintf(bb, "#cgo CXXFLAGS: -pipe -stdlib=libc++ -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -fvisibility-inlines-hidden -Wno-non-virtual-dtor -Wno-overloaded-virtual -Wno-exit-time-destructors -O2 -std=gnu++11 -isysroot %v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/%v -mios-simulator-version-min=7.0 -arch x86_64 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC\n", utils.XCODE_DIR(), utils.IPHONESIMULATOR_SDK_DIR())

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DDARWIN_NO_CARBON -DQT_COMPILER_SUPPORTS_SSE2 -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/%v/ios/mkspecs/common/uikit -I%v/%v/ios/include -I%v/%v/ios/mkspecs/macx-ios-clang\n", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/%v/ios/include/Qt%v", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,%v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/%v -mios-simulator-version-min=7.0 -arch x86_64\n", utils.XCODE_DIR(), utils.IPHONESIMULATOR_SDK_DIR())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v/%v/ios/plugins/platforms -framework Foundation -framework UIKit -framework QuartzCore -framework AudioToolbox -framework AssetsLibrary -L%v/%v/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework OpenGLES -framework CoreText -framework CoreGraphics -framework Security -framework SystemConfiguration -framework WebKit -framework CoreBluetooth -lqios -lQt5FontDatabaseSupport -lQt5GraphicsSupport -lQt5ClipboardSupport -lqtfreetype -L%v/%v/ios/plugins/imageformats -lqgif -lqicns -lqico -lqjpeg -lqmacjp2 -framework ImageIO -lqtga -lqtiff -lqwbmp -lqwebp -lqtlibpng -lqtharfbuzz -lm -lz -lqtpcre", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
	for _, m := range libs {
		if m != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}

	if !strings.Contains(bb.String(), "-lQt5Gui") {
		fmt.Fprint(bb, " -lQt5Gui")
	}

	switch module {
	case
		"Multimedia":
		{
			fmt.Fprint(bb, " -lQt5OpenGL")
		}

	case "TestLib":
		{
			fmt.Fprintf(bb, " -F%v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest", utils.XCODE_DIR())
		}

	case "Qml", "WebChannel", "Quick", "QuickControls2":
		{
			if module != "Quick" {
				fmt.Fprint(bb, " -lQt5Quick -lQt5QuickParticles -lQt5QuickTest -lQt5QuickWidgets")
			}
			fmt.Fprintf(bb, " -L%v/%v/ios/plugins/qmltooling -lqmldbg_debugger -lqmldbg_inspector -lqmldbg_local -lqmldbg_native -lqmldbg_profiler -lqmldbg_quickprofiler -lqmldbg_server -lQt5PacketProtocol -lqmldbg_tcp", utils.QT_DIR(), utils.QT_VERSION_MAJOR())

			fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtQuick.2 -lqtquick2plugin -L%v/%v/ios/qml/QtQuick/Layouts -lqquicklayoutsplugin -L%v/%v/ios/qml/QtQuick/Dialogs -ldialogplugin -L%v/%v/ios/qml/QtQuick/Controls -lqtquickcontrolsplugin -L%v/%v/ios/qml/Qt/labs/folderlistmodel -lqmlfolderlistmodelplugin -L%v/%v/ios/qml/Qt/labs/settings -lqmlsettingsplugin -L%v/%v/ios/qml/QtQuick/Dialogs/Private -ldialogsprivateplugin -L%v/%v/ios/qml/QtQuick/Window.2 -lwindowplugin -L%v/%v/ios/qml/QtQml/Models.2 -lmodelsplugin -L%v/%v/ios/qml/QtQuick/Extras -lqtquickextrasplugin -L%v/%v/ios/qml/QtGraphicalEffects/private -lqtgraphicaleffectsprivate", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
			fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtQuick/Controls.2 -lqtquickcontrols2plugin -L%v/%v/ios/qml/QtQuick/Controls.2/Material -lqtquickcontrols2materialstyleplugin -L%v/%v/ios/qml/QtQuick/Controls.2/Universal -lqtquickcontrols2universalstyleplugin -lQt5QuickControls2", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
			fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtQuick/Templates.2 -lqtquicktemplates2plugin -lQt5QuickTemplates2 -L%v/%v/ios/qml/QtGraphicalEffects -lqtgraphicaleffectsplugin", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		}

	case "Purchasing":
		{
			fmt.Fprint(bb, " -framework StoreKit")
		}

	case "WebView":
		{
			fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtWebView -ldeclarative_webview", utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		}
	}

	if module == "build_ios" {
		fmt.Fprint(bb, " -lQt5OpenGL")
		fmt.Fprintf(bb, " -F%v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest", utils.XCODE_DIR())
		fmt.Fprint(bb, " -lQt5Quick -lQt5QuickParticles -lQt5QuickTest -lQt5QuickWidgets")
		fmt.Fprintf(bb, " -L%v/%v/ios/plugins/qmltooling -lqmldbg_debugger -lqmldbg_inspector -lqmldbg_local -lqmldbg_native -lqmldbg_profiler -lqmldbg_quickprofiler -lqmldbg_server -lQt5PacketProtocol -lqmldbg_tcp", utils.QT_DIR(), utils.QT_VERSION_MAJOR())

		fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtQuick.2 -lqtquick2plugin -L%v/%v/ios/qml/QtQuick/Layouts -lqquicklayoutsplugin -L%v/%v/ios/qml/QtQuick/Dialogs -ldialogplugin -L%v/%v/ios/qml/QtQuick/Controls -lqtquickcontrolsplugin -L%v/%v/ios/qml/Qt/labs/folderlistmodel -lqmlfolderlistmodelplugin -L%v/%v/ios/qml/Qt/labs/settings -lqmlsettingsplugin -L%v/%v/ios/qml/QtQuick/Dialogs/Private -ldialogsprivateplugin -L%v/%v/ios/qml/QtQuick/Window.2 -lwindowplugin -L%v/%v/ios/qml/QtQml/Models.2 -lmodelsplugin -L%v/%v/ios/qml/QtQuick/Extras -lqtquickextrasplugin -L%v/%v/ios/qml/QtGraphicalEffects/private -lqtgraphicaleffectsprivate", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtQuick/Controls.2 -lqtquickcontrols2plugin -L%v/%v/ios/qml/QtQuick/Controls.2/Material -lqtquickcontrols2materialstyleplugin -L%v/%v/ios/qml/QtQuick/Controls.2/Universal -lqtquickcontrols2universalstyleplugin -lQt5QuickControls2", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtQuick/Templates.2 -lqtquicktemplates2plugin -lQt5QuickTemplates2 -L%v/%v/ios/qml/QtGraphicalEffects -lqtgraphicaleffectsplugin", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())

		fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtMultimedia -ldeclarative_multimedia -lQt5MultimediaQuick_p", utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		fmt.Fprintf(bb, " -L%v/%v/ios/plugins/mediaservice -lqtmedia_audioengine -L%v/%v/ios/plugins/audio -lqtaudio_coreaudio -L%v/%v/ios/plugins/playlistformats -lqtmultimedia_m3u -lQt5Multimedia", utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR(), utils.QT_DIR(), utils.QT_VERSION_MAJOR())
		fmt.Fprint(bb, " -lqavfcamera -framework AudioToolbox -framework CoreAudio -framework AVFoundation -framework CoreMedia -framework CoreVideo -lqavfmediaplayer -lQt5MultimediaWidgets")

		fmt.Fprintf(bb, " -L%v/%v/ios/qml/QtWebView -ldeclarative_webview", utils.QT_DIR(), utils.QT_VERSION_MAJOR())
	}

	fmt.Fprint(bb, "\n*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	if module == "build_ios" {
		return bb.String()
	}

	var path, prefix = func() (string, string) {
		switch {
		case mode == MOC:
			{
				return mocPath, "moc_"
			}

		case mode == MINIMAL:
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

func cgoSailfish(module, mocPath string, mode int, pkg string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module, mode)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build ${BUILDTARGET}%v\n\n", func() string {
		if mode == MINIMAL {
			return ",minimal"
		}
		return ""
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
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
		if !(m == "UiPlugin" || m == "Sailfish") {
			if parser.IsWhiteListedSailfishLib(m) {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = strings.Replace(bb.String(), "${BUILDTARGET}", "sailfish_emulator", -1)

	switch {
	case mode == RCC:
		{
			utils.Save(filepath.Join(mocPath, "rcc_cgo_sailfish_emulator_linux_386.go"), tmp)
		}

	case mode == MOC:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_sailfish_emulator_linux_386.go"), tmp)
		}

	case mode == MINIMAL:
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
	case mode == RCC:
		{
			utils.Save(filepath.Join(mocPath, "rcc_cgo_sailfish_linux_arm.go"), tmp)
		}

	case mode == MOC:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_sailfish_linux_arm.go"), tmp)
		}

	case mode == MINIMAL:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_sailfish_linux_arm.go"), tmp)
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_sailfish_linux_arm.go"), tmp)
		}
	}
}

func cgoRaspberryPi1(module, mocPath string, mode int, pkg string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module, mode)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build rpi1%v\n\n", func() string {
		if mode == MINIMAL {
			return ",minimal"
		}
		if mode == MOC {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
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
	case mode == MOC:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_rpi1_linux_arm.go"), tmp)
		}

	case mode == MINIMAL:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_rpi1_linux_arm.go"), tmp)
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_rpi1_linux_arm.go"), tmp)
		}
	}
}

func cgoRaspberryPi2(module, mocPath string, mode int, pkg string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module, mode)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build rpi2%v\n\n", func() string {
		if mode == MINIMAL {
			return ",minimal"
		}
		if mode == MOC {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
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
	case mode == MOC:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_rpi2_linux_arm.go"), tmp)
		}

	case mode == MINIMAL:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_rpi2_linux_arm.go"), tmp)
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_rpi2_linux_arm.go"), tmp)
		}
	}
}

func cgoRaspberryPi3(module, mocPath string, mode int, pkg string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module, mode)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build rpi3%v\n\n", func() string {
		if mode == MINIMAL {
			return ",minimal"
		}
		if mode == MOC {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
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
	case mode == MOC:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_rpi3_linux_arm.go"), tmp)
		}

	case mode == MINIMAL:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_rpi3_linux_arm.go"), tmp)
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_rpi3_linux_arm.go"), tmp)
		}
	}
}

func cleanLibs(module string, mode int) []string {

	if mode == MOC || module == "build_ios" {
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

//needed for static linking
func GetiOSClang(buildTarget, buildARM string) []string {
	var tmp = cgoIos("build_ios", "", NONE, "main")

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

func cgoAsteroid(module, mocPath string, mode int, pkg string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module, mode)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "// +build ${BUILDTARGET}%v\n\n", func() string {
		if mode == MINIMAL {
			return ",minimal"
		}
		if mode == MOC {
			return ""
		}
		return ",!minimal"
	}())

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if mode == MOC {
			return pkg
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -feliminate-unused-debug-types -fexceptions -fstack-protector --param=ssp-buffer-size=4 -Wformat -Wformat-security -fmessage-length=0 -march=armv7ve -mfloat-abi=softfp -mfpu=neon -mthumb -Wno-psabi -fPIC -fvisibility=hidden -Wall -W -D_REENTRANT -fPIE\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -feliminate-unused-debug-types -fexceptions -fstack-protector --param=ssp-buffer-size=4 -Wformat -Wformat-security -fmessage-length=0 -march=armv7ve -mfloat-abi=softfp -mfpu=neon -mthumb -Wno-psabi -fPIC -fvisibility=hidden -Wall -W -D_REENTRANT -fPIE\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%[1]s/usr/include/c++/6.2.0/arm-oe-linux-gnueabi -I%[1]s/usr/include/c++/6.2.0  -I%[1]s/usr/lib/mkspecs -I%[1]s/usr/include -I%[1]s/usr/include/mdeclarativecache5 -I%[1]s/usr/include/resource/qt5\n", os.Getenv("OECORE_TARGET_SYSROOT"))

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%s/usr/include/Qt%v", os.Getenv("OECORE_TARGET_SYSROOT"), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath-link,%[1]s/usr/lib -Wl,-rpath-link,%[1]s/lib\n", os.Getenv("OECORE_TARGET_SYSROOT"))

	fmt.Fprintf(bb, "#cgo LDFLAGS: -rdynamic -L%[1]s/usr/lib -L%[1]s/lib -lmdeclarativecache5", os.Getenv("OECORE_TARGET_SYSROOT"))
	for _, m := range libs {
		if m != "UiPlugin" {
			if parser.IsWhiteListedSailfishLib(m) {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var tmp = strings.Replace(bb.String(), "${BUILDTARGET}", "asteroid", -1)
	tmp = strings.Replace(tmp, "i486", "armv7ve", -1)

	switch {
	case mode == MOC:
		{
			utils.Save(filepath.Join(mocPath, "moc_cgo_asteroid_linux_arm.go"), tmp)
		}

	case mode == MINIMAL:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "minimal_cgo_asteroid_linux_arm.go"), tmp)
		}

	default:
		{
			utils.Save(utils.GoQtPkgPath(strings.ToLower(module), "cgo_asteroid_linux_arm.go"), tmp)
		}
	}
}
