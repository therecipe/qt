package uitools

/*
#cgo CFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -std=c++11 -O2 -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT -fPIC
#cgo CXXFLAGS: -DQT_NO_DEBUG -DQT_UITOOLS_LIB -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/android_armv7/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem /opt/android-ndk/platforms/android-9/arch-arm/usr/include -I/usr/local/Qt5.7.0/5.7/android_armv7/mkspecs/android-g++
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/android_armv7/include/QtUiTools -I/usr/local/Qt5.7.0/5.7/android_armv7/include/QtWidgets -I/usr/local/Qt5.7.0/5.7/android_armv7/include/QtGui -I/usr/local/Qt5.7.0/5.7/android_armv7/include/QtCore

#cgo LDFLAGS: --sysroot=/opt/android-ndk/platforms/android-9/arch-arm -Wl,-rpath=/usr/local/Qt5.7.0/5.7/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition -Wl,--allow-shlib-undefined
#cgo LDFLAGS: -L/opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android-ndk/platforms/android-9/arch-arm//usr/lib -L/usr/local/Qt5.7.0/5.7/android_armv7/lib -L/opt/android/ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android/ndk/platforms/android-9/arch-arm//usr/lib -lQt5UiTools -lQt5Widgets -lQt5Gui -lQt5Core -lGLESv2 -lgnustl_shared -llog -lz -lm -ldl -lc -lgcc
*/
import "C"
