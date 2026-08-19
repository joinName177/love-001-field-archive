#!/usr/bin/env sh
set -eu
platform=${1:?platform required}
docker buildx build --platform "$platform" --load -f benzhi.Dockerfile -t field-archive:benzhi .
