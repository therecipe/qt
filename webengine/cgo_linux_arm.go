package webengine

/*
#cgo CPPFLAGS: -pipe -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -Wformat -Wformat-security -fmessage-length=0 -march=armv7-a -mfloat-abi=hard -mfpu=neon -mthumb -Wno-psabi -fPIC -fvisibility=hidden -fvisibility-inlines-hidden -Wall -W -D_REENTRANT -fPIE
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_NETWORK_LIB -DQT_WEBCHANNEL_LIB -DQT_WIDGETS_LIB -DQT_WEBENGINECORE_LIB -DQT_WEBENGINEWIDGETS_LIB -DQT_WEBENGINE_LIB
#cgo CPPFLAGS: -I/srv/mer/targets/SailfishOS-armv7hl/usr/share/qt5/mkspecs/linux-g++ -I/srv/mer/targets/SailfishOS-armv7hl/usr/include -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/sailfishapp -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/mdeclarativecache5 -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5
#cgo CPPFLAGS: -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtCore -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtGui -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtNetwork -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtWebChannel -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtWidgets -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtWebEngineCore -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtWebEngineWidgets -I/srv/mer/targets/SailfishOS-armv7hl/usr/include/qt5/QtWebEngine

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath-link,/srv/mer/targets/SailfishOS-armv7hl/usr/lib -Wl,-rpath-link,/srv/mer/targets/SailfishOS-armv7hl/lib
#cgo LDFLAGS: -rdynamic -L/srv/mer/targets/SailfishOS-armv7hl/usr/lib -L/srv/mer/targets/SailfishOS-armv7hl/lib -lsailfishapp -lmdeclarativecache5 -lQt5Core -lQt5Gui -lQt5Network -lGLESv2 -lpthread
*/
import "C"
