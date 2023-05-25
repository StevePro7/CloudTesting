#!/usr/bin/env sh
set -e
cd /tmp
git clone https://github.com/libbpf/libbpf
cd libbpf/src
make
sudo make install