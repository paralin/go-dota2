#!/bin/bash
set -eo pipefail

REPO_ROOT=$(git rev-parse --show-toplevel)
GAME_DIR="Protobufs"
GAME_PATH="${REPO_ROOT}/generator/${GAME_DIR}"

cd ${REPO_ROOT}/generator
git submodule update --init ${GAME_DIR}

WORK_DIR=`mktemp -d`
# deletes the temp directory
function cleanup {
    sync || true
    if [ -d "$WORK_DIR" ]; then
        rm -rf "$WORK_DIR" || true
    fi
}
trap cleanup EXIT

cd ${GAME_PATH}/dota2
mkdir -p ${WORK_DIR}/orig ${WORK_DIR}/protos
cp \
    ./dota_gcmessages_*.proto \
    ./dota_client_enums.proto \
    ./network_connection.proto \
    ./base_gcmessages.proto \
    ./econ_*.proto \
    ./dota_match_metadata.proto \
    ./dota_shared_enums.proto \
    ./gcsdk_gcmessages.proto \
    ./steammessages.proto \
    ./valveextensions.proto \
    ./gcsystemmsgs.proto \
    ${WORK_DIR}/orig/

cd ${WORK_DIR}
# Add package lines to each protobuf file.
for f in ${WORK_DIR}/orig/*.proto ; do
    fname=$(basename $f)
    printf 'syntax = "proto2";\npackage protocol;\n\n' |\
        cat - $f |\
        sed -e "s/optional \./optional /g" \
            -e "s/required \./required /g" \
            -e "s/repeated \./repeated /g" \
            -e "s/\t\./\t/g" >\
            ${WORK_DIR}/protos/${fname}
done

# Generate protobufs
cd ${WORK_DIR}/protos
protoc -I $(pwd) --go_out=. $(pwd)/*.proto

# Move final files out.
rsync -rv --delete $(pwd)/ ${REPO_ROOT}/protocol/

