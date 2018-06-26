FROM therecipe/qt:js_ext
LABEL maintainer therecipe


RUN df -h


RUN git clone -b 5.11.0 https://code.qt.io/qt/qtcharts.git /opt/qtcharts
RUN git clone -b 5.11.0 https://code.qt.io/qt/qtxmlpatterns.git /opt/qtxmlpatterns
RUN git clone -b 5.11.0 https://code.qt.io/qt/qtcanvas3d.git /opt/qtcanvas3d
RUN git clone -b 5.11.0 https://code.qt.io/qt/qtdatavis3d.git /opt/qtdatavis3d
RUN git clone -b 5.11.0 https://code.qt.io/qt/qtmultimedia.git /opt/qtmultimedia


RUN echo "#!/bin/bash\ncd $HOME/emsdk && source ./emsdk_env.sh \
	&& cd /opt/qtcharts && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtxmlpatterns && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtcanvas3d && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtdatavis3d && /opt/qtbase/bin/qmake *.pro && make \
	&& cd /opt/qtmultimedia && /opt/qtbase/bin/qmake *.pro && make" > $HOME/build.sh \
	&& chmod +x $HOME/build.sh && $HOME/build.sh


RUN go get -v github.com/gopherjs/gopherjs


ENV QT_QMAKE_DIR /opt/qtbase/bin