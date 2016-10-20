## Screenshots

### Windows
![](internal/screens/windows.png)

### macOS
![](internal/screens/mac.png)

### Linux
![](internal/screens/linux.png)

### Android

#### Portrait
![](internal/screens/android_portrait.png)

#### Landscape
![](internal/screens/android_landscape.png)

### iOS

#### Portrait
![](internal/screens/ios_portrait.png)

#### Landscape
![](internal/screens/ios_landscape.png)

### Sailfish OS

![](internal/screens/sailfish_portrait.png)

---
---

## Installation

**You will need at least 8gb ram to install this binding!**

* Desktop
	* [Windows](#windows-1)
	* [macOS](#macos-1)
	* [Linux](#linux-1)

* Mobile
	* [Android](#android-1)
	* [iOS](#ios-1)
	* [Sailfish OS](#sailfish-os-1)

* Embedded
	* [Raspberry Pi](#raspberry-pi)

---

## Windows

1. Install Go >= 1.7.1 and setup a proper [**GOPATH**](https://golang.org/doc/code.html#GOPATH)

	* https://golang.org/doc/install?download=go1.7.1.windows-amd64.msi

2. Install Qt 5.7.0; you can also define a custom location with **QT_DIR**

	* https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-windows-x86-android-5.7.0.exe

3. Add the directory that contains **gcc** and **g++** to your **PATH**

	* `C:\Qt\Qt5.7.0\Tools\mingw530_32\bin`

4. Download the binding

	* `go get -d github.com/therecipe/qt`

5. Generate, install and test (20 min)

	* `cd %GOPATH%\src\github.com\therecipe\qt && setup.bat`

6. Create your first [application](#example)

## macOS

1. Install Go >= 1.7.1 and setup a proper [**GOPATH**](https://golang.org/doc/code.html#GOPATH)

	* https://golang.org/doc/install?download=go1.7.1.darwin-amd64.pkg

2. Install Qt 5.7.0; you can also define a custom location with **QT_DIR**

	* Install the official prebuild package; you can also define a custom location with **QT_DIR**
		* without iOS https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-mac-x64-android-5.7.0.dmg
		* with iOS https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-mac-x64-android-ios-5.7.0.dmg

	or

	* Install the Qt-dev package with Homebrew **(experimental)** and define the custom location with **QT_DIR** (usually: /usr/local/opt/qt5/)
		* `brew install qt5`

3. Install **Xcode** >= 7.2.0; you can also define a custom location with **XCODE_DIR**

	* https://itunes.apple.com/us/app/xcode/id497799835

4. Download the binding

	* `go get -d github.com/therecipe/qt`

5. Generate, install and test (20 min)

	* `cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh`

6. Create your first [application](#example)

## Linux

1. Install Go >= 1.7.1 and setup a proper [**GOPATH**](https://golang.org/doc/code.html#GOPATH)

	* https://golang.org/doc/install?download=go1.7.1.linux-amd64.tar.gz

2. Install Qt 5.7.0

	* Install the official prebuild package; you can also define a custom location with **QT_DIR**
		* https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-linux-x64-android-5.7.0.run

	or

	* Install the Qt-dev package with your system package manager; if you want to link against your system libs **(experimental)**
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

	* `go get -d github.com/therecipe/qt`

5. Generate, install and test (20 min)

	* `cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh`

6. Create your first [application](#example)

---

## Android

1. Install the desktop version for [Windows](#windows-1), [macOS](#macos-1) or [Linux](#linux-1)

2. Unzip the Android SDK in `C:\` or `$HOME`; you can also define a custom location with **ANDROID_SDK_DIR**
	* https://dl.google.com/android/android-sdk_r24.4.1-windows.zip
	* https://dl.google.com/android/android-sdk_r24.4.1-macosx.zip
	* https://dl.google.com/android/android-sdk_r24.4.1-linux.tgz

3. Install the SDK dependencies with `C:\android-sdk-windows\tools\android.bat` or `$HOME/android-sdk-{ macosx | linux }/tools/android`
	* Tools
		* Android SDK Build-tools (24.0.3)
	* Android 7.0 (API 24)
		* SDK Platform
	* Extras (Windows only)
		* Google USB Driver

4. Unzip the Android NDK in `C:\` or `$HOME`; you can also define a custom location with **ANDROID_NDK_DIR**
	* https://dl.google.com/android/repository/android-ndk-r12b-windows-x86_64.zip
	* https://dl.google.com/android/repository/android-ndk-r12b-darwin-x86_64.zip
	* https://dl.google.com/android/repository/android-ndk-r12b-linux-x86_64.zip

5. Install Java SE Development Kit >= 8 (Linux: install in `$HOME/jdk/`); you can also define a custom location with **JDK_DIR**
	* https://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html

6. Install and test (20 min)

	* `cd %GOPATH%\src\github.com\therecipe\qt && setup.bat android`

		or

	* `cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh android`

7. Create your first [application](#example)

## iOS

1. Install the desktop version for [macOS](#macos-1)

2. Install and test (20 min)

	* `cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh ios && ./setup.sh ios-simulator`

3. Create your first [application](#example)

## Sailfish OS

1. Install the desktop version for [Windows](#windows-1), [macOS](#macos-1) or [Linux](#linux-1)

2. Install VirtualBox; you can also define a custom location with **VIRTUALBOX_DIR**
	* http://download.virtualbox.org/virtualbox/5.1.8/VirtualBox-5.1.8-111374-Win.exe
	* http://download.virtualbox.org/virtualbox/5.1.8/VirtualBox-5.1.8-111374-OSX.dmg
	* http://download.virtualbox.org/virtualbox/5.1.8/VirtualBox-5.1.8-111374-Linux_amd64.run

3. Install the Sailfish OS SDK; you can also define a custom location with **SAILFISH_DIR**
	* https://releases.sailfishos.org/sdk/installers/1608/SailfishOSSDK-Beta-1608-Qt5-windows-offline.exe
	* https://releases.sailfishos.org/sdk/installers/1608/SailfishOSSDK-Beta-1608-Qt5-mac-offline.dmg
	* https://releases.sailfishos.org/sdk/installers/1608/SailfishOSSDK-Beta-1608-Qt5-linux-64-offline.run

4. Install and test (20 min)

	* `cd %GOPATH%\src\github.com\therecipe\qt && setup.bat sailfish && setup.bat sailfish-emulator`

		or

	* `cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh sailfish && ./setup.sh sailfish-emulator`

5. Create your first [application](#example)

---

## Raspberry Pi

1. Install the desktop version for [Linux](#linux-1)

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
		* `cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh rpi1`

	* Raspberry Pi 2
		* `cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh rpi2`

	* Raspberry Pi 3
		* `cd $GOPATH/src/github.com/therecipe/qt && ./setup.sh rpi3`

15. Notes

	* run `startx &` before starting an application with `-platform xcb` (qml/quick applications won't work; they may work with Qt 5.8)

	* run `weston &` or `weston --tty=1 &` (via ssh) or create your own [compositor](https://doc.qt.io/qt-5/qtwaylandcompositor-index.html) before starting an application with `-platform wayland-egl` (qml/quick applications won't work; they may work with Qt 5.8)

	* you can increase the available gpu memory by editing `/boot/config.txt`

16. Create your first [application](#example)

---
---

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

		//create a button and connect the clicked signal
		var button = widgets.NewQPushButton2("Click me!", nil)
		button.ConnectClicked(func(flag bool) {
			widgets.QMessageBox_Information(nil, "OK", "You clicked me!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		})

		//create a layout and add the button
		var layout = widgets.NewQVBoxLayout()
		layout.AddWidget(button, 0, core.Qt__AlignCenter)

		//create a window, add the layout and show the window
		var window = widgets.NewQMainWindow(nil, 0)
		window.SetWindowTitle("Hello World Example")
		window.SetMinimumSize2(200, 200)
		window.Layout().DestroyQObject()
		window.SetLayout(layout)
		window.Show()

		widgets.QApplication_Exec()
}
```

3. Open the { command-line | terminal | shell } in `[GOPATH]/src/qtExample` and run `qtdeploy build desktop`

4. You will find the application in `[GOPATH]/src/qtExample/deploy/[GOOS]_minimal/qtExample`

5. Take a look at the [examples](https://github.com/therecipe/qt/tree/master/internal/examples)

6. Make yourself familiar with the [Qt documentation](https://doc.qt.io/qt-5/classes.html)

7. Join our Slack Channel [#qt-binding](https://gophers.slack.com/messages/qt-binding/details/)
