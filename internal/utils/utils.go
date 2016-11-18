package utils

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
)

func Exists(name string) bool {
	var _, err = ioutil.ReadFile(name)
	return err == nil
}

func MakeFolder(dir string) error {
	var err = os.MkdirAll(dir, 0777)
	if err != nil {
		Log.WithError(err).Panicf("failed to create folder %v", dir)
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

func GetAbsPath(appPath string) string {
	var path, err = filepath.Abs(appPath)
	if err != nil {
		Log.WithError(err).Panicf("failed to get absolute path for %v", appPath)
	}
	return path
}

func GoQtPkgPath(s ...string) string {
	return filepath.Join(MustGoPath(), "src", "github.com", "therecipe", "qt", filepath.Join(s...))
}

func RunCmd(cmd *exec.Cmd, name string) string {
	fields := logrus.Fields{"func": "RunCmd", "name": name, "cmd": strings.Join(cmd.Args, "")}
        Log.WithFields(fields).Debug("Execute")
	var out, err = cmd.CombinedOutput()
	if err != nil {
		Log.WithError(err).WithFields(fields).WithField("output", out).Panic("failed to run command")
	}
	return string(out)
}

func RunCmdOptional(cmd *exec.Cmd, name string) string {
	var out, err = cmd.CombinedOutput()
	if err != nil {
		Log.WithError(err).Errorf("failed to %v\nerror: %s\ncmd: %v", name, out, cmd)
	}
	return string(out)
}
