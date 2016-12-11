package utils

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
)

func ExistsFile(name string) bool {
	var _, err = ioutil.ReadFile(name)
	return err == nil
}

func ExistsDir(name string) bool {
	var _, err = ioutil.ReadDir(name)
	return err == nil
}

func MkdirAll(dir string) error {
	var err = os.MkdirAll(dir, 0777)
	if err != nil {
		Log.WithError(err).Panicf("failed to create dir %v", dir)
	}
	return err
}

func RemoveAll(name string) error {
	var err = os.RemoveAll(name)
	if err != nil {
		Log.WithError(err).Panicf("failed to remove %v", name)
	}
	return err
}

func Save(name, data string) error {
	var err = ioutil.WriteFile(name, []byte(data), 0777)
	if err != nil {
		Log.WithError(err).Panicf("failed to save %v", name)
	}
	return err
}

func SaveBytes(name string, data []byte) error {
	var err = ioutil.WriteFile(name, data, 0777)
	if err != nil {
		Log.WithError(err).Panicf("failed to save %v", name)
	}
	return err
}

//TODO: export error
func Load(name string) string {
	var out, err = ioutil.ReadFile(name)
	if err != nil {
		Log.WithError(err).Errorf("failed to load %v", name)
	}
	return string(out)
}

func Abs(appPath string) (string, error) {
	var path, err = filepath.Abs(appPath)
	if err != nil {
		Log.WithError(err).Panicf("failed to get absolute path for %v", appPath)
	}
	return path, err
}

func GoQtPkgPath(s ...string) string {
	return filepath.Join(MustGoPath(), "src", "github.com", "therecipe", "qt", filepath.Join(s...))
}

func RunCmd(cmd *exec.Cmd, name string) string {
	fields := logrus.Fields{"func": "RunCmd", "name": name, "cmd": strings.Join(cmd.Args, " ")}
	Log.WithFields(fields).Debug("Execute")
	var out, err = cmd.CombinedOutput()
	if err != nil {
		Log.WithError(err).WithFields(fields).WithField("output", string(out)).Panic("failed to run command")
	}
	return string(out)
}

func RunCmdOptional(cmd *exec.Cmd, name string) string {
	fields := logrus.Fields{"func": "RunCmdOptional", "name": name, "cmd": strings.Join(cmd.Args, " ")}
	Log.WithFields(fields).Debug("Execute")
	var out, err = cmd.CombinedOutput()
	if err != nil {
		Log.WithError(err).WithFields(fields).WithField("output", string(out)).Error("failed to run command")
	}
	return string(out)
}
