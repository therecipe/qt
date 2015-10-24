package qml

//#include "qqmlfileselector.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QQmlFileSelector struct {
	core.QObject
}

type QQmlFileSelectorITF interface {
	core.QObjectITF
	QQmlFileSelectorPTR() *QQmlFileSelector
}

func PointerFromQQmlFileSelector(ptr QQmlFileSelectorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlFileSelectorPTR().Pointer()
	}
	return nil
}

func QQmlFileSelectorFromPointer(ptr unsafe.Pointer) *QQmlFileSelector {
	var n = new(QQmlFileSelector)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlFileSelector_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlFileSelector) QQmlFileSelectorPTR() *QQmlFileSelector {
	return ptr
}

func NewQQmlFileSelector(engine QQmlEngineITF, parent core.QObjectITF) *QQmlFileSelector {
	return QQmlFileSelectorFromPointer(unsafe.Pointer(C.QQmlFileSelector_NewQQmlFileSelector(C.QtObjectPtr(PointerFromQQmlEngine(engine)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func QQmlFileSelector_Get(engine QQmlEngineITF) *QQmlFileSelector {
	return QQmlFileSelectorFromPointer(unsafe.Pointer(C.QQmlFileSelector_QQmlFileSelector_Get(C.QtObjectPtr(PointerFromQQmlEngine(engine)))))
}

func (ptr *QQmlFileSelector) SetExtraSelectors(strin []string) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetExtraSelectors(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(strin, "|")))
	}
}

func (ptr *QQmlFileSelector) SetExtraSelectors2(strin []string) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetExtraSelectors2(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(strin, "|")))
	}
}

func (ptr *QQmlFileSelector) SetSelector(selector core.QFileSelectorITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetSelector(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQFileSelector(selector)))
	}
}

func (ptr *QQmlFileSelector) DestroyQQmlFileSelector() {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DestroyQQmlFileSelector(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
