#!/bin/bash

cd ${GOROOT}/src/pkg/net
make nuke
export CGO_ENABLED=0
make
make install
