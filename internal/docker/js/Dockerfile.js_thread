FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_API 5.12.0
ENV QT_DOCKER true
ENV QT_QMAKE_DIR /usr/local/Qt-5.13.0/bin

COPY --from=therecipe/qt:linux /usr/local/go /usr/local/go
COPY --from=therecipe/qt:linux $GOPATH/bin $GOPATH/bin
COPY --from=therecipe/qt:linux $GOPATH/src/github.com/therecipe/qt $GOPATH/src/github.com/therecipe/qt
COPY --from=therecipe/qt:js_base_thread $HOME/emsdk $HOME/emsdk
COPY --from=therecipe/qt:js_base_thread $HOME/.emscripten $HOME/.emscripten
COPY --from=therecipe/qt:js_base_thread /usr/local/Qt-5.13.0 /usr/local/Qt-5.13.0

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install python2.7 nodejs cmake default-jre && apt-get -qq clean
RUN ln -s /usr/bin/python2.7 /usr/bin/python
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install git && apt-get -qq clean
RUN go get github.com/gopherjs/gopherjs

RUN $GOPATH/bin/qtsetup prep
RUN $GOPATH/bin/qtsetup check js
RUN QT_STUB=true $GOPATH/bin/qtsetup generate
RUN $GOPATH/bin/qtsetup generate js
RUN $GOPATH/bin/qtsetup install js
RUN cd $GOPATH/src/github.com/therecipe/qt/internal/examples/widgets/line_edits && $GOPATH/bin/qtdeploy build js && rm -rf ./deploy

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates git pkg-config