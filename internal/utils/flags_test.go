package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendToFlag(t *testing.T) {

	flags := []string{"-p", "-pkgdir=/usr/lib/testdir", "-mod", "-modfile go.mod"}

	flags = AppendToFlag(flags, "-extldflags", "-v")
	require.Equal(t, "-extldflags=-v", flags[4], "Flag was not inserted")

	flags = AppendToFlag(flags, "-extldflags", "-Wl,soname,libname.so")
	require.Equal(t, "-extldflags=-v -Wl,soname,libname.so", flags[4], "Flag was not appended with new content")

	flags = AppendToFlag(flags, "-p", "test")
	require.Equal(t, "-p=test", flags[0], "Content was not appended to simple flag")
}
