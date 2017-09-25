FROM therecipe/qt:linux


RUN df -h


ENV JDK_DIR /usr/lib/jvm/java-8-openjdk-amd64
RUN add-apt-repository -y ppa:openjdk-r/ppa && apt-get -qq update && apt-get -y -qq install openjdk-8-jdk && apt-get -qq clean && update-ca-certificates -f


ENV SDK sdk-tools-linux-3859397.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$SDK https://dl.google.com/android/repository/$SDK && unzip -qq /tmp/$SDK -d $HOME/android-sdk-linux/ && rm -f /tmp/$SDK
RUN $HOME/android-sdk-linux/tools/bin/sdkmanager --list --verbose
RUN echo "y" | $HOME/android-sdk-linux/tools/bin/sdkmanager "platform-tools" "build-tools;26.0.0" "platforms;android-25"
RUN $HOME/android-sdk-linux/tools/bin/sdkmanager --update


ENV NDK android-ndk-r14b-linux-x86_64.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$NDK https://dl.google.com/android/repository/$NDK && unzip -qq /tmp/$NDK -d $HOME && rm -f /tmp/$NDK


ENV OSSL openssl-1.0.2k.tar.gz
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$OSSL https://www.openssl.org/source/$OSSL && tar -xzf /tmp/$OSSL -C $HOME && rm -f /tmp/$OSSL

ENV CC $HOME/android-ndk-r14b/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64/bin/arm-linux-androideabi-gcc
ENV AR $HOME/android-ndk-r14b/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64/bin/arm-linux-androideabi-ar
ENV ANDROID_DEV $HOME/android-ndk-r14b/platforms/android-16/arch-arm/usr
RUN sed -i 's/LIBCOMPATVERSIONS=/#LIBCOMPATVERSIONS=/' $HOME/openssl-1.0.2k/Makefile
RUN sed -i 's/LIBVERSION=/\\ #LIBVERSION=/' $HOME/openssl-1.0.2k/Makefile
RUN cd $HOME/openssl-1.0.2k && ./Configure shared android-armv7 && make build_libs $HOME/openssl-1.0.2k


RUN $GOPATH/bin/qtsetup check android
RUN $GOPATH/bin/qtsetup generate android
RUN $GOPATH/bin/qtsetup install android

RUN $GOPATH/bin/qtdeploy build android $GOPATH/src/github.com/therecipe/qt/internal/examples/androidextras/jni
