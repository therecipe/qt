package quick

/*
#cgo CPPFLAGS: -pipe -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.7 -Wall -W
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_NETWORK_LIB -DQT_WIDGETS_LIB -DQT_QML_LIB -DQT_QUICKWIDGETS_LIB -DQT_QUICK_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/clang_64/mkspecs/macx-clang -F/usr/local/Qt5.5.1/5.5/clang_64/lib
#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtCore.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtGui.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtNetwork.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtWidgets.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtQml.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtQuickWidgets.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtQuick.framework/Headers
#cgo CPPFLAGS: -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/AGL.framework/Headers

#cgo LDFLAGS: -headerpad_max_install_names -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.7
#cgo LDFLAGS: -F/usr/local/Qt5.5.1/5.5/clang_64/lib -framework QtCore -framework QtGui -framework QtNetwork -framework QtWidgets -framework QtQml -framework QtQuickWidgets -framework QtQuick -framework DiskArbitration -framework IOKit -framework OpenGL -framework AGL
*/
import "C"
