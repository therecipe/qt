package widgets

//#include "qgraphicsproxywidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsProxyWidget struct {
	QGraphicsWidget
}

type QGraphicsProxyWidgetITF interface {
	QGraphicsWidgetITF
	QGraphicsProxyWidgetPTR() *QGraphicsProxyWidget
}

func PointerFromQGraphicsProxyWidget(ptr QGraphicsProxyWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsProxyWidgetPTR().Pointer()
	}
	return nil
}

func QGraphicsProxyWidgetFromPointer(ptr unsafe.Pointer) *QGraphicsProxyWidget {
	var n = new(QGraphicsProxyWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsProxyWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsProxyWidget) QGraphicsProxyWidgetPTR() *QGraphicsProxyWidget {
	return ptr
}

func NewQGraphicsProxyWidget(parent QGraphicsItemITF, wFlags core.Qt__WindowType) *QGraphicsProxyWidget {
	return QGraphicsProxyWidgetFromPointer(unsafe.Pointer(C.QGraphicsProxyWidget_NewQGraphicsProxyWidget(C.QtObjectPtr(PointerFromQGraphicsItem(parent)), C.int(wFlags))))
}

func (ptr *QGraphicsProxyWidget) CreateProxyForChildWidget(child QWidgetITF) *QGraphicsProxyWidget {
	if ptr.Pointer() != nil {
		return QGraphicsProxyWidgetFromPointer(unsafe.Pointer(C.QGraphicsProxyWidget_CreateProxyForChildWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(child)))))
	}
	return nil
}

func (ptr *QGraphicsProxyWidget) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsProxyWidget) SetGeometry(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QGraphicsProxyWidget) SetWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_SetWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsProxyWidget) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsProxyWidget_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsProxyWidget) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QGraphicsProxyWidget_Widget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsProxyWidget) DestroyQGraphicsProxyWidget() {
	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DestroyQGraphicsProxyWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
