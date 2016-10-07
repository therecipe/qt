#!/bin/sh
set -ev

#check env
df -h

ls $HOME/*
du -sh $HOME/*

#needed for headless qt installation
export QT_QPA_PLATFORM=minimal

#replace gcc4 with gcc5
sudo add-apt-repository -y ppa:ubuntu-toolchain-r/test
sudo apt-get -qq update
sudo apt-get -y -qq install gcc-5 g++-5
sudo rm /usr/bin/gcc; sudo ln -s /usr/bin/gcc-5 /usr/bin/gcc
sudo rm /usr/bin/g++; sudo ln -s /usr/bin/g++-5 /usr/bin/g++

#download and install qt
QT=qt-opensource-linux-x64-android-5.7.0.run
curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT https://download.qt.io/official_releases/qt/5.7/5.7.0/$QT
chmod +x /tmp/$QT
/tmp/$QT --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
rm -f /tmp/$QT

#download and install android sdk
SDK=android-sdk_r24.4.1-linux.tgz
curl -sL --retry 10 --retry-delay 10 -o /tmp/$SDK https://dl.google.com/android/$SDK
tar -xzf /tmp/$SDK -C $HOME
rm -f /tmp/$SDK

#install deps for android sdk
echo "y" | $HOME/android-sdk-linux/tools/android -s update sdk -f -u -a -t 1,2,4,31

#download and install android ndk
NDK=android-ndk-r12b-linux-x86_64.zip
curl -sL --retry 10 --retry-delay 10 -o /tmp/$NDK https://dl.google.com/android/repository/$NDK
unzip -qq /tmp/$NDK -d $HOME
rm -f /tmp/$NDK

#prepare env
sudo chown $USER /usr/local/bin

#check env
df -h

ls $HOME/*
du -sh $HOME/*
