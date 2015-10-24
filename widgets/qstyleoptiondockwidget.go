package widgets

//#include "qstyleoptiondockwidget.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionDockWidget struct {
	QStyleOption
}

type QStyleOptionDockWidgetITF interface {
	QStyleOptionITF
	QStyleOptionDockWidgetPTR() *QStyleOptionDockWidget
}

func PointerFromQStyleOptionDockWidget(ptr QStyleOptionDockWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionDockWidgetPTR().Pointer()
	}
	return nil
}

func QStyleOptionDockWidgetFromPointer(ptr unsafe.Pointer) *QStyleOptionDockWidget {
	var n = new(QStyleOptionDockWidget)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionDockWidget) QStyleOptionDockWidgetPTR() *QStyleOptionDockWidget {
	return ptr
}

//QStyleOptionDockWidget::StyleOptionType
type QStyleOptionDockWidget__StyleOptionType int

var (
	QStyleOptionDockWidget__Type = QStyleOptionDockWidget__StyleOptionType(QStyleOption__SO_DockWidget)
)

//QStyleOptionDockWidget::StyleOptionVersion
type QStyleOptionDockWidget__StyleOptionVersion int

var (
	QStyleOptionDockWidget__Version = QStyleOptionDockWidget__StyleOptionVersion(2)
)

func NewQStyleOptionDockWidget() *QStyleOptionDockWidget {
	return QStyleOptionDockWidgetFromPointer(unsafe.Pointer(C.QStyleOptionDockWidget_NewQStyleOptionDockWidget()))
}

func NewQStyleOptionDockWidget2(other QStyleOptionDockWidgetITF) *QStyleOptionDockWidget {
	return QStyleOptionDockWidgetFromPointer(unsafe.Pointer(C.QStyleOptionDockWidget_NewQStyleOptionDockWidget2(C.QtObjectPtr(PointerFromQStyleOptionDockWidget(other)))))
}
