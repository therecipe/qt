FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/lib/mxe/usr/bin:/usr/local/go/bin:$PATH
ENV QT_DOCKER true
ENV QT_MXE true
ENV QT_MXE_ARCH amd64
ENV QT_MXE_STATIC false

COPY --from=bluszcz/cutego:linux /usr/local/go /usr/local/go
COPY --from=bluszcz/cutego:linux $GOPATH/bin $GOPATH/bin
COPY --from=bluszcz/cutego:linux $GOPATH/src/github.com/StarAurryon/qt $GOPATH/src/github.com/StarAurryon/qt
COPY --from=bluszcz/cutego:windows_64_shared_base /usr/lib/mxe /usr/lib/mxe

RUN $GOPATH/bin/qtsetup prep
RUN $GOPATH/bin/qtsetup check windows
RUN $GOPATH/bin/qtsetup generate windows
RUN $GOPATH/bin/qtsetup install windows
RUN cd $GOPATH/src/github.com/StarAurryon/qt/internal/examples/widgets/line_edits && $GOPATH/bin/qtdeploy build windows && rm -rf ./deploy

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates git pkg-config