package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	for len(n.ObjectName()) < len("QMacNativeWidget_") {
		n.SetObjectName("QMacNativeWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMacNativeWidget) QMacNativeWidget_PTR() *QMacNativeWidget {
	return ptr
}

func (ptr *QMacNativeWidget) DestroyQMacNativeWidget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacNativeWidget::~QMacNativeWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DestroyQMacNativeWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
