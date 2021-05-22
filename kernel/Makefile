.PHONY: all
all: download config build

.PHONY: download
download:
	git clone https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git -b v5.12 --depth=1

.PHONY: config
config:
	cd linux; make ARCH=x86_64 defconfig

.PHONY: build
build:
	cd linux; make -j32 bzImage
