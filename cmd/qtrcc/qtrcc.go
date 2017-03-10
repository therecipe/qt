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
	if utils.QT_QMAKE_CGO() == true {
		qmake_main()
		return
	}

	var appPath, _ = os.Getwd()
	if env_cwd := os.Getenv("QTRCC_CWD"); env_cwd != "" {
		appPath = env_cwd
	}

	var output_dir = flag.String("o", "", "define alternative output dir")
	if cmd.ParseFlags() {
		flag.PrintDefaults()
		return
	}

	if *output_dir != "" {
		if !filepath.IsAbs(*output_dir) {
			var tmp_output_dir, _ = filepath.Abs(*output_dir)
			output_dir = &tmp_output_dir
		}
	} else {
		env_output_dir := os.Getenv("QTRCC_OUTPUT_DIR")
		if env_output_dir != "" {
			if !filepath.IsAbs(env_output_dir) {
				var tmp_output_dir, _ = filepath.Abs(env_output_dir)
				output_dir = &tmp_output_dir
			} else {
				output_dir = &env_output_dir
			}
		}
	}

	var buildDocker bool

	switch flag.NArg() {
	case 1:
		{
			appPath = flag.Arg(0)
		}

	case 2:
		{
			appPath = flag.Arg(0)
			buildDocker = true
		}
	}
	if !filepath.IsAbs(appPath) {
		appPath, _ = filepath.Abs(appPath)
	}
	if _, err := ioutil.ReadDir(appPath); err != nil {
		utils.Log.Fatalln("usage:", "qtrcc", filepath.Join("path", "to", "project"), "[ docker ]")
	}

	if buildDocker {
		cmd.Docker([]string{"qtrcc", "-debug"}, "linux", appPath)
	} else {
		rcc.Rcc(appPath, "desktop", output_dir) //TODO: allow buildTarget specification
	}
}
