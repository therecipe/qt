package widgets

//#include "qitemeditorfactory.h"
import "C"
import (
	"unsafe"
)

type QItemEditorFactory struct {
	ptr unsafe.Pointer
}

type QItemEditorFactoryITF interface {
	QItemEditorFactoryPTR() *QItemEditorFactory
}

func (p *QItemEditorFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QItemEditorFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQItemEditorFactory(ptr QItemEditorFactoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemEditorFactoryPTR().Pointer()
	}
	return nil
}

func QItemEditorFactoryFromPointer(ptr unsafe.Pointer) *QItemEditorFactory {
	var n = new(QItemEditorFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QItemEditorFactory) QItemEditorFactoryPTR() *QItemEditorFactory {
	return ptr
}

func NewQItemEditorFactory() *QItemEditorFactory {
	return QItemEditorFactoryFromPointer(unsafe.Pointer(C.QItemEditorFactory_NewQItemEditorFactory()))
}

func (ptr *QItemEditorFactory) CreateEditor(userType int, parent QWidgetITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QItemEditorFactory_CreateEditor(C.QtObjectPtr(ptr.Pointer()), C.int(userType), C.QtObjectPtr(PointerFromQWidget(parent)))))
	}
	return nil
}

func QItemEditorFactory_DefaultFactory() *QItemEditorFactory {
	return QItemEditorFactoryFromPointer(unsafe.Pointer(C.QItemEditorFactory_QItemEditorFactory_DefaultFactory()))
}

func (ptr *QItemEditorFactory) RegisterEditor(userType int, creator QItemEditorCreatorBaseITF) {
	if ptr.Pointer() != nil {
		C.QItemEditorFactory_RegisterEditor(C.QtObjectPtr(ptr.Pointer()), C.int(userType), C.QtObjectPtr(PointerFromQItemEditorCreatorBase(creator)))
	}
}

func QItemEditorFactory_SetDefaultFactory(factory QItemEditorFactoryITF) {
	C.QItemEditorFactory_QItemEditorFactory_SetDefaultFactory(C.QtObjectPtr(PointerFromQItemEditorFactory(factory)))
}

func (ptr *QItemEditorFactory) DestroyQItemEditorFactory() {
	if ptr.Pointer() != nil {
		C.QItemEditorFactory_DestroyQItemEditorFactory(C.QtObjectPtr(ptr.Pointer()))
	}
}
