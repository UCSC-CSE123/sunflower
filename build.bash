#! /usr/bin/env bash

OSES="windows linux darwin"

build() {
  for OS in $OSES; do
    GOOS=$OS GOARCH=amd64 CGO_ENABLED=0 go build -o sunflower-$OS-amd64 -ldflags '-extldflags "-f no-PIC -static"' -tags 'osusergo netgo static_build'
    zip sunflower-$OS-amd64.zip sunflower-$OS-amd64
    rm sunflower-$OS-amd64
  done
}

clean() {
  go clean
  rm -f sunflower-*
}

if [ "$1" = "clean" ]; then
  clean
  exit
fi

if [ "$1" = "build" ]; then
  build
  exit
fi

echo "Invalid command. Usage: build.bash [clean | build]"
