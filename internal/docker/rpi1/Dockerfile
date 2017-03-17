FROM therecipe/qt:rpi


ENV QT qtrpi-rpi1_qt-5.7.0.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT http://www.qtrpi.com/downloads/qtrpi/rpi1/$QT && unzip -qq /tmp/$QT -d /opt/qtrpi && rm -f /tmp/$QT


ENV QT_QMAKE_DIR /opt/qtrpi/raspi/qt5/bin
ENV RPI_TOOLS_DIR /opt/qtrpi/raspi/tools
ENV RPI_COMPILER gcc-linaro-arm-linux-gnueabihf-raspbian-x64
RUN $GOPATH/bin/qtsetup check rpi1
RUN $GOPATH/bin/qtsetup generate rpi1
