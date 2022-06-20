#!/bin/sh

CONTAINER_REPOSITORY="docker.io"
CONTAINER_IMAGE="${CONTAINER_REPOSITORY}/vranystepan/workshop-app"

# build the first variant
docker build \
    -t "${CONTAINER_IMAGE}:working" \
    --platform linux/amd64 \
    --build-arg ARG_CONFIG_BREAK=false \
    --build-arg ARG_CONFIG_SLEEP=false \
    .

# build the first variant
docker build \
    -t "${CONTAINER_IMAGE}:breaking" \
    --platform linux/amd64 \
    --build-arg ARG_CONFIG_BREAK=true \
    --build-arg ARG_CONFIG_SLEEP=false \
    .

docker build \
    -t "${CONTAINER_IMAGE}:sleeping" \
    --platform linux/amd64 \
    --build-arg ARG_CONFIG_BREAK=false \
    --build-arg ARG_CONFIG_SLEEP=true \
    .

# push all the variants to ECR
docker push "${CONTAINER_IMAGE}:working"
docker push "${CONTAINER_IMAGE}:breaking"
docker push "${CONTAINER_IMAGE}:sleeping"
