package moc

import (
	"errors"
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

var temporaryFiles = []string{
	"moc.h", "moc.go", "moc.cpp", "moc_moc.h",
	"moc_cgo_desktop_darwin_amd64.go", "moc_cgo_desktop_windows_386.go", "moc_cgo_desktop_windows_amd64.go", "moc_cgo_desktop_linux_amd64.go",
	"moc_cgo_android_linux_arm.go",
	"moc_cgo_ios_simulator_darwin_amd64.go", "moc_cgo_ios_simulator_darwin_386.go", "moc_cgo_ios_darwin_arm64.go", "moc_cgo_ios_darwin_arm.go",
	"moc_cgo_sailfish_emulator_linux_386.go", "moc_cgo_sailfish_linux_arm.go",
	"moc_cgo_rpi1_linux_arm.go", "moc_cgo_rpi2_linux_arm.go", "moc_cgo_rpi3_linux_arm.go",
}

func init() {
	// TODO: dangerous
	templater.Moc = true
}

type appMoc struct {
	appPath string
	module  *parser.Module
}

func newAppMoc(appPath string) *appMoc {
	return &appMoc{
		//tmpFiles: make([]string, 0),
		appPath: appPath,
		module: &parser.Module{
			Project:   parser.MOC,
			Namespace: &parser.Namespace{
				Classes: make([]*parser.Class, 0),
			},
		},
	}
}

func cacheModules() (err error) {
	fields := logrus.Fields{
		"func": "cacheModules",
	}
	if len(parser.ClassMap) == 0 {
		for _, module := range templater.GetLibs() {
			modName := strings.ToLower(module)
			if _, err = parser.GetModule(modName); err != nil {
				return
			}
		}
	} else {
		utils.Log.WithFields(fields).Warn("empty ClassMap")
	}
	return
}

// return how many moc classes are in module, delete those that are not
func (m *appMoc) cleanupClassMap() (size int) {
	for _, class := range m.module.Namespace.Classes {
		if !class.IsQObjectSubClass() {
			delete(parser.ClassMap, class.Name)
		} else {
			size++
		}
	}
	return size
}

func (m *appMoc) parseGo(path string) error {
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
			m.module.Namespace.Classes = append(m.module.Namespace.Classes, class)
		}
	}
	return nil
}

func (m *appMoc) cleanup() (err error) {
	fields := logrus.Fields{"func": "Moc.cleanup"}
	files := append(temporaryFiles, "moc_cleanup.json")
	for _, n := range files {
		absoluteFilename := filepath.Join(m.appPath, n)
		if utils.Exists(absoluteFilename) {
			utils.Log.WithFields(fields).WithField("filename", n).Debug("File already exists, cleanup")
			if err = os.Remove(absoluteFilename); err != nil {
				return
			}
		}
	}
	return nil
}

func (m *appMoc) runQtMoc() (err error) {
	var (
		mocPath string
		output  []byte
		info    os.FileInfo
	)
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

	if info, err = os.Stat(mocPath); err != nil {
		return
	}
	if info.IsDir() {
		err = fmt.Errorf("%s is a directory", mocPath)
		return
	}

	if err = os.Chdir(m.appPath); err != nil {
		return
	}

	cmd := exec.Command(mocPath, filepath.Join(m.appPath, "moc.cpp"), "-o", filepath.Join(m.appPath, "moc_moc.h"))
	fields := logrus.Fields{
		"func": "runQtMoc",
		"cmd":  strings.Join(cmd.Args, " "),
	}
	utils.Log.WithFields(fields).Debug("Execute moc")
	if output, err = cmd.CombinedOutput(); err != nil {
		utils.Log.WithError(err).WithFields(fields).WithField("output", string(output)).Error("failed to run command")
	}
	return
}

func (m *appMoc) generate() error {
	files, err := ioutil.ReadDir(m.appPath)
	if err != nil {
		return err
	}
	fields := logrus.Fields{
		"func": "appMoc.generate",
	}

	for _, info := range files {
		filename := filepath.Join(m.appPath, info.Name())
		loopFields := logrus.Fields{"filename": filename}
		if info.IsDir() {
			utils.Log.WithFields(fields).WithFields(loopFields).Debug("Skip directory")
			continue
		}

		if filepath.Ext(filename) != ".go" {
			utils.Log.WithFields(fields).WithFields(loopFields).Debug("Skip non-go source")
			continue
		}

		if strings.HasPrefix(filename, "moc") {
			utils.Log.WithFields(fields).WithFields(loopFields).Debug("Skip moc output")
			continue
		}

		utils.Log.WithFields(fields).WithFields(loopFields).Debug("Process source")
		if err = m.parseGo(filename); err != nil {
			return err
		}
	}

	structsSize := len(m.module.Namespace.Classes)
	if structsSize  == 0 {
		utils.Log.WithFields(fields).Warning("failed to find moc structs")
		return nil
	}
	fields["structs"] = structsSize
	utils.Log.WithFields(fields).Debug("Found moc structs")

	if err = cacheModules(); err != nil {
		return err
	}

	// ?
	for _, c := range m.module.Namespace.Classes {
		for _, bc := range c.GetBases() {
			if isInClassArray(m.module.Namespace.Classes, bc) || isInClassMap(parser.ClassMap, bc) {
				c.Bases = bc
				break
			}
		}
	}

	m.module.Prepare()

	if m.cleanupClassMap() == 0 {
		//utils.Log.WithFields(fields).Debug("failed to find at least one valid moc struct")
		return errors.New("failed to find at least one valid moc struct")
	}

	for _, c := range m.module.Namespace.Classes {
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
	for _ = range m.module.Namespace.Classes {
		for _, c := range m.module.Namespace.Classes {
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

	if err = ioutil.WriteFile(filepath.Join(m.appPath, "moc.cpp"), templater.CppTemplate(parser.MOC), 0644); err != nil {
		return err
	}

	if err = ioutil.WriteFile(filepath.Join(m.appPath, "moc.h"), templater.HTemplate(parser.MOC), 0644); err != nil {
		return err
	}

	if err = ioutil.WriteFile(filepath.Join(m.appPath, "moc.go"), templater.GoTemplate(parser.MOC, false), 0644); err != nil {
		return err
	}
	templater.CopyCgoForMoc(parser.MOC, m.appPath)

	for _, c := range parser.ClassMap {
		if c.Module == parser.MOC {
			delete(parser.ClassMap, c.Name)
		}
	}

	return m.runQtMoc()
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

// MocTree process an application and all it's sub-packages and create moc files
func MocTree(appPath string) error {
	return filepath.Walk(
		appPath,
		utils.WalkOnlyDirectory(
			utils.WalkFilterBlacklist(appPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				am := newAppMoc(path)
				if cErr := am.cleanup(); cErr != nil {
					return cErr
				}
				return am.generate()
			}),
		),
	)
}
