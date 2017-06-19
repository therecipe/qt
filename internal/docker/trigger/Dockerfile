FROM ubuntu:14.04

ARG SECRET

RUN apt-get -qq update && apt-get -y -qq install curl && apt-get -qq clean

RUN sleep 65
RUN curl -H "Content-Type: application/json" --data '{"docker_tag": "linux"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 65
RUN curl -H "Content-Type: application/json" --data '{"docker_tag": "android"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 65
RUN curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_shared"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 65
RUN curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_static"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 65
RUN curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_shared"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 65
RUN curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_static"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 65
RUN curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 65
RUN curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi1"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 65
RUN curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi2"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 65
RUN curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi3"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 65

RUN exit 1
