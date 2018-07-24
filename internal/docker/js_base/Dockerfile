FROM therecipe/qt:linux
LABEL maintainer therecipe


RUN df -h


RUN apt-get -qq update && apt-get -y -qq install python2.7 nodejs cmake default-jre && apt-get -qq clean


RUN git clone https://github.com/juj/emsdk.git $HOME/emsdk && cd $HOME/emsdk && ./emsdk install latest && ./emsdk activate latest


RUN git clone -b wip/webassembly https://code.qt.io/qt/qtbase.git /opt/qtbase


RUN echo "#!/bin/bash\ncd $HOME/emsdk && source ./emsdk_env.sh \
	&& cd /opt/qtbase && ./configure -xplatform wasm-emscripten -developer-build -release -static -no-thread -nomake tests -nomake examples -no-dbus -no-headersclean -system-libpng -no-ssl -no-warnings-are-errors -confirm-license -opensource && make" > $HOME/build.sh \
	&& chmod +x $HOME/build.sh && $HOME/build.sh