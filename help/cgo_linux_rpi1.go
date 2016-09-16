// +build rpi1

package help

/*
#cgo CFLAGS: -pipe -marm -mfpu=vfp -mtune=arm1176jzf-s -march=armv6zk -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -pipe -marm -mfpu=vfp -mtune=arm1176jzf-s -march=armv6zk -mabi=aapcs-linux -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_NETWORK_LIB -DQT_SQL_LIB -DQT_CLUCENE_LIB -DQT_WIDGETS_LIB -DQT_HELP_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi1/include -I/home/user/raspi/sysroot/opt/vc/include -I/home/user/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/user/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I/usr/local/Qt5.7.0/5.7/rpi1/mkspecs/devices/linux-rasp-pi-g++
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi1/include/QtCore -I/usr/local/Qt5.7.0/5.7/rpi1/include/QtGui -I/usr/local/Qt5.7.0/5.7/rpi1/include/QtNetwork -I/usr/local/Qt5.7.0/5.7/rpi1/include/QtSql -I/usr/local/Qt5.7.0/5.7/rpi1/include/QtCLucene -I/usr/local/Qt5.7.0/5.7/rpi1/include/QtWidgets -I/usr/local/Qt5.7.0/5.7/rpi1/include/QtHelp

#cgo LDFLAGS: -Wl,-rpath-link,/home/user/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/user/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/user/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,/usr/local/Qt5.7.0/5.7/rpi1/lib -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin
#cgo LDFLAGS: -L/home/user/raspi/sysroot/opt/vc/lib -L/usr/local/Qt5.7.0/5.7/rpi1/lib -lQt5Core -lQt5Gui -lQt5Network -lQt5Sql -lQt5CLucene -lQt5Widgets -lQt5Help -lGLESv2 -lpthread
*/
import "C"
