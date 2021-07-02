#!/usr/bin/env bash

docker build -t my-golang-app . && docker run --net="host" my-golang-app:latest
