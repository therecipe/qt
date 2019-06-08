FROM ubuntu:16.04 as base

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install curl cmake file git build-essential xz-utils software-properties-common

RUN curl -sL --retry 10 --retry-delay 60 https://apt.llvm.org/llvm-snapshot.gpg.key | apt-key add -
RUN apt-add-repository "deb http://apt.llvm.org/xenial/ llvm-toolchain-xenial-6.0 main"
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install clang-6.0 && ln -s /usr/bin/clang-6.0 /usr/bin/clang && ln -s /usr/bin/clang++-6.0 /usr/bin/clang++

RUN git clone -q --depth 1 https://github.com/tpoechtrager/osxcross.git
#RUN sed -i -e 's,/usr/include/c++/v1,/usr/include/c++,g' /osxcross/wrapper/target.cpp
#RUN sed -i -e 's,test_compiler o,#test_compiler o,g' -e 's,-f "have,true || -f "have,g' /osxcross/build.sh

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install automake libxml2-dev libssl-dev
#RUN git clone -q --depth 1 https://github.com/mackyle/xar.git && cd /xar/xar && ./autogen.sh --prefix=/osxcross/target && make -j $(grep -c ^processor /proc/cpuinfo) && make install -j $(grep -c ^processor /proc/cpuinfo) && ./autogen.sh --prefix=/usr && make -j $(grep -c ^processor /proc/cpuinfo) && make install -j $(grep -c ^processor /proc/cpuinfo)

#ENV LD_LIBRARY_PATH=/osxcross/target/lib
#RUN git clone -q --depth 1 -b 1000.10.8 https://github.com/tpoechtrager/apple-libtapi.git && cd /apple-libtapi && INSTALLPREFIX=/osxcross/target ./build.sh && ./install.sh
#RUN git clone -q --depth 1 https://github.com/tpoechtrager/cctools-port.git && cd /cctools-port/cctools && ./configure --prefix=/osxcross/target --target=x86_64-apple-darwin18 --with-libtapi=/osxcross/target && make -j $(grep -c ^processor /proc/cpuinfo) && make install -j $(grep -c ^processor /proc/cpuinfo)


#RUN curl -sL --retry 10 --retry-delay 60 -o /osxcross/tarballs/MacOSX10.14.sdk.tar.xz https://.../MacOSX10.14.sdk.tar.xz
COPY MacOSX10.14.sdk.tar.xz /osxcross/tarballs/MacOSX10.14.sdk.tar.xz


RUN cd /osxcross && OSX_VERSION_MIN=10.12 UNATTENDED=1 ./build.sh
#RUN find /osxcross/target/ -name '*.tbd' -exec sed -i'' s/zippered/macosx/ {} \;
RUN cd /osxcross && ./target/bin/o64-clang++-libc++ -o osxcross ./oclang/test_libcxx.cpp


FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_API 5.12.0
ENV QT_DOCKER true
ENV QT_QMAKE_DIR $GOPATH/src/github.com/therecipe/env_darwin_amd64_512/5.12.0/clang_64/bin


RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install curl git build-essential libglib2.0-dev software-properties-common && apt-get -qq clean

COPY --from=base /osxcross/target /osxcross/target
RUN ln -s /osxcross/target/bin/xcrun /usr/bin/xcrun
RUN rm /usr/bin/strip && ln -s /osxcross/target/bin/x86_64-apple-darwin18-strip /usr/bin/strip
RUN ln -s /osxcross/target/bin/x86_64-apple-darwin18-otool /usr/bin/otool
RUN ln -s /osxcross/target/bin/x86_64-apple-darwin18-install_name_tool /usr/bin/install_name_tool

RUN GO=go1.12.4.linux-amd64.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/go/$GO && tar -xzf $GO -C /usr/local
RUN go get github.com/therecipe/qt/cmd/...
RUN GOOS=darwin go get github.com/therecipe/qt/cmd/... && DST=$GOPATH/src/github.com/therecipe/env_darwin_amd64_512/5.12.0/clang_64/bin && rm -r $DST && cp -r $GOPATH/src/github.com/therecipe/env_linux_amd64_512/5.12.0/gcc_64/bin $DST

RUN sed -i -e 's/isEmpty(QMAKE_XCODEBUILD_PATH)/if(false)/g' -e 's/!if/if/g' -e 's/!if/if/g' -e 's/isEmpty(QMAKE/if(false){#/g' $GOPATH/src/github.com/therecipe/env_darwin_amd64_512/5.12.0/clang_64/mkspecs/features/mac/default_pre.prf
RUN sed -i -e 's/$$QMAKE_MAC_SDK_PATH/$$(QMAKE_MAC_SDK_PATH)/g' $GOPATH/src/github.com/therecipe/env_darwin_amd64_512/5.12.0/clang_64/mkspecs/features/mac/default_post.prf
ENV QMAKE_MAC_SDK_PATH /osxcross/target/SDK/MacOSX10.14.sdk

RUN curl -sL --retry 10 --retry-delay 60 https://apt.llvm.org/llvm-snapshot.gpg.key | apt-key add -
RUN apt-add-repository "deb http://apt.llvm.org/xenial/ llvm-toolchain-xenial-6.0 main"
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install clang-6.0 && apt-get -qq clean && ln -s /usr/bin/clang-6.0 /usr/bin/clang && ln -s /usr/bin/clang++-6.0 /usr/bin/clang++

ENV LD_LIBRARY_PATH=$GOPATH/src/github.com/therecipe/env_linux_amd64_512/5.12.0/gcc_64/lib
ENV QT_PLUGIN_PATH=$GOPATH/src/github.com/therecipe/env_linux_amd64_512/5.12.0/gcc_64/plugins
RUN git clone -q --depth 1 -b 5.13 https://github.com/qt/qttools.git && \
cd /qttools/src/macdeployqt/macdeployqt && \ 
sed -i -e 's/LIBS/#LIBS/g' macdeployqt.pro && \
sed -i -e 's,QString binaryPath;,QString binaryPath = QString(appBundlePath); binaryPath.chop(4); auto bS = binaryPath.split("/"); binaryPath.append(".app/Contents/MacOS/").append(bS.at(bS.size()-1)); return binaryPath;,g' ../shared/shared.cpp && \
$GOPATH/src/github.com/therecipe/env_linux_amd64_512/5.12.0/gcc_64/bin/qmake macdeployqt.pro && make -j $(grep -c ^processor /proc/cpuinfo) && \
mv /qttools/bin/macdeployqt $GOPATH/src/github.com/therecipe/env_darwin_amd64_512/5.12.0/clang_64/bin/macdeployqt && rm -rf /qttools

ENV CC /osxcross/target/bin/x86_64-apple-darwin18-clang
ENV CXX /osxcross/target/bin/x86_64-apple-darwin18-clang++

RUN $GOPATH/bin/qtsetup prep darwin
RUN $GOPATH/bin/qtsetup check darwin
RUN $GOPATH/bin/qtsetup generate darwin
RUN cd $GOPATH/src/github.com/therecipe/qt/internal/examples/widgets/line_edits && $GOPATH/bin/qtdeploy build darwin && rm -rf ./deploy

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates git pkg-config