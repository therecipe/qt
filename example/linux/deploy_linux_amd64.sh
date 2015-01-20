export PATH=$PATH:/usr/local/go/bin

go build -o example ../

cp /usr/local/Qt5.4.0/5.4/gcc_64/lib/libQt5Core.so.5.4.0 ./libQt5Core.so.5
cp /usr/local/Qt5.4.0/5.4/gcc_64/lib/libQt5Gui.so.5.4.0 ./libQt5Gui.so.5
cp /usr/local/Qt5.4.0/5.4/gcc_64/lib/libQt5Widgets.so.5.4.0 ./libQt5Widgets.so.5
cp /usr/local/Qt5.4.0/5.4/gcc_64/lib/libQt5DBus.so.5.4.0 ./libQt5DBus.so.5
cp /usr/local/Qt5.4.0/5.4/gcc_64/lib/libicudata.so.53.1 ./libicudata.so.53
cp /usr/local/Qt5.4.0/5.4/gcc_64/lib/libicui18n.so.53.1 ./libicui18n.so.53
cp /usr/local/Qt5.4.0/5.4/gcc_64/lib/libicuuc.so.53.1 ./libicuuc.so.53

mkdir ./platforms
cp /usr/local/Qt5.4.0/5.4/gcc_64/plugins/platforms/libqxcb.so ./platforms/libqxcb.so

mkdir ./platformthemes
cp /usr/local/Qt5.4.0/5.4/gcc_64/plugins/platformthemes/libqgtk2.so ./platformthemes/libqgtk2.so

chmod 777 ./example.sh
./example.sh