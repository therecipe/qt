package help

//#include "qhelpsearchquerywidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QHelpSearchQueryWidget struct {
	widgets.QWidget
}

type QHelpSearchQueryWidgetITF interface {
	widgets.QWidgetITF
	QHelpSearchQueryWidgetPTR() *QHelpSearchQueryWidget
}

func PointerFromQHelpSearchQueryWidget(ptr QHelpSearchQueryWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQueryWidgetPTR().Pointer()
	}
	return nil
}

func QHelpSearchQueryWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchQueryWidget {
	var n = new(QHelpSearchQueryWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHelpSearchQueryWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpSearchQueryWidget) QHelpSearchQueryWidgetPTR() *QHelpSearchQueryWidget {
	return ptr
}

func (ptr *QHelpSearchQueryWidget) IsCompactMode() bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_IsCompactMode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQHelpSearchQueryWidget(parent widgets.QWidgetITF) *QHelpSearchQueryWidget {
	return QHelpSearchQueryWidgetFromPointer(unsafe.Pointer(C.QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(C.QtObjectPtr(widgets.PointerFromQWidget(parent)))))
}

func (ptr *QHelpSearchQueryWidget) CollapseExtendedSearch() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CollapseExtendedSearch(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QHelpSearchQueryWidget) ExpandExtendedSearch() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ExpandExtendedSearch(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSearch(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectSearch(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "search", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSearch() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectSearch(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "search")
	}
}

//export callbackQHelpSearchQueryWidgetSearch
func callbackQHelpSearchQueryWidgetSearch(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "search").(func())()
}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidget() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
