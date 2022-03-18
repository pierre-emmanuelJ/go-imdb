#!/bin/bash

docker run -u "$UID:$GID" --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:latest generate \
    -i /local/swagger.json \
    -g go \
    -o /local/ \
    --additional-properties=packageName=client

mv ./*.go ./client
