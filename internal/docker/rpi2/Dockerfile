FROM therecipe/qt:rpi


RUN df -h


ENV QT qtrpi-rpi2_qt-5.7.0.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT http://www.qtrpi.com/downloads/qtrpi/rpi2/$QT && unzip -qq /tmp/$QT -d /opt/qtrpi && rm -f /tmp/$QT


RUN $GOPATH/bin/qtsetup check rpi2
RUN $GOPATH/bin/qtsetup generate rpi2
RUN $GOPATH/bin/qtsetup install rpi2
