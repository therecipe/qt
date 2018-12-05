#!/bin/bash

set -ev

mkdir -p ./darwin/Contents/Frameworks
curl -sL https://github.com/sparkle-project/Sparkle/releases/download/1.21.0/Sparkle-1.21.0.tar.bz2 | tar -C ./darwin/Contents/Frameworks -x Sparkle.framework
