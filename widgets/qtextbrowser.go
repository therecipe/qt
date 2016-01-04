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
func callbackQTextBrowserSetSource(ptr unsafe.Pointer, ptrName *C.char, name unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextBrowser::setSource")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSource"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(name))
		return true
	}
	return false

}

func (ptr *QTextBrowser) SetSource(name core.QUrl_ITF) {
	defer qt.Recovering("QTextBrowser::setSource")

	if ptr.Pointer() != nil {
		C.QTextBrowser_SetSource(ptr.Pointer(), core.PointerFromQUrl(name))
	}
}

func (ptr *QTextBrowser) SetSourceDefault(name core.QUrl_ITF) {
	defer qt.Recovering("QTextBrowser::setSource")

	if ptr.Pointer() != nil {
		C.QTextBrowser_SetSourceDefault(ptr.Pointer(), core.PointerFromQUrl(name))
	}
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
func callbackQTextBrowserAnchorClicked(ptr unsafe.Pointer, ptrName *C.char, link unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::anchorClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "anchorClicked"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QTextBrowser) AnchorClicked(link core.QUrl_ITF) {
	defer qt.Recovering("QTextBrowser::anchorClicked")

	if ptr.Pointer() != nil {
		C.QTextBrowser_AnchorClicked(ptr.Pointer(), core.PointerFromQUrl(link))
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
func callbackQTextBrowserBackward(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QTextBrowser::backward")

	if signal := qt.GetSignal(C.GoString(ptrName), "backward"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTextBrowser) Backward() {
	defer qt.Recovering("QTextBrowser::backward")

	if ptr.Pointer() != nil {
		C.QTextBrowser_Backward(ptr.Pointer())
	}
}

func (ptr *QTextBrowser) BackwardDefault() {
	defer qt.Recovering("QTextBrowser::backward")

	if ptr.Pointer() != nil {
		C.QTextBrowser_BackwardDefault(ptr.Pointer())
	}
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
func callbackQTextBrowserBackwardAvailable(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QTextBrowser::backwardAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "backwardAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QTextBrowser) BackwardAvailable(available bool) {
	defer qt.Recovering("QTextBrowser::backwardAvailable")

	if ptr.Pointer() != nil {
		C.QTextBrowser_BackwardAvailable(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
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

func (ptr *QTextBrowser) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QTextBrowser::event")

	if ptr.Pointer() != nil {
		return C.QTextBrowser_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QTextBrowser) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QTextBrowser::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QTextBrowser_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
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
func callbackQTextBrowserFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQTextBrowserFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QTextBrowser) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTextBrowser::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QTextBrowser) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTextBrowser::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
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
func callbackQTextBrowserForward(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QTextBrowser::forward")

	if signal := qt.GetSignal(C.GoString(ptrName), "forward"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTextBrowser) Forward() {
	defer qt.Recovering("QTextBrowser::forward")

	if ptr.Pointer() != nil {
		C.QTextBrowser_Forward(ptr.Pointer())
	}
}

func (ptr *QTextBrowser) ForwardDefault() {
	defer qt.Recovering("QTextBrowser::forward")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ForwardDefault(ptr.Pointer())
	}
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
func callbackQTextBrowserForwardAvailable(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QTextBrowser::forwardAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "forwardAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QTextBrowser) ForwardAvailable(available bool) {
	defer qt.Recovering("QTextBrowser::forwardAvailable")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ForwardAvailable(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
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
func callbackQTextBrowserHighlighted2(ptr unsafe.Pointer, ptrName *C.char, link *C.char) {
	defer qt.Recovering("callback QTextBrowser::highlighted")

	if signal := qt.GetSignal(C.GoString(ptrName), "highlighted2"); signal != nil {
		signal.(func(string))(C.GoString(link))
	}

}

func (ptr *QTextBrowser) Highlighted2(link string) {
	defer qt.Recovering("QTextBrowser::highlighted")

	if ptr.Pointer() != nil {
		C.QTextBrowser_Highlighted2(ptr.Pointer(), C.CString(link))
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
func callbackQTextBrowserHighlighted(ptr unsafe.Pointer, ptrName *C.char, link unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::highlighted")

	if signal := qt.GetSignal(C.GoString(ptrName), "highlighted"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QTextBrowser) Highlighted(link core.QUrl_ITF) {
	defer qt.Recovering("QTextBrowser::highlighted")

	if ptr.Pointer() != nil {
		C.QTextBrowser_Highlighted(ptr.Pointer(), core.PointerFromQUrl(link))
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
func callbackQTextBrowserHistoryChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTextBrowser::historyChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "historyChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextBrowser) HistoryChanged() {
	defer qt.Recovering("QTextBrowser::historyChanged")

	if ptr.Pointer() != nil {
		C.QTextBrowser_HistoryChanged(ptr.Pointer())
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
func callbackQTextBrowserHome(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QTextBrowser::home")

	if signal := qt.GetSignal(C.GoString(ptrName), "home"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTextBrowser) Home() {
	defer qt.Recovering("QTextBrowser::home")

	if ptr.Pointer() != nil {
		C.QTextBrowser_Home(ptr.Pointer())
	}
}

func (ptr *QTextBrowser) HomeDefault() {
	defer qt.Recovering("QTextBrowser::home")

	if ptr.Pointer() != nil {
		C.QTextBrowser_HomeDefault(ptr.Pointer())
	}
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
func callbackQTextBrowserKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQTextBrowserFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QTextBrowser) KeyPressEvent(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTextBrowser::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QTextBrowser) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTextBrowser::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
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
func callbackQTextBrowserMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextBrowser::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextBrowser) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextBrowser::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQTextBrowserMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextBrowser::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextBrowser) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextBrowser::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQTextBrowserMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextBrowser::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextBrowser) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextBrowser::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQTextBrowserPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QTextBrowser::paintEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QTextBrowser) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QTextBrowser::paintEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
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
func callbackQTextBrowserReload(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QTextBrowser::reload")

	if signal := qt.GetSignal(C.GoString(ptrName), "reload"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTextBrowser) Reload() {
	defer qt.Recovering("QTextBrowser::reload")

	if ptr.Pointer() != nil {
		C.QTextBrowser_Reload(ptr.Pointer())
	}
}

func (ptr *QTextBrowser) ReloadDefault() {
	defer qt.Recovering("QTextBrowser::reload")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ReloadDefault(ptr.Pointer())
	}
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
func callbackQTextBrowserSourceChanged(ptr unsafe.Pointer, ptrName *C.char, src unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::sourceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sourceChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(src))
	}

}

func (ptr *QTextBrowser) SourceChanged(src core.QUrl_ITF) {
	defer qt.Recovering("QTextBrowser::sourceChanged")

	if ptr.Pointer() != nil {
		C.QTextBrowser_SourceChanged(ptr.Pointer(), core.PointerFromQUrl(src))
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
func callbackQTextBrowserContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QTextBrowser::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QTextBrowser) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QTextBrowser::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
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
func callbackQTextBrowserChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QTextBrowser::changeEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QTextBrowser) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QTextBrowser::changeEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
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
func callbackQTextBrowserDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) DragEnterEvent(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QTextBrowser::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QTextBrowser) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QTextBrowser::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
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
func callbackQTextBrowserDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QTextBrowser::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QTextBrowser) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QTextBrowser::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
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
func callbackQTextBrowserDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QTextBrowser::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QTextBrowser) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QTextBrowser::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
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
func callbackQTextBrowserDropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QTextBrowser::dropEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QTextBrowser) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QTextBrowser::dropEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
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
func callbackQTextBrowserFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTextBrowser::focusInEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QTextBrowser) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTextBrowser::focusInEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
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
func callbackQTextBrowserInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) InputMethodEvent(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QTextBrowser::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

func (ptr *QTextBrowser) InputMethodEventDefault(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QTextBrowser::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
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
func callbackQTextBrowserInsertFromMimeData(ptr unsafe.Pointer, ptrName *C.char, source unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::insertFromMimeData")

	if signal := qt.GetSignal(C.GoString(ptrName), "insertFromMimeData"); signal != nil {
		signal.(func(*core.QMimeData))(core.NewQMimeDataFromPointer(source))
	} else {
		NewQTextBrowserFromPointer(ptr).InsertFromMimeDataDefault(core.NewQMimeDataFromPointer(source))
	}
}

func (ptr *QTextBrowser) InsertFromMimeData(source core.QMimeData_ITF) {
	defer qt.Recovering("QTextBrowser::insertFromMimeData")

	if ptr.Pointer() != nil {
		C.QTextBrowser_InsertFromMimeData(ptr.Pointer(), core.PointerFromQMimeData(source))
	}
}

func (ptr *QTextBrowser) InsertFromMimeDataDefault(source core.QMimeData_ITF) {
	defer qt.Recovering("QTextBrowser::insertFromMimeData")

	if ptr.Pointer() != nil {
		C.QTextBrowser_InsertFromMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(source))
	}
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
func callbackQTextBrowserKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTextBrowser::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QTextBrowser) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTextBrowser::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQTextBrowserMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextBrowser::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextBrowser) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextBrowser::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQTextBrowserResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QTextBrowser::resizeEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QTextBrowser) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QTextBrowser::resizeEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
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
func callbackQTextBrowserScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QTextBrowser::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQTextBrowserFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QTextBrowser) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QTextBrowser::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QTextBrowser) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QTextBrowser::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
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
func callbackQTextBrowserShowEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
	} else {
		NewQTextBrowserFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(v))
	}
}

func (ptr *QTextBrowser) ShowEvent(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QTextBrowser::showEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QTextBrowser) ShowEventDefault(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QTextBrowser::showEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
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
func callbackQTextBrowserWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQTextBrowserFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QTextBrowser) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QTextBrowser::wheelEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QTextBrowser) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QTextBrowser::wheelEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
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
func callbackQTextBrowserSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQTextBrowserFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QTextBrowser) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QTextBrowser::setupViewport")

	if ptr.Pointer() != nil {
		C.QTextBrowser_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QTextBrowser) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QTextBrowser::setupViewport")

	if ptr.Pointer() != nil {
		C.QTextBrowser_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
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
func callbackQTextBrowserActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QTextBrowser::actionEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QTextBrowser) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QTextBrowser::actionEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQTextBrowserEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTextBrowser::enterEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTextBrowser) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTextBrowser::enterEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQTextBrowserHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QTextBrowser::hideEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QTextBrowser) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QTextBrowser::hideEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
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
func callbackQTextBrowserLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTextBrowser::leaveEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTextBrowser) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTextBrowser::leaveEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQTextBrowserMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QTextBrowser::moveEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QTextBrowser) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QTextBrowser::moveEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
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
func callbackQTextBrowserSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QTextBrowser::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QTextBrowser) SetVisible(visible bool) {
	defer qt.Recovering("QTextBrowser::setVisible")

	if ptr.Pointer() != nil {
		C.QTextBrowser_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTextBrowser) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QTextBrowser::setVisible")

	if ptr.Pointer() != nil {
		C.QTextBrowser_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQTextBrowserCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QTextBrowser::closeEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QTextBrowser) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QTextBrowser::closeEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
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
func callbackQTextBrowserInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQTextBrowserFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QTextBrowser) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QTextBrowser::initPainter")

	if ptr.Pointer() != nil {
		C.QTextBrowser_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QTextBrowser) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QTextBrowser::initPainter")

	if ptr.Pointer() != nil {
		C.QTextBrowser_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
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
func callbackQTextBrowserTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QTextBrowser::tabletEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QTextBrowser) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QTextBrowser::tabletEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
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
func callbackQTextBrowserTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTextBrowser::timerEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTextBrowser) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTextBrowser::timerEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQTextBrowserChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTextBrowser::childEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTextBrowser) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTextBrowser::childEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQTextBrowserCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextBrowser::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTextBrowserFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextBrowser) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTextBrowser::customEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTextBrowser) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTextBrowser::customEvent")

	if ptr.Pointer() != nil {
		C.QTextBrowser_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
