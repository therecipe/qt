package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
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
	return n
}

func (ptr *QItemEditorFactory) QItemEditorFactory_PTR() *QItemEditorFactory {
	return ptr
}

func NewQItemEditorFactory() *QItemEditorFactory {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemEditorFactory::QItemEditorFactory")
		}
	}()

	return NewQItemEditorFactoryFromPointer(C.QItemEditorFactory_NewQItemEditorFactory())
}

func (ptr *QItemEditorFactory) CreateEditor(userType int, parent QWidget_ITF) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemEditorFactory::createEditor")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QItemEditorFactory_CreateEditor(ptr.Pointer(), C.int(userType), PointerFromQWidget(parent)))
	}
	return nil
}

func QItemEditorFactory_DefaultFactory() *QItemEditorFactory {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemEditorFactory::defaultFactory")
		}
	}()

	return NewQItemEditorFactoryFromPointer(C.QItemEditorFactory_QItemEditorFactory_DefaultFactory())
}

func (ptr *QItemEditorFactory) RegisterEditor(userType int, creator QItemEditorCreatorBase_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemEditorFactory::registerEditor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemEditorFactory_RegisterEditor(ptr.Pointer(), C.int(userType), PointerFromQItemEditorCreatorBase(creator))
	}
}

func QItemEditorFactory_SetDefaultFactory(factory QItemEditorFactory_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemEditorFactory::setDefaultFactory")
		}
	}()

	C.QItemEditorFactory_QItemEditorFactory_SetDefaultFactory(PointerFromQItemEditorFactory(factory))
}

func (ptr *QItemEditorFactory) ValuePropertyName(userType int) *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemEditorFactory::valuePropertyName")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QItemEditorFactory_ValuePropertyName(ptr.Pointer(), C.int(userType)))
	}
	return nil
}

func (ptr *QItemEditorFactory) DestroyQItemEditorFactory() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemEditorFactory::~QItemEditorFactory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemEditorFactory_DestroyQItemEditorFactory(ptr.Pointer())
	}
}
