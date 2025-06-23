#!/bin/bash
set -eo pipefail

REPO_ROOT=$(git rev-parse --show-toplevel)
REPO_PROTOS=${REPO_ROOT}/protos
GAME_DIR="Protobufs"
GAME_PATH="${REPO_ROOT}/generator/${GAME_DIR}"

cd ${REPO_ROOT}/generator
git submodule update --init ${GAME_DIR}

WORK_DIR=`mktemp -d`
echo "Using working directory: ${WORK_DIR}"
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
    ./gcsdk_gcmessages.proto \
    ./econ_*.proto \
    ./dota_match_metadata.proto \
    ./dota_shared_enums.proto \
    ./steammessages.proto \
    ./steammessages_unified_base.steamworkssdk.proto \
    ./steammessages_steamlearn.steamworkssdk.proto \
    ./valveextensions.proto \
    ./gcsystemmsgs.proto \
    ${WORK_DIR}/orig/

mkdir -p ${WORK_DIR}/orig/google/protobuf
cp -ra ${GAME_PATH}/google/protobuf/. ${WORK_DIR}/orig/google/protobuf/

cd ${WORK_DIR}
# Add valve_extensions.proto
# cp ${REPO_PROTOS}/valve_extensions.proto ${WORK_DIR}/orig/
# Add package lines to each protobuf file.
for f in ${WORK_DIR}/orig/*.proto ; do
    fname=$(basename $f)
    printf 'syntax = "proto2";\npackage protocol;\n\n' |\
        cat - $f |\
        sed -e "s/optional \./optional /g" \
            -e "s/required \./required /g" \
            -e "s/repeated \./repeated /g" \
            -e "s#google/protobuf/valve_extensions.proto#valveextensions.proto#g" \
            -e "s/\t\./\t/g" >\
            ${WORK_DIR}/protos/${fname}
done

# Fixup some issues in the protos
sed -i -e "s/(.CMsgSteamLearn/(CMsgSteamLearn/g" ${WORK_DIR}/protos/steammessages_steamlearn.steamworkssdk.proto

# Remove the problematic CDotaMsgStructuredTooltipProperties message block
sed -i -e '/^message CDotaMsgStructuredTooltipProperties {/,/^}$/d' ${WORK_DIR}/protos/dota_gcmessages_common.proto

# Move final files out.
rsync -rv --delete $(pwd)/protos/ ${REPO_ROOT}/protocol/
echo "package protocol" > ${REPO_ROOT}/protocol/protocol.go
