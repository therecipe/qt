FROM clickable/ubuntu-sdk:15.04-armhf
LABEL maintainer therecipe


RUN df -h


ENV USER user
ENV HOME /home/$USER
ENV QT_DOCKER true


RUN apt-get -qq update && apt-get -y -qq install curl git && apt-get -qq clean
RUN apt-get -qq update && apt-get -y -qq install qt*5-doc-html && apt-get -qq clean


RUN rm -R /usr/local/go | true
ENV GO go1.10.3.linux-amd64.tar.gz
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$GO https://storage.googleapis.com/golang/$GO && tar -xzf /tmp/$GO -C /usr/local && rm -f /tmp/$GO
ENV PATH /usr/local/go/bin:$PATH
ENV GOPATH $HOME/work


ENV GOARCH amd64
ENV CC gcc
ENV CXX g++
RUN go get -v github.com/therecipe/qt/cmd/...
RUN go clean -i github.com/sirupsen/... && go clean -i golang.org/...
RUN rm -r $GOPATH/src/github.com/sirupsen/ && rm -r $GOPATH/src/golang.org/
ENV GOARCH arm
ENV CC arm-linux-gnueabihf-gcc
ENV CXX arm-linux-gnueabihf-g++


RUN ln -s /usr/lib/x86_64-linux-gnu/qt5/bin/rcc /usr/lib/arm-linux-gnueabihf/qt5/bin/ 
RUN ln -s /usr/lib/x86_64-linux-gnu/qt5/bin/moc /usr/lib/arm-linux-gnueabihf/qt5/bin/ 


ENV QT_PKG_CONFIG true
ENV QT_UBPORTS true
ENV QT_UBPORTS_ARCH arm
ENV QT_UBPORTS_VERSION vivid
RUN $GOPATH/bin/qtsetup prep ubports
RUN $GOPATH/bin/qtsetup check ubports
RUN $GOPATH/bin/qtsetup generate ubports
RUN $GOPATH/bin/qtsetup install ubports
