package main

import (
	"encoding/json"
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

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"
	"github.com/therecipe/qt/internal/utils"
)

func init() {
	templater.UsedFromMoc = true
}

var tmpFiles = make([]string, 0)

func main() {
	var (
		appPath string
		cleanup bool
	)

	switch len(os.Args) {
	case 1:
		{
			appPath, _ = os.Getwd()
		}

	case 2, 3:
		{
			appPath = os.Args[1]
			cleanup = len(os.Args) == 3
		}
	}
	if !filepath.IsAbs(appPath) {
		appPath = utils.GetAbsPath(appPath)
	}

	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !isBlacklisted(appPath, filepath.Join(path, info.Name())) {
			moc(path)

			for className, class := range parser.ClassMap {
				if class.Module == parser.MOC {
					delete(parser.ClassMap, className)
				}
			}
		}

		return nil
	}

	filepath.Walk(appPath, walkFunc)

	if cleanup {
		var b, err = json.Marshal(tmpFiles)
		if err == nil {
			utils.SaveBytes(filepath.Join(appPath, "moc_cleanup.json"), b)
		}
	}
}

func moc(appPath string) {

	for _, name := range []string{"moc.h", "moc.go", "moc.cpp", "moc_moc.h",
		"moc_cgo_desktop_darwin_amd64.go", "moc_cgo_desktop_windows_386.go", "moc_cgo_desktop_linux_amd64.go",
		"moc_cgo_android_linux_arm.go",
		"moc_cgo_ios_simulator_darwin_amd64.go", "moc_cgo_ios_simulator_darwin_386.go", "moc_cgo_ios_darwin_arm64.go", "moc_cgo_ios_darwin_arm.go",
		"moc_cgo_sailfish_emulator_linux_386.go", "moc_cgo_sailfish_linux_arm.go",
		"moc_cgo_rpi1_linux_arm.go", "moc_cgo_rpi2_linux_arm.go", "moc_cgo_rpi3_linux_arm.go"} {
		if utils.Exists(filepath.Join(appPath, name)) {
			utils.RemoveAll(filepath.Join(appPath, name))
		}
		tmpFiles = append(tmpFiles, filepath.Join(appPath, name))
	}
	utils.RemoveAll(filepath.Join(appPath, "moc_cleanup.json"))

	var module = &parser.Module{Project: parser.MOC, Namespace: &parser.Namespace{Classes: make([]*parser.Class, 0)}}

	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if !strings.HasPrefix(info.Name(), "moc") && strings.HasSuffix(info.Name(), ".go") && filepath.Dir(path) == appPath && !info.IsDir() {

			src, err := ioutil.ReadFile(path)
			if err != nil {
				utils.Log.WithError(err).Panicf("failed to load file %v", path)
			}

			file, err := goparser.ParseFile(token.NewFileSet(), path, nil, 0)
			if err != nil {
				utils.Log.WithError(err).Panicf("failed to parse file %v", path)
			}

			if !strings.Contains(string(src), "package main") {
				var plist = strings.Split(filepath.Clean(path), string(filepath.Separator))
				templater.MocModule = plist[len(plist)-2]
			}

			for _, d := range file.Decls {
				if typeDecl, ok := d.(*ast.GenDecl); ok {
					for _, s := range typeDecl.Specs {
						if typeSpec, ok := s.(*ast.TypeSpec); ok {

							var (
								class     = &parser.Class{Access: "public", Module: parser.MOC, Name: typeSpec.Name.String(), Status: "public", Functions: make([]*parser.Function, 0)}
								skipClass = true
							)

							if structDecl, ok := typeSpec.Type.(*ast.StructType); ok {
								for _, field := range structDecl.Fields.List {

									var fieldValue = string(src[field.Pos()-1 : field.End()-1])

									if !strings.Contains(fieldValue, " ") && fieldValue != "" && class.Bases == "" {
										if strings.Contains(fieldValue, ".") {
											class.Bases = strings.Split(fieldValue, ".")[1]
										} else {
											class.Bases = fieldValue
										}
										skipClass = strings.HasPrefix(fieldValue, "*")
									}

									for range field.Names {

										var _type = string(src[field.Type.Pos()-1 : field.Type.End()-1])
										var tag = ""
										if field.Tag != nil {
											tag = field.Tag.Value
										}

										tag = strings.Replace(tag, "\"", "", -1)
										tag = strings.Replace(tag, "`", "", -1)

										var meta string
										if strings.Contains(tag, "signal:") {
											meta = parser.SIGNAL
										}
										if strings.Contains(tag, "slot:") {
											meta = parser.SLOT
										}

										if meta != "" {
											var (
												name = strings.Replace(strings.Split(tag, ":")[1], "\"", "", -1)
												f    = &parser.Function{Access: "public", Fullname: class.Name + "::" + name, Meta: meta, Name: name, Output: "void", Status: "public", Virtual: "non", Signature: "()"}
											)
											f.Parameters = getParameters(_type)
											if f.Meta == parser.SLOT {
												f.Output = getCppTypeFromGoType(strings.TrimSpace(strings.Split(_type, ")")[1]))
											}
											class.Functions = append(class.Functions, f)
										}
									}
								}
							}

							if !skipClass {
								module.Namespace.Classes = append(module.Namespace.Classes, class)
							}
						}
					}
				}
			}
		}

		return nil
	}

	filepath.Walk(appPath, walkFunc)

	if len(module.Namespace.Classes) > 0 {

		var importedPkgMap = make(map[string]bool)

		var walkFuncImports = func(appPath string, info os.FileInfo, err error) error {
			if err == nil && !strings.HasPrefix(info.Name(), "moc") && !strings.HasPrefix(info.Name(), "rcc") && strings.HasSuffix(info.Name(), ".go") && !info.IsDir() {
				var pFile, errParse = goparser.ParseFile(token.NewFileSet(), appPath, nil, 0)
				if errParse != nil {
					utils.Log.WithError(errParse).Panicf("failed to parser file %v", appPath)
				} else {
					for _, i := range pFile.Imports {
						if iPath := strings.Replace(i.Path.Value, "\"", "", -1); iPath != "github.com/therecipe/qt" {
							var pkg string
							for _, pkgs := range templater.Libs {
								if strings.ToLower(pkgs) == strings.TrimPrefix(iPath, "github.com/therecipe/qt/") {
									pkg = pkgs
								}
							}
							if _, exists := templater.LibDeps[pkg]; exists {
								importedPkgMap[pkg] = true
								for _, dep := range templater.LibDeps[pkg] {
									if _, exists := templater.LibDeps[dep]; exists {
										importedPkgMap[dep] = true
									}
								}
							}
						}
					}
				}
			}
			return nil
		}
		filepath.Walk(appPath, walkFuncImports)

		var importedPkgs []string
		for _, dep := range templater.Libs {
			if _, exist := importedPkgMap[dep]; exist {
				importedPkgs = append(importedPkgs, dep)
			}
		}

		for _, module := range importedPkgs {
			utils.Log.Debugf("loading qt/%v", strings.ToLower(module))
			if _, err := parser.GetModule(strings.ToLower(module)); err != nil {
				utils.Log.WithError(err).Errorf("failed to load qt/%v", strings.ToLower(module))
			}
		}

		module.Prepare()

		for _, c := range parser.ClassMap {
			if c.Module == parser.MOC {
				for _, f := range c.Functions {
					for _, p := range f.Parameters {
						p.Value = getCppTypeFromGoType(p.Value)
					}
				}
			}
		}

		for i := 0; i <= len(module.Namespace.Classes); i++ {
			for _, c := range module.Namespace.Classes {
				if bc, exists := parser.ClassMap[c.Bases]; exists {
					for _, bcf := range bc.Functions {
						if bcf.Meta == parser.CONSTRUCTOR || bcf.Meta == parser.DESTRUCTOR {
							var f = *bcf
							f.Fullname = strings.Replace(f.Fullname, bcf.Class(), c.Name, -1)
							f.Name = strings.Replace(f.Name, bcf.Class(), c.Name, -1)

							var exists bool
							for _, cf := range c.Functions {
								if cf.Fullname == f.Fullname {
									exists = true
								}
							}
							if !exists {
								c.Functions = append(c.Functions, &f)
							}
						}
					}
				}
			}
		}

		for _, c := range parser.ClassMap {
			if c.Module == parser.MOC {
				if !c.IsQObjectSubClass() {
					delete(parser.ClassMap, c.Name)
				}
			}
		}

		var classCount int
		for _, class := range parser.ClassMap {
			if class.Module == parser.MOC {
				classCount++
			}
		}

		if classCount > 0 {
			utils.Log.Debugf("generating %v moc", parser.MOC)
			utils.SaveBytes(filepath.Join(appPath, "moc.cpp"), templater.CppTemplate(parser.MOC))
			utils.SaveBytes(filepath.Join(appPath, "moc.h"), templater.HTemplate(parser.MOC))
			utils.SaveBytes(filepath.Join(appPath, "moc.go"), templater.GoTemplate(parser.MOC, false))

			tmpFiles = append(tmpFiles, filepath.Join(appPath, "moc.cpp"))
			tmpFiles = append(tmpFiles, filepath.Join(appPath, "moc.h"))
			tmpFiles = append(tmpFiles, filepath.Join(appPath, "moc.go"))

			var mocPath string
			switch runtime.GOOS {
			case "darwin":
				{
					mocPath = filepath.Join(utils.QT_DARWIN_DIR(), "bin", "moc")
				}

			case "linux":
				{
					if utils.UsePkgConfig() {
						mocPath = filepath.Join(strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=host_bins", "Qt5Core"), "moc.LinuxPkgConfig_hostBins")), "moc")
					} else {
						mocPath = filepath.Join(utils.QT_DIR(), "5.7", "gcc_64", "bin", "moc")
					}
				}

			case "windows":
				{
					mocPath = filepath.Join(utils.QT_DIR(), "5.7", "mingw53_32", "bin", "moc")
				}
			}

			var moc = exec.Command(mocPath)
			moc.Args = append(moc.Args,
				filepath.Join(appPath, "moc.cpp"),
				"-o", filepath.Join(appPath, "moc_moc.h"))
			moc.Dir = filepath.Join(appPath)
			utils.RunCmd(moc, "moc.moc")

			tmpFiles = append(tmpFiles, filepath.Join(appPath, "moc_moc.h"))

			var gofmt = exec.Command("go", "fmt")
			gofmt.Dir = appPath
			utils.RunCmd(gofmt, "moc.fmt")

			templater.MocAppPath = appPath
			templater.CopyCgo(parser.MOC)
		}
	}

	parser.ClassMap = make(map[string]*parser.Class)
}

func getParameters(tag string) []*parser.Parameter {
	var out = make([]*parser.Parameter, 0)

	if strings.Contains(tag, "(") {
		var (
			sig       = strings.Split(strings.Split(tag, "(")[1], ")")[0]
			pairs     = strings.Split(sig, ",")
			lastValue string
		)

		for i := len(pairs) - 1; i >= 0; i-- {
			var (
				pairSplitted = strings.Split(strings.TrimSpace(pairs[i]), " ")
				p            *parser.Parameter
			)

			if len(pairSplitted) == 1 {
				p = &parser.Parameter{Name: fmt.Sprintf("v%v", i), Value: pairSplitted[0]}
				if getCppTypeFromGoType(p.Value) == "void" {
					p.Name = p.Value
					p.Value = lastValue
				}
			} else {
				p = &parser.Parameter{Name: pairSplitted[0], Value: pairSplitted[1]}
				lastValue = p.Value
			}

			if p.Name == "" && p.Value == "" {
			} else {
				out = append(out, p)
			}
		}

		var reverseOut = make([]*parser.Parameter, 0)
		for i := len(out) - 1; i >= 0; i-- {
			reverseOut = append(reverseOut, out[i])
		}
		return reverseOut
	}

	return out
}

func getCppTypeFromGoType(t string) string {
	t = strings.TrimPrefix(t, "*")
	switch t {
	case "string":
		{
			return "QString"
		}

	case "[]string":
		{
			return "QStringList"
		}

	case "bool":
		{
			return "bool"
		}

	case "int16":
		{
			return "qint16"
		}

	case "uint16":
		{
			return "quint16"
		}

	case "int", "int32":
		{
			return "qint32"
		}

	case "uint", "uint32":
		{
			return "quint32"
		}

	case "int64":
		{
			return "qint64"
		}

	case "uint64":
		{
			return "quint64"
		}

	case "float32":
		{
			return "float"
		}

	case "float64":
		{
			return "qreal"
		}

	case "uintptr":
		{
			return "quintptr"
		}

	case "unsafe.Pointer":
		{
			return "void*"
		}
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

	return "void"
}

func isBlacklisted(appPath, currentPath string) bool {

	for _, n := range []string{"deploy", "qml", "android", "ios", "ios-simulator", "sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3"} {
		if strings.Contains(filepath.Join(currentPath), filepath.Join(appPath, n)) {
			return true
		}
	}

	return false
}
