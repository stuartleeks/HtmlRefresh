#!/bin/bash

echo "docker-build.sh starting..."

export DOTNET_CLI_TELEMETRY_OPTOUT=1

echo "restoring packages..."
dotnet restore HtmlRefresh/HtmlRefresh.csproj 

echo "building project..."
dotnet build HtmlRefresh/HtmlRefresh.csproj 

echo "publishing project..."
rm -rf $(pwd)/publish
dotnet publish $(pwd)/HtmlRefresh/HtmlRefresh.csproj -o $(pwd)/publish/HtmlRefresh

echo "building docker image..."
docker build \
        --build-arg ARG_HTML_REFRESH_BGCOLOUR="blue" \
        --build-arg ARG_HTML_REFRESH_FGCOLOUR="white" \
        -t stuartleeks/htmlrefresh:blue .
docker build \
        --build-arg ARG_HTML_REFRESH_BGCOLOUR="green" \
        --build-arg ARG_HTML_REFRESH_FGCOLOUR="white" \
        -t stuartleeks/htmlrefresh:green .


