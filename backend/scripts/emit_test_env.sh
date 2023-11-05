#!/usr/bin/env bash

docker build -f Dockerfile.test . -t lastdisco/backend:latest
docker run --rm --net=host -p 3000:3000 lastdisco/backend:latest