FROM therecipe/qt:rpi
LABEL maintainer therecipe


RUN df -h


ENV QT qtrpi-rpi1_qt-5.7.0.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT http://www.qtrpi.com/downloads/qtrpi/rpi1/$QT && unzip -qq /tmp/$QT -d /opt/qtrpi && rm -f /tmp/$QT


RUN git clone -b 5.7 https://code.qt.io/qt/qtvirtualkeyboard.git /opt/qtrpi/raspi/qtvirtualkeyboard && cd /opt/qtrpi/raspi/qtvirtualkeyboard && ../qt5/bin/qmake qtvirtualkeyboard.pro CONFIG+=lang-all && make && make install && rm -r /opt/qtrpi/raspi/qtvirtualkeyboard


RUN $GOPATH/bin/qtsetup check rpi1
RUN $GOPATH/bin/qtsetup generate rpi1
RUN $GOPATH/bin/qtsetup install rpi1
