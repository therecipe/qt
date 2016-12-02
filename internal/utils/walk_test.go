package utils

import (
	"testing"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	assert "github.com/stretchr/testify/require"
	"github.com/Sirupsen/logrus"
)

var 	dummyData = []byte{0, 1, 2}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	Log.Level  = logrus.DebugLevel
}

type walkResult struct {
	output []string
	root string
}

func (w *walkResult) accumulate(path string, info os.FileInfo, err error) error {
	if err == nil {
		relPath, relErr := filepath.Rel(w.root, path)
		if relErr != nil {
			return relErr
		}
		w.output = append(w.output, relPath)
	}
	return err
}

func (w *walkResult) sorted() []string {
	output := w.output
	sort.Strings(output)
	return output
}

func newWalkResult(root string) *walkResult {
	return &walkResult{root: root}
}

func TestWalkFilterBlacklist(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "TestWalkFilterBlacklist")
	assert.NoError(t, err)
	assert.NotEmpty(t, tempDir)
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()

	blackListedFilename := filepath.Join(tempDir, "ios")
	assert.NoError(t, ioutil.WriteFile(blackListedFilename, dummyData, 0644))

	blackListedDir := filepath.Join(tempDir, ".git")
	assert.NoError(t, os.Mkdir(blackListedDir, 0755))
	blackListedSubFilename := filepath.Join(blackListedDir, "config")
	assert.NoError(t, ioutil.WriteFile(blackListedSubFilename, dummyData, 0644))

	whiteListedFilename := filepath.Join(tempDir, "whiteListedFile.dat")
	assert.NoError(t, ioutil.WriteFile(whiteListedFilename, dummyData, 0644))

	whiteListedDirectory := filepath.Join(tempDir, "whiteListedDir")
	assert.NoError(t, os.Mkdir(whiteListedDirectory, 0755))
	whiteListedSubFilename := filepath.Join(whiteListedDirectory, "whiteListedSubFilename")
	assert.NoError(t, ioutil.WriteFile(whiteListedSubFilename, dummyData, 0644))

	result := newWalkResult(tempDir)
	assert.NoError(t, filepath.Walk(tempDir, WalkFilterBlacklist(tempDir, result.accumulate)))
	output := result.sorted()
	assert.Len(t, output, 4)
	assert.Equal(t, output[0], ".")
	assert.Equal(t, output[1], "whiteListedDir")
	assert.Equal(t, output[2], "whiteListedDir/whiteListedSubFilename")
	assert.Equal(t, output[3], "whiteListedFile.dat")
}

func TestWalkOnlyDirectory(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "TestWalkOnlyDirectory")
	assert.NoError(t, err)
	assert.NotEmpty(t, tempDir)
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()

	file := filepath.Join(tempDir, "file")
	assert.NoError(t, ioutil.WriteFile(file, dummyData, 0644))

	dir := filepath.Join(tempDir, "dir")
	assert.NoError(t, os.Mkdir(dir, 0755))
	subFile := filepath.Join(dir, "subfile")
	assert.NoError(t, ioutil.WriteFile(subFile, dummyData, 0644))

	result := newWalkResult(tempDir)
	assert.NoError(t, filepath.Walk(tempDir, WalkOnlyDirectory(result.accumulate)))
	output := result.sorted()
	assert.Len(t, output, 2)
	assert.Equal(t, output[0], ".")
	assert.Equal(t, output[1], "dir")
}
