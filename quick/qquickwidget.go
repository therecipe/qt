package quick

//#include "qquickwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QQuickWidget struct {
	widgets.QWidget
}

type QQuickWidgetITF interface {
	widgets.QWidgetITF
	QQuickWidgetPTR() *QQuickWidget
}

func PointerFromQQuickWidget(ptr QQuickWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWidgetPTR().Pointer()
	}
	return nil
}

func QQuickWidgetFromPointer(ptr unsafe.Pointer) *QQuickWidget {
	var n = new(QQuickWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickWidget) QQuickWidgetPTR() *QQuickWidget {
	return ptr
}

//QQuickWidget::ResizeMode
type QQuickWidget__ResizeMode int

var (
	QQuickWidget__SizeViewToRootObject = QQuickWidget__ResizeMode(0)
	QQuickWidget__SizeRootObjectToView = QQuickWidget__ResizeMode(1)
)

//QQuickWidget::Status
type QQuickWidget__Status int

var (
	QQuickWidget__Null    = QQuickWidget__Status(0)
	QQuickWidget__Ready   = QQuickWidget__Status(1)
	QQuickWidget__Loading = QQuickWidget__Status(2)
	QQuickWidget__Error   = QQuickWidget__Status(3)
)
