package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/sys/unix"
)

func main() {
	fmt.Printf("env: %v\n", os.Environ())

	for {
		if err := mountProc(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to mount: %v\n", err)
		}

		tryDumpFile("/proc/cmdline")
		tryDumpFile("/proc/mounts")

		if err := readCmdline(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read cmdline: %v\n", err)
		}

		fmt.Fprintf(os.Stdout, "Hello world\n")
		time.Sleep(time.Second)
	}
}

func tryDumpFile(p string) {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to print file %q: %v\n", p, err)
	}
	fmt.Printf("%s: %q\n", p, string(b))
}

func readCmdline() error {
	cmdline, err := ioutil.ReadFile("/proc/cmdline")
	if err != nil {
		return fmt.Errorf("ReadFile(/proc/cmdline) failed: %w", err)
	}
	fmt.Printf("cmdline: %q\n", string(cmdline))

	return nil
}

func mountProc() error {
	_, err := os.Stat("/proc/self")
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("Stat(/proc/self) failed: %w", err)
	}

	if err := unix.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		return fmt.Errorf("unix.Mount(proc) failed: %w", err)
	}

	return nil
}
