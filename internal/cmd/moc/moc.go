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
	"sync"

	"golang.org/x/tools/imports"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/cmd"

	"github.com/therecipe/qt/internal/utils"
)

var (
	done   = make(map[string]struct{})
	rootWg = new(sync.WaitGroup)

	goFilesCache      = make(map[string][]string)
	goFilesCacheMutex = new(sync.Mutex)

	goImportsCache      = make(map[string][]string)
	goImportsCacheMutex = new(sync.Mutex)
)

func Moc(path, target, tags string, fast, slow bool) {
	moc(path, target, tags, fast, slow, true)
}

func moc(path, target, tags string, fast, slow, root bool) {
	utils.Log.WithField("path", path).WithField("target", target).Debug("start Moc")

	if root {
		ngr := 50
		if slow {
			ngr = 1
		}
		defer rootWg.Wait()
		wg := new(sync.WaitGroup)
		wc := make(chan bool, ngr)
		allImports := append([]string{path}, cmd.GetImports(path, target, tags, 0, false)...)

		wg.Add(len(allImports) * 2)
		for _, path := range allImports {
			wc <- true
			go func(path string) {
				files := cmd.GetGoFiles(path, target, tags)
				goFilesCacheMutex.Lock()
				goFilesCache[path] = files
				goFilesCacheMutex.Unlock()
				<-wc
				wg.Done()
			}(path)

			wc <- true
			go func(path string) {
				imports := cmd.GetImports(path, target, tags, 0, true)
				goImportsCacheMutex.Lock()
				goImportsCache[path] = imports
				goImportsCacheMutex.Unlock()
				<-wc
				wg.Done()
			}(path)
		}
		wg.Wait()
	}

	var classes []*parser.Class
	var otherclasses []*parser.Class
	var pkg string

	for i, path := range append([]string{path}, goImportsCache[path]...) {
		if _, ok := done[path]; !ok && i > 0 && !fast {
			done[path] = struct{}{}
			moc(path, target, tags, false, slow, false)
		}

		for _, fpath := range goFilesCache[path] {
			cls, ipkg, err := parse(fpath)
			if pkg == "" {
				pkg = ipkg
			}
			if err != nil {
				utils.Log.WithError(err)
			} else {
				if cls == nil {
					continue
				}

				if pkg == ipkg {
					classes = append(classes, cls...)
				} else {

					m := &parser.Module{
						Project:   "custom_" + ipkg,
						Pkg:       path,
						Namespace: &parser.Namespace{Classes: cls},
					}
					m.Prepare()

					otherclasses = append(otherclasses, cls...)
				}
			}
		}
	}

	c := len(classes)
	utils.Log.WithField("path", path).Debugln("found", c, "moc structs")
	if c == 0 {
		return
	}

	if _, ok := parser.State.ClassMap["QObject"]; !ok {
		parser.LoadModules()
	} else {
		utils.Log.Debug("modules already cached")
	}

	//find valid base classes
	for _, c := range classes {
		for _, bc := range c.GetBases() {
			var has bool
			for _, c := range append(classes, otherclasses...) {
				if c.Name == bc {
					has = true
					break
				}
			}
			if !has {
				for _, c := range parser.State.ClassMap {
					if c.Name != bc {
						has = true
						break
					}
				}
			}
			if has {
				c.Bases = bc
				break
			}
		}
	}

	//TODO: internal functions rely on the prepared state
	m := &parser.Module{
		Project:   parser.MOC,
		Namespace: &parser.Namespace{Classes: classes},
	}

	//copy properties + signals + slots
	utils.Log.Debug("start copy properties + signals + slots")
	for _ = range append(m.Namespace.Classes, otherclasses...) {
		for _, c := range append(m.Namespace.Classes, otherclasses...) {
			bc, ok := parser.State.ClassMap[c.Bases]
			if !ok || bc.Pkg == "" {
				continue
			}

			for _, bcp := range bc.Properties {
				var has bool
				for _, cp := range c.Properties {
					if cp.Name == bcp.Name {
						has = true
						break
					}
				}
				if !has {
					bcp.IsMocSynthetic = true
					c.Properties = append(c.Properties, bcp)
				}
			}

			for _, bcf := range bc.Functions {
				if !(bcf.Meta == parser.SIGNAL || bcf.Meta == parser.SLOT) {
					continue
				}

				f := *bcf
				f.Name = strings.Replace(f.Name, bcf.ClassName(), c.Name, -1)
				f.Fullname = strings.Replace(f.Fullname, bcf.ClassName(), c.Name, -1)
				if !c.HasFunctionWithNameAndOverloadNumber(f.Name, f.OverloadNumber) {
					c.Functions = append(c.Functions, &f)
				}
			}
		}
	}
	utils.Log.Debug("done copy properties + signals + slots")

	m.Prepare()

	var remaining int
	for _, class := range m.Namespace.Classes {
		if _, failed := class.GetAllBasesRecursiveCheckFailed(0); failed || !class.IsSubClassOfQObject() {
			delete(parser.State.ClassMap, class.Name)
		} else {
			remaining++
		}
	}
	utils.Log.WithField("path", path).Debugln("found", remaining, "remaining moc structs")
	if remaining == 0 {
		return
	}

	utils.Log.Debug("start converting types")
	for _, c := range m.Namespace.Classes {
		for _, f := range c.Functions {
			if f.NoMocDeduce {
				continue
			}
			for _, p := range f.Parameters {
				p.Value, p.PureGoType = cppTypeFromGoType(f, p.Value)
			}
			f.Output, f.PureGoOutput = cppTypeFromGoType(f, f.Output)
		}
		utils.Log.Debug("done converting func types")
		for _, p := range c.Properties {
			p.Output, p.PureGoType = cppTypeFromGoType(nil, p.Output)
		}
		utils.Log.Debug("done converting property types")

		//TODO: needed because only now the values are decuded to c++ types
		c.FixGenericHelper()
	}
	utils.Log.Debug("done converting types")

	//copy constructor and destructor
	utils.Log.Debug("start copy structors")
	for !hasStructors(m) {
		for _, c := range append(m.Namespace.Classes, otherclasses...) {
			bc, ok := parser.State.ClassMap[c.Bases]
			if !ok {
				continue
			}
			for _, bcf := range bc.Functions {
				if !(bcf.Meta == parser.CONSTRUCTOR || bcf.Meta == parser.DESTRUCTOR) {
					continue
				}
				f := *bcf
				f.Name = strings.Replace(f.Name, bcf.ClassName(), c.Name, -1)
				f.Fullname = strings.Replace(f.Fullname, bcf.ClassName(), c.Name, -1)
				if !c.HasFunctionWithNameAndOverloadNumber(f.Name, f.OverloadNumber) {
					c.Functions = append(c.Functions, &f)
				}
			}
		}
	}
	utils.Log.Debug("done copy structors")

	goFilesCacheMutex.Lock()
	files := goFilesCache[path]
	goFilesCacheMutex.Unlock()
	for _, fpath := range files {
		if filepath.Base(fpath) != "moc.go" {
			parseNonMocDeps(fpath)
		}
	}

	if err := utils.SaveBytes(filepath.Join(path, "moc.cpp"), templater.CppTemplate(parser.MOC, templater.MOC, target, tags)); err != nil {
		return
	}
	if err := utils.SaveBytes(filepath.Join(path, "moc.h"), templater.HTemplate(parser.MOC, templater.MOC, tags)); err != nil {
		return
	}
	fix := templater.GoTemplate(parser.MOC, false, templater.MOC, pkg, target, tags)

	rootWg.Add(1)
	go func() {
		defer rootWg.Done()
		var err error
		for i := 0; i < 5; i++ {
			fix, err = imports.Process(filepath.Join(path, "moc.go"), fix, nil)
			if err != nil {
				utils.Log.WithError(err).Fatal("failed to fix go imports")
				return
			}
		}
		if err := utils.SaveBytes(filepath.Join(path, "moc.go"), fix); err != nil {
			return
		}
	}()

	rootWg.Add(1)
	go func(libs []string) {
		templater.CgoTemplateSafe(parser.MOC, path, target, templater.MOC, pkg, tags, libs)
		rootWg.Done()
	}(parser.LibDeps[parser.MOC])

	//TODO: cleanup state -->
	for _, c := range parser.State.ClassMap {
		if c.Module == parser.MOC {
			delete(parser.State.ClassMap, c.Name)
		}
	}
	parser.LibDeps[parser.MOC] = make([]string, 0)
	//<--

	rootWg.Add(1)
	go func() {
		utils.RunCmd(exec.Command(utils.ToolPath("moc", target), filepath.Join(path, "moc.cpp"), "-o", filepath.Join(path, "moc_moc.h")), "run moc")
		if tags != "" {
			utils.Save(filepath.Join(path, "moc_moc.h"), "// +build "+tags+"\n\n"+utils.Load(filepath.Join(path, "moc_moc.h")))
		}

		if utils.QT_DOCKER() {
			if idug, ok := os.LookupEnv("IDUG"); ok {
				utils.RunCmd(exec.Command("chown", "-R", idug, path), "chown files to user")
			}
		}
		rootWg.Done()
	}()
}

func parse(path string) ([]*parser.Class, string, error) {
	utils.Log.WithField("path", path).Debug("parse")

	if strings.HasPrefix(path, filepath.Join(runtime.GOROOT(), "src")) {
		return nil, "", errors.New("path is in GOROOT/src")
	}

	src, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, "", err
	}

	file, err := goparser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		return nil, "", err
	}

	var classes []*parser.Class
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
				Access: "public",
				Module: parser.MOC,
				Name:   typeSpec.Name.String(),
				Status: "public",
			}

			//collect possible base classes
			var bases []string
			for _, field := range structDecl.Fields.List {
				value := string(src[field.Pos()-1 : field.End()-1])
				if value != "" && !strings.Contains(value, " ") && !strings.HasPrefix(value, "*") {
					if strings.Contains(value, ".") {
						value = strings.Split(value, ".")[1]
					}
					bases = append(bases, value)
				}

				if len(field.Names) > 0 {
					if field.Tag == nil {
						continue
					}
					tag := strings.Replace(strings.Replace(field.Tag.Value, "\"", "", -1), "`", "", -1)

					var meta string
					switch {
					case strings.HasPrefix(tag, "signal:"):
						meta = parser.SIGNAL
					case strings.HasPrefix(tag, "slot:"):
						meta = parser.SLOT
					case strings.HasPrefix(tag, "property:"):
						meta = parser.PROP
					case strings.HasPrefix(tag, "constructor:"): //TODO: more advanced constructor support (multiple constructors, custom inputs, error output, custom naming, ...)
						meta = parser.CONSTRUCTOR
					default:
						continue
					}

					switch typ := string(src[field.Type.Pos()-1 : field.Type.End()-1]); meta {
					case parser.SIGNAL, parser.SLOT:
						f := &parser.Function{
							Access:        "public",
							Fullname:      fmt.Sprintf("%v::%v", class.Name, strings.Split(tag, ":")[1]),
							Meta:          meta,
							Name:          strings.Split(tag, ":")[1],
							Status:        "public",
							Virtual:       parser.PURE,
							Signature:     "()",
							Output:        "void",
							Parameters:    parameters(typ),
							IsMocFunction: true,
						}
						if meta == parser.SLOT {
							if strings.Contains(typ, ") ") {
								f.Output = strings.Split(typ, ") ")[1]
							} else if strings.Contains(typ, ")") {
								f.Output = strings.Split(typ, ")")[1]
							}
							if strings.HasPrefix(f.Output, "(") {
								f.Output = strings.TrimPrefix(f.Output, "(")
								f.Output = strings.TrimSuffix(f.Output, ")")
								f.Output = strings.Join(strings.Split(f.Output, " ")[1:], " ")
							}
						}
						if len(f.Parameters[0].Value) == 0 {
							f.Parameters = make([]*parser.Parameter, 0)
						}
						class.Functions = append(class.Functions, f)

					case parser.PROP:
						class.Properties = append(class.Properties,
							&parser.Variable{
								Access:   "public",
								Fullname: fmt.Sprintf("%v::%v", class.Name, strings.Split(tag, ":")[1]),
								Name:     strings.Split(tag, ":")[1],
								Status:   "public",
								Output:   typ,
							})

					case parser.CONSTRUCTOR:
						class.Constructors = append(class.Constructors, strings.Split(tag, ":")[1])
					}

				}
			}
			class.Bases = strings.Join(bases, ",")
			classes = append(classes, class)
		}
	}

	return classes, file.Name.String(), nil
}

func parameters(tag string) []*parser.Parameter {
	if !strings.Contains(tag, "(") {
		return nil
	}

	tag = strings.TrimPrefix(tag, "func(")

	switch {
	case strings.Contains(tag, ") ("):
		tag = strings.Split(tag, ") (")[0]

	case strings.Contains(tag, ") func"):
		tag = strings.Split(tag, ") func")[0]

	case strings.Contains(tag, ") "):
		tag = strings.Split(tag, ") ")[0]

	default:
		tag = strings.TrimSuffix(tag, ")")
	}

	tag = strings.Replace(tag, " func", ";func", -1)
	tag = strings.Replace(tag, ";", " ", 1)
	tag = strings.Replace(tag, "<-chan ", "$RC_", -1)
	tag = strings.Replace(tag, "chan<- ", "$WC_", -1)
	tag = strings.Replace(tag, "chan ", "$C_", -1)

	var lv string
	var o []*parser.Parameter
	pairs := strings.Split(tag, ",")
	for i := len(pairs) - 1; i >= 0; i-- {

		var pOut *parser.Parameter
		ps := strings.Split(strings.TrimSpace(pairs[i]), " ")
		if len(ps) == 1 {
			pOut = &parser.Parameter{Name: fmt.Sprintf("v%v", i), Value: ps[0]}
			if lv != "" {
				pOut.Name = pOut.Value
				pOut.Value = lv
			}
		} else {
			pOut = &parser.Parameter{Name: ps[0], Value: ps[1]}
			lv = pOut.Value
		}

		o = append(o, pOut)
	}

	var ro []*parser.Parameter
	for i := len(o) - 1; i >= 0; i-- {
		o[i].Name = strings.Replace(o[i].Name, ";", " ", -1)
		o[i].Value = strings.Replace(o[i].Value, ";", " ", -1)
		ro = append(ro, o[i])
	}

	return ro
}

func cppTypeFromGoType(f *parser.Function, t string) (string, string) {
	tOld := t //TODO: also for differentiation of QVariant and *QVariant
	//t = strings.TrimPrefix(t, "*")

	if strings.Count(t, "[") == 1 || strings.HasSuffix(t, "][]string") {
		//TODO: multidimensional array and nested maps
		if strings.HasPrefix(t, "[]") && t != "[]string" {
			o, pureGoType := cppTypeFromGoType(f, strings.TrimPrefix(t, "[]"))
			if pureGoType == "" {
				return fmt.Sprintf("QList<%v>", o), ""
			}
		}
		if strings.HasPrefix(t, "map[") {
			var head = fmt.Sprintf("map[%v]", strings.Split(strings.TrimPrefix(t, "map["), "]")[0])
			o1, pureGoType1 := cppTypeFromGoType(f, strings.Split(strings.TrimPrefix(t, "map["), "]")[0])
			o2, pureGoType2 := cppTypeFromGoType(f, strings.TrimPrefix(t, head))
			if pureGoType1 == "" && pureGoType2 == "" {
				return fmt.Sprintf("QHash<%v, %v>", o1, o2), ""
			}
		}
	}

	if f != nil && t == "error" {
		f.AsError = true
	}

	switch t {
	case "string", "error":
		return "QString", ""
	case "[]string":
		return "QStringList", ""
	case "bool":
		return "bool", ""
	case "int8":
		return "qint8", ""
	case "uint8":
		return "quint8", ""
	case "int16":
		return "qint16", ""
	case "uint16":
		return "quint16", ""
	case "int", "int32":
		return "qint32", ""
	case "uint", "uint32":
		return "quint32", ""
	case "int64":
		return "qint64", ""
	case "uint64":
		return "quint64", ""
	case "float32":
		return "float", ""
	case "float64":
		return "qreal", ""
	case "uintptr":
		return "quintptr", ""
	case "unsafe.Pointer":
		return "void*", ""
	}

	if strings.Contains(t, ".") {
		t = strings.Split(t, ".")[1]
	}

	if strings.Contains(t, "__") {
		return strings.Replace(t, "_", ":", -1), ""
	}

	//TODO: differentiate between QVariant and *QVariant
	ttmp := strings.TrimPrefix(t, "*")
	if _, ok := parser.State.ClassMap[ttmp]; ok {
		if parser.State.ClassMap[ttmp].IsSubClassOfQObject() /*TODO: || f == nil && strings.HasPrefix(tOLD, "*")*/ {
			return ttmp + "*", ""
		}
		return ttmp, ""
	}

	//TODO: *_ITF support

	if tOld == "" || tOld == "void" {
		return "void", ""
	}

	tOld = strings.Replace(tOld, "$RC_", "<-chan ", -1)
	tOld = strings.Replace(tOld, "$WC_", "chan<- ", -1)
	tOld = strings.Replace(tOld, "$C_", "chan ", -1)

	return "quintptr", tOld
}

func hasStructors(m *parser.Module) bool {
	for _, c := range m.Namespace.Classes {
		if !c.IsSubClassOfQObject() {
			continue
		}
		if !c.HasConstructor() /*|| !c.HasDestructor()*/ {
			return false
		}
	}
	return true
}

func parseNonMocDeps(path string) {
	utils.Log.WithField("path", path).Debug("parseNonMocDeps")

	file, err := goparser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		return
	}

	for _, i := range file.Imports {
		if i.Path.Value == "\"C\"" {
			continue
		}

		if i.Name != nil {
			parser.State.MocImports[fmt.Sprintf("%v %v", i.Name.String(), i.Path.Value)] = struct{}{}
		} else {
			parser.State.MocImports[i.Path.Value] = struct{}{}
		}
	}
}
