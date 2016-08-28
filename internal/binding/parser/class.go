package parser

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

type Class struct {
	Name      string      `xml:"name,attr"`
	Status    string      `xml:"status,attr"`
	Access    string      `xml:"access,attr"`
	Abstract  bool        `xml:"abstract,attr"`
	Bases     string      `xml:"bases,attr"`
	Module    string      `xml:"module,attr"`
	Brief     string      `xml:"brief,attr"`
	Functions []*Function `xml:"function"`
	Enums     []*Enum     `xml:"enum"`
	Variables []*Variable `xml:"variable"`
	DocModule string
	Stub      bool
	WeakLink  map[string]bool
	Export    bool
}

func (c *Class) register(module string) {
	c.DocModule = c.Module
	c.Module = module
	ClassMap[c.Name] = c
}

func (c *Class) removeFunctions() {
	for i := len(c.Functions) - 1; i >= 0; i-- {
		var f = c.Functions[i]
		if (f.Status == "obsolete" || f.Status == "compat") || !(f.Access == "public" || f.Access == "protected") || strings.ContainsAny(f.Name, "&<>=/!()[]{}^|*+-") || strings.Contains(f.Name, "Operator") {
			c.Functions = append(c.Functions[:i], c.Functions[i+1:]...)
		}
	}
}

func (c *Class) removeEnums() {
	for i := len(c.Enums) - 1; i >= 0; i-- {
		var e = c.Enums[i]
		if (e.Status == "obsolete" || e.Status == "compat") || !(e.Access == "public" || e.Access == "protected") {
			c.Enums = append(c.Enums[:i], c.Enums[i+1:]...)
		}
	}
}

func (c *Class) Dump() {
	var bb = new(bytes.Buffer)

	fmt.Fprintln(bb, "########################################\t\t\tFUNCTIONS\t\t\t########################################")
	for _, f := range c.Functions {
		fmt.Fprintln(bb, f)
	}

	fmt.Fprintln(bb, "########################################\t\t\tENUMS\t\t\t########################################")
	for _, e := range c.Enums {
		fmt.Fprintln(bb, e)
	}

	utils.MakeFolder(utils.GetQtPkgPath("internal", "binding", "dump", c.Module))
	utils.SaveBytes(utils.GetQtPkgPath("internal", "binding", "dump", c.Module, fmt.Sprintf("%v.txt", c.Name)), bb.Bytes())
}

func (c *Class) GetBases() []string {

	if c.Bases == "" {
		return make([]string, 0)
	}

	if strings.Contains(c.Bases, ",") {
		return strings.Split(c.Bases, ",")
	}

	return []string{c.Bases}
}

func (c *Class) GetAllBases() []string { return c.getAllBases(make([]string, 0)) }

func (c *Class) getAllBases(input []string) []string {

	var bases = c.GetBases()

	switch len(bases) {
	case 0:
		{
			return input
		}

	case 1:
		{
			if sb, exists := ClassMap[bases[0]]; exists {
				input = append(input, bases[0])
				return sb.getAllBases(input)
			}
			return input
		}

	}

	for _, b := range bases {
		input = append(input, b)
		if bs, exists := ClassMap[b]; exists {
			for _, sb := range bs.GetAllBases() {
				input = append(input, sb)
			}
		}
	}
	return input
}

func (c *Class) IsQObjectSubClass() bool {

	if c != nil {

		if c.Name == "QObject" {
			return true
		}

		for _, b := range c.GetAllBases() {
			if b == "QObject" {
				return true
			}
		}

	}

	return false
}

func (c *Class) fix() {
	if c.Name == "QStyle" {

		var defFunction Function

		for _, f := range c.Functions {
			if f.Name == "standardIcon" {
				defFunction = *f
				break
			}
		}

		defFunction.Name = "standardPixmap"
		defFunction.Output = "QPixmap"
		defFunction.Fullname = fmt.Sprintf("%v::%v", c.Name, defFunction.Name)

		c.Functions = append(c.Functions, &defFunction)
	}

	if c.Name == "QScxmlCppDataModel" {
		for _, s := range []struct{ Name, Output string }{
			{"evaluateToString", "QString"},
			{"evaluateToBool", "bool"},
			{"evaluateToVariant", "QVariant"},
			{"evaluateToVoid", "void"},
			{"evaluateAssignment", "void"},
			{"evaluateInitialization", "void"},
		} {
			c.Functions = append(c.Functions, &Function{
				Name:     s.Name,
				Fullname: fmt.Sprintf("%v::%v", c.Name, s.Name),
				Access:   "public",
				Virtual:  PURE,
				Meta:     PLAIN,
				Output:   s.Output,
				Parameters: []*Parameter{
					&Parameter{
						Name:  "id",
						Value: "QScxmlExecutableContent::EvaluatorId",
					},
					&Parameter{
						Name:  "ok",
						Value: "bool*",
					}},
				Signature: "()",
				Export:    true,
			})
		}

		c.Functions = append(c.Functions, &Function{
			Name:     "evaluateForeach",
			Fullname: fmt.Sprintf("%v::evaluateForeach", c.Name),
			Access:   "public",
			Virtual:  PURE,
			Meta:     PLAIN,
			Output:   "void",
			Parameters: []*Parameter{
				&Parameter{
					Name:  "id",
					Value: "QScxmlExecutableContent::EvaluatorId",
				},
				&Parameter{
					Name:  "ok",
					Value: "bool*",
				},
				&Parameter{
					Name:  "body",
					Value: "ForeachLoopBody*",
				}},
			Signature: "()",
			Export:    true,
		})
	}
}

func (c *Class) fixBases() {
	if c.Module != MOC {

		var (
			prefixPath string
			midfixPath = "include"
			suffixPath = string(filepath.Separator)
		)

		switch runtime.GOOS {
		case "windows":
			{
				prefixPath = "C:\\Qt\\Qt5.7.0\\5.7\\mingw53_32"
			}

		case "darwin":
			{
				prefixPath = "/usr/local/Qt5.7.0/5.7/clang_64"
				midfixPath = "lib"
				suffixPath = ".framework/Versions/5/Headers/"
			}

		case "linux":
			{
				switch runtime.GOARCH {
				case "amd64":
					{
						prefixPath = "/usr/local/Qt5.7.0/5.7/gcc_64"
					}

				case "386":
					{
						prefixPath = "/usr/local/Qt5.7.0/5.7/gcc"
					}
				}
			}
		}

		switch c.Name {
		case "Qt", "QtGlobalStatic", "QUnicodeTools", "QHooks", "QModulesPrivate", "QtMetaTypePrivate", "QUnicodeTables", "QAndroidJniEnvironment", "QAndroidJniObject", "QAndroidActivityResultReceiver", "QSupportedWritingSystems", "QAbstractOpenGLFunctions":
			{
				c.Bases = ""
				return
			}

		case "QUiLoader", "QEGLNativeContext", "QWGLNativeContext", "QGLXNativeContext", "QEglFSFunctions", "QWindowsWindowFunctions", "QCocoaNativeContext", "QXcbWindowFunctions", "QCocoaWindowFunctions":
			{
				c.Bases = getBasesFromHeader(utils.Load(fmt.Sprintf(filepath.Join("%v", "include", "%v", "%v.h"), prefixPath, c.Module, strings.ToLower(c.Name))), c.Name, c.Module)
				return
			}

		case "QPlatformSystemTrayIcon", "QPlatformGraphicsBuffer":
			{
				c.Bases = getBasesFromHeader(utils.Load(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v5.7.0", "QtGui", "qpa", "%v.h"), prefixPath, midfixPath, c.Module, suffixPath, strings.ToLower(c.Name))), c.Name, c.Module)
				return
			}

		case "QColumnView", "QLCDNumber", "QWebEngineUrlSchemeHandler", "QWebEngineUrlRequestInterceptor", "QWebEngineCookieStore", "QWebEngineUrlRequestInfo", "QWebEngineUrlRequestJob":
			{
				for _, m := range append(LibDeps[strings.TrimPrefix(c.Module, "Qt")], strings.TrimPrefix(c.Module, "Qt")) {
					m = fmt.Sprintf("Qt%v", m)
					if utils.Exists(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v%v"), prefixPath, midfixPath, m, suffixPath, c.Name)) {

						c.Bases = getBasesFromHeader(utils.Load(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v%v.h"), prefixPath, midfixPath, m, suffixPath, strings.ToLower(c.Name))), c.Name, c.Module)
						return
					}
				}
			}

		case "QFutureWatcher", "QDBusAbstractInterface":
			{
				c.Bases = "QObject"
				return
			}

		case "QDBusPendingReply":
			{
				c.Bases = "QDBusPendingCall"
				return
			}

		case "QRasterPaintEngine":
			{
				c.Bases = "QPaintEngine"
				return
			}
		}

		var found bool
		for _, m := range append(LibDeps[strings.TrimPrefix(c.Module, "Qt")], strings.TrimPrefix(c.Module, "Qt")) {
			m = fmt.Sprintf("Qt%v", m)
			if utils.Exists(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v%v"), prefixPath, midfixPath, m, suffixPath, c.Name)) {

				var f = utils.Load(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v%v"), prefixPath, midfixPath, m, suffixPath, c.Name))
				if f != "" {
					found = true
					c.Bases = getBasesFromHeader(utils.Load(filepath.Join(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v"), prefixPath, midfixPath, m, suffixPath), strings.Split(f, `"`)[1])), c.Name, m)
				}
				break
			}
		}

		if !found && c.Name != "SailfishApp" {
			fmt.Println("HEADER FILE NOT FOUND:", c.Name, c.Module)
		}

		var bases = c.GetBases()
		for i := len(bases) - 1; i >= 0; i-- {
			if _, exists := ClassMap[bases[i]]; !exists {
				bases = append(bases[:i], bases[i+1:]...)
			}
		}
		c.Bases = strings.Join(bases, ",")

	}
}

func getBasesFromHeader(f string, n string, m string) string {

	f = strings.Replace(f, "\r", "", -1)

	if strings.HasSuffix(n, "Iterator") {
		return ""
	}

	for i, l := range strings.Split(f, "\n") {

		if strings.Contains(l, "class "+n) || strings.Contains(l, "class Q_"+strings.ToUpper(strings.TrimPrefix(m, "Qt"))+"_EXPORT "+n) || strings.Contains(l, "class Q"+strings.ToUpper(strings.TrimPrefix(m, "Qt"))+"_EXPORT "+n) || strings.Contains(l, "class QDESIGNER_SDK_EXPORT "+n) || strings.Contains(l, "class QDESIGNER_EXTENSION_EXPORT "+n) || strings.Contains(l, "class QDESIGNER_UILIB_EXPORT "+n) || strings.Contains(l, "class  "+n) || strings.Contains(l, "class Q_"+strings.ToUpper(strings.TrimPrefix(m, "Qt"))+"_EXPORT  "+n) || strings.Contains(l, "class Q"+strings.ToUpper(strings.TrimPrefix(m, "Qt"))+"_EXPORT  "+n) || strings.Contains(l, "class QDESIGNER_SDK_EXPORT  "+n) || strings.Contains(l, "class QDESIGNER_EXTENSION_EXPORT  "+n) || strings.Contains(l, "class QDESIGNER_UILIB_EXPORT  "+n) {

			if strings.Contains(l, n+" ") || strings.Contains(l, n+":") || strings.HasSuffix(l, n) {

				l = normalizedClassDeclaration(f, i)

				if !strings.Contains(l, ":") {
					return ""
				}

				if strings.Contains(l, "<") {
					l = strings.Split(l, "<")[0]
				}

				if strings.Contains(l, "/") {
					l = strings.Split(l, "/")[0]
				}

				var tmp = strings.Split(l, ":")[1]

				for _, s := range []string{"{", "}", "#ifndef", "QT_NO_QOBJECT", "#else", "#endif", "class", "Q_" + strings.ToUpper(strings.TrimPrefix(m, "Qt")) + "_EXPORT " + n, "public", "protected", "private", "  ", " "} {
					tmp = strings.Replace(tmp, s, "", -1)
				}

				return strings.TrimSpace(tmp)
			}
		}
	}

	for _, l := range strings.Split(f, "\n") {
		if strings.Contains(l, "struct "+n) || strings.Contains(l, "struct Q_"+strings.ToUpper(strings.TrimPrefix(m, "Qt"))+"_EXPORT "+n) {
			return ""
		}
	}

	for _, l := range strings.Split(f, "\n") {
		if strings.Contains(l, "namespace "+n) {
			return ""
		}
	}

	fmt.Println("HEADER FILE NOT SUCCESFULL PARSED:", m, n)
	return ""
}

func normalizedClassDeclaration(f string, is int) string {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	for i, l := range strings.Split(f, "\n") {
		if i >= is {
			fmt.Fprint(bb, l)
			if strings.Contains(l, "{") {
				break
			}
		}
	}
	return bb.String()
}

var LibDeps = map[string][]string{
	"Core":              []string{"Gui", "Widgets"},
	"AndroidExtras":     []string{"Core"},
	"Gui":               []string{"Core", "Widgets"},
	"Network":           []string{"Core"},
	"Sql":               []string{"Core", "Widgets"},
	"Xml":               []string{"Core", "XmlPatterns"},
	"DBus":              []string{"Core"},
	"Nfc":               []string{"Core"},
	"Script":            []string{"Core"},
	"Sensors":           []string{"Core"},
	"Positioning":       []string{"Core"},
	"Widgets":           []string{"Core", "Gui"},
	"MacExtras":         []string{"Core", "Gui"},
	"Qml":               []string{"Core", "Network"},
	"WebSockets":        []string{"Core", "Network"},
	"XmlPatterns":       []string{"Core", "Network"},
	"Bluetooth":         []string{"Core", "Concurrent"},
	"WebChannel":        []string{"Core", "Network", "Qml"},
	"Svg":               []string{"Core", "Gui", "Widgets"},
	"Multimedia":        []string{"Core", "Gui", "Network", "Widgets", "MultimediaWidgets"},
	"Quick":             []string{"Core", "Gui", "Network", "Widgets", "Qml", "QuickWidgets"},
	"Help":              []string{"Core", "Gui", "Network", "Sql", "CLucene", "Widgets"},
	"Location":          []string{"Core", "Gui", "Network", "Positioning", "Qml", "Quick"},
	"ScriptTools":       []string{"Core", "Gui", "Script", "Widgets"},
	"MultimediaWidgets": []string{"Core", "Gui", "Network", "Widgets", "OpenGL", "Multimedia"},
	"UiTools":           []string{"Core", "Gui", "Widgets"},
	"X11Extras":         []string{"Core"},
	"WinExtras":         []string{},
	"WebEngine":         []string{"Core", "Gui", "Network", "WebChannel", "Widgets", "WebEngineCore", "WebEngineWidgets"},
	"WebKit":            []string{"Core", "Gui", "Network", "WebChannel", "Widgets", "PrintSupport", "WebKitWidgets"},
	"TestLib":           []string{"Core", "Gui", "Widgets", "Test"},
	"SerialPort":        []string{"Core"},
	"SerialBus":         []string{"Core"},
	"PrintSupport":      []string{"Core", "Gui", "Widgets"},
	"PlatformHeaders":   []string{"Core"},
	"Designer":          []string{"Core", "Gui", "Widgets", "UiPlugin", "DesignerComponents"},
	"Scxml":             []string{"Core", "Network", "Qml"},
	"Gamepad":           []string{"Core", "Gui"},

	"Purchasing":        []string{"Core"},
	"DataVisualization": []string{"Core", "Gui"},
	"Charts":            []string{"Core", "Gui", "Widgets"},
	"Quick2DRenderer":   []string{"Core"},

	"Sailfish": []string{"Core"},

	MOC:         make([]string, 0),
	"build_ios": []string{"Core", "Gui", "Network", "Sql", "Xml", "DBus", "Nfc", "Script", "Sensors", "Positioning", "Widgets", "Qml", "WebSockets", "XmlPatterns", "Bluetooth", "WebChannel", "Svg", "Multimedia", "Quick", "Help", "Location", "ScriptTools", "MultimediaWidgets", "UiTools", "PrintSupport"},
}

func (c *Class) hasFunctionWithName(name string) bool {
	for _, f := range c.Functions {
		if strings.ToLower(f.Name) == strings.ToLower(name) {
			return true
		}
	}
	return false
}
