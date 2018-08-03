FROM ubuntu:16.04 as base

ENV USER user
ENV HOME /home/$USER

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl make openjdk-8-jdk perl unzip
RUN mkdir -p $HOME

RUN SDK=sdk-tools-linux-3859397.zip && curl -sL --retry 10 --retry-delay 10 -O https://dl.google.com/android/repository/$SDK && unzip -qq $SDK -d $HOME/android-sdk-linux
RUN $HOME/android-sdk-linux/tools/bin/sdkmanager --list --verbose
RUN echo "y" | $HOME/android-sdk-linux/tools/bin/sdkmanager "platform-tools" "build-tools;26.0.0" "platforms;android-25"
RUN $HOME/android-sdk-linux/tools/bin/sdkmanager --update

RUN NDK=android-ndk-r14b-linux-x86_64.zip && curl -sL --retry 10 --retry-delay 10 -O https://dl.google.com/android/repository/$NDK && unzip -qq $NDK -d $HOME

RUN OSSL=openssl-1.0.2o.tar.gz && curl -sL --retry 10 --retry-delay 10 -O https://www.openssl.org/source/$OSSL && tar -xzf $OSSL -C $HOME

ENV CC $HOME/android-ndk-r14b/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64/bin/arm-linux-androideabi-gcc
ENV AR $HOME/android-ndk-r14b/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64/bin/arm-linux-androideabi-ar
ENV ANDROID_DEV $HOME/android-ndk-r14b/platforms/android-16/arch-arm/usr
RUN sed -i 's/LIBCOMPATVERSIONS=/#LIBCOMPATVERSIONS=/' $HOME/openssl-1.0.2o/Makefile
RUN sed -i 's/LIBVERSION=/\\ #LIBVERSION=/' $HOME/openssl-1.0.2o/Makefile
RUN cd $HOME/openssl-1.0.2o && ./Configure shared android-armv7 && make build_libs $HOME/openssl-1.0.2o
RUN mkdir -p $HOME/openssl-1.0.2o-arm && mv $HOME/openssl-1.0.2o/*.so $HOME/openssl-1.0.2o-arm

RUN rm -rf $HOME/openssl-1.0.2o && tar -xzf openssl-1.0.2o.tar.gz -C $HOME

ENV CC $HOME/android-ndk-r14b/toolchains/x86-4.9/prebuilt/linux-x86_64/bin/i686-linux-android-gcc
ENV AR $HOME/android-ndk-r14b/toolchains/x86-4.9/prebuilt/linux-x86_64/bin/i686-linux-android-ar
ENV ANDROID_DEV $HOME/android-ndk-r14b/platforms/android-16/arch-x86/usr
RUN sed -i 's/LIBCOMPATVERSIONS=/#LIBCOMPATVERSIONS=/' $HOME/openssl-1.0.2o/Makefile
RUN sed -i 's/LIBVERSION=/\\ #LIBVERSION=/' $HOME/openssl-1.0.2o/Makefile
RUN cd $HOME/openssl-1.0.2o && ./Configure shared android-x86 && make build_libs $HOME/openssl-1.0.2o
RUN mkdir -p $HOME/openssl-1.0.2o-x86 && mv $HOME/openssl-1.0.2o/*.so $HOME/openssl-1.0.2o-x86


FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV JDK_DIR /usr/lib/jvm/java-8-openjdk-amd64
ENV PATH /usr/local/go/bin:$PATH
ENV QT_DIR /opt/Qt
ENV QT_DOCKER true

COPY --from=therecipe/qt:linux /usr/local/go /usr/local/go
COPY --from=therecipe/qt:linux $GOPATH/bin $GOPATH/bin
COPY --from=therecipe/qt:linux $GOPATH/src/github.com/therecipe/qt $GOPATH/src/github.com/therecipe/qt
COPY --from=therecipe/qt:linux /opt/Qt/5.11.1/gcc_64/include /opt/Qt/5.11.1/gcc_64/include
COPY --from=therecipe/qt:linux /opt/Qt/5.11.1/android_armv7 /opt/Qt/5.11.1/android_armv7
COPY --from=therecipe/qt:linux /opt/Qt/5.11.1/android_x86 /opt/Qt/5.11.1/android_x86
COPY --from=therecipe/qt:linux /opt/Qt/Docs /opt/Qt/Docs
COPY --from=base $HOME/android-sdk-linux $HOME/android-sdk-linux
COPY --from=base $HOME/android-ndk-r14b $HOME/android-ndk-r14b
COPY --from=base $HOME/openssl-1.0.2o-arm $HOME/openssl-1.0.2o-arm
COPY --from=base $HOME/openssl-1.0.2o-x86 $HOME/openssl-1.0.2o-x86

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates openjdk-8-jdk && apt-get -qq clean

RUN $GOPATH/bin/qtsetup prep

RUN $GOPATH/bin/qtsetup check android
RUN $GOPATH/bin/qtsetup generate android
RUN $GOPATH/bin/qtsetup install android
RUN cd $GOPATH/src/github.com/therecipe/qt/internal/examples/androidextras/jni && $GOPATH/bin/qtdeploy build android | true && rm -rf ./deploy

RUN $GOPATH/bin/qtsetup check android-emulator
RUN $GOPATH/bin/qtsetup generate android-emulator
RUN $GOPATH/bin/qtsetup install android-emulator
RUN cd $GOPATH/src/github.com/therecipe/qt/internal/examples/androidextras/jni && $GOPATH/bin/qtdeploy build android-emulator | true && rm -rf ./deploy
