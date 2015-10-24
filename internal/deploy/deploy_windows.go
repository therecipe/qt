package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/therecipe/qt/internal/utils"
)

func main() {

	var appPath = os.Args[1]
	var deployPath string = path.Join(appPath, "deploy_windows")
	var appName string = path.Base(appPath)

	utils.RemoveAll(deployPath)

	var out, err = exec.Command("go", "build", "-ldflags=\"-H=windowsgui\"", "-o", path.Join(deployPath, appName), fmt.Sprintf("%v/%v.go", appPath, appName)).CombinedOutput()
	if err != nil {
		fmt.Println("deploy.build", string(out), err)
	}

	out, err = exec.Command("windeployqt", path.Join(deployPath, appName), fmt.Sprintf("-qmldir=%v", path.Join(appPath, "qml")), "--force").CombinedOutput()
	if err != nil {
		fmt.Println("deploy.qtdeploy", string(out), err)
	}

	exec.Command("cp", "-R", path.Join(appPath, "qml/"), path.Join(deployPath, "qml/")).Start()

	exec.Command(path.Join(deployPath, appName))
}
