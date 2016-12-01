package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"
	"github.com/therecipe/qt/internal/utils"
)

var (
	tmpFiles  = make([]string, 0)
	blacklist = []string{
		"deploy",
		"qml",
		"android",
		"ios",
		"ios-simulator",
		"sailfish",
		"sailfish-emulator",
		"rpi1",
		"rpi2",
		"rpi3",
		"node_modules",
		".git",
	}
)

func init() {
	templater.Moc = true
}

func main() {
	var (
		appPath string
		cleanup bool
		debug = flag.Bool("debug", false, "Print debug logs")
		fields  = logrus.Fields{"func": "main"}
	)
	defaultAppPath, err := os.Getwd()
	if err != nil {
		utils.Log.WithFields(fields).Fatal("Can't get current working directory")
	}
	flag.BoolVar(&cleanup, "cleanup", false, "Cleanup (not default)")
	flag.Parse()

	if *debug {
		utils.Log.Level = logrus.DebugLevel
	}

	switch flag.NArg() {
	case 0:
		appPath = defaultAppPath
	case 1:
		appPath = flag.Args()[0]
	default:
		fmt.Println("Only specify one path")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// validate that path is readable and a directory
	if !filepath.IsAbs(appPath) {
		appPath = utils.GetAbsPath(appPath)
	}
	fields["app_path"] = appPath

	appPathInfo, err := os.Stat(appPath)
	if err != nil {
		utils.Log.WithFields(fields).Fatal("Invalid path")
		os.Exit(1)
	}
	if !appPathInfo.IsDir() {
		utils.Log.WithFields(fields).Fatal("Path is not a directory")
		os.Exit(1)
	}

	err = filepath.Walk(appPath, func(path string, info os.FileInfo, err error) error {
		walkFields := logrus.Fields{
			"path": path,
		}
		if err != nil {
			utils.Log.WithError(err).WithFields(walkFields).WithFields(fields).Debug("Error walk, skip")
			return nil
		}
		fields["filename"] = info.Name()

		if !info.IsDir() {
			utils.Log.WithError(err).WithFields(walkFields).WithFields(fields).Debug("Not directory, skip")
			return nil
		}

		if isBlacklisted(appPath, filepath.Join(path, info.Name())) {
			utils.Log.WithFields(walkFields).WithFields(fields).Debug("Blacklisted, skip")
		}
		return moc(path)
	})
	if err != nil {
		utils.Log.WithError(err).WithFields(fields).Fatalln("Can't process")
	}

	if cleanup {
		var b, err = json.Marshal(tmpFiles)
		if err == nil {
			utils.SaveBytes(filepath.Join(appPath, "moc_cleanup.json"), b)
		} else {
			utils.Log.WithError(err).WithFields(fields).Error("failed to save moc_cleanup.json")
		}
	}
}

func moc(appPath string) (err error) {
	fields := logrus.Fields{
		"func":     "moc",
		"app_path": appPath,
	}
	utils.Log.WithFields(fields).Debug("Initialize")

	for _, n := range []string{
		"moc.h", "moc.go", "moc.cpp", "moc_moc.h",
		"moc_cgo_desktop_darwin_amd64.go", "moc_cgo_desktop_windows_386.go", "moc_cgo_desktop_windows_amd64.go", "moc_cgo_desktop_linux_amd64.go",
		"moc_cgo_android_linux_arm.go",
		"moc_cgo_ios_simulator_darwin_amd64.go", "moc_cgo_ios_simulator_darwin_386.go", "moc_cgo_ios_darwin_arm64.go", "moc_cgo_ios_darwin_arm.go",
		"moc_cgo_sailfish_emulator_linux_386.go", "moc_cgo_sailfish_linux_arm.go",
		"moc_cgo_rpi1_linux_arm.go", "moc_cgo_rpi2_linux_arm.go", "moc_cgo_rpi3_linux_arm.go",
	} {
		if utils.Exists(filepath.Join(appPath, n)) {
			utils.Log.WithFields(fields).WithField("filename", n).Debug("File already exists, cleanup")
			if err = utils.RemoveAll(filepath.Join(appPath, n)); err != nil {
				return
			}
		}
		tmpFiles = append(tmpFiles, filepath.Join(appPath, n))
	}
	if err = utils.RemoveAll(filepath.Join(appPath, "moc_cleanup.json")); err != nil {
		return
	}

	module := &parser.Module{
		Project: parser.MOC,
		Namespace: &parser.Namespace{
			Classes: make([]*parser.Class, 0),
		},
	}

	err = filepath.Walk(appPath, func(path string, info os.FileInfo, wErr error) error {
		walkFields := logrus.Fields{
			"path": path,
		}
		if wErr != nil {
			utils.Log.WithError(err).WithFields(walkFields).WithFields(fields).Debug("Error, skip")
			return nil
		}
		filename := info.Name()
		walkFields["filename"] = filename

		if info.IsDir() {
			utils.Log.WithFields(walkFields).WithFields(fields).Debug("Directory, skip")
			return nil
		}

		if strings.HasPrefix(filename, "moc") {
			utils.Log.WithFields(walkFields).WithFields(fields).Debug("File with `moc` prefix, skip")
			return nil
		}

		extension := filepath.Ext(filename)
		walkFields["extension"] = extension
		if extension != ".go" {
			utils.Log.WithFields(walkFields).WithFields(fields).Debug("Not go source, skip")
			return nil
		}

		dirName := filepath.Dir(path)
		walkFields["dirname"] = dirName
		if dirName != appPath {
			utils.Log.WithFields(walkFields).WithFields(fields).Debug("Path mismatch")
			return nil
		}

		utils.Log.WithFields(walkFields).WithFields(fields).Debug("Process file")

		src, errRead := ioutil.ReadFile(path)
		if errRead != nil {
			return errRead
		}

		file, errParse := goparser.ParseFile(token.NewFileSet(), path, nil, 0)
		if errParse != nil {
			return errParse
		}

		templater.MocModule = file.Name.String()

		for _, d := range file.Decls {
			typeDecl, ok := d.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, s := range typeDecl.Specs {
				typeSpec, ok := s.(*ast.TypeSpec)
				if !ok {
					continue
				}
				structDecl, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				class := &parser.Class{
					Access:    "public",
					Module:    parser.MOC,
					Name:      typeSpec.Name.String(),
					Status:    "public",
					Functions: make([]*parser.Function, 0),
				}

				for _, field := range structDecl.Fields.List {
					//collect possible base classes
					var fieldValue = string(src[field.Pos()-1 : field.End()-1])
					if fieldValue != "" && !strings.Contains(fieldValue, " ") && !strings.HasPrefix(fieldValue, "*") {
						if strings.Contains(fieldValue, ".") {
							fieldValue = strings.Split(fieldValue, ".")[1]
						}
						class.Bases += fmt.Sprintf("%v,", fieldValue)
					}

					for range field.Names {
						if field.Tag == nil {
							continue
						}
						var meta string
						tag := strings.Replace(strings.Replace(field.Tag.Value, "\"", "", -1), "`", "", -1)

						switch {
						case strings.HasPrefix(tag, "signal:"):
							meta = parser.SIGNAL
						case strings.HasPrefix(tag, "slot:"):
							meta = parser.SLOT
						default:
							// only handle signal and slot
							continue
						}

						var (
							typ = string(src[field.Type.Pos()-1 : field.Type.End()-1])
							f   = &parser.Function{
								Access:     "public",
								Fullname:   fmt.Sprintf("%v::%v", class.Name, strings.Split(tag, ":")[1]),
								Meta:       meta,
								Name:       strings.Split(tag, ":")[1],
								Status:     "public",
								Virtual:    "non",
								Signature:  "()",
								Parameters: getParameters(typ),
								Output: func() string {
									if meta == parser.SLOT {
										var out = strings.TrimSpace(strings.Split(typ, ")")[1])
										if strings.Contains(out, "(") {
											return getParameters(out + ")")[0].Value
										}
										return out
									}
									return "void"
								}(),
							}
						)
						if len(f.Parameters[0].Value) == 0 {
							f.Parameters = make([]*parser.Parameter, 0)
						}
						class.Functions = append(class.Functions, f)

					}
				}
				module.Namespace.Classes = append(module.Namespace.Classes, class)
			}
		}
		return nil
	})
	if err != nil {
		utils.Log.WithError(err).WithFields(fields).Fatal("Can't parse code")
	}

	if len(module.Namespace.Classes) == 0 {
		utils.Log.WithFields(fields).Debug("failed to find moc structs")
		return nil
	}

	//cache modules
	if len(parser.ClassMap) == 0 {
		for _, module := range templater.GetLibs() {
			modName := strings.ToLower(module)
			loopFields := logrus.Fields{
				"module": fmt.Sprintf("qt/%v", modName),
			}
			utils.Log.WithFields(fields).WithFields(loopFields).Debug("loading module")
			if _, err := parser.GetModule(modName); err != nil {
				utils.Log.WithError(err).WithFields(fields).WithFields(loopFields).Error("failed to load module")
			}
		}
	}

	for _, c := range module.Namespace.Classes {
		for _, bc := range c.GetBases() {
			if isInClassArray(module.Namespace.Classes, bc) || isInClassMap(parser.ClassMap, bc) {
				c.Bases = bc
				break
			}
		}
	}

	module.Prepare()

	var atLeastOneMocClass bool
	for _, class := range module.Namespace.Classes {
		if !class.IsQObjectSubClass() {
			delete(parser.ClassMap, class.Name)
		} else {
			atLeastOneMocClass = true
		}
	}
	if !atLeastOneMocClass {
		utils.Log.WithFields(fields).Debug("failed to find at least one valid moc struct")
		return nil
	}

	for _, c := range module.Namespace.Classes {
		for _, f := range c.Functions {
			if !f.NoMocDeduce {
				for _, p := range f.Parameters {
					p.Value = getCppTypeFromGoType(p.Value)
				}
				f.Output = getCppTypeFromGoType(f.Output)
			}
		}
	}

	//copy constructor and destructor
	for _ = range module.Namespace.Classes {
		for _, c := range module.Namespace.Classes {
			if bc, exists := parser.ClassMap[c.Bases]; exists {
				for _, bcf := range bc.Functions {
					if bcf.Meta == parser.CONSTRUCTOR || bcf.Meta == parser.DESTRUCTOR {
						var f = *bcf
						f.Name = strings.Replace(f.Name, bcf.Class(), c.Name, -1)
						f.Fullname = strings.Replace(f.Fullname, bcf.Class(), c.Name, -1)

						if !c.HasFunctionWithName(f.Name) {
							c.Functions = append(c.Functions, &f)
						}
					}
				}
			}
		}
	}

	fields["parser_moc"] = parser.MOC
	utils.Log.WithFields(fields).Debug("generating moc")

	if err = utils.SaveBytes(filepath.Join(appPath, "moc.cpp"), templater.CppTemplate(parser.MOC)); err != nil {
		return
	}
	tmpFiles = append(tmpFiles, filepath.Join(appPath, "moc.cpp"))
	if err = utils.SaveBytes(filepath.Join(appPath, "moc.h"), templater.HTemplate(parser.MOC)); err != nil {
		return
	}
	tmpFiles = append(tmpFiles, filepath.Join(appPath, "moc.h"))
	if err = utils.SaveBytes(filepath.Join(appPath, "moc.go"), templater.GoTemplate(parser.MOC, false)); err != nil {
		return
	}
	tmpFiles = append(tmpFiles, filepath.Join(appPath, "moc.go"))

	templater.CopyCgoForMoc(parser.MOC, appPath)

	for _, c := range parser.ClassMap {
		if c.Module == parser.MOC {
			delete(parser.ClassMap, c.Name)
		}
	}

	var mocPath string
	switch runtime.GOOS {
	case "darwin":
		mocPath = filepath.Join(utils.QT_DARWIN_DIR(), "bin", "moc")
	case "linux":
		if utils.UsePkgConfig() {
			mocPath = filepath.Join(strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=host_bins", "Qt5Core"), "moc.LinuxPkgConfig_hostBins")), "moc")
		} else {
			mocPath = filepath.Join(utils.QT_DIR(), "5.7", "gcc_64", "bin", "moc")
		}
	case "windows":
		if utils.UseMsys2() {
			mocPath = filepath.Join(utils.QT_MSYS2_DIR(), "bin", "moc")
		} else {
			mocPath = filepath.Join(utils.QT_DIR(), "5.7", "mingw53_32", "bin", "moc")
		}
	}

	utils.RunCmd(exec.Command(mocPath, filepath.Join(appPath, "moc.cpp"), "-o", filepath.Join(appPath, "moc_moc.h")), "execute moc")
	tmpFiles = append(tmpFiles, filepath.Join(appPath, "moc_moc.h"))

	return nil
}

func getParameters(tag string) []*parser.Parameter {
	var out = make([]*parser.Parameter, 0)

	if strings.Contains(tag, "(") {
		var (
			pairs     = strings.Split(strings.Split(strings.Split(tag, "(")[1], ")")[0], ",")
			lastValue string
		)

		for i := len(pairs) - 1; i >= 0; i-- {
			var (
				pairSplitted = strings.Split(strings.TrimSpace(pairs[i]), " ")
				pOut         *parser.Parameter
			)

			if len(pairSplitted) == 1 {
				pOut = &parser.Parameter{Name: fmt.Sprintf("v%v", i), Value: pairSplitted[0]}
				if lastValue != "" {
					pOut.Name = pOut.Value
					pOut.Value = lastValue
				}
			} else {
				pOut = &parser.Parameter{Name: pairSplitted[0], Value: pairSplitted[1]}
				lastValue = pOut.Value
			}

			out = append(out, pOut)
		}

		var rOut = make([]*parser.Parameter, 0)
		for i := len(out) - 1; i >= 0; i-- {
			rOut = append(rOut, out[i])
		}
		return rOut
	}

	return out
}

func getCppTypeFromGoType(t string) string {
	t = strings.TrimPrefix(t, "*")
	switch t {
	case "string":
		return "QString"
	case "[]string":
		return "QStringList"
	case "bool":
		return "bool"
	case "int8":
		return "qint8"
	case "uint8":
		return "quint8"
	case "int16":
		return "qint16"
	case "uint16":
		return "quint16"
	case "int", "int32":
		return "qint32"
	case "uint", "uint32":
		return "quint32"
	case "int64":
		return "qint64"
	case "uint64":
		return "quint64"
	case "float32":
		return "float"
	case "float64":
		return "qreal"
	case "uintptr":
		return "quintptr"
	case "unsafe.Pointer":
		return "void*"
	}

	if strings.Contains(t, ".") {
		t = strings.Split(t, ".")[1]
	}

	if strings.Contains(t, "__") {
		return strings.Replace(t, "_", ":", -1)
	}

	if _, exists := parser.ClassMap[t]; exists {
		if parser.ClassMap[t].IsQObjectSubClass() {
			return t + "*"
		}
		return t
	}

	//TODO: *_ITF support

	return "void"
}

func isBlacklisted(appPath, currentPath string) bool {
	for _, n := range blacklist {
		if strings.Contains(filepath.Join(currentPath), filepath.Join(appPath, n)) {
			return true
		}
	}
	return false
}

func isInClassArray(classes []*parser.Class, className string) bool {
	for _, c := range classes {
		if c.Name == className {
			return true
		}
	}
	return false
}

func isInClassMap(classes map[string]*parser.Class, className string) bool {
	for _, c := range classes {
		if c.Name == className {
			return true
		}
	}
	return false
}
