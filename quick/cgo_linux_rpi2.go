// +build rpi2

package quick

/*
#cgo CFLAGS: -pipe -march=armv7-a -marm -mthumb-interwork -mfpu=neon-vfpv4 -mtune=cortex-a7 -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -pipe -march=armv7-a -marm -mthumb-interwork -mfpu=neon-vfpv4 -mtune=cortex-a7 -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_NETWORK_LIB -DQT_WIDGETS_LIB -DQT_QML_LIB -DQT_QUICKWIDGETS_LIB -DQT_QUICK_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi2/include -I/home/user/raspi/sysroot/opt/vc/include -I/home/user/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/user/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I/usr/local/Qt5.7.0/5.7/rpi2/mkspecs/devices/linux-rasp-pi2-g++
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtCore -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtGui -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtNetwork -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtWidgets -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtQml -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtQuickWidgets -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtQuick

#cgo LDFLAGS: -Wl,-rpath-link,/home/user/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/user/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/user/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,/usr/local/Qt5.7.0/5.7/rpi2/lib -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin
#cgo LDFLAGS: -L/home/user/raspi/sysroot/opt/vc/lib -L/usr/local/Qt5.7.0/5.7/rpi2/lib -lQt5Core -lQt5Gui -lQt5Network -lQt5Widgets -lQt5Qml -lQt5QuickWidgets -lQt5Quick -lGLESv2 -lpthread
*/
import "C"
