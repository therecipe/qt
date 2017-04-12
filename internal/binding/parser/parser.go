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

	Minimal bool //TODO:
	Target  string
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
		logFields = logrus.Fields{"module": m}
	)
	utils.Log.WithFields(logFields).Debug(logName)

	if m == "Sailfish" {
		return sailfishModule().Prepare()
	}

	module := new(Module)
	var err error
	switch {
	case utils.QT_WEBKIT() && m == "WebKit":
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt", "internal", "binding", "files", "docs", "5.7.1", fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_HOMEBREW(), utils.QT_MSYS2():
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt", "internal", "binding", "files", "docs", utils.QT_VERSION(), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_PKG_CONFIG():
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(utils.QT_DOC_DIR(), fmt.Sprintf("qt%v", strings.ToLower(m)), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	default:
		err = xml.Unmarshal([]byte(utils.Load(filepath.Join(utils.QT_DIR(), "Docs", fmt.Sprintf("Qt-%v", utils.QT_VERSION_MAJOR()), fmt.Sprintf("qt%v", strings.ToLower(m)), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
	}
	if err != nil {
		if m != "DataVisualization" && m != "Charts" {
			utils.Log.WithFields(logFields).WithError(err).Warn(logName)
		} else {
			utils.Log.WithFields(logFields).WithError(err).Debug(logName)
		}
		return err
	}

	return module.Prepare()
}
