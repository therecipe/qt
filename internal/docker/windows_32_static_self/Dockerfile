FROM therecipe/qt:windows_32_static_self_base_ext_2 as base

RUN cd /usr/lib/mxe && make MXE_TARGETS='i686-w64-mingw32.static' qt5
RUN rm -rf /usr/lib/mxe/log && rm -rf /usr/lib/mxe/pkg


FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/lib/mxe/usr/bin:/usr/local/go/bin:$PATH
ENV QT_DIR /opt/Qt
ENV QT_DOCKER true
ENV QT_MXE true
ENV QT_MXE_ARCH 386
ENV QT_MXE_STATIC true

COPY --from=therecipe/qt:linux /usr/local/go /usr/local/go
COPY --from=therecipe/qt:linux $GOPATH/bin $GOPATH/bin
COPY --from=therecipe/qt:linux $GOPATH/src/github.com/therecipe/qt $GOPATH/src/github.com/therecipe/qt
COPY --from=therecipe/qt:linux /opt/Qt/5.11.1/gcc_64/include /opt/Qt/5.11.1/gcc_64/include
COPY --from=base /usr/lib/mxe /usr/lib/mxe

RUN $GOPATH/bin/qtsetup prep
RUN $GOPATH/bin/qtsetup check windows
RUN $GOPATH/bin/qtsetup generate windows
RUN $GOPATH/bin/qtsetup install windows
RUN cd $GOPATH/src/github.com/therecipe/qt/internal/examples/widgets/line_edits && $GOPATH/bin/qtdeploy build windows && rm -rf ./deploy
