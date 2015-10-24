package widgets

//#include "qvboxlayout.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QVBoxLayout struct {
	QBoxLayout
}

type QVBoxLayoutITF interface {
	QBoxLayoutITF
	QVBoxLayoutPTR() *QVBoxLayout
}

func PointerFromQVBoxLayout(ptr QVBoxLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBoxLayoutPTR().Pointer()
	}
	return nil
}

func QVBoxLayoutFromPointer(ptr unsafe.Pointer) *QVBoxLayout {
	var n = new(QVBoxLayout)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QVBoxLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVBoxLayout) QVBoxLayoutPTR() *QVBoxLayout {
	return ptr
}

func NewQVBoxLayout() *QVBoxLayout {
	return QVBoxLayoutFromPointer(unsafe.Pointer(C.QVBoxLayout_NewQVBoxLayout()))
}

func NewQVBoxLayout2(parent QWidgetITF) *QVBoxLayout {
	return QVBoxLayoutFromPointer(unsafe.Pointer(C.QVBoxLayout_NewQVBoxLayout2(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QVBoxLayout) DestroyQVBoxLayout() {
	if ptr.Pointer() != nil {
		C.QVBoxLayout_DestroyQVBoxLayout(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
