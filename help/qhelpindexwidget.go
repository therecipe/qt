package help

//#include "qhelpindexwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QHelpIndexWidget struct {
	widgets.QListView
}

type QHelpIndexWidgetITF interface {
	widgets.QListViewITF
	QHelpIndexWidgetPTR() *QHelpIndexWidget
}

func PointerFromQHelpIndexWidget(ptr QHelpIndexWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexWidgetPTR().Pointer()
	}
	return nil
}

func QHelpIndexWidgetFromPointer(ptr unsafe.Pointer) *QHelpIndexWidget {
	var n = new(QHelpIndexWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHelpIndexWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpIndexWidget) QHelpIndexWidgetPTR() *QHelpIndexWidget {
	return ptr
}

func (ptr *QHelpIndexWidget) ActivateCurrentItem() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActivateCurrentItem(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QHelpIndexWidget) FilterIndices(filter string, wildcard string) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FilterIndices(C.QtObjectPtr(ptr.Pointer()), C.CString(filter), C.CString(wildcard))
	}
}

func (ptr *QHelpIndexWidget) ConnectLinkActivated(f func(link string, keyword string)) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ConnectLinkActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLinkActivated() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectLinkActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQHelpIndexWidgetLinkActivated
func callbackQHelpIndexWidgetLinkActivated(ptrName *C.char, link *C.char, keyword *C.char) {
	qt.GetSignal(C.GoString(ptrName), "linkActivated").(func(string, string))(C.GoString(link), C.GoString(keyword))
}
