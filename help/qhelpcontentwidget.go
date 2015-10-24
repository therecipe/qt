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

type QHelpContentWidgetITF interface {
	widgets.QTreeViewITF
	QHelpContentWidgetPTR() *QHelpContentWidget
}

func PointerFromQHelpContentWidget(ptr QHelpContentWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentWidgetPTR().Pointer()
	}
	return nil
}

func QHelpContentWidgetFromPointer(ptr unsafe.Pointer) *QHelpContentWidget {
	var n = new(QHelpContentWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHelpContentWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpContentWidget) QHelpContentWidgetPTR() *QHelpContentWidget {
	return ptr
}

func (ptr *QHelpContentWidget) IndexOf(link string) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QHelpContentWidget_IndexOf(C.QtObjectPtr(ptr.Pointer()), C.CString(link))))
	}
	return nil
}

func (ptr *QHelpContentWidget) ConnectLinkActivated(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ConnectLinkActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLinkActivated() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectLinkActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQHelpContentWidgetLinkActivated
func callbackQHelpContentWidgetLinkActivated(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "linkActivated").(func(string))(C.GoString(link))
}
