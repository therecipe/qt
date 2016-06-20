package main

import (
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

var modulesInited bool

func main() {
	var appPath string

	switch len(os.Args) {
	case 1:
		{
			appPath, _ = os.Getwd()
		}

	default:
		{
			appPath = os.Args[1]
		}
	}

	if !filepath.IsAbs(appPath) {
		appPath = utils.GetAbsPath(appPath)
	}

	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
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
}

func moc(appPath string) {

	for _, name := range []string{"moc.h", "moc.go", "moc.cpp", "moc_moc.h",
		"moc_cgo_windows_386.go", "moc_cgo_linux_arm.go", "moc_cgo_linux_amd64.go", "moc_cgo_linux_386.go",
		"moc_cgo_darwin_arm64.go", "moc_cgo_darwin_arm.go", "moc_cgo_darwin_arm.go", "moc_cgo_darwin_amd64.go",
		"moc_cgo_darwin_386.go", "moc_cgo_android_arm.go"} {
		if utils.Exists(filepath.Join(appPath, name)) {
			utils.RemoveAll(filepath.Join(appPath, name))
		}
	}

	var module = &parser.Module{Project: parser.MOC, Namespace: &parser.Namespace{Classes: make([]*parser.Class, 0)}}

	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if !strings.HasPrefix(info.Name(), "moc") && strings.HasSuffix(info.Name(), ".go") && filepath.Dir(path) == appPath && !info.IsDir() {

			src, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			file, err := goparser.ParseFile(token.NewFileSet(), path, nil, 0)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
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

									for _ = range field.Names {

										var _type = string(src[field.Type.Pos()-1 : field.Type.End()-1])
										var tag = ""
										if field.Tag != nil {
											tag = field.Tag.Value
										}

										tag = strings.Replace(tag, "\"", "", -1)
										tag = strings.Replace(tag, "`", "", -1)

										var meta string
										if strings.Contains(tag, "signal:") {
											meta = "signal"
										}
										if strings.Contains(tag, "slot:") {
											meta = "slot"
										}

										if meta != "" {
											var (
												name = strings.Split(tag, ":")[1]
												f    = &parser.Function{Access: "public", Fullname: class.Name + "::" + name, Meta: meta, Name: name, Output: "void", Status: "public", Virtual: "non", Signature: "()"}
											)
											f.Parameters = getParameters(_type)
											if f.Meta == "slot" {
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

		if !modulesInited {
			for _, module := range templater.GetLibs() {
				parser.GetModule(strings.ToLower(module))
			}
			modulesInited = true
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
						if bcf.Meta == "constructor" || bcf.Meta == "destructor" {
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
			utils.SaveBytes(filepath.Join(appPath, "moc.cpp"), templater.CppTemplate(parser.MOC))
			utils.SaveBytes(filepath.Join(appPath, "moc.h"), templater.HTemplate(parser.MOC))
			utils.SaveBytes(filepath.Join(appPath, "moc.go"), templater.GoTemplate(parser.MOC, false))

			var mocPath string

			switch runtime.GOOS {
			case "darwin":
				{
					mocPath = "/usr/local/Qt5.7.0/5.7/clang_64/bin/moc"
				}

			case "linux":
				{
					mocPath = "/usr/local/Qt5.7.0/5.7/gcc_64/bin/moc"
				}

			case "windows":
				{
					mocPath = "C:\\Qt\\Qt5.7.0\\5.7\\mingw53_32\\bin\\moc.exe"
				}
			}

			var moc = exec.Command(mocPath)
			moc.Args = append(moc.Args,
				filepath.Join(appPath, "moc.cpp"),
				"-o", filepath.Join(appPath, "moc_moc.h"))
			moc.Dir = filepath.Join(appPath)
			runCmd(moc, "moc.moc")

			var gofmt = exec.Command("go", "fmt")
			gofmt.Dir = appPath
			runCmd(gofmt, "moc.fmt")

			templater.MocAppPath = appPath
			templater.CopyCgo(parser.MOC)
		}
	}
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

	case "unsafe.Pointer", "uintptr":
		{
			return "void*"
		}

	case "bool":
		{
			return "bool"
		}

	case "int":
		{
			return "int"
		}

	case "float64":
		{
			return "qreal"
		}

	case "int64":
		{
			return "qint64"
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

func runCmd(cmd *exec.Cmd, n string) {
	var out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", n, out, err)
		os.Exit(1)
	}
}
