package minimal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/bluszcz/cutego/internal/binding/converter"
	"github.com/bluszcz/cutego/internal/binding/parser"
	"github.com/bluszcz/cutego/internal/binding/templater"

	"github.com/bluszcz/cutego/internal/cmd"
	"github.com/bluszcz/cutego/internal/cmd/moc"

	"github.com/bluszcz/cutego/internal/utils"
)

func Minimal(path, target, tags string, skipSetup bool) {
	rootPath := path
	defer func() {
		if cmd.ImportsQmlOrQuick() || cmd.ImportsInterop() { //TODO: and not deploying + reinstate on moc.moc with deploying ?
			if pkg_path := utils.GoQtPkgPath("internal/binding/runtime"); utils.QT_NOT_CACHED() || !utils.ExistsFile(filepath.Join(pkg_path, templater.CgoFileNames(pkg_path, target, templater.MOC)[0])) {
				filepath.Walk(pkg_path, func(path string, info os.FileInfo, err error) error {
					if err != nil || info.IsDir() {
						return err
					}
					if strings.HasPrefix(info.Name(), "moc_cgo") || strings.HasPrefix(info.Name(), "rcc_cgo") { //rcc_cgo for static linux workaround
						os.Remove(path)
					}
					return nil
				})
				moc.Moc(pkg_path, target, "", true, false, false, true)
			} else {
				moc.ResourceNamesMutex.Lock()
				moc.ResourceNames[filepath.Join(pkg_path, "moc.cpp")] = ""
				moc.ResourceNamesMutex.Unlock()
			}
		}
	}()

	if utils.UseGOMOD(path) && !skipSetup {
		if !utils.ExistsDir(filepath.Join(filepath.Dir(utils.GOMOD(path)), "vendor")) {
			cmd := exec.Command("go", "mod", "vendor")
			cmd.Dir = path
			utils.RunCmd(cmd, "go mod vendor")
		}
		if utils.QT_DOCKER() {
			cmd := exec.Command("go", "get", "-v", "-d", "github.com/bluszcz/cutego/internal/binding/files/docs/"+utils.QT_API(utils.QT_VERSION())+"@"+cmd.QtModVersion(filepath.Dir(utils.GOMOD(path)))) //TODO: needs to pull 5.8.0 if QT_WEBKIT
			cmd.Dir = path
			if !utils.QT_PKG_CONFIG() {
				utils.RunCmdOptional(cmd, "go get docs") //TODO: this can fail if QT_PKG_CONFIG
			}

			if strings.HasPrefix(target, "sailfish") || strings.HasPrefix(target, "android") { //TODO: generate android and sailfish minimal instead
				cmd := exec.Command(filepath.Join(utils.GOBIN(), "qtsetup"), "generate", target)
				cmd.Dir = path
				utils.RunCmd(cmd, "run setup")
			}
		}
	}

	if !(target == "js" || target == "wasm" || utils.QT_NOT_CACHED()) { //TODO: remove for module support + resolve dependencies
		env, tagsEnv, _, _ := cmd.BuildEnv(target, "", "")
		scmd := utils.GoList("{{.Stale}}|{{.StaleReason}}")
		scmd.Dir = path

		tagsEnv = append(tagsEnv, "minimal")

		if tags != "" {
			tagsEnv = append(tagsEnv, strings.Split(tags, " ")...)
		}
		scmd.Args = append(scmd.Args, utils.BuildTags(tagsEnv))

		if target != runtime.GOOS {
			scmd.Args = append(scmd.Args, []string{"-pkgdir", filepath.Join(utils.MustGoPath(), "pkg", fmt.Sprintf("%v_%v_%v", strings.Replace(target, "-", "_", -1), env["GOOS"], env["GOARCH"]))}...)
		}

		for key, value := range env {
			scmd.Env = append(scmd.Env, fmt.Sprintf("%v=%v", key, value))
		}

		if out := utils.RunCmdOptional(scmd, fmt.Sprintf("go check stale for %v on %v", target, runtime.GOOS)); strings.Contains(out, "but available in build cache") || strings.Contains(out, "false|") {
			utils.Log.WithField("path", path).Debug("skipping already cached minimal")
			return
		}
	}

	utils.Log.WithField("path", path).WithField("target", target).Debug("start Minimal")

	//TODO: cleanup state from moc for minimal first -->
	for _, c := range parser.State.ClassMap {
		if c.Module == parser.MOC || strings.HasPrefix(c.Module, "custom_") || c.ToBeCleanedUp {
			delete(parser.State.ClassMap, c.Name)
		}
	}
	parser.LibDeps[parser.MOC] = make([]string, 0)
	if target == "js" || target == "wasm" { //TODO: REVIEW
		if parser.LibDeps["build_static"][0] == "Qml" {
			parser.LibDeps["build_static"] = parser.LibDeps["build_static"][1:]
		}
	} else {
		parser.LibDeps["build_static"] = []string{"Qml"}
	}
	//<--

	wg := new(sync.WaitGroup)
	wc := make(chan bool, runtime.NumCPU()*2)

	var files []string
	fileMutex := new(sync.Mutex)

	allImports := append([]string{path}, cmd.GetImports(path, target, tags, 0, false)...)
	wg.Add(len(allImports))
	for _, path := range allImports {
		wc <- true
		go func(path string) {
			for _, path := range cmd.GetGoFiles(path, target, tags) {
				if base := filepath.Base(path); strings.HasPrefix(base, "rcc_cgo") || strings.HasPrefix(base, "moc_cgo") {
					continue
				}
				utils.Log.WithField("path", path).Debug("analyse for minimal")
				file := utils.Load(path)
				fileMutex.Lock()
				files = append(files, file)
				fileMutex.Unlock()
			}
			if target == "js" || cmd.ImportsQmlOrQuick() || cmd.ImportsInterop() { //TODO: wasm as well
				filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
					if err != nil || info.IsDir() || strings.HasPrefix(path, filepath.Join(rootPath, "deploy")) {
						return err
					}
					if filepath.Ext(path) == ".js" || (filepath.Ext(path) == ".ts" && !strings.HasSuffix(path, ".d.ts")) || filepath.Ext(path) == ".qml" ||
						filepath.Ext(path) == ".dart" || filepath.Ext(path) == ".hx" || filepath.Ext(path) == ".swift" {
						file := utils.Load(path)
						fileMutex.Lock()
						if strings.Contains(file, utils.PackageName) {
							files = append(files, file)
							utils.Log.WithField("path", path).Debug("analyse for minimal")
						} else {
							utils.Log.WithField("path", path).Debug("ignore for minimal")
						}
						fileMutex.Unlock()
					}
					return nil
				})
			}
			<-wc
			wg.Done()
		}(path)
	}
	wg.Wait()

	c := len(files)
	utils.Log.Debugln("found", c, "files to analyze")
	if c == 0 {
		return
	}

	if _, ok := parser.State.ClassMap["QObject"]; !ok {
		parser.LoadModulesM(target)
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
				continue
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
				continue
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
		exportClass(parser.State.ClassMap["QSvgWidget"], files)

	case "android", "android-emulator":
		exportClass(parser.State.ClassMap["QSvgWidget"], files)

	case "rpi1", "rpi2", "rpi3":
		if !utils.QT_RPI() {
			break
		}
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
			if since >= 5.8 || !parser.IsWhiteListedRaspberryLib(strings.TrimPrefix(c.Module, "Qt")) {
				c.Export = false
				delete(parser.State.ClassMap, c.Name)
				continue
			}

			for _, f := range c.Functions {
				since, err := strconv.ParseFloat(strings.TrimPrefix(f.Since, "Qt "), 64)
				if err != nil {
					continue
				}
				if since >= 5.8 {
					f.Export = false
				}
			}
		}
	case "js", "wasm":
		exportClass(parser.State.ClassMap["QSvgWidget"], files)
	}
	//"android", "android-emulator": //TODO: generate minimal androidextras instead, otherwise using the androidextras module might report missing classes when building for targets other than android
	exportClass(parser.State.ClassMap["QChildEvent"], files)
	exportClass(parser.State.ClassMap["QTimerEvent"], files)
	exportClass(parser.State.ClassMap["QMetaObject"], files)
	exportClass(parser.State.ClassMap["QEvent"], files)
	exportClass(parser.State.ClassMap["QMetaMethod"], files)
	exportClass(parser.State.ClassMap["QByteArray"], files)
	exportClass(parser.State.ClassMap["QVariant"], files)
	exportClass(parser.State.ClassMap["QObject"], files)

	if utils.QT_STATIC() {
		exportClass(parser.State.ClassMap["QSvgWidget"], files)
	}
	if utils.QT_FELGO() || utils.QT_GEN_GO_WRAPPER() {
		exportClass(parser.State.ClassMap["QCoreApplication"], files)
		exportFunction(parser.State.ClassMap["QCoreApplication"].GetFunction("instance"), files)
	}

	wg.Add(len(files))
	for _, f := range files {
		go func(f string) {
			for _, c := range parser.State.ClassMap {
				if strings.Contains(f, c.Name) &&
					strings.Contains(f, fmt.Sprintf("%v/%v", utils.PackageName, strings.ToLower(strings.TrimPrefix(c.Module, "Qt")))) {
					exportClass(c, files)
				}
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

	if _, ok := parser.State.ClassMap["QVariant"]; ok {
		exportClass(parser.State.ClassMap["QVariant"], files)
		for _, fn := range []string{"type", "canConvert", "toList", "toMap", "isValid", "toString", "isNull"} {
			exportFunction(parser.State.ClassMap["QVariant"].GetFunction(fn), files)
		}

		for _, v := range parser.State.ClassMap["QVariant"].Enums[0].Values {
			if f := parser.State.ClassMap["QVariant"].GetFunction("to" + v.Name); f != nil {
				if _, ok := parser.IsClass("Q" + v.Name); !ok ||
					(v.Name == "Map" ||
						v.Name == "String" ||
						v.Name == "StringList" ||
						v.Name == "Hash") {
					exportFunction(f, files)
				}
			}
		}
	}

	if _, ok := parser.State.ClassMap["QObject"]; ok {
		exportClass(parser.State.ClassMap["QObject"], files)
		for _, fn := range []string{"setObjectName", "objectNameChanged"} {
			exportFunction(parser.State.ClassMap["QObject"].GetFunction(fn), files)
		}
	}

	if _, ok := parser.State.ClassMap["QJSValue"]; ok {
		exportClass(parser.State.ClassMap["QJSValue"], files)
		for _, fn := range []string{"property", "setProperty", "toQObject", "isCallable", "isNull", "isUndefined", "toString", "call", "toInt", "isString", "isObject"} {
			exportFunction(parser.State.ClassMap["QJSValue"].GetFunction(fn), files)
		}
	}

	if _, ok := parser.State.ClassMap["QJSEngine"]; ok {
		exportClass(parser.State.ClassMap["QJSEngine"], files)
		for _, fn := range []string{"newQObject", "newObject", "qjsEngine", "toScriptValue", "globalObject", "newArray", "evaluate"} {
			exportFunction(parser.State.ClassMap["QJSEngine"].GetFunction(fn), files)
		}
	}

	if _, ok := parser.State.ClassMap["QQuickView"]; ok {
		exportClass(parser.State.ClassMap["QQuickView"], files)
		for _, fn := range []string{"engine"} {
			exportFunction(parser.State.ClassMap["QQuickView"].GetFunction(fn), files)
		}
	}

	//TODO: cleanup state
	parser.State.Minimal = true
	for _, m := range parser.GetLibs() {
		if !parser.ShouldBuildForTargetM(m, target) ||
			m == "AndroidExtras" || m == "Sailfish" {
			continue
		}

		utils.MkdirAll(utils.GoQtPkgPath(strings.ToLower(m)))
		if !utils.QT_GEN_GO_WRAPPER() {
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.cpp"), templater.CppTemplate(m, templater.MINIMAL, target, ""))
			utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.h"), templater.HTemplate(m, templater.MINIMAL, ""))
		} else {
			utils.RemoveAll(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.cpp"))
			utils.RemoveAll(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.h"))
			utils.RemoveAll(utils.GoQtPkgPath(strings.ToLower(m), templater.CgoFileNames(utils.GoQtPkgPath(strings.ToLower(m)), target, templater.MINIMAL)[0]))
		}
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.go"), templater.GoTemplate(m, false, templater.MINIMAL, m, target, ""))
	}

	if !utils.QT_GEN_GO_WRAPPER() {
		for _, m := range parser.GetLibs() {
			if !parser.ShouldBuildForTargetM(m, target) ||
				m == "AndroidExtras" || m == "Sailfish" {
				continue
			}

			wg.Add(1)
			go func(m string, libs []string) {
				templater.CgoTemplateSafe(m, "", target, templater.MINIMAL, m, "", libs)
				wg.Done()
			}(m, parser.LibDeps[m])
		}
	}
	wg.Wait()

	parser.State.Minimal = false
	for _, c := range parser.State.ClassMap {
		c.Export = false
		for _, f := range c.Functions {
			f.Export = false
		}
	}
}

func exportClass(c *parser.Class, files []string) {
	if c == nil {
		return
	}
	if c.Name == "FelgoLiveClient" && !utils.QT_FELGO_LIVE() {
		return
	}
	c.Lock()
	exp := c.Export
	c.Unlock()
	if exp {
		return
	}
	c.Lock()
	c.Export = true
	c.Unlock()

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

			if f.Name == "toVariant" {
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
	if f == nil {
		return
	}
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

	if c, ok := parser.State.ClassMap[parser.CleanValue(f.Output)]; ok && f.Virtual == parser.PURE {
		for _, f := range c.Functions {
			if f.Meta == parser.CONSTRUCTOR && len(f.Parameters) == 0 {
				exportFunction(f, files)
			}
		}
	}

	for _, ff := range parser.State.ClassMap[f.ClassName()].Functions {
		if f.Name == ff.Name && f.OverloadNumber != ff.OverloadNumber {
			exportFunction(ff, files)
		}
	}
}
