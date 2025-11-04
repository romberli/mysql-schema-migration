#!/bin/bash

MSM_CMD=./bin/msm

output=$(${MSM_CMD} show diff --config=./config/msm.yaml)
if [[ $? -ne 0 ]]; then
    echo "failed to show diff."
    echo "${output}"
    exit 1
fi

echo "${output}"