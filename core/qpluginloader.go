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

type QPluginLoader_ITF interface {
	QObject_ITF
	QPluginLoader_PTR() *QPluginLoader
}

func PointerFromQPluginLoader(ptr QPluginLoader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPluginLoader_PTR().Pointer()
	}
	return nil
}

func NewQPluginLoaderFromPointer(ptr unsafe.Pointer) *QPluginLoader {
	var n = new(QPluginLoader)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QPluginLoader_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPluginLoader) QPluginLoader_PTR() *QPluginLoader {
	return ptr
}

func (ptr *QPluginLoader) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPluginLoader_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPluginLoader) LoadHints() QLibrary__LoadHint {
	if ptr.Pointer() != nil {
		return QLibrary__LoadHint(C.QPluginLoader_LoadHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPluginLoader) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QPluginLoader_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QPluginLoader) SetLoadHints(loadHints QLibrary__LoadHint) {
	if ptr.Pointer() != nil {
		C.QPluginLoader_SetLoadHints(ptr.Pointer(), C.int(loadHints))
	}
}

func NewQPluginLoader(parent QObject_ITF) *QPluginLoader {
	return NewQPluginLoaderFromPointer(C.QPluginLoader_NewQPluginLoader(PointerFromQObject(parent)))
}

func NewQPluginLoader2(fileName string, parent QObject_ITF) *QPluginLoader {
	return NewQPluginLoaderFromPointer(C.QPluginLoader_NewQPluginLoader2(C.CString(fileName), PointerFromQObject(parent)))
}

func (ptr *QPluginLoader) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPluginLoader_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPluginLoader) Instance() *QObject {
	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QPluginLoader_Instance(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPluginLoader) IsLoaded() bool {
	if ptr.Pointer() != nil {
		return C.QPluginLoader_IsLoaded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPluginLoader) Load() bool {
	if ptr.Pointer() != nil {
		return C.QPluginLoader_Load(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPluginLoader) MetaData() *QJsonObject {
	if ptr.Pointer() != nil {
		return NewQJsonObjectFromPointer(C.QPluginLoader_MetaData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPluginLoader) Unload() bool {
	if ptr.Pointer() != nil {
		return C.QPluginLoader_Unload(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPluginLoader) DestroyQPluginLoader() {
	if ptr.Pointer() != nil {
		C.QPluginLoader_DestroyQPluginLoader(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
