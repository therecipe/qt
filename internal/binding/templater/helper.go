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

		f.Fullname == "QNdefFilter::appendRecord" && !f.Overload, //QNdefRecord::TypeNameFormat -> uint

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
		return f.Export || f.Meta == parser.DESTRUCTOR || f.Fullname == "QObject::destroyed" || strings.HasPrefix(f.Name, parser.TILDE)
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
	"Core":          false,
	"AndroidExtras": false,
	"Gui":           false,
	"Network":       false,
	"Xml":           false,
	"DBus":          false,
	"Nfc":           false,
	"Script":        false,
	"Sensors":       false,
	"Positioning":   false,
	"Widgets":       false,
	"Sql":           false,
	"MacExtras":     false,
	"Qml":           false,
	"WebSockets":    false,
	"XmlPatterns":   false,
	"Bluetooth":     false,
	"WebChannel":    false,
	"Svg":           false,
	"Multimedia":    false,
	"Quick":         false,
	"Help":          false,
	"Location":      false,
	"ScriptTools":   false,
	"UiTools":       false,
	"X11Extras":     false,
	"WinExtras":     false,
	"WebEngine":     false,
	"TestLib":       false,
	"SerialPort":    false,
	"SerialBus":     false,
	"PrintSupport":  false,
	//"PlatformHeaders": false,
	"Designer": false,
	"Scxml":    false,
	"Gamepad":  false,

	"Purchasing": false,
	//"DataVisualization": false,
	//"Charts":            false,
	//"Quick2DRenderer":   false,

	"Sailfish": false,
}

var Libs = []string{
	"Core",
	"AndroidExtras",
	"Gui",
	"Network",
	"Xml",
	"DBus",
	"Nfc",
	"Script", //depreached (planned) in 5.6
	"Sensors",
	"Positioning",
	"Widgets",
	"Sql",
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
	"UiTools",
	"X11Extras",
	"WinExtras",
	"WebEngine",
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
			runtime.GOOS != "windows" && Libs[i] == "WinExtras",
			runtime.GOOS != "darwin" && Libs[i] == "MacExtras",
			runtime.GOOS != "linux" && Libs[i] == "X11Extras":
			{
				Libs = append(Libs[:i], Libs[i+1:]...)
			}
		}
	}

	return Libs
}

var LibDeps = map[string][]string{
	"Core":          []string{"Widgets", "Gui"}, //Widgets, Gui
	"AndroidExtras": []string{"Core"},
	"Gui":           []string{"Widgets", "Core"}, //Widgets
	"Network":       []string{"Core"},
	"Xml":           []string{"XmlPatterns", "Core"}, //XmlPatterns
	"DBus":          []string{"Core"},
	"Nfc":           []string{"Core"},
	"Script":        []string{"Core"},
	"Sensors":       []string{"Core"},
	"Positioning":   []string{"Core"},
	"Widgets":       []string{"Gui", "Core"},
	"Sql":           []string{"Widgets", "Gui", "Core"}, //Widgets, Gui
	"MacExtras":     []string{"Gui", "Core"},
	"Qml":           []string{"Network", "Core"},
	"WebSockets":    []string{"Network", "Core"},
	"XmlPatterns":   []string{"Network", "Core"},
	"Bluetooth":     []string{"Core"},
	"WebChannel":    []string{"Network", "Qml", "Core"}, //Network (needed for static linking ios)
	"Svg":           []string{"Widgets", "Gui", "Core"},
	"Multimedia":    []string{"MultimediaWidgets", "Widgets", "Network", "Gui", "Core"},   //MultimediaWidgets, Widgets
	"Quick":         []string{"QuickWidgets", "Widgets", "Network", "Qml", "Gui", "Core"}, //QuickWidgets, Widgets, Network (needed for static linking ios)
	"Help":          []string{"Sql", "CLucene", "Network", "Widgets", "Gui", "Core"},      //Sql + CLucene + Network (needed for static linking ios)
	"Location":      []string{"Positioning", "Quick", "Gui", "Core"},
	"ScriptTools":   []string{"Script", "Widgets", "Core"}, //Script, Widgets
	"UiTools":       []string{"Widgets", "Gui", "Core"},
	"X11Extras":     []string{"Gui", "Core"},
	"WinExtras":     []string{"Gui", "Core"},
	"WebEngine":     []string{"Widgets", "WebEngineWidgets", "WebChannel", "Network", "WebEngineCore", "Quick", "Gui", "Qml", "Core"}, //Widgets, WebEngineWidgets, WebChannel, Network
	"TestLib":       []string{"Widgets", "Gui", "Core"},                                                                               //Widgets, Gui
	"SerialPort":    []string{"Core"},
	"SerialBus":     []string{"Core"},
	"PrintSupport":  []string{"Widgets", "Gui", "Core"},
	//"PlatformHeaders": []string{}, //TODO:
	"Designer": []string{"UiPlugin", "Widgets", "Gui", "Xml", "Core"},
	"Scxml":    []string{"Network", "Qml", "Core"}, //Network (needed for static linking ios)
	"Gamepad":  []string{"Gui", "Core"},

	"Purchasing": []string{"Core"},
	//"DataVisualization": []string{"Gui", "Core"},
	//"Charts":            []string{"Widgets", "Gui", "Core"},
	//"Quick2DRenderer":   []string{}, //TODO:

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

func classNeedsCallbackFunctions(class *parser.Class) bool {

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
