FROM therecipe/qt:linux


ENV JDK_DIR /usr/lib/jvm/java-8-openjdk-amd64
RUN add-apt-repository -y ppa:openjdk-r/ppa && apt-get -qq update && apt-get -y -qq install openjdk-8-jdk && apt-get -qq clean && update-ca-certificates -f


ENV SDK tools_r25.2.5-linux.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$SDK https://dl.google.com/android/repository/$SDK && unzip -qq /tmp/$SDK -d $HOME/android-sdk-linux/ && rm -f /tmp/$SDK && $HOME/android-sdk-linux/tools/android list sdk && echo "y" | $HOME/android-sdk-linux/tools/android -s update sdk -f -u -t 1,2,3,4,5,6 && rm -f /tmp/$SDK


ENV NDK android-ndk-r13b-linux-x86_64.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$NDK https://dl.google.com/android/repository/$NDK && unzip -qq /tmp/$NDK -d $HOME && rm -f /tmp/$NDK


ENV OSSL openssl-1.0.2k.tar.gz
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$OSSL https://www.openssl.org/source/$OSSL && tar -xzf /tmp/$OSSL -C $HOME && rm -f /tmp/$OSSL

ENV CC $HOME/android-ndk-r13b/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64/bin/arm-linux-androideabi-gcc
ENV AR $HOME/android-ndk-r13b/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64/bin/arm-linux-androideabi-ar
ENV ANDROID_DEV $HOME/android-ndk-r13b/platforms/android-16/arch-arm/usr
RUN sed -i 's/LIBCOMPATVERSIONS=/#LIBCOMPATVERSIONS=/' $HOME/openssl-1.0.2k/Makefile
RUN sed -i 's/LIBVERSION=/\\ #LIBVERSION=/' $HOME/openssl-1.0.2k/Makefile
RUN cd $HOME/openssl-1.0.2k && ./Configure shared android-armv7 && make build_libs $HOME/openssl-1.0.2k


RUN $GOPATH/bin/qtsetup check android
RUN $GOPATH/bin/qtsetup generate android
