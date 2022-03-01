#!/bin/bash
set -eo pipefail

go build -o apigen ./
echo "=== Type list ==="
./apigen print-type-list | tee snapshot-type-list.txt
echo "=== API codegen ==="
./apigen generate-api | tee snapshot-apigen.txt
echo "=== Messages list ==="
./apigen print-messages | tee snapshot-messages.txt
