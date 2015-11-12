# Qt binding for Go (Windows / Mac OS X / Linux / Android)

## Desktop (Windows / Mac OS X / Linux)

1. Install Go >= 1.5.1
	* https://storage.googleapis.com/golang/go1.5.1.windows-386.msi
	* https://storage.googleapis.com/golang/go1.5.1.darwin-amd64.pkg
	* https://storage.googleapis.com/golang/go1.5.1.linux-386.tar.gz
	* https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz

2. Install Qt 5.5.1 `C:\Qt\Qt5.5.1\` or `/usr/local/Qt5.5.1/`
	* https://download.qt.io/official_releases/qt/5.5/5.5.1/qt-opensource-windows-x86-android-5.5.1.exe
	* https://download.qt.io/official_releases/qt/5.5/5.5.1/qt-opensource-mac-x64-android-5.5.1.dmg
	* https://download.qt.io/official_releases/qt/5.5/5.5.1/qt-opensource-linux-x86-android-5.5.1.run
	* https://download.qt.io/official_releases/qt/5.5/5.5.1/qt-opensource-linux-x64-android-5.5.1.run

3. Setup the environment
	* Windows
		* Add the directory that contains "gcc" to your PATH

			`C:\Qt\Qt5.5.1\Tools\mingw492_32\bin\`

	* Mac OS X
		* Install Xcode >= 7.0.1

	* Linux
		* Install "g++"

			`sudo apt-get install g++`

		* Install OpenGL dependencies

			`sudo apt-get install mesa-common-dev`

4. Download the qt binding

 	`go get github.com/therecipe/qt`


5. Generate, install and test

  	`github.com/therecipe/qt/setup.bat`

  	or

  	`github.com/therecipe/qt/setup.sh`

## Mobile (Android)

1. Make sure the binding is working by setting up the Desktop version

2. Install the Android SDK `C:\android\android-sdk\` or `/opt/android-sdk/`
	* http://dl.google.com/android/installer_r24.4.1-windows.exe
	* http://dl.google.com/android/android-sdk_r24.4.1-macosx.zip
	* http://dl.google.com/android/android-sdk_r24.4.1-linux.tgz

3. Install the SDK dependencies `C:\android\android-sdk\tools\android.bat` or `/opt/android-sdk/tools/android`
	* Tools
		* Android SDK Build-tools (23.0.2)
	* Android 6.0 (API 23)
		* SDK Platform
	* Extras (Windows only)
		* Google USB Driver

4. Install the Android NDK `C:\android\android-ndk\` or `/opt/android-ndk/`
	* http://dl.google.com/android/ndk/android-ndk-r10e-windows-x86.exe
	* http://dl.google.com/android/ndk/android-ndk-r10e-darwin-x86_64.bin
	* http://dl.google.com/android/ndk/android-ndk-r10e-linux-x86.bin
	* http://dl.google.com/android/ndk/android-ndk-r10e-linux-x86_64.bin

5. Install Apache-Ant `C:\android\apache-ant\` or `/opt/apache-ant/`
	* http://mirror.synyx.de/apache/ant/binaries/apache-ant-1.9.6-bin.zip
	* http://mirror.synyx.de/apache/ant/binaries/apache-ant-1.9.6-bin.tar.bz2
	* http://mirror.synyx.de/apache/ant/binaries/apache-ant-1.9.6-bin.tar.gz

6. Install Java SE Development Kit (linux: `/opt/jdk/`)
	* http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html

7. Install and test

  	`github.com/therecipe/qt/setup.bat android`

    or

  	`github.com/therecipe/qt/setup.sh android`
