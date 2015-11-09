package xmlpatterns

/*
#cgo CPPFLAGS: -Wno-psabi -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -ffunction-sections -funwind-tables -fstack-protector -fno-short-enums -DANDROID -Wa,--noexecstack -fno-builtin-memmove -Os -fomit-frame-pointer -fno-strict-aliasing -finline-limit=64 -mthumb -Wall -Wno-psabi -W -D_REENTRANT
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_NETWORK_LIB -DQT_XMLPATTERNS_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/android_armv7/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem /opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -isystem /opt/android-ndk/platforms/android-9/arch-arm/usr/include -I/usr/local/Qt5.5.1/5.5/android_armv7/mkspecs/android-g++
#cgo CPPFLAGS: -I/usr/local/Qt5.5.1/5.5/android_armv7/include/QtCore -I/usr/local/Qt5.5.1/5.5/android_armv7/include/QtNetwork -I/usr/local/Qt5.5.1/5.5/android_armv7/include/QtXmlPatterns

#cgo LDFLAGS: --sysroot=/opt/android-ndk/platforms/android-9/arch-arm/ -Wl,-rpath=/usr/local/Qt5.5.1/5.5/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack -Wl,--allow-multiple-definition
#cgo LDFLAGS: -L/opt/android-ndk/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android-ndk/platforms/android-9/arch-arm//usr/lib -L/usr/local/Qt5.5.1/5.5/android_armv7/lib -L/opt/android/ndk/sources/cxx-stl/gnu-libstdc++/4.8/libs/armeabi-v7a -L/opt/android/ndk/platforms/android-9/arch-arm//usr/lib -lQt5Core -lQt5Network -lQt5XmlPatterns -lGLESv2 -lgnustl_shared -llog -lz -lm -ldl -lc -lgcc
*/
import "C"
