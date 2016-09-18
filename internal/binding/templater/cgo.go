package templater

import (
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
		tmp  string
		libs = cleanLibs(module)
	)

	tmp += fmt.Sprintf("package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	tmp += "/*\n"

	tmp += "#cgo CFLAGS: -pipe -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.8 -Wall -W -fPIC\n"
	tmp += "#cgo CXXFLAGS: -pipe -stdlib=libc++ -O2 -std=gnu++11 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.8 -Wall -W -fPIC\n"

	tmp += "#cgo CXXFLAGS: -DQT_NO_DEBUG"
	for i := len(libs) - 1; i >= 0; i-- {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(libs[i]))
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/clang_64/mkspecs/macx-clang -F/usr/local/Qt5.7.0/5.7/clang_64/lib\n"

	tmp += "#cgo CXXFLAGS:"
	for i := len(libs) - 1; i >= 0; i-- {
		if libs[i] == "UiTools" || libs[i] == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -I/usr/local/Qt5.7.0/5.7/clang_64/include/Qt%v", libs[i])
		} else {
			tmp += fmt.Sprintf(" -I/usr/local/Qt5.7.0/5.7/clang_64/lib/Qt%v.framework/Headers", libs[i])
		}
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/AGL.framework/Headers\n"
	tmp += "\n"

	tmp += "#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.8 -Wl,-rpath,/usr/local/Qt5.7.0/5.7/clang_64/lib\n"

	tmp += "#cgo LDFLAGS: -F/usr/local/Qt5.7.0/5.7/clang_64/lib"

	for i := len(libs) - 1; i >= 0; i-- {
		if libs[i] == "UiTools" || libs[i] == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -L/usr/local/Qt5.7.0/5.7/clang_64/lib -lQt5%v", libs[i])
		} else {
			if libs[i] != "UiPlugin" {
				tmp += fmt.Sprintf(" -framework Qt%v", libs[i])
			}
		}
	}

	tmp += " -framework DiskArbitration -framework IOKit -framework OpenGL -framework AGL\n"
	tmp += "*/\n"

	tmp += fmt.Sprintf("import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_darwin_amd64.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_darwin_amd64.go"), tmp)
	}
}

func cgoWindows(module string) {
	var (
		tmp  string
		libs = cleanLibs(module)
	)

	tmp += fmt.Sprintf("package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	tmp += "/*\n"

	tmp += "#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra\n"
	tmp += "#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -O2 -std=gnu++11 -frtti -Wall -Wextra -fexceptions -mthreads\n"

	tmp += "#cgo CXXFLAGS: -DUNICODE -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -IC:/Qt/Qt5.7.0/5.7/mingw53_32/include -IC:/Qt/Qt5.7.0/5.7/mingw53_32/mkspecs/win32-g++\n"

	tmp += "#cgo CXXFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -IC:/Qt/Qt5.7.0/5.7/mingw53_32/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition\n"

	tmp += "#cgo LDFLAGS: -LC:/Qt/Qt5.7.0/5.7/mingw53_32/lib"
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				tmp += fmt.Sprintf(" -lQt5%v", m)
			}
		}
	}

	tmp += " -lmingw32 -lqtmain -lshell32\n"
	tmp += "*/\n"

	tmp += fmt.Sprintf("import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_windows_386.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_windows_386.go"), tmp)
	}
}

func cgoLinux(module string) {
	var (
		tmp  string
		libs = cleanLibs(module)
	)

	tmp += fmt.Sprintf("package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	tmp += "/*\n"

	tmp += "#cgo CFLAGS: -pipe -O2 -Wall -W -D_REENTRANT -fPIC\n"
	tmp += "#cgo CXXFLAGS: -pipe -O2 -std=gnu++11 -Wall -W -D_REENTRANT -fPIC\n"

	tmp += "#cgo CXXFLAGS: -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/gcc/include -I/usr/local/Qt5.7.0/5.7/gcc/mkspecs/linux-g++\n"

	tmp += "#cgo CXXFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.7.0/5.7/gcc/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.7.0/5.7/gcc -Wl,-rpath,/usr/local/Qt5.7.0/5.7/gcc/lib\n"

	tmp += "#cgo LDFLAGS: -L/usr/local/Qt5.7.0/5.7/gcc/lib -L/usr/lib64"
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				tmp += fmt.Sprintf(" -lQt5%v", m)
			}
		}
	}

	tmp += " -lpthread\n"
	tmp += "*/\n"

	tmp += fmt.Sprintf("import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_amd64.go"), strings.Replace(tmp, "gcc", "gcc_64", -1))
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_amd64.go"), strings.Replace(tmp, "gcc", "gcc_64", -1))
	}
}

func cgoAndroidOnDarwinAndLinux(module string) {
	var (
		tmp  string
		libs = cleanLibs(module)
	)

	tmp += fmt.Sprintf("package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	tmp += "/*\n"

	tmp += "#cgo CFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC\n"
	tmp += "#cgo CXXFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -std=c++11 -O2 -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC\n"

	tmp += "#cgo CXXFLAGS: -DQT_NO_DEBUG"
	for i := len(libs) - 1; i >= 0; i-- {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(libs[i]))
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/android_armv7/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem /opt/android-ndk/platforms/android-9/arch-arm/usr/include -I/usr/local/Qt5.7.0/5.7/android_armv7/mkspecs/android-g++\n"

	tmp += "#cgo CXXFLAGS:"
	for i := len(libs) - 1; i >= 0; i-- {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.7.0/5.7/android_armv7/include/Qt%v", libs[i])
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: --sysroot=/opt/android-ndk/platforms/android-9/arch-arm -Wl,-rpath=/usr/local/Qt5.7.0/5.7/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition -Wl,--allow-shlib-undefined\n"

	tmp += "#cgo LDFLAGS: -L/opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android-ndk/platforms/android-9/arch-arm//usr/lib -L/usr/local/Qt5.7.0/5.7/android_armv7/lib -L/opt/android/ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android/ndk/platforms/android-9/arch-arm//usr/lib"

	for i := len(libs) - 1; i >= 0; i-- {
		if libs[i] != "UiPlugin" {
			tmp += fmt.Sprintf(" -lQt5%v", libs[i])
		}
	}

	tmp += " -lGLESv2 -lgnustl_shared -llog -lz -lm -ldl -lc -lgcc\n"
	tmp += "*/\n"

	tmp += fmt.Sprintf("import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_android_arm.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_android_arm.go"), tmp)
	}
}

func cgoAndroidOnWindows(module string) {
	var (
		tmp  string
		libs = cleanLibs(module)
	)

	tmp += fmt.Sprintf("package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	tmp += "/*\n"

	tmp += "#cgo CFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC\n"
	tmp += "#cgo CXXFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -std=c++11 -O2 -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC\n"

	tmp += "#cgo CXXFLAGS: -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -IC:/Qt/Qt5.7.0/5.7/android_armv7/include -isystem C:/android/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem C:/android/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem C:/android/android-ndk/platforms/android-9/arch-arm/usr/include -IC:/Qt/Qt5.7.0/5.7/android_armv7/mkspecs/android-g++\n"

	tmp += "#cgo CXXFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -IC:/Qt/Qt5.7.0/5.7/android_armv7/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: --sysroot=C:/android/android-ndk/platforms/android-9/arch-arm -Wl,-rpath=C:/Qt/Qt5.7.0/5.7/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition -Wl,--allow-shlib-undefined\n"

	tmp += "#cgo LDFLAGS: -LC:/android/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -LC:/android/android-ndk/platforms/android-9/arch-arm//usr/lib -LC:/Qt/Qt5.7.0/5.7/android_armv7/lib -LC:/android/android/ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -LC:/android/android/ndk/platforms/android-9/arch-arm//usr/lib"
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				tmp += fmt.Sprintf(" -lQt5%v", m)
			}
		}
	}

	tmp += " -lGLESv2 -lgnustl_shared -llog -lz -lm -ldl -lc -lgcc\n"
	tmp += "*/\n"

	tmp += fmt.Sprintf("import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_android_arm.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_android_arm.go"), tmp)
	}
}

func cgoIos(module string) string {
	var (
		tmp  string
		libs = cleanLibs(module)
	)

	tmp += fmt.Sprintf("package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	tmp += "/*\n"

	tmp += "#cgo CFLAGS: -pipe -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator9.3.sdk -mios-simulator-version-min=7.0 -arch i386 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC\n"
	tmp += "#cgo CXXFLAGS: -pipe -stdlib=libc++ -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -fvisibility-inlines-hidden -Wno-non-virtual-dtor -Wno-overloaded-virtual -Wno-exit-time-destructors -O2 -std=gnu++11 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator9.3.sdk -mios-simulator-version-min=7.0 -arch i386 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC\n"

	tmp += "#cgo CXXFLAGS: -DDARWIN_NO_CARBON -DQT_NO_PRINTER -DQT_NO_PRINTDIALOG -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/ios/mkspecs/macx-ios-clang/ios -I/usr/local/Qt5.7.0/5.7/ios/include -I/usr/local/Qt5.7.0/5.7/ios/mkspecs/macx-ios-clang\n"

	tmp += "#cgo CXXFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.7.0/5.7/ios/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator9.3.sdk -mios-simulator-version-min=7.0 -arch i386\n"

	tmp += "#cgo LDFLAGS: -L/usr/local/Qt5.7.0/5.7/ios/plugins/platforms -lqios_iphonesimulator -framework Foundation -framework UIKit -framework QuartzCore -framework AssetsLibrary -L/usr/local/Qt5.7.0/5.7/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework CoreText -framework CoreGraphics -framework OpenGLES -lqtfreetype_iphonesimulator -framework Security -framework SystemConfiguration -framework CoreBluetooth -L/usr/local/Qt5.7.0/5.7/ios/plugins/imageformats -lqdds_iphonesimulator -lqicns_iphonesimulator -lqico_iphonesimulator -lqtga_iphonesimulator -lqtiff_iphonesimulator -lqwbmp_iphonesimulator -lqwebp_iphonesimulator -lqtharfbuzzng_iphonesimulator -lz -lqtpcre_iphonesimulator -lm"
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -lQt5%v_iphonesimulator", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				tmp += fmt.Sprintf(" -lQt5%v_iphonesimulator", m)
			}
		}
	}

	if !strings.Contains(tmp, "-lQt5Gui_iphonesimulator") {
		tmp += " -lQt5Gui_iphonesimulator"
	}

	switch module {
	case
		"Multimedia":
		{
			tmp += " -lQt5OpenGL_iphonesimulator"
		}

	case "TestLib":
		{
			tmp += " -F/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest"
		}

	case "Qml", "WebChannel", "Quick":
		{
			if module != "Quick" {
				tmp += " -lQt5Quick_iphonesimulator -lQt5QuickParticles_iphonesimulator -lQt5QuickTest_iphonesimulator -lQt5QuickWidgets_iphonesimulator"
			}
			tmp += " -L/usr/local/Qt5.7.0/5.7/ios/plugins/qmltooling -lqmldbg_debugger_iphonesimulator -lqmldbg_inspector_iphonesimulator -lqmldbg_local_iphonesimulator -lqmldbg_native_iphonesimulator -lqmldbg_profiler_iphonesimulator -lqmldbg_quickprofiler_iphonesimulator -lqmldbg_server_iphonesimulator -lQt5PacketProtocol_iphonesimulator -lqmldbg_tcp_iphonesimulator"
		}

	case "Purchasing":
		{
			tmp += " -framework StoreKit"
		}
	}

	if module == "build_ios" {
		tmp += " -lQt5OpenGL_iphonesimulator"
		tmp += " -F/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest"
		tmp += " -lQt5Quick_iphonesimulator -lQt5QuickParticles_iphonesimulator -lQt5QuickTest_iphonesimulator -lQt5QuickWidgets_iphonesimulator"
		tmp += " -L/usr/local/Qt5.7.0/5.7/ios/plugins/qmltooling -lqmldbg_debugger_iphonesimulator -lqmldbg_inspector_iphonesimulator -lqmldbg_local_iphonesimulator -lqmldbg_native_iphonesimulator -lqmldbg_profiler_iphonesimulator -lqmldbg_quickprofiler_iphonesimulator -lqmldbg_server_iphonesimulator -lQt5PacketProtocol_iphonesimulator -lqmldbg_tcp_iphonesimulator"

		tmp += " -L/usr/local/Qt5.7.0/5.7/ios/qml/QtQuick.2 -lqtquick2plugin_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/qml/QtQuick/Layouts -lqquicklayoutsplugin_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/qml/QtQuick/Dialogs -ldialogplugin_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/qml/QtQuick/Controls -lqtquickcontrolsplugin_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/qml/Qt/labs/folderlistmodel -lqmlfolderlistmodelplugin_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/qml/Qt/labs/settings -lqmlsettingsplugin_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/qml/QtQuick/Dialogs/Private -ldialogsprivateplugin_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/qml/QtQuick/Window.2 -lwindowplugin_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/qml/QtQml/Models.2 -lmodelsplugin_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/qml/QtQuick/Extras -lqtquickextrasplugin_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/qml/QtGraphicalEffects/private -lqtgraphicaleffectsprivate_iphonesimulator"
		tmp += " -L/usr/local/Qt5.7.0/5.7/ios/qml/QtMultimedia -ldeclarative_multimedia_iphonesimulator -lQt5MultimediaQuick_p_iphonesimulator"
		tmp += " -L/usr/local/Qt5.7.0/5.7/ios/plugins/mediaservice -lqtmedia_audioengine_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/plugins/audio -lqtaudio_coreaudio_iphonesimulator -L/usr/local/Qt5.7.0/5.7/ios/plugins/playlistformats -lqtmultimedia_m3u_iphonesimulator -lQt5Multimedia_iphonesimulator"
		tmp += " -lqavfcamera_iphonesimulator -framework AudioToolbox -framework CoreAudio -framework AVFoundation -framework CoreMedia -framework CoreVideo -lqavfmediaplayer_iphonesimulator -lQt5MultimediaWidgets_iphonesimulator"
	}

	tmp += " -lQt5PlatformSupport_iphonesimulator\n"
	tmp += "*/\n"

	tmp += fmt.Sprintf("import \"C\"\n")

	if module == "build_ios" {
		return tmp
	}

	var path, prefix = func() (string, string) {
		if module == parser.MOC {
			return MocAppPath, "moc_"
		}
		return utils.GetQtPkgPath(strings.ToLower(module)), ""
	}()

	//386
	utils.Save(filepath.Join(path, prefix+"cgo_darwin_386.go"), tmp)

	//arm64
	tmp = strings.Replace(tmp, "_iphonesimulator", "", -1)
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
		tmp  = "// +build !android,!rpi1,!rpi2,!rpi3\n\n"
		libs = cleanLibs(module)
	)

	tmp += fmt.Sprintf("package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	tmp += "/*\n"

	tmp += "#cgo CFLAGS: -pipe -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -Wformat -Wformat-security -m32 -msse -msse2 -march=i686 -mfpmath=sse -mtune=generic -fno-omit-frame-pointer -fasynchronous-unwind-tables -fPIC -fvisibility=hidden -fvisibility-inlines-hidden -Wall -W -D_REENTRANT -fPIE\n"
	tmp += "#cgo CXXFLAGS: -pipe -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -Wformat -Wformat-security -m32 -msse -msse2 -march=i686 -mfpmath=sse -mtune=generic -fno-omit-frame-pointer -fasynchronous-unwind-tables -fPIC -fvisibility=hidden -fvisibility-inlines-hidden -Wall -W -D_REENTRANT -fPIE\n"

	tmp += "#cgo CXXFLAGS: -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -I/srv/mer/targets/SailfishOS-i486/usr/share/qt5/mkspecs/linux-g++ -I/srv/mer/targets/SailfishOS-i486/usr/include -I/srv/mer/targets/SailfishOS-i486/usr/include/sailfishapp -I/srv/mer/targets/SailfishOS-i486/usr/include/mdeclarativecache5 -I/srv/mer/targets/SailfishOS-i486/usr/include/qt5\n"

	tmp += "#cgo CXXFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -I/srv/mer/targets/SailfishOS-i486/usr/include/qt5/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath-link,/srv/mer/targets/SailfishOS-i486/usr/lib -Wl,-rpath-link,/srv/mer/targets/SailfishOS-i486/lib\n"

	tmp += "#cgo LDFLAGS: -rdynamic -L/srv/mer/targets/SailfishOS-i486/usr/lib -L/srv/mer/targets/SailfishOS-i486/lib -lsailfishapp -lmdeclarativecache5"
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			if IsWhiteListedSailfishLib(m) {
				tmp += fmt.Sprintf(" -lQt5%v", m)
			}
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				if IsWhiteListedSailfishLib(m) {
					tmp += fmt.Sprintf(" -lQt5%v", m)
				}
			}
		}
	}
	tmp += " -lGLESv2 -lpthread\n"
	tmp += "*/\n"

	tmp += fmt.Sprintf("import \"C\"\n")

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_386.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_386.go"), tmp)
	}

	tmp = strings.Replace(tmp, "-m32 -msse -msse2 -march=i686 -mfpmath=sse -mtune=generic -fno-omit-frame-pointer -fasynchronous-unwind-tables", "-fmessage-length=0 -march=armv7-a -mfloat-abi=hard -mfpu=neon -mthumb -Wno-psabi", -1)
	tmp = strings.Replace(tmp, "i486", "armv7hl", -1)

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_arm.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_arm.go"), tmp)
	}
}

func cgoRaspberryPi1(module string) {
	var (
		tmp  = "// +build rpi1\n\n"
		libs = cleanLibs(module)
	)

	tmp += fmt.Sprintf("package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	tmp += "/*\n"

	tmp += "#cgo CFLAGS: -pipe -marm -mfpu=vfp -mtune=arm1176jzf-s -march=armv6zk -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n"
	tmp += "#cgo CXXFLAGS: -pipe -marm -mfpu=vfp -mtune=arm1176jzf-s -march=armv6zk -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n"

	tmp += "#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi1/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I/usr/local/Qt5.7.0/5.7/rpi1/mkspecs/devices/linux-rasp-pi-g++\n"

	tmp += "#cgo CXXFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.7.0/5.7/rpi1/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,/usr/local/Qt5.7.0/5.7/rpi1/lib -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin\n"

	tmp += "#cgo LDFLAGS: -L/home/${USERNAME}/raspi/sysroot/opt/vc/lib -L/usr/local/Qt5.7.0/5.7/rpi1/lib"
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				tmp += fmt.Sprintf(" -lQt5%v", m)
			}
		}
	}

	tmp += " -lGLESv2 -lpthread\n"
	tmp += "*/\n"

	tmp += fmt.Sprintf("import \"C\"\n")

	var username = os.Getenv("USERNAME")
	if username == "" {
		username = "user"
	}

	tmp = strings.Replace(tmp, "${USERNAME}", username, -1)

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_rpi1.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_rpi1.go"), tmp)
	}
}

func cgoRaspberryPi2(module string) {
	var (
		tmp  = "// +build rpi2\n\n"
		libs = cleanLibs(module)
	)

	tmp += fmt.Sprintf("package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	tmp += "/*\n"

	tmp += "#cgo CFLAGS: -pipe -march=armv7-a -marm -mthumb-interwork -mfpu=neon-vfpv4 -mtune=cortex-a7 -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n"
	tmp += "#cgo CXXFLAGS: -pipe -march=armv7-a -marm -mthumb-interwork -mfpu=neon-vfpv4 -mtune=cortex-a7 -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n"

	tmp += "#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi2/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I/usr/local/Qt5.7.0/5.7/rpi2/mkspecs/devices/linux-rasp-pi2-g++\n"

	tmp += "#cgo CXXFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.7.0/5.7/rpi2/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,/usr/local/Qt5.7.0/5.7/rpi2/lib -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin\n"

	tmp += "#cgo LDFLAGS: -L/home/${USERNAME}/raspi/sysroot/opt/vc/lib -L/usr/local/Qt5.7.0/5.7/rpi2/lib"
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				tmp += fmt.Sprintf(" -lQt5%v", m)
			}
		}
	}

	tmp += " -lGLESv2 -lpthread\n"
	tmp += "*/\n"

	tmp += fmt.Sprintf("import \"C\"\n")

	var username = os.Getenv("USERNAME")
	if username == "" {
		username = "user"
	}

	tmp = strings.Replace(tmp, "${USERNAME}", username, -1)

	if module == parser.MOC {
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_rpi2.go"), tmp)
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_rpi2.go"), tmp)
	}
}

func cgoRaspberryPi3(module string) {
	var (
		tmp  = "// +build rpi3\n\n"
		libs = cleanLibs(module)
	)

	tmp += fmt.Sprintf("package %v\n\n", func() string {
		if MocModule != "" {
			return MocModule
		}
		return strings.ToLower(module)
	}())
	tmp += "/*\n"

	tmp += "#cgo CFLAGS: -march=armv8-a+crc -mtune=cortex-a53 -mfpu=crypto-neon-fp-armv8 -pipe -Os -mthumb -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n"
	tmp += "#cgo CXXFLAGS: -march=armv8-a+crc -mtune=cortex-a53 -mfpu=crypto-neon-fp-armv8 -pipe -Os -mthumb -std=c++1z -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC\n"

	tmp += "#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi3/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/${USERNAME}/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I/usr/local/Qt5.7.0/5.7/rpi3/mkspecs/devices/linux-rpi3-g++\n"

	tmp += "#cgo CXXFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.7.0/5.7/rpi3/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/${USERNAME}/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,/usr/local/Qt5.7.0/5.7/rpi3/lib -mfloat-abi=hard --sysroot=/home/${USERNAME}/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin\n"

	tmp += "#cgo LDFLAGS: -L/home/${USERNAME}/raspi/sysroot/opt/vc/lib -L/usr/local/Qt5.7.0/5.7/rpi3/lib"
	for _, m := range libs {
		if m == "UiTools" || m == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -lQt5%v", m)
		}
	}
	for _, m := range libs {
		if m != "UiTools" && m != "PlatformHeaders" {
			if m != "UiPlugin" {
				tmp += fmt.Sprintf(" -lQt5%v", m)
			}
		}
	}

	tmp += " -lGLESv2 -lpthread\n"
	tmp += "*/\n"

	tmp += fmt.Sprintf("import \"C\"\n")

	var username = os.Getenv("USERNAME")
	if username == "" {
		username = "user"
	}

	tmp = strings.Replace(tmp, "${USERNAME}", username, -1)

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
