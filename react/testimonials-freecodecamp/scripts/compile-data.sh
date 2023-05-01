#!/bin/sh

set -e

# Using yarn or npm, the env variable PWD is set to the project's root path
data_path=$PWD/src/assets/data

files=$(find $data_path -type f \( -iname "*.yaml" -o -iname "*.yml" \))

# check https://mikefarah.gitbook.io/yq/operators/multiply-merge to choose how to merge
# *d -> Deep Merge, the order is important
yq eval-all '. as $item ireduce ({}; . *d $item )' \
    $files --colors --output-format=json \
    | tee /dev/tty \
    | sed 's/\x1B\[[0-9;]\+[A-Za-z]//g' > $data_path/data.json
