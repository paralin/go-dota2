#!/bin/bash
set -eo pipefail

PERSIST_TEMPDIR="yes"

GAMETRACKING_REPO="https://github.com/SteamDatabase/GameTracking-Dota2/archive/master/gametracking-dota2.tar.gz"

if [ -n "$PERSIST_TEMPDIR" ]; then
    WORK_DIR=/tmp/dota-proto-update
    mkdir -p ${WORK_DIR}
else
    WORK_DIR=`mktemp -d`
    # deletes the temp directory
    function cleanup {
        sync || true
        if [ -d "$WORK_DIR" ]; then
            rm -rf "$WORK_DIR" || true
        fi
    }
    trap cleanup EXIT
fi

cd ${WORK_DIR}
if [ ! -d Protobufs ]; then
    echo "Pulling ${GAMETRACKING_REPO}"
    curl -L ${GAMETRACKING_REPO} | tar --strip-components=1 -zxf-
fi
cd -
ln -fs ${WORK_DIR}/Protobufs ./Protobufs
go build -v
./generator proto
rm Protobufs

