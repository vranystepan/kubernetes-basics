#!/bin/sh

AWS_ACCOUNT_ID="$(aws sts get-caller-identity | jq -r '.Account')"
CONTAINER_REPOSITORY="${AWS_ACCOUNT_ID}.dkr.ecr.eu-west-1.amazonaws.com"
CONTAINER_IMAGE="${CONTAINER_REPOSITORY}/training/application"

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
