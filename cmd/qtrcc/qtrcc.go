package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/rcc"
	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var appPath, _ = os.Getwd()

	var output_dir = flag.String("o", "", "define alternative output dir")
	cmd.ParseFlags()

	if *output_dir != "" {
		if !filepath.IsAbs(*output_dir) {
			var tmp_output_dir, _ = utils.Abs(*output_dir)
			output_dir = &tmp_output_dir
		}
	}

	switch flag.NArg() {
	case 1:
		{
			appPath = flag.Arg(0)
		}
	}
	if !filepath.IsAbs(appPath) {
		appPath, _ = utils.Abs(appPath)
	}
	if _, err := ioutil.ReadDir(appPath); err != nil {
		utils.Log.Fatalln("usage:", "qtrcc", filepath.Join("path", "to", "project"))
	}

	rcc.Rcc(appPath, output_dir)
}
