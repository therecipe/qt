#!/bin/bash

curl -H "Content-Type: application/json" --data '{"docker_tag": "linux"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
curl -H "Content-Type: application/json" --data '{"docker_tag": "android"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_shared"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_static"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_shared"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_static"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi1"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi2"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi3"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/

exit 1
