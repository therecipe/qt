// +build rpi2

package webengine

/*
#cgo CFLAGS: -pipe -march=armv7-a -marm -mthumb-interwork -mfpu=neon-vfpv4 -mtune=cortex-a7 -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -pipe -march=armv7-a -marm -mthumb-interwork -mfpu=neon-vfpv4 -mtune=cortex-a7 -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG -DQT_WEBENGINE_LIB -DQT_WIDGETS_LIB -DQT_WEBENGINEWIDGETS_LIB -DQT_WEBCHANNEL_LIB -DQT_NETWORK_LIB -DQT_WEBENGINECORE_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi2/include -I/home/user/raspi/sysroot/opt/vc/include -I/home/user/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/user/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I/usr/local/Qt5.7.0/5.7/rpi2/mkspecs/devices/linux-rasp-pi2-g++
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtWebEngine -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtWidgets -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtWebEngineWidgets -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtWebChannel -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtNetwork -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtWebEngineCore -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtQuick -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtGui -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtQml -I/usr/local/Qt5.7.0/5.7/rpi2/include/QtCore

#cgo LDFLAGS: -Wl,-rpath-link,/home/user/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/user/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/user/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,/usr/local/Qt5.7.0/5.7/rpi2/lib -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin
#cgo LDFLAGS: -L/home/user/raspi/sysroot/opt/vc/lib -L/usr/local/Qt5.7.0/5.7/rpi2/lib -lQt5WebEngine -lQt5Widgets -lQt5WebEngineWidgets -lQt5WebChannel -lQt5Network -lQt5WebEngineCore -lQt5Quick -lQt5Gui -lQt5Qml -lQt5Core -lGLESv2 -lpthread
*/
import "C"
