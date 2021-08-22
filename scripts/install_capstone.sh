#!/usr/bin/bash

CS_VERSION=4.0.2

cd /tmp

wget https://github.com/aquynh/capstone/archive/${CS_VERSION}.tar.gz
tar xzf ./${CS_VERSION}.tar.gz
rm ./${CS_VERSION}.tar.gz
mv ./capstone-${CS_VERSION} ./capstone
cd ./capstone
make
make install

cd $HOME
rm -rf /tmp/capstone