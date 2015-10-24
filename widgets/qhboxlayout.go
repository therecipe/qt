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

type QHBoxLayoutITF interface {
	QBoxLayoutITF
	QHBoxLayoutPTR() *QHBoxLayout
}

func PointerFromQHBoxLayout(ptr QHBoxLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBoxLayoutPTR().Pointer()
	}
	return nil
}

func QHBoxLayoutFromPointer(ptr unsafe.Pointer) *QHBoxLayout {
	var n = new(QHBoxLayout)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHBoxLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHBoxLayout) QHBoxLayoutPTR() *QHBoxLayout {
	return ptr
}

func NewQHBoxLayout() *QHBoxLayout {
	return QHBoxLayoutFromPointer(unsafe.Pointer(C.QHBoxLayout_NewQHBoxLayout()))
}

func NewQHBoxLayout2(parent QWidgetITF) *QHBoxLayout {
	return QHBoxLayoutFromPointer(unsafe.Pointer(C.QHBoxLayout_NewQHBoxLayout2(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QHBoxLayout) DestroyQHBoxLayout() {
	if ptr.Pointer() != nil {
		C.QHBoxLayout_DestroyQHBoxLayout(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
