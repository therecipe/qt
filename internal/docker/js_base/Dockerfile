FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install build-essential git

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install python2.7 nodejs cmake default-jre
RUN ln -s /usr/bin/python2.7 /usr/bin/python
RUN git clone https://github.com/juj/emsdk.git $HOME/emsdk && cd $HOME/emsdk && ./emsdk install latest && ./emsdk activate latest

RUN git clone -b wip/webassembly https://code.qt.io/qt/qtbase.git /opt/qtbase
RUN echo "#!/bin/bash\nsource $HOME/emsdk/emsdk_env.sh \
	&& cd /opt/qtbase && ./configure -xplatform wasm-emscripten -developer-build -release -static -no-thread -nomake tests -nomake examples -no-dbus -no-headersclean -system-libpng -no-ssl -no-warnings-are-errors -confirm-license -opensource && make" > $HOME/build.sh \
	&& chmod +x $HOME/build.sh && $HOME/build.sh
