package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QHelpSearchQueryWidget struct {
	widgets.QWidget
}

type QHelpSearchQueryWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget
}

func PointerFromQHelpSearchQueryWidget(ptr QHelpSearchQueryWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQueryWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchQueryWidget {
	var n = new(QHelpSearchQueryWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpSearchQueryWidget_") {
		n.SetObjectName("QHelpSearchQueryWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpSearchQueryWidget) QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget {
	return ptr
}

func (ptr *QHelpSearchQueryWidget) IsCompactMode() bool {
	defer qt.Recovering("QHelpSearchQueryWidget::isCompactMode")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_IsCompactMode(ptr.Pointer()) != 0
	}
	return false
}

func NewQHelpSearchQueryWidget(parent widgets.QWidget_ITF) *QHelpSearchQueryWidget {
	defer qt.Recovering("QHelpSearchQueryWidget::QHelpSearchQueryWidget")

	return NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(widgets.PointerFromQWidget(parent)))
}

func (ptr *QHelpSearchQueryWidget) CollapseExtendedSearch() {
	defer qt.Recovering("QHelpSearchQueryWidget::collapseExtendedSearch")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CollapseExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ExpandExtendedSearch() {
	defer qt.Recovering("QHelpSearchQueryWidget::expandExtendedSearch")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ExpandExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSearch(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::search")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectSearch(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "search", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSearch() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::search")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectSearch(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "search")
	}
}

//export callbackQHelpSearchQueryWidgetSearch
func callbackQHelpSearchQueryWidgetSearch(ptrName *C.char) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::search")

	if signal := qt.GetSignal(C.GoString(ptrName), "search"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidget() {
	defer qt.Recovering("QHelpSearchQueryWidget::~QHelpSearchQueryWidget")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
