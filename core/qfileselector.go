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

type QFileSelector_ITF interface {
	QObject_ITF
	QFileSelector_PTR() *QFileSelector
}

func PointerFromQFileSelector(ptr QFileSelector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileSelector_PTR().Pointer()
	}
	return nil
}

func NewQFileSelectorFromPointer(ptr unsafe.Pointer) *QFileSelector {
	var n = new(QFileSelector)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFileSelector_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFileSelector) QFileSelector_PTR() *QFileSelector {
	return ptr
}

func NewQFileSelector(parent QObject_ITF) *QFileSelector {
	return NewQFileSelectorFromPointer(C.QFileSelector_NewQFileSelector(PointerFromQObject(parent)))
}

func (ptr *QFileSelector) AllSelectors() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSelector_AllSelectors(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSelector) ExtraSelectors() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSelector_ExtraSelectors(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSelector) Select(filePath string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSelector_Select(ptr.Pointer(), C.CString(filePath)))
	}
	return ""
}

func (ptr *QFileSelector) SetExtraSelectors(list []string) {
	if ptr.Pointer() != nil {
		C.QFileSelector_SetExtraSelectors(ptr.Pointer(), C.CString(strings.Join(list, "|")))
	}
}

func (ptr *QFileSelector) DestroyQFileSelector() {
	if ptr.Pointer() != nil {
		C.QFileSelector_DestroyQFileSelector(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
