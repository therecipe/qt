package widgets

//#include "qtextbrowser.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QTextBrowser struct {
	QTextEdit
}

type QTextBrowser_ITF interface {
	QTextEdit_ITF
	QTextBrowser_PTR() *QTextBrowser
}

func PointerFromQTextBrowser(ptr QTextBrowser_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBrowser_PTR().Pointer()
	}
	return nil
}

func NewQTextBrowserFromPointer(ptr unsafe.Pointer) *QTextBrowser {
	var n = new(QTextBrowser)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QTextBrowser_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextBrowser) QTextBrowser_PTR() *QTextBrowser {
	return ptr
}

func (ptr *QTextBrowser) OpenExternalLinks() bool {
	if ptr.Pointer() != nil {
		return C.QTextBrowser_OpenExternalLinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBrowser) OpenLinks() bool {
	if ptr.Pointer() != nil {
		return C.QTextBrowser_OpenLinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBrowser) SearchPaths() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QTextBrowser_SearchPaths(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QTextBrowser) SetOpenExternalLinks(open bool) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_SetOpenExternalLinks(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QTextBrowser) SetOpenLinks(open bool) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_SetOpenLinks(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QTextBrowser) SetSearchPaths(paths []string) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_SetSearchPaths(ptr.Pointer(), C.CString(strings.Join(paths, "|")))
	}
}

func (ptr *QTextBrowser) SetSource(name core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_SetSource(ptr.Pointer(), core.PointerFromQUrl(name))
	}
}

func NewQTextBrowser(parent QWidget_ITF) *QTextBrowser {
	return NewQTextBrowserFromPointer(C.QTextBrowser_NewQTextBrowser(PointerFromQWidget(parent)))
}

func (ptr *QTextBrowser) Backward() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_Backward(ptr.Pointer())
	}
}

func (ptr *QTextBrowser) ConnectBackwardAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectBackwardAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "backwardAvailable", f)
	}
}

func (ptr *QTextBrowser) DisconnectBackwardAvailable() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectBackwardAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "backwardAvailable")
	}
}

//export callbackQTextBrowserBackwardAvailable
func callbackQTextBrowserBackwardAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "backwardAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextBrowser) BackwardHistoryCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBrowser_BackwardHistoryCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBrowser) ClearHistory() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ClearHistory(ptr.Pointer())
	}
}

func (ptr *QTextBrowser) Forward() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_Forward(ptr.Pointer())
	}
}

func (ptr *QTextBrowser) ConnectForwardAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectForwardAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "forwardAvailable", f)
	}
}

func (ptr *QTextBrowser) DisconnectForwardAvailable() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectForwardAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "forwardAvailable")
	}
}

//export callbackQTextBrowserForwardAvailable
func callbackQTextBrowserForwardAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "forwardAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextBrowser) ForwardHistoryCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBrowser_ForwardHistoryCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBrowser) ConnectHistoryChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectHistoryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "historyChanged", f)
	}
}

func (ptr *QTextBrowser) DisconnectHistoryChanged() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectHistoryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "historyChanged")
	}
}

//export callbackQTextBrowserHistoryChanged
func callbackQTextBrowserHistoryChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "historyChanged").(func())()
}

func (ptr *QTextBrowser) HistoryTitle(i int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBrowser_HistoryTitle(ptr.Pointer(), C.int(i)))
	}
	return ""
}

func (ptr *QTextBrowser) Home() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_Home(ptr.Pointer())
	}
}

func (ptr *QTextBrowser) IsBackwardAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QTextBrowser_IsBackwardAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBrowser) IsForwardAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QTextBrowser_IsForwardAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBrowser) LoadResource(ty int, name core.QUrl_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextBrowser_LoadResource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name)))
	}
	return nil
}

func (ptr *QTextBrowser) Reload() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_Reload(ptr.Pointer())
	}
}
