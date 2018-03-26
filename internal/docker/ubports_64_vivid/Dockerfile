FROM clickable/ubuntu-sdk:15.04-amd64


RUN df -h


ENV USER user
ENV HOME /home/$USER


RUN apt-get -qq update && apt-get -y -qq install curl git && apt-get -qq clean
RUN apt-get -qq update && apt-get -y -qq install qt*5-doc-html && apt-get -qq clean


RUN rm -R /usr/local/go | true
ENV GO go1.10.linux-amd64.tar.gz
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$GO https://storage.googleapis.com/golang/$GO && tar -xzf /tmp/$GO -C /usr/local && rm -f /tmp/$GO
ENV PATH /usr/local/go/bin:$PATH
ENV GOPATH $HOME/work


RUN go get -v github.com/therecipe/qt/cmd/...
RUN go clean -i github.com/sirupsen/... && go clean -i golang.org/...
RUN rm -r $GOPATH/src/github.com/sirupsen/ && rm -r $GOPATH/src/golang.org/


ENV QT_PKG_CONFIG true
RUN $GOPATH/bin/qtsetup prep linux
RUN $GOPATH/bin/qtsetup check linux
RUN $GOPATH/bin/qtsetup generate linux
RUN $GOPATH/bin/qtsetup install linux
