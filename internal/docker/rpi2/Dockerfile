FROM therecipe/qt:rpi


ENV QT qtrpi-rpi2_qt-5.7.0.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT http://www.qtrpi.com/downloads/qtrpi/rpi2/$QT && unzip -qq /tmp/$QT -d /opt/qtrpi && rm -f /tmp/$QT


ENV QT_QMAKE_DIR /opt/qtrpi/raspi/qt5/bin
ENV RPI_TOOLS_DIR /opt/qtrpi/raspi/tools
ENV RPI_COMPILER gcc-linaro-arm-linux-gnueabihf-raspbian-x64
RUN $GOPATH/bin/qtsetup check rpi2
RUN $GOPATH/bin/qtsetup generate rpi2
