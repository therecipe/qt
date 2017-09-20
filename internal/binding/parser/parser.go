package parser

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/therecipe/qt/internal/utils"
)

var State = &struct {
	ClassMap map[string]*Class

	MocImports map[string]struct{}
	Minimal    bool //TODO:
	Target     string
}{
	ClassMap:   make(map[string]*Class),
	MocImports: make(map[string]struct{}),
}

func LoadModules() {
	libs := GetLibs()
	modules := make([]*Module, len(libs))
	modulesMutex := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	wg.Add(len(GetLibs()))
	for i, m := range libs {
		go func(i int, m string) {
			mod := LoadModule(m)

			modulesMutex.Lock()
			modules[i] = mod
			modulesMutex.Unlock()
			wg.Done()
		}(i, m)
	}
	wg.Wait()

	for _, m := range modules {
		if m != nil {
			m.Prepare()
		}
	}
}

func LoadModule(m string) *Module {
	var (
		logName   = "parser.LoadModule"
		logFields = logrus.Fields{"module": m}
	)
	utils.Log.WithFields(logFields).Debug(logName)

	if m == "Sailfish" {
		return sailfishModule()
	}

	module := new(Module)
	var err error
	switch {
	case utils.QT_WEBKIT() && m == "WebKit":
		if utils.QT_HOMEBREW() {
			err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt", "internal", "binding", "files", "docs", "5.7.1", fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
		} else {
			err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt", "internal", "binding", "files", "docs", "5.8.0", fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
		}

	case utils.QT_MXE(), utils.QT_MSYS2() && utils.QT_MSYS2_STATIC():
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt", "internal", "binding", "files", "docs", "5.8.0", fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_HOMEBREW(), utils.QT_MSYS2():
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt", "internal", "binding", "files", "docs", "5.9.0", fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

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
		return nil
	}

	return module
}
