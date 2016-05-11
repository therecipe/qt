package quick

/*
#cgo CPPFLAGS: -pipe -O2 -Wall -W -D_REENTRANT
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_NETWORK_LIB -DQT_WIDGETS_LIB -DQT_QML_LIB -DQT_QUICKWIDGETS_LIB -DQT_QUICK_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc/include -I/usr/local/Qt5.6.0/5.6/gcc/mkspecs/linux-g++
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc/include/QtCore -I/usr/local/Qt5.6.0/5.6/gcc/include/QtGui -I/usr/local/Qt5.6.0/5.6/gcc/include/QtNetwork -I/usr/local/Qt5.6.0/5.6/gcc/include/QtWidgets -I/usr/local/Qt5.6.0/5.6/gcc/include/QtQml -I/usr/local/Qt5.6.0/5.6/gcc/include/QtQuickWidgets -I/usr/local/Qt5.6.0/5.6/gcc/include/QtQuick

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc/lib
#cgo LDFLAGS: -L/usr/local/Qt5.6.0/5.6/gcc/lib -L/usr/lib -lQt5Core -lQt5Gui -lQt5Network -lQt5Widgets -lQt5Qml -lQt5QuickWidgets -lQt5Quick -lpthread
*/
import "C"
