FROM therecipe/qt:linux


RUN df -h


RUN echo "deb http://pkg.mxe.cc/repos/apt/debian wheezy main" | tee --append /etc/apt/sources.list.d/mxeapt.list > /dev/null && apt-key adv --keyserver keyserver.ubuntu.com --recv-keys D43A795B73B16ABE9643FE1AFD8FFF16DB45C6AB && apt-get -qq update && apt-get -y -qq install mxe-x86-64-w64-mingw32.static-qt3d mxe-x86-64-w64-mingw32.static-qtactiveqt mxe-x86-64-w64-mingw32.static-qtbase mxe-x86-64-w64-mingw32.static-qtcanvas3d mxe-x86-64-w64-mingw32.static-qtcharts mxe-x86-64-w64-mingw32.static-qtconnectivity mxe-x86-64-w64-mingw32.static-qtdatavis3d mxe-x86-64-w64-mingw32.static-qtdeclarative mxe-x86-64-w64-mingw32.static-qtgamepad mxe-x86-64-w64-mingw32.static-qtgraphicaleffects mxe-x86-64-w64-mingw32.static-qtimageformats mxe-x86-64-w64-mingw32.static-qtlocation mxe-x86-64-w64-mingw32.static-qtmultimedia mxe-x86-64-w64-mingw32.static-qtofficeopenxml mxe-x86-64-w64-mingw32.static-qtpurchasing mxe-x86-64-w64-mingw32.static-qtquickcontrols mxe-x86-64-w64-mingw32.static-qtquickcontrols2 mxe-x86-64-w64-mingw32.static-qtscript mxe-x86-64-w64-mingw32.static-qtscxml mxe-x86-64-w64-mingw32.static-qtsensors mxe-x86-64-w64-mingw32.static-qtserialbus mxe-x86-64-w64-mingw32.static-qtserialport mxe-x86-64-w64-mingw32.static-qtservice mxe-x86-64-w64-mingw32.static-qtsvg mxe-x86-64-w64-mingw32.static-qtsystems mxe-x86-64-w64-mingw32.static-qttools mxe-x86-64-w64-mingw32.static-qttranslations mxe-x86-64-w64-mingw32.static-qtvirtualkeyboard mxe-x86-64-w64-mingw32.static-qtwebchannel mxe-x86-64-w64-mingw32.static-qtwebsockets mxe-x86-64-w64-mingw32.static-qtwinextras mxe-x86-64-w64-mingw32.static-qtxlsxwriter mxe-x86-64-w64-mingw32.static-qtxmlpatterns && apt-get -qq clean


ENV QT_MXE true
ENV QT_MXE_ARCH amd64
ENV QT_MXE_STATIC true
RUN $GOPATH/bin/qtsetup check windows
RUN $GOPATH/bin/qtsetup generate windows
RUN $GOPATH/bin/qtsetup install windows
