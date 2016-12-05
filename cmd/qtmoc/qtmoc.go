package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/therecipe/qt/internal/utils"
	"github.com/therecipe/qt/internal/moc"
)

func main() {
	var (
		appPath string
		cleanup bool
		debug   = flag.Bool("debug", false, "Print debug logs")
		fields  = logrus.Fields{"func": "main"}
	)
	defaultAppPath, err := os.Getwd()
	if err != nil {
		utils.Log.WithFields(fields).Fatal("Can't get current working directory")
	}
	flag.BoolVar(&cleanup, "cleanup", false, "Cleanup (not default)")
	flag.Parse()

	if *debug {
		utils.Log.Level = logrus.DebugLevel
	}

	switch flag.NArg() {
	case 0:
		appPath = defaultAppPath
	case 1:
		appPath = flag.Args()[0]
	default:
		fmt.Println("Only specify one path")
		fmt.Printf("%s path/to/app\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	// validate that path is readable and a directory
	if !filepath.IsAbs(appPath) {
		appPath = utils.GetAbsPath(appPath)
	}
	fields["app_path"] = appPath

	appPathInfo, err := os.Stat(appPath)
	if err != nil {
		utils.Log.WithFields(fields).Fatal("Invalid path")
		os.Exit(1)
	}
	if !appPathInfo.IsDir() {
		utils.Log.WithFields(fields).Fatal("Path is not a directory")
		os.Exit(1)
	}

	utils.Log.WithFields(fields).Debug("Running...")
	if err = moc.MocTree(appPath, cleanup); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
