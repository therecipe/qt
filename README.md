[![GoDoc](https://godoc.org/github.com/therecipe/qt?status.svg)](https://godoc.org/github.com/therecipe/qt)

Introduction
------------

[Qt](https://en.wikipedia.org/wiki/Qt_(software)) is a cross-platform application framework that is used for developing application software that can be run on various software and hardware platforms with little or no change in the underlying codebase, while still being a native application with native capabilities and speed.

[Go](https://en.wikipedia.org/wiki/Go_(programming_language)) (often referred to as golang) is a free and open source programming language created at Google.

This binding is a thin layer on top of Qt's C++ API (with a one-to-one mapping) and allows you to write Qt applications entirely in Go, it comes along with some tools to make working with it more convenient.

Beside the main effort to ease the development of application software by making Qt's API accessible from Go, the second biggest effort went into simplifying the development and deployment processes.

[Screenshots](screenshots.md)

Status
------

**WIP**

Most of Qt's API is accessible from Go.

It should already contain everything you need to build fully featured applications.

At the moment webengine/webview packages do not work on Windows since they can't be compiled with MinGW.

If you still miss something, please open an issue.

Also there have been no release so far, so please pin the repo to a specific commit that is known to work for you.

Targets
-------

The following targets are currently supported:

| Target                   | Arch          | Linkage                          | Docker Deployment | Host OS             |
|:------------------------:|:-------------:|:--------------------------------:|:-----------------:|:-------------------:|
|         Windows          |   (32 / 64)   | (dynamic / static / system libs) |        Yes        |         Any         |
|     Android (+Wear)      |      arm      |             dynamic              |        Yes        |         Any         |
| Android-Emulator (+Wear) |      32       |             dynamic              |        No         | Windows/macOS/Linux |
|          Linux           |      64       |     (dynamic / system libs)      |        Yes        |         Any         |
|   Raspberry Pi (1/2/3)   |      arm      |     (dynamic / system libs)      |        Yes        |         Any         |
|          macOS           |      64       |     (dynamic / system libs)      |        No         |        macOS        |
|           iOS            | (arm + arm64) |              static              |        No         |        macOS        |
|      iOS-Simulator       |   (32 + 64)   |              static              |        No         |        macOS        |
|        SailfishOS        |      arm      |           system libs            |        No         | Windows/macOS/Linux |
|   SailfishOS-Emulator    |      32       |           system libs            |        No         | Windows/macOS/Linux |
|        AsteroidOS        |      arm      |           system libs            |        No         |        Linux        |

FAQ
---

-	What are the system requirements?

	-	2.5gb free ram (only needed during the setup) and at least 5gb free disk space.

-	How long does it take to get started?

	-	It takes 15-30 min to download and install everything.

-	Why is the setup so slow?

	-	The setup spends most of it's time installing the packages to speed up the compilation of your applications later.

-	Can I make the setup faster or lower the system requirements?

	-	Yes, you can export `QT_STUB=true` prior running the setup, and thereby also lower the system requirements.

-	Why is qtdeploy so slow?

	-	qtdeploy is slow, because it is intended to be used for deploying and therefore takes some extra actions that may not be needed during development.

-	Can I make qtdeploy faster?

	-	Yes, you can use the `-fast` flag to skip most of the actions taken otherwise and so drastically reduce the time it takes to compile your application.

Docker
------

The easiest way to get started is to use the pre-built [docker images](https://hub.docker.com/r/therecipe/qt/tags).

| Tag               | Compressed Size |
|:-----------------:|:---------------:|
|       linux       |      1 GB       |
|      android      |      3 GB       |
| windows_32_shared |      1 GB       |
| windows_32_static |      2 GB       |
| windows_64_shared |      1 GB       |
| windows_64_static |      2 GB       |
|       rpi1        |      2 GB       |
|       rpi2        |      2 GB       |
|       rpi3        |      2 GB       |

(If you want to switch between the 32/64-bit and the dynamic/static linkage for the windows target, you can use the **QT_MXE_ARCH** and **QT_MXE_STATIC** environmental variables like you would do if install MXE directly on your machine)

You can download an image by running: `docker pull therecipe/qt:TAG`

These images contain Go, Qt and all dependencies that may be necessary to deploy to your target.

You can look into the Dockerfiles that are used to create these images [here](https://github.com/therecipe/qt/blob/master/internal/docker).

The images can be used for the following:

-	The minimal setup to make it easier to get started.

-	The docker deployment to various targets during development or for CI or CD with `qtdeploy -docker build ...`

-	To provide pre-built (disposable) environments for CI or CD.

-	To provide template images to build upon for CI or CD.

Installation
------------

If you want to build applications for macOS or iOS, then you need to take the regular setup on macOS.

### Minimal Setup

-	Install Go: https://golang.org/dl

-	Install Docker: https://store.docker.com/search?type=edition&offering=community

-	Pull the Linux image: `docker pull therecipe/qt:linux`

-	Share your **GOPATH** with docker if needed (for example on macOS)

-	Clone the repo: `go get -v github.com/therecipe/qt/cmd/...`

-	Run the docker setup: `$GOPATH/bin/qtsetup -docker`

-	Pull additional images if you want to deploy to targets other than Linux

-	Read the [usage instructions](#tools) and build the [example](#hello-world)

### Regular Setup

Please make the environmental variables persistent if you override them.

You can do this by either editing you `$HOME/.bash_profile` or `$HOME/.profile` on macOS or Linux.

On Windows you can simply use the advanced system settings.

These are some general environmental variables, you can find the target specific environmental variables [here](#target-specific-infos-and-settings).

| Variable     | Default                        | Type   | Note                                                         |
|:------------:|:------------------------------:|:------:|:------------------------------------------------------------:|
|  QT_VERSION  |             5.8.0              | string |       can also be set by using the `-qt_version` flag        |
|    QT_DIR    | $HOME/Qt5.8.0 or C:\Qt\Qt5.8.0 | string |         can also be set by using the `-qt_dir` flag          |
|   QT_STUB    |             false              |  bool  |          is set to `true` during the minimal setup           |
|   QT_DEBUG   |             false              |  bool  | set to `true` if you want to print function names at runtime |
| QT_QMAKE_DIR |                                | string |       can be used to make use of custom versions of Qt       |

#### Desktop

##### Windows

<details markdown="1"> <summary markdown="0">Official (with Android support)</summary>

-	Install Go: https://golang.org/doc/install?download=go1.8.1.windows-amd64.msi

-	Clone the repo: `go get -u -v github.com/therecipe/qt/cmd/...`

-	Install Qt; you can also define a custom location with **QT_DIR**

	-	without Android support: https://download.qt.io/official_releases/qt/5.8/5.8.0/qt-opensource-windows-x86-mingw530-5.8.0.exe

	-	with Android support: https://download.qt.io/official_releases/qt/5.8/5.8.0/qt-opensource-windows-x86-android-5.8.0.exe

-	Add the directory that contains **gcc** to your **PATH**: `C:\Qt\Qt5.8.0\Tools\mingw530_32\bin` (or similar)

-	Run the setup: `%GOPATH%\bin\qtsetup`

-	Read the [usage instructions](#tools) and build the [example](#hello-world)

</details>

<details markdown="1"> <summary markdown="0">MSYS2</summary>

-	Install Go: https://golang.org/doc/install?download=go1.8.1.windows-amd64.msi

-	Clone the repo: `go get -u -v github.com/therecipe/qt/cmd/...`

-	Install [MSYS2](https://msys2.github.io) and export **QT_MSYS2=true**

-	Open a MinGW shell (the 64 bit version, if you want to deploy 64 bit applications)

-	`pacman -Syyu`

-	Install Qt

	-	to deploy dynamically linked 32-bit applications: `pacman -S mingw-w64-i686-qt-creator mingw-w64-i686-qt5`

	-	to deploy dynamically linked 64-bit applications: `pacman -S mingw-w64-x86_64-qt-creator mingw-w64-x86_64-qt5`

	-	to deploy statically linked 32-bit applications: `pacman -S mingw-w64-i686-qt-creator mingw-w64-i686-qt5-static`

	-	to deploy statically linked 64-bit applications: `pacman -S mingw-w64-x86_64-qt-creator mingw-w64-x86_64-qt5-static`

-	`pacman -Scc`

-	Use the MinGW shell for the setup and deployments (the 64 bit version, if you want to deploy 64 bit applications)

-	Export **QT_MSYS2_STATIC=true** if you want to deploy statically linked applications.

-	Run the setup: `%GOPATH%\bin\qtsetup`

-	Read the [usage instructions](#tools) and build the [example](#hello-world)

</details>

##### macOS

<details markdown="1"> <summary markdown="0">Official (with iOS/Android support)</summary>

-	Install Go: https://golang.org/doc/install?download=go1.8.1.darwin-amd64.pkg

-	Install Xcode

-	Clone the repo: `go get -u -v github.com/therecipe/qt/cmd/...`

-	Install Qt; you can also define a custom location with **QT_DIR**

	-	without Android or iOS support: https://download.qt.io/official_releases/qt/5.8/5.8.0/qt-opensource-mac-x64-clang-5.8.0.dmg

	-	with Android support: https://download.qt.io/official_releases/qt/5.8/5.8.0/qt-opensource-mac-x64-android-5.8.0.dmg

	-	with Android and iOS support: https://download.qt.io/official_releases/qt/5.8/5.8.0/qt-opensource-mac-x64-android-ios-5.8.0.dmg

-	Run the setup: `$GOPATH/bin/qtsetup`

-	Read the [usage instructions](#tools) and build the [example](#hello-world)

</details>

<details markdown="1"> <summary markdown="0">HomeBrew (no deployments possible)</summary>

-	Install Go: https://golang.org/doc/install?download=go1.8.1.darwin-amd64.pkg

-	Install Xcode

-	Clone the repo: `go get -u -v github.com/therecipe/qt/cmd/...`

-	Install [HomeBrew](https://brew.sh) and export **QT_HOMEBREW=true**

-	Install Qt: `brew install qt5`

-	Run the setup: `$GOPATH/bin/qtsetup`

-	Read the [usage instructions](#tools) and build the [example](#hello-world)

</details>

#### Linux

<details markdown="1"> <summary markdown="0">Official (with Android support)</summary>

-	Install Go: https://golang.org/doc/install?download=go1.8.1.linux-amd64.tar.gz

-	Clone the repo: `go get -u -v github.com/therecipe/qt/cmd/...`

-	Install Qt; you can also define a custom location with **QT_DIR**

	-	without Android support: https://download.qt.io/official_releases/qt/5.8/5.8.0/qt-opensource-linux-x64-5.8.0.run

	-	with Android support: https://download.qt.io/official_releases/qt/5.8/5.8.0/qt-opensource-linux-x64-android-5.8.0.run

-	Install **g++** >= 5 and **OpenGL** dependencies

	-	Debian/Ubuntu (apt-get)

		-	`sudo apt-get -y install build-essential libgl1-mesa-dev libpulse-dev`

	-	Fedora/RHEL/CentOS (yum)

		-	`sudo yum -y groupinstall "C Development Tools and Libraries"`
		-	`sudo yum -y install mesa-libGL-devel`

	-	openSUSE (zypper)

		-	`sudo zypper -n install -t pattern devel_basis`

-	Run the setup: `$GOPATH/bin/qtsetup`

-	Read the [usage instructions](#tools) and build the [example](#hello-world) (use the *.sh script to start your application)

</details>

<details markdown="1"> <summary markdown="0">Pkg-Config (no deployments possible)</summary>

-	Install Go: https://golang.org/doc/install?download=go1.8.1.linux-amd64.tar.gz

-	Clone the repo: `go get -u -v github.com/therecipe/qt/cmd/...`

-	Install the Qt-dev package with the help of your systems package manager

-	Export **QT_PKG_CONFIG=true**

-	Install **g++** >= 5 and **OpenGL** dependencies

	-	Debian/Ubuntu (apt-get)

		-	`sudo apt-get -y install build-essential libgl1-mesa-dev libpulse-dev`

	-	Fedora/RHEL/CentOS (yum)

		-	`sudo yum -y groupinstall "C Development Tools and Libraries"`
		-	`sudo yum -y install mesa-libGL-devel`

	-	openSUSE (zypper)

		-	`sudo zypper -n install -t pattern devel_basis`

-	Run the setup: `$GOPATH/bin/qtsetup`

-	Read the [usage instructions](#tools) and build the [example](#hello-world) (use the *.sh script to start your application)

</details>

<details markdown="1"> <summary markdown="0">Cross compile to Windows on Debian/Ubuntu</summary>

-	Install Wine: `sudo apt-get -y install wine`

-	Install MXE

	-	`echo "deb https://pkg.mxe.cc/repos/apt/debian wheezy main" | sudo tee --append /etc/apt/sources.list.d/mxeapt.list > /dev/null`

	-	`sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys D43A795B73B16ABE9643FE1AFD8FFF16DB45C6AB`

	-	`sudo apt-get update`

-	Install Qt

	-	to deploy dynamically linked 32-bit applications `sudo apt-get -y -qq install mxe-i686-w64-mingw32.shared-qt3d mxe-i686-w64-mingw32.shared-qtactiveqt mxe-i686-w64-mingw32.shared-qtbase mxe-i686-w64-mingw32.shared-qtcanvas3d mxe-i686-w64-mingw32.shared-qtcharts mxe-i686-w64-mingw32.shared-qtconnectivity mxe-i686-w64-mingw32.shared-qtdatavis3d mxe-i686-w64-mingw32.shared-qtdeclarative mxe-i686-w64-mingw32.shared-qtgamepad mxe-i686-w64-mingw32.shared-qtgraphicaleffects mxe-i686-w64-mingw32.shared-qtimageformats mxe-i686-w64-mingw32.shared-qtlocation mxe-i686-w64-mingw32.shared-qtmultimedia mxe-i686-w64-mingw32.shared-qtofficeopenxml mxe-i686-w64-mingw32.shared-qtpurchasing mxe-i686-w64-mingw32.shared-qtquickcontrols mxe-i686-w64-mingw32.shared-qtquickcontrols2 mxe-i686-w64-mingw32.shared-qtscript mxe-i686-w64-mingw32.shared-qtscxml mxe-i686-w64-mingw32.shared-qtsensors mxe-i686-w64-mingw32.shared-qtserialbus mxe-i686-w64-mingw32.shared-qtserialport mxe-i686-w64-mingw32.shared-qtservice mxe-i686-w64-mingw32.shared-qtsvg mxe-i686-w64-mingw32.shared-qtsystems mxe-i686-w64-mingw32.shared-qttools mxe-i686-w64-mingw32.shared-qttranslations mxe-i686-w64-mingw32.shared-qtvirtualkeyboard mxe-i686-w64-mingw32.shared-qtwebchannel mxe-i686-w64-mingw32.shared-qtwebkit mxe-i686-w64-mingw32.shared-qtwebsockets mxe-i686-w64-mingw32.shared-qtwinextras mxe-i686-w64-mingw32.shared-qtxlsxwriter mxe-i686-w64-mingw32.shared-qtxmlpatterns`

	-	to deploy dynamically linked 64-bit applications `sudo apt-get -y -qq install mxe-x86-64-w64-mingw32.shared-qt3d mxe-x86-64-w64-mingw32.shared-qtactiveqt mxe-x86-64-w64-mingw32.shared-qtbase mxe-x86-64-w64-mingw32.shared-qtcanvas3d mxe-x86-64-w64-mingw32.shared-qtcharts mxe-x86-64-w64-mingw32.shared-qtconnectivity mxe-x86-64-w64-mingw32.shared-qtdatavis3d mxe-x86-64-w64-mingw32.shared-qtdeclarative mxe-x86-64-w64-mingw32.shared-qtgamepad mxe-x86-64-w64-mingw32.shared-qtgraphicaleffects mxe-x86-64-w64-mingw32.shared-qtimageformats mxe-x86-64-w64-mingw32.shared-qtlocation mxe-x86-64-w64-mingw32.shared-qtmultimedia mxe-x86-64-w64-mingw32.shared-qtofficeopenxml mxe-x86-64-w64-mingw32.shared-qtpurchasing mxe-x86-64-w64-mingw32.shared-qtquickcontrols mxe-x86-64-w64-mingw32.shared-qtquickcontrols2 mxe-x86-64-w64-mingw32.shared-qtscript mxe-x86-64-w64-mingw32.shared-qtscxml mxe-x86-64-w64-mingw32.shared-qtsensors mxe-x86-64-w64-mingw32.shared-qtserialbus mxe-x86-64-w64-mingw32.shared-qtserialport mxe-x86-64-w64-mingw32.shared-qtservice mxe-x86-64-w64-mingw32.shared-qtsvg mxe-x86-64-w64-mingw32.shared-qtsystems mxe-x86-64-w64-mingw32.shared-qttools mxe-x86-64-w64-mingw32.shared-qttranslations mxe-x86-64-w64-mingw32.shared-qtvirtualkeyboard mxe-x86-64-w64-mingw32.shared-qtwebchannel mxe-x86-64-w64-mingw32.shared-qtwebkit mxe-x86-64-w64-mingw32.shared-qtwebsockets mxe-x86-64-w64-mingw32.shared-qtwinextras mxe-x86-64-w64-mingw32.shared-qtxlsxwriter mxe-x86-64-w64-mingw32.shared-qtxmlpatterns`

	-	to deploy statically linked 32-bit applications `sudo apt-get -y -qq install mxe-i686-w64-mingw32.static-qt3d mxe-i686-w64-mingw32.static-qtactiveqt mxe-i686-w64-mingw32.static-qtbase mxe-i686-w64-mingw32.static-qtcanvas3d mxe-i686-w64-mingw32.static-qtcharts mxe-i686-w64-mingw32.static-qtconnectivity mxe-i686-w64-mingw32.static-qtdatavis3d mxe-i686-w64-mingw32.static-qtdeclarative mxe-i686-w64-mingw32.static-qtgamepad mxe-i686-w64-mingw32.static-qtgraphicaleffects mxe-i686-w64-mingw32.static-qtimageformats mxe-i686-w64-mingw32.static-qtlocation mxe-i686-w64-mingw32.static-qtmultimedia mxe-i686-w64-mingw32.static-qtofficeopenxml mxe-i686-w64-mingw32.static-qtpurchasing mxe-i686-w64-mingw32.static-qtquickcontrols mxe-i686-w64-mingw32.static-qtquickcontrols2 mxe-i686-w64-mingw32.static-qtscript mxe-i686-w64-mingw32.static-qtscxml mxe-i686-w64-mingw32.static-qtsensors mxe-i686-w64-mingw32.static-qtserialbus mxe-i686-w64-mingw32.static-qtserialport mxe-i686-w64-mingw32.static-qtservice mxe-i686-w64-mingw32.static-qtsvg mxe-i686-w64-mingw32.static-qtsystems mxe-i686-w64-mingw32.static-qttools mxe-i686-w64-mingw32.static-qttranslations mxe-i686-w64-mingw32.static-qtvirtualkeyboard mxe-i686-w64-mingw32.static-qtwebchannel mxe-i686-w64-mingw32.static-qtwebsockets mxe-i686-w64-mingw32.static-qtwinextras mxe-i686-w64-mingw32.static-qtxlsxwriter mxe-i686-w64-mingw32.static-qtxmlpatterns`

	-	to deploy statically linked 64-bit applications `sudo apt-get -y -qq install mxe-x86-64-w64-mingw32.static-qt3d mxe-x86-64-w64-mingw32.static-qtactiveqt mxe-x86-64-w64-mingw32.static-qtbase mxe-x86-64-w64-mingw32.static-qtcanvas3d mxe-x86-64-w64-mingw32.static-qtcharts mxe-x86-64-w64-mingw32.static-qtconnectivity mxe-x86-64-w64-mingw32.static-qtdatavis3d mxe-x86-64-w64-mingw32.static-qtdeclarative mxe-x86-64-w64-mingw32.static-qtgamepad mxe-x86-64-w64-mingw32.static-qtgraphicaleffects mxe-x86-64-w64-mingw32.static-qtimageformats mxe-x86-64-w64-mingw32.static-qtlocation mxe-x86-64-w64-mingw32.static-qtmultimedia mxe-x86-64-w64-mingw32.static-qtofficeopenxml mxe-x86-64-w64-mingw32.static-qtpurchasing mxe-x86-64-w64-mingw32.static-qtquickcontrols mxe-x86-64-w64-mingw32.static-qtquickcontrols2 mxe-x86-64-w64-mingw32.static-qtscript mxe-x86-64-w64-mingw32.static-qtscxml mxe-x86-64-w64-mingw32.static-qtsensors mxe-x86-64-w64-mingw32.static-qtserialbus mxe-x86-64-w64-mingw32.static-qtserialport mxe-x86-64-w64-mingw32.static-qtservice mxe-x86-64-w64-mingw32.static-qtsvg mxe-x86-64-w64-mingw32.static-qtsystems mxe-x86-64-w64-mingw32.static-qttools mxe-x86-64-w64-mingw32.static-qttranslations mxe-x86-64-w64-mingw32.static-qtvirtualkeyboard mxe-x86-64-w64-mingw32.static-qtwebchannel mxe-x86-64-w64-mingw32.static-qtwebsockets mxe-x86-64-w64-mingw32.static-qtwinextras mxe-x86-64-w64-mingw32.static-qtxlsxwriter mxe-x86-64-w64-mingw32.static-qtxmlpatterns`

-	Export **QT_MXE_ARCH=386** to deploy 32-bit applications or **QT_MXE_ARCH=amd64** to deploy 64-bit applications.

-	Export **QT_MXE_STATIC=true** to deploy statically linked applications.

-	Run the setup: `$GOPATH/bin/qtsetup`

-	Read the [usage instructions](#tools) and build the [example](#hello-world) (use the *.sh script to start your application)

</details>

#### Mobile

<details markdown="1"> <summary markdown="0">Android / Android-Emulator</summary>

-	Install the desktop version for Windows, macOS or Linux

-	Unzip the Android SDK in `C:\android-sdk-windows\` or `$HOME/android-sdk-macosx/` or `$HOME/android-sdk-linux/`; you can also define a custom location with **ANDROID_SDK_DIR**

	-	https://dl.google.com/android/repository/tools_r25.2.5-windows.zip
	-	https://dl.google.com/android/repository/tools_r25.2.5-macosx.zip
	-	https://dl.google.com/android/repository/tools_r25.2.5-linux.zip

-	Install the SDK dependencies with `C:\android-sdk-windows\tools\android.bat` or `$HOME/android-sdk-{macosx|linux }/tools/android`

	-	Tools
		-	Android SDK Build-tools (26.0.0)
	-	Android 7.1.1 (API 25)
		-	SDK Platform
	-	Extras (Windows only)
		-	Google USB Driver

-	Unzip the Android NDK in `C:\` or `$HOME`; you can also define a custom location with **ANDROID_NDK_DIR**

	-	https://dl.google.com/android/repository/android-ndk-r14b-windows-x86_64.zip
	-	https://dl.google.com/android/repository/android-ndk-r14b-darwin-x86_64.zip
	-	https://dl.google.com/android/repository/android-ndk-r14b-linux-x86_64.zip

-	Install the JDK >= 8 (Linux: install in `$HOME/jdk/`); you can also define a custom location with **JDK_DIR**

	-	https://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html

-	Run the setup: `%GOPATH%\bin\qtsetup full android` or `$GOPATH/bin/qtsetup full android`

-	Read the [usage instructions](#tools) and build the [example](#hello-world)

</details>

<details markdown="1"> <summary markdown="0">iOS</summary>

-	Install the desktop version for macOS

-	Run the setup: `$GOPATH/bin/qtsetup full ios && $GOPATH/bin/qtsetup full ios-simulator`

-	Read the [usage instructions](#tools) and build the [example](#hello-world)

</details>

<details markdown="1"> <summary markdown="0">SailfishOS</summary>

-	Install the desktop version for Windows, macOS or Linux (**Qt <= 5.7.x required**\)

-	Install VirtualBox; you can also define a custom location with **VIRTUALBOX_DIR**

	-	https://download.virtualbox.org/virtualbox/5.1.18/VirtualBox-5.1.18-114002-Win.exe
	-	https://download.virtualbox.org/virtualbox/5.1.18/VirtualBox-5.1.18-114002-OSX.dmg
	-	https://download.virtualbox.org/virtualbox/5.1.18/VirtualBox-5.1.18-114002-Linux_amd64.run

-	Install the SailfishOS SDK; you can also define a custom location with **SAILFISH_DIR**

	-	https://releases.sailfishos.org/sdk/installers/1611/SailfishOSSDK-Beta-1611-Qt5-windows-offline.exe
	-	https://releases.sailfishos.org/sdk/installers/1611/SailfishOSSDK-Beta-1611-Qt5-mac-offline.dmg
	-	https://releases.sailfishos.org/sdk/installers/1611/SailfishOSSDK-Beta-1611-Qt5-linux-64-offline.run

-	Run the setup:

	-	`%GOPATH%\bin\qtsetup full sailfish && %GOPATH%\bin\qtsetup full sailfish-emulator`

	-	`$GOPATH/bin/qtsetup full sailfish && $GOPATH/bin/qtsetup full sailfish-emulator`

-	Read the [usage instructions](#tools) and build the [example](#hello-world)

</details>

<details markdown="1"> <summary markdown="0">AsteroidOS</summary>

-	**TBA**

</details>

#### Embedded

##### Raspberry Pi

<details markdown="1"> <summary markdown="0">Raspbian + pre-built Qt 5.7</summary>

-	Install the desktop version for Linux

-	Install Raspbian on your rpi: https://downloads.raspberrypi.org/raspbian/images/raspbian-2017-03-03/2017-03-02-raspbian-jessie.zip

-	Follow the instructions: https://www.qtrpi.com/faq#howtoinstall

-	`export QT_QMAKE_DIR=/opt/qtrpi/raspi/qt5/bin`

-	`export RPI_TOOLS_DIR=/opt/qtrpi/raspi/tools`

-	`export RPI_COMPILER=gcc-linaro-arm-linux-gnueabihf-raspbian-x64`

-	Run the setup: `$GOPATH/bin/qtsetup full rpiX` (replace X with: 1, 2 or 3)

-	Read the [usage instructions](#tools) and build the [example](#hello-world) (use the *.sh script to start your application)

</details>

<details markdown="1"> <summary markdown="0">Arch Linux + self-compiled Qt 5.8</summary>

-	Install the desktop version for Linux

-	`mkdir $HOME/raspi`

-	Download and unpack the Qt source

	-	`cd $HOME/raspi && wget https://download.qt.io/official_releases/qt/5.8/5.8.0/single/qt-everywhere-opensource-src-5.8.0.tar.gz`
	-	`tar -xzf qt-everywhere-opensource-src-5.8.0.tar.gz qt-everywhere-opensource-src-5.8.0`

-	Download the cross compiler; you can also define a custom location with **RPI_TOOLS_DIR** (but then you might need to manually change commands from here on during the setup)

	-	`cd $HOME/raspi && git clone --depth 1 https://github.com/raspberrypi/tools.git`

-	Get dependencies and install Arch Linux on your SD card

	-	`sudo apt-get -y install bsdtar libwayland-dev flex bison gperf python`

	-	Raspberry Pi 1: https://archlinuxarm.org/platforms/armv6/raspberry-pi

	-	Raspberry Pi 2: https://archlinuxarm.org/platforms/armv7/broadcom/raspberry-pi-2

	-	Raspberry Pi 3: https://archlinuxarm.org/platforms/armv8/broadcom/raspberry-pi-3

-	Start your Raspberry Pi

-	Enable root login over ssh

	-	`export RASPI_IP=192.168.XXX.XXX` (replace XXX.XXX with the valid ip ending)

	-	`ssh alarm@$RASPI_IP` (password: alarm)

	-	`su` (password: root)

	-	`sed -i 's/#PermitRootLogin/PermitRootLogin/' /etc/ssh/sshd_config && sed -i 's/prohibit-password/yes/' /etc/ssh/sshd_config && systemctl restart sshd.service`

-	Update and install dependencies (5 min)

	-	`pacman -Syyu`

	-	`pacman -S fontconfig icu libinput libjpeg-turbo libproxy libsm libxi libxkbcommon-x11 libxrender tslib xcb-util-image xcb-util-keysyms xcb-util-wm freetds gtk3 libfbclient libmariadbclient mtdev postgresql-libs unixodbc assimp bluez-libs sdl2 jasper libmng libwebp gst-plugins-base-libs libpulse openal gst-plugins-bad hunspell libxcomposite wayland gst-plugins-base libxslt gst-plugins-good ffmpeg jsoncpp libevent libsrtp libvpx libxcursor libxrandr libxss libxtst nss opus protobuf snappy xcb-util xcb-util-cursor xcb-util-renderutil xcb-util-xrm libxfixes libxshmfence libxext libx11 libxcb libice weston ttf-freefont lxde gamin xorg-server xorg-xinit xorg-server-utils mesa xf86-video-fbdev xf86-video-vesa xorg-server-xwayland xf86-input-libinput gst-plugins-ugly sqlite2 cups xorg-server-devel rsync`

	-	`pacman -Scc`

	-	Raspberry Pi 1

		-	`sed -i 's/gpu_mem=64/gpu_mem=128/' /boot/config.txt`
		-	`echo "exec startlxde" >> $HOME/.xinitrc && mkdir $HOME/.config/ && echo -e "[core]\nbackend=fbdev-backend.so\nmodules=xwayland.so" >> $HOME/.config/weston.ini`

		**experimental**: enable OpenGL under X; will break most applications

		-	`echo "dtoverlay=vc4-kms-v3d,cma-128" >> /boot/config.txt && sed -i 's/fbdev-backend/drm-backend/' $HOME/.config/weston.ini`

	-	Raspberry Pi 2 or 3

		-	`sed -i 's/gpu_mem=64/gpu_mem=256/' /boot/config.txt`
		-	`echo "exec startlxde" >> $HOME/.xinitrc && mkdir $HOME/.config/ && echo -e "[core]\nbackend=fbdev-backend.so\nmodules=xwayland.so" >> $HOME/.config/weston.ini`

		**experimental**: enable OpenGL under X; will break most applications

		-	`echo "dtoverlay=vc4-kms-v3d,cma-256" >> /boot/config.txt && sed -i 's/fbdev-backend/drm-backend/' $HOME/.config/weston.ini`

	-	`reboot`

-	Get sysroot for cross compiling (password: root) (5 min)

	-	`cd $HOME/raspi && mkdir sysroot sysroot/usr sysroot/opt`

	-	`rsync -avz root@$RASPI_IP:/lib sysroot --delete`

	-	`rsync -avz root@$RASPI_IP:/usr/include sysroot/usr --delete`

	-	`rsync -avz root@$RASPI_IP:/usr/lib sysroot/usr --delete`

	-	`rsync -avz root@$RASPI_IP:/opt/vc sysroot/opt --delete`

-	Prepare sysroot; you can also define a custom location with **RPI1_SYSROOT_DIR**, **RPI2_SYSROOT_DIR** or **RPI3_SYSROOT_DIR** (but then you might need to manually change commands from here on during the setup)

	-	`cd $HOME/raspi && wget https://raw.githubusercontent.com/riscv/riscv-poky/master/scripts/sysroot-relativelinks.py`
	-	`chmod +x sysroot-relativelinks.py && ./sysroot-relativelinks.py sysroot`

-	Build Qt (2 hours)

	-	`cd $HOME/raspi/qt-everywhere-opensource-src-5.8.0`

	-	make sure **QT_DIR** points to your desktop installation of Qt; you may also want to tweak the configure command below, if you put the tools or the sysroot in an alternative location

	-	Raspberry Pi 1

		-	`./configure -opengl es2 -device linux-rasp-pi-g++ -device-option CROSS_COMPILE=$HOME/raspi/tools/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf- -sysroot $HOME/raspi/sysroot -opensource -confirm-license -make libs -skip webengine -nomake tools -nomake examples -extprefix $QT_DIR/5.8/rpi1 -I $HOME/raspi/sysroot/opt/vc/include -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I $HOME/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -silent`

	-	Raspberry Pi 2

		-	`./configure -opengl es2 -device linux-rasp-pi2-g++ -device-option CROSS_COMPILE=$HOME/raspi/tools/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf- -sysroot $HOME/raspi/sysroot -opensource -confirm-license -make libs -skip webengine -nomake tools -nomake examples -extprefix $QT_DIR/5.8/rpi2 -I $HOME/raspi/sysroot/opt/vc/include -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I $HOME/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -silent`

	-	Raspberry Pi 3

		-	`./configure -opengl es2 -device linux-rpi3-g++ -device-option CROSS_COMPILE=$HOME/raspi/tools/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf- -sysroot $HOME/raspi/sysroot -opensource -confirm-license -make libs -skip webengine -nomake tools -nomake examples -extprefix $QT_DIR/5.8/rpi3 -I $HOME/raspi/sysroot/opt/vc/include -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos -I $HOME/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I $HOME/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -silent`

	-	`make -k -i && sudo make -k -i install`

-	Prepare the Qt directory: `sudo chown -R $USER $QT_DIR`

-	Run the setup: `$GOPATH/bin/qtsetup full rpiX` (replace X with: 1, 2 or 3)

-	Read the [usage instructions](#tools) and build the [example](#hello-world) (use the *.sh script to start your application)

-	Notes:

	-	run `startx &` before starting an application with `-platform xcb`

	-	run `weston &` or `weston --tty=1 &` (via ssh) or create your own [compositor](https://doc.qt.io/qt-5/qtwaylandcompositor-index.html) before starting an application with `-platform wayland-egl`

	-	you can increase the available gpu memory by editing `/boot/config.txt`

</details>

##### Additional Modules

<details markdown="1"> <summary markdown="0">WebKit</summary>

-	Download the WebKit module:

	-	https://github.com/annulen/webkit/releases/download/qtwebkit-tp5/qtwebkit-tp5-qt58-mingw530-x86.zip
	-	https://github.com/annulen/webkit/releases/download/qtwebkit-tp5/qtwebkit-tp5-qt58-darwin-x64.tar.xz
	-	https://github.com/annulen/webkit/releases/download/qtwebkit-tp5/qtwebkit-tp5-qt58-linux-x64.tar.xz

-	Extract it inside `QT_DIR/5.8/{mingw53_32|clang_64|gcc_64}`

-	Export `QT_WEBKIT=true`

-	Run the setup

</details>

Target Specific Infos And Settings
----------------------------------

#### Desktop

<details markdown="1"> <summary markdown="0">Windows</summary>

### General

You may need to export `CGO_ENABLED=1` and `GOARCH=386` if you want to use `go build` on windows.

Depending on your editor or IDE you may also need to export `GOARCH=386` to make the code completion work.

It's recommended to use `qtdeploy -fast build path/to/your/project` instead of the direct usage of `go build`.

### Environmental Variables

#### Windows Host

| Variable        | Default                | Type   | Note                                                                            |
|:---------------:|:----------------------:|:------:|:-------------------------------------------------------------------------------:|
|    QT_MSYS2     |         false          |  bool  |            set to `true` if you want to use the msys2 version of Qt             |
|  QT_MSYS2_DIR   | C:\msys32 or C:\msys64 | string | `qtsetup check` will display this with the prefix detected by **QT_MSYS2_ARCH** |
|  QT_MSYS2_ARCH  |          386           | string |            set to `amd64` if you want to deploy 64 bit applications             |
| QT_MSYS2_STATIC |         false          |  bool  |       set to `true` if you want to deploy statically linked applications        |

-	**QT_MSYS2** and **QT_MSYS2_ARCH** may be removed as they are now automatically detected if you use the mingw shells

-	**QT_MSYS2_DIR** may be refined to display only the current value (without the arch dir)

#### Linux or macOS Host

| Variable      | Default      | Type   | Note                                                               |
|:-------------:|:------------:|:------:|:------------------------------------------------------------------:|
|  QT_MXE_DIR   | /usr/lib/mxe | string |         can be set to use mxe from an alternative location         |
|  QT_MXE_ARCH  |     386      | string |      set to `amd64` if you want to deploy 64 bit applications      |
| QT_MXE_STATIC |    false     |  bool  | set to `true` if you want to deploy statically linked applications |

### Branding

You can use the `windows` folder at the root of your project to let `qtdeploy` automatically bundle your assets.

If you want to use a custom icon for your binary, you may need to use `windres`.

To do so, create a file called `icon.rc` with the following content in your root project folder.

-	`IDI_ICON1 ICON DISCARDABLE "icon.ico"`

Rename your icon to `icon.ico` and place it alongside the `icon.rc`

Run `windres` like this: `path/to/windres path/to/project/icon.rc -o path/to/project/icon_windows.syso`

`windres` can be usually found inside the same folder as rcc on a windows system (for example: `Qt5.8.0/5.8/mingw53_32/bin/windres.exe`\)

Use `qtdeploy` or `go build` as usuall, the `*.syso` file should be automatically detected and be added to your binary.

You can find more infos here: https://doc.qt.io/qt-5/appicon.html

</details>

<details markdown="1"> <summary markdown="0">macOS</summary>

### Environmental Variables

| Variable    | Default                 | Type   | Note                                                        |
|:-----------:|:-----------------------:|:------:|:-----------------------------------------------------------:|
| QT_HOMEBREW |          false          |  bool  | set to `true` if you want to use the homebrew version of Qt |
|  XCODE_DIR  | /Applications/Xcode.app | string |    can be set to use Xcode from an alternative location     |

### Branding

You can use the `darwin` folder at the root of your project to let `qtdeploy` automatically bundle your assets.

```
project_name
├── darwin
│   └── Contents
│       ├── Info.plist
│       └── Resources
│           └── project_name.icns
├── project.go
└── qml
    └── project.qml
```

Furthermore you can use the `Info.plist` found inside `deploy/darwin/project_name.app` as a template to build upon.

</details>

<details markdown="1"> <summary markdown="0">Linux</summary>

### Environmental Variables

| Variable      | Default                        | Type   | Note                                                                                |
|:-------------:|:------------------------------:|:------:|:-----------------------------------------------------------------------------------:|
|   QT_DISTRO   | arch/fedora/suse/ubuntu/gentoo | string |  set this to another distro if you want to change the values for DOC and MISC dir   |
| QT_PKG_CONFIG |             false              |  bool  | set this to `true` if you want to use the Qt version of your systems package manger |
|  QT_DOC_DIR   |    "depends on the distro"     | string |               set this to the documentation (*.index) folder location               |
|  QT_MISC_DIR  |    "depends on the distro"     | string |                                                                                     |

### Branding

You can use the `linux` folder at the root of your project to let `qtdeploy` automatically bundle your assets.

</details>

#### Mobile

<details markdown="1"> <summary markdown="0">Android / Android-Emulator</summary>

### Environmental Variables

| Variable        | Default                                                                        | Type   | Note                                            |
|:---------------:|:------------------------------------------------------------------------------:|:------:|:-----------------------------------------------:|
|     JDK_DIR     |                                   JAVA_HOME                                    | string |    can be also set by setting **JAVA_HOME**     |
| ANDROID_SDK_DIR | $HOME/android-sdk-linux or $HOME/android-sdk-macosx or C:\\android-sdk-windows | string | can be also set by setting **ANDROID_SDK_ROOT** |
| ANDROID_NDK_DIR |                 $HOME/android-ndk-r14b or C:\android-ndk-r14b                  | string | can be also set by setting **ANDROID_NDK_ROOT** |

### Branding

You can use the `android` or `android-emulator` folder at the root of your project to let `qtdeploy` automatically bundle your assets.

```
project_name
├── android
│   ├── AndroidManifest.xml
│   ├── alias.txt
│   ├── build.gradle
│   ├── gradle
│   │   └── wrapper
│   │       └── gradle-wrapper.properties
│   ├── icon.png
│   ├── password.txt
│   ├── project_name.keystore
│   └── res
│       ├── drawable-hdpi
│       │   └── icon.png
│       ├── drawable-ldpi
│       │   └── icon.png
│       └── drawable-mdpi
│           └── icon.png
├── deploy
│   └── android
│       └── build
├── project.go
└── qml
    └── project.qml
```

The content of the `android` or `android-emulator` folder will be copied prior `gradle` is used to create an `*.apk`.

Furthermore you can use the `AndroidManifest.xml` found inside `deploy/android/build` as a template to build upon.

This can be useful if you want to change the package identifier, the icon or add more permissions.

You could for example add an `android:icon="@drawable/icon"` attribute to `application` in `AndroidManifest.xml` to set replace the default icon with your own.

If you want to have your `*.apk` signed automatically, then you need to create an `alias.txt`, `password.txt` and an `project_name.keystore` inside the `android` folder.

Or you could sign it manually as well.

### Java Native Interface (JNI)

You can call Java functions directly from Go with the help of the [androidextras](https://github.com/therecipe/qt/blob/master/androidextras) package.

Callbacks into Go code from Java are also possible (with some manual labor).

Take a look at the [examples](https://github.com/therecipe/qt/blob/master/internal/examples/androidextras) to see how it works.

### Android Wear

To get your app working on Android Wear, you just need to add `android:theme="@android:style/Theme.DeviceDefault"` to the `application` node of your custom `AndroidManifest.xml` before running qtdeploy.

You can find a working example [here](https://github.com/therecipe/qt/blob/master/internal/examples/canvas3d/quickitemtexture_wear).

</details>

<details markdown="1"> <summary markdown="0">iOS</summary>

### Branding

You can use the `ios` or `ios-simulator` folder at the root of your project to let `qtdeploy` automatically bundle your assets.

Furthermore you can sign your app with the help of Xcode.

</details>

<details markdown="1"> <summary markdown="0">SailfishOS</summary>

### Environmental Variables

| Variable       | Default                           | Type   | Note                                                                                |
|:--------------:|:---------------------------------:|:------:|:-----------------------------------------------------------------------------------:|
| VIRTUALBOX_DIR |     "depends on the host os"      | string |          set this to override the default path to the `vboxmanage` binary           |
|  SAILFISH_DIR  | $HOME/SailfishOS or C:\SailfishOS | string | set this to override the default path to the root dir of your sailfish installation |

### Branding

You can use the `sailfish` or `sailfish-emulator` folder at the root of your project to let `qtdeploy` automatically bundle your assets.

```
project_name
├── project.go
├── qml
│   └── project.qml
└── sailfish
    ├── harbour-project_name.png
    ├── project_name.desktop
    └── project_name.spec
```

Furthermore you can use the `*.desktop` and `*.spec` files found inside `deploy/sailfish` as a template to build upon.

</details>

<details markdown="1"> <summary markdown="0">AsteroidOS</summary>

-	**TBA**

</details>

#### Embedded

<details markdown="1"> <summary markdown="0">Raspberry Pi</summary>

### Environmental Variables

| Variable      | Default                       | Type   | Note                                                                        |
|:-------------:|:-----------------------------:|:------:|:---------------------------------------------------------------------------:|
| RPI_TOOLS_DIR |       $HOME/raspi/tools       | string |        set this to override the default path to the raspi tools repo        |
| RPI_COMPILER  | arm-rpi-4.9.3-linux-gnueabihf | string | `gcc-linaro-arm-linux-gnueabihf-raspbian-x64` is used for the docker images |

### Branding

You can use the `rpiX` folders at the root of your project to let `qtdeploy` automatically bundle your assets.

</details>

Tools
-----

The tools can be found either in `$GOPATH/bin` or `$GOBIN` and are symlinked into your **PATH** during the setup.

If you ran the minimal setup and/or depend on the docker images for the specific target, then you need to use the `-docker` flag when you use them.

<details markdown="1"> <summary markdown="0">qtsetup</summary>

The use of `qtsetup` is rarely necessary after you are done with the inital setup.

But you can use it to dump some environmental infos by running: `qtsetup check [TARGET]`

And it can be used to update the tools by running: `qtsetup update`

Or be used to upgrade the whole repo by running: `qtsetup upgrade` (you will need to re-run the setup after this)

Run `qtsetup -help` for more infos.

</details>

<details markdown="1"> <summary markdown="0">qtdeploy</summary>

`qtdeploy` is used to deploy your applications.

You can compile and deploy to your desktop target by running: `qtdeploy build desktop [path/to/your/project]`

(it's possible to obtain the path, if you are currently inside the project folder)

The other valid targets are:

| Target            |
|:-----------------:|
|      windows      |
|      darwin       |
|       linux       |
|      android      |
| android-emulator  |
|        ios        |
|   ios-simulator   |
|     sailfish      |
| sailfish-emulator |
|       rpi1        |
|       rpi2        |
|       rpi3        |
|     asteroid      |

So for example running:

`qtdeploy build android [path/to/your/project]`

will deploy your project for the android target.

and running:

`qtdeploy -docker build android [path/to/your/project]`

will use the docker image for it.

You can speedup the compilation by using the `-fast` flag. (only for the desktop target)

`qtdeploy` will also call `qtrcc`, `qtmoc` and `qtminimal` for you.

There will be various files created during the compilation, it's recommended to not manually edit or remove them.

Run `qtdeploy -help` for more infos.

</details>

<details markdown="1"> <summary markdown="0">qtrcc</summary>

`qtrcc` is a wrapper around Qt's [rcc](https://doc.qt.io/qt-5/resources.html) and can be used to bundle resources with your code.

You can use it to bundle resources for the desktop target like this: `qtrcc desktop path/to/your/project`

(it's possible to obtain the path, if you are currently inside the project folder)

`qtrcc` will then automatically bundle all resources from the `qml` sub-folder (`path/to/your/project/qml`).

If you want to re-use your current `.qrc` file(s) instead, or simply want to have more control about your resouces, then you can place your `.qrc` file(s) inside the root folder.

Or if you need even more control, then you can just use `rcc` directly as well, but make sure to obtain one of these `rcc_cgo_*.go` files to place them alongside with your `rcc` generated `.cpp` file.

Invoking `qtrcc` will generate at least 3 different files, it's recommended to not manually edit or remove them.

Futhermore `qtquickcontrols2.conf` files are automatically detected inside the root folder and bundled as well.

You can edit the `*.qrc` files by hand, or with the help of the `Qt Creator`.

Run `qtrcc -help` for more infos.

</details>

<details markdown="1"> <summary markdown="0">qtmoc</summary>

`qtmoc` is a wrapper around Qt's [moc](https://doc.qt.io/qt-5/moc.html) and can be used to create sub-classes of Qt classes and to extend them with your own constructors, signals, slots and properties.

It's recommended that you get yourself familiar with Qt's signal and slot mechanism first: https://doc.qt.io/qt-5/signalsandslots.html

You can use `qtmoc` like this: `qtmoc desktop path/to/your/project` or `qtmoc desktop path/to/your/sub-package`

(it's possible to obtain the path, if you are currently inside the project or sub-package folder)

For this to work you need to create a moc struct first:

```go
package main

import "github.com/therecipe/qt/core"

//this struct will let qtmoc generate the necessary Go and C++ code.
//such as the NewExampleStruct and DestroyExampleStruct functions
type exampleStruct struct {

	//let qtmoc know what class you want to sub-class
	//this can be any class that is derived from QObject
	//and could also be one of your own moc classes
	//it's important that you anonymously embed the type as a value here
	//qtmoc will ignore this struct otherwise
	core.QObject

	//this will let qtmoc know that you want to have the "init" function called
	//when you call NewExampleStruct
	_ func() `constructor:"init"`

	//this will let qtmoc know that you want a signal called "firstSignal"
	//so that the related Connect* Disconnect* functions can be created
	_ func() `signal:"firstSignal"`

	//the signal can also have parameters, but no return parameter
	//return parameters are only possible if you create a slot
	_ func(bool, int, string, []string, map[string]string) `signal:"secondSignal"`

	//you can also let the signal accept a single *core.QObject an array or a map
	_ func(*core.QObject, []*core.QObject, map[string]*core.QObject) `signal:"thirdSignal"`

	//as a special addition to the supported Go primitives, you can also use Go errors in signals/slots/properties
	_ func(error) `signals:"fourthSignal"`

	//this will let qtmoc know that you want a slot called "firstSlot"
	_ func() `slot:"firstSlot"`

	//a slot can be created in the same way as a signal, but it can additionally return a single argument
	_ func(string) string               `slot:"secondSlot"`
	_ func(*core.QObject) *core.QObject `slot:"thirdSlot"`
	_ func() error                      `slot:"fourthSlot"`

	//this will let qtmoc know that you want a property called "firstProperty"
	//there will be helper getter + setter functions and a changed signal created called:
	//FirstProperty (IsFirstProperty for bools), SetFirstProperty, FirstProperyChanged
	_ string `property:"firstProperty"`
}

//this function will be automatically called, when you use the `NewExampleStruct` function
func (s *exampleStruct) init() {
	//here you can do some initializing
	s.SetFirstProperty("defaultString")
	s.ConnectFirstSignal(func() { println("do something here") })
	s.ConnectSecondSignal(s.secondSignal)
}

func (s *exampleStruct) secondSignal(v0 bool, v1 int, v2 string, v3 []string, v4 map[string]string) {
	println("do something here")
}
```

There are some additions planned, such as:

-	custom constructors (with additional input and return parameters)
-	the option to automatically connect signals/slots
-	the option to make signals/slots/properties private
-	connect functions for properties

Run `qtmoc -help` for more infos.

</details>

<details markdown="1"> <summary markdown="0">qtminimal</summary>

`qtminimal` is not needed to be ran manually, as it is just used to reduce the overall compiled binary size, by analysing your code and re-generating a sub-set of the binding.

You can use it to generate the sub-set for the desktop target like this: `qtminimal desktop path/to/your/project`

(it's possible to obtain the path, if you are currently inside the project folder)

Then you could use `go build -tags=minimal ...` to compile against the slimmed down version of the binding.

**Note that qtminimal should be always ran after the optional use of qtrcc or qtmoc.**

Run `qtminimal -help` for more infos.

</details>

#### Speed up your development cycle

It's recommended to use `qtdeploy -fast build ...` if you want to have your project compiled as fast as possible.

This will skip most of the actions `qtdeploy` would take otherwise, such as running `qtmoc`, `qtminimal` and deploying your project.

Note that this option is only available for the `desktop` target, and it's mandatory to run the normal deployment at least once before using the `-fast` flag.

Beside the option to use the `-fast` flag, you could also just use `qtrcc`, `qtmoc` and `qtminimal` manually with `go build`.

Hello World
-----------

-	Run the [setup](#installation)

-	Choose one of the following examples and manually setup the `$GOPATH/src/qtexample` folder

#### Basic

<details markdown="1"> <summary markdown="0">widgets</summary>

`$GOPATH/src/qtexample/main.go`

```go
package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {

	widgets.NewQApplication(len(os.Args), os.Args)

	//create a window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize2(200, 200)

	//create a layout
	layout := widgets.NewQVBoxLayout()

	//create a widget and set the layout
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	//create a lineedit and add it to the layout
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("1. write something")
	layout.AddWidget(input, 0, 0)

	//create a button and add it to the layout
	button := widgets.NewQPushButton2("2. click me", nil)
	button.ConnectClicked(func(checked bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	layout.AddWidget(button, 0, 0)

	//add the widget as the central widget to the window
	window.SetCentralWidget(widget)

	//show the window
	window.Show()

	//enter the main event loop
	widgets.QApplication_Exec()
}
```

</details>

<details markdown="1"> <summary markdown="0">qml/quick</summary>

`$GOPATH/src/qtexample/main.go`

```go
package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

func main() {

	//enable high dpi scaling
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	//use the material style for qml controls
	//"universal" is also available
	quickcontrols2.QQuickStyle_SetStyle("material")

	//create a qml application
	view := qml.NewQQmlApplicationEngine(nil)

	//load the main qml file
	view.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	//enter the main event loop
	gui.QGuiApplication_Exec()
}
```

`$GOPATH/src/qtexample/qml/main.qml`

```qml
import QtQuick 2.7
import QtQuick.Controls 2.1

ApplicationWindow {
  id: window

  visible: true
  title: "Hello World Example"
  minimumWidth: 400
  minimumHeight: 400

  Column {
    anchors.centerIn: parent

    TextField {
      id: input

      anchors.horizontalCenter: parent.horizontalCenter

      placeholderText: "1. write something"
    }

    Button {
      anchors.horizontalCenter: parent.horizontalCenter

      text: "2. click me"
      onClicked: dialog.open()
    }
  }

  Dialog {
    contentWidth: window.width / 2
    contentHeight: window.height / 4

    id: dialog
    standardButtons: Dialog.Ok

    contentItem: Label {
      text: input.text
    }
  }
}
```

</details>

#### Advanced

<details markdown="1"> <summary markdown="0">widgets</summary>

`$GOPATH/src/qtexample/main.go`

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//go:generate qtmoc
type MocLabel struct {
	widgets.QLabel

	_ func(int) `signal:"updateLabel"`
}

func main() {

	widgets.NewQApplication(len(os.Args), os.Args)

	//create a window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize2(200, 200)

	//create a layout
	layout := widgets.NewQVBoxLayout()

	//create a widget and set the layout
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	//create a moc label
	label := NewMocLabel(nil, 0)
	label.SetAlignment(core.Qt__AlignCenter)

	//wrap the setText function with a custom signal
	label.ConnectUpdateLabel(func(s int) {

		//we are back in the main thread
		//so it's safe to update the label now
		label.SetText(fmt.Sprintf("%v second(s)", s))
	})

	//setup a ticker to update the label in the background
	t := time.Now()
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for _ = range ticker.C {
			label.UpdateLabel(int(time.Since(t).Seconds()))
		}
	}()

	//add the label to the layout
	layout.AddWidget(label, 0, 0)

	//create a button and add it to the layout
	button := widgets.NewQPushButton2("reset ticker", nil)
	button.ConnectClicked(func(checked bool) {
		label.UpdateLabel(0)
		t = time.Now()
	})
	layout.AddWidget(button, 0, 0)

	//add the widget as the central widget to the window
	window.SetCentralWidget(widget)

	//show the window
	window.Show()

	//enter the main event loop
	widgets.QApplication_Exec()
}
```

</details>

<details markdown="1"> <summary markdown="0">qml/quick</summary>

`$GOPATH/src/qtexample/main.go`

```go
package main

import (
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

//go:generate qtmoc
type MocBridge struct {
	core.QObject

	_ func(secs int) `signal:"updateLabel"`
	_ func()         `signal:"resetTicker"`
}

func main() {

	//enable high dpi scaling
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	//use the material style for qml controls
	//"universal" is also available
	quickcontrols2.QQuickStyle_SetStyle("material")

	//create a qml application
	app := qml.NewQQmlApplicationEngine(nil)

	//create a moc bridge
	bridge := NewMocBridge(nil)

	//setup a ticker to update the label in the background
	t := time.Now()
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for _ = range ticker.C {
			//call the connected "onUpdateLabel" signal on the qml side
			bridge.UpdateLabel(int(time.Since(t).Seconds()))
		}
	}()

	//connect the resetTicker function that is later called by qml
	bridge.ConnectResetTicker(func() {
		t = time.Now()
	})

	//make the bridge object accessible inside qml as "go"
	app.RootContext().SetContextProperty("go", bridge)

	//load the main qml file
	app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	//enter the main event loop
	gui.QGuiApplication_Exec()
}
```

`$GOPATH/src/qtexample/qml/main.qml`

```qml
import QtQuick 2.7
import QtQuick.Controls 2.1

ApplicationWindow {
  id: window

  visible: true
  title: "Hello World Example"
  minimumWidth: 400
  minimumHeight: 400

  Column {
    anchors.centerIn: parent

    Label {
      id: label

      anchors.horizontalCenter: parent.horizontalCenter

      Connections {
        target: go
        onUpdateLabel: {
          label.text = secs + " second(s)"
        }
      }
    }

    Button {
      anchors.horizontalCenter: parent.horizontalCenter

      text: "reset ticker"
      onClicked: {
        go.updateLabel(0)
        go.resetTicker()
      }
    }
  }
}
```

</details>

-	Run `qtdeploy test desktop $GOPATH/src/qtexample` to compile, deploy and start the application.

-	More examples can be found [here](https://github.com/therecipe/qt/blob/master/internal/examples)

Continuous integration and deployment
-------------------------------------

There are a few possible ways how you can integrate this dependency into your CI or CD plan.

-	You could create your own docker image(s) on top of the pre-built docker images.

-	You could pull the docker images for your target and then automate your work inside the (disposable) container.

-	You could use the minimal setup, which only depends on Docker + Go and then use `qtdeploy -docker build ...`

-	You could look into the [CI](https://github.com/therecipe/qt/blob/master/internal/ci) folder and the `*.yml` files in the root folder of this repo.

Misc
----

[Qt documentation](https://doc.qt.io/qt-5/classes.html)

[#qt-binding](https://gophers.slack.com/messages/qt-binding/details) Slack channel ([invite](https://invite.slack.golangbridge.org)\)

[An Video Overview of GUI technologies in Qt](https://www.youtube.com/watch?v=WIRRoPxIerc)

<details markdown="1"> <summary markdown="0">Debugging and Profiling Qml Applications</summary>

-	Export `QT_DEBUG_QML=true`

-	Run `qtdeploy` (`-fast` or `go build/run` won't work)

-	Manually start the created binary with: `-qmljsdebugger=port:8080,block`

-	Open the created `debug.pro`

-	For debugging: Debug > Start Debugging > Attach to QML Port...

	-	https://doc.qt.io/qt-5/qtquick-debugging.html

	-	https://doc.qt.io/qtcreator/creator-debugging-qml.html

	-	https://doc.qt.io/qtcreator/creator-qml-debugging-example.html

-	For profiling: Analyze > QML Profiler (External)

	-	https://doc.qt.io/qtcreator/creator-qml-performance-monitor.html

	-	https://doc.qt.io/qt-5/qtquick-visualcanvas-scenegraph-renderer.html#visualizing-overdraw

</details>

<details markdown="1"> <summary markdown="0">Drag & Drop WYSIWYG Editors</summary>

-	For the desktop widgets: https://doc.qt.io/qtcreator/creator-using-qt-designer.html

[Here](https://github.com/therecipe/qt/blob/master/internal/examples/uitools) is an example how to load the created `*.ui` files.

-	For the qml/quick widgets: https://doc.qt.io/qtcreator/creator-using-qt-quick-designer.html

The created `*.qml` files can later be used like hand written qml files.

</details>

<details markdown="1"> <summary markdown="0">Code Editor Settings (for code completion)</summary>

On windows, depending on your editor or IDE you may also need to export `GOARCH=386` to make the code completion work.

-	IntelliJ Idea / Gogland

	Help > Edit Custom Properties

	```
	# custom IntelliJ IDEA properties

	idea.max.intellisense.filesize=999999
	```

	Help > Edit Custom VM Options

	```
	# custom IntelliJ IDEA VM options

	-Xms128m
	-Xmx4000m
	-XX:ReservedCodeCacheSize=512m
	-XX:+UseCompressedOops
	```

-	Neovim

	Update your neocomplete.vim

	```
	let g:neocomplete#skip_auto_completion_time = ""
	```

</details>

<details markdown="1"> <summary markdown="0">3rd party examples / demos / applications</summary>

-	[5k3105/AM](https://github.com/5k3105/AM)
-	[5k3105/Pipeline-Editor](https://github.com/5k3105/Pipeline-Editor)
-	[5k3105/Sprite-Editor](https://github.com/5k3105/Sprite-Editor)
-	[gen2brain/url2img](https://github.com/gen2brain/url2img)
-	[joeblew99/qt-archi](https://github.com/joeblew99/qt-archi)
-	[jordanorelli/grpc-ui](https://github.com/jordanorelli/grpc-ui)
-	[Matt-Texier/local-mitigation-agent](https://github.com/Matt-Texier/local-mitigation-agent)

</details>

License
-------

This binding is available under LGPLv3.

Qt is available under multiple licenses: https://www.qt.io/licensing
