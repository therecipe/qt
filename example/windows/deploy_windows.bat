go build -ldflags -H=windowsgui -o example.exe ../

set PATH=C:\Qt\Qt5.4.0\5.4\mingw491_32\bin;%PATH%

windeployqt example.exe

example.exe