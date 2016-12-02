package utils

import (
	"testing"
	"io/ioutil"
	"os"
	"path/filepath"

	assert "github.com/stretchr/testify/require"
	"github.com/Sirupsen/logrus"
	"sort"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	Log.Level  = logrus.DebugLevel
}

func TestWalkFilterBlacklist(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "TestWalkFilterBlacklist")
	assert.NoError(t, err)
	assert.NotEmpty(t, tempDir)
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()

	dummyData := []byte{0, 1, 2}

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

	var whiteListed []string
	err = filepath.Walk(tempDir, WalkFilterBlacklist(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			t.Logf("%s is blacklisted", path)
		} else {
			relPath, relErr := filepath.Rel(tempDir, path)
			assert.NoError(t, relErr)
			whiteListed = append(whiteListed, relPath)
		}
		return nil
	}))
	assert.NoError(t, err)
	assert.Len(t, whiteListed, 4)
	sort.Strings(whiteListed)
	assert.Equal(t, whiteListed[0], ".")
	assert.Equal(t, whiteListed[1], "whiteListedDir")
	assert.Equal(t, whiteListed[2], "whiteListedDir/whiteListedSubFilename")
	assert.Equal(t, whiteListed[3], "whiteListedFile.dat")
}
