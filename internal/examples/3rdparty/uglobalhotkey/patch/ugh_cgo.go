package UGlobalHotkey

/*
#cgo CXXFLAGS: -DUGLOBALHOTKEY_LIBRARY

#cgo linux CXXFLAGS: -I/opt/Qt/5.13.0/gcc_64/include/QtGui/5.13.0/QtGui
#cgo linux LDFLAGS: -lxcb -lxcb-keysyms

#cgo darwin LDFLAGS: -framework Carbon

#cgo windows,!msvc LDFLAGS: -luser32
#cgo windows,msvc MSLDFLAGS: user32.lib
#cgo windows,msvc MSCXXFLAGS: -DUGLOBALHOTKEY_LIBRARY
*/
import "C"

import "github.com/therecipe/qt/widgets"

type stub struct{ widgets.QWidget }
