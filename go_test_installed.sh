#!/bin/sh
LD_LIBRARY_PATH=/usr/local/lib CGO_LDFLAGS='-Wl,-rpath,/usr/local/lib' go test . ./types/... ./errors ./utils ./config ./apis/... ./service/... ./tests
