FROM ubuntu:16.04
LABEL maintainer therecipe

ARG SECRET

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl && apt-get -qq clean

RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "linux"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "android"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "js_base_1"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "js_base_2"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "js_base_3"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "js_base_4"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "js"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi_base"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi1"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi2"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "rpi3"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "sailfish"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "ubports_arm_vivid"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "ubports_64_vivid"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "ubports_arm_xenial"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "ubports_64_xenial"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_static_base_1"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_static_base_2"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_static_base_3"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_static_base_4"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_static_base_5"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_static_base_6"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_static_base_7"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_static"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_shared_base_1"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_shared_base_2"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_shared_base_3"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_shared_base_4"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_shared_base_5"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_shared_base_6"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_shared_base_7"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_32_shared"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_static_base_1"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_static_base_2"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_static_base_3"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_static_base_4"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_static_base_5"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_static_base_6"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_static_base_7"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_static"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_shared_base_1"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_shared_base_2"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_shared_base_3"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_shared_base_4"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_shared_base_5"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_shared_base_6"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
#RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_shared_base_7"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/
RUN sleep 120 && curl -H "Content-Type: application/json" --data '{"docker_tag": "windows_64_shared"}' -X POST https://registry.hub.docker.com/u/therecipe/qt/trigger/$SECRET/

RUN exit 1
