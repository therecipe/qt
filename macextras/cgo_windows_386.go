package macextras

/*
#cgo CPPFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra
#cgo CPPFLAGS: -DUNICODE -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_MACEXTRAS_LIB
#cgo CPPFLAGS: -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include -IC:/Qt/Qt5.5.1/5.5/mingw492_32/mkspecs/win32-g++
#cgo CPPFLAGS: -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtCore -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtGui -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtMacExtras

#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition
#cgo LDFLAGS: -LC:/Qt/Qt5.5.1/5.5/mingw492_32/lib -lQt5Core -lQt5Gui -lQt5MacExtras -lmingw32 -lqtmain -lshell32
*/
import "C"
