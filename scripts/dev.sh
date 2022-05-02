#!/bin/bash
set -e
source scripts/build-config.sh

mkdir -p web/dist
echo "<html></html>" > web/dist/index.html

docker-compose up --build -V
