FROM therecipe/qt:js_base
LABEL maintainer therecipe


RUN df -h


RUN git clone -b wip/webassembly https://code.qt.io/qt/qtdeclarative.git /opt/qtdeclarative
RUN git clone -b wip/webassembly https://code.qt.io/qt/qtwebsockets.git /opt/qtwebsockets
RUN git clone -b 5.11.0 https://code.qt.io/qt/qtgraphicaleffects.git /opt/qtgraphicaleffects
RUN git clone -b 5.11.0 https://code.qt.io/qt/qtquickcontrols.git /opt/qtquickcontrols
RUN git clone -b 5.11.0 https://code.qt.io/qt/qtquickcontrols2.git /opt/qtquickcontrols2
RUN git clone -b 5.11.0 https://code.qt.io/qt/qtsvg.git /opt/qtsvg


RUN echo "#!/bin/bash\ncd $HOME/emsdk && source ./emsdk_env.sh \
	&& cd /opt/qtdeclarative && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtwebsockets && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtgraphicaleffects && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtquickcontrols && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtquickcontrols2 && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtsvg && /opt/qtbase/bin/qmake *.pro && make" > $HOME/build.sh \
	&& chmod +x $HOME/build.sh && $HOME/build.sh