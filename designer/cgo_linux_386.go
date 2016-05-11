package designer

/*
#cgo CPPFLAGS: -pipe -O2 -Wall -W -D_REENTRANT
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_DESIGNER_LIB -DQT_CORE_LIB -DQT_GUI_LIB -DQT_WIDGETS_LIB -DQT_UIPLUGIN_LIB -DQT_DESIGNERCOMPONENTS_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc/include -I/usr/local/Qt5.6.0/5.6/gcc/mkspecs/linux-g++
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc/include/QtDesigner -I/usr/local/Qt5.6.0/5.6/gcc/include/QtCore -I/usr/local/Qt5.6.0/5.6/gcc/include/QtGui -I/usr/local/Qt5.6.0/5.6/gcc/include/QtWidgets -I/usr/local/Qt5.6.0/5.6/gcc/include/QtUiPlugin -I/usr/local/Qt5.6.0/5.6/gcc/include/QtDesignerComponents

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc/lib
#cgo LDFLAGS: -L/usr/local/Qt5.6.0/5.6/gcc/lib -L/usr/lib -lQt5Designer -lQt5Core -lQt5Gui -lQt5Widgets -lQt5DesignerComponents -lpthread
*/
import "C"
