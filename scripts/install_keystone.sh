#!/usr/bin/bash

KS_VERSION=0.9.2

cd /tmp

wget https://github.com/keystone-engine/keystone/archive/${KS_VERSION}.tar.gz
tar xzf ./${KS_VERSION}.tar.gz
rm ./${KS_VERSION}.tar.gz
mv ./keystone-${KS_VERSION} ./keystone
cd ./keystone
mkdir build
cd build
cmake -DCMAKE_BUILD_TYPE=Release -D BUILD_LIBS_ONLY=1 -DBUILD_SHARED_LIBS=OFF -G "Unix Makefiles" ..
make -j8
make install

cd $HOME
rm -rf /tmp/keystone

# Keystone is installed in /usr/local, depending on your distribution (eg. Ubuntu) you might need to add /usr/local/lib to /etc/ld.so.conf. 
if [[ "$(awk -F= '/^ID=/{print $2}' /etc/os-release)" == "ubuntu" ]]; then
  ldconfig
fi