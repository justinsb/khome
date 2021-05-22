.PHONY: all
all: .build/initramfs.cpio
	echo "Done"

.PHONY: .build/init
.build/init:
	mkdir -p .build/
	cd init; go build -o ../.build/init main.go

.build/initramfs.cpio: initramfs.list .build/init
	../kernel/linux/usr/gen_init_cpio initramfs.list > .build/initramfs.cpio
