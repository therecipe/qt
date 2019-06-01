FROM ubuntu:16.04 as base

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl git
RUN GO=go1.12.4.linux-amd64.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/go/$GO && tar -xzf $GO -C /usr/local
RUN /usr/local/go/bin/go get -tags=no_env github.com/therecipe/qt/cmd/...


FROM therecipe/qt:linux_static_base as fcitx

ENV CMAKE_PREFIX_PATH /opt/Qt/5.12.0/gcc_64/lib/cmake/Qt5/

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install build-essential libglib2.0-dev libglu1-mesa-dev \
	&& apt-get --no-install-recommends -qq -y install ca-certificates cmake git libxkbcommon-dev

RUN git clone -q --depth 1 https://github.com/fcitx/fcitx-qt5.git
RUN FILE=/fcitx-qt5/platforminputcontext/CMakeLists.txt \ 
	&& echo "ADD_DEFINITIONS(-DQT_STATICPLUGIN)\n$(cat $FILE)" > $FILE \
	&& echo "find_package(Qt5 REQUIRED COMPONENTS Core Gui DBus)\n$(cat $FILE)" > $FILE \
	&& echo "$(head -n -5 $FILE)" > $FILE \
	&& sed -i -e 's/ MODULE / STATIC /g' $FILE
RUN cd /fcitx-qt5/platforminputcontext && cmake . &&  make -j $(grep -c ^processor /proc/cpuinfo)


FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_API 5.12.0
ENV QT_DIR /opt/Qt
ENV QT_DOCKER true
ENV QT_STATIC true


COPY --from=base /usr/local/go /usr/local/go
COPY --from=base $GOPATH/bin $GOPATH/bin
COPY --from=base $GOPATH/src/github.com/therecipe/qt $GOPATH/src/github.com/therecipe/qt
COPY --from=therecipe/qt:linux_static_base /opt/Qt/5.12.0 /opt/Qt/5.12.0
COPY --from=fcitx /fcitx-qt5/platforminputcontext/libfcitxplatforminputcontextplugin.a /opt/Qt/5.12.0/gcc_64/plugins/platforminputcontexts/libfcitxplatforminputcontextplugin.a

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install build-essential libglib2.0-dev libglu1-mesa-dev libpulse-dev \
	&& apt-get --no-install-recommends -qq -y install fontconfig libasound2 libegl1-mesa libnss3 libpci3 libxcomposite1 libxcursor1 libxi6 libxrandr2 libxtst6 && apt-get -qq clean
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install libfontconfig1-dev libfreetype6-dev libxrender-dev libxkbcommon-dev && apt-get -qq clean

RUN $GOPATH/bin/qtsetup prep
RUN $GOPATH/bin/qtsetup check
RUN $GOPATH/bin/qtsetup generate
RUN cd $GOPATH/src/github.com/therecipe/qt/internal/examples/widgets/line_edits && $GOPATH/bin/qtdeploy build linux && rm -rf ./deploy

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates git pkg-config