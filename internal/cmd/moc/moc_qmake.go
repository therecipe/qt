package moc

import (
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/utils"
)

func Moc(path, target string) {
	utils.Log.WithField("path", path).WithField("target", target).Debug("start Moc")

	//TODO: use getImports from qtminimal ?
	filepath.Walk(path, func(current string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != current && !isBlacklistedPath(path, current) {
			Moc(current, target)
		}
		return nil
	})

	fileList, err := ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Error("failed to read dir")
		return
	}

	var classes []*parser.Class
	var pkg string
	for _, file := range fileList {
		path := filepath.Join(path, file.Name())
		if file.IsDir() || filepath.Ext(path) != ".go" {
			utils.Log.Debugln("skipping", path)
			continue
		}
		cls, ipkg, err := parse(path)
		pkg = ipkg
		if err != nil {
			utils.Log.WithError(err)
		} else {
			classes = append(classes, cls...)
		}
	}

	c := len(classes)
	utils.Log.WithField("path", path).Debugln("found", c, "moc structs")
	if c == 0 {
		return
	}

	if len(parser.State.ClassMap) == 0 {
		parser.LoadModules()
	} else {
		utils.Log.Debug("modules already cached")
	}

	//find valid base classes
	for _, c := range classes {
		for _, bc := range c.GetBases() {
			var has bool
			for _, c := range classes {
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
	m.Prepare()

	var remaining int
	for _, class := range m.Namespace.Classes {
		if _, failed := class.GetAllBasesRecursiveCheckFailed(0); failed || !class.IsSubClassOfQObject() {
			delete(parser.State.ClassMap, class.Name)
		} else {
			remaining++
		}
	}
	utils.Log.WithField("path", path).Debugln("found", c, "remaining moc structs")
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
				p.Value = cppTypeFromGoType(f, p.Value)
			}
			f.Output = cppTypeFromGoType(f, f.Output)
		}
		utils.Log.Debug("done converting func types")
		for _, p := range c.Properties {
			p.Output = cppTypeFromGoType(nil, p.Output)
		}
		utils.Log.Debug("done converting property types")

		//TODO: needed because only now the values are decuded to c++ types
		c.FixGenericHelper()
	}
	utils.Log.Debug("done converting types")

	//copy constructor and destructor
	utils.Log.Debug("start copy structors")
	for _ = range m.Namespace.Classes {
		for _, c := range m.Namespace.Classes {
			if bc, ok := parser.State.ClassMap[c.Bases]; ok {
				for _, bcf := range bc.Functions {
					if bcf.Meta == parser.CONSTRUCTOR || bcf.Meta == parser.DESTRUCTOR {
						f := *bcf
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
	utils.Log.Debug("done copy structors")

	if err := utils.SaveBytes(filepath.Join(path, "moc.cpp"), templater.CppTemplate(parser.MOC, templater.MOC)); err != nil {
		return
	}
	if err := utils.SaveBytes(filepath.Join(path, "moc.h"), templater.HTemplate(parser.MOC, templater.MOC)); err != nil {
		return
	}
	if err := utils.SaveBytes(filepath.Join(path, "moc.go"), templater.GoTemplate(parser.MOC, false, templater.MOC, pkg)); err != nil {
		return
	}
	templater.CgoTemplate(parser.MOC, path, target, templater.MOC, pkg)

	//TODO: cleanup state
	for _, c := range parser.State.ClassMap {
		if c.Module == parser.MOC {
			delete(parser.State.ClassMap, c.Name)
		}
	}
	parser.LibDeps[parser.MOC] = make([]string, 0)

	utils.RunCmd(exec.Command(utils.ToolPath("moc", target), filepath.Join(path, "moc.cpp"), "-o", filepath.Join(path, "moc_moc.h")), "run moc")
}

func parse(path string) ([]*parser.Class, string, error) {
	utils.Log.WithField("path", path).Debug("parse")

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
							out := strings.TrimSpace(strings.Split(typ, ")")[1])
							if strings.Contains(out, "(") {
								f.Output = parameters(out + ")")[0].Value
							} else {
								f.Output = out
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

func isBlacklistedPath(path, current string) bool {
	for _, n := range []string{"deploy", "qml", "android", "ios", "ios-simulator", "sailfish", "sailfish-emulator", "rpi1", "rpi2", "rpi3", "asteroid", "node_modules", ".git"} {
		if strings.Contains(current, filepath.Join(path, n)) {
			return true
		}
	}
	return false
}

func CleanPath(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Fatal("failed to read dir")
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasPrefix(f.Name(), "moc") {
			utils.RemoveAll(filepath.Join(path, f.Name()))
		}
	}
}

func parameters(tag string) []*parser.Parameter {
	if !strings.Contains(tag, "(") {
		return nil
	}

	var lv string
	var o []*parser.Parameter
	pairs := strings.Split(strings.Split(strings.Split(tag, "(")[1], ")")[0], ",")
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
		ro = append(ro, o[i])
	}

	return ro
}

func cppTypeFromGoType(f *parser.Function, t string) string {
	//TODO: var tOld = t (for differentiation of QVariant and *QVariant)
	t = strings.TrimPrefix(t, "*")

	//TODO: multidimensional array and nested maps
	if strings.HasPrefix(t, "[]") && t != "[]string" {
		return fmt.Sprintf("QList<%v>", cppTypeFromGoType(f, strings.TrimPrefix(t, "[]")))
	}
	if strings.HasPrefix(t, "map[") {
		var head = fmt.Sprintf("map[%v]", strings.Split(strings.TrimPrefix(t, "map["), "]")[0])
		return fmt.Sprintf("QHash<%v, %v>",
			cppTypeFromGoType(f, strings.Split(strings.TrimPrefix(t, "map["), "]")[0]),
			cppTypeFromGoType(f, strings.TrimPrefix(t, head)),
		)
	}

	if f != nil && t == "error" {
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

	//TODO: differentiate between QVariant and *QVariant
	if _, ok := parser.State.ClassMap[t]; ok {
		if parser.State.ClassMap[t].IsSubClassOfQObject() /*TODO: || f == nil && strings.HasPrefix(tOLD, "*")*/ {
			return t + "*"
		}
		return t
	}

	//TODO: *_ITF support

	return "void"
}
