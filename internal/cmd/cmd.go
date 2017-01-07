package cmd

import (
	"flag"
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/therecipe/qt/internal/utils"
)

func ParseFlags() {
	var (
		debug      = flag.Bool("debug", false, "print debug logs")
		p          = flag.Int("p", runtime.NumCPU(), "the number of cpu's used")
		qt_dir     = flag.String("qt_dir", "", "define alternative qt dir")
		qt_version = flag.String("qt_version", "", "define alternative qt version")
	)
	flag.Parse()

	if *debug {
		utils.Log.Level = logrus.DebugLevel
		os.Setenv("QT_DEBUG", "true")
	}

	runtime.GOMAXPROCS(*p)

	if dir := *qt_dir; dir != "" {
		os.Setenv("QT_DIR", dir)
	}

	if version := *qt_version; version != "" {
		os.Setenv("QT_VERSION", version)
	}
}
