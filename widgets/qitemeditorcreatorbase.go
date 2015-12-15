package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QItemEditorCreatorBase struct {
	ptr unsafe.Pointer
}

type QItemEditorCreatorBase_ITF interface {
	QItemEditorCreatorBase_PTR() *QItemEditorCreatorBase
}

func (p *QItemEditorCreatorBase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QItemEditorCreatorBase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQItemEditorCreatorBase(ptr QItemEditorCreatorBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemEditorCreatorBase_PTR().Pointer()
	}
	return nil
}

func NewQItemEditorCreatorBaseFromPointer(ptr unsafe.Pointer) *QItemEditorCreatorBase {
	var n = new(QItemEditorCreatorBase)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QItemEditorCreatorBase_") {
		n.SetObjectNameAbs("QItemEditorCreatorBase_" + qt.Identifier())
	}
	return n
}

func (ptr *QItemEditorCreatorBase) QItemEditorCreatorBase_PTR() *QItemEditorCreatorBase {
	return ptr
}

func (ptr *QItemEditorCreatorBase) DestroyQItemEditorCreatorBase() {
	defer qt.Recovering("QItemEditorCreatorBase::~QItemEditorCreatorBase")

	if ptr.Pointer() != nil {
		C.QItemEditorCreatorBase_DestroyQItemEditorCreatorBase(ptr.Pointer())
	}
}

func (ptr *QItemEditorCreatorBase) CreateWidget(parent QWidget_ITF) *QWidget {
	defer qt.Recovering("QItemEditorCreatorBase::createWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QItemEditorCreatorBase_CreateWidget(ptr.Pointer(), PointerFromQWidget(parent)))
	}
	return nil
}

func (ptr *QItemEditorCreatorBase) ValuePropertyName() *core.QByteArray {
	defer qt.Recovering("QItemEditorCreatorBase::valuePropertyName")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QItemEditorCreatorBase_ValuePropertyName(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemEditorCreatorBase) ObjectNameAbs() string {
	defer qt.Recovering("QItemEditorCreatorBase::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QItemEditorCreatorBase_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemEditorCreatorBase) SetObjectNameAbs(name string) {
	defer qt.Recovering("QItemEditorCreatorBase::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QItemEditorCreatorBase_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
