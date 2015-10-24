#Qt binding for Go (Windows / Mac OS X / Linux / Android)

##Windows (x86)
(currently not working) -> https://support.microsoft.com/kb/830473

1. Install Go >= 1.5.1
> https://storage.googleapis.com/golang/go1.5.1.windows-386.msi

2. Install Qt 5.5.1 ("C:\Qt\Qt5.5.1")
> https://download.qt.io/official_releases/qt/5.5/5.5.1/qt-opensource-windows-x86-android-5.5.1.exe

3. Add the directory that contains "gcc" to your PATH
> C:\Qt\Qt5.5.1\Tools\mingw492_32\bin

4. Download the qt binding
> go get github.com/therecipe/qt

5. Generate, install and test
> github.com/therecipe/qt/internal/binding/make.sh

##Mac OS X (x64)

1. Install Go >= 1.5.1
> https://storage.googleapis.com/golang/go1.5.1.darwin-amd64.pkg

2. Install Qt 5.5.1 ("/usr/local/Qt5.5.1")
> https://download.qt.io/official_releases/qt/5.5/5.5.1/qt-opensource-mac-x64-android-5.5.1.dmg

3. Install Xcode >= 7.0.1

4. Download the qt binding
> go get github.com/therecipe/qt

5. Generate, install and test
> github.com/therecipe/qt/internal/binding/make.sh

##Linux (x86/x64)

1. Install Go >= 1.5.1

  >(x86) https://storage.googleapis.com/golang/go1.5.1.linux-386.tar.gz

  >(x64) https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz

2. Install Qt 5.5.1 ("/usr/local/Qt5.5.1")

  >(x86) https://download.qt.io/official_releases/qt/5.5/5.5.1/qt-opensource-linux-x86-android-5.5.1.run

  >(x64) https://download.qt.io/official_releases/qt/5.5/5.5.1/qt-opensource-linux-x64-android-5.5.1.run

3. Install "g++"
> sudo apt-get install g++

4. Install OpenGL dependencies
> sudo apt-get install mesa-common-dev

5. Download the qt binding
> go get github.com/therecipe/qt

6. Generate, install and test
> github.com/therecipe/qt/internal/binding/make.bat

##Android (armv7)

TBA
