#!/bin/sh
set -ev

#download and install qt
curl -sL --retry 3 -o /tmp/qt-opensource-mac-x64-android-ios-5.7.0.dmg https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-mac-x64-android-ios-5.7.0.dmg
hdiutil attach -noverify -noautofsck -quiet /tmp/qt-opensource-mac-x64-android-ios-5.7.0.dmg
/Volumes/qt-opensource-mac-x64-android-ios-5.7.0/qt-opensource-mac-x64-android-ios-5.7.0.app/Contents/MacOS/qt-opensource-mac-x64-android-ios-5.7.0 --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
diskutil unmountDisk disk1
rm -f /tmp/qt-opensource-mac-x64-android-ios-5.7.0.dmg

#download and install android sdk
curl -sL --retry 3 -o /tmp/android-sdk_r24.4.1-macosx.zip https://dl.google.com/android/android-sdk_r24.4.1-macosx.zip
unzip -qq /tmp/android-sdk_r24.4.1-macosx.zip -d /tmp
rm -f /tmp/android-sdk_r24.4.1-macosx.zip

#install deps for android sdk
echo "y" | /tmp/android-sdk-macosx/tools/android -s update sdk -f -u -a -t 1,2,4,31

#download and install android ndk
curl -sL --retry 3 -o /tmp/android-ndk-r12b-darwin-x86_64.zip https://dl.google.com/android/repository/android-ndk-r12b-darwin-x86_64.zip
unzip -qq /tmp/android-ndk-r12b-darwin-x86_64.zip -d /tmp
rm -f /tmp/android-ndk-r12b-darwin-x86_64.zip

#download and install virtualbox
curl -sL --retry 3 -o /tmp/VirtualBox-5.1.6-110634-OSX.dmg http://download.virtualbox.org/virtualbox/5.1.6/VirtualBox-5.1.6-110634-OSX.dmg
hdiutil attach -noverify -noautofsck -quiet /tmp/VirtualBox-5.1.6-110634-OSX.dmg
sudo installer -pkg /Volumes/VirtualBox/VirtualBox.pkg -target /
diskutil unmountDisk disk1
rm -f /tmp/VirtualBox-5.1.6-110634-OSX.dmg

#download and install sailfish sdk
curl -sL --retry 3 -o /tmp/SailfishOSSDK-Beta-1608-Qt5-mac-offline.dmg https://releases.sailfishos.org/sdk/installers/1608/SailfishOSSDK-Beta-1608-Qt5-mac-offline.dmg
hdiutil attach -noverify -noautofsck -quiet /tmp/SailfishOSSDK-Beta-1608-Qt5-mac-offline.dmg
/Volumes/SailfishOSSDK-mac-offline-160801/SailfishOSSDK-mac-offline-160801.app/Contents/MacOS/SailfishOSSDK-mac-offline-160801 --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
diskutil unmountDisk disk1
rm -f /tmp/SailfishOSSDK-Beta-1608-Qt5-mac-offline.dmg
