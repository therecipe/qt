package core

/*
#cgo CPPFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra
#cgo CPPFLAGS: -DUNICODE -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_CORE_LIB
#cgo CPPFLAGS: -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include -IC:/Qt/Qt5.5.1/5.5/mingw492_32/mkspecs/win32-g++
#cgo CPPFLAGS: -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtWidgets -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtCore

#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads
#cgo LDFLAGS: -LC:/Qt/Qt5.5.1/5.5/mingw492_32/lib -lQt5Widgets -lQt5Core -lmingw32 -lqtmain -lshell32
*/
import "C"
