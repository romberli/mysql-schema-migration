#!/bin/bash

MSM_CMD=./bin/msm

output=$(${MSM_CMD} show migration --config=./config/msm.yaml)
if [[ $? -ne 0 ]]; then
    echo "failed to show migration."
    echo "${output}"
    exit 1
fi

echo "${output}"