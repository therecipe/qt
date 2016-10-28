package utils

import (
	"os"
	"path/filepath"
)

func RPI_TOOLS_DIR() string {
	if dir := os.Getenv("RPI_TOOLS_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join(os.Getenv("HOME"), "raspi/tools")
}

func RPI1_SYSROOT_DIR() string {
	if dir := os.Getenv("RPI1_SYSROOT_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join(os.Getenv("HOME"), "raspi/sysroot")
}

func RPI2_SYSROOT_DIR() string {
	if dir := os.Getenv("RPI2_SYSROOT_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join(os.Getenv("HOME"), "raspi/sysroot")
}

func RPI3_SYSROOT_DIR() string {
	if dir := os.Getenv("RPI3_SYSROOT_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return filepath.Join(os.Getenv("HOME"), "raspi/sysroot")
}
