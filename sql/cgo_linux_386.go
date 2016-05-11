package sql

/*
#cgo CPPFLAGS: -pipe -O2 -Wall -W -D_REENTRANT
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_WIDGETS_LIB -DQT_SQL_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc/include -I/usr/local/Qt5.6.0/5.6/gcc/mkspecs/linux-g++
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc/include/QtCore -I/usr/local/Qt5.6.0/5.6/gcc/include/QtWidgets -I/usr/local/Qt5.6.0/5.6/gcc/include/QtSql

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc/lib
#cgo LDFLAGS: -L/usr/local/Qt5.6.0/5.6/gcc/lib -L/usr/lib -lQt5Core -lQt5Widgets -lQt5Sql -lpthread
*/
import "C"
