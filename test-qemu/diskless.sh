#!/bin/bash

TFTPROOT=`pwd`/tftp

qemu-system-x86_64 -nographic -serial mon:stdio \
  -netdev user,id=net0,net=192.168.76.0/24,dhcpstart=192.168.76.9,tftp=$TFTPROOT,bootfile=/ipxe-script \
  -device e1000,netdev=net0 \
  -m 512 --enable-kvm
