FROM ubuntu:16.04 as base

ENV USER user
ENV HOME /home/$USER

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl make openjdk-8-jdk perl unzip
RUN mkdir -p $HOME

RUN SDK=sdk-tools-linux-3859397.zip && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/android/repository/$SDK && unzip -qq $SDK -d $HOME/android-sdk-linux
RUN $HOME/android-sdk-linux/tools/bin/sdkmanager --list --verbose
RUN echo "y" | $HOME/android-sdk-linux/tools/bin/sdkmanager "platform-tools" "build-tools;28.0.3" "platforms;android-28"
RUN $HOME/android-sdk-linux/tools/bin/sdkmanager --update


RUN NDK=android-ndk-r18b-linux-x86_64.zip && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/android/repository/$NDK && unzip -qq $NDK -d $HOME

RUN OSSL=openssl-1.1.1a.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://www.openssl.org/source/$OSSL && tar -xzf $OSSL -C $HOME

ENV ANDROID_NDK $HOME/android-ndk-r18b
ENV PATH $ANDROID_NDK/toolchains/llvm/prebuilt/linux-x86_64/bin:$PATH

RUN cd $HOME/openssl-1.1.1a && ./Configure -D__ANDROID_API__=21 android-arm && make -j $(grep -c ^processor /proc/cpuinfo) SHLIB_VERSION_NUMBER= SHLIB_EXT=.so build_libs $HOME/openssl-1.1.1a
RUN mkdir -p $HOME/openssl-1.1.1a-arm && mv $HOME/openssl-1.1.1a/*.so* $HOME/openssl-1.1.1a-arm
RUN rm -rf $HOME/openssl-1.1.1a && tar -xzf openssl-1.1.1a.tar.gz -C $HOME
RUN cd $HOME/openssl-1.1.1a && ./Configure -D__ANDROID_API__=21 android-x86 && make -j $(grep -c ^processor /proc/cpuinfo) SHLIB_VERSION_NUMBER= SHLIB_EXT=.so build_libs $HOME/openssl-1.1.1a
RUN mkdir -p $HOME/openssl-1.1.1a-x86 && mv $HOME/openssl-1.1.1a/*.so* $HOME/openssl-1.1.1a-x86
RUN find $HOME/openssl-1.1.1a-* -type f -exec llvm-strip --strip-all {} \;


FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV JDK_DIR /usr/lib/jvm/java-8-openjdk-amd64
ENV PATH /usr/local/go/bin:$PATH
ENV QT_VERSION 5.13.0
ENV QT_DIR /opt/Qt
ENV QT_DOCKER true

COPY --from=therecipe/qt:linux /usr/local/go /usr/local/go
COPY --from=therecipe/qt:linux $GOPATH/bin $GOPATH/bin
COPY --from=therecipe/qt:linux $GOPATH/src/github.com/therecipe/qt $GOPATH/src/github.com/therecipe/qt
COPY --from=therecipe/qt:linux /opt/Qt/5.13.0/android_armv7 /opt/Qt/5.13.0/android_armv7
COPY --from=therecipe/qt:linux /opt/Qt/5.13.0/android_x86 /opt/Qt/5.13.0/android_x86
COPY --from=therecipe/qt:linux /opt/Qt/5.13.0/android_arm64_v8a /opt/Qt/5.13.0/android_arm64_v8a
COPY --from=therecipe/qt:linux /opt/Qt/Docs /opt/Qt/Docs
COPY --from=base $HOME/android-sdk-linux $HOME/android-sdk-linux
COPY --from=base $HOME/android-ndk-r18b $HOME/android-ndk-r18b
COPY --from=base $HOME/openssl-1.1.1a-arm $HOME/openssl-1.1.1a-arm
COPY --from=base $HOME/openssl-1.1.1a-x86 $HOME/openssl-1.1.1a-x86

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates openjdk-8-jdk && apt-get -qq clean

RUN $GOPATH/bin/qtsetup prep

RUN $GOPATH/bin/qtsetup check android
RUN $GOPATH/bin/qtsetup generate android
RUN $GOPATH/bin/qtsetup install android
RUN cd $GOPATH/src/github.com/therecipe/qt/internal/examples/androidextras/jni && $GOPATH/bin/qtdeploy build android && rm -rf ./deploy

RUN GOARCH=arm64 $GOPATH/bin/qtsetup check android
RUN GOARCH=arm64 $GOPATH/bin/qtsetup generate android
RUN GOARCH=arm64 $GOPATH/bin/qtsetup install android
RUN cd $GOPATH/src/github.com/therecipe/qt/internal/examples/androidextras/jni && GOARCH=arm64 $GOPATH/bin/qtdeploy build android && rm -rf ./deploy

RUN $GOPATH/bin/qtsetup check android-emulator
RUN $GOPATH/bin/qtsetup generate android-emulator
RUN $GOPATH/bin/qtsetup install android-emulator
RUN cd $GOPATH/src/github.com/therecipe/qt/internal/examples/androidextras/jni && $GOPATH/bin/qtdeploy build android-emulator && rm -rf ./deploy

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates git pkg-config