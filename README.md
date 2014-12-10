#Qt binding for Go

##Windows (x86):

1. Install Go >= 1.4 (x86)
> https://golang.org/dl/

2. Install Qt 5.4.0
> https://download.qt-project.org/official_releases/qt/5.4/5.4.0/qt-opensource-windows-x86-mingw491_opengl-5.4.0.exe

3. Add the directory that contains "gcc" to your PATH
> C:\Qt\Qt5.4.0\Tools\mingw491_32\bin

4. Download the qt binding
> go get -d github.com/therecipe/qt

5. Setup your build environment (edit the file if necessary)
> cgo_windows_386.go

6. Install the qt binding (failed ? -> go back to 5)
> go install github.com/therecipe/qt

7. Test the example (edit the file if necessary)
> example/windows/deploy_windows.bat

##Mac OS X (x64):

1. Install Go >= 1.4
> https://golang.org/dl/

2. Install Qt 5.4.0
> https://download.qt-project.org/official_releases/qt/5.4/5.4.0/qt-opensource-mac-x64-clang-5.4.0.dmg

3. Install Xcode >= 6.1

4. Download the qt binding
> go get -d github.com/therecipe/qt

5. Setup your build environment (edit the file if necessary)
> cgo_darwin_amd64.go

6. Install the qt binding (failed ? -> go back to 5)
> go install github.com/therecipe/qt

7. Test the example (edit the file if necessary)
> example/mac/deploy_mac.sh

##Linux (x86/x64):

1. Install Go >= 1.4 (x86/x64)
> https://golang.org/dl/

2. Install Qt 5.4.0 

  >(x86) https://download.qt-project.org/official_releases/qt/5.4/5.4.0/qt-opensource-linux-x86-5.4.0.run
  
  >(x64) https://download.qt-project.org/official_releases/qt/5.4/5.4.0/qt-opensource-linux-x64-5.4.0.run
  
3. Install "g++"
> sudo apt-get install build-essential

4. Install OpenGL libraries
> sudo apt-get install mesa-common-dev

5. Download the qt binding
> go get -d github.com/therecipe/qt

6. Setup your build environment (edit the file if necessary)
> cgo_linux_*.go

7. Install the qt binding (failed ? -> go back to 6)
> go install github.com/therecipe/qt

8. Test the example (edit the file if necessary)
> example/linux/deploy_linux.sh