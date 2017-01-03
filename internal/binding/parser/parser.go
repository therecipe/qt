package parser

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/therecipe/qt/internal/utils"
)

var State = &struct {
	ClassMap map[string]*Class

	Moc     bool
	Minimal bool

	Module string
}{
	ClassMap: make(map[string]*Class),
}

func LoadModules() {
	for _, m := range GetLibs() {
		_ = LoadModule(m)
	}
}

func LoadModule(m string) error {
	var (
		logName   = "parser.LoadModule"
		logFields = logrus.Fields{"0_module": m}
	)
	utils.Log.WithFields(logFields).Debug(logName)

	if m == "Sailfish" {
		return sailfishModule().Prepare()
	}

	var (
		module = new(Module)
		err    error
	)
	switch {
	case utils.UseHomeBrew(), utils.UseMsys2():
		{
			err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt", "internal", "binding", "files", "docs", utils.QT_VERSION(), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
		}

	case utils.UsePkgConfig():
		{
			err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(utils.QT_DOC_DIR(), fmt.Sprintf("qt%v", strings.ToLower(m)), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
		}

	default:
		{
			err = xml.Unmarshal([]byte(utils.Load(filepath.Join(utils.QT_DIR(), "Docs", "Qt-5.7", fmt.Sprintf("qt%v", strings.ToLower(m)), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
		}
	}
	if err != nil {
		if m != "datavisualization" && m != "charts" {
			utils.Log.WithFields(logFields).WithError(err).Warn(logName)
		} else {
			utils.Log.WithFields(logFields).WithError(err).Debug(logName)
		}
		return err
	}

	return module.Prepare()
}
