package templater

import (
	"runtime"
	"sort"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func functionIsSupported(_ *parser.Class, f *parser.Function) bool {

	switch {
	case
		(f.Class() == "QAccessibleObject" || f.Class() == "QAccessibleInterface" || f.Class() == "QAccessibleWidget" || //QAccessible::State -> quint64
			f.Class() == "QAccessibleStateChangeEvent") && (f.Name == "state" || f.Name == "changedStates" || f.Name == "m_changedStates" || f.Name == "setM_changedStates" || f.Meta == parser.CONSTRUCTOR),

		f.Fullname == "QPixmapCache::find" && f.OverloadNumber == "4", //Qt::Key -> int
		(f.Fullname == "QPixmapCache::remove" || f.Fullname == "QPixmapCache::insert") && f.OverloadNumber == "2",
		f.Fullname == "QPixmapCache::replace",

		f.Class() == "QSimpleXmlNodeModel" && f.Meta == parser.CONSTRUCTOR,

		f.Fullname == "QSGMaterialShader::attributeNames",

		f.Class() == "QVariant" && (f.Name == "value" || f.Name == "canConvert"), //needs template

		f.Fullname == "QNdefRecord::isRecordType", f.Fullname == "QScriptEngine::scriptValueFromQMetaObject", //needs template
		f.Fullname == "QScriptEngine::fromScriptValue", f.Fullname == "QJSEngine::fromScriptValue",

		f.Class() == "QMetaType" && //needs template
			(f.Name == "hasRegisteredComparators" || f.Name == "registerComparators" ||
				f.Name == "hasRegisteredConverterFunction" || f.Name == "registerConverter" ||
				f.Name == "registerEqualsComparator"),

		parser.ClassMap[f.Class()].Module == parser.MOC && f.Name == "metaObject", //needed for qtmoc

		f.Fullname == "QSignalBlocker::QSignalBlocker" && f.OverloadNumber == "3", //undefined symbol

		(f.Class() == "QCoreApplication" || f.Class() == "QGuiApplication" || f.Class() == "QApplication" ||
			f.Class() == "QAudioInput" || f.Class() == "QAudioOutput") && f.Name == "notify", //redeclared (name collision with QObject)

		f.Fullname == "QGraphicsItem::isBlockedByModalPanel", //** problem

		f.Name == "surfaceHandle", //QQuickWindow && QQuickView //unsupported_cppType(QPlatformSurface)

		f.Name == "readData", f.Name == "QNetworkReply", //TODO: char*

		f.Name == "QDesignerFormWindowInterface" || f.Name == "QDesignerFormWindowManagerInterface" || f.Name == "QDesignerWidgetBoxInterface", //unimplemented virtual

		strings.Contains(f.Access, "unsupported"), strings.ContainsAny(f.Signature, "<>"):

		{
			if !strings.Contains(f.Access, "unsupported") {
				f.Access = "unsupported_isBlockedFunction"
			}
			return false
		}
	}

	if f.Default {
		return functionIsSupportedDefault(f)
	}

	if Minimal {
		return f.Export || f.Meta == parser.DESTRUCTOR
	}

	return true
}

func functionIsSupportedDefault(f *parser.Function) bool {
	switch f.Fullname {
	case
		"QAnimationGroup::updateCurrentTime", "QAnimationGroup::duration",
		"QAbstractProxyModel::columnCount", "QAbstractTableModel::columnCount",
		"QAbstractListModel::data", "QAbstractTableModel::data",
		"QAbstractProxyModel::index", "QAbstractProxyModel::parent",
		"QAbstractListModel::rowCount", "QAbstractProxyModel::rowCount", "QAbstractTableModel::rowCount",

		"QPagedPaintDevice::paintEngine", "QAccessibleObject::childCount",
		"QAccessibleObject::indexOfChild", "QAccessibleObject::role",
		"QAccessibleObject::text", "QAccessibleObject::child",
		"QAccessibleObject::parent",

		"QAbstractGraphicsShapeItem::paint", "QGraphicsObject::paint",
		"QLayout::sizeHint", "QAbstractGraphicsShapeItem::boundingRect",
		"QGraphicsObject::boundingRect", "QGraphicsLayout::sizeHint",

		"QSimpleXmlNodeModel::typedValue", "QSimpleXmlNodeModel::documentUri",
		"QSimpleXmlNodeModel::compareOrder", "QSimpleXmlNodeModel::nextFromSimpleAxis",
		"QSimpleXmlNodeModel::kind", "QSimpleXmlNodeModel::root",

		"QAbstractPlanarVideoBuffer::unmap", "QAbstractPlanarVideoBuffer::mapMode",

		"QSGDynamicTexture::bind", "QSGDynamicTexture::hasMipmaps",
		"QSGDynamicTexture::textureSize", "QSGDynamicTexture::hasAlphaChannel",
		"QSGDynamicTexture::textureId",

		"QModbusClient::open", "QModbusClient::close", "QModbusServer::open", "QModbusServer::close":

		{
			return false
		}
	}

	if Minimal {
		return f.Export
	}

	return true
}

func classIsSupported(c *parser.Class) bool {

	switch c.Name {
	case
		"QString", "QStringList", "QByteArray", //mapped to primitive

		"QExplicitlySharedDataPointer", "QFuture", "QDBusPendingReply", "QDBusReply", "QFutureSynchronizer", //needs template
		"QGlobalStatic", "QMultiHash", "QQueue", "QMultiMap", "QScopedPointer", "QSharedDataPointer",
		"QScopedArrayPointer", "QSharedPointer", "QThreadStorage", "QScopedValueRollback", "QVarLengthArray",
		"QWeakPointer", "QWinEventNotifier",

		"QFlags", "QException", "QStandardItemEditorCreator", "QSGSimpleMaterialShader", "QGeoCodeReply", "QFutureWatcher", //other
		"QItemEditorCreator", "QGeoCodingManager", "QGeoCodingManagerEngine", "QQmlListProperty",

		"QPlatformGraphicsBuffer", "QPlatformSystemTrayIcon", "QRasterPaintEngine", "QSupportedWritingSystems", "QGeoLocation", //file not found or QPA API
		"QAbstractOpenGLFunctions",

		"QProcess", "QProcessEnvironment": //TODO: iOS

		{
			c.Access = "unsupported_isBlockedClass"
			return false
		}
	}

	switch {
	case
		strings.HasPrefix(c.Name, "QOpenGL"), strings.HasPrefix(c.Name, "QPlace"), //file not found or QPA API

		strings.HasPrefix(c.Name, "QAtomic"), //other

		strings.HasSuffix(c.Name, "terator"), strings.Contains(c.Brief, "emplate"): //needs template

		{
			c.Access = "unsupported_isBlockedClass"
			return false
		}
	}

	if Minimal {
		return c.Export
	}

	return true
}

func hasUnimplementedPureVirtualFunctions(className string) bool {
	for _, f := range parser.ClassMap[className].Functions {
		var f = *f
		cppFunction(&f)

		if f.Virtual == parser.PURE && !functionIsSupported(parser.ClassMap[className], &f) {
			return true
		}
	}
	return false
}

func ShouldBuild(module string) bool {
	return true //Build[module]
}

var Build = map[string]bool{
	"Core":              false,
	"AndroidExtras":     false,
	"Gui":               false,
	"Network":           false,
	"Sql":               false,
	"Xml":               false,
	"DBus":              false,
	"Nfc":               false,
	"Script":            false,
	"Sensors":           false,
	"Positioning":       false,
	"Widgets":           false,
	"MacExtras":         false,
	"Qml":               false,
	"WebSockets":        false,
	"XmlPatterns":       false,
	"Bluetooth":         false,
	"WebChannel":        false,
	"Svg":               false,
	"Multimedia":        false,
	"Quick":             false,
	"Help":              false,
	"Location":          false,
	"ScriptTools":       false,
	"MultimediaWidgets": false,
	"UiTools":           false,
	"X11Extras":         false,
	"WinExtras":         false,
	"WebEngine":         false,
	"WebKit":            false,
	"TestLib":           false,
	"SerialPort":        false,
	"SerialBus":         false,
	"PrintSupport":      false,
	"PlatformHeaders":   false,
	"Designer":          false,
	"Scxml":             false,
	"Gamepad":           false,
	"Purchasing":        false,
	"DataVisualization": false,
	"Charts":            false,
	"Quick2DRenderer":   false,
	"Sailfish":          false,
}

var Libs = []string{
	"Core",
	"AndroidExtras",
	"Gui",
	"Network",
	"Sql",
	"Xml",
	"DBus",
	"Nfc",
	"Script", //depreached (planned) in 5.6
	"Sensors",
	"Positioning",
	"Widgets",
	"MacExtras",
	"Qml",
	"WebSockets",
	"XmlPatterns",
	"Bluetooth",
	"WebChannel",
	"Svg",
	"Multimedia",
	"Quick",
	"Help",
	"Location",
	"ScriptTools", //depreached (planned) in 5.6
	//"MultimediaWidgets", //depreached (merged) in 5.6
	"UiTools",
	//"X11Extras", //linux/android only
	//"WinExtras", //windows only
	"WebEngine",
	//"WebKit", //depreached (full) in 5.6
	"TestLib",
	"SerialPort",
	"SerialBus",
	"PrintSupport",
	//"PlatformHeaders", //missing imports/guards
	"Designer",
	"Scxml",
	"Gamepad",

	"Purchasing", //GPLv3 & LGPLv3
	//"DataVisualization", //GPLv3
	//"Charts",            //GPLv3
	//"Quick2DRenderer",   //GPLv3

	"Sailfish",
}

func GetLibs() []string {
	for i := len(Libs) - 1; i >= 0; i-- {
		switch {
		case
			!(runtime.GOOS == "darwin" || runtime.GOOS == "linux") && Libs[i] == "WebEngine",
			runtime.GOOS != "darwin" && Libs[i] == "MacExtras",
			runtime.GOOS != "windows" && Libs[i] == "WinExtras":
			{
				Libs = append(Libs[:i], Libs[i+1:]...)
			}
		}
	}

	return Libs
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

	parser.MOC:  make([]string, 0),
	"build_ios": []string{"Core", "Gui", "Network", "Sql", "Xml", "DBus", "Nfc", "Script", "Sensors", "Positioning", "Widgets", "Qml", "WebSockets", "XmlPatterns", "Bluetooth", "WebChannel", "Svg", "Multimedia", "Quick", "Help", "Location", "ScriptTools", "MultimediaWidgets", "UiTools", "PrintSupport"},
}

func isGeneric(f *parser.Function) bool {

	if f.Class() == "QAndroidJniObject" {
		switch f.Name {
		case
			"callMethod",
			"callStaticMethod",

			"getField",
			//"setField", -> uses interface{} if not generic

			"getStaticField",
			//"setStaticField", -> uses interface{} if not generic

			"getObjectField",

			"getStaticObjectField",

			"callObjectMethod",
			"callStaticObjectMethod":
			{
				return true
			}

		case "setStaticField":
			{
				if f.OverloadNumber == "2" || f.OverloadNumber == "4" {
					return true
				}
			}
		}
	}

	return false
}

func needsCallbackFunctions(class *parser.Class) bool {
	for _, function := range class.Functions {
		if function.Virtual == parser.IMPURE || function.Virtual == parser.PURE || function.Meta == parser.SIGNAL || function.Meta == parser.SLOT {
			return true
		}
	}

	return false
}

func shortModule(module string) string {
	return strings.ToLower(strings.TrimPrefix(module, "Qt"))
}

func getSortedClassNamesForModule(module string) []string {
	var output = make([]string, 0)
	for _, class := range parser.ClassMap {
		if class.Module == module {
			output = append(output, class.Name)
		}
	}
	sort.Stable(sort.StringSlice(output))
	return output
}

func getSortedClassesForModule(module string) []*parser.Class {
	var (
		classNames = getSortedClassNamesForModule(module)
		output     = make([]*parser.Class, len(classNames))
	)
	for i, name := range classNames {
		output[i] = parser.ClassMap[name]
	}
	return output
}

//TODO: move to parser
func addCallbackNameFunctions(c *parser.Class) {
	if !c.IsQObjectSubClass() && needsCallbackFunctions(c) {
		c.Functions = append(c.Functions, &parser.Function{
			Name:     "objectNameAbs",
			Fullname: c.Name + "::" + "objectNameAbs",
			Access:   "public",
			Meta:     parser.PLAIN,
			Output:   "QString",
			Export:   true,
		})
		c.Functions = append(c.Functions, &parser.Function{
			Name:     "setObjectNameAbs",
			Fullname: c.Name + "::" + "setObjectNameAbs",
			Access:   "public",
			Meta:     parser.PLAIN,
			Output:   parser.VOID,
			Parameters: []*parser.Parameter{&parser.Parameter{
				Name:  "name",
				Value: "QString",
			}},
			Export: true,
		})
	}
}

func manualWeakLink(module string) {
	for _, class := range getSortedClassesForModule(module) {
		class.WeakLink = make(map[string]bool)

		switch class.Module {
		case "QtCore":
			{
				class.WeakLink["QtWidgets"] = true
			}

		case "QtGui":
			{
				class.WeakLink["QtWidgets"] = true
				class.WeakLink["QtMultimedia"] = true
			}
		}
	}
}

func classNeedsDestructor(c *parser.Class) bool {
	for _, f := range c.Functions {
		if f.Meta == parser.DESTRUCTOR {
			return false
		}
	}
	return true
}
