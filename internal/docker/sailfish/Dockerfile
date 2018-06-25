FROM i386/ubuntu:16.04
LABEL maintainer therecipe


RUN df -h


ENV USER user
ENV HOME /home/$USER
ENV QT_DOCKER true
ENV QT_SAILFISH true


RUN apt-get -qq update && apt-get -y -qq install curl git && apt-get -qq clean
RUN apt-get -qq update && apt-get -y -qq purge python && apt-get -qq clean


ENV GO go1.10.3.linux-386.tar.gz
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$GO https://storage.googleapis.com/golang/$GO && tar -xzf /tmp/$GO -C /usr/local && rm -f /tmp/$GO
ENV PATH /usr/local/go/bin:$PATH
ENV GOPATH $HOME/work


RUN go get -v github.com/therecipe/qt/cmd/...
RUN go clean -i github.com/sirupsen/... && go clean -i golang.org/...
RUN rm -r $GOPATH/src/github.com/sirupsen/ && rm -r $GOPATH/src/golang.org/


ENV SF_SDK_TOOLING Jolla-latest-Sailfish_SDK_Tooling-i486.tar.bz2
ENV SF_SDK_TOOLING_DIR /srv/mer/toolings/SailfishOS-2.1.4.13
RUN mkdir -p $SF_SDK_TOOLING_DIR
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$SF_SDK_TOOLING http://releases.sailfishos.org/sdk/latest/$SF_SDK_TOOLING && tar --numeric-owner -p -xjf /tmp/$SF_SDK_TOOLING -C $SF_SDK_TOOLING_DIR && rm -f /tmp/$SF_SDK_TOOLING


ENV SF_SDK_TARGET Jolla-latest-Sailfish_SDK_Target-i486.tar.bz2
RUN mkdir -p /srv/mer/targets/SailfishOS-2.1.4.13-i486
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$SF_SDK_TARGET http://releases.sailfishos.org/sdk/latest/$SF_SDK_TARGET && tar --numeric-owner -p -xjf /tmp/$SF_SDK_TARGET -C /srv/mer/targets/SailfishOS-2.1.4.13-i486 && rm -f /tmp/$SF_SDK_TARGET


ENV SF_SDK_TARGET Jolla-latest-Sailfish_SDK_Target-armv7hl.tar.bz2
RUN mkdir -p /srv/mer/targets/SailfishOS-2.1.4.13-armv7hl
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$SF_SDK_TARGET http://releases.sailfishos.org/sdk/latest/$SF_SDK_TARGET && tar --numeric-owner -p -xjf /tmp/$SF_SDK_TARGET -C /srv/mer/targets/SailfishOS-2.1.4.13-armv7hl && rm -f /tmp/$SF_SDK_TARGET


ENV PATH $SF_SDK_TOOLING_DIR/usr/bin/:$PATH


RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libmpc.so.3 /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libmpfr.so.4 /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libgmp.so.10 /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libpthread_nonshared.a /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libc_nonshared.a /usr/lib/


RUN ln -s $SF_SDK_TOOLING_DIR/opt/cross/bin/i486-meego-linux-gnu-as $SF_SDK_TOOLING_DIR/opt/cross/libexec/gcc/i486-meego-linux-gnu/4.8.3/as
RUN ln -s $SF_SDK_TOOLING_DIR/opt/cross/bin/i486-meego-linux-gnu-ld $SF_SDK_TOOLING_DIR/opt/cross/libexec/gcc/i486-meego-linux-gnu/4.8.3/ld

RUN ln -s $SF_SDK_TOOLING_DIR/opt/cross/bin/armv7hl-meego-linux-gnueabi-as $SF_SDK_TOOLING_DIR/opt/cross/libexec/gcc/armv7hl-meego-linux-gnueabi/4.8.3/as
RUN ln -s $SF_SDK_TOOLING_DIR/opt/cross/bin/armv7hl-meego-linux-gnueabi-ld $SF_SDK_TOOLING_DIR/opt/cross/libexec/gcc/armv7hl-meego-linux-gnueabi/4.8.3/ld


RUN cd /srv/mer/targets/SailfishOS-2.1.4.13-i486/ && $SF_SDK_TOOLING_DIR/usr/bin/sb2-init -L "--sysroot=/" -C "--sysroot=/" -n -N -t / i486-meego-linux-gnu $SF_SDK_TOOLING_DIR/opt/cross/bin/i486-meego-linux-gnu-gcc

RUN cd /srv/mer/targets/SailfishOS-2.1.4.13-i486/ && $SF_SDK_TOOLING_DIR/usr/bin/sb2-init -L "--sysroot=/" -C "--sysroot=/" -n -N -t / armv7hl-meego-linux $SF_SDK_TOOLING_DIR/opt/cross/bin/i486-meego-linux-gnu-gcc


RUN sed -i 's/--target=$build_tgt/--target=$OPT_TARGET/g' /srv/mer/toolings/SailfishOS-2.1.4.13/usr/bin/mb2


RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/librpmbuild.so.2 /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/librpm.so.2 /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/librpmio.so.2 /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libdb-4.8.so /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/liblua-5.1.so /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libpython2.7.so.1.0 /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libbfd-2.25.so /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libnss3.so /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libelf.so.1 /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libnssutil3.so /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libplc4.so /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libplds4.so /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libnspr4.so /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libglib-2.0.so.0 /usr/lib/
RUN ln -s $SF_SDK_TOOLING_DIR/usr/lib/libpcre.so.1 /usr/lib/


RUN $GOPATH/bin/qtsetup prep sailfish
RUN $GOPATH/bin/qtsetup generate sailfish
RUN $GOPATH/bin/qtsetup install sailfish

RUN $GOPATH/bin/qtsetup prep sailfish-emulator
RUN $GOPATH/bin/qtsetup generate sailfish-emulator
RUN $GOPATH/bin/qtsetup install sailfish-emulator
