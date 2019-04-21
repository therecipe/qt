FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_API=5.7.0
ENV QT_DOCKER true
ENV QT_RPI true
ENV QT_QMAKE_DIR /opt/qtrpi/raspi/qt5/bin
ENV QT_VERSION 5.7.0
ENV RPI_TOOLS_DIR /opt/qtrpi/raspi/tools
ENV RPI_COMPILER gcc-linaro-arm-linux-gnueabihf-raspbian-x64


COPY --from=therecipe/qt:linux /usr/local/go /usr/local/go
COPY --from=therecipe/qt:linux $GOPATH/bin $GOPATH/bin
COPY --from=therecipe/qt:linux $GOPATH/src/github.com/therecipe/qt $GOPATH/src/github.com/therecipe/qt

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl git make && apt-get -qq clean

RUN SYSROOT=qtrpi-sysroot-minimal-latest.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://github.com/therecipe/files/releases/download/v0.0.0/$SYSROOT && tar -xzf $SYSROOT -C / && rm -f $SYSROOT
RUN ln -s /opt/qtrpi/raspbian/sysroot-minimal /opt/qtrpi/raspbian/sysroot

RUN git clone -q --depth 1 https://github.com/raspberrypi/tools.git /opt/qtrpi/raspi/tools
