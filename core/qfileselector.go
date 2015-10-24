package core

//#include "qfileselector.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QFileSelector struct {
	QObject
}

type QFileSelectorITF interface {
	QObjectITF
	QFileSelectorPTR() *QFileSelector
}

func PointerFromQFileSelector(ptr QFileSelectorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileSelectorPTR().Pointer()
	}
	return nil
}

func QFileSelectorFromPointer(ptr unsafe.Pointer) *QFileSelector {
	var n = new(QFileSelector)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFileSelector_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFileSelector) QFileSelectorPTR() *QFileSelector {
	return ptr
}

func NewQFileSelector(parent QObjectITF) *QFileSelector {
	return QFileSelectorFromPointer(unsafe.Pointer(C.QFileSelector_NewQFileSelector(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QFileSelector) AllSelectors() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSelector_AllSelectors(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSelector) ExtraSelectors() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSelector_ExtraSelectors(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSelector) Select(filePath string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSelector_Select(C.QtObjectPtr(ptr.Pointer()), C.CString(filePath)))
	}
	return ""
}

func (ptr *QFileSelector) Select2(filePath string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSelector_Select2(C.QtObjectPtr(ptr.Pointer()), C.CString(filePath)))
	}
	return ""
}

func (ptr *QFileSelector) SetExtraSelectors(list []string) {
	if ptr.Pointer() != nil {
		C.QFileSelector_SetExtraSelectors(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(list, "|")))
	}
}

func (ptr *QFileSelector) DestroyQFileSelector() {
	if ptr.Pointer() != nil {
		C.QFileSelector_DestroyQFileSelector(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
