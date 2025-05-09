#!/usr/bin/env bash
set -euo pipefail

GO_VERSION=$(go version | cut -d' ' -f3-)
VERSION="${VERSION:-v0.0.1-alpha}"
OUT_DIR="dist"

LDFLAGS=(
    "-w" "-s"
    "-X main.ActualVersion=${VERSION}"
    "-extldflags=-static"
)

printf "\e[32m→\e[0m Go version: ${GO_VERSION}\r\n"
printf "\e[32m→\e[0m Building version: ${VERSION}\r\n"

if [ ! -d "$OUT_DIR" ]; then
    mkdir -p ${OUT_DIR}
else
    rm -rf ${OUT_DIR}
fi

go build \
    -ldflags="${LDFLAGS[*]}" \
    -o ${OUT_DIR}/mvs \
    .

printf "\e[32m→\e[0m Done build process!\r\n"
