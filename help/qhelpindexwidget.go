package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QHelpIndexWidget struct {
	widgets.QListView
}

type QHelpIndexWidget_ITF interface {
	widgets.QListView_ITF
	QHelpIndexWidget_PTR() *QHelpIndexWidget
}

func PointerFromQHelpIndexWidget(ptr QHelpIndexWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpIndexWidgetFromPointer(ptr unsafe.Pointer) *QHelpIndexWidget {
	var n = new(QHelpIndexWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpIndexWidget_") {
		n.SetObjectName("QHelpIndexWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpIndexWidget) QHelpIndexWidget_PTR() *QHelpIndexWidget {
	return ptr
}

func (ptr *QHelpIndexWidget) ActivateCurrentItem() {
	defer qt.Recovering("QHelpIndexWidget::activateCurrentItem")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActivateCurrentItem(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) FilterIndices(filter string, wildcard string) {
	defer qt.Recovering("QHelpIndexWidget::filterIndices")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FilterIndices(ptr.Pointer(), C.CString(filter), C.CString(wildcard))
	}
}
