#!/bin/bash

./tools/geninitramfs.sh
TERM=linux sudo qemu-system-x86_64 -kernel /boot/vmlinuz-2.6.38-8-server -initrd ./initramfs -curses -net none -no-kvm
