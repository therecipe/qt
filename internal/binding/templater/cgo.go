package templater

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func CopyCgo(module string) {
	createCgoDarwin(module)
	createCgoWindows(module)
	createCgoLinux(module)

	switch runtime.GOOS {
	case "darwin", "linux":
		createCgoAndroidDarwin(module)
	}
}

func createCgoDarwin(module string) {
	var tmp string

	tmp += fmt.Sprintf("package %v\n", strings.ToLower(module))
	tmp += "/*\n"

	tmp += "#cgo CPPFLAGS: -pipe -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.7 -Wall -W\n"

	tmp += "#cgo CPPFLAGS: -DQT_NO_DEBUG"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/clang_64/mkspecs/macx-clang -F/usr/local/Qt5.5.1/5.5/clang_64/lib\n"

	tmp += "#cgo CPPFLAGS:"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.5.1/5.5/clang_64/lib/Qt%v.framework/Headers", m)
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/AGL.framework/Headers\n"
	tmp += "\n"

	tmp += "#cgo LDFLAGS: -headerpad_max_install_names -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.7\n"

	tmp += "#cgo LDFLAGS: -F/usr/local/Qt5.5.1/5.5/clang_64/lib"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -framework Qt%v", m)
	}
	tmp += " -framework DiskArbitration -framework IOKit -framework OpenGL -framework AGL\n"

	tmp += "*/\n"
	tmp += fmt.Sprintf("import \"C\"\n")

	utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_darwin_amd64.go"), tmp)
}

func createCgoWindows(module string) {
	var tmp string

	tmp += fmt.Sprintf("package %v\n", strings.ToLower(module))
	tmp += "/*\n"

	tmp += "#cgo CPPFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra -fexceptions -mthreads\n"

	tmp += "#cgo CPPFLAGS: -DUNICODE -DQT_NO_DEBUG"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include -IC:/Qt/Qt5.5.1/5.5/mingw492_32/mkspecs/win32-g++\n"

	tmp += "#cgo CPPFLAGS:"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads\n"

	tmp += "#cgo LDFLAGS: -LC:/Qt/Qt5.5.1/5.5/mingw492_32/bin"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -lQt5%v", m)
	}

	tmp += "\n*/\n"
	tmp += fmt.Sprintf("import \"C\"\n")

	utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_windows_386.go"), tmp)
}

func createCgoLinux(module string) {
	var tmp string

	tmp += fmt.Sprintf("package %v\n", strings.ToLower(module))
	tmp += "/*\n"

	tmp += "#cgo CPPFLAGS: -pipe -O2 -Wall -W -D_REENTRANT\n"

	tmp += "#cgo CPPFLAGS: -DQT_NO_DEBUG"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/gcc/include -I/usr/local/Qt5.5.1/5.5/gcc/mkspecs/linux-g++\n"

	tmp += "#cgo CPPFLAGS:"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.5.1/5.5/gcc/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.5.1/5.5/gcc -Wl,-rpath,/usr/local/Qt5.5.1/5.5/gcc/lib\n"

	tmp += "#cgo LDFLAGS: -L/usr/local/Qt5.5.1/5.5/gcc/lib -L/usr/lib64"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -lQt5%v", m)
	}

	tmp += " -lpthread\n"

	tmp += "*/\n"
	tmp += fmt.Sprintf("import \"C\"\n")

	utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_386.go"), strings.Replace(tmp, "lib64", "lib", -1))
	utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_amd64.go"), strings.Replace(tmp, "gcc", "gcc_64", -1))
}

func createCgoAndroidDarwin(module string) {
	var tmp string

	tmp += fmt.Sprintf("package %v\n", strings.ToLower(module))
	tmp += "/*\n"

	tmp += "#cgo CPPFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT\n"

	tmp += "#cgo CPPFLAGS: -DQT_NO_DEBUG"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/android_armv7/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem /opt/android-ndk/platforms/android-9/arch-arm/usr/include -I/usr/local/Qt5.5.1/5.5/android_armv7/mkspecs/android-g++\n"

	tmp += "#cgo CPPFLAGS:"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.5.1/5.5/android_armv7/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: --sysroot=/opt/android-ndk/platforms/android-9/arch-arm/ -Wl,-rpath=/usr/local/Qt5.5.1/5.5/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition\n"

	tmp += "#cgo LDFLAGS: -L/opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android-ndk/platforms/android-9/arch-arm//usr/lib -L/usr/local/Qt5.5.1/5.5/android_armv7/lib -L/opt/android/ndk/sources/cxx-stl/gnu-libstdc++/4.8/libs/armeabi-v7a -L/opt/android/ndk/platforms/android-9/arch-arm//usr/lib"
	for _, m := range append(LibDeps[module], module) {
		tmp += fmt.Sprintf(" -lQt5%v", m)
	}

	tmp += " -lGLESv2 -lgnustl_shared -llog -lz -lm -ldl -lc -lgcc\n"

	tmp += "*/\n"
	tmp += fmt.Sprintf("import \"C\"\n")

	utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_android_arm.go"), tmp)
}
