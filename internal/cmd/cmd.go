package cmd

import (
	"flag"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/therecipe/qt/internal/utils"
)

func ParseFlags() {
	var (
		debug = flag.Bool("debug", false, "print debug logs")
		p     = flag.Int("p", runtime.NumCPU(), "the number of cpu's used")
	)
	flag.Parse()

	if *debug {
		utils.Log.Level = logrus.DebugLevel
	}

	runtime.GOMAXPROCS(*p)
}
