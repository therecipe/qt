package webkit

/*
#cgo CPPFLAGS: -pipe -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.7 -Wall -W -fPIC
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_WEBKIT_LIB -DQT_WEBKITWIDGETS_LIB -DQT_PRINTSUPPORT_LIB -DQT_WIDGETS_LIB -DQT_WEBCHANNEL_LIB -DQT_NETWORK_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/clang_64/mkspecs/macx-clang -F/usr/local/Qt5.5.1/5.5/clang_64/lib
#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtWebKit.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtWebKitWidgets.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtPrintSupport.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtWidgets.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtWebChannel.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtNetwork.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtGui.framework/Headers -I/usr/local/Qt5.5.1/5.5/clang_64/lib/QtCore.framework/Headers
#cgo CPPFLAGS: -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/AGL.framework/Headers

#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.7 -Wl,-rpath,/usr/local/Qt5.5.1/5.5/clang_64/lib
#cgo LDFLAGS: -F/usr/local/Qt5.5.1/5.5/clang_64/lib -framework QtWebKit -framework QtWebKitWidgets -framework QtPrintSupport -framework QtWidgets -framework QtWebChannel -framework QtNetwork -framework QtGui -framework QtCore -framework DiskArbitration -framework IOKit -framework OpenGL -framework AGL
*/
import "C"
