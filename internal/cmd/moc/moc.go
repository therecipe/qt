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

type appMoc struct {
	appPath     string
	buildTarget string
	module      *parser.Module
}

func newAppMoc(appPath, buildTarget string) *appMoc {
	return &appMoc{
		appPath:     appPath,
		buildTarget: buildTarget,
		module: &parser.Module{
			Project: parser.MOC,
			Namespace: &parser.Namespace{
				Classes: make([]*parser.Class, 0),
			},
		},
	}
}

func cacheModules() (err error) {

	if len(parser.State.ClassMap) != 0 {
		utils.Log.WithField("func", "cacheModules").Debug("modules already cached, skipping caching of modules")
		return
	}

	if len(parser.State.ClassMap) == 0 {
		parser.LoadModules()
	}

	return
}

// return how many moc classes are in module, delete those that are not
func (m *appMoc) cleanupClassMap() (size int) {
	for _, class := range m.module.Namespace.Classes {
		if _, failed := class.GetAllBasesRecursiveCheckFailed(0); failed || !class.IsSubClassOfQObject() {
			delete(parser.State.ClassMap, class.Name)
		} else {
			size++
		}
	}
	return
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

	parser.State.Module = file.Name.String()

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
				Access:     "public",
				Module:     parser.MOC,
				Name:       typeSpec.Name.String(),
				Status:     "public",
				Functions:  make([]*parser.Function, 0),
				Properties: make([]*parser.Variable, 0),
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
					case strings.HasPrefix(tag, "property:"):
						meta = parser.PROP
					case strings.HasPrefix(tag, "constructor:"):
						meta = parser.CONSTRUCTOR
					default:
						// only handle signal and slot and properties
						continue
					}

					var typ = string(src[field.Type.Pos()-1 : field.Type.End()-1])

					switch meta {
					case parser.SIGNAL, parser.SLOT:
						{
							var f = &parser.Function{
								Access:     "public",
								Fullname:   fmt.Sprintf("%v::%v", class.Name, strings.Split(tag, ":")[1]),
								Meta:       meta,
								Name:       strings.Split(tag, ":")[1],
								Status:     "public",
								Virtual:    parser.PURE,
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
								IsMocFunction: true,
							}

							if len(f.Parameters[0].Value) == 0 {
								f.Parameters = make([]*parser.Parameter, 0)
							}
							class.Functions = append(class.Functions, f)
						}

					case parser.PROP:
						{
							var p = &parser.Variable{
								Access:   "public",
								Fullname: fmt.Sprintf("%v::%v", class.Name, strings.Split(tag, ":")[1]),
								Name:     strings.Split(tag, ":")[1],
								Status:   "public",
								Output:   typ,
							}
							class.Properties = append(class.Properties, p)
						}

					case parser.CONSTRUCTOR:
						{
							class.Constructors = append(class.Constructors, strings.Split(tag, ":")[1])
						}
					}
				}

			}
			m.module.Namespace.Classes = append(m.module.Namespace.Classes, class)
		}
	}
	return nil
}

func CleanPath(path string) (err error) {
	var (
		tmpFileNames = []string{
			"moc.h", "moc.go", "moc.cpp", "moc_moc.h",
			"moc_cgo_darwin_darwin_amd64.go", "moc_cgo_windows_windows_386.go", "moc_cgo_windows_windows_amd64.go", "moc_cgo_linux_linux_amd64.go",
			"moc_cgo_android_linux_arm.go",
			"moc_cgo_ios_simulator_darwin_amd64.go", "moc_cgo_ios_simulator_darwin_386.go", "moc_cgo_ios_darwin_arm64.go", "moc_cgo_ios_darwin_arm.go",
			"moc_cgo_sailfish_emulator_linux_386.go", "moc_cgo_sailfish_linux_arm.go", "moc_cgo_asteroid_linux_arm.go",
			"moc_cgo_rpi1_linux_arm.go", "moc_cgo_rpi2_linux_arm.go", "moc_cgo_rpi3_linux_arm.go",
		}

		fields = logrus.Fields{"func": "CleanPath", "path": path}
	)

	pathInfo, err := os.Stat(path)
	if err != nil {
		utils.Log.WithFields(fields).Error("Invalid path")
		return err
	}
	if !pathInfo.IsDir() {
		utils.Log.WithFields(fields).Error("Path is not a directory")
		return errors.New("Path is not a directory")
	}

	err = filepath.Walk(
		path,

		utils.WalkOnlyFile(
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				for _, tFile := range tmpFileNames {
					if tFile == info.Name() {
						utils.Log.WithFields(fields).WithField("filename", tFile).Debug("Remove file")
						return utils.RemoveAll(path)
					}
				}

				return nil
			},
		),
	)

	if err != nil {
		utils.Log.WithFields(fields).WithError(err).Error("Failed to cleanup temp moc files")
	}

	return
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
		if m.buildTarget == "windows" || (utils.QT_DOCKER() && os.Getenv("QT_MXE_ARCH") != "") {
			var prefix = "i686"
			if utils.QT_MXE_ARCH() == "amd64" {
				prefix = "x86_64"
			}
			var suffix = "shared"
			if utils.QT_MXE_STATIC() {
				suffix = "static"
			}
			mocPath = filepath.Join("/usr", "lib", "mxe", "usr", fmt.Sprintf("%v-w64-mingw32.%v", prefix, suffix), "qt5", "bin", "moc")
		} else if utils.UsePkgConfig() {
			mocPath = filepath.Join(strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=host_bins", "Qt5Core"), "moc.LinuxPkgConfig_hostBins")), "moc")
		} else {
			mocPath = filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "gcc_64", "bin", "moc")
		}
	case "windows":
		if utils.UseMsys2() {
			mocPath = filepath.Join(utils.QT_MSYS2_DIR(), "bin", "moc")
		} else {
			mocPath = filepath.Join(utils.QT_DIR(), utils.QT_VERSION_MAJOR(), "mingw53_32", "bin", "moc")
		}
	}

	if runtime.GOOS != "windows" { //TODO: os.Stat fails on windows
		if info, err = os.Stat(mocPath); err != nil {
			return
		}
		if info.IsDir() {
			err = fmt.Errorf("%s is a directory", mocPath)
			return
		}
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
	if structsSize == 0 {
		utils.Log.WithFields(fields).Debug("failed to find moc structs")
		return nil
	}
	fields["structs"] = structsSize
	utils.Log.WithFields(fields).Debug("found moc structs")

	if err = cacheModules(); err != nil {
		return err
	}

	//find valid base classes for the moc classes in moc namespace and global namespace
	for _, c := range m.module.Namespace.Classes {
		for _, bc := range c.GetBases() {
			if isInClassArray(m.module.Namespace.Classes, bc) || isInClassMap(parser.State.ClassMap, bc) {
				c.Bases = bc
				break
			}
		}
	}

	m.module.Prepare()

	if m.cleanupClassMap() == 0 {
		utils.Log.WithFields(fields).Debug("failed to find at least one valid moc struct")
		return nil
		//return errors.New("failed to find at least one valid moc struct")
	}

	utils.Log.Debug("deducing go types to c++ types - start")

	for _, c := range m.module.Namespace.Classes {
		for _, f := range c.Functions {
			if !f.NoMocDeduce {
				for _, p := range f.Parameters {
					p.Value = getCppTypeFromGoType(f, p.Value)
				}
				f.Output = getCppTypeFromGoType(f, f.Output)
			}
		}
		utils.Log.Debug("deducing go types to c++ types - funcs done")
		for _, p := range c.Properties {
			p.Output = getCppTypeFromGoType(nil, p.Output)
		}
		utils.Log.Debug("deducing go types to c++ types - props done")

		//TODO: needed because only now the values are decuded to c++ types
		c.FixGenericHelper()
	}

	utils.Log.Debug("deducing go types to c++ types - done")

	//copy constructor and destructor
	for _ = range m.module.Namespace.Classes {
		for _, c := range m.module.Namespace.Classes {
			if bc, ok := parser.State.ClassMap[c.Bases]; ok {
				for _, bcf := range bc.Functions {
					if bcf.Meta == parser.CONSTRUCTOR || bcf.Meta == parser.DESTRUCTOR {
						var f = *bcf
						f.Name = strings.Replace(f.Name, bcf.ClassName(), c.Name, -1)
						f.Fullname = strings.Replace(f.Fullname, bcf.ClassName(), c.Name, -1)

						if !c.HasFunctionWithNameAndOverloadNumber(f.Name, f.OverloadNumber) {
							c.Functions = append(c.Functions, &f)
						}
					}
				}
			}
		}
	}

	if err = utils.SaveBytes(filepath.Join(m.appPath, "moc.cpp"), templater.CppTemplate(parser.MOC)); err != nil {
		return err
	}

	if err = utils.SaveBytes(filepath.Join(m.appPath, "moc.h"), templater.HTemplate(parser.MOC)); err != nil {
		return err
	}

	if err = utils.SaveBytes(filepath.Join(m.appPath, "moc.go"), templater.GoTemplate(parser.MOC, false)); err != nil {
		return err
	}
	templater.CgoTemplate(parser.MOC, m.appPath, m.buildTarget)

	for _, c := range parser.State.ClassMap {
		if c.Module == parser.MOC {
			delete(parser.State.ClassMap, c.Name)
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

func getCppTypeFromGoType(f *parser.Function, t string) string {
	//TODO: var tOld = t
	t = strings.TrimPrefix(t, "*")

	if strings.HasPrefix(t, "[]") && t != "[]string" {
		return fmt.Sprintf("QList<%v>", getCppTypeFromGoType(f, strings.TrimPrefix(t, "[]")))
	}

	if strings.HasPrefix(t, "map[") {
		var head = fmt.Sprintf("map[%v]", strings.Split(strings.TrimPrefix(t, "map["), "]")[0])
		return fmt.Sprintf("QHash<%v, %v>",
			getCppTypeFromGoType(f, strings.Split(strings.TrimPrefix(t, "map["), "]")[0]),
			getCppTypeFromGoType(f, strings.TrimPrefix(t, head)),
		)
	}

	if t == "error" && f != nil {
		f.AsError = true
	}

	switch t {
	case "string", "error":
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

	if _, ok := parser.State.ClassMap[t]; ok {
		if parser.State.ClassMap[t].IsSubClassOfQObject() /*TODO: || f == nil && strings.HasPrefix(tOLD, "*")*/ {
			return t + "*"
		}
		return t
	}

	//TODO: *_ITF support

	return "void"
}

// MocTree process an application and all it's sub-packages and create moc files
func MocTree(appPath, buildTarget string) (err error) {
	parser.State.Moc = true

	err = filepath.Walk(
		appPath,
		utils.WalkOnlyDirectory(
			utils.WalkFilterBlacklist(appPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				return newAppMoc(path, buildTarget).generate()
			}),
		),
	)

	return err
}
