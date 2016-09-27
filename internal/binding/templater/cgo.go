package templater

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

var (
	MocAppPath string
	MocModule  string
)

func CopyCgo(module string) {

	if Minimal {
		return
	}

	if !(strings.Contains(module, "droid") || strings.Contains(module, "fish")) {
		cgoDarwin(module)
		cgoWindows(module)
		cgoLinux(module)
		cgoSailfish(module)
		cgoRaspberryPi1(module)
		cgoRaspberryPi2(module)
		cgoRaspberryPi3(module)

		if runtime.GOOS == "darwin" {
			cgoIos(module)
		}
	}

	if strings.Contains(module, "fish") {
		cgoSailfish(module)
	}

	if !strings.Contains(module, "fish") {
		switch runtime.GOOS {
		case "darwin", "linux":
			{
				cgoAndroidOnDarwinAndLinux(module)
			}

		case "windows":
			{
				cgoAndroidOnWindows(module)
			}
		}
	}
}

func cgoDarwin(module string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk -mmacosx-version-min=10.8 -Wall -W -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -stdlib=libc++ -O2 -std=gnu++11 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk -mmacosx-version-min=10.8 -Wall -W -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_DEBUG")
	for i := len(libs) - 1; i >= 0; i-- {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(libs[i]))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/5.7/clang_64/mkspecs/macx-clang -F%v/5.7/clang_64/lib\n", utils.QtInstallDir(), utils.QtInstallDir())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for i := len(libs) - 1; i >= 0; i-- {
		if libs[i] == "UiTools" || libs[i] == "PlatformHeaders" {
			fmt.Fprintf(bb, " -I%v/5.7/clang_64/include/Qt%v", utils.QtInstallDir(), libs[i])
		} else {
			fmt.Fprintf(bb, " -I%v/5.7/clang_64/lib/Qt%v.framework/Headers", utils.QtInstallDir(), libs[i])
		}
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/System/Library/Frameworks/AGL.framework/Headers\n")
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk -mmacosx-version-min=10.8 -Wl,-rpath,%v/5.7/clang_64/lib\n", utils.QtInstallDir())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -F%v/5.7/clang_64/lib", utils.QtInstallDir())

	for i := len(libs) - 1; i >= 0; i-- {
		if libs[i] == "UiTools" || libs[i] == "PlatformHeaders" {
			fmt.Fprintf(bb, " -L%v/5.7/clang_64/lib -lQt5%v", utils.QtInstallDir(), libs[i])
		} else {
			if libs[i] != "UiPlugin" {
				fmt.Fprintf(bb, " -framework Qt%v", libs[i])
			}
		}
	}

	fmt.Fprint(bb, " -framework DiskArbitration -framework IOKit -framework OpenGL -framework AGL\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_darwin_amd64.go"), bb.String())
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_darwin_amd64.go"), bb.String())
	}
}

func cgoWindows(module string) {
	var (
		bb           = new(bytes.Buffer)
		libs         = cleanLibs(module)
		QtInstallDir = func() string { return strings.Replace(utils.QtInstallDirWindows(), "\\", "/", -1) }
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
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
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/5.7/mingw53_32/include -I%v/5.7/mingw53_32/mkspecs/win32-g++\n", QtInstallDir(), QtInstallDir())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/5.7/mingw53_32/include/Qt%v", QtInstallDir(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprint(bb, "#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v/5.7/mingw53_32/lib", QtInstallDir())
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
	}

	fmt.Fprint(bb, " -lmingw32 -lqtmain -lshell32\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_windows_386.go"), bb.String())
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_windows_386.go"), bb.String())
	}
}

func cgoLinux(module string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
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

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/5.7/gcc_64/include -I%v/5.7/gcc_64/mkspecs/linux-g++\n", utils.QtInstallDir(), utils.QtInstallDir())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/5.7/gcc_64/include/Qt%v", utils.QtInstallDir(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,%v/5.7/gcc_64 -Wl,-rpath,%v/5.7/gcc_64/lib\n", utils.QtInstallDir(), utils.QtInstallDir())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v/5.7/gcc_64/lib -L/usr/lib64", utils.QtInstallDir())
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
	}

	fmt.Fprint(bb, " -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_amd64.go"), bb.String())
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_amd64.go"), bb.String())
	}
}

func cgoAndroidOnDarwinAndLinux(module string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	fmt.Fprint(bb, "#cgo CFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -std=c++11 -O2 -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_DEBUG")
	for i := len(libs) - 1; i >= 0; i-- {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(libs[i]))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/5.7/android_armv7/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem /opt/android-ndk/platforms/android-9/arch-arm/usr/include -I%v/5.7/android_armv7/mkspecs/android-g++\n", utils.QtInstallDir(), utils.QtInstallDir())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for i := len(libs) - 1; i >= 0; i-- {
		fmt.Fprintf(bb, " -I%v/5.7/android_armv7/include/Qt%v", utils.QtInstallDir(), libs[i])
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: --sysroot=/opt/android-ndk/platforms/android-9/arch-arm -Wl,-rpath=%v/5.7/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition -Wl,--allow-shlib-undefined\n", utils.QtInstallDir())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L/opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android-ndk/platforms/android-9/arch-arm//usr/lib -L%v/5.7/android_armv7/lib -L/opt/android/ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android/ndk/platforms/android-9/arch-arm//usr/lib", utils.QtInstallDir())

	for i := len(libs) - 1; i >= 0; i-- {
		if libs[i] != "UiPlugin" {
			fmt.Fprintf(bb, " -lQt5%v", libs[i])
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lgnustl_shared -llog -lz -lm -ldl -lc -lgcc\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_android_arm.go"), bb.String())
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_android_arm.go"), bb.String())
	}
}

func cgoAndroidOnWindows(module string) {
	var (
		bb           = new(bytes.Buffer)
		libs         = cleanLibs(module)
		QtInstallDir = func() string { return strings.Replace(utils.QtInstallDirWindows(), "\\", "/", -1) }
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
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

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/5.7/android_armv7/include -isystem C:/android/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem C:/android/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem C:/android/android-ndk/platforms/android-9/arch-arm/usr/include -I%v/5.7/android_armv7/mkspecs/android-g++\n", QtInstallDir(), QtInstallDir())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/5.7/android_armv7/include/Qt%v", QtInstallDir(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: --sysroot=C:/android/android-ndk/platforms/android-9/arch-arm -Wl,-rpath=%v/5.7/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition -Wl,--allow-shlib-undefined\n", QtInstallDir())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -LC:/android/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -LC:/android/android-ndk/platforms/android-9/arch-arm//usr/lib -L%v/5.7/android_armv7/lib -LC:/android/android/ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -LC:/android/android/ndk/platforms/android-9/arch-arm//usr/lib", QtInstallDir())
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lgnustl_shared -llog -lz -lm -ldl -lc -lgcc\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_android_arm.go"), bb.String())
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_android_arm.go"), bb.String())
	}
}

func cgoIos(module string) string {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator.sdk -mios-simulator-version-min=7.0 -arch i386 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -stdlib=libc++ -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -fvisibility-inlines-hidden -Wno-non-virtual-dtor -Wno-overloaded-virtual -Wno-exit-time-destructors -O2 -std=gnu++11 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator.sdk -mios-simulator-version-min=7.0 -arch i386 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DDARWIN_NO_CARBON -DQT_NO_PRINTER -DQT_NO_PRINTDIALOG -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/5.7/ios/mkspecs/macx-ios-clang/ios -I%v/5.7/ios/include -I%v/5.7/ios/mkspecs/macx-ios-clang\n", utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/5.7/ios/include/Qt%v", utils.QtInstallDir(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprint(bb, "#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator.sdk -mios-simulator-version-min=7.0 -arch i386\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L%v/5.7/ios/plugins/platforms -lqios_iphonesimulator -framework Foundation -framework UIKit -framework QuartzCore -framework AssetsLibrary -L%v/5.7/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework CoreText -framework CoreGraphics -framework OpenGLES -lqtfreetype_iphonesimulator -framework Security -framework SystemConfiguration -framework CoreBluetooth -L%v/5.7/ios/plugins/imageformats -lqdds_iphonesimulator -lqicns_iphonesimulator -lqico_iphonesimulator -lqtga_iphonesimulator -lqtiff_iphonesimulator -lqwbmp_iphonesimulator -lqwebp_iphonesimulator -lqtharfbuzzng_iphonesimulator -lz -lqtpcre_iphonesimulator -lm", utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir())
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			fmt.Fprintf(bb, " -lQt5%v_iphonesimulator", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				fmt.Fprintf(bb, " -lQt5%v_iphonesimulator", m)
			}
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
			fmt.Fprint(bb, " -F/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest")
		}

	case "Qml", "WebChannel", "Quick":
		{
			if module != "Quick" {
				fmt.Fprint(bb, " -lQt5Quick_iphonesimulator -lQt5QuickParticles_iphonesimulator -lQt5QuickTest_iphonesimulator -lQt5QuickWidgets_iphonesimulator")
			}
			fmt.Fprintf(bb, " -L%v/5.7/ios/plugins/qmltooling -lqmldbg_debugger_iphonesimulator -lqmldbg_inspector_iphonesimulator -lqmldbg_local_iphonesimulator -lqmldbg_native_iphonesimulator -lqmldbg_profiler_iphonesimulator -lqmldbg_quickprofiler_iphonesimulator -lqmldbg_server_iphonesimulator -lQt5PacketProtocol_iphonesimulator -lqmldbg_tcp_iphonesimulator", utils.QtInstallDir())
		}

	case "Purchasing":
		{
			fmt.Fprint(bb, " -framework StoreKit")
		}
	}

	if module == "build_ios" {
		fmt.Fprint(bb, " -lQt5OpenGL_iphonesimulator")
		fmt.Fprint(bb, " -F/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest")
		fmt.Fprint(bb, " -lQt5Quick_iphonesimulator -lQt5QuickParticles_iphonesimulator -lQt5QuickTest_iphonesimulator -lQt5QuickWidgets_iphonesimulator")
		fmt.Fprintf(bb, " -L%v/5.7/ios/plugins/qmltooling -lqmldbg_debugger_iphonesimulator -lqmldbg_inspector_iphonesimulator -lqmldbg_local_iphonesimulator -lqmldbg_native_iphonesimulator -lqmldbg_profiler_iphonesimulator -lqmldbg_quickprofiler_iphonesimulator -lqmldbg_server_iphonesimulator -lQt5PacketProtocol_iphonesimulator -lqmldbg_tcp_iphonesimulator", utils.QtInstallDir())

		fmt.Fprintf(bb, " -L%v/5.7/ios/qml/QtQuick.2 -lqtquick2plugin_iphonesimulator -L%v/5.7/ios/qml/QtQuick/Layouts -lqquicklayoutsplugin_iphonesimulator -L%v/5.7/ios/qml/QtQuick/Dialogs -ldialogplugin_iphonesimulator -L%v/5.7/ios/qml/QtQuick/Controls -lqtquickcontrolsplugin_iphonesimulator -L%v/5.7/ios/qml/Qt/labs/folderlistmodel -lqmlfolderlistmodelplugin_iphonesimulator -L%v/5.7/ios/qml/Qt/labs/settings -lqmlsettingsplugin_iphonesimulator -L%v/5.7/ios/qml/QtQuick/Dialogs/Private -ldialogsprivateplugin_iphonesimulator -L%v/5.7/ios/qml/QtQuick/Window.2 -lwindowplugin_iphonesimulator -L%v/5.7/ios/qml/QtQml/Models.2 -lmodelsplugin_iphonesimulator -L%v/5.7/ios/qml/QtQuick/Extras -lqtquickextrasplugin_iphonesimulator -L%v/5.7/ios/qml/QtGraphicalEffects/private -lqtgraphicaleffectsprivate_iphonesimulator", utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir())
		fmt.Fprintf(bb, " -L%v/5.7/ios/qml/QtMultimedia -ldeclarative_multimedia_iphonesimulator -lQt5MultimediaQuick_p_iphonesimulator", utils.QtInstallDir())
		fmt.Fprintf(bb, " -L%v/5.7/ios/plugins/mediaservice -lqtmedia_audioengine_iphonesimulator -L%v/5.7/ios/plugins/audio -lqtaudio_coreaudio_iphonesimulator -L%v/5.7/ios/plugins/playlistformats -lqtmultimedia_m3u_iphonesimulator -lQt5Multimedia_iphonesimulator", utils.QtInstallDir(), utils.QtInstallDir(), utils.QtInstallDir())
		fmt.Fprint(bb, " -lqavfcamera_iphonesimulator -framework AudioToolbox -framework CoreAudio -framework AVFoundation -framework CoreMedia -framework CoreVideo -lqavfmediaplayer_iphonesimulator -lQt5MultimediaWidgets_iphonesimulator")
	}

	fmt.Fprint(bb, " -lQt5PlatformSupport_iphonesimulator\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	if module == "build_ios" {
		return bb.String()
	}

	var path, prefix = func() (string, string) {
		if module == parser.MOC {
			return MocAppPath, "moc_"
		}
		return utils.GetQtPkgPath(strings.ToLower(module)), ""
	}()

	//386
	utils.Save(filepath.Join(path, prefix+"cgo_darwin_386.go"), bb.String())

	//arm64
	var tmp = strings.Replace(bb.String(), "_iphonesimulator", "", -1)
	tmp = strings.Replace(tmp, "i386", "arm64", -1)
	tmp = strings.Replace(tmp, "iPhoneSimulator", "iPhoneOS", -1)
	tmp = strings.Replace(tmp, "ios-simulator", "iphoneos", -1)
	utils.Save(filepath.Join(path, prefix+"cgo_darwin_arm64.go"), tmp)

	//armv7
	utils.Save(filepath.Join(path, prefix+"cgo_darwin_arm.go"), strings.Replace(tmp, "arm64", "armv7", -1))

	return ""
}

func cgoSailfish(module string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()

	fmt.Fprint(bb, "// +build !android,!rpi1,!rpi2,!rpi3\n\n")

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
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
		if m == "UiTools" || m == "PlatformHeaders" {
			if IsWhiteListedSailfishLib(m) {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				if IsWhiteListedSailfishLib(m) {
					fmt.Fprintf(bb, " -lQt5%v", m)
				}
			}
		}
	}
	fmt.Fprint(bb, " -lGLESv2 -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_386.go"), bb.String())
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_386.go"), bb.String())
	}

	var tmp = strings.Replace(bb.String(), "-m32 -msse -msse2 -march=i686 -mfpmath=sse -mtune=generic -fno-omit-frame-pointer -fasynchronous-unwind-tables", "-fmessage-length=0 -march=armv7-a -mfloat-abi=hard -mfpu=neon -mthumb -Wno-psabi", -1)
	tmp = strings.Replace(tmp, "i486", "armv7hl", -1)

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_arm.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_arm.go"), tmp)
	}
}

func cgoRaspberryPi1(module string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()
	fmt.Fprint(bb, "// +build rpi1\n\n")

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -marm -mfpu=vfp -mtune=arm1176jzf-s -march=armv6zk -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -marm -mfpu=vfp -mtune=arm1176jzf-s -march=armv6zk -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/5.7/rpi1/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I%v/5.7/rpi1/mkspecs/devices/linux-rasp-pi-g++\n", utils.QtInstallDir(), utils.QtInstallDir())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/5.7/rpi1/include/Qt%v", utils.QtInstallDir(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,%v/5.7/rpi1/lib -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin\n", utils.QtInstallDir())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L/home/${USERNAME}/raspi/sysroot/opt/vc/lib -L%v/5.7/rpi1/lib", utils.QtInstallDir())
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var username = os.Getenv("USERNAME")
	if username == "" {
		username = "user"
	}

	var tmp = strings.Replace(bb.String(), "${USERNAME}", username, -1)

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_rpi1.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_rpi1.go"), tmp)
	}
}

func cgoRaspberryPi2(module string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()
	fmt.Fprint(bb, "// +build rpi2\n\n")

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	fmt.Fprint(bb, "#cgo CFLAGS: -pipe -march=armv7-a -marm -mthumb-interwork -mfpu=neon-vfpv4 -mtune=cortex-a7 -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -pipe -march=armv7-a -marm -mthumb-interwork -mfpu=neon-vfpv4 -mtune=cortex-a7 -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/5.7/rpi2/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I%v/5.7/rpi2/mkspecs/devices/linux-rasp-pi2-g++\n", utils.QtInstallDir(), utils.QtInstallDir())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/5.7/rpi2/include/Qt%v", utils.QtInstallDir(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,%v/5.7/rpi2/lib -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin\n", utils.QtInstallDir())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L/home/${USERNAME}/raspi/sysroot/opt/vc/lib -L%v/5.7/rpi2/lib", utils.QtInstallDir())
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var username = os.Getenv("USERNAME")
	if username == "" {
		username = "user"
	}

	var tmp = strings.Replace(bb.String(), "${USERNAME}", username, -1)

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_rpi2.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_rpi2.go"), tmp)
	}
}

func cgoRaspberryPi3(module string) {
	var (
		bb   = new(bytes.Buffer)
		libs = cleanLibs(module)
	)
	defer bb.Reset()
	fmt.Fprint(bb, "// +build rpi3\n\n")

	fmt.Fprintf(bb, "package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	fmt.Fprint(bb, "/*\n")

	fmt.Fprint(bb, "#cgo CFLAGS: -march=armv8-a+crc -mtune=cortex-a53 -mfpu=crypto-neon-fp-armv8 -pipe -Os -mthumb -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")
	fmt.Fprint(bb, "#cgo CXXFLAGS: -march=armv8-a+crc -mtune=cortex-a53 -mfpu=crypto-neon-fp-armv8 -pipe -Os -mthumb -std=c++11 -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n")

	fmt.Fprint(bb, "#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG")
	for _, m := range libs {
		fmt.Fprintf(bb, " -DQT_%v_LIB", strings.ToUpper(m))
	}
	fmt.Fprint(bb, "\n")

	fmt.Fprintf(bb, "#cgo CXXFLAGS: -I%v/5.7/rpi3/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I%v/5.7/rpi3/mkspecs/devices/linux-rpi3-g++\n", utils.QtInstallDir(), utils.QtInstallDir())

	fmt.Fprint(bb, "#cgo CXXFLAGS:")
	for _, m := range libs {
		fmt.Fprintf(bb, " -I%v/5.7/rpi3/include/Qt%v", utils.QtInstallDir(), m)
	}
	fmt.Fprint(bb, "\n\n")

	fmt.Fprintf(bb, "#cgo LDFLAGS: -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,%v/5.7/rpi3/lib -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin\n", utils.QtInstallDir())

	fmt.Fprintf(bb, "#cgo LDFLAGS: -L/home/${USERNAME}/raspi/sysroot/opt/vc/lib -L%v/5.7/rpi3/lib", utils.QtInstallDir())
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			fmt.Fprintf(bb, " -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				fmt.Fprintf(bb, " -lQt5%v", m)
			}
		}
	}

	fmt.Fprint(bb, " -lGLESv2 -lpthread\n")
	fmt.Fprint(bb, "*/\n")

	fmt.Fprint(bb, "import \"C\"\n")

	var username = os.Getenv("USERNAME")
	if username == "" {
		username = "user"
	}

	var tmp = strings.Replace(bb.String(), "${USERNAME}", username, -1)

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_rpi3.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_rpi3.go"), tmp)
	}
}

func cleanLibs(module string) []string {

	if module == parser.MOC || module == "build_ios" {
		return LibDeps[module]
	}

	if module == "Designer" {
		return append([]string{module}, LibDeps[module]...)
	}

	var libs = append(LibDeps[module], module)
	for i, k := range libs {
		if k == "TestLib" {
			libs[i] = "Test"
		}
	}
	return libs
}

func GetiOSClang(buildTarget, buildARM string) []string {
	var tmp = cgoIos("build_ios")

	tmp = strings.Split(tmp, "/*")[1]
	tmp = strings.Split(tmp, "*/")[0]

	tmp = strings.Replace(tmp, "#cgo CFLAGS: ", "", -1)
	tmp = strings.Replace(tmp, "#cgo CXXFLAGS: ", "", -1)
	tmp = strings.Replace(tmp, "#cgo LDFLAGS: ", "", -1)
	tmp = strings.Replace(tmp, "\n", " ", -1)

	if buildTarget == "ios" {
		tmp = strings.Replace(tmp, "_iphonesimulator", "", -1)
		tmp = strings.Replace(tmp, "i386", "arm64", -1)
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
