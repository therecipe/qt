package parser

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

type Class struct {
	Name       string      `xml:"name,attr"`
	Status     string      `xml:"status,attr"`
	Access     string      `xml:"access,attr"`
	Abstract   bool        `xml:"abstract,attr"`
	Bases      string      `xml:"bases,attr"`
	Module     string      `xml:"module,attr"`
	Brief      string      `xml:"brief,attr"`
	Functions  []*Function `xml:"function"`
	Enums      []*Enum     `xml:"enum"`
	Variables  []*Variable `xml:"variable"`
	Properties []*Variable `xml:"property"`
	DocModule  string
	Stub       bool
	WeakLink   map[string]bool
	Export     bool
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
	defer bb.Reset()

	fmt.Fprintln(bb, "########################################\t\t\tFUNCTIONS\t\t\t########################################")
	for _, f := range c.Functions {
		fmt.Fprintln(bb, f)
	}

	fmt.Fprintln(bb, "########################################\t\t\tENUMS\t\t\t########################################")
	for _, e := range c.Enums {
		fmt.Fprintln(bb, e)
	}

	utils.MakeFolder(utils.GoQtPkgPath("internal", "binding", "dump", c.Module))
	utils.SaveBytes(utils.GoQtPkgPath("internal", "binding", "dump", c.Module, fmt.Sprintf("%v.txt", c.Name)), bb.Bytes())
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

func (c *Class) add() {

	//TODO: needed until input + cgo support for generic Containers<T> to ignore virtuals with moc
	if c.Module == MOC {
		for _, sbc := range c.GetBases() {
			for _, sbcf := range ClassMap[sbc].Functions {
				if IsPackedList(sbcf.Output) {
					sbcf.Virtual = "non"
					sbcf.Meta = PLAIN
				}
			}
		}

		//add generic qRegisterMetaType functions
		if !c.HasFunctionWithName("qRegisterMetaType") {
			var tmpF = &Function{
				Name:           "qRegisterMetaType",
				Fullname:       fmt.Sprintf("%v::qRegisterMetaType", c.Name),
				Access:         "public",
				Virtual:        "non",
				Meta:           PLAIN,
				NonMember:      true,
				NonMoc:         true,
				Output:         fmt.Sprintf("int"),
				Parameters:     []*Parameter{},
				Signature:      "()",
				TemplateModeGo: fmt.Sprintf("%v*", c.Name),
			}
			c.Functions = append(c.Functions, tmpF)

			var f = *tmpF
			f.Overload = true
			f.OverloadNumber = "2"
			f.Parameters = []*Parameter{{Name: "typeName", Value: "const char *"}}
			f.Signature = "(const char *typeName)"
			c.Functions = append(c.Functions, &f)
		}
	}

	switch c.Name {
	case "QColor", "QFont", "QImage":
		{
			c.Functions = append(c.Functions, &Function{
				Name:       "toVariant",
				Fullname:   fmt.Sprintf("%v::toVariant", c.Name),
				Access:     "public",
				Virtual:    "non",
				Meta:       PLAIN,
				Output:     "QVariant",
				Parameters: []*Parameter{},
				Signature:  "()",
			})
		}

	case "QVariant":
		{
			for _, name := range []string{"toColor", "toFont", "toImage"} {
				c.Functions = append(c.Functions, &Function{
					Name:       name,
					Fullname:   fmt.Sprintf("%v::%v", c.Name, name),
					Access:     "public",
					Virtual:    "non",
					Meta:       PLAIN,
					Output:     strings.Replace(name, "to", "Q", -1),
					Parameters: []*Parameter{},
					Signature:  "()",
				})
			}
		}
	}

	for _, f := range c.Functions {
		switch f.Output {
		case "QModelIndexList":
			{
				f.Output = "QList<QModelIndex>"
			}

		case "QVariantList":
			{
				f.Output = "QList<QVariant>"
			}

		case "QObjectList":
			{
				f.Output = "QList<QObject *>"
			}

		case "QMediaResourceList":
			{
				f.Output = "QList<QMediaResource>"
			}

		case "QFileInfoList":
			{
				f.Output = "QList<QFileInfo>"
			}

		case "QWidgetList":
			{
				f.Output = "QList<QWidget *>"
			}

		/* TODO:
		case "QCameraFocusZoneList":
		{
			f.Output = "QList<QCameraFocusZone *>"
		}
		*/

		//generics

		case "QList<T>":
			{
				f.TemplateModeGo = "QObject*"
				f.Output = "QList<QObject*>"
			}

		case "T":
			{
				if f.Class() == "QObject" || f.Class() == "QMediaService" {
					f.TemplateModeGo = fmt.Sprintf("%v*", f.Class())
					f.Output = fmt.Sprintf("%v*", f.Class())
				}
			}
		}

		if IsPackedList(f.Output) {
			var b bool
			for _, p := range f.Parameters {
				if strings.ContainsAny(p.Value, "<>") {
					b = true
					break
				}
			}

			if !b && !c.HasFunctionWithName(fmt.Sprintf("%v_atList", f.Name)) {
				var newF = &Function{
					Name:       fmt.Sprintf("%v_atList", f.Name),
					Fullname:   fmt.Sprintf("%v::%v_atList", c.Name, f.Name),
					Access:     "public",
					Virtual:    "non",
					Meta:       PLAIN,
					Output:     fmt.Sprintf("const %v", strings.Split(strings.Split(f.Output, "<")[1], ">")[0]),
					Parameters: []*Parameter{{Name: "i", Value: "int"}},
					Signature:  "()",
					Container:  strings.Split(f.Output, "<")[0],
				}
				c.Functions = append(c.Functions, newF)
				f.Child = newF
			}
		}
	}
}

func IsPackedList(value string) bool {
	return (strings.HasPrefix(value, "QList<Q") || strings.HasPrefix(value, "QVector<Q") || strings.HasPrefix(value, "QStack<Q") || strings.HasPrefix(value, "QQueue<Q")) && strings.Count(value, "<") == 1 && !strings.Contains(value, ":") && ClassMap[UnpackedList(value)] != nil
}

func UnpackedList(value string) string {
	var CleanValue = func(value string) string {
		for _, b := range []string{"*", "const", "&amp", "&", ";"} {
			value = strings.Replace(value, b, "", -1)
		}
		return strings.TrimSpace(value)
	}
	return CleanValue(strings.Split(strings.Split(value, "<")[1], ">")[0])
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
					{
						Name:  "id",
						Value: "QScxmlExecutableContent::EvaluatorId",
					},
					{
						Name:  "ok",
						Value: "bool*",
					}},
				Signature: "()",
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
				{
					Name:  "id",
					Value: "QScxmlExecutableContent::EvaluatorId",
				},
				{
					Name:  "ok",
					Value: "bool*",
				},
				{
					Name:  "body",
					Value: "ForeachLoopBody*",
				}},
			Signature: "()",
		})
	}
}

func (c *Class) fixBases() {
	if c.Module != MOC {

		var (
			prefixPath string
			infixPath  = "include"
			suffixPath = string(filepath.Separator)
		)

		switch runtime.GOOS {
		case "windows":
			{
				if utils.UseMsys2() {
					prefixPath = utils.QT_MSYS2_DIR()
				} else {
					prefixPath = filepath.Join(utils.QT_DIR(), "5.7", "mingw53_32")
				}
			}

		case "darwin":
			{
				prefixPath = utils.QT_DARWIN_DIR()
				infixPath = "lib"
				suffixPath = ".framework/Versions/5/Headers/"
			}

		case "linux":
			{
				if utils.UsePkgConfig() {
					prefixPath = strings.TrimSpace(utils.RunCmd(exec.Command("pkg-config", "--variable=includedir", "Qt5Core"), "parser.class_includedir"))
				} else {
					prefixPath = filepath.Join(utils.QT_DIR(), "5.7", "gcc_64")
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
				if utils.UsePkgConfig() {
					c.Bases = getBasesFromHeader(utils.Load(filepath.Join(prefixPath, c.Module, fmt.Sprintf("%v.h", strings.ToLower(c.Name)))), c.Name, c.Module)
				} else {
					c.Bases = getBasesFromHeader(utils.Load(fmt.Sprintf(filepath.Join("%v", "include", "%v", "%v.h"), prefixPath, c.Module, strings.ToLower(c.Name))), c.Name, c.Module)
				}
				return
			}

		case "QPlatformSystemTrayIcon", "QPlatformGraphicsBuffer":
			{
				if utils.UsePkgConfig() {
					c.Bases = getBasesFromHeader(utils.Load(filepath.Join(prefixPath, c.Module, "5.7.0", c.Module, "qpa", fmt.Sprintf("%v.h", strings.ToLower(c.Name)))), c.Name, c.Module)
				} else {
					c.Bases = getBasesFromHeader(utils.Load(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v5.7.0", "QtGui", "qpa", "%v.h"), prefixPath, infixPath, c.Module, suffixPath, strings.ToLower(c.Name))), c.Name, c.Module)
				}
				return
			}

		case "QColumnView", "QLCDNumber", "QWebEngineUrlSchemeHandler", "QWebEngineUrlRequestInterceptor", "QWebEngineCookieStore", "QWebEngineUrlRequestInfo", "QWebEngineUrlRequestJob":
			{
				for _, m := range append(LibDeps[strings.TrimPrefix(c.Module, "Qt")], strings.TrimPrefix(c.Module, "Qt")) {
					m = fmt.Sprintf("Qt%v", m)
					if utils.UsePkgConfig() {
						if utils.Exists(filepath.Join(prefixPath, m, fmt.Sprintf("%v.h", strings.ToLower(c.Name)))) {
							c.Bases = getBasesFromHeader(utils.Load(filepath.Join(prefixPath, m, fmt.Sprintf("%v.h", strings.ToLower(c.Name)))), c.Name, c.Module)
							return
						}
					} else {
						if utils.Exists(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v%v"), prefixPath, infixPath, m, suffixPath, c.Name)) {
							c.Bases = getBasesFromHeader(utils.Load(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v%v.h"), prefixPath, infixPath, m, suffixPath, strings.ToLower(c.Name))), c.Name, c.Module)
							return
						}
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

		var libs = append(LibDeps[strings.TrimPrefix(c.Module, "Qt")], strings.TrimPrefix(c.Module, "Qt"))
		for i, v := range libs {
			if v == "TestLib" {
				libs[i] = "Test"
			}
		}

		var found bool
		for _, m := range libs {
			m = fmt.Sprintf("Qt%v", m)
			if utils.UsePkgConfig() {
				if utils.Exists(filepath.Join(prefixPath, m, c.Name)) {

					var f = utils.Load(filepath.Join(prefixPath, m, c.Name))
					if f != "" {
						found = true
						c.Bases = getBasesFromHeader(utils.Load(filepath.Join(prefixPath, m, strings.Split(f, "\"")[1])), c.Name, m)
					}
					break
				}
			} else {
				if utils.Exists(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v%v"), prefixPath, infixPath, m, suffixPath, c.Name)) {

					var f = utils.Load(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v%v"), prefixPath, infixPath, m, suffixPath, c.Name))
					if f != "" {
						found = true
						c.Bases = getBasesFromHeader(utils.Load(filepath.Join(fmt.Sprintf(filepath.Join("%v", "%v", "%v%v"), prefixPath, infixPath, m, suffixPath), strings.Split(f, "\"")[1])), c.Name, m)
					}
					break
				}
			}
		}

		if !found && c.Name != "SailfishApp" {
			utils.Log.Errorln("failed to find header file for:", c.Name, c.Module)
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

		//TODO:
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

	utils.Log.Errorln("failed to parse header for:", m, n)
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
	"Core":          {"Widgets", "Gui"}, //Widgets, Gui
	"AndroidExtras": {"Core"},
	"Gui":           {"Widgets", "Core"}, //Widgets
	"Network":       {"Core"},
	"Xml":           {"XmlPatterns", "Core"}, //XmlPatterns
	"DBus":          {"Core"},
	"Nfc":           {"Core"},
	"Script":        {"Core"},
	"Sensors":       {"Core"},
	"Positioning":   {"Core"},
	"Widgets":       {"Gui", "Core"},
	"Sql":           {"Widgets", "Gui", "Core"}, //Widgets, Gui
	"MacExtras":     {"Gui", "Core"},
	"Qml":           {"Network", "Core"},
	"WebSockets":    {"Network", "Core"},
	"XmlPatterns":   {"Network", "Core"},
	"Bluetooth":     {"Core"},
	"WebChannel":    {"Network", "Qml", "Core"}, //Network (needed for static linking ios)
	"Svg":           {"Widgets", "Gui", "Core"},
	"Multimedia":    {"MultimediaWidgets", "Widgets", "Network", "Gui", "Core"},   //MultimediaWidgets, Widgets
	"Quick":         {"QuickWidgets", "Widgets", "Network", "Qml", "Gui", "Core"}, //QuickWidgets, Widgets, Network (needed for static linking ios)
	"Help":          {"Sql", "CLucene", "Network", "Widgets", "Gui", "Core"},      //Sql + CLucene + Network (needed for static linking ios)
	"Location":      {"Positioning", "Quick", "Gui", "Core"},
	"ScriptTools":   {"Script", "Widgets", "Core"}, //Script, Widgets
	"UiTools":       {"Widgets", "Gui", "Core"},
	"X11Extras":     {"Gui", "Core"},
	"WinExtras":     {"Gui", "Core"},
	"WebEngine":     {"Widgets", "WebEngineWidgets", "WebChannel", "Network", "WebEngineCore", "Quick", "Gui", "Qml", "Core"}, //Widgets, WebEngineWidgets, WebChannel, Network
	"TestLib":       {"Widgets", "Gui", "Core"},                                                                               //Widgets, Gui
	"SerialPort":    {"Core"},
	"SerialBus":     {"Core"},
	"PrintSupport":  {"Widgets", "Gui", "Core"},
	//"PlatformHeaders": []string{}, //TODO:
	"Designer": {"UiPlugin", "Widgets", "Gui", "Xml", "Core"},
	"Scxml":    {"Network", "Qml", "Core"}, //Network (needed for static linking ios)
	"Gamepad":  {"Gui", "Core"},

	"Purchasing": {"Core"},
	//"DataVisualization": []string{"Gui", "Core"},
	//"Charts":            []string{"Widgets", "Gui", "Core"},
	//"Quick2DRenderer":   []string{}, //TODO:

	"Sailfish": {"Core"},

	MOC:         make([]string, 0),
	"build_ios": {"Core", "Gui", "Network", "Sql", "Xml", "DBus", "Nfc", "Script", "Sensors", "Positioning", "Widgets", "Qml", "WebSockets", "XmlPatterns", "Bluetooth", "WebChannel", "Svg", "Multimedia", "Quick", "Help", "Location", "ScriptTools", "MultimediaWidgets", "UiTools", "PrintSupport"},
}

func (c *Class) HasFunctionWithName(name string) bool {
	for _, f := range c.Functions {
		if strings.ToLower(f.Name) == strings.ToLower(name) {
			return true
		}
	}
	return false
}
