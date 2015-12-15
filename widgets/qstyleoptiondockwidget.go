package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStyleOptionDockWidget struct {
	QStyleOption
}

type QStyleOptionDockWidget_ITF interface {
	QStyleOption_ITF
	QStyleOptionDockWidget_PTR() *QStyleOptionDockWidget
}

func PointerFromQStyleOptionDockWidget(ptr QStyleOptionDockWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionDockWidget_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionDockWidgetFromPointer(ptr unsafe.Pointer) *QStyleOptionDockWidget {
	var n = new(QStyleOptionDockWidget)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionDockWidget) QStyleOptionDockWidget_PTR() *QStyleOptionDockWidget {
	return ptr
}

//QStyleOptionDockWidget::StyleOptionType
type QStyleOptionDockWidget__StyleOptionType int64

var (
	QStyleOptionDockWidget__Type = QStyleOptionDockWidget__StyleOptionType(QStyleOption__SO_DockWidget)
)

//QStyleOptionDockWidget::StyleOptionVersion
type QStyleOptionDockWidget__StyleOptionVersion int64

var (
	QStyleOptionDockWidget__Version = QStyleOptionDockWidget__StyleOptionVersion(2)
)

func NewQStyleOptionDockWidget() *QStyleOptionDockWidget {
	defer qt.Recovering("QStyleOptionDockWidget::QStyleOptionDockWidget")

	return NewQStyleOptionDockWidgetFromPointer(C.QStyleOptionDockWidget_NewQStyleOptionDockWidget())
}

func NewQStyleOptionDockWidget2(other QStyleOptionDockWidget_ITF) *QStyleOptionDockWidget {
	defer qt.Recovering("QStyleOptionDockWidget::QStyleOptionDockWidget")

	return NewQStyleOptionDockWidgetFromPointer(C.QStyleOptionDockWidget_NewQStyleOptionDockWidget2(PointerFromQStyleOptionDockWidget(other)))
}
