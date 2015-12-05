package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QHBoxLayout struct {
	QBoxLayout
}

type QHBoxLayout_ITF interface {
	QBoxLayout_ITF
	QHBoxLayout_PTR() *QHBoxLayout
}

func PointerFromQHBoxLayout(ptr QHBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQHBoxLayoutFromPointer(ptr unsafe.Pointer) *QHBoxLayout {
	var n = new(QHBoxLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHBoxLayout_") {
		n.SetObjectName("QHBoxLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHBoxLayout) QHBoxLayout_PTR() *QHBoxLayout {
	return ptr
}

func NewQHBoxLayout() *QHBoxLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHBoxLayout::QHBoxLayout")
		}
	}()

	return NewQHBoxLayoutFromPointer(C.QHBoxLayout_NewQHBoxLayout())
}

func NewQHBoxLayout2(parent QWidget_ITF) *QHBoxLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHBoxLayout::QHBoxLayout")
		}
	}()

	return NewQHBoxLayoutFromPointer(C.QHBoxLayout_NewQHBoxLayout2(PointerFromQWidget(parent)))
}

func (ptr *QHBoxLayout) DestroyQHBoxLayout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHBoxLayout::~QHBoxLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHBoxLayout_DestroyQHBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
