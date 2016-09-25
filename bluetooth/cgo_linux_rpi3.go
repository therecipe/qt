// +build rpi3

package bluetooth

/*
#cgo CFLAGS: -march=armv8-a+crc -mtune=cortex-a53 -mfpu=crypto-neon-fp-armv8 -pipe -Os -mthumb -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -O2 -fno-exceptions -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -march=armv8-a+crc -mtune=cortex-a53 -mfpu=crypto-neon-fp-armv8 -pipe -Os -mthumb -std=c++11 -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -O2 -std=gnu++11 -fno-exceptions -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -DQT_NO_EXCEPTIONS -D_LARGEFILE64_SOURCE -D_LARGEFILE_SOURCE -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_CONCURRENT_LIB -DQT_BLUETOOTH_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi3/include -I/home/user/raspi/sysroot/opt/vc/include -I/home/user/raspi/sysroot/opt/vc/include/interface/vcos/pthreads -I/home/user/raspi/sysroot/opt/vc/include/interface/vmcs_host/linux -I/usr/local/Qt5.7.0/5.7/rpi3/mkspecs/devices/linux-rpi3-g++
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/rpi3/include/QtCore -I/usr/local/Qt5.7.0/5.7/rpi3/include/QtConcurrent -I/usr/local/Qt5.7.0/5.7/rpi3/include/QtBluetooth

#cgo LDFLAGS: -Wl,-rpath-link,/home/user/raspi/sysroot/opt/vc/lib -Wl,-rpath-link,/home/user/raspi/sysroot/usr/lib/arm-linux-gnueabihf -Wl,-rpath-link,/home/user/raspi/sysroot/lib/arm-linux-gnueabihf -Wl,-rpath-link,/usr/local/Qt5.7.0/5.7/rpi3/lib -mfloat-abi=hard --sysroot=/home/user/raspi/sysroot -Wl,-O1 -Wl,--enable-new-dtags -Wl,-z,origin
#cgo LDFLAGS: -L/home/user/raspi/sysroot/opt/vc/lib -L/usr/local/Qt5.7.0/5.7/rpi3/lib -lQt5Core -lQt5Concurrent -lQt5Bluetooth -lGLESv2 -lpthread
*/
import "C"
