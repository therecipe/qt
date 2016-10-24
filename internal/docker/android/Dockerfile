FROM therecipe/qt:base


ENV JDK_DIR /usr/lib/jvm/java-8-openjdk-amd64
RUN add-apt-repository -y ppa:openjdk-r/ppa && apt-get -qq update && apt-get -y -qq install openjdk-8-jdk && apt-get -qq clean


ENV SDK android-sdk_r24.4.1-linux.tgz
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$SDK https://dl.google.com/android/$SDK && tar -xzf /tmp/$SDK -C $HOME && rm -f /tmp/$SDK && $HOME/android-sdk-linux/tools/android list sdk && echo "y" | $HOME/android-sdk-linux/tools/android -s update sdk -f -u -t 1,2,3,5 && && rm -f /tmp/$SDK
RUN unzip $HOME/android-sdk-linux/temp/tools_r25.2.2-linux.zip -d $HOME/android-sdk-linux/


ENV NDK android-ndk-r12b-linux-x86_64.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$NDK https://dl.google.com/android/repository/$NDK && unzip -qq /tmp/$NDK -d $HOME && rm -f /tmp/$NDK


RUN go run $GOPATH/src/github.com/therecipe/qt/internal/setup/check.go android
RUN go run $GOPATH/src/github.com/therecipe/qt/internal/setup/generate.go
