
::enable delayed expansion
setlocal enabledelayedexpansion


::disable updates
reg add "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\WindowsUpdate\Auto Update" /v AUOptions /t REG_DWORD /d 1 /f
sc config wuauserv start= disabled


::install curl
expand -f:* c:\tmp\curl.cab c:\tmp\
mv -f /cygdrive/c/tmp/AMD64/* "/cygdrive/c/Program Files/OpenSSH/bin/"


::install 7z
set SZ=7z1604-x64.exe
curl -sL --retry 10 --retry-delay 10 -o %TMP%\%SZ% http://7-zip.org/a/%SZ%
%TMP%\%SZ% /S
del %TMP%\%SZ% /Q
setx /M PATH "%PATH%;C:\Progra~1\7-Zip"
set PATH=%PATH%;C:\Progra~1\7-Zip


::install Git
set GIT=Git-2.13.2-64-bit.exe
curl -sL --retry 10 --retry-delay 10 -o %TMP%\%GIT% https://github.com/git-for-windows/git/releases/download/v2.13.2.windows.1/%GIT%
%TMP%\%GIT% /silent /norestart
del %TMP%\%GIT% /Q
setx /M PATH "%PATH%;C:\Progra~1\Git\bin"
set PATH=%PATH%;C:\Progra~1\Git\bin


::install Go + pull repo
set GO=go1.8.3.windows-amd64.msi
curl -sL --retry 10 --retry-delay 10 -o %TMP%\%GO% http://storage.googleapis.com/golang/%GO%
%TMP%\%GO% /passive /norestart
del %TMP%\%GO% /Q
setx /M PATH "%PATH%;C:\Go\bin"
set PATH=%PATH%;C:\Go\bin
setx /M GOPATH "C:\gopath"
set GOPATH=C:\gopath

go get -v github.com/therecipe/qt/cmd/...


::install VC++ 2015 Redis
set VC=vc_redist.x64.exe
curl -sL --retry 10 --retry-delay 10 -o %TMP%\%VC% https://download.microsoft.com/download/9/3/F/93FCF1E7-E6A4-478B-96E7-D4B285925B00/%VC%
%TMP%\%VC% /passive /norestart
del %TMP%\%VC% /Q


if "%QT_MSYS2%" == "true" (
  setx /M QT_MSYS2 "%QT_MSYS2%"
  setx /M QT_MSYS2_STATIC "%QT_MSYS2_STATIC%"
  setx /M QT_MSYS2_ARCH "%QT_MSYS2_ARCH%"

  if "%QT_MSYS2_ARCH%" == "386" (
    setx /M MSYSTEM "MINGW32"
    echo MSYSTEM=MINGW32>> C:\Users\vagrant\.ssh\environment
  ) else (
    setx /M MSYSTEM "MINGW64"
    echo MSYSTEM=MINGW64>> C:\Users\vagrant\.ssh\environment
  )

  echo QT_MSYS2=true>> C:\Users\vagrant\.ssh\environment
  echo QT_MSYS2_STATIC=%QT_MSYS2_STATIC%>> C:\Users\vagrant\.ssh\environment
  echo QT_MSYS2_ARCH=%QT_MSYS2_ARCH%>> C:\Users\vagrant\.ssh\environment


  ::install msys2
  set MSYS2=msys2-x86_64-20161025.exe
  set AI=auto-install.js
  curl -sL --retry 10 --retry-delay 10 -o %TMP%\!MSYS2! http://repo.msys2.org/distrib/x86_64/!MSYS2!
  curl -sL --retry 10 --retry-delay 10 -o %TMP%\!AI! https://raw.githubusercontent.com/msys2/msys2-installer/master/!AI!
  %TMP%\!MSYS2! --script %TMP%\!AI!
  del %TMP%\!MSYS2! /Q
  del %TMP%\!AI! /Q


  C:\msys64\usr\bin\bash -l -c "pacman -Q"
  C:\msys64\usr\bin\bash -l -c "pacman -Syyu --noconfirm --noprogress"
  C:\msys64\usr\bin\bash -l -c "pacman -Syyu --noconfirm --noprogress"

  if "%QT_MSYS2_ARCH%" == "386" (
    if "%QT_MSYS2_STATIC%" == "true" (
      C:\msys64\usr\bin\bash -l -c "pacman -S --noconfirm --needed --noprogress mingw-w64-i686-qt-creator mingw-w64-i686-qt5-static"
    ) else (
      C:\msys64\usr\bin\bash -l -c "pacman -S --noconfirm --needed --noprogress mingw-w64-i686-qt-creator mingw-w64-i686-qt5"
      C:\msys64\usr\bin\bash -l -c "pacman -S --noconfirm --needed --noprogress mingw-w64-i686-qtwebkit"
    )
  ) else (
    if "%QT_MSYS2_STATIC%" == "true" (
      C:\msys64\usr\bin\bash -l -c "pacman -S --noconfirm --needed --noprogress mingw-w64-x86_64-qt-creator mingw-w64-x86_64-qt5-static"
    ) else (
      C:\msys64\usr\bin\bash -l -c "pacman -S --noconfirm --needed --noprogress mingw-w64-x86_64-qt-creator mingw-w64-x86_64-qt5"
      C:\msys64\usr\bin\bash -l -c "pacman -S --noconfirm --needed --noprogress mingw-w64-x86-qtwebkit"
    )
  )

  C:\msys64\usr\bin\bash -l -c "pacman -Q"
  C:\msys64\usr\bin\bash -l -c "pacman -Scc --noconfirm --noprogress"
) else (
  ::install Qt
  set QT=qt-opensource-windows-x86-android-5.8.0.exe
  curl -sL --retry 10 --retry-delay 10 -o %TMP%\!QT! https://download.qt.io/official_releases/qt/5.8/5.8.0/!QT!
  %TMP%\!QT! --script %GOPATH%\src\github.com\therecipe\qt\internal\ci\iscript.qs
  del %TMP%\!QT! /Q
  setx /M PATH "%PATH%;C:\Qt\Qt5.8.0\Tools\mingw530_32\bin"
  set PATH=%PATH%;C:\Qt\Qt5.8.0\Tools\mingw530_32\bin
)


::update ssh env variables
echo TMP=C:/tmp>> C:\Users\vagrant\.ssh\environment
net stop "OpenSSH Server"
net start "OpenSSH Server"


if "%ANDROID%" == "true" (
  ::install JDK
  set JDK=jdk-8u131-windows-x64.exe
  curl -sL --retry 10 --retry-delay 10 -o %TMP%\!JDK! -H "Cookie: oraclelicense=accept-securebackup-cookie" http://download.oracle.com/otn-pub/java/jdk/8u131-b11/d54c1d3a095b4ff2b6607d096fa80163/!JDK!
  %TMP%\!JDK! /s
  del %TMP%\!JDK! /Q
  setx /M JAVA_HOME "C:\Progra~1\Java\jdk1.8.0_131"
  set JAVA_HOME=C:\Progra~1\Java\jdk1.8.0_131


  ::install Android SDK
  set SDK=tools_r25.2.5-windows.zip
  curl -sL --retry 10 --retry-delay 10 -o %TMP%\!SDK! https://dl.google.com/android/repository/!SDK!
  7z x %TMP%\!SDK! -oC:\android-sdk-windows\
  del %TMP%\!SDK! /Q

  cmd /C "C:\android-sdk-windows\tools\android.bat list sdk"
  cmd /C "echo y | C:\android-sdk-windows\tools\android.bat -s update sdk -f -u -t 1,2,3,4,5,6"
  cmd /C "echo y | C:\android-sdk-windows\tools\android.bat -s update sdk -a -f -u -t 5"


  ::install Android NDK
  set NDK=android-ndk-r14b-windows-x86_64.zip
  curl -sL --retry 10 --retry-delay 10 -o %TMP%\!NDK! https://dl.google.com/android/repository/!NDK!
  7z x %TMP%\!NDK! -oC:\
  del %TMP%\!NDK! /Q
)


::qtsetup
if "%QT_MSYS2%" == "true" (
  if "%QT_MSYS2_ARCH%" == "386" (
    set MSYSTEM=MINGW32
    C:\msys64\usr\bin\bash -l -c "$GOPATH/bin/qtsetup full desktop"
  ) else (
    set MSYSTEM=MINGW64
    C:\msys64\usr\bin\bash -l -c "$GOPATH/bin/qtsetup full desktop"
  )
) else (
  if "%DESKTOP%" == "true" "%GOPATH%\bin\qtsetup" full desktop
  if "%ANDROID%" == "true" "%GOPATH%\bin\qtsetup" full android
)
