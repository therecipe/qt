#!/bin/bash
set -ev

#sudo pkg ins -y kde5 sddm xorg && sudo sysrc sddm_enable=YES && sudo service sddm start

sudo pkg ins -y xfce slim xorg && echo "exec xfce4-session" > $HOME/.xinitrc && \
sudo sysrc dbus_enable=YES && sudo sysrc hald_enable=YES && sudo sysrc slim_enable=YES && \
sudo service dbus start && sudo service hald start && sudo service slim start

sudo pkg ins -y git go pkgconf qt5
go get -v github.com/therecipe/qt/cmd/...
QT_PKG_CONFIG=true $(go env GOPATH)/bin/qtsetup test

exit 0