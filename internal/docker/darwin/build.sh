#!/bin/bash

if [ ! -f MacOSX*.sdk.tar.xz ]; then
	GP=gen_sdk_package.sh && curl -sL --retry 10 --retry-delay 60 -O https://raw.githubusercontent.com/tpoechtrager/osxcross/master/tools/$GP && chmod +x $GP && ./$GP && rm $GP
fi

docker build -f Dockerfile -t therecipe/qt:darwin .
