#!/bin/sh
set -ev

#check env
df -h
diskutil list

ls $HOME/*
du -sh $HOME/*

#download and install qt
if [ "$IOS" == "true" ] || [ "$IOS_SIMULATOR" == "true" ]
then
QT=qt-opensource-mac-x64-android-ios-5.7.0
else
QT=qt-opensource-mac-x64-android-5.7.0
fi
curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT.dmg https://download.qt.io/official_releases/qt/5.7/5.7.0/$QT.dmg
hdiutil attach -noverify -noautofsck -quiet /tmp/$QT.dmg
/Volumes/$QT/$QT.app/Contents/MacOS/$QT --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
diskutil unmountDisk disk1
rm -f /tmp/$QT.dmg

if [ "$ANDROID" == "true" ]
then
#download and install android sdk
SDK=android-sdk_r24.4.1-macosx.zip
curl -sL --retry 10 --retry-delay 10 -o /tmp/$SDK https://dl.google.com/android/$SDK
unzip -qq /tmp/$SDK -d $HOME
rm -f /tmp/$SDK

#install deps for android sdk
$HOME/android-sdk-macosx/tools/android list sdk
echo "y" | $HOME/android-sdk-macosx/tools/android -s update sdk -f -u -t 1,2,3,5

#download and install android ndk
NDK=android-ndk-r12b-darwin-x86_64.zip
curl -sL --retry 10 --retry-delay 10 -o /tmp/$NDK https://dl.google.com/android/repository/$NDK
unzip -qq /tmp/$NDK -d $HOME
rm -f /tmp/$NDK
fi

#prepare env
sudo chown $USER /usr/local/bin
sudo chown $USER $GOROOT/pkg | true

#check env
df -h
diskutil list

ls $HOME/*
du -sh $HOME/*
