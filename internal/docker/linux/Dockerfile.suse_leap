FROM opensuse/leap:latest as base

ENV GOPATH $HOME/work

RUN zypper -q ref && zypper -n -q install --no-recommends curl git gzip tar
RUN GO=go1.12.4.linux-amd64.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/go/$GO && tar -xzf $GO -C /usr/local
RUN /usr/local/go/bin/go get -tags=no_env github.com/therecipe/qt/cmd/...

FROM opensuse/leap:latest
LABEL maintainer therecipe

ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_DOCKER true
ENV QT_PKG_CONFIG true

COPY --from=base /usr/local/go /usr/local/go
COPY --from=base $GOPATH/bin $GOPATH/bin
COPY --from=base $GOPATH/src/github.com/therecipe/qt $GOPATH/src/github.com/therecipe/qt

RUN zypper -q ref && zypper -n -q install --no-recommends -t pattern devel_basis && zypper clean -a
RUN zypper -q ref && zypper -n -q install --no-recommends libqt5-qt*-devel && zypper clean -a

RUN $GOPATH/bin/qtsetup prep
RUN $GOPATH/bin/qtsetup check
RUN $GOPATH/bin/qtsetup generate
RUN cd $GOPATH/src/github.com/therecipe/qt/internal/examples/widgets/line_edits && $GOPATH/bin/qtdeploy build linux && rm -rf ./deploy

RUN zypper -q ref && zypper -n -q install --no-recommends git pkg-config