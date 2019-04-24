#!/bin/bash

docker save therecipe/qt:darwin | gzip -n > darwin.tar.gz
