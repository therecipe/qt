package widgets

//#include "qitemeditorcreatorbase.h"
import "C"
import (
	"unsafe"
)

type QItemEditorCreatorBase struct {
	ptr unsafe.Pointer
}

type QItemEditorCreatorBaseITF interface {
	QItemEditorCreatorBasePTR() *QItemEditorCreatorBase
}

func (p *QItemEditorCreatorBase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QItemEditorCreatorBase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQItemEditorCreatorBase(ptr QItemEditorCreatorBaseITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemEditorCreatorBasePTR().Pointer()
	}
	return nil
}

func QItemEditorCreatorBaseFromPointer(ptr unsafe.Pointer) *QItemEditorCreatorBase {
	var n = new(QItemEditorCreatorBase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QItemEditorCreatorBase) QItemEditorCreatorBasePTR() *QItemEditorCreatorBase {
	return ptr
}

func (ptr *QItemEditorCreatorBase) DestroyQItemEditorCreatorBase() {
	if ptr.Pointer() != nil {
		C.QItemEditorCreatorBase_DestroyQItemEditorCreatorBase(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QItemEditorCreatorBase) CreateWidget(parent QWidgetITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QItemEditorCreatorBase_CreateWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(parent)))))
	}
	return nil
}
