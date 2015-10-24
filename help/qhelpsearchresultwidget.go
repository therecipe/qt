package help

//#include "qhelpsearchresultwidget.h"
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

type QHelpSearchResultWidgetITF interface {
	widgets.QWidgetITF
	QHelpSearchResultWidgetPTR() *QHelpSearchResultWidget
}

func PointerFromQHelpSearchResultWidget(ptr QHelpSearchResultWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchResultWidgetPTR().Pointer()
	}
	return nil
}

func QHelpSearchResultWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchResultWidget {
	var n = new(QHelpSearchResultWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHelpSearchResultWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpSearchResultWidget) QHelpSearchResultWidgetPTR() *QHelpSearchResultWidget {
	return ptr
}

func (ptr *QHelpSearchResultWidget) LinkAt(point core.QPointITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpSearchResultWidget_LinkAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point))))
	}
	return ""
}

func (ptr *QHelpSearchResultWidget) ConnectRequestShowLink(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectRequestShowLink(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "requestShowLink", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRequestShowLink() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectRequestShowLink(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "requestShowLink")
	}
}

//export callbackQHelpSearchResultWidgetRequestShowLink
func callbackQHelpSearchResultWidgetRequestShowLink(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "requestShowLink").(func(string))(C.GoString(link))
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidget() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
