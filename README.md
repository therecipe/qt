[![GoDoc](https://godoc.org/github.com/therecipe/qt?status.svg)](https://godoc.org/github.com/therecipe/qt)
[![GoCover](http://gocover.io/_badge/github.com/therecipe/qt)](http://gocover.io/github.com/therecipe/qt)
[![GoReportCard](https://goreportcard.com/badge/github.com/therecipe/qt)](https://goreportcard.com/report/github.com/therecipe/qt)

[![TravisCI](https://travis-ci.org/therecipe/qt.svg?branch=master)](https://travis-ci.org/therecipe/qt)
[![AppveyorCI](https://ci.appveyor.com/api/projects/status/github/therecipe/qt?branch=master&svg=true)](https://ci.appveyor.com/project/therecipe/qt)
[![SemaphoreCI](https://semaphoreci.com/api/v1/therecipe/qt/branches/master/shields_badge.svg)](https://semaphoreci.com/therecipe/qt)
[![CircleCI](https://circleci.com/gh/therecipe/qt/tree/master.svg?style=svg)](https://circleci.com/gh/therecipe/qt/tree/master)

## Overview

[Qt](https://en.wikipedia.org/wiki/Qt_(software)) is a cross-platform application framework.

[Go](https://en.wikipedia.org/wiki/Go_(programming_language)) (often referred to as golang) is a free and open source programming language created at Google.

This binding aims to make it as simple as possible to write applications for all operating systems supported by Qt in Go.

The project is pretty much a WIP and **not** recommended to be used in production yet.

However it should already contain everything you need to build fully featured Qt applications in Go.

[Screenshots of the Line Edits example](internal/screenshots)

## Installation

#### Full or stub installation?

The full installation requires at least **8gb** free ram and takes 20 min.

The stub installation requires only **1gb** free ram and takes 10 min. (**experimental**)

The only difference between those two version is, that you **won't** be able to use `go run/build` to build your applications if you choose to install the stub version.
You are therefore **limited** to the use of `qtdeploy` to build your application.

To build the stub version export `QT_STUB=true` system wide and proceed with the installation as usual.

#### Environmental variables

If you define environmental variables during the installation export them system wide, as they are needed not only during the installation but also later.

Add the environmental variables into your `$HOME/.bash_profile` or `$HOME/.profile` on macOS or Linux.

Or if you are on Windows simply use the advanced system settings.

#### Docker Images (**experimental**)

You can also use Docker images from [hub.docker](https://hub.docker.com/r/therecipe/qt/).

To cross compile to Windows, Linux and Android.

**You will still need to install the desktop version for your host system first!**

#### Command line debugging

You can export `QT_DEBUG=true` before "qtdeploying" your application to enable printing of the current function name at runtime.

#### Speedup your development

Instead of using `qtdeploy` during development, you might want to use `qtrcc`, `qtmoc`, `qtminimal` manually with `go build/run -tags=minimal ...`

In general `qtminimal` should always be called after the use of the optional use of `qtrcc` and `qtmoc`.

`qtrcc` is used to bundle your `project/qml` folder so you can access the files within by using the `qrc` prefix, like this `qrc:///qml/somefile` from various Qt functions.

`qtmoc` is used to subclass core.QObject based classes and to add your own signals/slots as shown in some examples.

`qtminimal` is used to create a small subset of the binding tailored to your code to reduce the binary size. (Use the `-tags=minimal` flag to make use of it)

If you use the STUB version, you are forced to use `qtminimal` otherwise it's optional like `qtrcc` and `qtmoc`.

#### Instructions

* Desktop
	* [Windows](#windows)
	* [macOS](#macos)
	* [Linux](#linux)

* Mobile
	* [Android](#android)
	* [iOS](#ios)
	* [Sailfish OS](#sailfish-os)

* Embedded
	* [Raspberry Pi](#raspberry-pi)

* [Docker](#docker)

## Windows

1. Install Go >= 1.7.1 and setup a proper [**GOPATH**](https://golang.org/doc/code.html#GOPATH)

	* https://golang.org/doc/install?download=go1.7.4.windows-amd64.msi

2. Install Qt 5.7.0 (with android support)

	* Install the official prebuilt package; you can also define a custom location with **QT_DIR**.
		* https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-windows-x86-android-5.7.0.exe

	or (without android support)

	* Install the Qt-dev package with [MSYS2](http://msys2.github.io) and define **QT_MSYS2=true** or define a custom Qt location with **QT_MSYS2_DIR** (usually: C:\msys32\ or C:\msys64\);
		* `pacman -Syyu`
		* if you want to deploy 32-bit applications `pacman -S mingw-w64-i686-qt-creator mingw-w64-i686-qt5`
		* if you want to deploy 64-bit applications `pacman -S mingw-w64-x86_64-qt-creator mingw-w64-x86_64-qt5`
		* `pacman -Scc`

3. Add the directory that contains **gcc** and **g++** to your **PATH** (not needed with MSYS2)

	* `C:\Qt\Qt5.7.0\Tools\mingw530_32\bin`

4. Download the binding

	* `go get -u -v github.com/therecipe/qt/cmd/...`

5. Generate, install and test (20 min) (MSYS2: run within the MSYS2 MinGW 32-bit or MSYS2 MinGW 64-bit shell)

	* `%GOPATH%\bin\qtsetup`

6. Create your first [application](#example)

7. Deploy applications with `%GOPATH%\bin\qtdeploy build desktop path\to\your\project` (MSYS2: run within the MSYS2 MinGW 32-bit or MSYS2 MinGW 64-bit shell)

## macOS

1. Install Go >= 1.7.1 and setup a proper [**GOPATH**](https://golang.org/doc/code.html#GOPATH)

	* https://golang.org/doc/install?download=go1.7.4.darwin-amd64.pkg

2. Install Qt 5.7.0 (with android/iOS support)

	* Install the official prebuilt package; you can also define a custom location with **QT_DIR**.
		* without iOS https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-mac-x64-android-5.7.0.dmg
		* with iOS https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-mac-x64-android-ios-5.7.0.dmg

	or (without android/iOS support)

	* Install the Qt-dev package with Homebrew and define **QT_HOMEBREW=true** or define a custom Qt location with **QT_DIR** (usually: /usr/local/opt/qt5/); if you want to link against homebrews Qt libs (**experimental**)
		* `brew install qt5`

3. Install **Xcode** >= 7.2.0; you can also define a custom location with **XCODE_DIR**

	* https://itunes.apple.com/us/app/xcode/id497799835

4. Download the binding

	* `go get -u -v github.com/therecipe/qt/cmd/...`

5. Generate, install and test (20 min)

	* `$GOPATH/bin/qtsetup`

6. Create your first [application](#example)

7. Deploy applications with `$GOPATH/bin/qtdeploy build desktop path/to/your/project`

## Linux

1. Install Go >= 1.7.1 and setup a proper [**GOPATH**](https://golang.org/doc/code.html#GOPATH)

	* https://golang.org/doc/install?download=go1.7.4.linux-amd64.tar.gz

2. Install Qt 5.7.0 (with android support)

	* Install the official prebuilt package; you can also define a custom location with **QT_DIR**.
		* https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-linux-x64-android-5.7.0.run

	or (without android support)

	* Install the Qt-dev package with your system package manager; if you want to link against your system Qt libs (**experimental**)
		* add **export QT_PKG_CONFIG=true** to your .profile or .bash_profile
		* if needed you can also define custom locations for the misc and doc dir with **QT_MISC_DIR** and/or **QT_DOC_DIR**
		* and you may want to define a custom pkg-config search path with **PKG_CONFIG_PATH**, if the default path points to old Qt pkg-config files

3. Install **g++** >= 5 and **OpenGL** dependencies

	* Debian/Ubuntu (apt-get)
		* `sudo apt-get -y install build-essential libgl1-mesa-dev`

	* Fedora/RHEL/CentOS (yum)
		* `sudo yum -y groupinstall "C Development Tools and Libraries"`
		* `sudo yum -y install mesa-libGL-devel`

	* openSUSE (zypper)
		* `sudo zypper -n install -t pattern devel_basis`

4. Download the binding

	* `go get -u -v github.com/therecipe/qt/cmd/...`

5. Generate, install and test (20 min)

	* `$GOPATH/bin/qtsetup`

6. Create your first [application](#example)

7. Deploy applications with `$GOPATH/bin/qtdeploy build desktop path/to/your/project` (use the *.sh file to start your application)

#### **Optional:** Cross compile for Windows on Debian/Ubuntu (**experimental**)

1. Install Wine

	* `sudo apt-get -y install wine`

2. Install MXE (M cross environment)

	* `echo "deb http://pkg.mxe.cc/repos/apt/debian wheezy main" | sudo tee --append /etc/apt/sources.list.d/mxeapt.list > /dev/null`

	* `sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys D43A795B73B16ABE9643FE1AFD8FFF16DB45C6AB`

	* `sudo apt-get update`

	* if you want to deploy 32-bit applications `sudo apt-get -y install mxe-i686-w64-mingw32.shared-qt3d mxe-i686-w64-mingw32.shared-qtactiveqt mxe-i686-w64-mingw32.shared-qtbase mxe-i686-w64-mingw32.shared-qtcanvas3d mxe-i686-w64-mingw32.shared-qtcharts mxe-i686-w64-mingw32.shared-qtconnectivity mxe-i686-w64-mingw32.shared-qtdatavis3d mxe-i686-w64-mingw32.shared-qtdeclarative mxe-i686-w64-mingw32.shared-qtdeclarative-render2d mxe-i686-w64-mingw32.shared-qtgamepad mxe-i686-w64-mingw32.shared-qtgraphicaleffects mxe-i686-w64-mingw32.shared-qtimageformats mxe-i686-w64-mingw32.shared-qtlocation mxe-i686-w64-mingw32.shared-qtmultimedia mxe-i686-w64-mingw32.shared-qtofficeopenxml mxe-i686-w64-mingw32.shared-qtpurchasing mxe-i686-w64-mingw32.shared-qtquickcontrols mxe-i686-w64-mingw32.shared-qtquickcontrols2 mxe-i686-w64-mingw32.shared-qtscript mxe-i686-w64-mingw32.shared-qtscxml mxe-i686-w64-mingw32.shared-qtsensors mxe-i686-w64-mingw32.shared-qtserialbus mxe-i686-w64-mingw32.shared-qtserialport mxe-i686-w64-mingw32.shared-qtservice mxe-i686-w64-mingw32.shared-qtsvg mxe-i686-w64-mingw32.shared-qtsystems mxe-i686-w64-mingw32.shared-qttools mxe-i686-w64-mingw32.shared-qttranslations mxe-i686-w64-mingw32.shared-qtvirtualkeyboard mxe-i686-w64-mingw32.shared-qtwebchannel mxe-i686-w64-mingw32.shared-qtwebkit mxe-i686-w64-mingw32.shared-qtwebsockets mxe-i686-w64-mingw32.shared-qtwinextras mxe-i686-w64-mingw32.shared-qtxlsxwriter mxe-i686-w64-mingw32.shared-qtxmlpatterns`

	* if you want to deploy 64-bit applications `sudo apt-get -y install mxe-x86-64-w64-mingw32.shared-qt3d mxe-x86-64-w64-mingw32.shared-qtactiveqt mxe-x86-64-w64-mingw32.shared-qtbase mxe-x86-64-w64-mingw32.shared-qtcanvas3d mxe-x86-64-w64-mingw32.shared-qtcharts mxe-x86-64-w64-mingw32.shared-qtconnectivity mxe-x86-64-w64-mingw32.shared-qtdatavis3d mxe-x86-64-w64-mingw32.shared-qtdeclarative mxe-x86-64-w64-mingw32.shared-qtdeclarative-render2d mxe-x86-64-w64-mingw32.shared-qtgamepad mxe-x86-64-w64-mingw32.shared-qtgraphicaleffects mxe-x86-64-w64-mingw32.shared-qtimageformats mxe-x86-64-w64-mingw32.shared-qtlocation mxe-x86-64-w64-mingw32.shared-qtmultimedia mxe-x86-64-w64-mingw32.shared-qtofficeopenxml mxe-x86-64-w64-mingw32.shared-qtpurchasing mxe-x86-64-w64-mingw32.shared-qtquickcontrols mxe-x86-64-w64-mingw32.shared-qtquickcontrols2 mxe-x86-64-w64-mingw32.shared-qtscript mxe-x86-64-w64-mingw32.shared-qtscxml mxe-x86-64-w64-mingw32.shared-qtsensors mxe-x86-64-w64-mingw32.shared-qtserialbus mxe-x86-64-w64-mingw32.shared-qtserialport mxe-x86-64-w64-mingw32.shared-qtservice mxe-x86-64-w64-mingw32.shared-qtsvg mxe-x86-64-w64-mingw32.shared-qtsystems mxe-x86-64-w64-mingw32.shared-qttools mxe-x86-64-w64-mingw32.shared-qttranslations mxe-x86-64-w64-mingw32.shared-qtvirtualkeyboard mxe-x86-64-w64-mingw32.shared-qtwebchannel mxe-x86-64-w64-mingw32.shared-qtwebkit mxe-x86-64-w64-mingw32.shared-qtwebsockets mxe-x86-64-w64-mingw32.shared-qtwinextras mxe-x86-64-w64-mingw32.shared-qtxlsxwriter mxe-x86-64-w64-mingw32.shared-qtxmlpatterns`

3. Export `QT_MXE_ARCH=386` to deploy 32-bit applications or `QT_MXE_ARCH=amd64` to deploy 64-bit applications.

4. Generate, install and test (20 min)

	* `$GOPATH/bin/qtsetup windows`

5. Deploy applications with `$GOPATH/bin/qtdeploy build windows path/to/your/project`

## Android

1. Install the desktop version for [Windows](#windows), [macOS](#macos) or [Linux](#linux)

2. Unzip the Android SDK in `C:\android-sdk-windows\` or `$HOME/android-sdk-macosx/` or `$HOME/android-sdk-linux/`; you can also define a custom location with **ANDROID_SDK_DIR**
	* https://dl.google.com/android/repository/tools_r25.2.4-windows.zip
	* https://dl.google.com/android/repository/tools_r25.2.4-macosx.zip
	* https://dl.google.com/android/repository/tools_r25.2.4-linux.zip

3. Install the SDK dependencies with `C:\android-sdk-windows\tools\android.bat` or `$HOME/android-sdk-{ macosx | linux }/tools/android`
	* Tools
		* Android SDK Build-tools (25.0.2)
	* Android 7.1.1 (API 25)
		* SDK Platform
	* Extras (Windows only)
		* Google USB Driver

4. Unzip the Android NDK in `C:\` or `$HOME`; you can also define a custom location with **ANDROID_NDK_DIR**
	* https://dl.google.com/android/repository/android-ndk-r13b-windows-x86_64.zip
	* https://dl.google.com/android/repository/android-ndk-r13b-darwin-x86_64.zip
	* https://dl.google.com/android/repository/android-ndk-r13b-linux-x86_64.zip

5. Install Java SE Development Kit >= 8 (Linux: install in `$HOME/jdk/`); you can also define a custom location with **JDK_DIR**
	* https://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html

6. Install and test (20 min)

	* `%GOPATH%\bin\qtsetup android`

		or

	* `$GOPATH/bin/qtsetup android`

7. Create your first [application](#example)

8. Deploy applications with `[GOPATH]/bin/qtdeploy build android path/to/your/project`

## iOS

1. Install the desktop version for [macOS](#macos)

2. Install and test (20 min)

	* `$GOPATH/bin/qtsetup ios && $GOPATH/bin/qtsetup ios-simulator`

3. Create your first [application](#example)

4. Deploy applications with `$GOPATH/bin/qtdeploy build ios path/to/your/project` or `$GOPATH/bin/qtdeploy build ios-simulator path/to/your/project`

## Sailfish OS

1. Install the desktop version for [Windows](#windows), [macOS](#macos) or [Linux](#linux)

2. Install VirtualBox; you can also define a custom location with **VIRTUALBOX_DIR**
	* http://download.virtualbox.org/virtualbox/5.1.10/VirtualBox-5.1.10-112026-Win.exe
	* http://download.virtualbox.org/virtualbox/5.1.10/VirtualBox-5.1.10-112026-OSX.dmg
	* http://download.virtualbox.org/virtualbox/5.1.10/VirtualBox-5.1.10-112026-Linux_amd64.run

3. Install the Sailfish OS SDK; you can also define a custom location with **SAILFISH_DIR**
	* https://releases.sailfishos.org/sdk/installers/1611/SailfishOSSDK-Beta-1611-Qt5-windows-offline.exe
	* https://releases.sailfishos.org/sdk/installers/1611/SailfishOSSDK-Beta-1611-Qt5-mac-offline.dmg
	* https://releases.sailfishos.org/sdk/installers/1611/SailfishOSSDK-Beta-1611-Qt5-linux-64-offline.run

4. Install and test (20 min)

	* `%GOPATH%\bin\qtsetup sailfish && %GOPATH%\bin\qtsetup sailfish-emulator`

		or

	* `$GOPATH/bin/qtsetup sailfish && $GOPATH/bin/qtsetup sailfish-emulator`

5. Create your first [application](#example)

6. Deploy applications with `[GOPATH]/bin/qtdeploy build sailfish path/to/your/project` or `[GOPATH]/bin/qtdeploy build sailfish-emulator path/to/your/project`

## Raspberry Pi

1. Install the desktop version for [Linux](#linux)

2. Create a folder `$HOME/raspi`

	* `mkdir $HOME/raspi`

3. Download and unpack the Qt source

	* `cd $HOME/raspi && wget https://download.qt.io/official_releases/qt/5.7/5.7.0/single/qt-everywhere-opensource-src-5.7.0.tar.gz`
	* `tar -xzf qt-everywhere-opensource-src-5.7.0.tar.gz qt-everywhere-opensource-src-5.7.0`

4. Patch Qt Source

	* `cd $HOME/raspi/qt-everywhere-opensource-src-5.7.0/qtbase && sed -i 's/c++1z/c++11/' ./mkspecs/devices/linux-rpi3-g++/qmake.conf`

	* `cd $HOME/raspi/qt-everywhere-opensource-src-5.7.0/qtwayland && wget https://github.com/qtproject/qtwayland/commit/75294be3.patch && patch -p1 -i 75294be3.patch`

5. Download the cross compiler; you can also define a custom location with **RPI_TOOLS_DIR**

	* `cd $HOME/raspi && git clone --depth 1 https://github.com/raspberrypi/tools.git`

6. Get dependencies and install Arch Linux on your SD card

	* `sudo apt-get -y install bsdtar libwayland-dev flex bison gperf python`

	* Raspberry Pi 1
		* https://archlinuxarm.org/platforms/armv6/raspberry-pi

	* Raspberry Pi 2
		* https://archlinuxarm.org/platforms/armv7/broadcom/raspberry-pi-2

	* Raspberry Pi 3
		* https://archlinuxarm.org/platforms/armv8/broadcom/raspberry-pi-3

7. Start your Raspberry Pi

8. Enable root login over ssh

	* `export RASPI_IP=192.168.XXX.XXX` (replace XXX.XXX with the valid ip ending)

	* `ssh alarm@$RASPI_IP` (password: alarm)

	* `su` (password: root)

	* `sed -i 's/#PermitRootLogin/PermitRootLogin/' /etc/ssh/sshd_config && sed -i 's/prohibit-password/yes/' /etc/ssh/sshd_config && systemctl restart sshd.service`

9. Update and install dependencies (5 min)

	* `pacman -Syyu`

	* `pacman -S fontconfig icu libinput libjpeg-turbo libproxy libsm libxi libxkbcommon-x11 libxrender tslib xcb-util-image xcb-util-keysyms xcb-util-wm freetds gtk3 libfbclient libmariadbclient mtdev postgresql-libs unixodbc assimp bluez-libs sdl2 jasper libmng libwebp gst-plugins-base-libs libpulse openal gst-plugins-bad hunspell libxcomposite wayland gst-plugins-base libxslt gst-plugins-good ffmpeg jsoncpp libevent libsrtp libvpx libxcursor libxrandr libxss libxtst nss opus protobuf snappy xcb-util xcb-util-cursor xcb-util-renderutil xcb-util-xrm libxfixes libxshmfence libxext libx11 libxcb libice weston ttf-freefont lxde gamin xorg-server xorg-xinit xorg-server-utils mesa xf86-video-fbdev xf86-video-vesa xorg-server-xwayland xf86-input-libinput gst-plugins-ugly sqlite2 cups xorg-server-devel rsync`

	* `pacman -Scc`

	* Raspberry Pi 1
		* `sed -i 's/gpu_mem=64/gpu_mem=128/' /boot/config.txt`
		* `echo "exec startlxde" >> $HOME/.xinitrc && mkdir $HOME/.config/ && echo -e "[core]\nbackend=fbdev-backend.so\nmodules=xwayland.so" >> $HOME/.config/weston.ini`
		* `echo "dtoverlay=vc4-kms-v3d,cma-128" >> /boot/config.txt && sed -i 's/fbdev-backend/drm-backend/' $HOME/.config/weston.ini` (**experimental**: enable OpenGL under X; will break most applications)

	* Raspberry Pi 2 or 3
		* `sed -i 's/gpu_mem=64/gpu_mem=256/' /boot/config.txt`
		* `echo "exec startlxde" >> $HOME/.xinitrc && mkdir $HOME/.config/ && echo -e "[core]\nbackend=fbdev-backend.so\nmodules=xwayland.so" >> $HOME/.config/weston.ini`
		* `echo "dtoverlay=vc4-kms-v3d,cma-256" >> /boot/config.txt && sed -i 's/fbdev-backend/drm-backend/' $HOME/.config/weston.ini` (**experimental**: enable OpenGL under X; will break most applications)

	* `reboot`

10. Get sysroot for cross compiling (password: root) (5 min)

	* `cd $HOME/raspi && mkdir sysroot sysroot/usr sysroot/opt`

	* `rsync -avz root@$RASPI_IP:/lib sysroot --delete`

	* `rsync -avz root@$RASPI_IP:/usr/include sysroot/usr --delete`

	* `rsync -avz root@$RASPI_IP:/usr/lib sysroot/usr --delete`

	* `rsync -avz root@$RASPI_IP:/opt/vc sysroot/opt --delete`

11. Prepare sysroot; you can also define a custom location with **RPI1_SYSROOT_DIR**, **RPI2_SYSROOT_DIR** or **RPI3_SYSROOT_DIR**

	* `cd $HOME/raspi && wget https://raw.githubusercontent.com/riscv/riscv-poky/master/scripts/sysroot-relativelinks.py`
	* `chmod +x sysroot-relativelinks.py && ./sysroot-relativelinks.py sysroot`

12. Build Qt (2 h)

	* `cd $HOME/raspi/qt-everywhere-opensource-src-5.7.0`

	* Raspberry Pi 1
		* `./configure -opengl es2 -device linux-rasp-pi-g++ -device-option CROSS_COMPILE=$HOME/raspi/tools/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf- -sysroot $HOME/raspi/sysroot -opensource -confirm-license -make libs -skip webengine -nomake tools -nomake examples -extprefix /usr/local/Qt5.7.0/5.7/rpi1 -I $HOME/raspi/sysroot/opt/vc/include -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I $HOME/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -silent`

	* Raspberry Pi 2
		* `./configure -opengl es2 -device linux-rasp-pi2-g++ -device-option CROSS_COMPILE=$HOME/raspi/tools/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf- -sysroot $HOME/raspi/sysroot -opensource -confirm-license -make libs -skip webengine -nomake tools -nomake examples -extprefix /usr/local/Qt5.7.0/5.7/rpi2 -I $HOME/raspi/sysroot/opt/vc/include -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I $HOME/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -silent`

	* Raspberry Pi 3
		* `./configure -opengl es2 -device linux-rpi3-g++ -device-option CROSS_COMPILE=$HOME/raspi/tools/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf- -sysroot $HOME/raspi/sysroot -opensource -confirm-license -make libs -skip webengine -nomake tools -nomake examples -extprefix /usr/local/Qt5.7.0/5.7/rpi3 -I $HOME/raspi/sysroot/opt/vc/include -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I $HOME/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -silent`

	* `make -k -i && sudo make -k -i install`

13. Prepare the Qt directory

	* `sudo chown -R $USER $HOME/Qt5.7.0/`

14. Install and test the binding (20 min)

	* Raspberry Pi 1
		* `$GOPATH/bin/qtsetup rpi1`

	* Raspberry Pi 2
		* `$GOPATH/bin/qtsetup rpi2`

	* Raspberry Pi 3
		* `$GOPATH/bin/qtsetup rpi3`

15. Notes

	* run `startx &` before starting an application with `-platform xcb` (qml/quick applications won't work; they may work with Qt 5.8)

	* run `weston &` or `weston --tty=1 &` (via ssh) or create your own [compositor](https://doc.qt.io/qt-5/qtwaylandcompositor-index.html) before starting an application with `-platform wayland-egl` (qml/quick applications won't work; they may work with Qt 5.8)

	* you can increase the available gpu memory by editing `/boot/config.txt`

16. Create your first [application](#example)

17. Deploy applications with `$GOPATH/bin/qtdeploy build rpiX path/to/your/project` (replace X with 1, 2 or 3)

## Docker

1. Install the desktop version for [Windows](#windows), [macOS](#macos) or [Linux](#linux)

2. Install and start [Docker](https://www.docker.com)

3. Share your **GOPATH** host directory as a [data volume](https://docs.docker.com/engine/tutorials/dockervolumes/#/mount-a-host-directory-as-a-data-volume) with Docker.
And make sure your Project folder is in your **GOPATH**.

4. Download the necessary Docker image(s) and run the setup.


* For cross compiling to Windows:

	* `docker pull therecipe/qt:base_windows` (to deploy 32-bit applications)

	* `docker pull therecipe/qt:base_windows_64` (to deploy 64-bit applications)

	* **Optional:** Install [Wine](https://www.winehq.org) to test your applications.

	* define `QT_MXE_ARCH=386` to deploy 32-bit applications or `QT_MXE_ARCH=amd64` to deploy 64-bit applications

	* `$GOPATH/bin/qtsetup windows-docker`

	* Deploy applications with `$GOPATH/bin/qtdeploy build windows path/to/your/project docker`

* For cross compiling to Linux:

	* `docker pull therecipe/qt:base`

	* `$GOPATH/bin/qtsetup linux-docker`

	* Deploy applications with `$GOPATH/bin/qtdeploy build linux path/to/your/project docker`

* For cross compiling to Android:

	* `docker pull therecipe/qt:base_android`

	* `$GOPATH/bin/qtsetup android-docker`

	* Deploy applications with `$GOPATH/bin/qtdeploy build android path/to/your/project docker`

## Code Editor Settings (for code completion)

* IntelliJ Idea

    Help > Edit Custom Properties
    ```
    # custom IntelliJ IDEA properties

    idea.max.intellisense.filesize=999999
    ```

    Help > Edit Custom VM Options

    ```
    # custom IntelliJ IDEA VM options

    -Xms128m
    -Xmx2000m
    -XX:ReservedCodeCacheSize=240m
    -XX:+UseCompressedOops
    ```

* Neovim

    Update your neocomplete.vim
    ```
    let g:neocomplete#skip_auto_completion_time = ""
    ```

## Example

1. Create a project folder `[GOPATH]/src/qtExample`

2. Create a file `[GOPATH]/src/qtExample/main.go`
	```go
package main

import (
		"os"

		"github.com/therecipe/qt/core"
		"github.com/therecipe/qt/widgets"
)

func main() {
		widgets.NewQApplication(len(os.Args), os.Args)

		//create a window
		var window = widgets.NewQMainWindow(nil, 0)
		window.SetWindowTitle("Hello World Example")
		window.SetMinimumSize2(200, 200)

		//create a layout
		var layout = widgets.NewQVBoxLayout()

		//add the layout to the centralWidget
		var centralWidget = widgets.NewQWidget(window, 0)
		centralWidget.SetLayout(layout)

		//create a button and connect the clicked signal
		var button = widgets.NewQPushButton2("Click me!", nil)
		button.ConnectClicked(func(flag bool) {
			widgets.QMessageBox_Information(nil, "OK", "You clicked me!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		})

		//add the button to the layout
		layout.AddWidget(button, 0, core.Qt__AlignCenter)

		//add the centralWidget to the window and show the window
		window.SetCentralWidget(centralWidget)
		window.Show()

		widgets.QApplication_Exec()
}
```

3. Open the terminal in `[GOPATH]/src` and run `[GOPATH]/bin/qtdeploy build desktop ./qtExample`

4. You will find the application in `[GOPATH]/src/qtExample/deploy/[GOOS]_minimal/`

5. Take a look at the [examples](https://github.com/therecipe/qt/tree/master/internal/examples)

6. Make yourself familiar with the [Qt documentation](https://doc.qt.io/qt-5/classes.html)

7. Join our Slack Channel [#qt-binding](https://gophers.slack.com/messages/qt-binding/details/)

## More extensive examples / demos / applications

* [5k3105/AM](https://github.com/5k3105/AM)
* [5k3105/Pipeline-Editor](https://github.com/5k3105/Pipeline-Editor)
* [gen2brain/url2img](https://github.com/gen2brain/url2img)
* [joeblew99/qt-archi](https://github.com/joeblew99/qt-archi)
* [jordanorelli/grpc-ui](https://github.com/jordanorelli/grpc-ui)
