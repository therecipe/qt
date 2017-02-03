FROM therecipe/qt:base


ENV JDK_DIR /usr/lib/jvm/java-8-openjdk-amd64
RUN add-apt-repository -y ppa:openjdk-r/ppa && apt-get -qq update && apt-get -y -qq install openjdk-8-jdk && apt-get -qq clean && update-ca-certificates -f


ENV SDK tools_r25.2.5-linux.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$SDK https://dl.google.com/android/repository/$SDK && unzip -qq /tmp/$SDK -d $HOME/android-sdk-linux/ && rm -f /tmp/$SDK && $HOME/android-sdk-linux/tools/android list sdk && echo "y" | $HOME/android-sdk-linux/tools/android -s update sdk -f -u -t 1,2,3,4,5,6 && rm -f /tmp/$SDK


ENV NDK android-ndk-r13b-linux-x86_64.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$NDK https://dl.google.com/android/repository/$NDK && unzip -qq /tmp/$NDK -d $HOME && rm -f /tmp/$NDK


RUN $GOPATH/bin/qtsetup check android
RUN $GOPATH/bin/qtsetup generate
