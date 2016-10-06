#!/bin/sh
set -ev

#
df -h
diskutil list

#download and install qt
if [ "$IOS" == "true" ]
then
QT=qt-opensource-mac-x64-android-ios-5.7.0
else
QT=qt-opensource-mac-x64-android-5.7.0
fi
curl -sL --retry 3 -o /tmp/$QT.dmg https://download.qt.io/official_releases/qt/5.7/5.7.0/$QT.dmg
hdiutil attach -noverify -noautofsck -quiet /tmp/$QT.dmg
/Volumes/$QT/$QT.app/Contents/MacOS/$QT --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
diskutil unmountDisk disk1
rm -f /tmp/$QT.dmg

if [ "$ANDROID" == "true" ]
then
#download and install android sdk
SDK=android-sdk_r24.4.1-macosx.zip
curl -sL --retry 3 -o /tmp/$SDK https://dl.google.com/android/$SDK
unzip -qq /tmp/$SDK -d $HOME
rm -f /tmp/$SDK

#install deps for android sdk
echo "y" | $HOME/android-sdk-macosx/tools/android -s update sdk -f -u -a -t 1,2,4,31

#download and install android ndk
NDK=android-ndk-r12b-darwin-x86_64.zip
curl -sL --retry 3 -o /tmp/$NDK https://dl.google.com/android/repository/$NDK
unzip -qq /tmp/$NDK -d $HOME
rm -f /tmp/$NDK
fi

if [ "$SAILFISH" == "true" ]
then
#download and install virtualbox
VBOX=VirtualBox-5.1.6-110634-OSX.dmg
curl -sL --retry 3 -o /tmp/$VBOX http://download.virtualbox.org/virtualbox/5.1.6/$VBOX
hdiutil attach -noverify -noautofsck -quiet /tmp/$VBOX
sudo installer -pkg /Volumes/VirtualBox/VirtualBox.pkg -target /
diskutil unmountDisk disk1
rm -f /tmp/$VBOX

#download and install sailfish sdk
SFDK=SailfishOSSDK-Beta-1608-Qt5-mac-offline.dmg
curl -sL --retry 3 -o /tmp/$SFDK https://releases.sailfishos.org/sdk/installers/1608/$SFDK
hdiutil attach -noverify -noautofsck -quiet /tmp/$SFDK
/Volumes/SailfishOSSDK-mac-offline-160801/SailfishOSSDK-mac-offline-160801.app/Contents/MacOS/SailfishOSSDK-mac-offline-160801 --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
diskutil unmountDisk disk1
rm -f /tmp/$SFDK
fi

#
df -h
diskutil list

ls $HOME
ls $HOME/*
ls $HOME/$SDK
ls $HOME/$NDK

du -sh $HOME
du -sh $HOME/*
du -sh $HOME/$SDK
du -sh $HOME/$NDK
