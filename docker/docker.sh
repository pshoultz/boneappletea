#!/bin/bash

value=$1

if [ "$value" == "build" ]; then
    cd ~/go/src/github.com/boneappletea/docker/client && docker image build -t bat-client-img . --no-cache &&
    cd ~/go/src/github.com/boneappletea/docker/api && docker image build -t bat-api-img . --no-cache &&
    cd ~/go/src/github.com/boneappletea/docker/compose && docker-compose up -d
fi

if [ "$value" == "destroy" ]; then
    cd ~/go/src/github.com/boneappletea/docker/compose && docker-compose down &&
    docker image rm bat-client-img &&
    docker image rm bat-api-img 
fi
