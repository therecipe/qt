#!/bin/bash

set -ev

###

curl -sL https://github.com/vslavik/winsparkle/releases/download/v0.6.0/WinSparkle-0.6.0.zip | tar -x

openssl dsaparam 4096 > dsaparam.pem

openssl gendsa dsaparam.pem -out dsa_priv.pem
rm -rf dsaparam.pem
openssl dsa -in dsa_priv.pem -pubout -out dsa_pub.pem

###

qtdeploy -docker build windows_32_shared
cp ./WinSparkle*/Release/WinSparkle.dll ./deploy/windows/
mv ./deploy/windows ./deploy/windows_32

qtdeploy -docker build windows_64_shared
cp ./WinSparkle*/x64/Release/WinSparkle.dll ./deploy/windows/
mv ./deploy/windows ./deploy/windows_64
