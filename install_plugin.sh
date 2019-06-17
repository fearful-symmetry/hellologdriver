#!/bin/bash

docker build -t rootfsimage .
id=$(docker create rootfsimage true)
rm -rf rootfs
mkdir rootfs
docker export "$id" | tar -x -C rootfs

docker rm -vf "$id"

docker plugin create ossifrage/hellologdriver:0.0.1 .
docker plugin enable ossifrage/hellologdriver:0.0.1