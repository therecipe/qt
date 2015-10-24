package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func main() {

	var appPath = os.Args[1]
	var deployPath string = path.Join(appPath, "deploy_linux")
	var appName string = path.Base(appPath)
	var libraryPath string

	utils.RemoveAll(deployPath)

	var out, err = exec.Command("go", "build", "-ldflags=\"-s\" \"-w\"", "-o", path.Join(deployPath, appName), fmt.Sprintf("%v/%v.go", appPath, appName)).CombinedOutput()
	if err != nil {
		fmt.Println("deploy.build", string(out), err)
	}

	utils.Save(path.Join(deployPath, fmt.Sprintf("%v.sh", appName)), linuxSH(appName))

	exec.Command("cp", "-R", path.Join(appPath, "qml/"), path.Join(deployPath, "qml/")).Run()

	out, err = exec.Command("ldd", path.Join(deployPath, appName)).CombinedOutput()
	if err != nil {
		fmt.Println("deploy.ldd", string(out), err)
	} else {
		for _, dep := range strings.Split(string(out), "\n") {
			if strings.Contains(dep, "libQt5") || strings.Contains(dep, "libicu") {
				var libraryP, libName = path.Split(strings.Split(dep, " ")[2])
				libraryPath = libraryP
				exec.Command("cp", "-L", path.Join(libraryPath, libName), path.Join(deployPath, libName)).Run()
			}
		}

		exec.Command("cp", "-L", path.Join(libraryPath, "libQt5DBus.so.5"), path.Join(deployPath, "libQt5DBus.so.5")).Run()
		exec.Command("cp", "-L", path.Join(libraryPath, "libQt5XcbQpa.so.5"), path.Join(deployPath, "libQt5XcbQpa.so.5")).Run()
		exec.Command("cp", "-L", path.Join(libraryPath, "libQt5Quick.so.5"), path.Join(deployPath, "libQt5Quick.so.5")).Run()

		libraryPath = strings.TrimSuffix(libraryPath, "lib/")

		exec.Command("cp", "-R", path.Join(libraryPath, "qml/"), deployPath).Run()

		utils.MakeFolder(path.Join(deployPath, "platforms"))
		exec.Command("cp", "-L", path.Join(libraryPath, "plugins", "platforms", "libqxcb.so"), path.Join(deployPath, "platforms", "libqxcb.so")).Run()

		utils.MakeFolder(path.Join(deployPath, "platformthemes"))
		exec.Command("cp", "-L", path.Join(libraryPath, "plugins", "platformthemes", "libqgtk2.so"), path.Join(deployPath, "platformthemes", "libqgtk2.so")).Run()

		utils.MakeFolder(path.Join(deployPath, "xcbglintegrations"))
		exec.Command("cp", "-L", path.Join(libraryPath, "plugins", "xcbglintegrations", "libqxcb-glx-integration.so"), path.Join(deployPath, "xcbglintegrations", "libqxcb-glx-integration.so")).Run()

	}

	var cmd = exec.Command(path.Join(deployPath, fmt.Sprintf("%v.sh", appName)))
	cmd.Dir = deployPath
	cmd.Start()
}

func linuxSH(appName string) (o string) {

	o += "#!/bin/sh\n"
	o += "appname=`basename $0 | sed s,\\.sh$,,`\n\n"
	o += "dirname=`dirname $0`\n"
	o += "tmp=\"${dirname#?}\"\n\n"
	o += "if [ \"${dirname%$tmp}\" != \"/\" ]; then\n"
	o += "dirname=$PWD/$dirname\n"
	o += "fi\n"

	o += "LD_LIBRARY_PATH=$dirname\n"
	o += "export LD_LIBRARY_PATH\n"

	o += "QML_IMPORT_PATH=$dirname/\"qml\"\n"
	o += "export QML_IMPORT_PATH\n"

	o += "QML2_IMPORT_PATH=$dirname/\"qml\"\n"
	o += "export QML2_IMPORT_PATH\n"

	o += "$dirname/$appname \"$@\"\n"

	return
}
