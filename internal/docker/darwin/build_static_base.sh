#!/bin/bash

set -ev

sudo xcode-select --switch /Applications/Xcode_9.4.1.app && xcrun -sdk macosx --show-sdk-path
git clone -q --depth 1 -b 5.13.2 --recursive https://code.qt.io/qt/qt5.git && cd qt5 && ./configure -prefix $HOME/Qt/5.13.0/clang_64 -opensource -confirm-license -release -nomake tests -nomake examples -qt-zlib -qt-libjpeg -qt-libpng -platform macx-clang -sysconfdir /Library/Preferences/Qt -dbus-runtime -static -optimize-size -skip qtwebengine -skip qtfeedback -skip qtpim -skip qtdocgallery -skip qtconnectivity -nomake tests -nomake examples && make -j $(sysctl -n hw.ncpu) && make install -j $(sysctl -n hw.ncpu) && cd .. && rm -rf qt5



QT_ROOT=$HOME
QT_VERSION=5.13.0
PREF=env_darwin_amd64_513

rm -rf ./$PREF
mkdir ./$PREF

rsync -avz $QT_ROOT/Qt/${QT_VERSION}/clang_64 ./$PREF/${QT_VERSION}/


rm -rf ./$PREF/${QT_VERSION}/clang_64/{doc,phrasebooks}
rm -rf ./$PREF/${QT_VERSION}/clang_64/lib/{cmake,pkgconfig,libQt5Bootstrap.a}


set +e
for v in *.jsc *.log *.pro *.pro.user *.qmake.stash *.qmlc .DS_Store *_debug* *.dSYM *.la; do
	find . -maxdepth 8 -name ${v} -exec rm -rf {} \;
done
set -e

mkdir -p ./$PREF/${QT_VERSION}/clang_64/_bin
for v in moc qmake qmlcachegen qmlimportscanner rcc uic; do
	mv ./$PREF/${QT_VERSION}/clang_64/bin/${v} ./$PREF/${QT_VERSION}/clang_64/_bin/
done
rm -rf ./$PREF/${QT_VERSION}/clang_64/bin && mv ./$PREF/${QT_VERSION}/clang_64/_bin ./$PREF/${QT_VERSION}/clang_64/bin

find ./$PREF/${QT_VERSION}/clang_64/bin -type f ! -name "qt.conf" -exec strip -x {} \;





mv $QT_ROOT/Qt $QT_ROOT/Qt_orig

go run ./patch_static_base.go



du -sh ./$PREF

#QT_STATIC=true QT_QMAKE_DIR=$(pwd)/$PREF/${QT_VERSION}/clang_64/bin $(go env GOPATH)/bin/qtsetup

zip -r ${PREF}.zip $PREF && rm -rf $PREF

du -sh ${PREF}.zip
