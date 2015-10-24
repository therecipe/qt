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

type QScriptExtensionPluginITF interface {
	core.QObjectITF
	QScriptExtensionPluginPTR() *QScriptExtensionPlugin
}

func PointerFromQScriptExtensionPlugin(ptr QScriptExtensionPluginITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionPluginPTR().Pointer()
	}
	return nil
}

func QScriptExtensionPluginFromPointer(ptr unsafe.Pointer) *QScriptExtensionPlugin {
	var n = new(QScriptExtensionPlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QScriptExtensionPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScriptExtensionPlugin) QScriptExtensionPluginPTR() *QScriptExtensionPlugin {
	return ptr
}

func (ptr *QScriptExtensionPlugin) Initialize(key string, engine QScriptEngineITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_Initialize(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.QtObjectPtr(PointerFromQScriptEngine(engine)))
	}
}

func (ptr *QScriptExtensionPlugin) Keys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptExtensionPlugin_Keys(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptExtensionPlugin) DestroyQScriptExtensionPlugin() {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
