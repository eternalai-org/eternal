#!/bin/bash

docker build -f build.Dockerfile --no-cache --output build .
chmod +x ./build/workersv