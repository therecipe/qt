package quick

/*
#cgo CFLAGS: -pipe -O2 -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -pipe -O2 -std=gnu++11 -Wall -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -DQT_NO_DEBUG -DQT_QUICK_LIB -DQT_QUICKWIDGETS_LIB -DQT_WIDGETS_LIB -DQT_NETWORK_LIB -DQT_QML_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/gcc_64/include -I/usr/local/Qt5.7.0/5.7/gcc_64/mkspecs/linux-g++
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtQuick -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtQuickWidgets -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtWidgets -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtNetwork -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtQml -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtGui -I/usr/local/Qt5.7.0/5.7/gcc_64/include/QtCore

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.7.0/5.7/gcc_64 -Wl,-rpath,/usr/local/Qt5.7.0/5.7/gcc_64/lib
#cgo LDFLAGS: -L/usr/local/Qt5.7.0/5.7/gcc_64/lib -L/usr/lib64 -lQt5Quick -lQt5QuickWidgets -lQt5Widgets -lQt5Network -lQt5Qml -lQt5Gui -lQt5Core -lpthread
*/
import "C"
