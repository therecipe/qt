package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QItemEditorFactory struct {
	ptr unsafe.Pointer
}

type QItemEditorFactory_ITF interface {
	QItemEditorFactory_PTR() *QItemEditorFactory
}

func (p *QItemEditorFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QItemEditorFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQItemEditorFactory(ptr QItemEditorFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemEditorFactory_PTR().Pointer()
	}
	return nil
}

func NewQItemEditorFactoryFromPointer(ptr unsafe.Pointer) *QItemEditorFactory {
	var n = new(QItemEditorFactory)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QItemEditorFactory_") {
		n.SetObjectNameAbs("QItemEditorFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QItemEditorFactory) QItemEditorFactory_PTR() *QItemEditorFactory {
	return ptr
}

func NewQItemEditorFactory() *QItemEditorFactory {
	defer qt.Recovering("QItemEditorFactory::QItemEditorFactory")

	return NewQItemEditorFactoryFromPointer(C.QItemEditorFactory_NewQItemEditorFactory())
}

func (ptr *QItemEditorFactory) CreateEditor(userType int, parent QWidget_ITF) *QWidget {
	defer qt.Recovering("QItemEditorFactory::createEditor")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QItemEditorFactory_CreateEditor(ptr.Pointer(), C.int(userType), PointerFromQWidget(parent)))
	}
	return nil
}

func QItemEditorFactory_DefaultFactory() *QItemEditorFactory {
	defer qt.Recovering("QItemEditorFactory::defaultFactory")

	return NewQItemEditorFactoryFromPointer(C.QItemEditorFactory_QItemEditorFactory_DefaultFactory())
}

func (ptr *QItemEditorFactory) RegisterEditor(userType int, creator QItemEditorCreatorBase_ITF) {
	defer qt.Recovering("QItemEditorFactory::registerEditor")

	if ptr.Pointer() != nil {
		C.QItemEditorFactory_RegisterEditor(ptr.Pointer(), C.int(userType), PointerFromQItemEditorCreatorBase(creator))
	}
}

func QItemEditorFactory_SetDefaultFactory(factory QItemEditorFactory_ITF) {
	defer qt.Recovering("QItemEditorFactory::setDefaultFactory")

	C.QItemEditorFactory_QItemEditorFactory_SetDefaultFactory(PointerFromQItemEditorFactory(factory))
}

func (ptr *QItemEditorFactory) ValuePropertyName(userType int) *core.QByteArray {
	defer qt.Recovering("QItemEditorFactory::valuePropertyName")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QItemEditorFactory_ValuePropertyName(ptr.Pointer(), C.int(userType)))
	}
	return nil
}

func (ptr *QItemEditorFactory) DestroyQItemEditorFactory() {
	defer qt.Recovering("QItemEditorFactory::~QItemEditorFactory")

	if ptr.Pointer() != nil {
		C.QItemEditorFactory_DestroyQItemEditorFactory(ptr.Pointer())
	}
}

func (ptr *QItemEditorFactory) ObjectNameAbs() string {
	defer qt.Recovering("QItemEditorFactory::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QItemEditorFactory_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QItemEditorFactory) SetObjectNameAbs(name string) {
	defer qt.Recovering("QItemEditorFactory::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QItemEditorFactory_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
