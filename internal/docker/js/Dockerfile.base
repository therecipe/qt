FROM ubuntu:16.04 as base

ENV USER user
ENV HOME /home/$USER

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install build-essential git

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install python2.7 nodejs cmake default-jre
RUN ln -s /usr/bin/python2.7 /usr/bin/python
RUN git clone -q --depth 1 https://github.com/juj/emsdk.git $HOME/emsdk && cd $HOME/emsdk && ./emsdk install latest && ./emsdk activate latest

RUN git clone -q --depth 1 -b 5.13 --recursive https://code.qt.io/qt/qt5.git /opt/qt5
RUN sed -i -e 's/Q_VKB_IMPORT_PLUGIN(QmlFolderListModelPlugin)//g' /opt/qt5/qtvirtualkeyboard/src/import/qtquickvirtualkeyboardplugin.cpp

RUN echo "#!/bin/bash\nsource $HOME/emsdk/emsdk_env.sh \
	&& cd /opt/qt5 && ./configure -xplatform wasm-emscripten -optimize-size -nomake tests -nomake examples -skip qtpim -skip qtfeedback -skip qtwinextras -skip qttools -confirm-license -opensource && make -j $(grep -c ^processor /proc/cpuinfo) && make install -j $(grep -c ^processor /proc/cpuinfo)" > $HOME/build.sh \
	&& chmod +x $HOME/build.sh && $HOME/build.sh


FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER

COPY --from=base $HOME/emsdk $HOME/emsdk
COPY --from=base $HOME/.emscripten $HOME/.emscripten
COPY --from=base /usr/local/Qt-5.13.0 /usr/local/Qt-5.13.0
