#!/bin/bash

echo "building docker image (blue)..."
docker build \
        --build-arg ARG_HTML_REFRESH_BGCOLOUR="blue" \
        --build-arg ARG_HTML_REFRESH_FGCOLOUR="white" \
        -t stuartleeks/htmlrefresh:blue .

echo " "
echo "building docker image (green)..."
docker build \
        --build-arg ARG_HTML_REFRESH_BGCOLOUR="green" \
        --build-arg ARG_HTML_REFRESH_FGCOLOUR="white" \
        -t stuartleeks/htmlrefresh:green .


