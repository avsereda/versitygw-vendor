#!/bin/sh

if [ -z "$VGW_ACCESS" ]; then
    echo >&2 "error: VGW_ACCESS= not specified"
    exit 1
fi

if [ -z "$VGW_SECRET" ]; then
    echo >&2 "error: VGW_SECRET= not specified"
    exit 1
fi

exec /versitygw --access "$VGW_ACCESS" --secret "$VGW_SECRET" --iam-dir /iam --webui 0.0.0.0:7071 posix /data "$@"