#!/bin/bash


sudo apt update && sudo apt install -y build-essential libtool autoconf automake pkg-config git
git clone --depth 1 --branch Hamlib-4.7 https://github.com/Hamlib/Hamlib.git /tmp/hamlib-build
cd /tmp/hamlib-build

./bootstrap
./configure --prefix=/tmp/hamlib-static --disable-shared --enable-static

make -j4
#make install






HAMLIB_VERSION="4.7.2"
INSTALL_DIR="/tmp/hamlib"
mkdir -p "$INSTALL_DIR"

echo "Downloading fixed Hamlib version $HAMLIB_VERSION from https://github.com/Hamlib/Hamlib/releases/download/${HAMLIB_VERSION}/hamlib-${HAMLIB_VERSION}.tar.gz..."

# curl -sL "https://downloads.sourceforge.net/project/hamlib/hamlib/${HAMLIB_VERSION}/hamlib-linux-x86_64-${HAMLIB_VERSION}.tar.gz" # | tar -xz -C "$INSTALL_DIR" --strip-components=1
curl -sL "https://github.com/Hamlib/Hamlib/releases/download/${HAMLIB_VERSION}/hamlib-${HAMLIB_VERSION}.tar.gz" | tar -xz -C "$INSTALL_DIR" --strip-components=1




