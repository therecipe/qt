package templater

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

var (
	MocAppPath string
	MocModule string
)

func CopyCgo(module string) {

	if !strings.Contains(module, "droid") {
		createCgoDarwin(module)
		createCgoWindows(module)
		createCgoLinux(module)

		if runtime.GOOS == "darwin" {
			createCgoiOS(module)
		}
	}

	switch runtime.GOOS {
	case "darwin", "linux":
		{
			createCgoandroidDarwinAndLinux(module)
		}

	case "windows":
		{
			createCgoandroidWindows(module)
		}
	}
}

func createCgoDarwin(module string) {
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

	tmp += "#cgo CPPFLAGS: -pipe -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.8 -Wall -W -fPIC\n"

	tmp += "#cgo CPPFLAGS: -DQT_NO_DEBUG"
	for i := len(libs) - 1; i >= 0; i-- {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(libs[i]))
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/clang_64/mkspecs/macx-clang -F/usr/local/Qt5.6.0/5.6/clang_64/lib\n"

	tmp += "#cgo CPPFLAGS:"
	for i := len(libs) - 1; i >= 0; i-- {
		if libs[i] == "UiTools" || libs[i] == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -I/usr/local/Qt5.6.0/5.6/clang_64/include/Qt%v", libs[i])
		} else {
			tmp += fmt.Sprintf(" -I/usr/local/Qt5.6.0/5.6/clang_64/lib/Qt%v.framework/Headers", libs[i])
		}
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/AGL.framework/Headers\n"
	tmp += "\n"

	tmp += "#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.8 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/clang_64/lib\n"

	tmp += "#cgo LDFLAGS: -F/usr/local/Qt5.6.0/5.6/clang_64/lib"

	for i := len(libs) - 1; i >= 0; i-- {
		if libs[i] == "UiTools" || libs[i] == "PlatformHeaders" {
			tmp += fmt.Sprintf(" -L/usr/local/Qt5.6.0/5.6/clang_64/lib -lQt5%v", libs[i])
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

func createCgoWindows(module string) {
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

	tmp += "#cgo CPPFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra\n"

	tmp += "#cgo CPPFLAGS: -DUNICODE -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -IC:/Qt/Qt5.6.0/5.6/mingw49_32/include -IC:/Qt/Qt5.6.0/5.6/mingw49_32/mkspecs/win32-g++\n"

	tmp += "#cgo CPPFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -IC:/Qt/Qt5.6.0/5.6/mingw49_32/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition\n"

	tmp += "#cgo LDFLAGS: -LC:/Qt/Qt5.6.0/5.6/mingw49_32/lib"
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

func createCgoLinux(module string) {
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

	tmp += "#cgo CPPFLAGS: -pipe -O2 -Wall -W -D_REENTRANT\n"

	tmp += "#cgo CPPFLAGS: -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc/include -I/usr/local/Qt5.6.0/5.6/gcc/mkspecs/linux-g++\n"

	tmp += "#cgo CPPFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.6.0/5.6/gcc/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc/lib\n"

	tmp += "#cgo LDFLAGS: -L/usr/local/Qt5.6.0/5.6/gcc/lib -L/usr/lib64"
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
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_386.go"), strings.Replace(tmp, "lib64", "lib", -1))
		utils.Save(filepath.Join(MocAppPath, "moc_cgo_linux_amd64.go"), strings.Replace(tmp, "gcc", "gcc_64", -1))
	} else {
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_386.go"), strings.Replace(tmp, "lib64", "lib", -1))
		utils.Save(utils.GetQtPkgPath(strings.ToLower(module), "cgo_linux_amd64.go"), strings.Replace(tmp, "gcc", "gcc_64", -1))
	}
}

func createCgoandroidDarwinAndLinux(module string) {
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

	tmp += "#cgo CPPFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC\n"

	tmp += "#cgo CPPFLAGS: -DQT_NO_DEBUG"
	for i := len(libs) - 1; i >= 0; i-- {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(libs[i]))
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/android_armv7/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem /opt/android-ndk/platforms/android-9/arch-arm/usr/include -I/usr/local/Qt5.6.0/5.6/android_armv7/mkspecs/android-g++\n"

	tmp += "#cgo CPPFLAGS:"
	for i := len(libs) - 1; i >= 0; i-- {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.6.0/5.6/android_armv7/include/Qt%v", libs[i])
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: --sysroot=/opt/android-ndk/platforms/android-9/arch-arm/ -Wl,-rpath=/usr/local/Qt5.6.0/5.6/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition -Wl,--allow-shlib-undefined\n"

	tmp += "#cgo LDFLAGS: -L/opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android-ndk/platforms/android-9/arch-arm//usr/lib -L/usr/local/Qt5.6.0/5.6/android_armv7/lib -L/opt/android/ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android/ndk/platforms/android-9/arch-arm//usr/lib"

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

func createCgoandroidWindows(module string) {
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

	tmp += "#cgo CPPFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT\n"

	tmp += "#cgo CPPFLAGS: -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -IC:/Qt/Qt5.6.0/5.6/android_armv7/include -isystem C:/android/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem C:/android/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem C:/android/android-ndk/platforms/android-9/arch-arm/usr/include -IC:/Qt/Qt5.6.0/5.6/android_armv7/mkspecs/android-g++\n"

	tmp += "#cgo CPPFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -IC:/Qt/Qt5.6.0/5.6/android_armv7/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: --sysroot=C:/android/android-ndk/platforms/android-9/arch-arm/ -Wl,-rpath=C:/Qt/Qt5.6.0/5.6/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition -Wl,--allow-shlib-undefined\n"

	tmp += "#cgo LDFLAGS: -LC:/android/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -LC:/android/android-ndk/platforms/android-9/arch-arm//usr/lib -LC:/Qt/Qt5.6.0/5.6/android_armv7/lib -LC:/android/android/ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -LC:/android/android/ndk/platforms/android-9/arch-arm//usr/lib"
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

func createCgoiOS(module string) string {
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

	tmp += "#cgo CPPFLAGS: -pipe -fpascal-strings -fmessage-length=0 -Wno-trigraphs -Wreturn-type -Wparentheses -Wswitch -Wno-unused-parameter -Wunused-variable -Wunused-value -Wno-shorten-64-to-32 -Wno-sign-conversion -fexceptions -fasm-blocks -Wno-missing-field-initializers -Wno-missing-prototypes -Wno-implicit-atomic-properties -Wformat -Wno-missing-braces -Wno-unused-function -Wno-unused-label -Wuninitialized -Wno-unknown-pragmas -Wno-shadow -Wno-four-char-constants -Wno-sign-compare -Wpointer-sign -Wno-newline-eof -Wdeprecated-declarations -Winvalid-offsetof -Wno-conversion -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator9.3.sdk -mios-simulator-version-min=7.0 -arch i386 -fobjc-nonfragile-abi -fobjc-legacy-dispatch -Wno-deprecated-implementations -Wprotocol -Wno-selector -Wno-strict-selector-match -Wno-undeclared-selector -Wall -W -fPIC\n"

	tmp += "#cgo CPPFLAGS: -DDARWIN_NO_CARBON -DQT_NO_PRINTER -DQT_NO_PRINTDIALOG -DQT_NO_DEBUG"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -DQT_%v_LIB", strings.ToUpper(m))
	}
	tmp += "\n"

	tmp += "#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/ios/mkspecs/macx-ios-clang/ios -I/usr/local/Qt5.6.0/5.6/ios/include -I/usr/local/Qt5.6.0/5.6/ios/mkspecs/macx-ios-clang\n"

	tmp += "#cgo CPPFLAGS:"
	for _, m := range libs {
		tmp += fmt.Sprintf(" -I/usr/local/Qt5.6.0/5.6/ios/include/Qt%v", m)
	}
	tmp += "\n\n"

	tmp += "#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator9.3.sdk -mios-simulator-version-min=7.0 -arch i386\n"

	tmp += "#cgo LDFLAGS: -L/usr/local/Qt5.6.0/5.6/ios/plugins/platforms -lqios_iphonesimulator -framework Foundation -framework UIKit -framework QuartzCore -framework AssetsLibrary -L/usr/local/Qt5.6.0/5.6/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework CoreText -framework CoreGraphics -framework OpenGLES -lqtfreetype_iphonesimulator -framework Security -framework SystemConfiguration -framework CoreBluetooth -L/usr/local/Qt5.6.0/5.6/ios/plugins/imageformats -lqdds_iphonesimulator -lqicns_iphonesimulator -lqico_iphonesimulator -lqtga_iphonesimulator -lqtiff_iphonesimulator -lqwbmp_iphonesimulator -lqwebp_iphonesimulator -lqtharfbuzzng_iphonesimulator -lz -lqtpcre_iphonesimulator -lm"
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
			tmp += " -L/usr/local/Qt5.6.0/5.6/ios/plugins/qmltooling -lqmldbg_debugger_iphonesimulator -lqmldbg_inspector_iphonesimulator -lqmldbg_local_iphonesimulator -lqmldbg_native_iphonesimulator -lqmldbg_profiler_iphonesimulator -lqmldbg_server_iphonesimulator -lqmldbg_tcp_iphonesimulator"
		}
	}

	if module == "build_ios" {
		tmp += " -lQt5OpenGL_iphonesimulator"
		tmp += " -F/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest"
		tmp += " -lQt5Quick_iphonesimulator -lQt5QuickParticles_iphonesimulator -lQt5QuickTest_iphonesimulator -lQt5QuickWidgets_iphonesimulator"
		tmp += " -L/usr/local/Qt5.6.0/5.6/ios/plugins/qmltooling -lqmldbg_debugger_iphonesimulator -lqmldbg_inspector_iphonesimulator -lqmldbg_local_iphonesimulator -lqmldbg_native_iphonesimulator -lqmldbg_profiler_iphonesimulator -lqmldbg_server_iphonesimulator -lqmldbg_tcp_iphonesimulator"
		tmp += " -L/usr/local/Qt5.6.0/5.6/ios/qml/QtQuick.2 -lqtquick2plugin_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/qml/QtQuick/Layouts -lqquicklayoutsplugin_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/qml/QtQuick/Dialogs -ldialogplugin_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/qml/QtQuick/Controls -lqtquickcontrolsplugin_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/qml/Qt/labs/folderlistmodel -lqmlfolderlistmodelplugin_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/qml/Qt/labs/settings -lqmlsettingsplugin_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/qml/QtQuick/Dialogs/Private -ldialogsprivateplugin_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/qml/QtQuick/Window.2 -lwindowplugin_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/qml/QtQml/Models.2 -lmodelsplugin_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/qml/QtQuick/Extras -lqtquickextrasplugin_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/qml/QtGraphicalEffects/private -lqtgraphicaleffectsprivate_iphonesimulator"
		tmp += " -L/usr/local/Qt5.6.0/5.6/ios/qml/QtMultimedia -ldeclarative_multimedia_iphonesimulator -lQt5MultimediaQuick_p_iphonesimulator"
		tmp += " -L/usr/local/Qt5.6.0/5.6/ios/plugins/mediaservice -lqtmedia_audioengine_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/plugins/audio -lqtaudio_coreaudio_iphonesimulator -L/usr/local/Qt5.6.0/5.6/ios/plugins/playlistformats -lqtmultimedia_m3u_iphonesimulator -lQt5Multimedia_iphonesimulator"
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
	var tmp = createCgoiOS("build_ios")

	tmp = strings.Split(tmp, "/*")[1]
	tmp = strings.Split(tmp, "*/")[0]

	tmp = strings.Replace(tmp, "#cgo LDFLAGS: ", "", -1)
	tmp = strings.Replace(tmp, "#cgo CPPFLAGS: ", "", -1)
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
