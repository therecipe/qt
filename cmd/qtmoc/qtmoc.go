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
	cmd.ParseFlags()

	var (
		appPath string
		fields  = logrus.Fields{"func": "main"}
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
	default:
		fmt.Println("Only specify one path")
		fmt.Printf("%s path/to/app\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	// validate that path is readable and a directory
	if !filepath.IsAbs(appPath) {
		appPath, _ = utils.Abs(appPath)
	}
	fields["app_path"] = appPath

	appPathInfo, err := os.Stat(appPath)
	if err != nil {
		utils.Log.WithFields(fields).Fatal("Invalid path")
	}
	if !appPathInfo.IsDir() {
		utils.Log.WithFields(fields).Fatal("Path is not a directory")
	}

	utils.Log.WithFields(fields).Debug("Running...")
	if err = moc.MocTree(appPath); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
