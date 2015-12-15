package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QVBoxLayout struct {
	QBoxLayout
}

type QVBoxLayout_ITF interface {
	QBoxLayout_ITF
	QVBoxLayout_PTR() *QVBoxLayout
}

func PointerFromQVBoxLayout(ptr QVBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQVBoxLayoutFromPointer(ptr unsafe.Pointer) *QVBoxLayout {
	var n = new(QVBoxLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVBoxLayout_") {
		n.SetObjectName("QVBoxLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QVBoxLayout) QVBoxLayout_PTR() *QVBoxLayout {
	return ptr
}

func NewQVBoxLayout() *QVBoxLayout {
	defer qt.Recovering("QVBoxLayout::QVBoxLayout")

	return NewQVBoxLayoutFromPointer(C.QVBoxLayout_NewQVBoxLayout())
}

func NewQVBoxLayout2(parent QWidget_ITF) *QVBoxLayout {
	defer qt.Recovering("QVBoxLayout::QVBoxLayout")

	return NewQVBoxLayoutFromPointer(C.QVBoxLayout_NewQVBoxLayout2(PointerFromQWidget(parent)))
}

func (ptr *QVBoxLayout) DestroyQVBoxLayout() {
	defer qt.Recovering("QVBoxLayout::~QVBoxLayout")

	if ptr.Pointer() != nil {
		C.QVBoxLayout_DestroyQVBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
