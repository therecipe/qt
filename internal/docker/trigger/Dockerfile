FROM ubuntu:14.04

ARG SECRET

RUN apt-get -qq update && apt-get -y -qq install curl && apt-get -qq clean

ADD trigger.sh /
