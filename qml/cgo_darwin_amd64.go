package qml

/*
#cgo CPPFLAGS: -pipe -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.8 -Wall -W -fPIC
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_CORE_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/clang_64/mkspecs/macx-clang -F/usr/local/Qt5.6.0/5.6/clang_64/lib
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/clang_64/lib/QtQml.framework/Headers -I/usr/local/Qt5.6.0/5.6/clang_64/lib/QtNetwork.framework/Headers -I/usr/local/Qt5.6.0/5.6/clang_64/lib/QtCore.framework/Headers
#cgo CPPFLAGS: -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk/System/Library/Frameworks/AGL.framework/Headers

#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.11.sdk -mmacosx-version-min=10.8 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/clang_64/lib
#cgo LDFLAGS: -F/usr/local/Qt5.6.0/5.6/clang_64/lib -framework QtQml -framework QtNetwork -framework QtCore -framework DiskArbitration -framework IOKit -framework OpenGL -framework AGL
*/
import "C"
