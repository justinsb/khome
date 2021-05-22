.PHONY: build
build:
	go get go.universe.tf/netboot/cmd/pixiecore

.PHONY: run
run: build
	sudo ~/go/bin/pixiecore boot ../test-qemu/tftp/bzImage ../test-qemu/tftp/initramfs.cpio 

