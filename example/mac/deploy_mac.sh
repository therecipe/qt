go build -o example

mkdir example.app
mkdir example.app/Contents
mkdir example.app/Contents/MacOS
mkdir example.app/Contents/Resources

mv example example.app/Contents/MacOS
cp example_sh example.app/Contents/MacOS
cp Info.plist example.app/Contents

chmod 777 example.app/Contents/MacOS/example
chmod 777 example.app/Contents/MacOS/example_sh

/usr/local/qt/5.4/clang_64/bin/macdeployqt example.app/ -always-overwrite

mv example.app/Contents/MacOS/example example.app/Contents/MacOS/example_app
mv example.app/Contents/MacOS/example_sh example.app/Contents/MacOS/example

open example.app