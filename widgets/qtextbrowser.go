package widgets

//#include "qtextbrowser.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QTextBrowser struct {
	QTextEdit
}

type QTextBrowserITF interface {
	QTextEditITF
	QTextBrowserPTR() *QTextBrowser
}

func PointerFromQTextBrowser(ptr QTextBrowserITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBrowserPTR().Pointer()
	}
	return nil
}

func QTextBrowserFromPointer(ptr unsafe.Pointer) *QTextBrowser {
	var n = new(QTextBrowser)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextBrowser_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextBrowser) QTextBrowserPTR() *QTextBrowser {
	return ptr
}

func (ptr *QTextBrowser) OpenExternalLinks() bool {
	if ptr.Pointer() != nil {
		return C.QTextBrowser_OpenExternalLinks(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextBrowser) OpenLinks() bool {
	if ptr.Pointer() != nil {
		return C.QTextBrowser_OpenLinks(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextBrowser) SearchPaths() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QTextBrowser_SearchPaths(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QTextBrowser) SetOpenExternalLinks(open bool) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_SetOpenExternalLinks(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QTextBrowser) SetOpenLinks(open bool) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_SetOpenLinks(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QTextBrowser) SetSearchPaths(paths []string) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_SetSearchPaths(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(paths, "|")))
	}
}

func (ptr *QTextBrowser) SetSource(name string) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_SetSource(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QTextBrowser) Source() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBrowser_Source(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQTextBrowser(parent QWidgetITF) *QTextBrowser {
	return QTextBrowserFromPointer(unsafe.Pointer(C.QTextBrowser_NewQTextBrowser(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QTextBrowser) ConnectAnchorClicked(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectAnchorClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "anchorClicked", f)
	}
}

func (ptr *QTextBrowser) DisconnectAnchorClicked() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectAnchorClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "anchorClicked")
	}
}

//export callbackQTextBrowserAnchorClicked
func callbackQTextBrowserAnchorClicked(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "anchorClicked").(func(string))(C.GoString(link))
}

func (ptr *QTextBrowser) Backward() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_Backward(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextBrowser) ConnectBackwardAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectBackwardAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "backwardAvailable", f)
	}
}

func (ptr *QTextBrowser) DisconnectBackwardAvailable() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectBackwardAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "backwardAvailable")
	}
}

//export callbackQTextBrowserBackwardAvailable
func callbackQTextBrowserBackwardAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "backwardAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextBrowser) BackwardHistoryCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBrowser_BackwardHistoryCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBrowser) ClearHistory() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ClearHistory(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextBrowser) Forward() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_Forward(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextBrowser) ConnectForwardAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectForwardAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "forwardAvailable", f)
	}
}

func (ptr *QTextBrowser) DisconnectForwardAvailable() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectForwardAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "forwardAvailable")
	}
}

//export callbackQTextBrowserForwardAvailable
func callbackQTextBrowserForwardAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "forwardAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextBrowser) ForwardHistoryCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBrowser_ForwardHistoryCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBrowser) ConnectHighlighted(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectHighlighted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "highlighted", f)
	}
}

func (ptr *QTextBrowser) DisconnectHighlighted() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectHighlighted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted")
	}
}

//export callbackQTextBrowserHighlighted
func callbackQTextBrowserHighlighted(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "highlighted").(func(string))(C.GoString(link))
}

func (ptr *QTextBrowser) ConnectHistoryChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectHistoryChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "historyChanged", f)
	}
}

func (ptr *QTextBrowser) DisconnectHistoryChanged() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectHistoryChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "historyChanged")
	}
}

//export callbackQTextBrowserHistoryChanged
func callbackQTextBrowserHistoryChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "historyChanged").(func())()
}

func (ptr *QTextBrowser) HistoryTitle(i int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBrowser_HistoryTitle(C.QtObjectPtr(ptr.Pointer()), C.int(i)))
	}
	return ""
}

func (ptr *QTextBrowser) HistoryUrl(i int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBrowser_HistoryUrl(C.QtObjectPtr(ptr.Pointer()), C.int(i)))
	}
	return ""
}

func (ptr *QTextBrowser) Home() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_Home(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextBrowser) IsBackwardAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QTextBrowser_IsBackwardAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextBrowser) IsForwardAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QTextBrowser_IsForwardAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextBrowser) LoadResource(ty int, name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBrowser_LoadResource(C.QtObjectPtr(ptr.Pointer()), C.int(ty), C.CString(name)))
	}
	return ""
}

func (ptr *QTextBrowser) Reload() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_Reload(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextBrowser) ConnectSourceChanged(f func(src string)) {
	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectSourceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QTextBrowser) DisconnectSourceChanged() {
	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectSourceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQTextBrowserSourceChanged
func callbackQTextBrowserSourceChanged(ptrName *C.char, src *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sourceChanged").(func(string))(C.GoString(src))
}
