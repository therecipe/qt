// +build rpi3

package designer

/*
#cgo CFLAGS: -march=armv8-a+crc -mtune=cortex-a53 -mfpu=crypto-neon-fp-armv8 -pipe -Os -mthumb -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -march=armv8-a+crc -mtune=cortex-a53 -mfpu=crypto-neon-fp-armv8 -pipe -Os -mthumb -std=c++1z -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG -DQT_DESIGNER_LIB -DQT_CORE_LIB -DQT_GUI_LIB -DQT_WIDGETS_LIB -DQT_UIPLUGIN_LIB -DQT_DESIGNERCOMPONENTS_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi3/include -I/home/user/raspi/sysroot/opt/vc/include -I/home/user/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/user/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I/usr/local/Qt5.7.0/5.7/rpi3/mkspecs/devices/linux-rpi3-g++
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi3/include/QtDesigner -I/usr/local/Qt5.7.0/5.7/rpi3/include/QtCore -I/usr/local/Qt5.7.0/5.7/rpi3/include/QtGui -I/usr/local/Qt5.7.0/5.7/rpi3/include/QtWidgets -I/usr/local/Qt5.7.0/5.7/rpi3/include/QtUiPlugin -I/usr/local/Qt5.7.0/5.7/rpi3/include/QtDesignerComponents

#cgo LDFLAGS: -Wl,-rpath-link,/home/user/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/user/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/user/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,/usr/local/Qt5.7.0/5.7/rpi3/lib -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin
#cgo LDFLAGS: -L/home/user/raspi/sysroot/opt/vc/lib -L/usr/local/Qt5.7.0/5.7/rpi3/lib -lQt5Designer -lQt5Core -lQt5Gui -lQt5Widgets -lQt5DesignerComponents -lGLESv2 -lpthread
*/
import "C"
