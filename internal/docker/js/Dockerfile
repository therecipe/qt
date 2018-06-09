FROM therecipe/qt:linux
LABEL maintainer therecipe


RUN df -h


RUN apt-get -qq update && apt-get -y -qq install python2.7 nodejs cmake default-jre && apt-get -qq clean


RUN git clone https://github.com/juj/emsdk.git $HOME/emsdk && cd $HOME/emsdk && ./emsdk install latest && ./emsdk activate latest


RUN git clone -b wip/webassembly https://code.qt.io/qt/qtbase.git /opt/qtbase #8a63cd94fbba0afafc00b467277a2648143d2890
RUN git clone -b wip/webassembly https://code.qt.io/qt/qtdeclarative.git /opt/qtdeclarative #5f4698c4e7952d64e78e3a06de5ac28d68415b19
RUN git clone -b wip/webassembly https://code.qt.io/qt/qtwebsockets.git /opt/qtwebsockets #5c64a39482cba6ad4bc4337bd7475f96cfad9109
RUN git clone https://code.qt.io/qt/qtgraphicaleffects.git /opt/qtgraphicaleffects
RUN git clone https://code.qt.io/qt/qtquickcontrols.git /opt/qtquickcontrols
RUN git clone https://code.qt.io/qt/qtquickcontrols2.git /opt/qtquickcontrols2
RUN git clone https://code.qt.io/qt/qtsvg.git /opt/qtsvg
RUN git clone https://code.qt.io/qt/qtcharts.git /opt/qtcharts
RUN git clone https://code.qt.io/qt/qtxmlpatterns.git /opt/qtxmlpatterns


RUN echo "#!/bin/bash\ncd $HOME/emsdk && source ./emsdk_env.sh \
	&& cd /opt/qtbase && ./configure -xplatform emscripten -developer-build -release -static -no-thread -nomake tests -nomake examples -no-dbus -no-headersclean -system-libpng -no-ssl -no-warnings-are-errors -confirm-license -opensource && make \
	&& cd /opt/qtdeclarative && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtwebsockets && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtgraphicaleffects && git checkout 80b9da7ac42468453b7d130bd0a322a1312ad0e3 && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtquickcontrols && git checkout 88ff352d5145b5761179ee42164a0bfcd4f1d54a && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtquickcontrols2 && git checkout 60f3106b3fe537f1f26230bf0b59525cf98d6789 && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtsvg && git checkout b1ca1bb725fd8768782d8105b782932f95b70908 && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtcharts && git checkout 4f66d1ffeecd95b4bf6666b71536c42dea7450ed && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtxmlpatterns && git checkout 17c45f3235f0bcb224379b0d1059c7669eb91487 && /opt/qtbase/bin/qmake *.pro && make" > $HOME/build.sh \
	&& chmod +x $HOME/build.sh && $HOME/build.sh


RUN go get -v github.com/gopherjs/gopherjs


ENV QT_QMAKE_DIR /opt/qtbase/bin