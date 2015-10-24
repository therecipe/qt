package macextras

/*
#cgo CPPFLAGS: -pipe -O2 -Wall -W -D_REENTRANT
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_MACEXTRAS_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/gcc/include -I/usr/local/Qt5.5.1/5.5/gcc/mkspecs/linux-g++
#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/gcc/include/QtCore -I/usr/local/Qt5.5.1/5.5/gcc/include/QtGui -I/usr/local/Qt5.5.1/5.5/gcc/include/QtMacExtras

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.5.1/5.5/gcc -Wl,-rpath,/usr/local/Qt5.5.1/5.5/gcc/lib
#cgo LDFLAGS: -L/usr/local/Qt5.5.1/5.5/gcc/lib -L/usr/lib -lQt5Core -lQt5Gui -lQt5MacExtras -lpthread
*/
import "C"
