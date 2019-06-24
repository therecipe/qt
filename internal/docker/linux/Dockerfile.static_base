FROM ubuntu:16.04 as base

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates git pkg-config

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install build-essential libglib2.0-dev libglu1-mesa-dev libpulse-dev \
	&& apt-get --no-install-recommends -qq -y install fontconfig libasound2 libegl1-mesa libnss3 libpci3 libxcomposite1 libxcursor1 libxi6 libxrandr2 libxtst6
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install libdbus-1-dev libssl-dev libxkbcommon-dev

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install libfontconfig1-dev libfreetype6-dev libx11-dev libxext-dev libxfixes-dev libxi-dev libxrender-dev libxcb1-dev libx11-xcb-dev libxcb-glx0-dev
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install libxcb-keysyms1-dev libxcb-image0-dev libxcb-shm0-dev libxcb-icccm4-dev libxcb-sync0-dev libxcb-xfixes0-dev libxcb-shape0-dev libxcb-randr0-dev libxcb-render-util0-dev libxkbcommon-dev

RUN git clone -q --depth 1 -b 5.12.4 --recursive https://code.qt.io/qt/qt5.git /opt/qt5
RUN cd /opt/qt5 && ./configure -prefix /opt/Qt/5.12.0/gcc_64 -confirm-license -opensource -static -qt-zlib -qt-libpng -qt-libjpeg -qt-xcb -sysconfdir /etc/xdg -dbus-runtime -openssl-runtime -opengl -optimize-size -skip qtwebengine -skip qtfeedback -nomake tests -nomake examples && make -j $(grep -c ^processor /proc/cpuinfo) && make install -j $(grep -c ^processor /proc/cpuinfo)


FROM ubuntu:16.04
LABEL maintainer therecipe

COPY --from=base /opt/Qt/5.12.0 /opt/Qt/5.12.0
