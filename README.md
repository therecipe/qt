# Screenshots

## Windows
![](internal/screens/windows.png)

## Mac OS X
![](internal/screens/mac.png)

## Linux
![](internal/screens/linux.png)

## Android

### Portrait
![](internal/screens/android_portrait.png)

### Landscape
![](internal/screens/android_landscape.png)

## iOS

### Portrait
![](internal/screens/ios_portrait.png)

### Landscape
![](internal/screens/ios_landscape.png)

## Sailfish OS

![](internal/screens/sailfish_portrait.png)

[source](http://github.com/therecipe/qt/blob/master/internal/examples/widgets/line_edits/line_edits.go)

---

# Getting Started

## Windows / Mac OS X / Linux

1. Install Go >= 1.7.0 and setup a proper [**GOPATH**](http://golang.org/doc/code.html#GOPATH)
	* http://storage.googleapis.com/golang/go1.7.0.windows-amd64.msi
	* http://storage.googleapis.com/golang/go1.7.0.darwin-amd64.pkg
	* http://storage.googleapis.com/golang/go1.7.0.linux-amd64.tar.gz

2. Install Qt 5.7.0 in `C:\Qt\Qt5.7.0\` or `/usr/local/Qt5.7.0/`
	* http://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-windows-x86-android-5.7.0.exe
	* http://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-mac-x64-android-5.7.0.dmg [(**with iOS**)](http://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-mac-x64-android-ios-5.7.0.dmg)
	* http://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-linux-x64-android-5.7.0.run

3. Setup the environment
	* Windows
		* Add the directory that contains **g++.exe** to your **PATH**

			`C:\Qt\Qt5.7.0\Tools\mingw530_32\bin`

	* Mac OS X
		* Install Xcode >= 7.3.1

	* Linux
		* Install g++

			`sudo apt-get install g++`

		* Install OpenGL dependencies

			`sudo apt-get install mesa-common-dev`

4. Download the binding

 	`go get github.com/therecipe/qt`

5. Generate, install and test

	`cd %GOPATH%\src\github.com\therecipe\qt && setup.bat` **(run as admin)**

	or

	`cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh`

---

## Android

1. Set up the desktop version

2. Install the Android SDK in `C:\android\android-sdk\` or `/opt/android-sdk/`
	* http://dl.google.com/android/android-sdk_r24.4.1-windows.zip
	* http://dl.google.com/android/android-sdk_r24.4.1-macosx.zip
	* http://dl.google.com/android/android-sdk_r24.4.1-linux.tgz

3. Install the SDK dependencies with `C:\android\android-sdk\tools\android.bat` or `/opt/android-sdk/tools/android`
	* Tools
		* Android SDK Build-tools (24.0.1)
	* Android 7.0 (API 24)
		* SDK Platform
	* Extras (Windows only)
		* Google USB Driver

4. Install the Android NDK in `C:\android\android-ndk\` or `/opt/android-ndk/`
	* http://dl.google.com/android/repository/android-ndk-r12b-windows-x86_64.zip
	* http://dl.google.com/android/repository/android-ndk-r12b-darwin-x86_64.zip
	* http://dl.google.com/android/repository/android-ndk-r12b-linux-x86_64.zip

5. Install Java SE Development Kit (Linux: install in `/opt/jdk/`)
	* http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html

6. Install and test

	`cd %GOPATH%\src\github.com\therecipe\qt && setup.bat android` **(run as admin)**

	or

	`cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh android`

---

## iOS

1. Set up the desktop version on Mac OS X

2. Install and test

	`cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh ios && ./setup.sh ios-simulator`

---

## Sailfish OS

1. Set up the desktop version

2. Install VirtualBox
	* http://download.virtualbox.org/virtualbox/5.1.4/VirtualBox-5.1.4-110228-Win.exe
	* http://download.virtualbox.org/virtualbox/5.1.4/VirtualBox-5.1.4-110228-OSX.dmg
	* http://download.virtualbox.org/virtualbox/5.1.4/VirtualBox-5.1.4-110228-Linux_amd64.run

3. Install the Sailfish OS SDK in `C:\SailfishOS\` or `/opt/SailfishOS/`
	* http://releases.sailfishos.org/sdk/installers/1608/SailfishOSSDK-Beta-1608-Qt5-windows-offline.exe
	* http://releases.sailfishos.org/sdk/installers/1608/SailfishOSSDK-Beta-1608-Qt5-mac-offline.dmg
	* http://releases.sailfishos.org/sdk/installers/1608/SailfishOSSDK-Beta-1608-Qt5-linux-64-offline.run

4. Install and test

	`cd %GOPATH%\src\github.com\therecipe\qt && setup.bat sailfish && setup.bat sailfish-emulator` **(run as admin)**

	or

	`cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh sailfish && ./setup.sh sailfish-emulator`

---

# Quick Start

1. Create a folder `[GOPATH]/src/qtExample`

2. Create a file `[GOPATH]/src/qtExample/main.go`
	```go
package main

import (
		"os"

		"github.com/therecipe/qt/widgets"
)

func main() {
		widgets.NewQApplication(len(os.Args), os.Args)

		var btn = widgets.NewQPushButton2("Hello World", nil)
		btn.Resize2(180, 44)
		btn.ConnectClicked(func(flag bool) {
			widgets.QMessageBox_Information(nil, "OK", "You Clicked me!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		})

		var window = widgets.NewQMainWindow(nil, 0)
		window.SetWindowTitle("Hello World Example")
		window.Layout().AddWidget(btn)
		window.Show()

		widgets.QApplication_Exec()
}
```

3. Open the command line in `[GOPATH]/src` and run
`qtdeploy build desktop qtExample`

4. You will find the binary file here
`[GOPATH]/src/qtExample/deploy/[GOOS]/qtExample`

5. Take a look at the [other examples](http://github.com/therecipe/qt/tree/master/internal/examples)

6. Make yourself familiar with the [qt documentation](http://doc.qt.io/qt-5/classes.html)
