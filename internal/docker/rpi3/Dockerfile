FROM therecipe/qt:rpi
LABEL maintainer therecipe


RUN df -h


ENV QT qtrpi-rpi3_qt-5.7.0.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT http://www.qtrpi.com/downloads/qtrpi/rpi3/$QT && unzip -qq /tmp/$QT -d /opt/qtrpi && rm -f /tmp/$QT


RUN git clone -b 5.7 https://code.qt.io/qt/qtvirtualkeyboard.git /opt/qtrpi/raspi/qtvirtualkeyboard && cd /opt/qtrpi/raspi/qtvirtualkeyboard && ../qt5/bin/qmake qtvirtualkeyboard.pro CONFIG+=lang-all && make && make install && rm -r /opt/qtrpi/raspi/qtvirtualkeyboard


RUN $GOPATH/bin/qtsetup check rpi3
RUN $GOPATH/bin/qtsetup generate rpi3
RUN $GOPATH/bin/qtsetup install rpi3
