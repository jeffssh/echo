#!/usr/bin/env bash

docker build -t jeffssh/echo:latest .
docker push jeffssh/echo:latest
