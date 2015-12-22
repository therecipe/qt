package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	for len(n.ObjectName()) < len("QHelpSearchResultWidget_") {
		n.SetObjectName("QHelpSearchResultWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpSearchResultWidget) QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget {
	return ptr
}

func (ptr *QHelpSearchResultWidget) LinkAt(point core.QPoint_ITF) *core.QUrl {
	defer qt.Recovering("QHelpSearchResultWidget::linkAt")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QHelpSearchResultWidget_LinkAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) ConnectRequestShowLink(f func(link *core.QUrl)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectRequestShowLink(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestShowLink", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRequestShowLink() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectRequestShowLink(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestShowLink")
	}
}

//export callbackQHelpSearchResultWidgetRequestShowLink
func callbackQHelpSearchResultWidgetRequestShowLink(ptrName *C.char, link unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::requestShowLink")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestShowLink"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidget() {
	defer qt.Recovering("QHelpSearchResultWidget::~QHelpSearchResultWidget")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
