package core

//#include "qpluginloader.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPluginLoader struct {
	QObject
}

type QPluginLoaderITF interface {
	QObjectITF
	QPluginLoaderPTR() *QPluginLoader
}

func PointerFromQPluginLoader(ptr QPluginLoaderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPluginLoaderPTR().Pointer()
	}
	return nil
}

func QPluginLoaderFromPointer(ptr unsafe.Pointer) *QPluginLoader {
	var n = new(QPluginLoader)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPluginLoader_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPluginLoader) QPluginLoaderPTR() *QPluginLoader {
	return ptr
}

func (ptr *QPluginLoader) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPluginLoader_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QPluginLoader) LoadHints() QLibrary__LoadHint {
	if ptr.Pointer() != nil {
		return QLibrary__LoadHint(C.QPluginLoader_LoadHints(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPluginLoader) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QPluginLoader_SetFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName))
	}
}

func (ptr *QPluginLoader) SetLoadHints(loadHints QLibrary__LoadHint) {
	if ptr.Pointer() != nil {
		C.QPluginLoader_SetLoadHints(C.QtObjectPtr(ptr.Pointer()), C.int(loadHints))
	}
}

func NewQPluginLoader(parent QObjectITF) *QPluginLoader {
	return QPluginLoaderFromPointer(unsafe.Pointer(C.QPluginLoader_NewQPluginLoader(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQPluginLoader2(fileName string, parent QObjectITF) *QPluginLoader {
	return QPluginLoaderFromPointer(unsafe.Pointer(C.QPluginLoader_NewQPluginLoader2(C.CString(fileName), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QPluginLoader) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPluginLoader_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QPluginLoader) Instance() *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QPluginLoader_Instance(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPluginLoader) IsLoaded() bool {
	if ptr.Pointer() != nil {
		return C.QPluginLoader_IsLoaded(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPluginLoader) Load() bool {
	if ptr.Pointer() != nil {
		return C.QPluginLoader_Load(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPluginLoader) Unload() bool {
	if ptr.Pointer() != nil {
		return C.QPluginLoader_Unload(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPluginLoader) DestroyQPluginLoader() {
	if ptr.Pointer() != nil {
		C.QPluginLoader_DestroyQPluginLoader(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
