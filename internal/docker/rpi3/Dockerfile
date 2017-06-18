FROM therecipe/qt:rpi


ONBUILD RUN go get -v github.com/Sirupsen/logrus && go get -v golang.org/x/crypto/ssh && go get -v golang.org/x/tools/imports
ONBUILD RUN $GOPATH/bin/qtsetup upgrade
ONBUILD RUN $GOPATH/bin/qtsetup prep linux
ONBUILD RUN $GOPATH/bin/qtsetup check linux
ONBUILD RUN $GOPATH/bin/qtsetup generate linux
ONBUILD RUN $GOPATH/bin/qtsetup generate rpi3
ONBUILD RUN $GOPATH/bin/qtsetup install rpi3
ONBUILD RUN rm -r /home/user/work/src/github.com/Sirupsen/ && rm -r /home/user/work/src/golang.org/
ONBUILD RUN rm -r /home/user/work/pkg/linux_amd64/github.com/Sirupsen/ && rm -r /home/user/work/pkg/linux_amd64/golang.org/


ENV QT qtrpi-rpi3_qt-5.7.0.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT http://www.qtrpi.com/downloads/qtrpi/rpi3/$QT && unzip -qq /tmp/$QT -d /opt/qtrpi && rm -f /tmp/$QT


RUN $GOPATH/bin/qtsetup check rpi3
RUN $GOPATH/bin/qtsetup generate rpi3
RUN $GOPATH/bin/qtsetup install rpi3

RUN rm -r /home/user/work/src/github.com/Sirupsen/ && rm -r /home/user/work/src/golang.org/
RUN rm -r /home/user/work/pkg/linux_amd64/github.com/Sirupsen/ && rm -r /home/user/work/pkg/linux_amd64/golang.org/
