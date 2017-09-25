FROM ubuntu:14.04
LABEL maintainer therecipe


RUN df -h


ENV USER user
ENV HOME /home/$USER
ENV QT_DOCKER true


RUN apt-get -qq update && apt-get -y -qq install curl git software-properties-common libgl1-mesa-dev fontconfig unzip && apt-get -qq clean
RUN apt-get -qq update && apt-get -y -qq install bison build-essential gperf flex ruby python libasound2-dev libbz2-dev libcap-dev libcups2-dev libdrm-dev libegl1-mesa-dev libgcrypt11-dev libnss3-dev libpci-dev libpulse-dev libudev-dev libxtst-dev gyp ninja-build && apt-get -qq clean
RUN apt-get -qq update && apt-get -y -qq install libssl-dev libxcursor-dev libxcomposite-dev libxdamage-dev libxrandr-dev libfontconfig1-dev libxss-dev libsrtp0-dev libwebp-dev libjsoncpp-dev libopus-dev libavutil-dev libavformat-dev libavcodec-dev libevent-dev libxslt1-dev && apt-get -qq clean


RUN add-apt-repository -y ppa:ubuntu-toolchain-r/test && apt-get -qq update && apt-get -y -qq install gcc-5 g++-5 && apt-get -qq clean
RUN rm -f /usr/bin/gcc; sudo ln -s /usr/bin/gcc-5 /usr/bin/gcc
RUN rm -f /usr/bin/g++; sudo ln -s /usr/bin/g++-5 /usr/bin/g++


ENV GO go1.9.linux-amd64.tar.gz
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$GO https://storage.googleapis.com/golang/$GO && tar -xzf /tmp/$GO -C /usr/local && rm -f /tmp/$GO
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH $HOME/work


RUN go get -v github.com/therecipe/qt/cmd/...
RUN rm -r /home/user/work/src/github.com/Sirupsen/ && rm -r /home/user/work/src/golang.org/
RUN rm -r /home/user/work/pkg/linux_amd64/github.com/Sirupsen/ && rm -r /home/user/work/pkg/linux_amd64/golang.org/


ENV QT qt-unified-linux-x64-online.run
ENV QT_QPA_PLATFORM minimal
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT https://download.qt.io/official_releases/online_installers/$QT && chmod +x /tmp/$QT && /tmp/$QT --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs && rm -f /tmp/$QT


ENV QT_DIR /opt/Qt
RUN $GOPATH/bin/qtsetup prep linux
RUN $GOPATH/bin/qtsetup check linux
RUN $GOPATH/bin/qtsetup generate linux
RUN $GOPATH/bin/qtsetup install linux
