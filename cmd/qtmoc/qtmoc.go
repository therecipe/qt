package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/moc"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	if utils.QT_QMAKE_CGO() == true {
		qmake_main()
		return
	}
	if cmd.ParseFlags() {
		flag.PrintDefaults()
		return
	}

	var (
		appPath     string
		buildDocker bool
		fields      = logrus.Fields{"func": "main"}
	)
	defaultAppPath, err := os.Getwd()
	if err != nil {
		utils.Log.WithFields(fields).Fatal("Can't get current working directory")
	}

	switch flag.NArg() {
	case 0:
		appPath = defaultAppPath
	case 1:
		appPath = flag.Arg(0)
	case 2:
		appPath = flag.Arg(0)
		buildDocker = true
	default:
		fmt.Println("Only specify one path")
		fmt.Printf("%s path/to/app [ docker ]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	// validate that path is readable and a directory
	if !filepath.IsAbs(appPath) {
		appPath, _ = filepath.Abs(appPath)
	}
	fields["app_path"] = appPath

	appPathInfo, err := os.Stat(appPath)
	if err != nil {
		utils.Log.WithFields(fields).Fatal("Invalid path")
	}
	if !appPathInfo.IsDir() {
		utils.Log.WithFields(fields).Fatal("Path is not a directory")
	}

	if buildDocker {
		cmd.Docker([]string{"qtmoc", "-debug"}, "linux", appPath)
	} else {
		utils.Log.WithFields(fields).Debug("Running...")
		if err = moc.MocTree(appPath, "desktop"); err != nil { //TODO: allow buildTarget specification
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
