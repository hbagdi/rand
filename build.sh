#!/bin/bash
# Build rand for various platforms and OSes

build () {
    if [ -z $1 ] || [ -z $2 ]
    then
        echo "Provide GOOS and GOARCH"
        exit 1
    fi
    echo "Building with GOOS=$1 and GOARCH=$2"
    GOOS=$1 GOARCH=$2 go build
    if [ $? -ne 0 ] 
    then
        exit $?
    fi
}

build linux amd64
#build linux 386
#build linux arm
#build linux arm64

build windows amd64
#build windows 386

build darwin amd64
#build darwin 386

#build freebsd amd64
#build freebsd 386
