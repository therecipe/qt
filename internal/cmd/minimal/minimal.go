package minimal

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/cmd"

	"github.com/therecipe/qt/internal/utils"
)

func Minimal(path, target string) {
	utils.Log.WithField("path", path).WithField("target", target).Debug("start Minimal")

	//TODO: cleanup state from moc for minimal first -->
	for _, c := range parser.State.ClassMap {
		if c.Module == parser.MOC || strings.HasPrefix(c.Module, "custom_") {
			delete(parser.State.ClassMap, c.Name)
		}
	}
	parser.LibDeps[parser.MOC] = make([]string, 0)
	//<--

	var files []string
	for _, path := range append([]string{path}, cmd.GetImports(path, target, 0, false)...) {
		fileList, err := ioutil.ReadDir(path)
		if err != nil {
			utils.Log.WithError(err).Error("failed to read dir")
			continue
		}

		for _, file := range fileList {
			path := filepath.Join(path, file.Name())
			if !file.IsDir() && filepath.Ext(path) == ".go" {
				utils.Log.WithField("path", path).Debug("analyse for minimal")
				files = append(files, utils.Load(path))
			}
		}
	}

	c := len(files)
	utils.Log.Debugln("found", c, "files to analyse")
	if c == 0 {
		return
	}

	if len(parser.State.ClassMap) == 0 {
		parser.LoadModules()
	} else {
		utils.Log.Debug("modules already cached")
	}

	//TODO: merge sailfish and asteroid
	switch target {
	case "sailfish", "sailfish-emulator":
		for _, bl := range []string{"TestCase", "QQuickWidget", "QLatin1String", "QStringRef"} {
			if c, ok := parser.State.ClassMap[bl]; ok {
				c.Export = false
				delete(parser.State.ClassMap, c.Name)
			}
		}

		for _, c := range parser.State.ClassMap {
			since, err := strconv.ParseFloat(strings.TrimPrefix(c.Since, "Qt "), 64)
			if err != nil {
				continue
			}
			if since >= 5.3 || !parser.IsWhiteListedSailfishLib(strings.TrimPrefix(c.Module, "Qt")) {
				c.Export = false
				delete(parser.State.ClassMap, c.Name)
			}

			for _, f := range c.Functions {
				since, err := strconv.ParseFloat(strings.TrimPrefix(f.Since, "Qt "), 64)
				if err != nil {
					continue
				}
				if since >= 5.3 {
					f.Export = false
				}
			}
		}

	case "asteroid":
		for _, bl := range []string{"TestCase", "QQuickWidget"} {
			if c, ok := parser.State.ClassMap[bl]; ok {
				c.Export = false
				delete(parser.State.ClassMap, c.Name)
			}
		}

		for _, c := range parser.State.ClassMap {
			since, err := strconv.ParseFloat(strings.TrimPrefix(c.Since, "Qt "), 64)
			if err != nil {
				continue
			}
			if since >= 5.7 || !parser.IsWhiteListedSailfishLib(strings.TrimPrefix(c.Module, "Qt")) {
				c.Export = false
				delete(parser.State.ClassMap, c.Name)
			}

			for _, f := range c.Functions {
				since, err := strconv.ParseFloat(strings.TrimPrefix(f.Since, "Qt "), 64)
				if err != nil {
					continue
				}
				if since >= 5.7 {
					f.Export = false
				}
			}
		}

	case "ios", "ios-simulator":
		for _, bl := range []string{"QProcess", "QProcessEnvironment"} {
			if c, ok := parser.State.ClassMap[bl]; ok {
				c.Export = false
				delete(parser.State.ClassMap, bl)
			}
		}
	}

	for _, f := range files {
		for _, c := range parser.State.ClassMap {
			if strings.Contains(f, c.Name) &&
				strings.Contains(f, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(strings.TrimPrefix(c.Module, "Qt")))) {
				exportClass(c, files)
			}
		}
	}

	//TODO: cleanup state
	parser.State.Minimal = true
	for _, m := range parser.GetLibs() {
		templater.GenModule(m, target, templater.MINIMAL)
	}
	parser.State.Minimal = false
	for _, c := range parser.State.ClassMap {
		c.Export = false
		for _, f := range c.Functions {
			f.Export = false
		}
	}
}

func exportClass(c *parser.Class, files []string) {
	if c.Export {
		return
	}
	c.Export = true

	for _, file := range files {
		for _, f := range c.Functions {

			switch {
			case f.Virtual == parser.IMPURE, f.Virtual == parser.PURE, f.Meta == parser.SIGNAL, f.Meta == parser.SLOT:
				for _, mode := range []string{parser.CONNECT, parser.DISCONNECT, ""} {
					f.SignalMode = mode
					if strings.Contains(file, converter.GoHeaderName(f)) {
						exportFunction(f, files)
					}
				}

			default:
				if f.Static {
					if strings.Contains(file, converter.GoHeaderName(f)) {
						exportFunction(f, files)
					}
					f.Static = false
					if strings.Contains(file, converter.GoHeaderName(f)) {
						exportFunction(f, files)
					}
					f.Static = true
				} else {
					if strings.Contains(file, converter.GoHeaderName(f)) {
						exportFunction(f, files)
					}
				}
			}

			if strings.HasPrefix(f.Name, "__") || f.Meta == parser.CONSTRUCTOR ||
				f.Meta == parser.DESTRUCTOR || f.Virtual == parser.PURE {
				exportFunction(f, files)
			}

		}
	}

	for _, b := range c.GetAllBases() {
		if c, ok := parser.State.ClassMap[b]; ok {
			exportClass(c, files)
		}
	}
}

func exportFunction(f *parser.Function, files []string) {
	if f.Export {
		return
	}
	f.Export = true

	for _, p := range f.Parameters {
		if c, ok := parser.State.ClassMap[parser.CleanValue(p.Value)]; ok {
			exportClass(c, files)
		}
		if parser.IsPackedList(p.Value) {
			if c, ok := parser.State.ClassMap[parser.UnpackedList(p.Value)]; ok {
				exportClass(c, files)
			}
		}
		if parser.IsPackedMap(p.Value) {
			key, value := parser.UnpackedMap(p.Value)
			if c, ok := parser.State.ClassMap[key]; ok {
				exportClass(c, files)
			}
			if c, ok := parser.State.ClassMap[value]; ok {
				exportClass(c, files)
			}
		}
	}

	if c, ok := parser.State.ClassMap[parser.CleanValue(f.Output)]; ok {
		exportClass(c, files)
	}
	if parser.IsPackedList(f.Output) {
		if c, ok := parser.State.ClassMap[parser.UnpackedList(f.Output)]; ok {
			exportClass(c, files)
		}
	}
	if parser.IsPackedMap(f.Output) {
		key, value := parser.UnpackedMap(f.Output)
		if c, ok := parser.State.ClassMap[key]; ok {
			exportClass(c, files)
		}
		if c, ok := parser.State.ClassMap[value]; ok {
			exportClass(c, files)
		}
	}
}
