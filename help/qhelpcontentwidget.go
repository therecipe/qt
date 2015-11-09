package help

//#include "qhelpcontentwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QHelpContentWidget struct {
	widgets.QTreeView
}

type QHelpContentWidget_ITF interface {
	widgets.QTreeView_ITF
	QHelpContentWidget_PTR() *QHelpContentWidget
}

func PointerFromQHelpContentWidget(ptr QHelpContentWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentWidgetFromPointer(ptr unsafe.Pointer) *QHelpContentWidget {
	var n = new(QHelpContentWidget)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QHelpContentWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpContentWidget) QHelpContentWidget_PTR() *QHelpContentWidget {
	return ptr
}

func (ptr *QHelpContentWidget) IndexOf(link core.QUrl_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexOf(ptr.Pointer(), core.PointerFromQUrl(link)))
	}
	return nil
}
