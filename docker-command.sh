#!/usr/bin/env bash

# docker build -t my-golang-app . && docker run --net="host" my-golang-app:latest
docker build -t wiki-races . && docker run -it -p 8080:80 wiki-races
