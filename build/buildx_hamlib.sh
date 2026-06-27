#!/bin/bash
set -e

HAMLIB_VERSION="4.7.1" #"master" #"Hamlib-4.7"
OUTPUT_DIR="./binaries"

echo "Compiling Hamlib ${HAMLIB_VERSION}..."
mkdir -p "${OUTPUT_DIR}/linux/amd64"
mkdir -p "${OUTPUT_DIR}/linux/arm64"

# If Docker buildx is not yet configured for multi-arch
docker buildx create --use --name hamlib_builder || true

echo "--> Building for Linux x86_64 (amd64)..."
docker buildx build \
    --platform linux/amd64 \
    --build-arg HAMLIB_TAG=${HAMLIB_VERSION} \
    --file Dockerfile.hamlib \
    --output type=local,dest=${OUTPUT_DIR}/linux/amd64 .

echo "--> Building for Raspberry Pi / ARM (arm64)..."
docker buildx build \
    --platform linux/arm64 \
    --build-arg HAMLIB_TAG=${HAMLIB_VERSION} \
    --file Dockerfile.hamlib \
    --output type=local,dest=${OUTPUT_DIR}/linux/arm64 .

echo "Done! The compiled binaries are located in: ${OUTPUT_DIR}"
ls -R ${OUTPUT_DIR}