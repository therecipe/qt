#!/bin/sh
set -e

#replace gcc4 with gcc5
sudo add-apt-repository -y ppa:ubuntu-toolchain-r/test
sudo apt-get update
sudo apt-get -y install gcc-5 g++-5
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-5 90
sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-5 90

#download and install qt
curl -sL -o /tmp/qt-opensource-linux-x64-android-5.7.0.run https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-linux-x64-android-5.7.0.run
chmod +x /tmp/qt-opensource-linux-x64-android-5.7.0.run && QT_QPA_PLATFORM=minimal /tmp/qt-opensource-linux-x64-android-5.7.0.run --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
