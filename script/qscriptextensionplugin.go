package script

//#include "qscriptextensionplugin.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QScriptExtensionPlugin struct {
	core.QObject
}

type QScriptExtensionPlugin_ITF interface {
	core.QObject_ITF
	QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin
}

func PointerFromQScriptExtensionPlugin(ptr QScriptExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQScriptExtensionPluginFromPointer(ptr unsafe.Pointer) *QScriptExtensionPlugin {
	var n = new(QScriptExtensionPlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QScriptExtensionPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScriptExtensionPlugin) QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin {
	return ptr
}

func (ptr *QScriptExtensionPlugin) Initialize(key string, engine QScriptEngine_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_Initialize(ptr.Pointer(), C.CString(key), PointerFromQScriptEngine(engine))
	}
}

func (ptr *QScriptExtensionPlugin) Keys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptExtensionPlugin_Keys(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptExtensionPlugin) SetupPackage(key string, engine QScriptEngine_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptExtensionPlugin_SetupPackage(ptr.Pointer(), C.CString(key), PointerFromQScriptEngine(engine)))
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) DestroyQScriptExtensionPlugin() {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
