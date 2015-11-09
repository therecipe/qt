package help

//#include "qhelpsearchresultwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QHelpSearchResultWidget struct {
	widgets.QWidget
}

type QHelpSearchResultWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget
}

func PointerFromQHelpSearchResultWidget(ptr QHelpSearchResultWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchResultWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchResultWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchResultWidget {
	var n = new(QHelpSearchResultWidget)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QHelpSearchResultWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpSearchResultWidget) QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget {
	return ptr
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidget() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
