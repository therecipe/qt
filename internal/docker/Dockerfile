FROM ubuntu:14.04
MAINTAINER therecipe


ENV USER user
ENV HOME /home/$USER


RUN apt-get -qq update && apt-get -y -qq install curl git software-properties-common libgl1-mesa-dev fontconfig && apt-get -qq clean


RUN add-apt-repository -y ppa:ubuntu-toolchain-r/test && apt-get -qq update && apt-get -y -qq install gcc-5 g++-5 && apt-get -qq clean
RUN ln -s /usr/bin/gcc-5 /usr/bin/gcc && ln -s /usr/bin/g++-5 /usr/bin/g++


ENV QT qt-opensource-linux-x64-android-5.7.0.run
ENV QT_SCRIPT iscript.qs
ENV QT_QPA_PLATFORM minimal
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT https://download.qt.io/official_releases/qt/5.7/5.7.0/$QT && chmod +x /tmp/$QT && curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT_SCRIPT https://raw.githubusercontent.com/therecipe/qt/master/internal/ci/iscript.qs && /tmp/$QT --script /tmp/$QT_SCRIPT && rm -f /tmp/$QT && rm -f /tmp/$QT_SCRIPT


ENV GO go1.7.3.linux-amd64.tar.gz
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$GO https://storage.googleapis.com/golang/$GO && tar -xzf /tmp/$GO -C /usr/local && rm -f /tmp/$GO
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH $HOME/work


RUN go get -d github.com/therecipe/qt
RUN go get -d golang.org/x/crypto/ssh
RUN go get -d github.com/emirpasic/gods/lists/arraylist
RUN go build -o /usr/local/bin/qtdeploy $GOPATH/src/github.com/therecipe/qt/internal/deploy
RUN go build -o /usr/local/bin/qtmoc $GOPATH/src/github.com/therecipe/qt/internal/moc


ENV QT_DIR /opt/Qt5.7.0
RUN go run $GOPATH/src/github.com/therecipe/qt/internal/setup/check.go
RUN go run $GOPATH/src/github.com/therecipe/qt/internal/setup/generate.go
