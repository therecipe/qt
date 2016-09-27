package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func MakeFolder(dir string) {
	var err = os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println("file.MakeFolder", err)
	}
}

func RemoveAll(name string) {
	var err = os.RemoveAll(name)
	if err != nil {
		fmt.Println("file.RemoveAll", err)
	}
}

func Save(name, data string) {
	var err = ioutil.WriteFile(name, []byte(data), 0777)
	if err != nil {
		fmt.Println("file.Save", err)
	}
}

func SaveBytes(name string, data []byte) {
	var err = ioutil.WriteFile(name, data, 0777)
	if err != nil {
		fmt.Println("file.SaveBytes", err)
	}
}

func Load(name string) string {
	var b, err = ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("file.Load", err)
	}
	return string(b)
}

func GetAbsPath(appPath string) string {
	var wd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Clean(filepath.Join(wd, appPath))
}

func GetQtPkgPath(s ...string) string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "therecipe", "qt", filepath.Join(s...))
}

func Exists(name string) bool {
	var _, err = ioutil.ReadFile(name)
	return err == nil
}

func QtInstallDir() string {
	switch runtime.GOOS {
	case "darwin", "linux":
		{
			return QtInstallDirDarwinAndLinux()
		}

	default:
		{
			return QtInstallDirWindows()
		}
	}
}

func QtInstallDirDarwinAndLinux() string {
	if dir := os.Getenv("QT_INSTALL_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return "/usr/local/Qt5.7.0"
}

func QtInstallDirWindows() string {
	if dir := os.Getenv("QT_INSTALL_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	return "C:\\Qt\\Qt5.7.0"
}
