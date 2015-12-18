package templater

import (
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func isSupportedFunction(c *parser.Class, f *parser.Function) bool {
	return (f.Access == "public" || (f.Access == "protected" && strings.Contains(f.Virtual, "impure"))) && f.Status != "obsolete" && isNotAbstractConstructor(c, f) && isNotAtomic(f) && !isBlocked(f) && !(strings.Contains(f.Virtual, "impure") && f.Overload)
}

func isNotAbstractConstructor(c *parser.Class, f *parser.Function) bool {
	return !(hasPureVirtualFunctions(c.Name) && f.Meta == "constructor")
}

func hasPureVirtualFunctions(c string) bool {

	if c == "QGraphicsObject" {
		return true
	}

	if parser.AbstractMap[c] {
		return true
	}

	if cl, exists := parser.ClassMap[c]; exists {
		if b := cl.Bases; b != "" && !strings.Contains(b, ",") {
			if parser.AbstractMap[b] {
				if bl, exists := parser.ClassMap[b]; exists {
					for _, fBase := range bl.Functions {
						if fBase.Virtual == "pure" {
							if !classHasRealFunction(c, fBase.Name) {
								return true
							}
						}
					}
				}
			} else {
				return hasPureVirtualFunctions(b)
			}
		}
	}

	return false
}

func classHasRealFunction(c string, n string) bool {
	if c != "" && c != "Qt" {
		if cl, exists := parser.ClassMap[c]; exists {
			for _, f := range cl.Functions {
				if f.Name == n {
					return true
				}
			}
		}
	}
	return false
}

func isBlocked(f *parser.Function) bool {
	for _, n := range []string{"relations", "scriptValueFromQMetaObject", "fromScriptValue", "QPlaceProposedSearchResult", "evaluateTo", "detected", "isRecordType", "replace", "insert", "remove", "find", "changedStates", "requestTexture", "draw", "setTabPositions", "setExtraSelections", "disconnect", "QJsonObject", "QJsonArray", "QAccessibleStateChangeEvent", "hitTest", "setupUi", "setEditFocus", "toUnicode", "registerConverter", "registerEqualsComparator", "registerComparators", "hasRegisteredConverterFunction", "hasRegisteredComparators", "setNavigationMode", "navigationMode", "setNativeArguments", "setAlphaChannel", "setDefaultAction", "unregisterEventNotifier", "QXmlStreamWriter", "hasEditFocus", "QTextStream", "QStringRef", "QSignalBlocker", "defaultAction", "canConvert", "queryItemValue", "hasQueryItem", "hasEncodedQueryItem", "hasLocalData", "registerEventNotifier", "registerTimer", "setYMD", "nativeArguments"} {
		if f.Name == n {
			f.Access = "unsupported_isBlocked"
			return true
		}
	}

	for _, n := range []string{"QAccessibleInterface::state", "QAccessibleWidget::state"} {
		if f.Fullname == n {
			f.Access = "unsupported_isBlocked"
			return true
		}
	}

	//Android Only
	for _, blockedAndroid := range []string{"setAsDockMenu", "setSelect"} {
		if f.Name == blockedAndroid {
			f.Access = "unsupported_isBlocked_Android"
			return true
		}
	}

	if f.Name == "exec" && runtime.GOOS == "linux" && !strings.Contains(f.Class(), "Sql") {
		f.Output = "void"
	}

	if f.Name == "value" && f.Class() == "QVariant" {
		f.Access = "unsupported_goFunction"
		return true
	}

	if f.Class() == "QAudioBuffer" && strings.Contains(f.Output, "T") {
		f.Access = "unsupported_goFunction"
		return true
	}

	return false
}

func isObjectSubClass(b string) bool {
	if bl, exists := parser.ClassMap[b]; exists {
		return bl.IsQObjectSubClass()
	}
	return false
}

func isNotTemplate(c *parser.Class) bool {
	return !strings.Contains(c.Brief, "emplate")
}

func isNotAtomic(f *parser.Function) bool {
	return !strings.Contains(f.Fullname, "tomic")
}

func isSupportedClass(c *parser.Class) bool {
	return c.Access == "public" && c.Status != "obsolete" && isNotTemplate(c) && !isBlockedClass(c) && isSupportedFile(c)
}

func isBlockedClass(c *parser.Class) bool {
	for _, n := range []string{"QException", "QPlaceContentRequest", "QPlaceContentReply", "QPlaceContent", "QPlaceContactDetail", "QPlaceCategory", "QPlaceAttribute", "QPlace", "QGeoCodingManagerEngine", "QGeoCodingManager", "QHash", "QGenericMatrix", "QContiguousCache", "QStringList", "QCache", "QString", "QItemEditorCreator", "QImageIOPlugin", "QImageIOHandler", "QIconEnginePlugin", "QGraphicsTransform", "QGraphicsLayoutItem", "QGraphicsLayout", "QGraphicsItem", "QGraphicsEffect", "QGestureRecognizer", "QGenericPlugin", "QDebug", "QAnimationGroup", "QAccessiblePlugin", "QOpenGLFunctions_ES2", "Qt", "QSharedDataPointer", "QScopedValueRollback", "QScopedArrayPointer", "QQueue", "QMultiMap", "QWinEventNotifier", "QWeakPointer", "QVarLengthArray", "QThreadStorage", "QSharedPointer", "QScopedPointer", "QMutableHashIterator", "QMultiHash", "QMapIterator", "QListIterator", "QLinkedListIterator", "QHashIterator", "QGlobalStatic", "QFutureWatcher", "QFutureSynchronizer", "QFutureIterator", "QFuture", "QEnableSharedFromThis", "QExplicitlySharedDataPointer", "QFlags"} {
		if c.Name == n || strings.Contains(c.Name, "Iterator") {
			return true
		}
	}

	if strings.Contains(c.Name, "QPlace") {
		return true
	}

	//Android Only
	if strings.Contains(c.Name, "OpenGL") {
		return true
	}

	return false
}

func isSupportedEnum(e *parser.Enum) bool {
	return (e.Access == "public" || e.Access == "protected") && e.Status != "obsolete"
}

func needsCppValueGlue(v *parser.Value) bool {
	return strings.ContainsAny(v.Value, "()<>~+") || v.Value == ""
}

func isSupportedFile(c *parser.Class) bool {
	switch c.Name {
	case
		"QStandardItemEditorCreator",
		"QSupportedWritingSystems",
		"QRasterPaintEngine",
		"QPlatformSystemTrayIcon",
		"QPlatformGraphicsBuffer",
		"QAbstractOpenGLFunctions",
		"QQmlListProperty",
		"QSGSimpleMaterialShader",
		"QSqlRelationalDelegate",
		"QDBusPendingReply",
		"QDBusReply",
		"QDBus",
		"QGeoLocation",
		"QGeoCodeReply",
		"QGeoCodingManager",
		"QGeoCodingManagerEngine":
		return false
	}

	if strings.Contains(c.Name, "QPlace") {
		return false
	}

	if _, exists := parser.SubnamespaceMap[c.Name]; exists {
		return false
	}

	//Android Only
	if strings.Contains(c.Name, "OpenGL") {
		return false
	}

	return true
}

func ShouldBuild(module string) bool {
	return true //Build[module]
}

var Build = map[string]bool{
	"Core":              true,
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
	"Script",
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
	"ScriptTools",
	"MultimediaWidgets",
}

func GetLibs() []string {
	for i := len(Libs) - 1; i >= 0; i-- {
		switch {
		case
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
	"Core":              []string{"Widgets"},
	"AndroidExtras":     []string{"Core"},
	"Gui":               []string{"Core", "Widgets"},
	"Network":           []string{"Core"},
	"Sql":               []string{"Core"},
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
	"Multimedia":        []string{"Core", "Gui", "Network", "MultimediaWidgets"},
	"Quick":             []string{"Core", "Gui", "Network", "Widgets", "Qml", "QuickWidgets"},
	"Help":              []string{"Core", "Gui", "Network", "Sql", "CLucene", "Widgets"},
	"Location":          []string{"Core", "Gui", "Network", "Positioning", "Qml", "Quick"},
	"ScriptTools":       []string{"Core", "Gui", "Script", "Widgets"},
	"MultimediaWidgets": []string{"Core", "Gui", "Network", "Widgets", "OpenGL", "Multimedia"},

	/*
		CLucene
		Designer
		OpenGL
		Concurrent
		WinExtras
		UiTools
	*/
}

func isGeneric(f *parser.Function) bool {

	switch f.Fullname {
	case "QAndroidJniObject::getStaticField", "QAndroidJniObject::getField", "QAndroidJniObject::callStaticMethod", "QAndroidJniObject::callMethod":
		return true

	case "QAndroidJniObject::setStaticField":
		{
			if f.OverloadNumber == "2" || f.OverloadNumber == "4" {
				return true
			}
		}
	}

	return false
}

func jniGenericModes(f *parser.Function) []string {

	switch f.Name {
	case "callMethod", "callStaticMethod":
		return []string{"Int", "Boolean", "Void"} //TODO: add std string function

	case "getField", "getStaticField", "setStaticField":
		return []string{"Int", "Boolean"} //TODO: add std string function

	default:
		return make([]string, 0)
	}

}

func hasVirtualFunction(c *parser.Class) bool {

	for _, f := range c.Functions {
		if f.Virtual == "impure" {
			return true
		}
	}

	return false
}

func hasSignalFunction(c *parser.Class) bool {

	for _, f := range c.Functions {
		if f.Meta == "signal" {
			return true
		}
	}

	return false
}
