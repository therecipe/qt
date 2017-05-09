FROM ubuntu:14.04
LABEL maintainer therecipe


ONBUILD RUN $GOPATH/bin/qtsetup update


ENV USER user
ENV HOME /home/$USER
ENV QT_DOCKER true


RUN apt-get -qq update && apt-get -y -qq install curl git software-properties-common libgl1-mesa-dev fontconfig unzip && apt-get -qq clean
RUN apt-get -qq update && apt-get -y -qq install bison build-essential gperf flex ruby python libasound2-dev libbz2-dev libcap-dev libcups2-dev libdrm-dev libegl1-mesa-dev libgcrypt11-dev libnss3-dev libpci-dev libpulse-dev libudev-dev libxtst-dev gyp ninja-build && apt-get -qq clean
RUN apt-get -qq update && apt-get -y -qq install libssl-dev libxcursor-dev libxcomposite-dev libxdamage-dev libxrandr-dev libfontconfig1-dev libxss-dev libsrtp0-dev libwebp-dev libjsoncpp-dev libopus-dev libavutil-dev libavformat-dev libavcodec-dev libevent-dev libxslt1-dev libpulse-dev && apt-get -qq clean


RUN add-apt-repository -y ppa:ubuntu-toolchain-r/test && apt-get -qq update && apt-get -y -qq install gcc-5 g++-5 && apt-get -qq clean
RUN rm /usr/bin/gcc; sudo ln -s /usr/bin/gcc-5 /usr/bin/gcc
RUN rm /usr/bin/g++; sudo ln -s /usr/bin/g++-5 /usr/bin/g++


ENV QT qt-opensource-linux-x64-android-5.8.0.run
ENV QT_SCRIPT iscript.qs
ENV QT_QPA_PLATFORM minimal
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT https://download.qt.io/official_releases/qt/5.8/5.8.0/$QT && chmod +x /tmp/$QT && curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT_SCRIPT https://raw.githubusercontent.com/therecipe/qt/master/internal/ci/iscript.qs && /tmp/$QT --script /tmp/$QT_SCRIPT && rm -f /tmp/$QT && rm -f /tmp/$QT_SCRIPT


ENV GO go1.8.1.linux-amd64.tar.gz
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$GO https://storage.googleapis.com/golang/$GO && tar -xzf /tmp/$GO -C /usr/local && rm -f /tmp/$GO
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH $HOME/work


RUN go get -v github.com/therecipe/qt/cmd/...


ENV QT_DIR /opt/Qt5.8.0
RUN $GOPATH/bin/qtsetup prep linux
RUN $GOPATH/bin/qtsetup check linux
RUN $GOPATH/bin/qtsetup generate linux
RUN $GOPATH/bin/qtsetup install linux
