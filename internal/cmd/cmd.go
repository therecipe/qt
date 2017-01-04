package cmd

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/therecipe/qt/internal/utils"
)

func ParseFlags() {
	var (
		debug = flag.Bool("debug", false, "print debug logs")
	)
	flag.Parse()

	if *debug {
		utils.Log.Level = logrus.DebugLevel
	}
}
