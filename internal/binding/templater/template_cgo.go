package templater

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

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
			//fmt.Fprintf(bb, " -F%v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest", utils.XCODE_DIR())
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
		//fmt.Fprintf(bb, " -F%v/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/Library/Frameworks -weak_framework XCTest", utils.XCODE_DIR())
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

	if utils.QT_VERSION_NUM() >= 5090 {
		tmp := bb.String()
		bb.Reset()
		tmp = strings.Replace(tmp, "min=7.0", "min=8.0", -1)
		bb.WriteString(strings.Replace(tmp, "-lqtpcre", "-lqtpcre2", -1))
	}

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
		if mode == NONE {
			return ",!minimal"
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

func cleanLibs(module string, mode int) []string {

	var out []string

	switch {
	case mode == RCC:
		out = []string{"Core"}
	case mode == MOC, module == "build_ios":
		out = parser.LibDeps[module]
	case mode == MINIMAL, mode == NONE:
		out = append([]string{module}, parser.LibDeps[module]...)
	}

	for i, v := range out {
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
	}

	if buildARM == "armv7" {
		tmp = strings.Replace(tmp, "arm64", "armv7", -1)
	}
	if buildARM == "i386" {
		tmp = strings.Replace(tmp, "x86_64", "i386", -1)
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
