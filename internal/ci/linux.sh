#!/bin/sh
set -ev

#replace gcc4 with gcc5
sudo add-apt-repository -y ppa:ubuntu-toolchain-r/test
sudo apt-get -qq update
sudo apt-get -y -qq install gcc-5 g++-5
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-5 90
sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-5 90

#download and install qt
curl -sL -o /tmp/qt-opensource-linux-x64-android-5.7.0.run https://download.qt.io/official_releases/qt/5.7/5.7.0/qt-opensource-linux-x64-android-5.7.0.run
chmod +x /tmp/qt-opensource-linux-x64-android-5.7.0.run
/tmp/qt-opensource-linux-x64-android-5.7.0.run --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
rm -f /tmp/qt-opensource-linux-x64-android-5.7.0.run

#download and install android sdk
curl -sL -o /tmp/android-sdk_r24.4.1-linux.tgz https://dl.google.com/android/android-sdk_r24.4.1-linux.tgz
tar -xzf /tmp/android-sdk_r24.4.1-linux.tgz -C /tmp
rm -f /tmp/android-sdk_r24.4.1-linux.tgz

#install deps for android sdk
echo "y" | /tmp/android-sdk-linux/tools/android -s update sdk -f -u -a -t 1,2,4,31

#download and install android ndk
curl -sL -o /tmp/android-ndk-r12b-linux-x86_64.zip https://dl.google.com/android/repository/android-ndk-r12b-linux-x86_64.zip
unzip -qq /tmp/android-ndk-r12b-linux-x86_64.zip -d /tmp
rm -f /tmp/android-ndk-r12b-linux-x86_64.zip

#download and install virtualbox
#sudo apt-get -y -qq install linux-headers-generic
#curl -sL -o /tmp/VirtualBox-5.1.6-110634-Linux_amd64.run http://download.virtualbox.org/virtualbox/5.1.6/VirtualBox-5.1.6-110634-Linux_amd64.run
#chmod +x /tmp/VirtualBox-5.1.6-110634-Linux_amd64.run
#sudo /tmp/VirtualBox-5.1.6-110634-Linux_amd64.run
#rm -f /tmp/VirtualBox-5.1.6-110634-Linux_amd64.run

#download and install sailfish sdk
#curl -sL -o /tmp/SailfishOSSDK-Beta-1608-Qt5-linux-64-offline.run https://releases.sailfishos.org/sdk/installers/1608/SailfishOSSDK-Beta-1608-Qt5-linux-64-offline.run
#chmod +x /tmp/SailfishOSSDK-Beta-1608-Qt5-linux-64-offline.run
#/tmp/SailfishOSSDK-Beta-1608-Qt5-linux-64-offline.run --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
#rm -f /tmp/SailfishOSSDK-Beta-1608-Qt5-linux-64-offline.run
