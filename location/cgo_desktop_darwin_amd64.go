// +build !ios

package location

/*
#cgo CFLAGS: -pipe -O2 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk -mmacosx-version-min=10.8 -Wall -W -fPIC
#cgo CXXFLAGS: -pipe -stdlib=libc++ -O2 -std=gnu++11 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk -mmacosx-version-min=10.8 -Wall -W -fPIC
#cgo CXXFLAGS: -DQT_NO_DEBUG -DQT_LOCATION_LIB -DQT_POSITIONING_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/clang_64/mkspecs/macx-clang -F/usr/local/Qt5.7.0/5.7/clang_64/lib
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/clang_64/lib/QtLocation.framework/Headers -I/usr/local/Qt5.7.0/5.7/clang_64/lib/QtPositioning.framework/Headers -I/usr/local/Qt5.7.0/5.7/clang_64/lib/QtQuick.framework/Headers -I/usr/local/Qt5.7.0/5.7/clang_64/lib/QtGui.framework/Headers -I/usr/local/Qt5.7.0/5.7/clang_64/lib/QtCore.framework/Headers
#cgo CXXFLAGS: -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk/System/Library/Frameworks/AGL.framework/Headers

#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk -mmacosx-version-min=10.8 -Wl,-rpath,/usr/local/Qt5.7.0/5.7/clang_64/lib
#cgo LDFLAGS: -F/usr/local/Qt5.7.0/5.7/clang_64/lib -framework QtLocation -framework QtPositioning -framework QtQuick -framework QtGui -framework QtCore -framework DiskArbitration -framework IOKit -framework OpenGL -framework AGL
*/
import "C"
