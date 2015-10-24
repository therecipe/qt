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
	var deployPath string = path.Join(appPath, "deploy_darwin")
	var appName string = path.Base(appPath)

	utils.RemoveAll(deployPath)

	var out, err = exec.Command("go", "build", "-ldflags=\"-w\" \"-r=/usr/local/Qt5.5.1/5.5/clang_64/lib\"", "-o", path.Join(deployPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName)), fmt.Sprintf("%v/%v.go", appPath, appName)).CombinedOutput()
	if err != nil {
		fmt.Println("deploy.build", string(out), err)
	}

	utils.Save(path.Join(deployPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_sh", appName, appName)), macSH(appName))
	utils.Save(path.Join(deployPath, fmt.Sprintf("%v.app/Contents/MacOS/Info.plist", appName)), macPLIST(appName))

	var macdeploy = exec.Command("/usr/local/Qt5.5.1/5.5/clang_64/bin/macdeployqt", path.Join(deployPath, fmt.Sprintf("%v.app/", appName)), fmt.Sprintf("-qmldir=%v", path.Join(appPath, "qml")), "-always-overwrite")
	macdeploy.Dir = "/usr/local/Qt5.5.1/5.5/clang_64/bin/"
	out, err = macdeploy.CombinedOutput()
	if err != nil {
		fmt.Println("deploy.qtdeploy", string(out), err)
	}

	exec.Command("cp", "-R", path.Join(appPath, "qml/"), path.Join(deployPath, fmt.Sprintf("%v.app/Contents/MacOS/qml/", appName))).Run()

	exec.Command("mv", path.Join(deployPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName)), path.Join(deployPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_app", appName, appName))).Run()
	exec.Command("mv", path.Join(deployPath, fmt.Sprintf("%v.app/Contents/MacOS/%v_sh", appName, appName)), path.Join(deployPath, fmt.Sprintf("%v.app/Contents/MacOS/%v", appName, appName))).Run()

	exec.Command("open", path.Join(deployPath, fmt.Sprintf("%v.app/", appName))).Start()

}

func macSH(appName string) (o string) {
	o += "#!/bin/bash\n"
	o += "cd \"${0%/*}\"\n"
	o += fmt.Sprintf("./%v_app", appName)
	return
}

func macPLIST(appName string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>English</string>

	<key>CFBundleExecutable</key>
	<string>%v</string>

	<key>CFBundleIconFile</key>
	<string>%v.icns</string>

	<key>CFBundleIdentifier</key>
	<string>com.identifier.%v</string>

	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>

	<key>CFBundleName</key>
	<string>%v</string>

	<key>CFBundlePackageType</key>
	<string>APPL</string>

	<key>CFBundleShortVersionString</key>
	<string>1.0.0.0</string>

	<key>CFBundleVersion</key>
	<string>1</string>

	<key>NSHighResolutionCapable</key>
	<string>True</string>
</dict>
</plist>
`, appName, appName, appName, appName)
}
