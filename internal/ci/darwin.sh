#!/bin/sh
set -ev

#download and install qt
curl -sL -o /tmp/qt-opensource-mac-x64-android-ios-5.7.0.dmg https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-mac-x64-android-ios-5.7.0.dmg
open /tmp/qt-opensource-mac-x64-android-ios-5.7.0.dmg
QT_QPA_PLATFORM=minimal /Volumes/qt-opensource-mac-x64-android-ios-5.7.0/qt-opensource-mac-x64-android-ios-5.7.0.app/Contents/MacOS/qt-opensource-mac-x64-android-ios-5.7.0 --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
diskutil unmountDisk disk2
