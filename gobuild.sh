#!/bin/bash

set -euo pipefail

export GOPATH="$(pwd)/.gobuild"
SRCDIR="${GOPATH}/src/github.com/audy/deskclean"

[ -d ${GOPATH} ] && rm -rf ${GOPATH}
mkdir -p ${GOPATH}/{src,pkg,bin}
mkdir -p ${SRCDIR}
cp deskclean.go ${SRCDIR}

echo ${GOPATH}
cd ${SRCDIR}
go get .
go install .
