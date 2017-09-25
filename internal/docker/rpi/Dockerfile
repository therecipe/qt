FROM therecipe/qt:linux


RUN df -h


ENV SYSROOT qtrpi-sysroot-minimal-latest.zip
RUN curl -sL --retry 10 --retry-delay 10 -o /tmp/$SYSROOT http://www.qtrpi.com/downloads/sysroot/$SYSROOT && unzip -qq /tmp/$SYSROOT -d /opt/qtrpi && rm -f /tmp/$SYSROOT
RUN ln -s /opt/qtrpi/raspbian/sysroot-minimal /opt/qtrpi/raspbian/sysroot


RUN git clone -q --depth 1 https://github.com/raspberrypi/tools.git /opt/qtrpi/raspi/tools


ENV QT_QMAKE_DIR /opt/qtrpi/raspi/qt5/bin
ENV RPI_TOOLS_DIR /opt/qtrpi/raspi/tools
ENV RPI_COMPILER gcc-linaro-arm-linux-gnueabihf-raspbian-x64
