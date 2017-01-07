package rcc

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func Rcc(appPath string, output_dir *string) {
	if output_dir == nil {
		var toutput_dir string
		output_dir = &toutput_dir
	}

	var appName = filepath.Base(appPath)

	utils.MkdirAll(filepath.Join(appPath, "qml"))

	var (
		rccPath string
		qmlGo   = filepath.Join(func() string {
			if *output_dir != "" {
				return *output_dir
			}
			return appPath
		}(), "rrc.go")

		qmlQrc = filepath.Join(appPath, "rrc.qrc")

		qmlCpp = filepath.Join(func() string {
			if *output_dir != "" {
				return *output_dir
			}
			return appPath
		}(), "rrc.cpp")
	)

	switch runtime.GOOS {
	case "darwin":
		{
			rccPath = filepath.Join(utils.QT_DARWIN_DIR(), "bin", "rcc")
		}

	case "linux":
		{
			if utils.UsePkgConfig() {
				rccPath = filepath.Join(strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=host_bins", "Qt5Core"), fmt.Sprintf("find rccPath with pkg-config on %v", runtime.GOOS))), "rcc")
			} else {
				rccPath = filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "gcc_64", "bin", "rcc")
			}
		}

	case "windows":
		{
			if utils.UseMsys2() {
				rccPath = filepath.Join(utils.QT_MSYS2_DIR(), "bin", "rcc")
			} else {
				rccPath = filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "mingw53_32", "bin", "rcc")
			}
		}
	}

	utils.Save(qmlGo, qmlHeader(appName))

	var rcc = exec.Command(rccPath, "-project", "-o", qmlQrc)
	rcc.Dir = filepath.Join(appPath, "qml")
	utils.RunCmd(rcc, fmt.Sprintf("execute rcc.1 on %v", runtime.GOOS))

	utils.Save(qmlQrc, strings.Replace(utils.Load(qmlQrc), "<file>./", "<file>qml/", -1))

	if utils.ExistsFile(filepath.Join(appPath, "qtquickcontrols2.conf")) {
		utils.Save(qmlQrc, strings.Replace(utils.Load(qmlQrc), "<qresource>", "<qresource>\n<file>qtquickcontrols2.conf</file>", -1))
	}

	var files, err = ioutil.ReadDir(appPath)
	if err != nil {
		utils.Log.WithError(err).Fatalln("failed to read dir")
	}

	var fileList = make([]string, 0)
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".qrc") {
			fileList = append(fileList, filepath.Join(appPath, file.Name()))
		}
	}

	rcc = exec.Command(rccPath, "-o", qmlCpp)
	rcc.Args = append(rcc.Args, fileList...)

	utils.RunCmd(rcc, fmt.Sprintf("execute rcc.2 on %v", runtime.GOOS))
}

func qmlHeader(appName string) string {

	return strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(`package main

/*
#cgo +build windows,386 windows,amd64 LDFLAGS: -L${QT_WINDOWS_DIR} -lQt5Core

#cgo +build !ios,darwin,amd64 LDFLAGS: -F${QT_DARWIN_DIR}/lib -framework QtCore

#cgo +build linux,amd64 LDFLAGS: -Wl,-rpath,${QT_LINUX_DIR} -L${QT_LINUX_DIR} -lQt5Core


#cgo +build android,linux,arm LDFLAGS: -L${QT_DIR}/${QT_VERSION_MAJOR}/android_armv7/lib -lQt5Core


#cgo +build ios,darwin,amd64 LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,${XCODE_DIR}/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/${IPHONESIMULATOR_SDK_DIR} -mios-simulator-version-min=7.0 -arch x86_64
#cgo +build ios,darwin,amd64 LDFLAGS: -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/plugins/platforms -lqios_iphonesimulator -framework Foundation -framework UIKit -framework QuartzCore -framework AssetsLibrary -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework CoreText -framework CoreGraphics -framework OpenGLES -lqtfreetype_iphonesimulator -framework Security -framework SystemConfiguration -framework CoreBluetooth -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/plugins/imageformats -lqdds_iphonesimulator -lqicns_iphonesimulator -lqico_iphonesimulator -lqtga_iphonesimulator -lqtiff_iphonesimulator -lqwbmp_iphonesimulator -lqwebp_iphonesimulator -lqtharfbuzzng_iphonesimulator -lz -lqtpcre_iphonesimulator -lm -lQt5Core_iphonesimulator -lQt5Widgets_iphonesimulator -lQt5Gui_iphonesimulator -lQt5PlatformSupport_iphonesimulator

#cgo +build ios,darwin,386 LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,${XCODE_DIR}/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/${IPHONESIMULATOR_SDK_DIR} -mios-simulator-version-min=7.0 -arch i386
#cgo +build ios,darwin,386 LDFLAGS: -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/plugins/platforms -lqios_iphonesimulator -framework Foundation -framework UIKit -framework QuartzCore -framework AssetsLibrary -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework CoreText -framework CoreGraphics -framework OpenGLES -lqtfreetype_iphonesimulator -framework Security -framework SystemConfiguration -framework CoreBluetooth -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/plugins/imageformats -lqdds_iphonesimulator -lqicns_iphonesimulator -lqico_iphonesimulator -lqtga_iphonesimulator -lqtiff_iphonesimulator -lqwbmp_iphonesimulator -lqwebp_iphonesimulator -lqtharfbuzzng_iphonesimulator -lz -lqtpcre_iphonesimulator -lm -lQt5Core_iphonesimulator -lQt5Widgets_iphonesimulator -lQt5Gui_iphonesimulator -lQt5PlatformSupport_iphonesimulator

#cgo +build ios,darwin,arm64 LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,${XCODE_DIR}/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/${IPHONEOS_SDK_DIR} -miphoneos-version-min=7.0 -arch arm64
#cgo +build ios,darwin,arm64 LDFLAGS: -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/plugins/platforms -lqios -framework Foundation -framework UIKit -framework QuartzCore -framework AssetsLibrary -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework CoreText -framework CoreGraphics -framework OpenGLES -lqtfreetype -framework Security -framework SystemConfiguration -framework CoreBluetooth -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/plugins/imageformats -lqdds -lqicns -lqico -lqtga -lqtiff -lqwbmp -lqwebp -lqtharfbuzzng -lz -lqtpcre -lm -lQt5Core -lQt5Widgets -lQt5Gui -lQt5PlatformSupport

#cgo +build ios,darwin,arm LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,${XCODE_DIR}/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/${IPHONEOS_SDK_DIR} -miphoneos-version-min=7.0 -arch armv7
#cgo +build ios,darwin,arm LDFLAGS: -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/plugins/platforms -lqios -framework Foundation -framework UIKit -framework QuartzCore -framework AssetsLibrary -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/lib -framework MobileCoreServices -framework CoreFoundation -framework CoreText -framework CoreGraphics -framework OpenGLES -lqtfreetype -framework Security -framework SystemConfiguration -framework CoreBluetooth -L${QT_DIR}/${QT_VERSION_MAJOR}/ios/plugins/imageformats -lqdds -lqicns -lqico -lqtga -lqtiff -lqwbmp -lqwebp -lqtharfbuzzng -lz -lqtpcre -lm -lQt5Core -lQt5Widgets -lQt5Gui -lQt5PlatformSupport


#cgo +build sailfish_emulator,linux,386 LDFLAGS: -Wl,-rpath,/usr/share/harbour-${APPNAME}/lib -Wl,-rpath-link,/srv/mer/targets/SailfishOS-i486/usr/lib -Wl,-rpath-link,/srv/mer/targets/SailfishOS-i486/lib -L/srv/mer/targets/SailfishOS-i486/usr/lib -L/srv/mer/targets/SailfishOS-i486/lib -lQt5Core
#cgo +build sailfish,linux,arm LDFLAGS: -Wl,-rpath,/usr/share/harbour-${APPNAME}/lib -Wl,-rpath-link,/srv/mer/targets/SailfishOS-armv7hl/usr/lib -Wl,-rpath-link,/srv/mer/targets/SailfishOS-armv7hl/lib -L/srv/mer/targets/SailfishOS-armv7hl/usr/lib -L/srv/mer/targets/SailfishOS-armv7hl/lib -lQt5Core


#cgo +build rpi1,linux,arm LDFLAGS: -Wl,-rpath-link,${RPI1_SYSROOT_DIR}/opt/vc/lib -Wl,-rpath-link,${RPI1_SYSROOT_DIR}/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,${RPI1_SYSROOT_DIR}/lib/arm-linux-gnueabihf -Wl,-rpath-link,${QT_DIR}/${QT_VERSION_MAJOR}/rpi1/lib -mfloat-abi=hard --sysroot=${RPI1_SYSROOT_DIR} -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin -L${RPI1_SYSROOT_DIR}/opt/vc/lib -L${QT_DIR}/${QT_VERSION_MAJOR}/rpi1/lib -lQt5Core
#cgo +build rpi2,linux,arm LDFLAGS: -Wl,-rpath-link,${RPI2_SYSROOT_DIR}/opt/vc/lib -Wl,-rpath-link,${RPI2_SYSROOT_DIR}/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,${RPI2_SYSROOT_DIR}/lib/arm-linux-gnueabihf -Wl,-rpath-link,${QT_DIR}/${QT_VERSION_MAJOR}/rpi2/lib -mfloat-abi=hard --sysroot=${RPI2_SYSROOT_DIR} -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin -L${RPI2_SYSROOT_DIR}/opt/vc/lib -L${QT_DIR}/${QT_VERSION_MAJOR}/rpi2/lib -lQt5Core
#cgo +build rpi3,linux,arm LDFLAGS: -Wl,-rpath-link,${RPI3_SYSROOT_DIR}/opt/vc/lib -Wl,-rpath-link,${RPI3_SYSROOT_DIR}/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,${RPI3_SYSROOT_DIR}/lib/arm-linux-gnueabihf -Wl,-rpath-link,${QT_DIR}/${QT_VERSION_MAJOR}/rpi3/lib -mfloat-abi=hard --sysroot=${RPI3_SYSROOT_DIR} -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin -L${RPI3_SYSROOT_DIR}/opt/vc/lib -L${QT_DIR}/${QT_VERSION_MAJOR}/rpi3/lib -lQt5Core
*/
import "C"`,
		"${QT_WINDOWS_DIR}", func() string {
			if runtime.GOOS == "linux" {
				if utils.QT_MXE_ARCH() == "amd64" {
					return "/usr/lib/mxe/usr/x86_64-w64-mingw32.shared/qt5/lib"
				}
				return "/usr/lib/mxe/usr/i686-w64-mingw32.shared/qt5/lib"
			}
			if utils.UseMsys2() {
				return filepath.Join(utils.QT_MSYS2_DIR(), "lib")
			}
			return "${QT_DIR}/${QT_VERSION_MAJOR}/mingw53_32/lib"
		}(), -1),
		"${QT_DARWIN_DIR}", utils.QT_DARWIN_DIR(), -1),
		"${QT_LINUX_DIR}", func() string {
			if utils.UsePkgConfig() {
				return strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=libdir", "Qt5Core"), fmt.Sprintf("find linux pkg-config lib dir on %v", runtime.GOOS)))
			}
			return "${QT_DIR}/${QT_VERSION_MAJOR}/gcc_64/lib"
		}(), -1),
		"${QT_DIR}", utils.QT_DIR(), -1),
		"${RPI1_SYSROOT_DIR}", utils.RPI1_SYSROOT_DIR(), -1),
		"${RPI2_SYSROOT_DIR}", utils.RPI2_SYSROOT_DIR(), -1),
		"${RPI3_SYSROOT_DIR}", utils.RPI3_SYSROOT_DIR(), -1),
		"${XCODE_DIR}", utils.XCODE_DIR(), -1),
		"${IPHONEOS_SDK_DIR}", utils.IPHONEOS_SDK_DIR(), -1),
		"${IPHONESIMULATOR_SDK_DIR}", utils.IPHONESIMULATOR_SDK_DIR(), -1),
		"${APPNAME}", appName, -1),
		"${QT_VERSION_MAJOR}", utils.QT_VERSION_MAJOR(), -1),
		"\\", "/", -1)
}
