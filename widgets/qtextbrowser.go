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

func (ptr *QTextBrowser) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QTextBrowser::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QTextBrowser::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQTextBrowserContextMenuEvent
func callbackQTextBrowserContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QTextBrowser::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QTextBrowser::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQTextBrowserChangeEvent
func callbackQTextBrowserChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectDragEnterEvent(f func(e *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QTextBrowser::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QTextBrowser::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQTextBrowserDragEnterEvent
func callbackQTextBrowserDragEnterEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QTextBrowser::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QTextBrowser::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQTextBrowserDragLeaveEvent
func callbackQTextBrowserDragLeaveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QTextBrowser::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QTextBrowser::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQTextBrowserDragMoveEvent
func callbackQTextBrowserDragMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QTextBrowser::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QTextBrowser::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQTextBrowserDropEvent
func callbackQTextBrowserDropEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTextBrowser::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QTextBrowser::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQTextBrowserFocusInEvent
func callbackQTextBrowserFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectInputMethodEvent(f func(e *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QTextBrowser::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QTextBrowser::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQTextBrowserInputMethodEvent
func callbackQTextBrowserInputMethodEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectInsertFromMimeData(f func(source *core.QMimeData)) {
	defer qt.Recovering("connect QTextBrowser::insertFromMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "insertFromMimeData", f)
	}
}

func (ptr *QTextBrowser) DisconnectInsertFromMimeData() {
	defer qt.Recovering("disconnect QTextBrowser::insertFromMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "insertFromMimeData")
	}
}

//export callbackQTextBrowserInsertFromMimeData
func callbackQTextBrowserInsertFromMimeData(ptrName *C.char, source unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::insertFromMimeData")

	if signal := qt.GetSignal(C.GoString(ptrName), "insertFromMimeData"); signal != nil {
		signal.(func(*core.QMimeData))(core.NewQMimeDataFromPointer(source))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTextBrowser::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QTextBrowser::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQTextBrowserKeyReleaseEvent
func callbackQTextBrowserKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTextBrowser::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QTextBrowser::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQTextBrowserMouseDoubleClickEvent
func callbackQTextBrowserMouseDoubleClickEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QTextBrowser::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QTextBrowser::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQTextBrowserResizeEvent
func callbackQTextBrowserResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QTextBrowser::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QTextBrowser) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QTextBrowser::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQTextBrowserScrollContentsBy
func callbackQTextBrowserScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QTextBrowser::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QTextBrowser::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QTextBrowser::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQTextBrowserShowEvent
func callbackQTextBrowserShowEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QTextBrowser::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QTextBrowser::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQTextBrowserWheelEvent
func callbackQTextBrowserWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QTextBrowser::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QTextBrowser) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QTextBrowser::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQTextBrowserSetupViewport
func callbackQTextBrowserSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QTextBrowser::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QTextBrowser::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQTextBrowserActionEvent
func callbackQTextBrowserActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextBrowser::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QTextBrowser::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQTextBrowserEnterEvent
func callbackQTextBrowserEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QTextBrowser::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QTextBrowser::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQTextBrowserHideEvent
func callbackQTextBrowserHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextBrowser::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QTextBrowser::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQTextBrowserLeaveEvent
func callbackQTextBrowserLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QTextBrowser::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QTextBrowser::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQTextBrowserMoveEvent
func callbackQTextBrowserMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QTextBrowser::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QTextBrowser) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QTextBrowser::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQTextBrowserSetVisible
func callbackQTextBrowserSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QTextBrowser::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QTextBrowser::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QTextBrowser::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQTextBrowserCloseEvent
func callbackQTextBrowserCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QTextBrowser::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QTextBrowser) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QTextBrowser::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQTextBrowserInitPainter
func callbackQTextBrowserInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QTextBrowser::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QTextBrowser::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQTextBrowserTabletEvent
func callbackQTextBrowserTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTextBrowser::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTextBrowser::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTextBrowserTimerEvent
func callbackQTextBrowserTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTextBrowser::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTextBrowser::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTextBrowserChildEvent
func callbackQTextBrowserChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextBrowser) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextBrowser::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTextBrowser) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTextBrowser::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTextBrowserCustomEvent
func callbackQTextBrowserCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
