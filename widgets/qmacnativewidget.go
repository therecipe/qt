package widgets

//#include "qmacnativewidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMacNativeWidget struct {
	QWidget
}

type QMacNativeWidget_ITF interface {
	QWidget_ITF
	QMacNativeWidget_PTR() *QMacNativeWidget
}

func PointerFromQMacNativeWidget(ptr QMacNativeWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacNativeWidget_PTR().Pointer()
	}
	return nil
}

func NewQMacNativeWidgetFromPointer(ptr unsafe.Pointer) *QMacNativeWidget {
	var n = new(QMacNativeWidget)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QMacNativeWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMacNativeWidget) QMacNativeWidget_PTR() *QMacNativeWidget {
	return ptr
}

func (ptr *QMacNativeWidget) DestroyQMacNativeWidget() {
	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DestroyQMacNativeWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
