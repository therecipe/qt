package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
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
	for len(n.ObjectName()) < len("QTextBrowser_") {
		n.SetObjectName("QTextBrowser_" + qt.Identifier())
	}
	return n
}

func (ptr *QTextBrowser) QTextBrowser_PTR() *QTextBrowser {
	return ptr
}

func (ptr *QTextBrowser) OpenExternalLinks() bool {
	defer qt.Recovering("QTextBrowser::openExternalLinks")

	if ptr.Pointer() != nil {
		return C.QTextBrowser_OpenExternalLinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBrowser) OpenLinks() bool {
	defer qt.Recovering("QTextBrowser::openLinks")

	if ptr.Pointer() != nil {
		return C.QTextBrowser_OpenLinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBrowser) SearchPaths() []string {
	defer qt.Recovering("QTextBrowser::searchPaths")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QTextBrowser_SearchPaths(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QTextBrowser) SetOpenExternalLinks(open bool) {
	defer qt.Recovering("QTextBrowser::setOpenExternalLinks")

	if ptr.Pointer() != nil {
		C.QTextBrowser_SetOpenExternalLinks(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QTextBrowser) SetOpenLinks(open bool) {
	defer qt.Recovering("QTextBrowser::setOpenLinks")

	if ptr.Pointer() != nil {
		C.QTextBrowser_SetOpenLinks(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QTextBrowser) SetSearchPaths(paths []string) {
	defer qt.Recovering("QTextBrowser::setSearchPaths")

	if ptr.Pointer() != nil {
		C.QTextBrowser_SetSearchPaths(ptr.Pointer(), C.CString(strings.Join(paths, ",,,")))
	}
}

func (ptr *QTextBrowser) ConnectSetSource(f func(name *core.QUrl)) {
	defer qt.Recovering("connect QTextBrowser::setSource")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSource", f)
	}
}

func (ptr *QTextBrowser) DisconnectSetSource() {
	defer qt.Recovering("disconnect QTextBrowser::setSource")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSource")
	}
}

//export callbackQTextBrowserSetSource
func callbackQTextBrowserSetSource(ptrName *C.char, name unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::setSource")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSource"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(name))
		return true
	}
	return false

}

func (ptr *QTextBrowser) Source() *core.QUrl {
	defer qt.Recovering("QTextBrowser::source")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QTextBrowser_Source(ptr.Pointer()))
	}
	return nil
}

func NewQTextBrowser(parent QWidget_ITF) *QTextBrowser {
	defer qt.Recovering("QTextBrowser::QTextBrowser")

	return NewQTextBrowserFromPointer(C.QTextBrowser_NewQTextBrowser(PointerFromQWidget(parent)))
}

func (ptr *QTextBrowser) ConnectAnchorClicked(f func(link *core.QUrl)) {
	defer qt.Recovering("connect QTextBrowser::anchorClicked")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectAnchorClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "anchorClicked", f)
	}
}

func (ptr *QTextBrowser) DisconnectAnchorClicked() {
	defer qt.Recovering("disconnect QTextBrowser::anchorClicked")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectAnchorClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "anchorClicked")
	}
}

//export callbackQTextBrowserAnchorClicked
func callbackQTextBrowserAnchorClicked(ptrName *C.char, link unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::anchorClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "anchorClicked"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QTextBrowser) ConnectBackward(f func()) {
	defer qt.Recovering("connect QTextBrowser::backward")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "backward", f)
	}
}

func (ptr *QTextBrowser) DisconnectBackward() {
	defer qt.Recovering("disconnect QTextBrowser::backward")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "backward")
	}
}

//export callbackQTextBrowserBackward
func callbackQTextBrowserBackward(ptrName *C.char) bool {
	defer qt.Recovering("callback QTextBrowser::backward")

	if signal := qt.GetSignal(C.GoString(ptrName), "backward"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectBackwardAvailable(f func(available bool)) {
	defer qt.Recovering("connect QTextBrowser::backwardAvailable")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectBackwardAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "backwardAvailable", f)
	}
}

func (ptr *QTextBrowser) DisconnectBackwardAvailable() {
	defer qt.Recovering("disconnect QTextBrowser::backwardAvailable")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectBackwardAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "backwardAvailable")
	}
}

//export callbackQTextBrowserBackwardAvailable
func callbackQTextBrowserBackwardAvailable(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QTextBrowser::backwardAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "backwardAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QTextBrowser) BackwardHistoryCount() int {
	defer qt.Recovering("QTextBrowser::backwardHistoryCount")

	if ptr.Pointer() != nil {
		return int(C.QTextBrowser_BackwardHistoryCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBrowser) ClearHistory() {
	defer qt.Recovering("QTextBrowser::clearHistory")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ClearHistory(ptr.Pointer())
	}
}

func (ptr *QTextBrowser) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTextBrowser::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QTextBrowser::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQTextBrowserFocusOutEvent
func callbackQTextBrowserFocusOutEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectForward(f func()) {
	defer qt.Recovering("connect QTextBrowser::forward")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "forward", f)
	}
}

func (ptr *QTextBrowser) DisconnectForward() {
	defer qt.Recovering("disconnect QTextBrowser::forward")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "forward")
	}
}

//export callbackQTextBrowserForward
func callbackQTextBrowserForward(ptrName *C.char) bool {
	defer qt.Recovering("callback QTextBrowser::forward")

	if signal := qt.GetSignal(C.GoString(ptrName), "forward"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectForwardAvailable(f func(available bool)) {
	defer qt.Recovering("connect QTextBrowser::forwardAvailable")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectForwardAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "forwardAvailable", f)
	}
}

func (ptr *QTextBrowser) DisconnectForwardAvailable() {
	defer qt.Recovering("disconnect QTextBrowser::forwardAvailable")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectForwardAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "forwardAvailable")
	}
}

//export callbackQTextBrowserForwardAvailable
func callbackQTextBrowserForwardAvailable(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QTextBrowser::forwardAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "forwardAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QTextBrowser) ForwardHistoryCount() int {
	defer qt.Recovering("QTextBrowser::forwardHistoryCount")

	if ptr.Pointer() != nil {
		return int(C.QTextBrowser_ForwardHistoryCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBrowser) ConnectHighlighted2(f func(link string)) {
	defer qt.Recovering("connect QTextBrowser::highlighted")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectHighlighted2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "highlighted2", f)
	}
}

func (ptr *QTextBrowser) DisconnectHighlighted2() {
	defer qt.Recovering("disconnect QTextBrowser::highlighted")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectHighlighted2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted2")
	}
}

//export callbackQTextBrowserHighlighted2
func callbackQTextBrowserHighlighted2(ptrName *C.char, link *C.char) {
	defer qt.Recovering("callback QTextBrowser::highlighted")

	if signal := qt.GetSignal(C.GoString(ptrName), "highlighted2"); signal != nil {
		signal.(func(string))(C.GoString(link))
	}

}

func (ptr *QTextBrowser) ConnectHighlighted(f func(link *core.QUrl)) {
	defer qt.Recovering("connect QTextBrowser::highlighted")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectHighlighted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "highlighted", f)
	}
}

func (ptr *QTextBrowser) DisconnectHighlighted() {
	defer qt.Recovering("disconnect QTextBrowser::highlighted")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectHighlighted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted")
	}
}

//export callbackQTextBrowserHighlighted
func callbackQTextBrowserHighlighted(ptrName *C.char, link unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::highlighted")

	if signal := qt.GetSignal(C.GoString(ptrName), "highlighted"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QTextBrowser) ConnectHistoryChanged(f func()) {
	defer qt.Recovering("connect QTextBrowser::historyChanged")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectHistoryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "historyChanged", f)
	}
}

func (ptr *QTextBrowser) DisconnectHistoryChanged() {
	defer qt.Recovering("disconnect QTextBrowser::historyChanged")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectHistoryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "historyChanged")
	}
}

//export callbackQTextBrowserHistoryChanged
func callbackQTextBrowserHistoryChanged(ptrName *C.char) {
	defer qt.Recovering("callback QTextBrowser::historyChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "historyChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextBrowser) HistoryTitle(i int) string {
	defer qt.Recovering("QTextBrowser::historyTitle")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBrowser_HistoryTitle(ptr.Pointer(), C.int(i)))
	}
	return ""
}

func (ptr *QTextBrowser) HistoryUrl(i int) *core.QUrl {
	defer qt.Recovering("QTextBrowser::historyUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QTextBrowser_HistoryUrl(ptr.Pointer(), C.int(i)))
	}
	return nil
}

func (ptr *QTextBrowser) ConnectHome(f func()) {
	defer qt.Recovering("connect QTextBrowser::home")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "home", f)
	}
}

func (ptr *QTextBrowser) DisconnectHome() {
	defer qt.Recovering("disconnect QTextBrowser::home")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "home")
	}
}

//export callbackQTextBrowserHome
func callbackQTextBrowserHome(ptrName *C.char) bool {
	defer qt.Recovering("callback QTextBrowser::home")

	if signal := qt.GetSignal(C.GoString(ptrName), "home"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTextBrowser) IsBackwardAvailable() bool {
	defer qt.Recovering("QTextBrowser::isBackwardAvailable")

	if ptr.Pointer() != nil {
		return C.QTextBrowser_IsBackwardAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBrowser) IsForwardAvailable() bool {
	defer qt.Recovering("QTextBrowser::isForwardAvailable")

	if ptr.Pointer() != nil {
		return C.QTextBrowser_IsForwardAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBrowser) ConnectKeyPressEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTextBrowser::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QTextBrowser::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQTextBrowserKeyPressEvent
func callbackQTextBrowserKeyPressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QTextBrowser) LoadResource(ty int, name core.QUrl_ITF) *core.QVariant {
	defer qt.Recovering("QTextBrowser::loadResource")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextBrowser_LoadResource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name)))
	}
	return nil
}

func (ptr *QTextBrowser) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTextBrowser::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QTextBrowser::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQTextBrowserMouseMoveEvent
func callbackQTextBrowserMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTextBrowser::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QTextBrowser::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQTextBrowserMousePressEvent
func callbackQTextBrowserMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTextBrowser::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QTextBrowser::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQTextBrowserMouseReleaseEvent
func callbackQTextBrowserMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QTextBrowser::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QTextBrowser::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQTextBrowserPaintEvent
func callbackQTextBrowserPaintEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectReload(f func()) {
	defer qt.Recovering("connect QTextBrowser::reload")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reload", f)
	}
}

func (ptr *QTextBrowser) DisconnectReload() {
	defer qt.Recovering("disconnect QTextBrowser::reload")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reload")
	}
}

//export callbackQTextBrowserReload
func callbackQTextBrowserReload(ptrName *C.char) bool {
	defer qt.Recovering("callback QTextBrowser::reload")

	if signal := qt.GetSignal(C.GoString(ptrName), "reload"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectSourceChanged(f func(src *core.QUrl)) {
	defer qt.Recovering("connect QTextBrowser::sourceChanged")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QTextBrowser) DisconnectSourceChanged() {
	defer qt.Recovering("disconnect QTextBrowser::sourceChanged")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQTextBrowserSourceChanged
func callbackQTextBrowserSourceChanged(ptrName *C.char, src unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::sourceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sourceChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(src))
	}

}
