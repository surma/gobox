#!/bin/bash

./tools/geninitramfs.sh
#"virtio", "i82551", "i82557b", "i82559er", "ne2k_pci", "ne2k_isa", "pcnet", "rtl8139", "e1000", "smc91c111", "lance" and "mcf_fec".
#ne2k_pci,i82551,i82557b,i82559er,rtl8139,e1000,pcnet,virtio
TERM=linux sudo qemu-system-x86_64 -kernel "/boot/vmlinuz-2.6.38-8-server" -append "console=ttyS0" -initrd "./initramfs" -nographic -net nic,vlan=0,model=virtio -net tap,vlan=0,name=tap0,script="tools/qemu/up.sh",downscript="tools/qemu/down.sh" -no-kvm

