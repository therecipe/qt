package parser

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/therecipe/qt/internal/utils"
)

var State = &struct {
	ClassMap   map[string]*Class
	GoClassMap map[string]*Class

	MocImports map[string]struct{}
	Minimal    bool //TODO:
	Target     string
}{
	ClassMap:   make(map[string]*Class),
	GoClassMap: make(map[string]*Class),
	MocImports: make(map[string]struct{}),
}

func LoadModules(target string) {
	State.Target = target

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

func LoadModulesM(target string) {
	State.Target = target

	libs := GetLibs()
	modules := make([]*Module, len(libs))
	modulesMutex := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	var i int
	for _, m := range libs {
		if !ShouldBuildForTargetM(m, target) {
			continue
		}

		wg.Add(1)
		go func(i int, m string) {
			mod := LoadModule(m)

			modulesMutex.Lock()
			modules[i] = mod
			modulesMutex.Unlock()
			wg.Done()
		}(i, m)
		i++
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

	if utils.QT_API_NUM(utils.QT_VERSION()) >= 5120 && m == "QuickControls2" {
		m = "QuickControls"
	}

	module := new(Module)
	var err error
	switch {
	case utils.QT_WEBKIT() && m == "WebKit":
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/5.8.0", "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_MXE():
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API("5.13.0"), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_HOMEBREW():
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API("5.12.0"), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_MACPORTS(), utils.QT_NIX(), utils.QT_FELGO():
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API("5.11.1"), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_MSYS2():
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API("5.12.0"), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_UBPORTS_VERSION() == "xenial":
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API("5.9.0"), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_SAILFISH():
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API("5.6.3"), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_RPI():
		err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API("5.7.0"), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)

	case utils.QT_PKG_CONFIG():
		if utils.QT_API("") != "" {
			err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API(utils.QT_VERSION()), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
		} else {
			err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(utils.QT_DOC_DIR(), fmt.Sprintf("qt%v", strings.ToLower(m)), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
			if err != nil {
				utils.Log.WithFields(logFields).Debug("falling back to bundled *.index files")
				err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API("5.13.0"), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
			}
		}

	default:
		if utils.QT_API("") != "" {
			err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API(utils.QT_VERSION()), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
		} else {
			path := filepath.Join(utils.QT_DIR(), "Docs", fmt.Sprintf("Qt-%v", utils.QT_VERSION_MAJOR()), fmt.Sprintf("qt%v", strings.ToLower(m)), fmt.Sprintf("qt%v.index", strings.ToLower(m)))
			if !utils.ExistsDir(filepath.Join(utils.QT_DIR(), "Docs", fmt.Sprintf("Qt-%v", utils.QT_VERSION_MAJOR()))) {
				path = filepath.Join(utils.QT_DIR(), "Docs", fmt.Sprintf("Qt-%v", utils.QT_VERSION()), fmt.Sprintf("qt%v", strings.ToLower(m)), fmt.Sprintf("qt%v.index", strings.ToLower(m)))
			}
			err = xml.Unmarshal([]byte(utils.LoadOptional(path)), &module)
			if err != nil {
				utils.Log.WithFields(logFields).Debug("falling back to bundled *.index files")
				err = xml.Unmarshal([]byte(utils.LoadOptional(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API(utils.QT_VERSION()), "-find", "get doc dir")), fmt.Sprintf("qt%v.index", strings.ToLower(m))))), &module)
			}
		}
	}
	if err != nil {
		switch m {
		case "Charts", "DataVisualization", "VirtualKeyboard":
			utils.Log.WithFields(logFields).WithError(err).Debug(logName)
		default:
			utils.Log.WithFields(logFields).WithError(err).Warn(logName)
		}
		return nil
	}

	return module
}
