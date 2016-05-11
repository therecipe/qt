package serialport

/*
#cgo CPPFLAGS: -pipe -O2 -Wall -W -D_REENTRANT
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_SERIALPORT_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc/include -I/usr/local/Qt5.6.0/5.6/gcc/mkspecs/linux-g++
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc/include/QtCore -I/usr/local/Qt5.6.0/5.6/gcc/include/QtSerialPort

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc/lib
#cgo LDFLAGS: -L/usr/local/Qt5.6.0/5.6/gcc/lib -L/usr/lib -lQt5Core -lQt5SerialPort -lpthread
*/
import "C"
