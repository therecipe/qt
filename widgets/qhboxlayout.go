package widgets

//#include "qhboxlayout.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QHBoxLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHBoxLayout) QHBoxLayout_PTR() *QHBoxLayout {
	return ptr
}

func NewQHBoxLayout() *QHBoxLayout {
	return NewQHBoxLayoutFromPointer(C.QHBoxLayout_NewQHBoxLayout())
}

func NewQHBoxLayout2(parent QWidget_ITF) *QHBoxLayout {
	return NewQHBoxLayoutFromPointer(C.QHBoxLayout_NewQHBoxLayout2(PointerFromQWidget(parent)))
}

func (ptr *QHBoxLayout) DestroyQHBoxLayout() {
	if ptr.Pointer() != nil {
		C.QHBoxLayout_DestroyQHBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
