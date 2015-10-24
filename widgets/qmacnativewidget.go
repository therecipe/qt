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

type QMacNativeWidgetITF interface {
	QWidgetITF
	QMacNativeWidgetPTR() *QMacNativeWidget
}

func PointerFromQMacNativeWidget(ptr QMacNativeWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacNativeWidgetPTR().Pointer()
	}
	return nil
}

func QMacNativeWidgetFromPointer(ptr unsafe.Pointer) *QMacNativeWidget {
	var n = new(QMacNativeWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMacNativeWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMacNativeWidget) QMacNativeWidgetPTR() *QMacNativeWidget {
	return ptr
}

func (ptr *QMacNativeWidget) DestroyQMacNativeWidget() {
	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DestroyQMacNativeWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
