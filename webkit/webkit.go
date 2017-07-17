// +build !minimal

package webkit

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "webkit.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/network"
	"github.com/therecipe/qt/printsupport"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtWebKit_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QGraphicsWebView struct {
	widgets.QGraphicsWidget
}

type QGraphicsWebView_ITF interface {
	widgets.QGraphicsWidget_ITF
	QGraphicsWebView_PTR() *QGraphicsWebView
}

func (ptr *QGraphicsWebView) QGraphicsWebView_PTR() *QGraphicsWebView {
	return ptr
}

func (ptr *QGraphicsWebView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsWebView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsWidget_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsWebView(ptr QGraphicsWebView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsWebView_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsWebViewFromPointer(ptr unsafe.Pointer) *QGraphicsWebView {
	var n = new(QGraphicsWebView)
	n.SetPointer(ptr)
	return n
}
func NewQGraphicsWebView(parent widgets.QGraphicsItem_ITF) *QGraphicsWebView {
	var tmpValue = NewQGraphicsWebViewFromPointer(C.QGraphicsWebView_NewQGraphicsWebView(widgets.PointerFromQGraphicsItem(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGraphicsWebView_ItemChange
func callbackQGraphicsWebView_ItemChange(ptr unsafe.Pointer, change C.longlong, value unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemChange"); signal != nil {
		return core.PointerFromQVariant(signal.(func(widgets.QGraphicsItem__GraphicsItemChange, *core.QVariant) *core.QVariant)(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
	}

	return core.PointerFromQVariant(NewQGraphicsWebViewFromPointer(ptr).ItemChangeDefault(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
}

func (ptr *QGraphicsWebView) ItemChangeDefault(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QGraphicsWebView_ItemChangeDefault(ptr.Pointer(), C.longlong(change), core.PointerFromQVariant(value)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsWebView_Event
func callbackQGraphicsWebView_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsWebView) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) FindText(subString string, options QWebPage__FindFlag) bool {
	if ptr.Pointer() != nil {
		var subStringC *C.char
		if subString != "" {
			subStringC = C.CString(subString)
			defer C.free(unsafe.Pointer(subStringC))
		}
		return C.QGraphicsWebView_FindText(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: subStringC, len: C.longlong(len(subString))}, C.longlong(options)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_FocusNextPrevChild
func callbackQGraphicsWebView_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QGraphicsWebView) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQGraphicsWebView_SceneEvent
func callbackQGraphicsWebView_SceneEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "sceneEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).SceneEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsWebView) SceneEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_SceneEventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_Back
func callbackQGraphicsWebView_Back(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "back"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).BackDefault()
	}
}

func (ptr *QGraphicsWebView) ConnectBack(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "back"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "back", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "back", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectBack() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "back")
	}
}

func (ptr *QGraphicsWebView) Back() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Back(ptr.Pointer())
	}
}

func (ptr *QGraphicsWebView) BackDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_BackDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_ContextMenuEvent
func callbackQGraphicsWebView_ContextMenuEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneContextMenuEvent))(widgets.NewQGraphicsSceneContextMenuEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ContextMenuEventDefault(widgets.NewQGraphicsSceneContextMenuEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ContextMenuEventDefault(ev widgets.QGraphicsSceneContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ContextMenuEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneContextMenuEvent(ev))
	}
}

//export callbackQGraphicsWebView_DragEnterEvent
func callbackQGraphicsWebView_DragEnterEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DragEnterEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) DragEnterEventDefault(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DragEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

//export callbackQGraphicsWebView_DragLeaveEvent
func callbackQGraphicsWebView_DragLeaveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DragLeaveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) DragLeaveEventDefault(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DragLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

//export callbackQGraphicsWebView_DragMoveEvent
func callbackQGraphicsWebView_DragMoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DragMoveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) DragMoveEventDefault(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DragMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

//export callbackQGraphicsWebView_DropEvent
func callbackQGraphicsWebView_DropEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DropEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) DropEventDefault(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DropEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

//export callbackQGraphicsWebView_FocusInEvent
func callbackQGraphicsWebView_FocusInEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQGraphicsWebView_FocusOutEvent
func callbackQGraphicsWebView_FocusOutEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQGraphicsWebView_Forward
func callbackQGraphicsWebView_Forward(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "forward"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ForwardDefault()
	}
}

func (ptr *QGraphicsWebView) ConnectForward(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "forward"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "forward", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "forward", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectForward() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "forward")
	}
}

func (ptr *QGraphicsWebView) Forward() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Forward(ptr.Pointer())
	}
}

func (ptr *QGraphicsWebView) ForwardDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ForwardDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_HoverLeaveEvent
func callbackQGraphicsWebView_HoverLeaveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverLeaveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).HoverLeaveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) HoverLeaveEventDefault(ev widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HoverLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(ev))
	}
}

//export callbackQGraphicsWebView_HoverMoveEvent
func callbackQGraphicsWebView_HoverMoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).HoverMoveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) HoverMoveEventDefault(ev widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HoverMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(ev))
	}
}

//export callbackQGraphicsWebView_IconChanged
func callbackQGraphicsWebView_IconChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsWebView) ConnectIconChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "iconChanged") {
			C.QGraphicsWebView_ConnectIconChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "iconChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "iconChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "iconChanged", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectIconChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectIconChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "iconChanged")
	}
}

func (ptr *QGraphicsWebView) IconChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_IconChanged(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_InputMethodEvent
func callbackQGraphicsWebView_InputMethodEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) InputMethodEventDefault(ev gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(ev))
	}
}

//export callbackQGraphicsWebView_KeyPressEvent
func callbackQGraphicsWebView_KeyPressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackQGraphicsWebView_KeyReleaseEvent
func callbackQGraphicsWebView_KeyReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) KeyReleaseEventDefault(ev gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackQGraphicsWebView_LinkClicked
func callbackQGraphicsWebView_LinkClicked(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "linkClicked"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QGraphicsWebView) ConnectLinkClicked(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linkClicked") {
			C.QGraphicsWebView_ConnectLinkClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linkClicked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "linkClicked", func(url *core.QUrl) {
				signal.(func(*core.QUrl))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linkClicked", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectLinkClicked() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectLinkClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "linkClicked")
	}
}

func (ptr *QGraphicsWebView) LinkClicked(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_LinkClicked(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QGraphicsWebView) Load2(request network.QNetworkRequest_ITF, operation network.QNetworkAccessManager__Operation, body core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Load2(ptr.Pointer(), network.PointerFromQNetworkRequest(request), C.longlong(operation), core.PointerFromQByteArray(body))
	}
}

func (ptr *QGraphicsWebView) Load(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQGraphicsWebView_LoadFinished
func callbackQGraphicsWebView_LoadFinished(ptr unsafe.Pointer, ok C.char) {
	if signal := qt.GetSignal(ptr, "loadFinished"); signal != nil {
		signal.(func(bool))(int8(ok) != 0)
	}

}

func (ptr *QGraphicsWebView) ConnectLoadFinished(f func(ok bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadFinished") {
			C.QGraphicsWebView_ConnectLoadFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", func(ok bool) {
				signal.(func(bool))(ok)
				f(ok)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectLoadFinished() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadFinished")
	}
}

func (ptr *QGraphicsWebView) LoadFinished(ok bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_LoadFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))
	}
}

//export callbackQGraphicsWebView_LoadProgress
func callbackQGraphicsWebView_LoadProgress(ptr unsafe.Pointer, progress C.int) {
	if signal := qt.GetSignal(ptr, "loadProgress"); signal != nil {
		signal.(func(int))(int(int32(progress)))
	}

}

func (ptr *QGraphicsWebView) ConnectLoadProgress(f func(progress int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadProgress") {
			C.QGraphicsWebView_ConnectLoadProgress(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadProgress"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadProgress", func(progress int) {
				signal.(func(int))(progress)
				f(progress)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadProgress", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectLoadProgress() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadProgress")
	}
}

func (ptr *QGraphicsWebView) LoadProgress(progress int) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_LoadProgress(ptr.Pointer(), C.int(int32(progress)))
	}
}

//export callbackQGraphicsWebView_LoadStarted
func callbackQGraphicsWebView_LoadStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsWebView) ConnectLoadStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadStarted") {
			C.QGraphicsWebView_ConnectLoadStarted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadStarted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectLoadStarted() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadStarted")
	}
}

func (ptr *QGraphicsWebView) LoadStarted() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_LoadStarted(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_MouseDoubleClickEvent
func callbackQGraphicsWebView_MouseDoubleClickEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).MouseDoubleClickEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) MouseDoubleClickEventDefault(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MouseDoubleClickEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

//export callbackQGraphicsWebView_MouseMoveEvent
func callbackQGraphicsWebView_MouseMoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).MouseMoveEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) MouseMoveEventDefault(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MouseMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

//export callbackQGraphicsWebView_MousePressEvent
func callbackQGraphicsWebView_MousePressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).MousePressEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) MousePressEventDefault(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MousePressEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

//export callbackQGraphicsWebView_MouseReleaseEvent
func callbackQGraphicsWebView_MouseReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).MouseReleaseEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) MouseReleaseEventDefault(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MouseReleaseEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

//export callbackQGraphicsWebView_Paint
func callbackQGraphicsWebView_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsWebView) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQGraphicsWebView_Reload
func callbackQGraphicsWebView_Reload(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reload"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ReloadDefault()
	}
}

func (ptr *QGraphicsWebView) ConnectReload(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reload"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "reload", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reload", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectReload() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reload")
	}
}

func (ptr *QGraphicsWebView) Reload() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Reload(ptr.Pointer())
	}
}

func (ptr *QGraphicsWebView) ReloadDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ReloadDefault(ptr.Pointer())
	}
}

func (ptr *QGraphicsWebView) SetContent(data core.QByteArray_ITF, mimeType string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var mimeTypeC *C.char
		if mimeType != "" {
			mimeTypeC = C.CString(mimeType)
			defer C.free(unsafe.Pointer(mimeTypeC))
		}
		C.QGraphicsWebView_SetContent(ptr.Pointer(), core.PointerFromQByteArray(data), C.struct_QtWebKit_PackedString{data: mimeTypeC, len: C.longlong(len(mimeType))}, core.PointerFromQUrl(baseUrl))
	}
}

//export callbackQGraphicsWebView_SetGeometry
func callbackQGraphicsWebView_SetGeometry(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		signal.(func(*core.QRectF))(core.NewQRectFFromPointer(rect))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).SetGeometryDefault(core.NewQRectFFromPointer(rect))
	}
}

func (ptr *QGraphicsWebView) SetGeometryDefault(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsWebView) SetHtml(html string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var htmlC *C.char
		if html != "" {
			htmlC = C.CString(html)
			defer C.free(unsafe.Pointer(htmlC))
		}
		C.QGraphicsWebView_SetHtml(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: htmlC, len: C.longlong(len(html))}, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QGraphicsWebView) SetPage(page QWebPage_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetPage(ptr.Pointer(), PointerFromQWebPage(page))
	}
}

func (ptr *QGraphicsWebView) SetRenderHint(hint gui.QPainter__RenderHint, enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetRenderHint(ptr.Pointer(), C.longlong(hint), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QGraphicsWebView) SetRenderHints(hints gui.QPainter__RenderHint) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetRenderHints(ptr.Pointer(), C.longlong(hints))
	}
}

func (ptr *QGraphicsWebView) SetResizesToContents(enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetResizesToContents(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QGraphicsWebView) SetTiledBackingStoreFrozen(frozen bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetTiledBackingStoreFrozen(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(frozen))))
	}
}

func (ptr *QGraphicsWebView) SetUrl(vqu core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetUrl(ptr.Pointer(), core.PointerFromQUrl(vqu))
	}
}

func (ptr *QGraphicsWebView) SetZoomFactor(vqr float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetZoomFactor(ptr.Pointer(), C.double(vqr))
	}
}

//export callbackQGraphicsWebView_StatusBarMessage
func callbackQGraphicsWebView_StatusBarMessage(ptr unsafe.Pointer, text C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "statusBarMessage"); signal != nil {
		signal.(func(string))(cGoUnpackString(text))
	}

}

func (ptr *QGraphicsWebView) ConnectStatusBarMessage(f func(text string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusBarMessage") {
			C.QGraphicsWebView_ConnectStatusBarMessage(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "statusBarMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "statusBarMessage", func(text string) {
				signal.(func(string))(text)
				f(text)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusBarMessage", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectStatusBarMessage() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectStatusBarMessage(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusBarMessage")
	}
}

func (ptr *QGraphicsWebView) StatusBarMessage(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QGraphicsWebView_StatusBarMessage(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQGraphicsWebView_Stop
func callbackQGraphicsWebView_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).StopDefault()
	}
}

func (ptr *QGraphicsWebView) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stop", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QGraphicsWebView) Stop() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Stop(ptr.Pointer())
	}
}

func (ptr *QGraphicsWebView) StopDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_StopDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_TitleChanged
func callbackQGraphicsWebView_TitleChanged(ptr unsafe.Pointer, vqs C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "titleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	}

}

func (ptr *QGraphicsWebView) ConnectTitleChanged(f func(vqs string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleChanged") {
			C.QGraphicsWebView_ConnectTitleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", func(vqs string) {
				signal.(func(string))(vqs)
				f(vqs)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleChanged")
	}
}

func (ptr *QGraphicsWebView) TitleChanged(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QGraphicsWebView_TitleChanged(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QGraphicsWebView) TriggerPageAction(action QWebPage__WebAction, checked bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_TriggerPageAction(ptr.Pointer(), C.longlong(action), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

//export callbackQGraphicsWebView_UpdateGeometry
func callbackQGraphicsWebView_UpdateGeometry(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateGeometry"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).UpdateGeometryDefault()
	}
}

func (ptr *QGraphicsWebView) UpdateGeometryDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UpdateGeometryDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_UrlChanged
func callbackQGraphicsWebView_UrlChanged(ptr unsafe.Pointer, vqu unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "urlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(vqu))
	}

}

func (ptr *QGraphicsWebView) ConnectUrlChanged(f func(vqu *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "urlChanged") {
			C.QGraphicsWebView_ConnectUrlChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "urlChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "urlChanged", func(vqu *core.QUrl) {
				signal.(func(*core.QUrl))(vqu)
				f(vqu)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "urlChanged", f)
		}
	}
}

func (ptr *QGraphicsWebView) DisconnectUrlChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "urlChanged")
	}
}

func (ptr *QGraphicsWebView) UrlChanged(vqu core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(vqu))
	}
}

//export callbackQGraphicsWebView_WheelEvent
func callbackQGraphicsWebView_WheelEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneWheelEvent))(widgets.NewQGraphicsSceneWheelEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).WheelEventDefault(widgets.NewQGraphicsSceneWheelEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) WheelEventDefault(ev widgets.QGraphicsSceneWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_WheelEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneWheelEvent(ev))
	}
}

func (ptr *QGraphicsWebView) DestroyQGraphicsWebView() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DestroyQGraphicsWebView(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGraphicsWebView) PageAction(action QWebPage__WebAction) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QGraphicsWebView_PageAction(ptr.Pointer(), C.longlong(action)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQIconFromPointer(C.QGraphicsWebView_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) RenderHints() gui.QPainter__RenderHint {
	if ptr.Pointer() != nil {
		return gui.QPainter__RenderHint(C.QGraphicsWebView_RenderHints(ptr.Pointer()))
	}
	return 0
}

//export callbackQGraphicsWebView_SizeHint
func callbackQGraphicsWebView_SizeHint(ptr unsafe.Pointer, which C.longlong, constraint unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSizeF(signal.(func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
	}

	return core.PointerFromQSizeF(NewQGraphicsWebViewFromPointer(ptr).SizeHintDefault(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
}

func (ptr *QGraphicsWebView) SizeHintDefault(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFFromPointer(C.QGraphicsWebView_SizeHintDefault(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		runtime.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGraphicsWebView_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsWebView) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QGraphicsWebView_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsWebView_InputMethodQuery
func callbackQGraphicsWebView_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQGraphicsWebViewFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QGraphicsWebView) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QGraphicsWebView_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) History() *QWebHistory {
	if ptr.Pointer() != nil {
		return NewQWebHistoryFromPointer(C.QGraphicsWebView_History(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWebView) Page() *QWebPage {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebPageFromPointer(C.QGraphicsWebView_Page(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) Settings() *QWebSettings {
	if ptr.Pointer() != nil {
		return NewQWebSettingsFromPointer(C.QGraphicsWebView_Settings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWebView) IsModified() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) IsTiledBackingStoreFrozen() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_IsTiledBackingStoreFrozen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) ResizesToContents() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_ResizesToContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) ZoomFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsWebView_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWebView) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QGraphicsWebView___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QGraphicsWebView) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QGraphicsWebView___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QGraphicsWebView) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QGraphicsWebView___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QGraphicsWebView) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___actions_newList(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QGraphicsWebView___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGraphicsWebView) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGraphicsWebView___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsWebView) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGraphicsWebView___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsWebView) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGraphicsWebView___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsWebView) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___findChildren_newList(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGraphicsWebView___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsWebView) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___children_newList(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __setTransformations_transformations_atList(i int) *widgets.QGraphicsTransform {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQGraphicsTransformFromPointer(C.QGraphicsWebView___setTransformations_transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) __setTransformations_transformations_setList(i widgets.QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___setTransformations_transformations_setList(ptr.Pointer(), widgets.PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QGraphicsWebView) __setTransformations_transformations_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___setTransformations_transformations_newList(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __childItems_atList(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QGraphicsWebView___childItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QGraphicsWebView) __childItems_setList(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___childItems_setList(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QGraphicsWebView) __childItems_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___childItems_newList(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __collidingItems_atList(i int) *widgets.QGraphicsItem {
	if ptr.Pointer() != nil {
		return widgets.NewQGraphicsItemFromPointer(C.QGraphicsWebView___collidingItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QGraphicsWebView) __collidingItems_setList(i widgets.QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___collidingItems_setList(ptr.Pointer(), widgets.PointerFromQGraphicsItem(i))
	}
}

func (ptr *QGraphicsWebView) __collidingItems_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___collidingItems_newList(ptr.Pointer()))
}

func (ptr *QGraphicsWebView) __transformations_atList(i int) *widgets.QGraphicsTransform {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQGraphicsTransformFromPointer(C.QGraphicsWebView___transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWebView) __transformations_setList(i widgets.QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView___transformations_setList(ptr.Pointer(), widgets.PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QGraphicsWebView) __transformations_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGraphicsWebView___transformations_newList(ptr.Pointer()))
}

//export callbackQGraphicsWebView_Close
func callbackQGraphicsWebView_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).CloseDefault())))
}

func (ptr *QGraphicsWebView) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQGraphicsWebView_WindowFrameEvent
func callbackQGraphicsWebView_WindowFrameEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "windowFrameEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).WindowFrameEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsWebView) WindowFrameEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_WindowFrameEventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_ChangeEvent
func callbackQGraphicsWebView_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_CloseEvent
func callbackQGraphicsWebView_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQGraphicsWebView_GeometryChanged
func callbackQGraphicsWebView_GeometryChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "geometryChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsWebView_GrabKeyboardEvent
func callbackQGraphicsWebView_GrabKeyboardEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "grabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).GrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) GrabKeyboardEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_GrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_GrabMouseEvent
func callbackQGraphicsWebView_GrabMouseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "grabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).GrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) GrabMouseEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_GrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_HideEvent
func callbackQGraphicsWebView_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQGraphicsWebView_MoveEvent
func callbackQGraphicsWebView_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMoveEvent))(widgets.NewQGraphicsSceneMoveEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).MoveEventDefault(widgets.NewQGraphicsSceneMoveEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) MoveEventDefault(event widgets.QGraphicsSceneMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMoveEvent(event))
	}
}

//export callbackQGraphicsWebView_PaintWindowFrame
func callbackQGraphicsWebView_PaintWindowFrame(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintWindowFrame"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).PaintWindowFrameDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsWebView) PaintWindowFrameDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_PaintWindowFrameDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQGraphicsWebView_PolishEvent
func callbackQGraphicsWebView_PolishEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "polishEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).PolishEventDefault()
	}
}

func (ptr *QGraphicsWebView) PolishEventDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_PolishEventDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_ResizeEvent
func callbackQGraphicsWebView_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneResizeEvent))(widgets.NewQGraphicsSceneResizeEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ResizeEventDefault(widgets.NewQGraphicsSceneResizeEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ResizeEventDefault(event widgets.QGraphicsSceneResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ResizeEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneResizeEvent(event))
	}
}

//export callbackQGraphicsWebView_ShowEvent
func callbackQGraphicsWebView_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQGraphicsWebView_UngrabKeyboardEvent
func callbackQGraphicsWebView_UngrabKeyboardEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ungrabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).UngrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) UngrabKeyboardEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UngrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_UngrabMouseEvent
func callbackQGraphicsWebView_UngrabMouseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ungrabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).UngrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) UngrabMouseEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UngrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_Shape
func callbackQGraphicsWebView_Shape(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "shape"); signal != nil {
		return gui.PointerFromQPainterPath(signal.(func() *gui.QPainterPath)())
	}

	return gui.PointerFromQPainterPath(NewQGraphicsWebViewFromPointer(ptr).ShapeDefault())
}

func (ptr *QGraphicsWebView) ShapeDefault() *gui.QPainterPath {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPainterPathFromPointer(C.QGraphicsWebView_ShapeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsWebView_BoundingRect
func callbackQGraphicsWebView_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQGraphicsWebViewFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QGraphicsWebView) BoundingRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QGraphicsWebView_BoundingRectDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsWebView_WindowFrameSectionAt
func callbackQGraphicsWebView_WindowFrameSectionAt(ptr unsafe.Pointer, pos unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "windowFrameSectionAt"); signal != nil {
		return C.longlong(signal.(func(*core.QPointF) core.Qt__WindowFrameSection)(core.NewQPointFFromPointer(pos)))
	}

	return C.longlong(NewQGraphicsWebViewFromPointer(ptr).WindowFrameSectionAtDefault(core.NewQPointFFromPointer(pos)))
}

func (ptr *QGraphicsWebView) WindowFrameSectionAtDefault(pos core.QPointF_ITF) core.Qt__WindowFrameSection {
	if ptr.Pointer() != nil {
		return core.Qt__WindowFrameSection(C.QGraphicsWebView_WindowFrameSectionAtDefault(ptr.Pointer(), core.PointerFromQPointF(pos)))
	}
	return 0
}

//export callbackQGraphicsWebView_Type
func callbackQGraphicsWebView_Type(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQGraphicsWebViewFromPointer(ptr).TypeDefault()))
}

func (ptr *QGraphicsWebView) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGraphicsWebView_TypeDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQGraphicsWebView_GetContentsMargins
func callbackQGraphicsWebView_GetContentsMargins(ptr unsafe.Pointer, left C.double, top C.double, right C.double, bottom C.double) {
	if signal := qt.GetSignal(ptr, "getContentsMargins"); signal != nil {
		signal.(func(float64, float64, float64, float64))(float64(left), float64(top), float64(right), float64(bottom))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).GetContentsMarginsDefault(float64(left), float64(top), float64(right), float64(bottom))
	}
}

func (ptr *QGraphicsWebView) GetContentsMarginsDefault(left float64, top float64, right float64, bottom float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_GetContentsMarginsDefault(ptr.Pointer(), C.double(left), C.double(top), C.double(right), C.double(bottom))
	}
}

//export callbackQGraphicsWebView_InitStyleOption
func callbackQGraphicsWebView_InitStyleOption(ptr unsafe.Pointer, option unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initStyleOption"); signal != nil {
		signal.(func(*widgets.QStyleOption))(widgets.NewQStyleOptionFromPointer(option))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).InitStyleOptionDefault(widgets.NewQStyleOptionFromPointer(option))
	}
}

func (ptr *QGraphicsWebView) InitStyleOptionDefault(option widgets.QStyleOption_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_InitStyleOptionDefault(ptr.Pointer(), widgets.PointerFromQStyleOption(option))
	}
}

//export callbackQGraphicsWebView_EnabledChanged
func callbackQGraphicsWebView_EnabledChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enabledChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsWebView_OpacityChanged
func callbackQGraphicsWebView_OpacityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsWebView_ParentChanged
func callbackQGraphicsWebView_ParentChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "parentChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsWebView_RotationChanged
func callbackQGraphicsWebView_RotationChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rotationChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsWebView_ScaleChanged
func callbackQGraphicsWebView_ScaleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scaleChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsWebView_UpdateMicroFocus
func callbackQGraphicsWebView_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QGraphicsWebView) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_VisibleChanged
func callbackQGraphicsWebView_VisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsWebView_XChanged
func callbackQGraphicsWebView_XChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "xChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsWebView_YChanged
func callbackQGraphicsWebView_YChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "yChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsWebView_ZChanged
func callbackQGraphicsWebView_ZChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "zChanged"); signal != nil {
		signal.(func())()
	}

}

//export callbackQGraphicsWebView_EventFilter
func callbackQGraphicsWebView_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsWebView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_ChildEvent
func callbackQGraphicsWebView_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGraphicsWebView_ConnectNotify
func callbackQGraphicsWebView_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsWebView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsWebView_CustomEvent
func callbackQGraphicsWebView_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_DeleteLater
func callbackQGraphicsWebView_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGraphicsWebView) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGraphicsWebView_Destroyed
func callbackQGraphicsWebView_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGraphicsWebView_DisconnectNotify
func callbackQGraphicsWebView_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsWebView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsWebView_ObjectNameChanged
func callbackQGraphicsWebView_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQGraphicsWebView_TimerEvent
func callbackQGraphicsWebView_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGraphicsWebView_MetaObject
func callbackQGraphicsWebView_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGraphicsWebViewFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGraphicsWebView) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGraphicsWebView_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsWebView_SceneEventFilter
func callbackQGraphicsWebView_SceneEventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "sceneEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem, *core.QEvent) bool)(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).SceneEventFilterDefault(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsWebView) SceneEventFilterDefault(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_SceneEventFilterDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_Advance
func callbackQGraphicsWebView_Advance(ptr unsafe.Pointer, phase C.int) {
	if signal := qt.GetSignal(ptr, "advance"); signal != nil {
		signal.(func(int))(int(int32(phase)))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).AdvanceDefault(int(int32(phase)))
	}
}

func (ptr *QGraphicsWebView) AdvanceDefault(phase int) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_AdvanceDefault(ptr.Pointer(), C.int(int32(phase)))
	}
}

//export callbackQGraphicsWebView_HoverEnterEvent
func callbackQGraphicsWebView_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverEnterEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).HoverEnterEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) HoverEnterEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HoverEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsWebView_OpaqueArea
func callbackQGraphicsWebView_OpaqueArea(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "opaqueArea"); signal != nil {
		return gui.PointerFromQPainterPath(signal.(func() *gui.QPainterPath)())
	}

	return gui.PointerFromQPainterPath(NewQGraphicsWebViewFromPointer(ptr).OpaqueAreaDefault())
}

func (ptr *QGraphicsWebView) OpaqueAreaDefault() *gui.QPainterPath {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPainterPathFromPointer(C.QGraphicsWebView_OpaqueAreaDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsWebView_CollidesWithItem
func callbackQGraphicsWebView_CollidesWithItem(ptr unsafe.Pointer, other unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "collidesWithItem"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem, core.Qt__ItemSelectionMode) bool)(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).CollidesWithItemDefault(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode)))))
}

func (ptr *QGraphicsWebView) CollidesWithItemDefault(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_CollidesWithItemDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(other), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_CollidesWithPath
func callbackQGraphicsWebView_CollidesWithPath(ptr unsafe.Pointer, path unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "collidesWithPath"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*gui.QPainterPath, core.Qt__ItemSelectionMode) bool)(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).CollidesWithPathDefault(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode)))))
}

func (ptr *QGraphicsWebView) CollidesWithPathDefault(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_CollidesWithPathDefault(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_Contains
func callbackQGraphicsWebView_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point)))))
}

func (ptr *QGraphicsWebView) ContainsDefault(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_IsObscuredBy
func callbackQGraphicsWebView_IsObscuredBy(ptr unsafe.Pointer, item unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isObscuredBy"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem) bool)(widgets.NewQGraphicsItemFromPointer(item)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).IsObscuredByDefault(widgets.NewQGraphicsItemFromPointer(item)))))
}

func (ptr *QGraphicsWebView) IsObscuredByDefault(item widgets.QGraphicsItem_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_IsObscuredByDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

type QWebDatabase struct {
	ptr unsafe.Pointer
}

type QWebDatabase_ITF interface {
	QWebDatabase_PTR() *QWebDatabase
}

func (ptr *QWebDatabase) QWebDatabase_PTR() *QWebDatabase {
	return ptr
}

func (ptr *QWebDatabase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebDatabase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebDatabase(ptr QWebDatabase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebDatabase_PTR().Pointer()
	}
	return nil
}

func NewQWebDatabaseFromPointer(ptr unsafe.Pointer) *QWebDatabase {
	var n = new(QWebDatabase)
	n.SetPointer(ptr)
	return n
}
func NewQWebDatabase(other QWebDatabase_ITF) *QWebDatabase {
	var tmpValue = NewQWebDatabaseFromPointer(C.QWebDatabase_NewQWebDatabase(PointerFromQWebDatabase(other)))
	runtime.SetFinalizer(tmpValue, (*QWebDatabase).DestroyQWebDatabase)
	return tmpValue
}

func QWebDatabase_RemoveAllDatabases() {
	C.QWebDatabase_QWebDatabase_RemoveAllDatabases()
}

func (ptr *QWebDatabase) RemoveAllDatabases() {
	C.QWebDatabase_QWebDatabase_RemoveAllDatabases()
}

func QWebDatabase_RemoveDatabase(db QWebDatabase_ITF) {
	C.QWebDatabase_QWebDatabase_RemoveDatabase(PointerFromQWebDatabase(db))
}

func (ptr *QWebDatabase) RemoveDatabase(db QWebDatabase_ITF) {
	C.QWebDatabase_QWebDatabase_RemoveDatabase(PointerFromQWebDatabase(db))
}

func (ptr *QWebDatabase) DestroyQWebDatabase() {
	if ptr.Pointer() != nil {
		C.QWebDatabase_DestroyQWebDatabase(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebDatabase) DisplayName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebDatabase_DisplayName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebDatabase) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebDatabase_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebDatabase) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebDatabase_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebDatabase) Origin() *QWebSecurityOrigin {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebSecurityOriginFromPointer(C.QWebDatabase_Origin(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebSecurityOrigin).DestroyQWebSecurityOrigin)
		return tmpValue
	}
	return nil
}

func (ptr *QWebDatabase) ExpectedSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebDatabase_ExpectedSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebDatabase) Size() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebDatabase_Size(ptr.Pointer()))
	}
	return 0
}

type QWebElement struct {
	ptr unsafe.Pointer
}

type QWebElement_ITF interface {
	QWebElement_PTR() *QWebElement
}

func (ptr *QWebElement) QWebElement_PTR() *QWebElement {
	return ptr
}

func (ptr *QWebElement) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebElement) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebElement(ptr QWebElement_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebElement_PTR().Pointer()
	}
	return nil
}

func NewQWebElementFromPointer(ptr unsafe.Pointer) *QWebElement {
	var n = new(QWebElement)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QWebElement__StyleResolveStrategy
//QWebElement::StyleResolveStrategy
type QWebElement__StyleResolveStrategy int64

var (
	QWebElement__InlineStyle   QWebElement__StyleResolveStrategy = QWebElement__StyleResolveStrategy(0)
	QWebElement__CascadedStyle QWebElement__StyleResolveStrategy = QWebElement__StyleResolveStrategy(1)
	QWebElement__ComputedStyle QWebElement__StyleResolveStrategy = QWebElement__StyleResolveStrategy(2)
)

func (ptr *QWebElement) EvaluateJavaScript(scriptSource string) *core.QVariant {
	if ptr.Pointer() != nil {
		var scriptSourceC *C.char
		if scriptSource != "" {
			scriptSourceC = C.CString(scriptSource)
			defer C.free(unsafe.Pointer(scriptSourceC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QWebElement_EvaluateJavaScript(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: scriptSourceC, len: C.longlong(len(scriptSource))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) TakeFromDocument() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElement_TakeFromDocument(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func NewQWebElement() *QWebElement {
	var tmpValue = NewQWebElementFromPointer(C.QWebElement_NewQWebElement())
	runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
	return tmpValue
}

func NewQWebElement2(other QWebElement_ITF) *QWebElement {
	var tmpValue = NewQWebElementFromPointer(C.QWebElement_NewQWebElement2(PointerFromQWebElement(other)))
	runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
	return tmpValue
}

func (ptr *QWebElement) AddClass(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QWebElement_AddClass(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QWebElement) AppendInside(markup string) {
	if ptr.Pointer() != nil {
		var markupC *C.char
		if markup != "" {
			markupC = C.CString(markup)
			defer C.free(unsafe.Pointer(markupC))
		}
		C.QWebElement_AppendInside(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: markupC, len: C.longlong(len(markup))})
	}
}

func (ptr *QWebElement) AppendInside2(element QWebElement_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElement_AppendInside2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) AppendOutside(markup string) {
	if ptr.Pointer() != nil {
		var markupC *C.char
		if markup != "" {
			markupC = C.CString(markup)
			defer C.free(unsafe.Pointer(markupC))
		}
		C.QWebElement_AppendOutside(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: markupC, len: C.longlong(len(markup))})
	}
}

func (ptr *QWebElement) AppendOutside2(element QWebElement_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElement_AppendOutside2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) EncloseContentsWith2(markup string) {
	if ptr.Pointer() != nil {
		var markupC *C.char
		if markup != "" {
			markupC = C.CString(markup)
			defer C.free(unsafe.Pointer(markupC))
		}
		C.QWebElement_EncloseContentsWith2(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: markupC, len: C.longlong(len(markup))})
	}
}

func (ptr *QWebElement) EncloseContentsWith(element QWebElement_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElement_EncloseContentsWith(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) EncloseWith(markup string) {
	if ptr.Pointer() != nil {
		var markupC *C.char
		if markup != "" {
			markupC = C.CString(markup)
			defer C.free(unsafe.Pointer(markupC))
		}
		C.QWebElement_EncloseWith(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: markupC, len: C.longlong(len(markup))})
	}
}

func (ptr *QWebElement) EncloseWith2(element QWebElement_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElement_EncloseWith2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) PrependInside(markup string) {
	if ptr.Pointer() != nil {
		var markupC *C.char
		if markup != "" {
			markupC = C.CString(markup)
			defer C.free(unsafe.Pointer(markupC))
		}
		C.QWebElement_PrependInside(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: markupC, len: C.longlong(len(markup))})
	}
}

func (ptr *QWebElement) PrependInside2(element QWebElement_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElement_PrependInside2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) PrependOutside(markup string) {
	if ptr.Pointer() != nil {
		var markupC *C.char
		if markup != "" {
			markupC = C.CString(markup)
			defer C.free(unsafe.Pointer(markupC))
		}
		C.QWebElement_PrependOutside(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: markupC, len: C.longlong(len(markup))})
	}
}

func (ptr *QWebElement) PrependOutside2(element QWebElement_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElement_PrependOutside2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) RemoveAllChildren() {
	if ptr.Pointer() != nil {
		C.QWebElement_RemoveAllChildren(ptr.Pointer())
	}
}

func (ptr *QWebElement) RemoveAttribute(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QWebElement_RemoveAttribute(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QWebElement) RemoveAttributeNS(namespaceUri string, name string) {
	if ptr.Pointer() != nil {
		var namespaceUriC *C.char
		if namespaceUri != "" {
			namespaceUriC = C.CString(namespaceUri)
			defer C.free(unsafe.Pointer(namespaceUriC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QWebElement_RemoveAttributeNS(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: namespaceUriC, len: C.longlong(len(namespaceUri))}, C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QWebElement) RemoveClass(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QWebElement_RemoveClass(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QWebElement) RemoveFromDocument() {
	if ptr.Pointer() != nil {
		C.QWebElement_RemoveFromDocument(ptr.Pointer())
	}
}

func (ptr *QWebElement) Render(painter gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElement_Render(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QWebElement) Render2(painter gui.QPainter_ITF, clip core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElement_Render2(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(clip))
	}
}

func (ptr *QWebElement) Replace(markup string) {
	if ptr.Pointer() != nil {
		var markupC *C.char
		if markup != "" {
			markupC = C.CString(markup)
			defer C.free(unsafe.Pointer(markupC))
		}
		C.QWebElement_Replace(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: markupC, len: C.longlong(len(markup))})
	}
}

func (ptr *QWebElement) Replace2(element QWebElement_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElement_Replace2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) SetAttribute(name string, value string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QWebElement_SetAttribute(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtWebKit_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *QWebElement) SetAttributeNS(namespaceUri string, name string, value string) {
	if ptr.Pointer() != nil {
		var namespaceUriC *C.char
		if namespaceUri != "" {
			namespaceUriC = C.CString(namespaceUri)
			defer C.free(unsafe.Pointer(namespaceUriC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QWebElement_SetAttributeNS(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: namespaceUriC, len: C.longlong(len(namespaceUri))}, C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtWebKit_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *QWebElement) SetFocus() {
	if ptr.Pointer() != nil {
		C.QWebElement_SetFocus(ptr.Pointer())
	}
}

func (ptr *QWebElement) SetInnerXml(markup string) {
	if ptr.Pointer() != nil {
		var markupC *C.char
		if markup != "" {
			markupC = C.CString(markup)
			defer C.free(unsafe.Pointer(markupC))
		}
		C.QWebElement_SetInnerXml(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: markupC, len: C.longlong(len(markup))})
	}
}

func (ptr *QWebElement) SetOuterXml(markup string) {
	if ptr.Pointer() != nil {
		var markupC *C.char
		if markup != "" {
			markupC = C.CString(markup)
			defer C.free(unsafe.Pointer(markupC))
		}
		C.QWebElement_SetOuterXml(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: markupC, len: C.longlong(len(markup))})
	}
}

func (ptr *QWebElement) SetPlainText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QWebElement_SetPlainText(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QWebElement) SetStyleProperty(name string, value string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QWebElement_SetStyleProperty(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtWebKit_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *QWebElement) ToggleClass(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QWebElement_ToggleClass(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QWebElement) DestroyQWebElement() {
	if ptr.Pointer() != nil {
		C.QWebElement_DestroyQWebElement(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebElement) Geometry() *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QWebElement_Geometry(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) Attribute(name string, defaultValue string) string {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var defaultValueC *C.char
		if defaultValue != "" {
			defaultValueC = C.CString(defaultValue)
			defer C.free(unsafe.Pointer(defaultValueC))
		}
		return cGoUnpackString(C.QWebElement_Attribute(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtWebKit_PackedString{data: defaultValueC, len: C.longlong(len(defaultValue))}))
	}
	return ""
}

func (ptr *QWebElement) AttributeNS(namespaceUri string, name string, defaultValue string) string {
	if ptr.Pointer() != nil {
		var namespaceUriC *C.char
		if namespaceUri != "" {
			namespaceUriC = C.CString(namespaceUri)
			defer C.free(unsafe.Pointer(namespaceUriC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var defaultValueC *C.char
		if defaultValue != "" {
			defaultValueC = C.CString(defaultValue)
			defer C.free(unsafe.Pointer(defaultValueC))
		}
		return cGoUnpackString(C.QWebElement_AttributeNS(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: namespaceUriC, len: C.longlong(len(namespaceUri))}, C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtWebKit_PackedString{data: defaultValueC, len: C.longlong(len(defaultValue))}))
	}
	return ""
}

func (ptr *QWebElement) LocalName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebElement_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) NamespaceUri() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebElement_NamespaceUri(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) Prefix() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebElement_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) StyleProperty(name string, strategy QWebElement__StyleResolveStrategy) string {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return cGoUnpackString(C.QWebElement_StyleProperty(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(strategy)))
	}
	return ""
}

func (ptr *QWebElement) TagName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebElement_TagName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) ToInnerXml() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebElement_ToInnerXml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) ToOuterXml() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebElement_ToOuterXml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) ToPlainText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebElement_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) AttributeNames(namespaceUri string) []string {
	if ptr.Pointer() != nil {
		var namespaceUriC *C.char
		if namespaceUri != "" {
			namespaceUriC = C.CString(namespaceUri)
			defer C.free(unsafe.Pointer(namespaceUriC))
		}
		return strings.Split(cGoUnpackString(C.QWebElement_AttributeNames(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: namespaceUriC, len: C.longlong(len(namespaceUri))})), "|")
	}
	return make([]string, 0)
}

func (ptr *QWebElement) Classes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QWebElement_Classes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QWebElement) Clone() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElement_Clone(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) Document() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElement_Document(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) FindFirst(selectorQuery string) *QWebElement {
	if ptr.Pointer() != nil {
		var selectorQueryC *C.char
		if selectorQuery != "" {
			selectorQueryC = C.CString(selectorQuery)
			defer C.free(unsafe.Pointer(selectorQueryC))
		}
		var tmpValue = NewQWebElementFromPointer(C.QWebElement_FindFirst(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: selectorQueryC, len: C.longlong(len(selectorQuery))}))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) FirstChild() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElement_FirstChild(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) LastChild() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElement_LastChild(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) NextSibling() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElement_NextSibling(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) Parent() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElement_Parent(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) PreviousSibling() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElement_PreviousSibling(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) FindAll(selectorQuery string) *QWebElementCollection {
	if ptr.Pointer() != nil {
		var selectorQueryC *C.char
		if selectorQuery != "" {
			selectorQueryC = C.CString(selectorQuery)
			defer C.free(unsafe.Pointer(selectorQueryC))
		}
		var tmpValue = NewQWebElementCollectionFromPointer(C.QWebElement_FindAll(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: selectorQueryC, len: C.longlong(len(selectorQuery))}))
		runtime.SetFinalizer(tmpValue, (*QWebElementCollection).DestroyQWebElementCollection)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) WebFrame() *QWebFrame {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebFrameFromPointer(C.QWebElement_WebFrame(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebElement) HasAttribute(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return C.QWebElement_HasAttribute(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))}) != 0
	}
	return false
}

func (ptr *QWebElement) HasAttributeNS(namespaceUri string, name string) bool {
	if ptr.Pointer() != nil {
		var namespaceUriC *C.char
		if namespaceUri != "" {
			namespaceUriC = C.CString(namespaceUri)
			defer C.free(unsafe.Pointer(namespaceUriC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return C.QWebElement_HasAttributeNS(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: namespaceUriC, len: C.longlong(len(namespaceUri))}, C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))}) != 0
	}
	return false
}

func (ptr *QWebElement) HasAttributes() bool {
	if ptr.Pointer() != nil {
		return C.QWebElement_HasAttributes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebElement) HasClass(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return C.QWebElement_HasClass(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))}) != 0
	}
	return false
}

func (ptr *QWebElement) HasFocus() bool {
	if ptr.Pointer() != nil {
		return C.QWebElement_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebElement) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QWebElement_IsNull(ptr.Pointer()) != 0
	}
	return false
}

type QWebElementCollection struct {
	ptr unsafe.Pointer
}

type QWebElementCollection_ITF interface {
	QWebElementCollection_PTR() *QWebElementCollection
}

func (ptr *QWebElementCollection) QWebElementCollection_PTR() *QWebElementCollection {
	return ptr
}

func (ptr *QWebElementCollection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebElementCollection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebElementCollection(ptr QWebElementCollection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebElementCollection_PTR().Pointer()
	}
	return nil
}

func NewQWebElementCollectionFromPointer(ptr unsafe.Pointer) *QWebElementCollection {
	var n = new(QWebElementCollection)
	n.SetPointer(ptr)
	return n
}
func NewQWebElementCollection() *QWebElementCollection {
	var tmpValue = NewQWebElementCollectionFromPointer(C.QWebElementCollection_NewQWebElementCollection())
	runtime.SetFinalizer(tmpValue, (*QWebElementCollection).DestroyQWebElementCollection)
	return tmpValue
}

func NewQWebElementCollection2(contextElement QWebElement_ITF, query string) *QWebElementCollection {
	var queryC *C.char
	if query != "" {
		queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
	}
	var tmpValue = NewQWebElementCollectionFromPointer(C.QWebElementCollection_NewQWebElementCollection2(PointerFromQWebElement(contextElement), C.struct_QtWebKit_PackedString{data: queryC, len: C.longlong(len(query))}))
	runtime.SetFinalizer(tmpValue, (*QWebElementCollection).DestroyQWebElementCollection)
	return tmpValue
}

func NewQWebElementCollection3(other QWebElementCollection_ITF) *QWebElementCollection {
	var tmpValue = NewQWebElementCollectionFromPointer(C.QWebElementCollection_NewQWebElementCollection3(PointerFromQWebElementCollection(other)))
	runtime.SetFinalizer(tmpValue, (*QWebElementCollection).DestroyQWebElementCollection)
	return tmpValue
}

func (ptr *QWebElementCollection) Append(other QWebElementCollection_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElementCollection_Append(ptr.Pointer(), PointerFromQWebElementCollection(other))
	}
}

func (ptr *QWebElementCollection) DestroyQWebElementCollection() {
	if ptr.Pointer() != nil {
		C.QWebElementCollection_DestroyQWebElementCollection(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebElementCollection) ToList() []*QWebElement {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []*QWebElement {
			var out = make([]*QWebElement, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebElementCollectionFromPointer(l.data).__toList_atList(i)
			}
			return out
		}(C.QWebElementCollection_ToList(ptr.Pointer()))
	}
	return make([]*QWebElement, 0)
}

func (ptr *QWebElementCollection) At(i int) *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElementCollection_At(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElementCollection) First() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElementCollection_First(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElementCollection) Last() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElementCollection_Last(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElementCollection) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebElementCollection_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebElementCollection) __toList_atList(i int) *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebElementCollection___toList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebElementCollection) __toList_setList(i QWebElement_ITF) {
	if ptr.Pointer() != nil {
		C.QWebElementCollection___toList_setList(ptr.Pointer(), PointerFromQWebElement(i))
	}
}

func (ptr *QWebElementCollection) __toList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebElementCollection___toList_newList(ptr.Pointer()))
}

type QWebFrame struct {
	core.QObject
}

type QWebFrame_ITF interface {
	core.QObject_ITF
	QWebFrame_PTR() *QWebFrame
}

func (ptr *QWebFrame) QWebFrame_PTR() *QWebFrame {
	return ptr
}

func (ptr *QWebFrame) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebFrame) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebFrame(ptr QWebFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebFrame_PTR().Pointer()
	}
	return nil
}

func NewQWebFrameFromPointer(ptr unsafe.Pointer) *QWebFrame {
	var n = new(QWebFrame)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QWebFrame__RenderLayer
//QWebFrame::RenderLayer
type QWebFrame__RenderLayer int64

const (
	QWebFrame__ContentsLayer  QWebFrame__RenderLayer = QWebFrame__RenderLayer(0x10)
	QWebFrame__ScrollBarLayer QWebFrame__RenderLayer = QWebFrame__RenderLayer(0x20)
	QWebFrame__PanIconLayer   QWebFrame__RenderLayer = QWebFrame__RenderLayer(0x40)
	QWebFrame__AllLayers      QWebFrame__RenderLayer = QWebFrame__RenderLayer(0xff)
)

//go:generate stringer -type=QWebFrame__ValueOwnership
//QWebFrame::ValueOwnership
type QWebFrame__ValueOwnership int64

const (
	QWebFrame__QtOwnership     QWebFrame__ValueOwnership = QWebFrame__ValueOwnership(0)
	QWebFrame__ScriptOwnership QWebFrame__ValueOwnership = QWebFrame__ValueOwnership(1)
	QWebFrame__AutoOwnership   QWebFrame__ValueOwnership = QWebFrame__ValueOwnership(2)
)

//export callbackQWebFrame_Event
func callbackQWebFrame_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebFrameFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebFrame) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebFrame_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebFrame) AddToJavaScriptWindowObject(name string, object core.QObject_ITF, own QWebFrame__ValueOwnership) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QWebFrame_AddToJavaScriptWindowObject(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQObject(object), C.longlong(own))
	}
}

//export callbackQWebFrame_EvaluateJavaScript
func callbackQWebFrame_EvaluateJavaScript(ptr unsafe.Pointer, scriptSource C.struct_QtWebKit_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "evaluateJavaScript"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(cGoUnpackString(scriptSource)))
	}

	return core.PointerFromQVariant(NewQWebFrameFromPointer(ptr).EvaluateJavaScriptDefault(cGoUnpackString(scriptSource)))
}

func (ptr *QWebFrame) ConnectEvaluateJavaScript(f func(scriptSource string) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "evaluateJavaScript"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "evaluateJavaScript", func(scriptSource string) *core.QVariant {
				signal.(func(string) *core.QVariant)(scriptSource)
				return f(scriptSource)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "evaluateJavaScript", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectEvaluateJavaScript() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "evaluateJavaScript")
	}
}

func (ptr *QWebFrame) EvaluateJavaScript(scriptSource string) *core.QVariant {
	if ptr.Pointer() != nil {
		var scriptSourceC *C.char
		if scriptSource != "" {
			scriptSourceC = C.CString(scriptSource)
			defer C.free(unsafe.Pointer(scriptSourceC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QWebFrame_EvaluateJavaScript(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: scriptSourceC, len: C.longlong(len(scriptSource))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) EvaluateJavaScriptDefault(scriptSource string) *core.QVariant {
	if ptr.Pointer() != nil {
		var scriptSourceC *C.char
		if scriptSource != "" {
			scriptSourceC = C.CString(scriptSource)
			defer C.free(unsafe.Pointer(scriptSourceC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QWebFrame_EvaluateJavaScriptDefault(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: scriptSourceC, len: C.longlong(len(scriptSource))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQWebFrame_ContentsSizeChanged
func callbackQWebFrame_ContentsSizeChanged(ptr unsafe.Pointer, size unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contentsSizeChanged"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(size))
	}

}

func (ptr *QWebFrame) ConnectContentsSizeChanged(f func(size *core.QSize)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "contentsSizeChanged") {
			C.QWebFrame_ConnectContentsSizeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "contentsSizeChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contentsSizeChanged", func(size *core.QSize) {
				signal.(func(*core.QSize))(size)
				f(size)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentsSizeChanged", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectContentsSizeChanged() {
	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectContentsSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "contentsSizeChanged")
	}
}

func (ptr *QWebFrame) ContentsSizeChanged(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_ContentsSizeChanged(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

//export callbackQWebFrame_IconChanged
func callbackQWebFrame_IconChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebFrame) ConnectIconChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "iconChanged") {
			C.QWebFrame_ConnectIconChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "iconChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "iconChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "iconChanged", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectIconChanged() {
	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectIconChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "iconChanged")
	}
}

func (ptr *QWebFrame) IconChanged() {
	if ptr.Pointer() != nil {
		C.QWebFrame_IconChanged(ptr.Pointer())
	}
}

//export callbackQWebFrame_InitialLayoutCompleted
func callbackQWebFrame_InitialLayoutCompleted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initialLayoutCompleted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebFrame) ConnectInitialLayoutCompleted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "initialLayoutCompleted") {
			C.QWebFrame_ConnectInitialLayoutCompleted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "initialLayoutCompleted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "initialLayoutCompleted", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initialLayoutCompleted", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectInitialLayoutCompleted() {
	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectInitialLayoutCompleted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "initialLayoutCompleted")
	}
}

func (ptr *QWebFrame) InitialLayoutCompleted() {
	if ptr.Pointer() != nil {
		C.QWebFrame_InitialLayoutCompleted(ptr.Pointer())
	}
}

//export callbackQWebFrame_JavaScriptWindowObjectCleared
func callbackQWebFrame_JavaScriptWindowObjectCleared(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "javaScriptWindowObjectCleared"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebFrame) ConnectJavaScriptWindowObjectCleared(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "javaScriptWindowObjectCleared") {
			C.QWebFrame_ConnectJavaScriptWindowObjectCleared(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "javaScriptWindowObjectCleared"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptWindowObjectCleared", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptWindowObjectCleared", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectJavaScriptWindowObjectCleared() {
	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectJavaScriptWindowObjectCleared(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "javaScriptWindowObjectCleared")
	}
}

func (ptr *QWebFrame) JavaScriptWindowObjectCleared() {
	if ptr.Pointer() != nil {
		C.QWebFrame_JavaScriptWindowObjectCleared(ptr.Pointer())
	}
}

func (ptr *QWebFrame) Load2(req network.QNetworkRequest_ITF, operation network.QNetworkAccessManager__Operation, body core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_Load2(ptr.Pointer(), network.PointerFromQNetworkRequest(req), C.longlong(operation), core.PointerFromQByteArray(body))
	}
}

func (ptr *QWebFrame) Load(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebFrame_LoadFinished
func callbackQWebFrame_LoadFinished(ptr unsafe.Pointer, ok C.char) {
	if signal := qt.GetSignal(ptr, "loadFinished"); signal != nil {
		signal.(func(bool))(int8(ok) != 0)
	}

}

func (ptr *QWebFrame) ConnectLoadFinished(f func(ok bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadFinished") {
			C.QWebFrame_ConnectLoadFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", func(ok bool) {
				signal.(func(bool))(ok)
				f(ok)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectLoadFinished() {
	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadFinished")
	}
}

func (ptr *QWebFrame) LoadFinished(ok bool) {
	if ptr.Pointer() != nil {
		C.QWebFrame_LoadFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))
	}
}

//export callbackQWebFrame_LoadStarted
func callbackQWebFrame_LoadStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebFrame) ConnectLoadStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadStarted") {
			C.QWebFrame_ConnectLoadStarted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadStarted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectLoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadStarted")
	}
}

func (ptr *QWebFrame) LoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebFrame_LoadStarted(ptr.Pointer())
	}
}

//export callbackQWebFrame_PageChanged
func callbackQWebFrame_PageChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pageChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebFrame) ConnectPageChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pageChanged") {
			C.QWebFrame_ConnectPageChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pageChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pageChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pageChanged", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectPageChanged() {
	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectPageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pageChanged")
	}
}

func (ptr *QWebFrame) PageChanged() {
	if ptr.Pointer() != nil {
		C.QWebFrame_PageChanged(ptr.Pointer())
	}
}

func (ptr *QWebFrame) Render2(painter gui.QPainter_ITF, layer QWebFrame__RenderLayer, clip gui.QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_Render2(ptr.Pointer(), gui.PointerFromQPainter(painter), C.longlong(layer), gui.PointerFromQRegion(clip))
	}
}

func (ptr *QWebFrame) Render(painter gui.QPainter_ITF, clip gui.QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_Render(ptr.Pointer(), gui.PointerFromQPainter(painter), gui.PointerFromQRegion(clip))
	}
}

func (ptr *QWebFrame) Scroll(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QWebFrame_Scroll(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

func (ptr *QWebFrame) ScrollToAnchor(anchor string) {
	if ptr.Pointer() != nil {
		var anchorC *C.char
		if anchor != "" {
			anchorC = C.CString(anchor)
			defer C.free(unsafe.Pointer(anchorC))
		}
		C.QWebFrame_ScrollToAnchor(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: anchorC, len: C.longlong(len(anchor))})
	}
}

func (ptr *QWebFrame) SetContent(data core.QByteArray_ITF, mimeType string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var mimeTypeC *C.char
		if mimeType != "" {
			mimeTypeC = C.CString(mimeType)
			defer C.free(unsafe.Pointer(mimeTypeC))
		}
		C.QWebFrame_SetContent(ptr.Pointer(), core.PointerFromQByteArray(data), C.struct_QtWebKit_PackedString{data: mimeTypeC, len: C.longlong(len(mimeType))}, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebFrame) SetFocus() {
	if ptr.Pointer() != nil {
		C.QWebFrame_SetFocus(ptr.Pointer())
	}
}

func (ptr *QWebFrame) SetHtml(html string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var htmlC *C.char
		if html != "" {
			htmlC = C.CString(html)
			defer C.free(unsafe.Pointer(htmlC))
		}
		C.QWebFrame_SetHtml(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: htmlC, len: C.longlong(len(html))}, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebFrame) SetScrollBarPolicy(orientation core.Qt__Orientation, policy core.Qt__ScrollBarPolicy) {
	if ptr.Pointer() != nil {
		C.QWebFrame_SetScrollBarPolicy(ptr.Pointer(), C.longlong(orientation), C.longlong(policy))
	}
}

func (ptr *QWebFrame) SetScrollBarValue(orientation core.Qt__Orientation, value int) {
	if ptr.Pointer() != nil {
		C.QWebFrame_SetScrollBarValue(ptr.Pointer(), C.longlong(orientation), C.int(int32(value)))
	}
}

func (ptr *QWebFrame) SetScrollPosition(pos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_SetScrollPosition(ptr.Pointer(), core.PointerFromQPoint(pos))
	}
}

func (ptr *QWebFrame) SetTextSizeMultiplier(factor float64) {
	if ptr.Pointer() != nil {
		C.QWebFrame_SetTextSizeMultiplier(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebFrame) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebFrame) SetZoomFactor(factor float64) {
	if ptr.Pointer() != nil {
		C.QWebFrame_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

//export callbackQWebFrame_TitleChanged
func callbackQWebFrame_TitleChanged(ptr unsafe.Pointer, title C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "titleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

func (ptr *QWebFrame) ConnectTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleChanged") {
			C.QWebFrame_ConnectTitleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", func(title string) {
				signal.(func(string))(title)
				f(title)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleChanged")
	}
}

func (ptr *QWebFrame) TitleChanged(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QWebFrame_TitleChanged(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

//export callbackQWebFrame_UrlChanged
func callbackQWebFrame_UrlChanged(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "urlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebFrame) ConnectUrlChanged(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "urlChanged") {
			C.QWebFrame_ConnectUrlChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "urlChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "urlChanged", func(url *core.QUrl) {
				signal.(func(*core.QUrl))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "urlChanged", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectUrlChanged() {
	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "urlChanged")
	}
}

func (ptr *QWebFrame) UrlChanged(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebFrame) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQIconFromPointer(C.QWebFrame_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) ChildFrames() []*QWebFrame {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []*QWebFrame {
			var out = make([]*QWebFrame, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebFrameFromPointer(l.data).__childFrames_atList(i)
			}
			return out
		}(C.QWebFrame_ChildFrames(ptr.Pointer()))
	}
	return make([]*QWebFrame, 0)
}

func (ptr *QWebFrame) MetaData() map[string]string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) map[string]string {
			var out = make(map[string]string, int(l.len))
			for _, i := range NewQWebFrameFromPointer(l.data).__metaData_keyList() {
				out[i] = NewQWebFrameFromPointer(l.data).__metaData_atList(i)
			}
			return out
		}(C.QWebFrame_MetaData(ptr.Pointer()))
	}
	return make(map[string]string, 0)
}

func (ptr *QWebFrame) Pos() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.QWebFrame_Pos(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) ScrollPosition() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.QWebFrame_ScrollPosition(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) Geometry() *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QWebFrame_Geometry(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) ScrollBarGeometry(orientation core.Qt__Orientation) *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QWebFrame_ScrollBarGeometry(ptr.Pointer(), C.longlong(orientation)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) ContentsSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebFrame_ContentsSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) FrameName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebFrame_FrameName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebFrame) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebFrame_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebFrame) ToHtml() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebFrame_ToHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebFrame) ToPlainText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebFrame_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebFrame) BaseUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebFrame_BaseUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) RequestedUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebFrame_RequestedUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebFrame_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) DocumentElement() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebFrame_DocumentElement(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) FindFirstElement(selectorQuery string) *QWebElement {
	if ptr.Pointer() != nil {
		var selectorQueryC *C.char
		if selectorQuery != "" {
			selectorQueryC = C.CString(selectorQuery)
			defer C.free(unsafe.Pointer(selectorQueryC))
		}
		var tmpValue = NewQWebElementFromPointer(C.QWebFrame_FindFirstElement(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: selectorQueryC, len: C.longlong(len(selectorQuery))}))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) OwnerElement() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebFrame_OwnerElement(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) FindAllElements(selectorQuery string) *QWebElementCollection {
	if ptr.Pointer() != nil {
		var selectorQueryC *C.char
		if selectorQuery != "" {
			selectorQueryC = C.CString(selectorQuery)
			defer C.free(unsafe.Pointer(selectorQueryC))
		}
		var tmpValue = NewQWebElementCollectionFromPointer(C.QWebFrame_FindAllElements(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: selectorQueryC, len: C.longlong(len(selectorQuery))}))
		runtime.SetFinalizer(tmpValue, (*QWebElementCollection).DestroyQWebElementCollection)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) ParentFrame() *QWebFrame {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebFrameFromPointer(C.QWebFrame_ParentFrame(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) HitTestContent(pos core.QPoint_ITF) *QWebHitTestResult {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebHitTestResultFromPointer(C.QWebFrame_HitTestContent(ptr.Pointer(), core.PointerFromQPoint(pos)))
		runtime.SetFinalizer(tmpValue, (*QWebHitTestResult).DestroyQWebHitTestResult)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) Page() *QWebPage {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebPageFromPointer(C.QWebFrame_Page(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) SecurityOrigin() *QWebSecurityOrigin {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebSecurityOriginFromPointer(C.QWebFrame_SecurityOrigin(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebSecurityOrigin).DestroyQWebSecurityOrigin)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) ScrollBarPolicy(orientation core.Qt__Orientation) core.Qt__ScrollBarPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QWebFrame_ScrollBarPolicy(ptr.Pointer(), C.longlong(orientation)))
	}
	return 0
}

func (ptr *QWebFrame) HasFocus() bool {
	if ptr.Pointer() != nil {
		return C.QWebFrame_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebFrame) ScrollBarMaximum(orientation core.Qt__Orientation) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebFrame_ScrollBarMaximum(ptr.Pointer(), C.longlong(orientation))))
	}
	return 0
}

func (ptr *QWebFrame) ScrollBarMinimum(orientation core.Qt__Orientation) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebFrame_ScrollBarMinimum(ptr.Pointer(), C.longlong(orientation))))
	}
	return 0
}

func (ptr *QWebFrame) ScrollBarValue(orientation core.Qt__Orientation) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebFrame_ScrollBarValue(ptr.Pointer(), C.longlong(orientation))))
	}
	return 0
}

func (ptr *QWebFrame) TextSizeMultiplier() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWebFrame_TextSizeMultiplier(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebFrame) ZoomFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWebFrame_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebFrame_Print
func callbackQWebFrame_Print(ptr unsafe.Pointer, printer unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "print"); signal != nil {
		signal.(func(*printsupport.QPrinter))(printsupport.NewQPrinterFromPointer(printer))
	} else {
		NewQWebFrameFromPointer(ptr).PrintDefault(printsupport.NewQPrinterFromPointer(printer))
	}
}

func (ptr *QWebFrame) ConnectPrint(f func(printer *printsupport.QPrinter)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "print"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "print", func(printer *printsupport.QPrinter) {
				signal.(func(*printsupport.QPrinter))(printer)
				f(printer)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "print", f)
		}
	}
}

func (ptr *QWebFrame) DisconnectPrint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "print")
	}
}

func (ptr *QWebFrame) Print(printer printsupport.QPrinter_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_Print(ptr.Pointer(), printsupport.PointerFromQPrinter(printer))
	}
}

func (ptr *QWebFrame) PrintDefault(printer printsupport.QPrinter_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_PrintDefault(ptr.Pointer(), printsupport.PointerFromQPrinter(printer))
	}
}

func (ptr *QWebFrame) __childFrames_atList(i int) *QWebFrame {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebFrameFromPointer(C.QWebFrame___childFrames_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) __childFrames_setList(i QWebFrame_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame___childFrames_setList(ptr.Pointer(), PointerFromQWebFrame(i))
	}
}

func (ptr *QWebFrame) __childFrames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebFrame___childFrames_newList(ptr.Pointer()))
}

func (ptr *QWebFrame) __metaData_atList(i string) string {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		return cGoUnpackString(C.QWebFrame___metaData_atList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))}))
	}
	return ""
}

func (ptr *QWebFrame) __metaData_setList(key string, i string) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QWebFrame___metaData_setList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: keyC, len: C.longlong(len(key))}, C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QWebFrame) __metaData_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebFrame___metaData_newList(ptr.Pointer()))
}

func (ptr *QWebFrame) __metaData_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebFrameFromPointer(l.data).____metaData_keyList_atList(i)
			}
			return out
		}(C.QWebFrame___metaData_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QWebFrame) ____metaData_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebFrame_____metaData_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QWebFrame) ____metaData_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QWebFrame_____metaData_keyList_setList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QWebFrame) ____metaData_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebFrame_____metaData_keyList_newList(ptr.Pointer()))
}

func (ptr *QWebFrame) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QWebFrame___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebFrame) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebFrame___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QWebFrame) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebFrame___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebFrame) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QWebFrame___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QWebFrame) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebFrame___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebFrame) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QWebFrame___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QWebFrame) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebFrame___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebFrame) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebFrame___findChildren_newList(ptr.Pointer()))
}

func (ptr *QWebFrame) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebFrame___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebFrame) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebFrame) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebFrame___children_newList(ptr.Pointer()))
}

//export callbackQWebFrame_EventFilter
func callbackQWebFrame_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebFrameFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebFrame) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebFrame_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebFrame_ChildEvent
func callbackQWebFrame_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebFrameFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebFrame) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebFrame_ConnectNotify
func callbackQWebFrame_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebFrameFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebFrame) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebFrame_CustomEvent
func callbackQWebFrame_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebFrameFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebFrame) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebFrame_DeleteLater
func callbackQWebFrame_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebFrameFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebFrame) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebFrame_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebFrame_Destroyed
func callbackQWebFrame_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebFrame_DisconnectNotify
func callbackQWebFrame_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebFrameFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebFrame) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebFrame_ObjectNameChanged
func callbackQWebFrame_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQWebFrame_TimerEvent
func callbackQWebFrame_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebFrameFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebFrame) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebFrame_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebFrame_MetaObject
func callbackQWebFrame_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebFrameFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebFrame) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebFrame_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebHistory struct {
	ptr unsafe.Pointer
}

type QWebHistory_ITF interface {
	QWebHistory_PTR() *QWebHistory
}

func (ptr *QWebHistory) QWebHistory_PTR() *QWebHistory {
	return ptr
}

func (ptr *QWebHistory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebHistory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebHistory(ptr QWebHistory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebHistory_PTR().Pointer()
	}
	return nil
}

func NewQWebHistoryFromPointer(ptr unsafe.Pointer) *QWebHistory {
	var n = new(QWebHistory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebHistory) DestroyQWebHistory() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebHistory) Back() {
	if ptr.Pointer() != nil {
		C.QWebHistory_Back(ptr.Pointer())
	}
}

func (ptr *QWebHistory) Clear() {
	if ptr.Pointer() != nil {
		C.QWebHistory_Clear(ptr.Pointer())
	}
}

func (ptr *QWebHistory) Forward() {
	if ptr.Pointer() != nil {
		C.QWebHistory_Forward(ptr.Pointer())
	}
}

func (ptr *QWebHistory) GoToItem(item QWebHistoryItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistory_GoToItem(ptr.Pointer(), PointerFromQWebHistoryItem(item))
	}
}

func (ptr *QWebHistory) SetMaximumItemCount(count int) {
	if ptr.Pointer() != nil {
		C.QWebHistory_SetMaximumItemCount(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QWebHistory) BackItems(maxItems int) []*QWebHistoryItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []*QWebHistoryItem {
			var out = make([]*QWebHistoryItem, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebHistoryFromPointer(l.data).__backItems_atList(i)
			}
			return out
		}(C.QWebHistory_BackItems(ptr.Pointer(), C.int(int32(maxItems))))
	}
	return make([]*QWebHistoryItem, 0)
}

func (ptr *QWebHistory) ForwardItems(maxItems int) []*QWebHistoryItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []*QWebHistoryItem {
			var out = make([]*QWebHistoryItem, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebHistoryFromPointer(l.data).__forwardItems_atList(i)
			}
			return out
		}(C.QWebHistory_ForwardItems(ptr.Pointer(), C.int(int32(maxItems))))
	}
	return make([]*QWebHistoryItem, 0)
}

func (ptr *QWebHistory) Items() []*QWebHistoryItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []*QWebHistoryItem {
			var out = make([]*QWebHistoryItem, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebHistoryFromPointer(l.data).__items_atList(i)
			}
			return out
		}(C.QWebHistory_Items(ptr.Pointer()))
	}
	return make([]*QWebHistoryItem, 0)
}

func (ptr *QWebHistory) ToMap() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) map[string]*core.QVariant {
			var out = make(map[string]*core.QVariant, int(l.len))
			for _, i := range NewQWebHistoryFromPointer(l.data).__toMap_keyList() {
				out[i] = NewQWebHistoryFromPointer(l.data).__toMap_atList(i)
			}
			return out
		}(C.QWebHistory_ToMap(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *QWebHistory) BackItem() *QWebHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebHistoryItemFromPointer(C.QWebHistory_BackItem(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebHistoryItem).DestroyQWebHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistory) CurrentItem() *QWebHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebHistoryItemFromPointer(C.QWebHistory_CurrentItem(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebHistoryItem).DestroyQWebHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistory) ForwardItem() *QWebHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebHistoryItemFromPointer(C.QWebHistory_ForwardItem(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebHistoryItem).DestroyQWebHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistory) ItemAt(i int) *QWebHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebHistoryItemFromPointer(C.QWebHistory_ItemAt(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebHistoryItem).DestroyQWebHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistory) CanGoBack() bool {
	if ptr.Pointer() != nil {
		return C.QWebHistory_CanGoBack(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHistory) CanGoForward() bool {
	if ptr.Pointer() != nil {
		return C.QWebHistory_CanGoForward(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHistory) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebHistory_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebHistory) CurrentItemIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebHistory_CurrentItemIndex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebHistory) MaximumItemCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebHistory_MaximumItemCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebHistory) __loadFromMap_map_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QWebHistory___loadFromMap_map_atList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistory) __loadFromMap_map_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QWebHistory___loadFromMap_map_setList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QWebHistory) __loadFromMap_map_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistory___loadFromMap_map_newList(ptr.Pointer()))
}

func (ptr *QWebHistory) __loadFromMap_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebHistoryFromPointer(l.data).____loadFromMap_keyList_atList(i)
			}
			return out
		}(C.QWebHistory___loadFromMap_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QWebHistory) __backItems_atList(i int) *QWebHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebHistoryItemFromPointer(C.QWebHistory___backItems_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebHistoryItem).DestroyQWebHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistory) __backItems_setList(i QWebHistoryItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistory___backItems_setList(ptr.Pointer(), PointerFromQWebHistoryItem(i))
	}
}

func (ptr *QWebHistory) __backItems_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistory___backItems_newList(ptr.Pointer()))
}

func (ptr *QWebHistory) __forwardItems_atList(i int) *QWebHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebHistoryItemFromPointer(C.QWebHistory___forwardItems_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebHistoryItem).DestroyQWebHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistory) __forwardItems_setList(i QWebHistoryItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistory___forwardItems_setList(ptr.Pointer(), PointerFromQWebHistoryItem(i))
	}
}

func (ptr *QWebHistory) __forwardItems_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistory___forwardItems_newList(ptr.Pointer()))
}

func (ptr *QWebHistory) __items_atList(i int) *QWebHistoryItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebHistoryItemFromPointer(C.QWebHistory___items_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebHistoryItem).DestroyQWebHistoryItem)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistory) __items_setList(i QWebHistoryItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistory___items_setList(ptr.Pointer(), PointerFromQWebHistoryItem(i))
	}
}

func (ptr *QWebHistory) __items_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistory___items_newList(ptr.Pointer()))
}

func (ptr *QWebHistory) __toMap_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QWebHistory___toMap_atList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistory) __toMap_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QWebHistory___toMap_setList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QWebHistory) __toMap_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistory___toMap_newList(ptr.Pointer()))
}

func (ptr *QWebHistory) __toMap_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebHistoryFromPointer(l.data).____toMap_keyList_atList(i)
			}
			return out
		}(C.QWebHistory___toMap_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QWebHistory) ____loadFromMap_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebHistory_____loadFromMap_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QWebHistory) ____loadFromMap_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QWebHistory_____loadFromMap_keyList_setList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QWebHistory) ____loadFromMap_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistory_____loadFromMap_keyList_newList(ptr.Pointer()))
}

func (ptr *QWebHistory) ____toMap_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebHistory_____toMap_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QWebHistory) ____toMap_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QWebHistory_____toMap_keyList_setList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QWebHistory) ____toMap_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistory_____toMap_keyList_newList(ptr.Pointer()))
}

type QWebHistoryInterface struct {
	core.QObject
}

type QWebHistoryInterface_ITF interface {
	core.QObject_ITF
	QWebHistoryInterface_PTR() *QWebHistoryInterface
}

func (ptr *QWebHistoryInterface) QWebHistoryInterface_PTR() *QWebHistoryInterface {
	return ptr
}

func (ptr *QWebHistoryInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebHistoryInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebHistoryInterface(ptr QWebHistoryInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebHistoryInterface_PTR().Pointer()
	}
	return nil
}

func NewQWebHistoryInterfaceFromPointer(ptr unsafe.Pointer) *QWebHistoryInterface {
	var n = new(QWebHistoryInterface)
	n.SetPointer(ptr)
	return n
}
func QWebHistoryInterface_DefaultInterface() *QWebHistoryInterface {
	var tmpValue = NewQWebHistoryInterfaceFromPointer(C.QWebHistoryInterface_QWebHistoryInterface_DefaultInterface())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWebHistoryInterface) DefaultInterface() *QWebHistoryInterface {
	var tmpValue = NewQWebHistoryInterfaceFromPointer(C.QWebHistoryInterface_QWebHistoryInterface_DefaultInterface())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQWebHistoryInterface(parent core.QObject_ITF) *QWebHistoryInterface {
	var tmpValue = NewQWebHistoryInterfaceFromPointer(C.QWebHistoryInterface_NewQWebHistoryInterface(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebHistoryInterface_AddHistoryEntry
func callbackQWebHistoryInterface_AddHistoryEntry(ptr unsafe.Pointer, url C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "addHistoryEntry"); signal != nil {
		signal.(func(string))(cGoUnpackString(url))
	}

}

func (ptr *QWebHistoryInterface) ConnectAddHistoryEntry(f func(url string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addHistoryEntry"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addHistoryEntry", func(url string) {
				signal.(func(string))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addHistoryEntry", f)
		}
	}
}

func (ptr *QWebHistoryInterface) DisconnectAddHistoryEntry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addHistoryEntry")
	}
}

func (ptr *QWebHistoryInterface) AddHistoryEntry(url string) {
	if ptr.Pointer() != nil {
		var urlC *C.char
		if url != "" {
			urlC = C.CString(url)
			defer C.free(unsafe.Pointer(urlC))
		}
		C.QWebHistoryInterface_AddHistoryEntry(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: urlC, len: C.longlong(len(url))})
	}
}

func QWebHistoryInterface_SetDefaultInterface(defaultInterface QWebHistoryInterface_ITF) {
	C.QWebHistoryInterface_QWebHistoryInterface_SetDefaultInterface(PointerFromQWebHistoryInterface(defaultInterface))
}

func (ptr *QWebHistoryInterface) SetDefaultInterface(defaultInterface QWebHistoryInterface_ITF) {
	C.QWebHistoryInterface_QWebHistoryInterface_SetDefaultInterface(PointerFromQWebHistoryInterface(defaultInterface))
}

func (ptr *QWebHistoryInterface) DestroyQWebHistoryInterface() {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_DestroyQWebHistoryInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebHistoryInterface_HistoryContains
func callbackQWebHistoryInterface_HistoryContains(ptr unsafe.Pointer, url C.struct_QtWebKit_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "historyContains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QWebHistoryInterface) ConnectHistoryContains(f func(url string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "historyContains"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "historyContains", func(url string) bool {
				signal.(func(string) bool)(url)
				return f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "historyContains", f)
		}
	}
}

func (ptr *QWebHistoryInterface) DisconnectHistoryContains() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "historyContains")
	}
}

func (ptr *QWebHistoryInterface) HistoryContains(url string) bool {
	if ptr.Pointer() != nil {
		var urlC *C.char
		if url != "" {
			urlC = C.CString(url)
			defer C.free(unsafe.Pointer(urlC))
		}
		return C.QWebHistoryInterface_HistoryContains(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: urlC, len: C.longlong(len(url))}) != 0
	}
	return false
}

func (ptr *QWebHistoryInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QWebHistoryInterface___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebHistoryInterface) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistoryInterface___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QWebHistoryInterface) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebHistoryInterface___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryInterface) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebHistoryInterface) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistoryInterface___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QWebHistoryInterface) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebHistoryInterface___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryInterface) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebHistoryInterface) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistoryInterface___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QWebHistoryInterface) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebHistoryInterface___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryInterface) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebHistoryInterface) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistoryInterface___findChildren_newList(ptr.Pointer()))
}

func (ptr *QWebHistoryInterface) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebHistoryInterface___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryInterface) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebHistoryInterface) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistoryInterface___children_newList(ptr.Pointer()))
}

//export callbackQWebHistoryInterface_Event
func callbackQWebHistoryInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebHistoryInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebHistoryInterface) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebHistoryInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebHistoryInterface_EventFilter
func callbackQWebHistoryInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebHistoryInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebHistoryInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebHistoryInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebHistoryInterface_ChildEvent
func callbackQWebHistoryInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebHistoryInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebHistoryInterface_ConnectNotify
func callbackQWebHistoryInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebHistoryInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebHistoryInterface_CustomEvent
func callbackQWebHistoryInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebHistoryInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebHistoryInterface_DeleteLater
func callbackQWebHistoryInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebHistoryInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebHistoryInterface_Destroyed
func callbackQWebHistoryInterface_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebHistoryInterface_DisconnectNotify
func callbackQWebHistoryInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebHistoryInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebHistoryInterface_ObjectNameChanged
func callbackQWebHistoryInterface_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQWebHistoryInterface_TimerEvent
func callbackQWebHistoryInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebHistoryInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebHistoryInterface_MetaObject
func callbackQWebHistoryInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebHistoryInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebHistoryInterface) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebHistoryInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebHistoryItem struct {
	ptr unsafe.Pointer
}

type QWebHistoryItem_ITF interface {
	QWebHistoryItem_PTR() *QWebHistoryItem
}

func (ptr *QWebHistoryItem) QWebHistoryItem_PTR() *QWebHistoryItem {
	return ptr
}

func (ptr *QWebHistoryItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebHistoryItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebHistoryItem(ptr QWebHistoryItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebHistoryItem_PTR().Pointer()
	}
	return nil
}

func NewQWebHistoryItemFromPointer(ptr unsafe.Pointer) *QWebHistoryItem {
	var n = new(QWebHistoryItem)
	n.SetPointer(ptr)
	return n
}
func NewQWebHistoryItem(other QWebHistoryItem_ITF) *QWebHistoryItem {
	var tmpValue = NewQWebHistoryItemFromPointer(C.QWebHistoryItem_NewQWebHistoryItem(PointerFromQWebHistoryItem(other)))
	runtime.SetFinalizer(tmpValue, (*QWebHistoryItem).DestroyQWebHistoryItem)
	return tmpValue
}

func (ptr *QWebHistoryItem) SetUserData(userData core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QWebHistoryItem_SetUserData(ptr.Pointer(), core.PointerFromQVariant(userData))
	}
}

func (ptr *QWebHistoryItem) DestroyQWebHistoryItem() {
	if ptr.Pointer() != nil {
		C.QWebHistoryItem_DestroyQWebHistoryItem(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebHistoryItem) LastVisited() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QWebHistoryItem_LastVisited(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryItem) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQIconFromPointer(C.QWebHistoryItem_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryItem) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebHistoryItem_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebHistoryItem) OriginalUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebHistoryItem_OriginalUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryItem) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebHistoryItem_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryItem) UserData() *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QWebHistoryItem_UserData(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryItem) ToMap() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) map[string]*core.QVariant {
			var out = make(map[string]*core.QVariant, int(l.len))
			for _, i := range NewQWebHistoryItemFromPointer(l.data).__toMap_keyList() {
				out[i] = NewQWebHistoryItemFromPointer(l.data).__toMap_atList(i)
			}
			return out
		}(C.QWebHistoryItem_ToMap(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *QWebHistoryItem) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QWebHistoryItem_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHistoryItem) __loadFromMap_map_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QWebHistoryItem___loadFromMap_map_atList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryItem) __loadFromMap_map_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QWebHistoryItem___loadFromMap_map_setList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QWebHistoryItem) __loadFromMap_map_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistoryItem___loadFromMap_map_newList(ptr.Pointer()))
}

func (ptr *QWebHistoryItem) __loadFromMap_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebHistoryItemFromPointer(l.data).____loadFromMap_keyList_atList(i)
			}
			return out
		}(C.QWebHistoryItem___loadFromMap_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QWebHistoryItem) __toMap_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QWebHistoryItem___toMap_atList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHistoryItem) __toMap_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QWebHistoryItem___toMap_setList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QWebHistoryItem) __toMap_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistoryItem___toMap_newList(ptr.Pointer()))
}

func (ptr *QWebHistoryItem) __toMap_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebHistoryItemFromPointer(l.data).____toMap_keyList_atList(i)
			}
			return out
		}(C.QWebHistoryItem___toMap_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QWebHistoryItem) ____loadFromMap_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebHistoryItem_____loadFromMap_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QWebHistoryItem) ____loadFromMap_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QWebHistoryItem_____loadFromMap_keyList_setList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QWebHistoryItem) ____loadFromMap_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistoryItem_____loadFromMap_keyList_newList(ptr.Pointer()))
}

func (ptr *QWebHistoryItem) ____toMap_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebHistoryItem_____toMap_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QWebHistoryItem) ____toMap_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QWebHistoryItem_____toMap_keyList_setList(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QWebHistoryItem) ____toMap_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebHistoryItem_____toMap_keyList_newList(ptr.Pointer()))
}

type QWebHitTestResult struct {
	ptr unsafe.Pointer
}

type QWebHitTestResult_ITF interface {
	QWebHitTestResult_PTR() *QWebHitTestResult
}

func (ptr *QWebHitTestResult) QWebHitTestResult_PTR() *QWebHitTestResult {
	return ptr
}

func (ptr *QWebHitTestResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebHitTestResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebHitTestResult(ptr QWebHitTestResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebHitTestResult_PTR().Pointer()
	}
	return nil
}

func NewQWebHitTestResultFromPointer(ptr unsafe.Pointer) *QWebHitTestResult {
	var n = new(QWebHitTestResult)
	n.SetPointer(ptr)
	return n
}
func NewQWebHitTestResult() *QWebHitTestResult {
	var tmpValue = NewQWebHitTestResultFromPointer(C.QWebHitTestResult_NewQWebHitTestResult())
	runtime.SetFinalizer(tmpValue, (*QWebHitTestResult).DestroyQWebHitTestResult)
	return tmpValue
}

func NewQWebHitTestResult2(other QWebHitTestResult_ITF) *QWebHitTestResult {
	var tmpValue = NewQWebHitTestResultFromPointer(C.QWebHitTestResult_NewQWebHitTestResult2(PointerFromQWebHitTestResult(other)))
	runtime.SetFinalizer(tmpValue, (*QWebHitTestResult).DestroyQWebHitTestResult)
	return tmpValue
}

func (ptr *QWebHitTestResult) DestroyQWebHitTestResult() {
	if ptr.Pointer() != nil {
		C.QWebHitTestResult_DestroyQWebHitTestResult(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebHitTestResult) Pixmap() *gui.QPixmap {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPixmapFromPointer(C.QWebHitTestResult_Pixmap(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) Pos() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.QWebHitTestResult_Pos(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) BoundingRect() *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QWebHitTestResult_BoundingRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) AlternateText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebHitTestResult_AlternateText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebHitTestResult) LinkText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebHitTestResult_LinkText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebHitTestResult) LinkTitleString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebHitTestResult_LinkTitleString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebHitTestResult) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebHitTestResult_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebHitTestResult) ImageUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebHitTestResult_ImageUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) LinkUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebHitTestResult_LinkUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) MediaUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebHitTestResult_MediaUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) Element() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebHitTestResult_Element(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) EnclosingBlockElement() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebHitTestResult_EnclosingBlockElement(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) LinkElement() *QWebElement {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebElementFromPointer(C.QWebHitTestResult_LinkElement(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QWebElement).DestroyQWebElement)
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) Frame() *QWebFrame {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebFrameFromPointer(C.QWebHitTestResult_Frame(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) LinkTargetFrame() *QWebFrame {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebFrameFromPointer(C.QWebHitTestResult_LinkTargetFrame(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebHitTestResult) IsContentEditable() bool {
	if ptr.Pointer() != nil {
		return C.QWebHitTestResult_IsContentEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHitTestResult) IsContentSelected() bool {
	if ptr.Pointer() != nil {
		return C.QWebHitTestResult_IsContentSelected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHitTestResult) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QWebHitTestResult_IsNull(ptr.Pointer()) != 0
	}
	return false
}

type QWebInspector struct {
	widgets.QWidget
}

type QWebInspector_ITF interface {
	widgets.QWidget_ITF
	QWebInspector_PTR() *QWebInspector
}

func (ptr *QWebInspector) QWebInspector_PTR() *QWebInspector {
	return ptr
}

func (ptr *QWebInspector) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebInspector) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQWebInspector(ptr QWebInspector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebInspector_PTR().Pointer()
	}
	return nil
}

func NewQWebInspectorFromPointer(ptr unsafe.Pointer) *QWebInspector {
	var n = new(QWebInspector)
	n.SetPointer(ptr)
	return n
}
func NewQWebInspector(parent widgets.QWidget_ITF) *QWebInspector {
	var tmpValue = NewQWebInspectorFromPointer(C.QWebInspector_NewQWebInspector(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebInspector_Event
func callbackQWebInspector_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev)))))
}

func (ptr *QWebInspector) EventDefault(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebInspector_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

//export callbackQWebInspector_CloseEvent
func callbackQWebInspector_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QWebInspector) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQWebInspector_HideEvent
func callbackQWebInspector_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QWebInspector) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQWebInspector_ResizeEvent
func callbackQWebInspector_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWebInspector) SetPage(page QWebPage_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_SetPage(ptr.Pointer(), PointerFromQWebPage(page))
	}
}

//export callbackQWebInspector_ShowEvent
func callbackQWebInspector_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QWebInspector) DestroyQWebInspector() {
	if ptr.Pointer() != nil {
		C.QWebInspector_DestroyQWebInspector(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebInspector_SizeHint
func callbackQWebInspector_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebInspectorFromPointer(ptr).SizeHintDefault())
}

func (ptr *QWebInspector) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebInspector_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWebInspector) Page() *QWebPage {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebPageFromPointer(C.QWebInspector_Page(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebInspector) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebInspector___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebInspector) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QWebInspector) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebInspector___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *QWebInspector) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebInspector___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebInspector) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QWebInspector) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebInspector___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *QWebInspector) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebInspector___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebInspector) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QWebInspector) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebInspector___actions_newList(ptr.Pointer()))
}

func (ptr *QWebInspector) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QWebInspector___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebInspector) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebInspector) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebInspector___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QWebInspector) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebInspector___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebInspector) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebInspector) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QWebInspector___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QWebInspector) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebInspector___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebInspector) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebInspector) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QWebInspector___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QWebInspector) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebInspector___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebInspector) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebInspector) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebInspector___findChildren_newList(ptr.Pointer()))
}

func (ptr *QWebInspector) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebInspector___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebInspector) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebInspector) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebInspector___children_newList(ptr.Pointer()))
}

//export callbackQWebInspector_Close
func callbackQWebInspector_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).CloseDefault())))
}

func (ptr *QWebInspector) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QWebInspector_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebInspector_FocusNextPrevChild
func callbackQWebInspector_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QWebInspector) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QWebInspector_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQWebInspector_NativeEvent
func callbackQWebInspector_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QWebInspector) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QWebInspector_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQWebInspector_ActionEvent
func callbackQWebInspector_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQWebInspector_ChangeEvent
func callbackQWebInspector_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebInspector_ContextMenuEvent
func callbackQWebInspector_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQWebInspector_CustomContextMenuRequested
func callbackQWebInspector_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQWebInspector_DragEnterEvent
func callbackQWebInspector_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QWebInspector) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQWebInspector_DragLeaveEvent
func callbackQWebInspector_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QWebInspector) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQWebInspector_DragMoveEvent
func callbackQWebInspector_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QWebInspector) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQWebInspector_DropEvent
func callbackQWebInspector_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QWebInspector) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQWebInspector_EnterEvent
func callbackQWebInspector_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebInspector) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebInspector_FocusInEvent
func callbackQWebInspector_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWebInspector) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQWebInspector_FocusOutEvent
func callbackQWebInspector_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWebInspector) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQWebInspector_Hide
func callbackQWebInspector_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).HideDefault()
	}
}

func (ptr *QWebInspector) HideDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_HideDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_InputMethodEvent
func callbackQWebInspector_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QWebInspector) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQWebInspector_KeyPressEvent
func callbackQWebInspector_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWebInspector) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWebInspector_KeyReleaseEvent
func callbackQWebInspector_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWebInspector) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWebInspector_LeaveEvent
func callbackQWebInspector_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebInspector) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebInspector_Lower
func callbackQWebInspector_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QWebInspector) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_LowerDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_MouseDoubleClickEvent
func callbackQWebInspector_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebInspector) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebInspector_MouseMoveEvent
func callbackQWebInspector_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebInspector) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebInspector_MousePressEvent
func callbackQWebInspector_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebInspector) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebInspector_MouseReleaseEvent
func callbackQWebInspector_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebInspector) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebInspector_MoveEvent
func callbackQWebInspector_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QWebInspector) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQWebInspector_PaintEvent
func callbackQWebInspector_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QWebInspector) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQWebInspector_Raise
func callbackQWebInspector_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QWebInspector) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_Repaint
func callbackQWebInspector_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QWebInspector) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_SetDisabled
func callbackQWebInspector_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQWebInspectorFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QWebInspector) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QWebInspector_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQWebInspector_SetEnabled
func callbackQWebInspector_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQWebInspectorFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QWebInspector) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebInspector_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQWebInspector_SetFocus2
func callbackQWebInspector_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QWebInspector) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QWebInspector_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQWebInspector_SetHidden
func callbackQWebInspector_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQWebInspectorFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QWebInspector) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QWebInspector_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQWebInspector_SetStyleSheet
func callbackQWebInspector_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQWebInspectorFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QWebInspector) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QWebInspector_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQWebInspector_SetVisible
func callbackQWebInspector_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQWebInspectorFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QWebInspector) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QWebInspector_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQWebInspector_SetWindowModified
func callbackQWebInspector_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQWebInspectorFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QWebInspector) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebInspector_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQWebInspector_SetWindowTitle
func callbackQWebInspector_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQWebInspectorFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QWebInspector) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QWebInspector_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQWebInspector_Show
func callbackQWebInspector_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWebInspector) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_ShowDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_ShowFullScreen
func callbackQWebInspector_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QWebInspector) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_ShowMaximized
func callbackQWebInspector_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QWebInspector) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_ShowMinimized
func callbackQWebInspector_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QWebInspector) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_ShowNormal
func callbackQWebInspector_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QWebInspector) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_TabletEvent
func callbackQWebInspector_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QWebInspector) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQWebInspector_Update
func callbackQWebInspector_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QWebInspector) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_UpdateMicroFocus
func callbackQWebInspector_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QWebInspector) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_WheelEvent
func callbackQWebInspector_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QWebInspector) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQWebInspector_WindowIconChanged
func callbackQWebInspector_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQWebInspector_WindowTitleChanged
func callbackQWebInspector_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQWebInspector_PaintEngine
func callbackQWebInspector_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQWebInspectorFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QWebInspector) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWebInspector_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebInspector_MinimumSizeHint
func callbackQWebInspector_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebInspectorFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QWebInspector) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebInspector_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQWebInspector_InputMethodQuery
func callbackQWebInspector_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQWebInspectorFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QWebInspector) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QWebInspector_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQWebInspector_HasHeightForWidth
func callbackQWebInspector_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QWebInspector) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QWebInspector_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebInspector_HeightForWidth
func callbackQWebInspector_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQWebInspectorFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QWebInspector) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebInspector_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQWebInspector_Metric
func callbackQWebInspector_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQWebInspectorFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QWebInspector) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebInspector_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQWebInspector_EventFilter
func callbackQWebInspector_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebInspector) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebInspector_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebInspector_ChildEvent
func callbackQWebInspector_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebInspector_ConnectNotify
func callbackQWebInspector_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebInspectorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebInspector) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebInspector_CustomEvent
func callbackQWebInspector_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebInspector) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebInspector_DeleteLater
func callbackQWebInspector_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebInspector) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebInspector_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebInspector_Destroyed
func callbackQWebInspector_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebInspector_DisconnectNotify
func callbackQWebInspector_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebInspectorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebInspector) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebInspector_ObjectNameChanged
func callbackQWebInspector_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQWebInspector_TimerEvent
func callbackQWebInspector_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebInspector) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebInspector_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebInspector_MetaObject
func callbackQWebInspector_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebInspectorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebInspector) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebInspector_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebPage struct {
	core.QObject
}

type QWebPage_ITF interface {
	core.QObject_ITF
	QWebPage_PTR() *QWebPage
}

func (ptr *QWebPage) QWebPage_PTR() *QWebPage {
	return ptr
}

func (ptr *QWebPage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebPage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebPage(ptr QWebPage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebPage_PTR().Pointer()
	}
	return nil
}

func NewQWebPageFromPointer(ptr unsafe.Pointer) *QWebPage {
	var n = new(QWebPage)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QWebPage__ErrorDomain
//QWebPage::ErrorDomain
type QWebPage__ErrorDomain int64

const (
	QWebPage__QtNetwork QWebPage__ErrorDomain = QWebPage__ErrorDomain(0)
	QWebPage__Http      QWebPage__ErrorDomain = QWebPage__ErrorDomain(1)
	QWebPage__WebKit    QWebPage__ErrorDomain = QWebPage__ErrorDomain(2)
)

//go:generate stringer -type=QWebPage__Extension
//QWebPage::Extension
type QWebPage__Extension int64

const (
	QWebPage__ChooseMultipleFilesExtension QWebPage__Extension = QWebPage__Extension(0)
	QWebPage__ErrorPageExtension           QWebPage__Extension = QWebPage__Extension(1)
)

//go:generate stringer -type=QWebPage__Feature
//QWebPage::Feature
type QWebPage__Feature int64

const (
	QWebPage__Notifications QWebPage__Feature = QWebPage__Feature(0)
	QWebPage__Geolocation   QWebPage__Feature = QWebPage__Feature(1)
)

//go:generate stringer -type=QWebPage__FindFlag
//QWebPage::FindFlag
type QWebPage__FindFlag int64

const (
	QWebPage__FindBackward                      QWebPage__FindFlag = QWebPage__FindFlag(1)
	QWebPage__FindCaseSensitively               QWebPage__FindFlag = QWebPage__FindFlag(2)
	QWebPage__FindWrapsAroundDocument           QWebPage__FindFlag = QWebPage__FindFlag(4)
	QWebPage__HighlightAllOccurrences           QWebPage__FindFlag = QWebPage__FindFlag(8)
	QWebPage__FindAtWordBeginningsOnly          QWebPage__FindFlag = QWebPage__FindFlag(16)
	QWebPage__TreatMedialCapitalAsWordBeginning QWebPage__FindFlag = QWebPage__FindFlag(32)
	QWebPage__FindBeginsInSelection             QWebPage__FindFlag = QWebPage__FindFlag(64)
)

//go:generate stringer -type=QWebPage__LinkDelegationPolicy
//QWebPage::LinkDelegationPolicy
type QWebPage__LinkDelegationPolicy int64

const (
	QWebPage__DontDelegateLinks     QWebPage__LinkDelegationPolicy = QWebPage__LinkDelegationPolicy(0)
	QWebPage__DelegateExternalLinks QWebPage__LinkDelegationPolicy = QWebPage__LinkDelegationPolicy(1)
	QWebPage__DelegateAllLinks      QWebPage__LinkDelegationPolicy = QWebPage__LinkDelegationPolicy(2)
)

//go:generate stringer -type=QWebPage__MessageLevel
//QWebPage::MessageLevel
type QWebPage__MessageLevel int64

const (
	QWebPage__LogMessageLevel     QWebPage__MessageLevel = QWebPage__MessageLevel(1)
	QWebPage__WarningMessageLevel QWebPage__MessageLevel = QWebPage__MessageLevel(2)
	QWebPage__ErrorMessageLevel   QWebPage__MessageLevel = QWebPage__MessageLevel(3)
	QWebPage__DebugMessageLevel   QWebPage__MessageLevel = QWebPage__MessageLevel(4)
	QWebPage__InfoMessageLevel    QWebPage__MessageLevel = QWebPage__MessageLevel(5)
)

//go:generate stringer -type=QWebPage__MessageSource
//QWebPage::MessageSource
type QWebPage__MessageSource int64

const (
	QWebPage__XmlMessageSource            QWebPage__MessageSource = QWebPage__MessageSource(0)
	QWebPage__JSMessageSource             QWebPage__MessageSource = QWebPage__MessageSource(1)
	QWebPage__NetworkMessageSource        QWebPage__MessageSource = QWebPage__MessageSource(2)
	QWebPage__ConsoleAPIMessageSource     QWebPage__MessageSource = QWebPage__MessageSource(3)
	QWebPage__StorageMessageSource        QWebPage__MessageSource = QWebPage__MessageSource(4)
	QWebPage__AppCacheMessageSource       QWebPage__MessageSource = QWebPage__MessageSource(5)
	QWebPage__RenderingMessageSource      QWebPage__MessageSource = QWebPage__MessageSource(6)
	QWebPage__CSSMessageSource            QWebPage__MessageSource = QWebPage__MessageSource(7)
	QWebPage__SecurityMessageSource       QWebPage__MessageSource = QWebPage__MessageSource(8)
	QWebPage__ContentBlockerMessageSource QWebPage__MessageSource = QWebPage__MessageSource(9)
	QWebPage__OtherMessageSource          QWebPage__MessageSource = QWebPage__MessageSource(10)
)

//go:generate stringer -type=QWebPage__NavigationType
//QWebPage::NavigationType
type QWebPage__NavigationType int64

const (
	QWebPage__NavigationTypeLinkClicked     QWebPage__NavigationType = QWebPage__NavigationType(0)
	QWebPage__NavigationTypeFormSubmitted   QWebPage__NavigationType = QWebPage__NavigationType(1)
	QWebPage__NavigationTypeBackOrForward   QWebPage__NavigationType = QWebPage__NavigationType(2)
	QWebPage__NavigationTypeReload          QWebPage__NavigationType = QWebPage__NavigationType(3)
	QWebPage__NavigationTypeFormResubmitted QWebPage__NavigationType = QWebPage__NavigationType(4)
	QWebPage__NavigationTypeOther           QWebPage__NavigationType = QWebPage__NavigationType(5)
)

//go:generate stringer -type=QWebPage__PermissionPolicy
//QWebPage::PermissionPolicy
type QWebPage__PermissionPolicy int64

const (
	QWebPage__PermissionUnknown       QWebPage__PermissionPolicy = QWebPage__PermissionPolicy(0)
	QWebPage__PermissionGrantedByUser QWebPage__PermissionPolicy = QWebPage__PermissionPolicy(1)
	QWebPage__PermissionDeniedByUser  QWebPage__PermissionPolicy = QWebPage__PermissionPolicy(2)
)

//go:generate stringer -type=QWebPage__VisibilityState
//QWebPage::VisibilityState
type QWebPage__VisibilityState int64

const (
	QWebPage__VisibilityStateVisible   QWebPage__VisibilityState = QWebPage__VisibilityState(0)
	QWebPage__VisibilityStateHidden    QWebPage__VisibilityState = QWebPage__VisibilityState(1)
	QWebPage__VisibilityStatePrerender QWebPage__VisibilityState = QWebPage__VisibilityState(2)
	QWebPage__VisibilityStateUnloaded  QWebPage__VisibilityState = QWebPage__VisibilityState(3)
)

//go:generate stringer -type=QWebPage__WebAction
//QWebPage::WebAction
type QWebPage__WebAction int64

const (
	QWebPage__NoWebAction                 QWebPage__WebAction = QWebPage__WebAction(-1)
	QWebPage__OpenLink                    QWebPage__WebAction = QWebPage__WebAction(0)
	QWebPage__OpenLinkInNewWindow         QWebPage__WebAction = QWebPage__WebAction(1)
	QWebPage__OpenFrameInNewWindow        QWebPage__WebAction = QWebPage__WebAction(2)
	QWebPage__DownloadLinkToDisk          QWebPage__WebAction = QWebPage__WebAction(3)
	QWebPage__CopyLinkToClipboard         QWebPage__WebAction = QWebPage__WebAction(4)
	QWebPage__OpenImageInNewWindow        QWebPage__WebAction = QWebPage__WebAction(5)
	QWebPage__DownloadImageToDisk         QWebPage__WebAction = QWebPage__WebAction(6)
	QWebPage__CopyImageToClipboard        QWebPage__WebAction = QWebPage__WebAction(7)
	QWebPage__Back                        QWebPage__WebAction = QWebPage__WebAction(8)
	QWebPage__Forward                     QWebPage__WebAction = QWebPage__WebAction(9)
	QWebPage__Stop                        QWebPage__WebAction = QWebPage__WebAction(10)
	QWebPage__Reload                      QWebPage__WebAction = QWebPage__WebAction(11)
	QWebPage__Cut                         QWebPage__WebAction = QWebPage__WebAction(12)
	QWebPage__Copy                        QWebPage__WebAction = QWebPage__WebAction(13)
	QWebPage__Paste                       QWebPage__WebAction = QWebPage__WebAction(14)
	QWebPage__Undo                        QWebPage__WebAction = QWebPage__WebAction(15)
	QWebPage__Redo                        QWebPage__WebAction = QWebPage__WebAction(16)
	QWebPage__MoveToNextChar              QWebPage__WebAction = QWebPage__WebAction(17)
	QWebPage__MoveToPreviousChar          QWebPage__WebAction = QWebPage__WebAction(18)
	QWebPage__MoveToNextWord              QWebPage__WebAction = QWebPage__WebAction(19)
	QWebPage__MoveToPreviousWord          QWebPage__WebAction = QWebPage__WebAction(20)
	QWebPage__MoveToNextLine              QWebPage__WebAction = QWebPage__WebAction(21)
	QWebPage__MoveToPreviousLine          QWebPage__WebAction = QWebPage__WebAction(22)
	QWebPage__MoveToStartOfLine           QWebPage__WebAction = QWebPage__WebAction(23)
	QWebPage__MoveToEndOfLine             QWebPage__WebAction = QWebPage__WebAction(24)
	QWebPage__MoveToStartOfBlock          QWebPage__WebAction = QWebPage__WebAction(25)
	QWebPage__MoveToEndOfBlock            QWebPage__WebAction = QWebPage__WebAction(26)
	QWebPage__MoveToStartOfDocument       QWebPage__WebAction = QWebPage__WebAction(27)
	QWebPage__MoveToEndOfDocument         QWebPage__WebAction = QWebPage__WebAction(28)
	QWebPage__SelectNextChar              QWebPage__WebAction = QWebPage__WebAction(29)
	QWebPage__SelectPreviousChar          QWebPage__WebAction = QWebPage__WebAction(30)
	QWebPage__SelectNextWord              QWebPage__WebAction = QWebPage__WebAction(31)
	QWebPage__SelectPreviousWord          QWebPage__WebAction = QWebPage__WebAction(32)
	QWebPage__SelectNextLine              QWebPage__WebAction = QWebPage__WebAction(33)
	QWebPage__SelectPreviousLine          QWebPage__WebAction = QWebPage__WebAction(34)
	QWebPage__SelectStartOfLine           QWebPage__WebAction = QWebPage__WebAction(35)
	QWebPage__SelectEndOfLine             QWebPage__WebAction = QWebPage__WebAction(36)
	QWebPage__SelectStartOfBlock          QWebPage__WebAction = QWebPage__WebAction(37)
	QWebPage__SelectEndOfBlock            QWebPage__WebAction = QWebPage__WebAction(38)
	QWebPage__SelectStartOfDocument       QWebPage__WebAction = QWebPage__WebAction(39)
	QWebPage__SelectEndOfDocument         QWebPage__WebAction = QWebPage__WebAction(40)
	QWebPage__DeleteStartOfWord           QWebPage__WebAction = QWebPage__WebAction(41)
	QWebPage__DeleteEndOfWord             QWebPage__WebAction = QWebPage__WebAction(42)
	QWebPage__SetTextDirectionDefault     QWebPage__WebAction = QWebPage__WebAction(43)
	QWebPage__SetTextDirectionLeftToRight QWebPage__WebAction = QWebPage__WebAction(44)
	QWebPage__SetTextDirectionRightToLeft QWebPage__WebAction = QWebPage__WebAction(45)
	QWebPage__ToggleBold                  QWebPage__WebAction = QWebPage__WebAction(46)
	QWebPage__ToggleItalic                QWebPage__WebAction = QWebPage__WebAction(47)
	QWebPage__ToggleUnderline             QWebPage__WebAction = QWebPage__WebAction(48)
	QWebPage__InspectElement              QWebPage__WebAction = QWebPage__WebAction(49)
	QWebPage__InsertParagraphSeparator    QWebPage__WebAction = QWebPage__WebAction(50)
	QWebPage__InsertLineSeparator         QWebPage__WebAction = QWebPage__WebAction(51)
	QWebPage__SelectAll                   QWebPage__WebAction = QWebPage__WebAction(52)
	QWebPage__ReloadAndBypassCache        QWebPage__WebAction = QWebPage__WebAction(53)
	QWebPage__PasteAndMatchStyle          QWebPage__WebAction = QWebPage__WebAction(54)
	QWebPage__RemoveFormat                QWebPage__WebAction = QWebPage__WebAction(55)
	QWebPage__ToggleStrikethrough         QWebPage__WebAction = QWebPage__WebAction(56)
	QWebPage__ToggleSubscript             QWebPage__WebAction = QWebPage__WebAction(57)
	QWebPage__ToggleSuperscript           QWebPage__WebAction = QWebPage__WebAction(58)
	QWebPage__InsertUnorderedList         QWebPage__WebAction = QWebPage__WebAction(59)
	QWebPage__InsertOrderedList           QWebPage__WebAction = QWebPage__WebAction(60)
	QWebPage__Indent                      QWebPage__WebAction = QWebPage__WebAction(61)
	QWebPage__Outdent                     QWebPage__WebAction = QWebPage__WebAction(62)
	QWebPage__AlignCenter                 QWebPage__WebAction = QWebPage__WebAction(63)
	QWebPage__AlignJustified              QWebPage__WebAction = QWebPage__WebAction(64)
	QWebPage__AlignLeft                   QWebPage__WebAction = QWebPage__WebAction(65)
	QWebPage__AlignRight                  QWebPage__WebAction = QWebPage__WebAction(66)
	QWebPage__StopScheduledPageRefresh    QWebPage__WebAction = QWebPage__WebAction(67)
	QWebPage__CopyImageUrlToClipboard     QWebPage__WebAction = QWebPage__WebAction(68)
	QWebPage__OpenLinkInThisWindow        QWebPage__WebAction = QWebPage__WebAction(69)
	QWebPage__DownloadMediaToDisk         QWebPage__WebAction = QWebPage__WebAction(70)
	QWebPage__CopyMediaUrlToClipboard     QWebPage__WebAction = QWebPage__WebAction(71)
	QWebPage__ToggleMediaControls         QWebPage__WebAction = QWebPage__WebAction(72)
	QWebPage__ToggleMediaLoop             QWebPage__WebAction = QWebPage__WebAction(73)
	QWebPage__ToggleMediaPlayPause        QWebPage__WebAction = QWebPage__WebAction(74)
	QWebPage__ToggleMediaMute             QWebPage__WebAction = QWebPage__WebAction(75)
	QWebPage__ToggleVideoFullscreen       QWebPage__WebAction = QWebPage__WebAction(76)
	QWebPage__WebActionCount              QWebPage__WebAction = QWebPage__WebAction(77)
)

//go:generate stringer -type=QWebPage__WebWindowType
//QWebPage::WebWindowType
type QWebPage__WebWindowType int64

const (
	QWebPage__WebBrowserWindow QWebPage__WebWindowType = QWebPage__WebWindowType(0)
	QWebPage__WebModalDialog   QWebPage__WebWindowType = QWebPage__WebWindowType(1)
)

func (ptr *QWebPage) CreateStandardContextMenu() *widgets.QMenu {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQMenuFromPointer(C.QWebPage_CreateStandardContextMenu(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebPage_CreatePlugin
func callbackQWebPage_CreatePlugin(ptr unsafe.Pointer, classid C.struct_QtWebKit_PackedString, url unsafe.Pointer, paramNames C.struct_QtWebKit_PackedString, paramValues C.struct_QtWebKit_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createPlugin"); signal != nil {
		return core.PointerFromQObject(signal.(func(string, *core.QUrl, []string, []string) *core.QObject)(cGoUnpackString(classid), core.NewQUrlFromPointer(url), strings.Split(cGoUnpackString(paramNames), "|"), strings.Split(cGoUnpackString(paramValues), "|")))
	}

	return core.PointerFromQObject(NewQWebPageFromPointer(ptr).CreatePluginDefault(cGoUnpackString(classid), core.NewQUrlFromPointer(url), strings.Split(cGoUnpackString(paramNames), "|"), strings.Split(cGoUnpackString(paramValues), "|")))
}

func (ptr *QWebPage) ConnectCreatePlugin(f func(classid string, url *core.QUrl, paramNames []string, paramValues []string) *core.QObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createPlugin"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createPlugin", func(classid string, url *core.QUrl, paramNames []string, paramValues []string) *core.QObject {
				signal.(func(string, *core.QUrl, []string, []string) *core.QObject)(classid, url, paramNames, paramValues)
				return f(classid, url, paramNames, paramValues)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createPlugin", f)
		}
	}
}

func (ptr *QWebPage) DisconnectCreatePlugin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createPlugin")
	}
}

func (ptr *QWebPage) CreatePlugin(classid string, url core.QUrl_ITF, paramNames []string, paramValues []string) *core.QObject {
	if ptr.Pointer() != nil {
		var classidC *C.char
		if classid != "" {
			classidC = C.CString(classid)
			defer C.free(unsafe.Pointer(classidC))
		}
		var paramNamesC = C.CString(strings.Join(paramNames, "|"))
		defer C.free(unsafe.Pointer(paramNamesC))
		var paramValuesC = C.CString(strings.Join(paramValues, "|"))
		defer C.free(unsafe.Pointer(paramValuesC))
		var tmpValue = core.NewQObjectFromPointer(C.QWebPage_CreatePlugin(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: classidC, len: C.longlong(len(classid))}, core.PointerFromQUrl(url), C.struct_QtWebKit_PackedString{data: paramNamesC, len: C.longlong(len(strings.Join(paramNames, "|")))}, C.struct_QtWebKit_PackedString{data: paramValuesC, len: C.longlong(len(strings.Join(paramValues, "|")))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) CreatePluginDefault(classid string, url core.QUrl_ITF, paramNames []string, paramValues []string) *core.QObject {
	if ptr.Pointer() != nil {
		var classidC *C.char
		if classid != "" {
			classidC = C.CString(classid)
			defer C.free(unsafe.Pointer(classidC))
		}
		var paramNamesC = C.CString(strings.Join(paramNames, "|"))
		defer C.free(unsafe.Pointer(paramNamesC))
		var paramValuesC = C.CString(strings.Join(paramValues, "|"))
		defer C.free(unsafe.Pointer(paramValuesC))
		var tmpValue = core.NewQObjectFromPointer(C.QWebPage_CreatePluginDefault(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: classidC, len: C.longlong(len(classid))}, core.PointerFromQUrl(url), C.struct_QtWebKit_PackedString{data: paramNamesC, len: C.longlong(len(strings.Join(paramNames, "|")))}, C.struct_QtWebKit_PackedString{data: paramValuesC, len: C.longlong(len(strings.Join(paramValues, "|")))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebPage_ChooseFile
func callbackQWebPage_ChooseFile(ptr unsafe.Pointer, parentFrame unsafe.Pointer, suggestedFile C.struct_QtWebKit_PackedString) C.struct_QtWebKit_PackedString {
	if signal := qt.GetSignal(ptr, "chooseFile"); signal != nil {
		tempVal := signal.(func(*QWebFrame, string) string)(NewQWebFrameFromPointer(parentFrame), cGoUnpackString(suggestedFile))
		return C.struct_QtWebKit_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQWebPageFromPointer(ptr).ChooseFileDefault(NewQWebFrameFromPointer(parentFrame), cGoUnpackString(suggestedFile))
	return C.struct_QtWebKit_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QWebPage) ConnectChooseFile(f func(parentFrame *QWebFrame, suggestedFile string) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "chooseFile"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "chooseFile", func(parentFrame *QWebFrame, suggestedFile string) string {
				signal.(func(*QWebFrame, string) string)(parentFrame, suggestedFile)
				return f(parentFrame, suggestedFile)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "chooseFile", f)
		}
	}
}

func (ptr *QWebPage) DisconnectChooseFile() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "chooseFile")
	}
}

func (ptr *QWebPage) ChooseFile(parentFrame QWebFrame_ITF, suggestedFile string) string {
	if ptr.Pointer() != nil {
		var suggestedFileC *C.char
		if suggestedFile != "" {
			suggestedFileC = C.CString(suggestedFile)
			defer C.free(unsafe.Pointer(suggestedFileC))
		}
		return cGoUnpackString(C.QWebPage_ChooseFile(ptr.Pointer(), PointerFromQWebFrame(parentFrame), C.struct_QtWebKit_PackedString{data: suggestedFileC, len: C.longlong(len(suggestedFile))}))
	}
	return ""
}

func (ptr *QWebPage) ChooseFileDefault(parentFrame QWebFrame_ITF, suggestedFile string) string {
	if ptr.Pointer() != nil {
		var suggestedFileC *C.char
		if suggestedFile != "" {
			suggestedFileC = C.CString(suggestedFile)
			defer C.free(unsafe.Pointer(suggestedFileC))
		}
		return cGoUnpackString(C.QWebPage_ChooseFileDefault(ptr.Pointer(), PointerFromQWebFrame(parentFrame), C.struct_QtWebKit_PackedString{data: suggestedFileC, len: C.longlong(len(suggestedFile))}))
	}
	return ""
}

//export callbackQWebPage_CreateWindow
func callbackQWebPage_CreateWindow(ptr unsafe.Pointer, ty C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createWindow"); signal != nil {
		return PointerFromQWebPage(signal.(func(QWebPage__WebWindowType) *QWebPage)(QWebPage__WebWindowType(ty)))
	}

	return PointerFromQWebPage(NewQWebPageFromPointer(ptr).CreateWindowDefault(QWebPage__WebWindowType(ty)))
}

func (ptr *QWebPage) ConnectCreateWindow(f func(ty QWebPage__WebWindowType) *QWebPage) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createWindow", func(ty QWebPage__WebWindowType) *QWebPage {
				signal.(func(QWebPage__WebWindowType) *QWebPage)(ty)
				return f(ty)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createWindow", f)
		}
	}
}

func (ptr *QWebPage) DisconnectCreateWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createWindow")
	}
}

func (ptr *QWebPage) CreateWindow(ty QWebPage__WebWindowType) *QWebPage {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebPageFromPointer(C.QWebPage_CreateWindow(ptr.Pointer(), C.longlong(ty)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) CreateWindowDefault(ty QWebPage__WebWindowType) *QWebPage {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebPageFromPointer(C.QWebPage_CreateWindowDefault(ptr.Pointer(), C.longlong(ty)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQWebPage(parent core.QObject_ITF) *QWebPage {
	var tmpValue = NewQWebPageFromPointer(C.QWebPage_NewQWebPage(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebPage_AcceptNavigationRequest
func callbackQWebPage_AcceptNavigationRequest(ptr unsafe.Pointer, frame unsafe.Pointer, request unsafe.Pointer, ty C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "acceptNavigationRequest"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QWebFrame, *network.QNetworkRequest, QWebPage__NavigationType) bool)(NewQWebFrameFromPointer(frame), network.NewQNetworkRequestFromPointer(request), QWebPage__NavigationType(ty)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).AcceptNavigationRequestDefault(NewQWebFrameFromPointer(frame), network.NewQNetworkRequestFromPointer(request), QWebPage__NavigationType(ty)))))
}

func (ptr *QWebPage) ConnectAcceptNavigationRequest(f func(frame *QWebFrame, request *network.QNetworkRequest, ty QWebPage__NavigationType) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "acceptNavigationRequest"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "acceptNavigationRequest", func(frame *QWebFrame, request *network.QNetworkRequest, ty QWebPage__NavigationType) bool {
				signal.(func(*QWebFrame, *network.QNetworkRequest, QWebPage__NavigationType) bool)(frame, request, ty)
				return f(frame, request, ty)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "acceptNavigationRequest", f)
		}
	}
}

func (ptr *QWebPage) DisconnectAcceptNavigationRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "acceptNavigationRequest")
	}
}

func (ptr *QWebPage) AcceptNavigationRequest(frame QWebFrame_ITF, request network.QNetworkRequest_ITF, ty QWebPage__NavigationType) bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_AcceptNavigationRequest(ptr.Pointer(), PointerFromQWebFrame(frame), network.PointerFromQNetworkRequest(request), C.longlong(ty)) != 0
	}
	return false
}

func (ptr *QWebPage) AcceptNavigationRequestDefault(frame QWebFrame_ITF, request network.QNetworkRequest_ITF, ty QWebPage__NavigationType) bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_AcceptNavigationRequestDefault(ptr.Pointer(), PointerFromQWebFrame(frame), network.PointerFromQNetworkRequest(request), C.longlong(ty)) != 0
	}
	return false
}

//export callbackQWebPage_Event
func callbackQWebPage_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev)))))
}

func (ptr *QWebPage) EventDefault(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QWebPage) FindText(subString string, options QWebPage__FindFlag) bool {
	if ptr.Pointer() != nil {
		var subStringC *C.char
		if subString != "" {
			subStringC = C.CString(subString)
			defer C.free(unsafe.Pointer(subStringC))
		}
		return C.QWebPage_FindText(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: subStringC, len: C.longlong(len(subString))}, C.longlong(options)) != 0
	}
	return false
}

func (ptr *QWebPage) FocusNextPrevChild(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQWebPage_JavaScriptConfirm
func callbackQWebPage_JavaScriptConfirm(ptr unsafe.Pointer, frame unsafe.Pointer, msg C.struct_QtWebKit_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "javaScriptConfirm"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QWebFrame, string) bool)(NewQWebFrameFromPointer(frame), cGoUnpackString(msg)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).JavaScriptConfirmDefault(NewQWebFrameFromPointer(frame), cGoUnpackString(msg)))))
}

func (ptr *QWebPage) ConnectJavaScriptConfirm(f func(frame *QWebFrame, msg string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "javaScriptConfirm"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptConfirm", func(frame *QWebFrame, msg string) bool {
				signal.(func(*QWebFrame, string) bool)(frame, msg)
				return f(frame, msg)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptConfirm", f)
		}
	}
}

func (ptr *QWebPage) DisconnectJavaScriptConfirm() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "javaScriptConfirm")
	}
}

func (ptr *QWebPage) JavaScriptConfirm(frame QWebFrame_ITF, msg string) bool {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		return C.QWebPage_JavaScriptConfirm(ptr.Pointer(), PointerFromQWebFrame(frame), C.struct_QtWebKit_PackedString{data: msgC, len: C.longlong(len(msg))}) != 0
	}
	return false
}

func (ptr *QWebPage) JavaScriptConfirmDefault(frame QWebFrame_ITF, msg string) bool {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		return C.QWebPage_JavaScriptConfirmDefault(ptr.Pointer(), PointerFromQWebFrame(frame), C.struct_QtWebKit_PackedString{data: msgC, len: C.longlong(len(msg))}) != 0
	}
	return false
}

//export callbackQWebPage_JavaScriptPrompt
func callbackQWebPage_JavaScriptPrompt(ptr unsafe.Pointer, frame unsafe.Pointer, msg C.struct_QtWebKit_PackedString, defaultValue C.struct_QtWebKit_PackedString, result C.struct_QtWebKit_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "javaScriptPrompt"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QWebFrame, string, string, string) bool)(NewQWebFrameFromPointer(frame), cGoUnpackString(msg), cGoUnpackString(defaultValue), cGoUnpackString(result)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).JavaScriptPromptDefault(NewQWebFrameFromPointer(frame), cGoUnpackString(msg), cGoUnpackString(defaultValue), cGoUnpackString(result)))))
}

func (ptr *QWebPage) ConnectJavaScriptPrompt(f func(frame *QWebFrame, msg string, defaultValue string, result string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "javaScriptPrompt"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptPrompt", func(frame *QWebFrame, msg string, defaultValue string, result string) bool {
				signal.(func(*QWebFrame, string, string, string) bool)(frame, msg, defaultValue, result)
				return f(frame, msg, defaultValue, result)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptPrompt", f)
		}
	}
}

func (ptr *QWebPage) DisconnectJavaScriptPrompt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "javaScriptPrompt")
	}
}

func (ptr *QWebPage) JavaScriptPrompt(frame QWebFrame_ITF, msg string, defaultValue string, result string) bool {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		var defaultValueC *C.char
		if defaultValue != "" {
			defaultValueC = C.CString(defaultValue)
			defer C.free(unsafe.Pointer(defaultValueC))
		}
		var resultC *C.char
		if result != "" {
			resultC = C.CString(result)
			defer C.free(unsafe.Pointer(resultC))
		}
		return C.QWebPage_JavaScriptPrompt(ptr.Pointer(), PointerFromQWebFrame(frame), C.struct_QtWebKit_PackedString{data: msgC, len: C.longlong(len(msg))}, C.struct_QtWebKit_PackedString{data: defaultValueC, len: C.longlong(len(defaultValue))}, C.struct_QtWebKit_PackedString{data: resultC, len: C.longlong(len(result))}) != 0
	}
	return false
}

func (ptr *QWebPage) JavaScriptPromptDefault(frame QWebFrame_ITF, msg string, defaultValue string, result string) bool {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		var defaultValueC *C.char
		if defaultValue != "" {
			defaultValueC = C.CString(defaultValue)
			defer C.free(unsafe.Pointer(defaultValueC))
		}
		var resultC *C.char
		if result != "" {
			resultC = C.CString(result)
			defer C.free(unsafe.Pointer(resultC))
		}
		return C.QWebPage_JavaScriptPromptDefault(ptr.Pointer(), PointerFromQWebFrame(frame), C.struct_QtWebKit_PackedString{data: msgC, len: C.longlong(len(msg))}, C.struct_QtWebKit_PackedString{data: defaultValueC, len: C.longlong(len(defaultValue))}, C.struct_QtWebKit_PackedString{data: resultC, len: C.longlong(len(result))}) != 0
	}
	return false
}

//export callbackQWebPage_ShouldInterruptJavaScript
func callbackQWebPage_ShouldInterruptJavaScript(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "shouldInterruptJavaScript"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).ShouldInterruptJavaScriptDefault())))
}

func (ptr *QWebPage) ConnectShouldInterruptJavaScript(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "shouldInterruptJavaScript"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "shouldInterruptJavaScript", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shouldInterruptJavaScript", f)
		}
	}
}

func (ptr *QWebPage) DisconnectShouldInterruptJavaScript() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "shouldInterruptJavaScript")
	}
}

func (ptr *QWebPage) ShouldInterruptJavaScript() bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_ShouldInterruptJavaScript(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) ShouldInterruptJavaScriptDefault() bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_ShouldInterruptJavaScriptDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) SwallowContextMenuEvent(event gui.QContextMenuEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_SwallowContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event)) != 0
	}
	return false
}

//export callbackQWebPage_ApplicationCacheQuotaExceeded
func callbackQWebPage_ApplicationCacheQuotaExceeded(ptr unsafe.Pointer, origin unsafe.Pointer, defaultOriginQuota C.ulonglong, totalSpaceNeeded C.ulonglong) {
	if signal := qt.GetSignal(ptr, "applicationCacheQuotaExceeded"); signal != nil {
		signal.(func(*QWebSecurityOrigin, uint64, uint64))(NewQWebSecurityOriginFromPointer(origin), uint64(defaultOriginQuota), uint64(totalSpaceNeeded))
	}

}

func (ptr *QWebPage) ConnectApplicationCacheQuotaExceeded(f func(origin *QWebSecurityOrigin, defaultOriginQuota uint64, totalSpaceNeeded uint64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "applicationCacheQuotaExceeded") {
			C.QWebPage_ConnectApplicationCacheQuotaExceeded(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "applicationCacheQuotaExceeded"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "applicationCacheQuotaExceeded", func(origin *QWebSecurityOrigin, defaultOriginQuota uint64, totalSpaceNeeded uint64) {
				signal.(func(*QWebSecurityOrigin, uint64, uint64))(origin, defaultOriginQuota, totalSpaceNeeded)
				f(origin, defaultOriginQuota, totalSpaceNeeded)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "applicationCacheQuotaExceeded", f)
		}
	}
}

func (ptr *QWebPage) DisconnectApplicationCacheQuotaExceeded() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectApplicationCacheQuotaExceeded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "applicationCacheQuotaExceeded")
	}
}

func (ptr *QWebPage) ApplicationCacheQuotaExceeded(origin QWebSecurityOrigin_ITF, defaultOriginQuota uint64, totalSpaceNeeded uint64) {
	if ptr.Pointer() != nil {
		C.QWebPage_ApplicationCacheQuotaExceeded(ptr.Pointer(), PointerFromQWebSecurityOrigin(origin), C.ulonglong(defaultOriginQuota), C.ulonglong(totalSpaceNeeded))
	}
}

//export callbackQWebPage_ConsoleMessageReceived
func callbackQWebPage_ConsoleMessageReceived(ptr unsafe.Pointer, source C.longlong, level C.longlong, message C.struct_QtWebKit_PackedString, lineNumber C.int, sourceID C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "consoleMessageReceived"); signal != nil {
		signal.(func(QWebPage__MessageSource, QWebPage__MessageLevel, string, int, string))(QWebPage__MessageSource(source), QWebPage__MessageLevel(level), cGoUnpackString(message), int(int32(lineNumber)), cGoUnpackString(sourceID))
	}

}

func (ptr *QWebPage) ConnectConsoleMessageReceived(f func(source QWebPage__MessageSource, level QWebPage__MessageLevel, message string, lineNumber int, sourceID string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "consoleMessageReceived") {
			C.QWebPage_ConnectConsoleMessageReceived(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "consoleMessageReceived"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "consoleMessageReceived", func(source QWebPage__MessageSource, level QWebPage__MessageLevel, message string, lineNumber int, sourceID string) {
				signal.(func(QWebPage__MessageSource, QWebPage__MessageLevel, string, int, string))(source, level, message, lineNumber, sourceID)
				f(source, level, message, lineNumber, sourceID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "consoleMessageReceived", f)
		}
	}
}

func (ptr *QWebPage) DisconnectConsoleMessageReceived() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectConsoleMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "consoleMessageReceived")
	}
}

func (ptr *QWebPage) ConsoleMessageReceived(source QWebPage__MessageSource, level QWebPage__MessageLevel, message string, lineNumber int, sourceID string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		var sourceIDC *C.char
		if sourceID != "" {
			sourceIDC = C.CString(sourceID)
			defer C.free(unsafe.Pointer(sourceIDC))
		}
		C.QWebPage_ConsoleMessageReceived(ptr.Pointer(), C.longlong(source), C.longlong(level), C.struct_QtWebKit_PackedString{data: messageC, len: C.longlong(len(message))}, C.int(int32(lineNumber)), C.struct_QtWebKit_PackedString{data: sourceIDC, len: C.longlong(len(sourceID))})
	}
}

//export callbackQWebPage_ContentsChanged
func callbackQWebPage_ContentsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contentsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectContentsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "contentsChanged") {
			C.QWebPage_ConnectContentsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "contentsChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contentsChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentsChanged", f)
		}
	}
}

func (ptr *QWebPage) DisconnectContentsChanged() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectContentsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "contentsChanged")
	}
}

func (ptr *QWebPage) ContentsChanged() {
	if ptr.Pointer() != nil {
		C.QWebPage_ContentsChanged(ptr.Pointer())
	}
}

//export callbackQWebPage_DatabaseQuotaExceeded
func callbackQWebPage_DatabaseQuotaExceeded(ptr unsafe.Pointer, frame unsafe.Pointer, databaseName C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "databaseQuotaExceeded"); signal != nil {
		signal.(func(*QWebFrame, string))(NewQWebFrameFromPointer(frame), cGoUnpackString(databaseName))
	}

}

func (ptr *QWebPage) ConnectDatabaseQuotaExceeded(f func(frame *QWebFrame, databaseName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "databaseQuotaExceeded") {
			C.QWebPage_ConnectDatabaseQuotaExceeded(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "databaseQuotaExceeded"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "databaseQuotaExceeded", func(frame *QWebFrame, databaseName string) {
				signal.(func(*QWebFrame, string))(frame, databaseName)
				f(frame, databaseName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "databaseQuotaExceeded", f)
		}
	}
}

func (ptr *QWebPage) DisconnectDatabaseQuotaExceeded() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectDatabaseQuotaExceeded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "databaseQuotaExceeded")
	}
}

func (ptr *QWebPage) DatabaseQuotaExceeded(frame QWebFrame_ITF, databaseName string) {
	if ptr.Pointer() != nil {
		var databaseNameC *C.char
		if databaseName != "" {
			databaseNameC = C.CString(databaseName)
			defer C.free(unsafe.Pointer(databaseNameC))
		}
		C.QWebPage_DatabaseQuotaExceeded(ptr.Pointer(), PointerFromQWebFrame(frame), C.struct_QtWebKit_PackedString{data: databaseNameC, len: C.longlong(len(databaseName))})
	}
}

//export callbackQWebPage_DownloadRequested
func callbackQWebPage_DownloadRequested(ptr unsafe.Pointer, request unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "downloadRequested"); signal != nil {
		signal.(func(*network.QNetworkRequest))(network.NewQNetworkRequestFromPointer(request))
	}

}

func (ptr *QWebPage) ConnectDownloadRequested(f func(request *network.QNetworkRequest)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "downloadRequested") {
			C.QWebPage_ConnectDownloadRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "downloadRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "downloadRequested", func(request *network.QNetworkRequest) {
				signal.(func(*network.QNetworkRequest))(request)
				f(request)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "downloadRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectDownloadRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectDownloadRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "downloadRequested")
	}
}

func (ptr *QWebPage) DownloadRequested(request network.QNetworkRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_DownloadRequested(ptr.Pointer(), network.PointerFromQNetworkRequest(request))
	}
}

//export callbackQWebPage_FeaturePermissionRequestCanceled
func callbackQWebPage_FeaturePermissionRequestCanceled(ptr unsafe.Pointer, frame unsafe.Pointer, feature C.longlong) {
	if signal := qt.GetSignal(ptr, "featurePermissionRequestCanceled"); signal != nil {
		signal.(func(*QWebFrame, QWebPage__Feature))(NewQWebFrameFromPointer(frame), QWebPage__Feature(feature))
	}

}

func (ptr *QWebPage) ConnectFeaturePermissionRequestCanceled(f func(frame *QWebFrame, feature QWebPage__Feature)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "featurePermissionRequestCanceled") {
			C.QWebPage_ConnectFeaturePermissionRequestCanceled(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "featurePermissionRequestCanceled"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "featurePermissionRequestCanceled", func(frame *QWebFrame, feature QWebPage__Feature) {
				signal.(func(*QWebFrame, QWebPage__Feature))(frame, feature)
				f(frame, feature)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "featurePermissionRequestCanceled", f)
		}
	}
}

func (ptr *QWebPage) DisconnectFeaturePermissionRequestCanceled() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectFeaturePermissionRequestCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "featurePermissionRequestCanceled")
	}
}

func (ptr *QWebPage) FeaturePermissionRequestCanceled(frame QWebFrame_ITF, feature QWebPage__Feature) {
	if ptr.Pointer() != nil {
		C.QWebPage_FeaturePermissionRequestCanceled(ptr.Pointer(), PointerFromQWebFrame(frame), C.longlong(feature))
	}
}

//export callbackQWebPage_FeaturePermissionRequested
func callbackQWebPage_FeaturePermissionRequested(ptr unsafe.Pointer, frame unsafe.Pointer, feature C.longlong) {
	if signal := qt.GetSignal(ptr, "featurePermissionRequested"); signal != nil {
		signal.(func(*QWebFrame, QWebPage__Feature))(NewQWebFrameFromPointer(frame), QWebPage__Feature(feature))
	}

}

func (ptr *QWebPage) ConnectFeaturePermissionRequested(f func(frame *QWebFrame, feature QWebPage__Feature)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "featurePermissionRequested") {
			C.QWebPage_ConnectFeaturePermissionRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "featurePermissionRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "featurePermissionRequested", func(frame *QWebFrame, feature QWebPage__Feature) {
				signal.(func(*QWebFrame, QWebPage__Feature))(frame, feature)
				f(frame, feature)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "featurePermissionRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectFeaturePermissionRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectFeaturePermissionRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "featurePermissionRequested")
	}
}

func (ptr *QWebPage) FeaturePermissionRequested(frame QWebFrame_ITF, feature QWebPage__Feature) {
	if ptr.Pointer() != nil {
		C.QWebPage_FeaturePermissionRequested(ptr.Pointer(), PointerFromQWebFrame(frame), C.longlong(feature))
	}
}

//export callbackQWebPage_FocusedElementChanged
func callbackQWebPage_FocusedElementChanged(ptr unsafe.Pointer, element unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusedElementChanged"); signal != nil {
		signal.(func(*QWebElement))(NewQWebElementFromPointer(element))
	}

}

func (ptr *QWebPage) ConnectFocusedElementChanged(f func(element *QWebElement)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "focusedElementChanged") {
			C.QWebPage_ConnectFocusedElementChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "focusedElementChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "focusedElementChanged", func(element *QWebElement) {
				signal.(func(*QWebElement))(element)
				f(element)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "focusedElementChanged", f)
		}
	}
}

func (ptr *QWebPage) DisconnectFocusedElementChanged() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectFocusedElementChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "focusedElementChanged")
	}
}

func (ptr *QWebPage) FocusedElementChanged(element QWebElement_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_FocusedElementChanged(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

//export callbackQWebPage_FrameCreated
func callbackQWebPage_FrameCreated(ptr unsafe.Pointer, frame unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "frameCreated"); signal != nil {
		signal.(func(*QWebFrame))(NewQWebFrameFromPointer(frame))
	}

}

func (ptr *QWebPage) ConnectFrameCreated(f func(frame *QWebFrame)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "frameCreated") {
			C.QWebPage_ConnectFrameCreated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "frameCreated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "frameCreated", func(frame *QWebFrame) {
				signal.(func(*QWebFrame))(frame)
				f(frame)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "frameCreated", f)
		}
	}
}

func (ptr *QWebPage) DisconnectFrameCreated() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectFrameCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "frameCreated")
	}
}

func (ptr *QWebPage) FrameCreated(frame QWebFrame_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_FrameCreated(ptr.Pointer(), PointerFromQWebFrame(frame))
	}
}

//export callbackQWebPage_GeometryChangeRequested
func callbackQWebPage_GeometryChangeRequested(ptr unsafe.Pointer, geom unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "geometryChangeRequested"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(geom))
	}

}

func (ptr *QWebPage) ConnectGeometryChangeRequested(f func(geom *core.QRect)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "geometryChangeRequested") {
			C.QWebPage_ConnectGeometryChangeRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "geometryChangeRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "geometryChangeRequested", func(geom *core.QRect) {
				signal.(func(*core.QRect))(geom)
				f(geom)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "geometryChangeRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectGeometryChangeRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectGeometryChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "geometryChangeRequested")
	}
}

func (ptr *QWebPage) GeometryChangeRequested(geom core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_GeometryChangeRequested(ptr.Pointer(), core.PointerFromQRect(geom))
	}
}

//export callbackQWebPage_JavaScriptAlert
func callbackQWebPage_JavaScriptAlert(ptr unsafe.Pointer, frame unsafe.Pointer, msg C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "javaScriptAlert"); signal != nil {
		signal.(func(*QWebFrame, string))(NewQWebFrameFromPointer(frame), cGoUnpackString(msg))
	} else {
		NewQWebPageFromPointer(ptr).JavaScriptAlertDefault(NewQWebFrameFromPointer(frame), cGoUnpackString(msg))
	}
}

func (ptr *QWebPage) ConnectJavaScriptAlert(f func(frame *QWebFrame, msg string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "javaScriptAlert"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptAlert", func(frame *QWebFrame, msg string) {
				signal.(func(*QWebFrame, string))(frame, msg)
				f(frame, msg)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptAlert", f)
		}
	}
}

func (ptr *QWebPage) DisconnectJavaScriptAlert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "javaScriptAlert")
	}
}

func (ptr *QWebPage) JavaScriptAlert(frame QWebFrame_ITF, msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.QWebPage_JavaScriptAlert(ptr.Pointer(), PointerFromQWebFrame(frame), C.struct_QtWebKit_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

func (ptr *QWebPage) JavaScriptAlertDefault(frame QWebFrame_ITF, msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.QWebPage_JavaScriptAlertDefault(ptr.Pointer(), PointerFromQWebFrame(frame), C.struct_QtWebKit_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

//export callbackQWebPage_JavaScriptConsoleMessage
func callbackQWebPage_JavaScriptConsoleMessage(ptr unsafe.Pointer, message C.struct_QtWebKit_PackedString, lineNumber C.int, sourceID C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "javaScriptConsoleMessage"); signal != nil {
		signal.(func(string, int, string))(cGoUnpackString(message), int(int32(lineNumber)), cGoUnpackString(sourceID))
	} else {
		NewQWebPageFromPointer(ptr).JavaScriptConsoleMessageDefault(cGoUnpackString(message), int(int32(lineNumber)), cGoUnpackString(sourceID))
	}
}

func (ptr *QWebPage) ConnectJavaScriptConsoleMessage(f func(message string, lineNumber int, sourceID string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "javaScriptConsoleMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptConsoleMessage", func(message string, lineNumber int, sourceID string) {
				signal.(func(string, int, string))(message, lineNumber, sourceID)
				f(message, lineNumber, sourceID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "javaScriptConsoleMessage", f)
		}
	}
}

func (ptr *QWebPage) DisconnectJavaScriptConsoleMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "javaScriptConsoleMessage")
	}
}

func (ptr *QWebPage) JavaScriptConsoleMessage(message string, lineNumber int, sourceID string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		var sourceIDC *C.char
		if sourceID != "" {
			sourceIDC = C.CString(sourceID)
			defer C.free(unsafe.Pointer(sourceIDC))
		}
		C.QWebPage_JavaScriptConsoleMessage(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: messageC, len: C.longlong(len(message))}, C.int(int32(lineNumber)), C.struct_QtWebKit_PackedString{data: sourceIDC, len: C.longlong(len(sourceID))})
	}
}

func (ptr *QWebPage) JavaScriptConsoleMessageDefault(message string, lineNumber int, sourceID string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		var sourceIDC *C.char
		if sourceID != "" {
			sourceIDC = C.CString(sourceID)
			defer C.free(unsafe.Pointer(sourceIDC))
		}
		C.QWebPage_JavaScriptConsoleMessageDefault(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: messageC, len: C.longlong(len(message))}, C.int(int32(lineNumber)), C.struct_QtWebKit_PackedString{data: sourceIDC, len: C.longlong(len(sourceID))})
	}
}

//export callbackQWebPage_LinkClicked
func callbackQWebPage_LinkClicked(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "linkClicked"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebPage) ConnectLinkClicked(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linkClicked") {
			C.QWebPage_ConnectLinkClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linkClicked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "linkClicked", func(url *core.QUrl) {
				signal.(func(*core.QUrl))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linkClicked", f)
		}
	}
}

func (ptr *QWebPage) DisconnectLinkClicked() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectLinkClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "linkClicked")
	}
}

func (ptr *QWebPage) LinkClicked(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_LinkClicked(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebPage_LinkHovered
func callbackQWebPage_LinkHovered(ptr unsafe.Pointer, link C.struct_QtWebKit_PackedString, title C.struct_QtWebKit_PackedString, textContent C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "linkHovered"); signal != nil {
		signal.(func(string, string, string))(cGoUnpackString(link), cGoUnpackString(title), cGoUnpackString(textContent))
	}

}

func (ptr *QWebPage) ConnectLinkHovered(f func(link string, title string, textContent string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linkHovered") {
			C.QWebPage_ConnectLinkHovered(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linkHovered"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "linkHovered", func(link string, title string, textContent string) {
				signal.(func(string, string, string))(link, title, textContent)
				f(link, title, textContent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linkHovered", f)
		}
	}
}

func (ptr *QWebPage) DisconnectLinkHovered() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectLinkHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "linkHovered")
	}
}

func (ptr *QWebPage) LinkHovered(link string, title string, textContent string) {
	if ptr.Pointer() != nil {
		var linkC *C.char
		if link != "" {
			linkC = C.CString(link)
			defer C.free(unsafe.Pointer(linkC))
		}
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		var textContentC *C.char
		if textContent != "" {
			textContentC = C.CString(textContent)
			defer C.free(unsafe.Pointer(textContentC))
		}
		C.QWebPage_LinkHovered(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: linkC, len: C.longlong(len(link))}, C.struct_QtWebKit_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWebKit_PackedString{data: textContentC, len: C.longlong(len(textContent))})
	}
}

//export callbackQWebPage_LoadFinished
func callbackQWebPage_LoadFinished(ptr unsafe.Pointer, ok C.char) {
	if signal := qt.GetSignal(ptr, "loadFinished"); signal != nil {
		signal.(func(bool))(int8(ok) != 0)
	}

}

func (ptr *QWebPage) ConnectLoadFinished(f func(ok bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadFinished") {
			C.QWebPage_ConnectLoadFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", func(ok bool) {
				signal.(func(bool))(ok)
				f(ok)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", f)
		}
	}
}

func (ptr *QWebPage) DisconnectLoadFinished() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadFinished")
	}
}

func (ptr *QWebPage) LoadFinished(ok bool) {
	if ptr.Pointer() != nil {
		C.QWebPage_LoadFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))
	}
}

//export callbackQWebPage_LoadProgress
func callbackQWebPage_LoadProgress(ptr unsafe.Pointer, progress C.int) {
	if signal := qt.GetSignal(ptr, "loadProgress"); signal != nil {
		signal.(func(int))(int(int32(progress)))
	}

}

func (ptr *QWebPage) ConnectLoadProgress(f func(progress int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadProgress") {
			C.QWebPage_ConnectLoadProgress(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadProgress"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadProgress", func(progress int) {
				signal.(func(int))(progress)
				f(progress)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadProgress", f)
		}
	}
}

func (ptr *QWebPage) DisconnectLoadProgress() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadProgress")
	}
}

func (ptr *QWebPage) LoadProgress(progress int) {
	if ptr.Pointer() != nil {
		C.QWebPage_LoadProgress(ptr.Pointer(), C.int(int32(progress)))
	}
}

//export callbackQWebPage_LoadStarted
func callbackQWebPage_LoadStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectLoadStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadStarted") {
			C.QWebPage_ConnectLoadStarted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadStarted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", f)
		}
	}
}

func (ptr *QWebPage) DisconnectLoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadStarted")
	}
}

func (ptr *QWebPage) LoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebPage_LoadStarted(ptr.Pointer())
	}
}

//export callbackQWebPage_MenuBarVisibilityChangeRequested
func callbackQWebPage_MenuBarVisibilityChangeRequested(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "menuBarVisibilityChangeRequested"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QWebPage) ConnectMenuBarVisibilityChangeRequested(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "menuBarVisibilityChangeRequested") {
			C.QWebPage_ConnectMenuBarVisibilityChangeRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "menuBarVisibilityChangeRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "menuBarVisibilityChangeRequested", func(visible bool) {
				signal.(func(bool))(visible)
				f(visible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "menuBarVisibilityChangeRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectMenuBarVisibilityChangeRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectMenuBarVisibilityChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "menuBarVisibilityChangeRequested")
	}
}

func (ptr *QWebPage) MenuBarVisibilityChangeRequested(visible bool) {
	if ptr.Pointer() != nil {
		C.QWebPage_MenuBarVisibilityChangeRequested(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQWebPage_MicroFocusChanged
func callbackQWebPage_MicroFocusChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "microFocusChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectMicroFocusChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "microFocusChanged") {
			C.QWebPage_ConnectMicroFocusChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "microFocusChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "microFocusChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "microFocusChanged", f)
		}
	}
}

func (ptr *QWebPage) DisconnectMicroFocusChanged() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectMicroFocusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "microFocusChanged")
	}
}

func (ptr *QWebPage) MicroFocusChanged() {
	if ptr.Pointer() != nil {
		C.QWebPage_MicroFocusChanged(ptr.Pointer())
	}
}

//export callbackQWebPage_PrintRequested
func callbackQWebPage_PrintRequested(ptr unsafe.Pointer, frame unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "printRequested"); signal != nil {
		signal.(func(*QWebFrame))(NewQWebFrameFromPointer(frame))
	}

}

func (ptr *QWebPage) ConnectPrintRequested(f func(frame *QWebFrame)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "printRequested") {
			C.QWebPage_ConnectPrintRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "printRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "printRequested", func(frame *QWebFrame) {
				signal.(func(*QWebFrame))(frame)
				f(frame)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "printRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectPrintRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectPrintRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "printRequested")
	}
}

func (ptr *QWebPage) PrintRequested(frame QWebFrame_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_PrintRequested(ptr.Pointer(), PointerFromQWebFrame(frame))
	}
}

//export callbackQWebPage_RecentlyAudibleChanged
func callbackQWebPage_RecentlyAudibleChanged(ptr unsafe.Pointer, recentlyAudible C.char) {
	if signal := qt.GetSignal(ptr, "recentlyAudibleChanged"); signal != nil {
		signal.(func(bool))(int8(recentlyAudible) != 0)
	}

}

func (ptr *QWebPage) ConnectRecentlyAudibleChanged(f func(recentlyAudible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "recentlyAudibleChanged") {
			C.QWebPage_ConnectRecentlyAudibleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "recentlyAudibleChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "recentlyAudibleChanged", func(recentlyAudible bool) {
				signal.(func(bool))(recentlyAudible)
				f(recentlyAudible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "recentlyAudibleChanged", f)
		}
	}
}

func (ptr *QWebPage) DisconnectRecentlyAudibleChanged() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectRecentlyAudibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "recentlyAudibleChanged")
	}
}

func (ptr *QWebPage) RecentlyAudibleChanged(recentlyAudible bool) {
	if ptr.Pointer() != nil {
		C.QWebPage_RecentlyAudibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(recentlyAudible))))
	}
}

//export callbackQWebPage_RepaintRequested
func callbackQWebPage_RepaintRequested(ptr unsafe.Pointer, dirtyRect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaintRequested"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(dirtyRect))
	}

}

func (ptr *QWebPage) ConnectRepaintRequested(f func(dirtyRect *core.QRect)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "repaintRequested") {
			C.QWebPage_ConnectRepaintRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "repaintRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "repaintRequested", func(dirtyRect *core.QRect) {
				signal.(func(*core.QRect))(dirtyRect)
				f(dirtyRect)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "repaintRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectRepaintRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectRepaintRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "repaintRequested")
	}
}

func (ptr *QWebPage) RepaintRequested(dirtyRect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_RepaintRequested(ptr.Pointer(), core.PointerFromQRect(dirtyRect))
	}
}

//export callbackQWebPage_RestoreFrameStateRequested
func callbackQWebPage_RestoreFrameStateRequested(ptr unsafe.Pointer, frame unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "restoreFrameStateRequested"); signal != nil {
		signal.(func(*QWebFrame))(NewQWebFrameFromPointer(frame))
	}

}

func (ptr *QWebPage) ConnectRestoreFrameStateRequested(f func(frame *QWebFrame)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "restoreFrameStateRequested") {
			C.QWebPage_ConnectRestoreFrameStateRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "restoreFrameStateRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "restoreFrameStateRequested", func(frame *QWebFrame) {
				signal.(func(*QWebFrame))(frame)
				f(frame)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "restoreFrameStateRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectRestoreFrameStateRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectRestoreFrameStateRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "restoreFrameStateRequested")
	}
}

func (ptr *QWebPage) RestoreFrameStateRequested(frame QWebFrame_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_RestoreFrameStateRequested(ptr.Pointer(), PointerFromQWebFrame(frame))
	}
}

//export callbackQWebPage_SaveFrameStateRequested
func callbackQWebPage_SaveFrameStateRequested(ptr unsafe.Pointer, frame unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "saveFrameStateRequested"); signal != nil {
		signal.(func(*QWebFrame, *QWebHistoryItem))(NewQWebFrameFromPointer(frame), NewQWebHistoryItemFromPointer(item))
	}

}

func (ptr *QWebPage) ConnectSaveFrameStateRequested(f func(frame *QWebFrame, item *QWebHistoryItem)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "saveFrameStateRequested") {
			C.QWebPage_ConnectSaveFrameStateRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "saveFrameStateRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "saveFrameStateRequested", func(frame *QWebFrame, item *QWebHistoryItem) {
				signal.(func(*QWebFrame, *QWebHistoryItem))(frame, item)
				f(frame, item)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "saveFrameStateRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectSaveFrameStateRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectSaveFrameStateRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "saveFrameStateRequested")
	}
}

func (ptr *QWebPage) SaveFrameStateRequested(frame QWebFrame_ITF, item QWebHistoryItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_SaveFrameStateRequested(ptr.Pointer(), PointerFromQWebFrame(frame), PointerFromQWebHistoryItem(item))
	}
}

//export callbackQWebPage_ScrollRequested
func callbackQWebPage_ScrollRequested(ptr unsafe.Pointer, dx C.int, dy C.int, rectToScroll unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scrollRequested"); signal != nil {
		signal.(func(int, int, *core.QRect))(int(int32(dx)), int(int32(dy)), core.NewQRectFromPointer(rectToScroll))
	}

}

func (ptr *QWebPage) ConnectScrollRequested(f func(dx int, dy int, rectToScroll *core.QRect)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "scrollRequested") {
			C.QWebPage_ConnectScrollRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "scrollRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "scrollRequested", func(dx int, dy int, rectToScroll *core.QRect) {
				signal.(func(int, int, *core.QRect))(dx, dy, rectToScroll)
				f(dx, dy, rectToScroll)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scrollRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectScrollRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectScrollRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "scrollRequested")
	}
}

func (ptr *QWebPage) ScrollRequested(dx int, dy int, rectToScroll core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_ScrollRequested(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)), core.PointerFromQRect(rectToScroll))
	}
}

//export callbackQWebPage_SelectionChanged
func callbackQWebPage_SelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionChanged") {
			C.QWebPage_ConnectSelectionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", f)
		}
	}
}

func (ptr *QWebPage) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionChanged")
	}
}

func (ptr *QWebPage) SelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebPage_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QWebPage) SetContentEditable(editable bool) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetContentEditable(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(editable))))
	}
}

func (ptr *QWebPage) SetFeaturePermission(frame QWebFrame_ITF, feature QWebPage__Feature, policy QWebPage__PermissionPolicy) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetFeaturePermission(ptr.Pointer(), PointerFromQWebFrame(frame), C.longlong(feature), C.longlong(policy))
	}
}

func (ptr *QWebPage) SetForwardUnsupportedContent(forward bool) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetForwardUnsupportedContent(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(forward))))
	}
}

func (ptr *QWebPage) SetLinkDelegationPolicy(policy QWebPage__LinkDelegationPolicy) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetLinkDelegationPolicy(ptr.Pointer(), C.longlong(policy))
	}
}

func (ptr *QWebPage) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QWebPage) SetPalette(palette gui.QPalette_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetPalette(ptr.Pointer(), gui.PointerFromQPalette(palette))
	}
}

func (ptr *QWebPage) SetPluginFactory(factory QWebPluginFactory_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetPluginFactory(ptr.Pointer(), PointerFromQWebPluginFactory(factory))
	}
}

func (ptr *QWebPage) SetView(view widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetView(ptr.Pointer(), widgets.PointerFromQWidget(view))
	}
}

func (ptr *QWebPage) SetVisibilityState(vvi QWebPage__VisibilityState) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetVisibilityState(ptr.Pointer(), C.longlong(vvi))
	}
}

//export callbackQWebPage_StatusBarMessage
func callbackQWebPage_StatusBarMessage(ptr unsafe.Pointer, text C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "statusBarMessage"); signal != nil {
		signal.(func(string))(cGoUnpackString(text))
	}

}

func (ptr *QWebPage) ConnectStatusBarMessage(f func(text string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusBarMessage") {
			C.QWebPage_ConnectStatusBarMessage(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "statusBarMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "statusBarMessage", func(text string) {
				signal.(func(string))(text)
				f(text)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusBarMessage", f)
		}
	}
}

func (ptr *QWebPage) DisconnectStatusBarMessage() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectStatusBarMessage(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusBarMessage")
	}
}

func (ptr *QWebPage) StatusBarMessage(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QWebPage_StatusBarMessage(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQWebPage_StatusBarVisibilityChangeRequested
func callbackQWebPage_StatusBarVisibilityChangeRequested(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "statusBarVisibilityChangeRequested"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QWebPage) ConnectStatusBarVisibilityChangeRequested(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusBarVisibilityChangeRequested") {
			C.QWebPage_ConnectStatusBarVisibilityChangeRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "statusBarVisibilityChangeRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "statusBarVisibilityChangeRequested", func(visible bool) {
				signal.(func(bool))(visible)
				f(visible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusBarVisibilityChangeRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectStatusBarVisibilityChangeRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectStatusBarVisibilityChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusBarVisibilityChangeRequested")
	}
}

func (ptr *QWebPage) StatusBarVisibilityChangeRequested(visible bool) {
	if ptr.Pointer() != nil {
		C.QWebPage_StatusBarVisibilityChangeRequested(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQWebPage_ToolBarVisibilityChangeRequested
func callbackQWebPage_ToolBarVisibilityChangeRequested(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "toolBarVisibilityChangeRequested"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	}

}

func (ptr *QWebPage) ConnectToolBarVisibilityChangeRequested(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "toolBarVisibilityChangeRequested") {
			C.QWebPage_ConnectToolBarVisibilityChangeRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "toolBarVisibilityChangeRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "toolBarVisibilityChangeRequested", func(visible bool) {
				signal.(func(bool))(visible)
				f(visible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "toolBarVisibilityChangeRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectToolBarVisibilityChangeRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectToolBarVisibilityChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "toolBarVisibilityChangeRequested")
	}
}

func (ptr *QWebPage) ToolBarVisibilityChangeRequested(visible bool) {
	if ptr.Pointer() != nil {
		C.QWebPage_ToolBarVisibilityChangeRequested(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQWebPage_TriggerAction
func callbackQWebPage_TriggerAction(ptr unsafe.Pointer, action C.longlong, checked C.char) {
	if signal := qt.GetSignal(ptr, "triggerAction"); signal != nil {
		signal.(func(QWebPage__WebAction, bool))(QWebPage__WebAction(action), int8(checked) != 0)
	} else {
		NewQWebPageFromPointer(ptr).TriggerActionDefault(QWebPage__WebAction(action), int8(checked) != 0)
	}
}

func (ptr *QWebPage) ConnectTriggerAction(f func(action QWebPage__WebAction, checked bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "triggerAction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "triggerAction", func(action QWebPage__WebAction, checked bool) {
				signal.(func(QWebPage__WebAction, bool))(action, checked)
				f(action, checked)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "triggerAction", f)
		}
	}
}

func (ptr *QWebPage) DisconnectTriggerAction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "triggerAction")
	}
}

func (ptr *QWebPage) TriggerAction(action QWebPage__WebAction, checked bool) {
	if ptr.Pointer() != nil {
		C.QWebPage_TriggerAction(ptr.Pointer(), C.longlong(action), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

func (ptr *QWebPage) TriggerActionDefault(action QWebPage__WebAction, checked bool) {
	if ptr.Pointer() != nil {
		C.QWebPage_TriggerActionDefault(ptr.Pointer(), C.longlong(action), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

//export callbackQWebPage_UnsupportedContent
func callbackQWebPage_UnsupportedContent(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "unsupportedContent"); signal != nil {
		signal.(func(*network.QNetworkReply))(network.NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QWebPage) ConnectUnsupportedContent(f func(reply *network.QNetworkReply)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "unsupportedContent") {
			C.QWebPage_ConnectUnsupportedContent(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "unsupportedContent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "unsupportedContent", func(reply *network.QNetworkReply) {
				signal.(func(*network.QNetworkReply))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "unsupportedContent", f)
		}
	}
}

func (ptr *QWebPage) DisconnectUnsupportedContent() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectUnsupportedContent(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "unsupportedContent")
	}
}

func (ptr *QWebPage) UnsupportedContent(reply network.QNetworkReply_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_UnsupportedContent(ptr.Pointer(), network.PointerFromQNetworkReply(reply))
	}
}

func (ptr *QWebPage) UpdatePositionDependentActions(pos core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_UpdatePositionDependentActions(ptr.Pointer(), core.PointerFromQPoint(pos))
	}
}

//export callbackQWebPage_ViewportChangeRequested
func callbackQWebPage_ViewportChangeRequested(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "viewportChangeRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectViewportChangeRequested(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "viewportChangeRequested") {
			C.QWebPage_ConnectViewportChangeRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "viewportChangeRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "viewportChangeRequested", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "viewportChangeRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectViewportChangeRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectViewportChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "viewportChangeRequested")
	}
}

func (ptr *QWebPage) ViewportChangeRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_ViewportChangeRequested(ptr.Pointer())
	}
}

//export callbackQWebPage_WindowCloseRequested
func callbackQWebPage_WindowCloseRequested(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowCloseRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectWindowCloseRequested(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "windowCloseRequested") {
			C.QWebPage_ConnectWindowCloseRequested(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "windowCloseRequested"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "windowCloseRequested", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "windowCloseRequested", f)
		}
	}
}

func (ptr *QWebPage) DisconnectWindowCloseRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectWindowCloseRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "windowCloseRequested")
	}
}

func (ptr *QWebPage) WindowCloseRequested() {
	if ptr.Pointer() != nil {
		C.QWebPage_WindowCloseRequested(ptr.Pointer())
	}
}

func (ptr *QWebPage) DestroyQWebPage() {
	if ptr.Pointer() != nil {
		C.QWebPage_DestroyQWebPage(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebPage) LinkDelegationPolicy() QWebPage__LinkDelegationPolicy {
	if ptr.Pointer() != nil {
		return QWebPage__LinkDelegationPolicy(C.QWebPage_LinkDelegationPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebPage) Action(action QWebPage__WebAction) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebPage_Action(ptr.Pointer(), C.longlong(action)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) CustomAction(action int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebPage_CustomAction(ptr.Pointer(), C.int(int32(action))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkAccessManagerFromPointer(C.QWebPage_NetworkAccessManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) Palette() *gui.QPalette {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPaletteFromPointer(C.QWebPage_Palette(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPalette).DestroyQPalette)
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) PreferredContentsSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebPage_PreferredContentsSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) ViewportSize() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebPage_ViewportSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) SelectedHtml() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebPage_SelectedHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebPage) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebPage_SelectedText(ptr.Pointer()))
	}
	return ""
}

//export callbackQWebPage_UserAgentForUrl
func callbackQWebPage_UserAgentForUrl(ptr unsafe.Pointer, url unsafe.Pointer) C.struct_QtWebKit_PackedString {
	if signal := qt.GetSignal(ptr, "userAgentForUrl"); signal != nil {
		tempVal := signal.(func(*core.QUrl) string)(core.NewQUrlFromPointer(url))
		return C.struct_QtWebKit_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQWebPageFromPointer(ptr).UserAgentForUrlDefault(core.NewQUrlFromPointer(url))
	return C.struct_QtWebKit_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QWebPage) ConnectUserAgentForUrl(f func(url *core.QUrl) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "userAgentForUrl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "userAgentForUrl", func(url *core.QUrl) string {
				signal.(func(*core.QUrl) string)(url)
				return f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "userAgentForUrl", f)
		}
	}
}

func (ptr *QWebPage) DisconnectUserAgentForUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "userAgentForUrl")
	}
}

func (ptr *QWebPage) UserAgentForUrl(url core.QUrl_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebPage_UserAgentForUrl(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return ""
}

func (ptr *QWebPage) UserAgentForUrlDefault(url core.QUrl_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebPage_UserAgentForUrlDefault(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return ""
}

func (ptr *QWebPage) SupportedContentTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QWebPage_SupportedContentTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QWebPage) UndoStack() *widgets.QUndoStack {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQUndoStackFromPointer(C.QWebPage_UndoStack(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QWebPage_InputMethodQuery(ptr.Pointer(), C.longlong(property)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) CurrentFrame() *QWebFrame {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebFrameFromPointer(C.QWebPage_CurrentFrame(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) FrameAt(pos core.QPoint_ITF) *QWebFrame {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebFrameFromPointer(C.QWebPage_FrameAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) MainFrame() *QWebFrame {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebFrameFromPointer(C.QWebPage_MainFrame(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) History() *QWebHistory {
	if ptr.Pointer() != nil {
		return NewQWebHistoryFromPointer(C.QWebPage_History(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebPage) PluginFactory() *QWebPluginFactory {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebPluginFactoryFromPointer(C.QWebPage_PluginFactory(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) Settings() *QWebSettings {
	if ptr.Pointer() != nil {
		return NewQWebSettingsFromPointer(C.QWebPage_Settings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebPage) View() *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QWebPage_View(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) VisibilityState() QWebPage__VisibilityState {
	if ptr.Pointer() != nil {
		return QWebPage__VisibilityState(C.QWebPage_VisibilityState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebPage) ForwardUnsupportedContent() bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_ForwardUnsupportedContent(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) HasSelection() bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) IsContentEditable() bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_IsContentEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) IsModified() bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) RecentlyAudible() bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_RecentlyAudible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) SupportsContentType(mimeType string) bool {
	if ptr.Pointer() != nil {
		var mimeTypeC *C.char
		if mimeType != "" {
			mimeTypeC = C.CString(mimeType)
			defer C.free(unsafe.Pointer(mimeTypeC))
		}
		return C.QWebPage_SupportsContentType(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: mimeTypeC, len: C.longlong(len(mimeType))}) != 0
	}
	return false
}

//export callbackQWebPage_SupportsExtension
func callbackQWebPage_SupportsExtension(ptr unsafe.Pointer, extension C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "supportsExtension"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(QWebPage__Extension) bool)(QWebPage__Extension(extension)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).SupportsExtensionDefault(QWebPage__Extension(extension)))))
}

func (ptr *QWebPage) ConnectSupportsExtension(f func(extension QWebPage__Extension) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportsExtension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "supportsExtension", func(extension QWebPage__Extension) bool {
				signal.(func(QWebPage__Extension) bool)(extension)
				return f(extension)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportsExtension", f)
		}
	}
}

func (ptr *QWebPage) DisconnectSupportsExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "supportsExtension")
	}
}

func (ptr *QWebPage) SupportsExtension(extension QWebPage__Extension) bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_SupportsExtension(ptr.Pointer(), C.longlong(extension)) != 0
	}
	return false
}

func (ptr *QWebPage) SupportsExtensionDefault(extension QWebPage__Extension) bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_SupportsExtensionDefault(ptr.Pointer(), C.longlong(extension)) != 0
	}
	return false
}

func (ptr *QWebPage) BytesReceived() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QWebPage_BytesReceived(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebPage) TotalBytes() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QWebPage_TotalBytes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebPage) SetActualVisibleContentRect(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetActualVisibleContentRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWebPage) SetPreferredContentsSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetPreferredContentsSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWebPage) SetViewportSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_SetViewportSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWebPage) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QWebPage___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebPage) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebPage___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QWebPage) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebPage___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebPage) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QWebPage___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QWebPage) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebPage___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebPage) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QWebPage___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QWebPage) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebPage___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebPage) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebPage___findChildren_newList(ptr.Pointer()))
}

func (ptr *QWebPage) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebPage___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPage) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebPage) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebPage___children_newList(ptr.Pointer()))
}

//export callbackQWebPage_EventFilter
func callbackQWebPage_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebPage) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebPage_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebPage_ChildEvent
func callbackQWebPage_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebPageFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebPage) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebPage_ConnectNotify
func callbackQWebPage_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebPageFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebPage) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebPage_CustomEvent
func callbackQWebPage_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebPageFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebPage) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebPage_DeleteLater
func callbackQWebPage_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebPageFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebPage) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebPage_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebPage_Destroyed
func callbackQWebPage_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebPage_DisconnectNotify
func callbackQWebPage_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebPageFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebPage) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebPage_ObjectNameChanged
func callbackQWebPage_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQWebPage_TimerEvent
func callbackQWebPage_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebPageFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebPage) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPage_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebPage_MetaObject
func callbackQWebPage_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebPageFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebPage) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebPage_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebPluginFactory struct {
	core.QObject
}

type QWebPluginFactory_ITF interface {
	core.QObject_ITF
	QWebPluginFactory_PTR() *QWebPluginFactory
}

func (ptr *QWebPluginFactory) QWebPluginFactory_PTR() *QWebPluginFactory {
	return ptr
}

func (ptr *QWebPluginFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebPluginFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebPluginFactory(ptr QWebPluginFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebPluginFactory_PTR().Pointer()
	}
	return nil
}

func NewQWebPluginFactoryFromPointer(ptr unsafe.Pointer) *QWebPluginFactory {
	var n = new(QWebPluginFactory)
	n.SetPointer(ptr)
	return n
}

//export callbackQWebPluginFactory_RefreshPlugins
func callbackQWebPluginFactory_RefreshPlugins(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "refreshPlugins"); signal != nil {
		signal.(func())()
	} else {
		NewQWebPluginFactoryFromPointer(ptr).RefreshPluginsDefault()
	}
}

func (ptr *QWebPluginFactory) ConnectRefreshPlugins(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "refreshPlugins"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "refreshPlugins", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "refreshPlugins", f)
		}
	}
}

func (ptr *QWebPluginFactory) DisconnectRefreshPlugins() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "refreshPlugins")
	}
}

func (ptr *QWebPluginFactory) RefreshPlugins() {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory_RefreshPlugins(ptr.Pointer())
	}
}

func (ptr *QWebPluginFactory) RefreshPluginsDefault() {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory_RefreshPluginsDefault(ptr.Pointer())
	}
}

//export callbackQWebPluginFactory_DestroyQWebPluginFactory
func callbackQWebPluginFactory_DestroyQWebPluginFactory(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebPluginFactory"); signal != nil {
		signal.(func())()
	} else {
		NewQWebPluginFactoryFromPointer(ptr).DestroyQWebPluginFactoryDefault()
	}
}

func (ptr *QWebPluginFactory) ConnectDestroyQWebPluginFactory(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebPluginFactory"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QWebPluginFactory", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebPluginFactory", f)
		}
	}
}

func (ptr *QWebPluginFactory) DisconnectDestroyQWebPluginFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWebPluginFactory")
	}
}

func (ptr *QWebPluginFactory) DestroyQWebPluginFactory() {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory_DestroyQWebPluginFactory(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebPluginFactory) DestroyQWebPluginFactoryDefault() {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory_DestroyQWebPluginFactoryDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebPluginFactory_Create
func callbackQWebPluginFactory_Create(ptr unsafe.Pointer, mimeType C.struct_QtWebKit_PackedString, url unsafe.Pointer, argumentNames C.struct_QtWebKit_PackedString, argumentValues C.struct_QtWebKit_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "create"); signal != nil {
		return core.PointerFromQObject(signal.(func(string, *core.QUrl, []string, []string) *core.QObject)(cGoUnpackString(mimeType), core.NewQUrlFromPointer(url), strings.Split(cGoUnpackString(argumentNames), "|"), strings.Split(cGoUnpackString(argumentValues), "|")))
	}

	return core.PointerFromQObject(core.NewQObject(nil))
}

func (ptr *QWebPluginFactory) ConnectCreate(f func(mimeType string, url *core.QUrl, argumentNames []string, argumentValues []string) *core.QObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "create"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "create", func(mimeType string, url *core.QUrl, argumentNames []string, argumentValues []string) *core.QObject {
				signal.(func(string, *core.QUrl, []string, []string) *core.QObject)(mimeType, url, argumentNames, argumentValues)
				return f(mimeType, url, argumentNames, argumentValues)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "create", f)
		}
	}
}

func (ptr *QWebPluginFactory) DisconnectCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "create")
	}
}

func (ptr *QWebPluginFactory) Create(mimeType string, url core.QUrl_ITF, argumentNames []string, argumentValues []string) *core.QObject {
	if ptr.Pointer() != nil {
		var mimeTypeC *C.char
		if mimeType != "" {
			mimeTypeC = C.CString(mimeType)
			defer C.free(unsafe.Pointer(mimeTypeC))
		}
		var argumentNamesC = C.CString(strings.Join(argumentNames, "|"))
		defer C.free(unsafe.Pointer(argumentNamesC))
		var argumentValuesC = C.CString(strings.Join(argumentValues, "|"))
		defer C.free(unsafe.Pointer(argumentValuesC))
		var tmpValue = core.NewQObjectFromPointer(C.QWebPluginFactory_Create(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: mimeTypeC, len: C.longlong(len(mimeType))}, core.PointerFromQUrl(url), C.struct_QtWebKit_PackedString{data: argumentNamesC, len: C.longlong(len(strings.Join(argumentNames, "|")))}, C.struct_QtWebKit_PackedString{data: argumentValuesC, len: C.longlong(len(strings.Join(argumentValues, "|")))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPluginFactory) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QWebPluginFactory___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebPluginFactory) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebPluginFactory) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebPluginFactory___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QWebPluginFactory) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebPluginFactory___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPluginFactory) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebPluginFactory) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QWebPluginFactory___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QWebPluginFactory) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebPluginFactory___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPluginFactory) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebPluginFactory) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QWebPluginFactory___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QWebPluginFactory) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebPluginFactory___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPluginFactory) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebPluginFactory) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebPluginFactory___findChildren_newList(ptr.Pointer()))
}

func (ptr *QWebPluginFactory) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebPluginFactory___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebPluginFactory) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebPluginFactory) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebPluginFactory___children_newList(ptr.Pointer()))
}

//export callbackQWebPluginFactory_Event
func callbackQWebPluginFactory_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebPluginFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebPluginFactory) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebPluginFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebPluginFactory_EventFilter
func callbackQWebPluginFactory_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebPluginFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebPluginFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebPluginFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebPluginFactory_ChildEvent
func callbackQWebPluginFactory_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebPluginFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebPluginFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebPluginFactory_ConnectNotify
func callbackQWebPluginFactory_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebPluginFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebPluginFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebPluginFactory_CustomEvent
func callbackQWebPluginFactory_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebPluginFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebPluginFactory) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebPluginFactory_DeleteLater
func callbackQWebPluginFactory_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebPluginFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebPluginFactory) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebPluginFactory_Destroyed
func callbackQWebPluginFactory_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebPluginFactory_DisconnectNotify
func callbackQWebPluginFactory_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebPluginFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebPluginFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebPluginFactory_ObjectNameChanged
func callbackQWebPluginFactory_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQWebPluginFactory_TimerEvent
func callbackQWebPluginFactory_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebPluginFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebPluginFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebPluginFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebPluginFactory_MetaObject
func callbackQWebPluginFactory_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebPluginFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebPluginFactory) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebPluginFactory_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebSecurityOrigin struct {
	ptr unsafe.Pointer
}

type QWebSecurityOrigin_ITF interface {
	QWebSecurityOrigin_PTR() *QWebSecurityOrigin
}

func (ptr *QWebSecurityOrigin) QWebSecurityOrigin_PTR() *QWebSecurityOrigin {
	return ptr
}

func (ptr *QWebSecurityOrigin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebSecurityOrigin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebSecurityOrigin(ptr QWebSecurityOrigin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSecurityOrigin_PTR().Pointer()
	}
	return nil
}

func NewQWebSecurityOriginFromPointer(ptr unsafe.Pointer) *QWebSecurityOrigin {
	var n = new(QWebSecurityOrigin)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QWebSecurityOrigin__SubdomainSetting
//QWebSecurityOrigin::SubdomainSetting
type QWebSecurityOrigin__SubdomainSetting int64

const (
	QWebSecurityOrigin__AllowSubdomains    QWebSecurityOrigin__SubdomainSetting = QWebSecurityOrigin__SubdomainSetting(0)
	QWebSecurityOrigin__DisallowSubdomains QWebSecurityOrigin__SubdomainSetting = QWebSecurityOrigin__SubdomainSetting(1)
)

func QWebSecurityOrigin_AllOrigins() []*QWebSecurityOrigin {
	return func(l C.struct_QtWebKit_PackedList) []*QWebSecurityOrigin {
		var out = make([]*QWebSecurityOrigin, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQWebSecurityOriginFromPointer(l.data).__allOrigins_atList(i)
		}
		return out
	}(C.QWebSecurityOrigin_QWebSecurityOrigin_AllOrigins())
}

func (ptr *QWebSecurityOrigin) AllOrigins() []*QWebSecurityOrigin {
	return func(l C.struct_QtWebKit_PackedList) []*QWebSecurityOrigin {
		var out = make([]*QWebSecurityOrigin, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQWebSecurityOriginFromPointer(l.data).__allOrigins_atList(i)
		}
		return out
	}(C.QWebSecurityOrigin_QWebSecurityOrigin_AllOrigins())
}

func QWebSecurityOrigin_LocalSchemes() []string {
	return strings.Split(cGoUnpackString(C.QWebSecurityOrigin_QWebSecurityOrigin_LocalSchemes()), "|")
}

func (ptr *QWebSecurityOrigin) LocalSchemes() []string {
	return strings.Split(cGoUnpackString(C.QWebSecurityOrigin_QWebSecurityOrigin_LocalSchemes()), "|")
}

func NewQWebSecurityOrigin(url core.QUrl_ITF) *QWebSecurityOrigin {
	var tmpValue = NewQWebSecurityOriginFromPointer(C.QWebSecurityOrigin_NewQWebSecurityOrigin(core.PointerFromQUrl(url)))
	runtime.SetFinalizer(tmpValue, (*QWebSecurityOrigin).DestroyQWebSecurityOrigin)
	return tmpValue
}

func NewQWebSecurityOrigin2(other QWebSecurityOrigin_ITF) *QWebSecurityOrigin {
	var tmpValue = NewQWebSecurityOriginFromPointer(C.QWebSecurityOrigin_NewQWebSecurityOrigin2(PointerFromQWebSecurityOrigin(other)))
	runtime.SetFinalizer(tmpValue, (*QWebSecurityOrigin).DestroyQWebSecurityOrigin)
	return tmpValue
}

func (ptr *QWebSecurityOrigin) AddAccessWhitelistEntry(scheme string, host string, subdomainSetting QWebSecurityOrigin__SubdomainSetting) {
	if ptr.Pointer() != nil {
		var schemeC *C.char
		if scheme != "" {
			schemeC = C.CString(scheme)
			defer C.free(unsafe.Pointer(schemeC))
		}
		var hostC *C.char
		if host != "" {
			hostC = C.CString(host)
			defer C.free(unsafe.Pointer(hostC))
		}
		C.QWebSecurityOrigin_AddAccessWhitelistEntry(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: schemeC, len: C.longlong(len(scheme))}, C.struct_QtWebKit_PackedString{data: hostC, len: C.longlong(len(host))}, C.longlong(subdomainSetting))
	}
}

func QWebSecurityOrigin_AddLocalScheme(scheme string) {
	var schemeC *C.char
	if scheme != "" {
		schemeC = C.CString(scheme)
		defer C.free(unsafe.Pointer(schemeC))
	}
	C.QWebSecurityOrigin_QWebSecurityOrigin_AddLocalScheme(C.struct_QtWebKit_PackedString{data: schemeC, len: C.longlong(len(scheme))})
}

func (ptr *QWebSecurityOrigin) AddLocalScheme(scheme string) {
	var schemeC *C.char
	if scheme != "" {
		schemeC = C.CString(scheme)
		defer C.free(unsafe.Pointer(schemeC))
	}
	C.QWebSecurityOrigin_QWebSecurityOrigin_AddLocalScheme(C.struct_QtWebKit_PackedString{data: schemeC, len: C.longlong(len(scheme))})
}

func (ptr *QWebSecurityOrigin) RemoveAccessWhitelistEntry(scheme string, host string, subdomainSetting QWebSecurityOrigin__SubdomainSetting) {
	if ptr.Pointer() != nil {
		var schemeC *C.char
		if scheme != "" {
			schemeC = C.CString(scheme)
			defer C.free(unsafe.Pointer(schemeC))
		}
		var hostC *C.char
		if host != "" {
			hostC = C.CString(host)
			defer C.free(unsafe.Pointer(hostC))
		}
		C.QWebSecurityOrigin_RemoveAccessWhitelistEntry(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: schemeC, len: C.longlong(len(scheme))}, C.struct_QtWebKit_PackedString{data: hostC, len: C.longlong(len(host))}, C.longlong(subdomainSetting))
	}
}

func QWebSecurityOrigin_RemoveLocalScheme(scheme string) {
	var schemeC *C.char
	if scheme != "" {
		schemeC = C.CString(scheme)
		defer C.free(unsafe.Pointer(schemeC))
	}
	C.QWebSecurityOrigin_QWebSecurityOrigin_RemoveLocalScheme(C.struct_QtWebKit_PackedString{data: schemeC, len: C.longlong(len(scheme))})
}

func (ptr *QWebSecurityOrigin) RemoveLocalScheme(scheme string) {
	var schemeC *C.char
	if scheme != "" {
		schemeC = C.CString(scheme)
		defer C.free(unsafe.Pointer(schemeC))
	}
	C.QWebSecurityOrigin_QWebSecurityOrigin_RemoveLocalScheme(C.struct_QtWebKit_PackedString{data: schemeC, len: C.longlong(len(scheme))})
}

func (ptr *QWebSecurityOrigin) SetApplicationCacheQuota(quota int64) {
	if ptr.Pointer() != nil {
		C.QWebSecurityOrigin_SetApplicationCacheQuota(ptr.Pointer(), C.longlong(quota))
	}
}

func (ptr *QWebSecurityOrigin) SetDatabaseQuota(quota int64) {
	if ptr.Pointer() != nil {
		C.QWebSecurityOrigin_SetDatabaseQuota(ptr.Pointer(), C.longlong(quota))
	}
}

func (ptr *QWebSecurityOrigin) DestroyQWebSecurityOrigin() {
	if ptr.Pointer() != nil {
		C.QWebSecurityOrigin_DestroyQWebSecurityOrigin(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebSecurityOrigin) Databases() []*QWebDatabase {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebKit_PackedList) []*QWebDatabase {
			var out = make([]*QWebDatabase, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebSecurityOriginFromPointer(l.data).__databases_atList(i)
			}
			return out
		}(C.QWebSecurityOrigin_Databases(ptr.Pointer()))
	}
	return make([]*QWebDatabase, 0)
}

func (ptr *QWebSecurityOrigin) Host() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSecurityOrigin_Host(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSecurityOrigin) Scheme() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSecurityOrigin_Scheme(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSecurityOrigin) Port() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebSecurityOrigin_Port(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebSecurityOrigin) DatabaseQuota() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebSecurityOrigin_DatabaseQuota(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSecurityOrigin) DatabaseUsage() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebSecurityOrigin_DatabaseUsage(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSecurityOrigin) __allOrigins_atList(i int) *QWebSecurityOrigin {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebSecurityOriginFromPointer(C.QWebSecurityOrigin___allOrigins_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebSecurityOrigin).DestroyQWebSecurityOrigin)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSecurityOrigin) __allOrigins_setList(i QWebSecurityOrigin_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSecurityOrigin___allOrigins_setList(ptr.Pointer(), PointerFromQWebSecurityOrigin(i))
	}
}

func (ptr *QWebSecurityOrigin) __allOrigins_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebSecurityOrigin___allOrigins_newList(ptr.Pointer()))
}

func (ptr *QWebSecurityOrigin) __databases_atList(i int) *QWebDatabase {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebDatabaseFromPointer(C.QWebSecurityOrigin___databases_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QWebDatabase).DestroyQWebDatabase)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSecurityOrigin) __databases_setList(i QWebDatabase_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSecurityOrigin___databases_setList(ptr.Pointer(), PointerFromQWebDatabase(i))
	}
}

func (ptr *QWebSecurityOrigin) __databases_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebSecurityOrigin___databases_newList(ptr.Pointer()))
}

type QWebSettings struct {
	ptr unsafe.Pointer
}

type QWebSettings_ITF interface {
	QWebSettings_PTR() *QWebSettings
}

func (ptr *QWebSettings) QWebSettings_PTR() *QWebSettings {
	return ptr
}

func (ptr *QWebSettings) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebSettings) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebSettings(ptr QWebSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSettings_PTR().Pointer()
	}
	return nil
}

func NewQWebSettingsFromPointer(ptr unsafe.Pointer) *QWebSettings {
	var n = new(QWebSettings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebSettings) DestroyQWebSettings() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QWebSettings__FontFamily
//QWebSettings::FontFamily
type QWebSettings__FontFamily int64

const (
	QWebSettings__StandardFont  QWebSettings__FontFamily = QWebSettings__FontFamily(0)
	QWebSettings__FixedFont     QWebSettings__FontFamily = QWebSettings__FontFamily(1)
	QWebSettings__SerifFont     QWebSettings__FontFamily = QWebSettings__FontFamily(2)
	QWebSettings__SansSerifFont QWebSettings__FontFamily = QWebSettings__FontFamily(3)
	QWebSettings__CursiveFont   QWebSettings__FontFamily = QWebSettings__FontFamily(4)
	QWebSettings__FantasyFont   QWebSettings__FontFamily = QWebSettings__FontFamily(5)
)

//go:generate stringer -type=QWebSettings__FontSize
//QWebSettings::FontSize
type QWebSettings__FontSize int64

const (
	QWebSettings__MinimumFontSize        QWebSettings__FontSize = QWebSettings__FontSize(0)
	QWebSettings__MinimumLogicalFontSize QWebSettings__FontSize = QWebSettings__FontSize(1)
	QWebSettings__DefaultFontSize        QWebSettings__FontSize = QWebSettings__FontSize(2)
	QWebSettings__DefaultFixedFontSize   QWebSettings__FontSize = QWebSettings__FontSize(3)
)

//go:generate stringer -type=QWebSettings__ThirdPartyCookiePolicy
//QWebSettings::ThirdPartyCookiePolicy
type QWebSettings__ThirdPartyCookiePolicy int64

const (
	QWebSettings__AlwaysAllowThirdPartyCookies       QWebSettings__ThirdPartyCookiePolicy = QWebSettings__ThirdPartyCookiePolicy(0)
	QWebSettings__AlwaysBlockThirdPartyCookies       QWebSettings__ThirdPartyCookiePolicy = QWebSettings__ThirdPartyCookiePolicy(1)
	QWebSettings__AllowThirdPartyWithExistingCookies QWebSettings__ThirdPartyCookiePolicy = QWebSettings__ThirdPartyCookiePolicy(2)
)

//go:generate stringer -type=QWebSettings__WebAttribute
//QWebSettings::WebAttribute
type QWebSettings__WebAttribute int64

var (
	QWebSettings__AutoLoadImages                    QWebSettings__WebAttribute = QWebSettings__WebAttribute(0)
	QWebSettings__JavascriptEnabled                 QWebSettings__WebAttribute = QWebSettings__WebAttribute(1)
	QWebSettings__JavaEnabled                       QWebSettings__WebAttribute = QWebSettings__WebAttribute(2)
	QWebSettings__PluginsEnabled                    QWebSettings__WebAttribute = QWebSettings__WebAttribute(3)
	QWebSettings__PrivateBrowsingEnabled            QWebSettings__WebAttribute = QWebSettings__WebAttribute(4)
	QWebSettings__JavascriptCanOpenWindows          QWebSettings__WebAttribute = QWebSettings__WebAttribute(5)
	QWebSettings__JavascriptCanAccessClipboard      QWebSettings__WebAttribute = QWebSettings__WebAttribute(6)
	QWebSettings__DeveloperExtrasEnabled            QWebSettings__WebAttribute = QWebSettings__WebAttribute(7)
	QWebSettings__LinksIncludedInFocusChain         QWebSettings__WebAttribute = QWebSettings__WebAttribute(8)
	QWebSettings__ZoomTextOnly                      QWebSettings__WebAttribute = QWebSettings__WebAttribute(9)
	QWebSettings__PrintElementBackgrounds           QWebSettings__WebAttribute = QWebSettings__WebAttribute(10)
	QWebSettings__OfflineStorageDatabaseEnabled     QWebSettings__WebAttribute = QWebSettings__WebAttribute(11)
	QWebSettings__OfflineWebApplicationCacheEnabled QWebSettings__WebAttribute = QWebSettings__WebAttribute(12)
	QWebSettings__LocalStorageEnabled               QWebSettings__WebAttribute = QWebSettings__WebAttribute(13)
	QWebSettings__LocalStorageDatabaseEnabled       QWebSettings__WebAttribute = QWebSettings__WebAttribute(QWebSettings__LocalStorageEnabled)
	QWebSettings__LocalContentCanAccessRemoteUrls   QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_LocalContentCanAccessRemoteUrls_Type())
	QWebSettings__DnsPrefetchEnabled                QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_DnsPrefetchEnabled_Type())
	QWebSettings__XSSAuditingEnabled                QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_XSSAuditingEnabled_Type())
	QWebSettings__AcceleratedCompositingEnabled     QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_AcceleratedCompositingEnabled_Type())
	QWebSettings__SpatialNavigationEnabled          QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_SpatialNavigationEnabled_Type())
	QWebSettings__LocalContentCanAccessFileUrls     QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_LocalContentCanAccessFileUrls_Type())
	QWebSettings__TiledBackingStoreEnabled          QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_TiledBackingStoreEnabled_Type())
	QWebSettings__FrameFlatteningEnabled            QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_FrameFlatteningEnabled_Type())
	QWebSettings__SiteSpecificQuirksEnabled         QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_SiteSpecificQuirksEnabled_Type())
	QWebSettings__JavascriptCanCloseWindows         QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_JavascriptCanCloseWindows_Type())
	QWebSettings__WebGLEnabled                      QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_WebGLEnabled_Type())
	QWebSettings__CSSRegionsEnabled                 QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_CSSRegionsEnabled_Type())
	QWebSettings__HyperlinkAuditingEnabled          QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_HyperlinkAuditingEnabled_Type())
	QWebSettings__CSSGridLayoutEnabled              QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_CSSGridLayoutEnabled_Type())
	QWebSettings__ScrollAnimatorEnabled             QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_ScrollAnimatorEnabled_Type())
	QWebSettings__CaretBrowsingEnabled              QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_CaretBrowsingEnabled_Type())
	QWebSettings__NotificationsEnabled              QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_NotificationsEnabled_Type())
	QWebSettings__WebAudioEnabled                   QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_WebAudioEnabled_Type())
	QWebSettings__Accelerated2dCanvasEnabled        QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_Accelerated2dCanvasEnabled_Type())
	QWebSettings__MediaSourceEnabled                QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_MediaSourceEnabled_Type())
	QWebSettings__MediaEnabled                      QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_MediaEnabled_Type())
	QWebSettings__WebSecurityEnabled                QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_WebSecurityEnabled_Type())
	QWebSettings__FullScreenSupportEnabled          QWebSettings__WebAttribute = QWebSettings__WebAttribute(C.QWebSettings_FullScreenSupportEnabled_Type())
)

//go:generate stringer -type=QWebSettings__WebGraphic
//QWebSettings::WebGraphic
type QWebSettings__WebGraphic int64

const (
	QWebSettings__MissingImageGraphic              QWebSettings__WebGraphic = QWebSettings__WebGraphic(0)
	QWebSettings__MissingPluginGraphic             QWebSettings__WebGraphic = QWebSettings__WebGraphic(1)
	QWebSettings__DefaultFrameIconGraphic          QWebSettings__WebGraphic = QWebSettings__WebGraphic(2)
	QWebSettings__TextAreaSizeGripCornerGraphic    QWebSettings__WebGraphic = QWebSettings__WebGraphic(3)
	QWebSettings__DeleteButtonGraphic              QWebSettings__WebGraphic = QWebSettings__WebGraphic(4)
	QWebSettings__InputSpeechButtonGraphic         QWebSettings__WebGraphic = QWebSettings__WebGraphic(5)
	QWebSettings__SearchCancelButtonGraphic        QWebSettings__WebGraphic = QWebSettings__WebGraphic(6)
	QWebSettings__SearchCancelButtonPressedGraphic QWebSettings__WebGraphic = QWebSettings__WebGraphic(7)
)

func QWebSettings_IconForUrl(url core.QUrl_ITF) *gui.QIcon {
	var tmpValue = gui.NewQIconFromPointer(C.QWebSettings_QWebSettings_IconForUrl(core.PointerFromQUrl(url)))
	runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
	return tmpValue
}

func (ptr *QWebSettings) IconForUrl(url core.QUrl_ITF) *gui.QIcon {
	var tmpValue = gui.NewQIconFromPointer(C.QWebSettings_QWebSettings_IconForUrl(core.PointerFromQUrl(url)))
	runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
	return tmpValue
}

func QWebSettings_WebGraphic(ty QWebSettings__WebGraphic) *gui.QPixmap {
	var tmpValue = gui.NewQPixmapFromPointer(C.QWebSettings_QWebSettings_WebGraphic(C.longlong(ty)))
	runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
	return tmpValue
}

func (ptr *QWebSettings) WebGraphic(ty QWebSettings__WebGraphic) *gui.QPixmap {
	var tmpValue = gui.NewQPixmapFromPointer(C.QWebSettings_QWebSettings_WebGraphic(C.longlong(ty)))
	runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
	return tmpValue
}

func QWebSettings_IconDatabasePath() string {
	return cGoUnpackString(C.QWebSettings_QWebSettings_IconDatabasePath())
}

func (ptr *QWebSettings) IconDatabasePath() string {
	return cGoUnpackString(C.QWebSettings_QWebSettings_IconDatabasePath())
}

func QWebSettings_OfflineStoragePath() string {
	return cGoUnpackString(C.QWebSettings_QWebSettings_OfflineStoragePath())
}

func (ptr *QWebSettings) OfflineStoragePath() string {
	return cGoUnpackString(C.QWebSettings_QWebSettings_OfflineStoragePath())
}

func QWebSettings_OfflineWebApplicationCachePath() string {
	return cGoUnpackString(C.QWebSettings_QWebSettings_OfflineWebApplicationCachePath())
}

func (ptr *QWebSettings) OfflineWebApplicationCachePath() string {
	return cGoUnpackString(C.QWebSettings_QWebSettings_OfflineWebApplicationCachePath())
}

func QWebSettings_PluginSearchPaths() []string {
	return strings.Split(cGoUnpackString(C.QWebSettings_QWebSettings_PluginSearchPaths()), "|")
}

func (ptr *QWebSettings) PluginSearchPaths() []string {
	return strings.Split(cGoUnpackString(C.QWebSettings_QWebSettings_PluginSearchPaths()), "|")
}

func QWebSettings_GlobalSettings() *QWebSettings {
	return NewQWebSettingsFromPointer(C.QWebSettings_QWebSettings_GlobalSettings())
}

func (ptr *QWebSettings) GlobalSettings() *QWebSettings {
	return NewQWebSettingsFromPointer(C.QWebSettings_QWebSettings_GlobalSettings())
}

func QWebSettings_MaximumPagesInCache() int {
	return int(int32(C.QWebSettings_QWebSettings_MaximumPagesInCache()))
}

func (ptr *QWebSettings) MaximumPagesInCache() int {
	return int(int32(C.QWebSettings_QWebSettings_MaximumPagesInCache()))
}

func QWebSettings_OfflineStorageDefaultQuota() int64 {
	return int64(C.QWebSettings_QWebSettings_OfflineStorageDefaultQuota())
}

func (ptr *QWebSettings) OfflineStorageDefaultQuota() int64 {
	return int64(C.QWebSettings_QWebSettings_OfflineStorageDefaultQuota())
}

func QWebSettings_OfflineWebApplicationCacheQuota() int64 {
	return int64(C.QWebSettings_QWebSettings_OfflineWebApplicationCacheQuota())
}

func (ptr *QWebSettings) OfflineWebApplicationCacheQuota() int64 {
	return int64(C.QWebSettings_QWebSettings_OfflineWebApplicationCacheQuota())
}

func QWebSettings_ClearIconDatabase() {
	C.QWebSettings_QWebSettings_ClearIconDatabase()
}

func (ptr *QWebSettings) ClearIconDatabase() {
	C.QWebSettings_QWebSettings_ClearIconDatabase()
}

func QWebSettings_ClearMemoryCaches() {
	C.QWebSettings_QWebSettings_ClearMemoryCaches()
}

func (ptr *QWebSettings) ClearMemoryCaches() {
	C.QWebSettings_QWebSettings_ClearMemoryCaches()
}

func QWebSettings_EnablePersistentStorage(path string) {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	C.QWebSettings_QWebSettings_EnablePersistentStorage(C.struct_QtWebKit_PackedString{data: pathC, len: C.longlong(len(path))})
}

func (ptr *QWebSettings) EnablePersistentStorage(path string) {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	C.QWebSettings_QWebSettings_EnablePersistentStorage(C.struct_QtWebKit_PackedString{data: pathC, len: C.longlong(len(path))})
}

func (ptr *QWebSettings) ResetAttribute(attribute QWebSettings__WebAttribute) {
	if ptr.Pointer() != nil {
		C.QWebSettings_ResetAttribute(ptr.Pointer(), C.longlong(attribute))
	}
}

func (ptr *QWebSettings) ResetFontFamily(which QWebSettings__FontFamily) {
	if ptr.Pointer() != nil {
		C.QWebSettings_ResetFontFamily(ptr.Pointer(), C.longlong(which))
	}
}

func (ptr *QWebSettings) ResetFontSize(ty QWebSettings__FontSize) {
	if ptr.Pointer() != nil {
		C.QWebSettings_ResetFontSize(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QWebSettings) SetAttribute(attribute QWebSettings__WebAttribute, on bool) {
	if ptr.Pointer() != nil {
		C.QWebSettings_SetAttribute(ptr.Pointer(), C.longlong(attribute), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QWebSettings) SetCSSMediaType(ty string) {
	if ptr.Pointer() != nil {
		var tyC *C.char
		if ty != "" {
			tyC = C.CString(ty)
			defer C.free(unsafe.Pointer(tyC))
		}
		C.QWebSettings_SetCSSMediaType(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: tyC, len: C.longlong(len(ty))})
	}
}

func (ptr *QWebSettings) SetDefaultTextEncoding(encoding string) {
	if ptr.Pointer() != nil {
		var encodingC *C.char
		if encoding != "" {
			encodingC = C.CString(encoding)
			defer C.free(unsafe.Pointer(encodingC))
		}
		C.QWebSettings_SetDefaultTextEncoding(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: encodingC, len: C.longlong(len(encoding))})
	}
}

func (ptr *QWebSettings) SetFontFamily(which QWebSettings__FontFamily, family string) {
	if ptr.Pointer() != nil {
		var familyC *C.char
		if family != "" {
			familyC = C.CString(family)
			defer C.free(unsafe.Pointer(familyC))
		}
		C.QWebSettings_SetFontFamily(ptr.Pointer(), C.longlong(which), C.struct_QtWebKit_PackedString{data: familyC, len: C.longlong(len(family))})
	}
}

func (ptr *QWebSettings) SetFontSize(ty QWebSettings__FontSize, size int) {
	if ptr.Pointer() != nil {
		C.QWebSettings_SetFontSize(ptr.Pointer(), C.longlong(ty), C.int(int32(size)))
	}
}

func QWebSettings_SetIconDatabasePath(path string) {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	C.QWebSettings_QWebSettings_SetIconDatabasePath(C.struct_QtWebKit_PackedString{data: pathC, len: C.longlong(len(path))})
}

func (ptr *QWebSettings) SetIconDatabasePath(path string) {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	C.QWebSettings_QWebSettings_SetIconDatabasePath(C.struct_QtWebKit_PackedString{data: pathC, len: C.longlong(len(path))})
}

func (ptr *QWebSettings) SetLocalStoragePath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QWebSettings_SetLocalStoragePath(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func QWebSettings_SetMaximumPagesInCache(pages int) {
	C.QWebSettings_QWebSettings_SetMaximumPagesInCache(C.int(int32(pages)))
}

func (ptr *QWebSettings) SetMaximumPagesInCache(pages int) {
	C.QWebSettings_QWebSettings_SetMaximumPagesInCache(C.int(int32(pages)))
}

func QWebSettings_SetObjectCacheCapacities(cacheMinDeadCapacity int, cacheMaxDead int, totalCapacity int) {
	C.QWebSettings_QWebSettings_SetObjectCacheCapacities(C.int(int32(cacheMinDeadCapacity)), C.int(int32(cacheMaxDead)), C.int(int32(totalCapacity)))
}

func (ptr *QWebSettings) SetObjectCacheCapacities(cacheMinDeadCapacity int, cacheMaxDead int, totalCapacity int) {
	C.QWebSettings_QWebSettings_SetObjectCacheCapacities(C.int(int32(cacheMinDeadCapacity)), C.int(int32(cacheMaxDead)), C.int(int32(totalCapacity)))
}

func QWebSettings_SetOfflineStorageDefaultQuota(maximumSize int64) {
	C.QWebSettings_QWebSettings_SetOfflineStorageDefaultQuota(C.longlong(maximumSize))
}

func (ptr *QWebSettings) SetOfflineStorageDefaultQuota(maximumSize int64) {
	C.QWebSettings_QWebSettings_SetOfflineStorageDefaultQuota(C.longlong(maximumSize))
}

func QWebSettings_SetOfflineStoragePath(path string) {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	C.QWebSettings_QWebSettings_SetOfflineStoragePath(C.struct_QtWebKit_PackedString{data: pathC, len: C.longlong(len(path))})
}

func (ptr *QWebSettings) SetOfflineStoragePath(path string) {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	C.QWebSettings_QWebSettings_SetOfflineStoragePath(C.struct_QtWebKit_PackedString{data: pathC, len: C.longlong(len(path))})
}

func QWebSettings_SetOfflineWebApplicationCachePath(path string) {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	C.QWebSettings_QWebSettings_SetOfflineWebApplicationCachePath(C.struct_QtWebKit_PackedString{data: pathC, len: C.longlong(len(path))})
}

func (ptr *QWebSettings) SetOfflineWebApplicationCachePath(path string) {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	C.QWebSettings_QWebSettings_SetOfflineWebApplicationCachePath(C.struct_QtWebKit_PackedString{data: pathC, len: C.longlong(len(path))})
}

func QWebSettings_SetOfflineWebApplicationCacheQuota(maximumSize int64) {
	C.QWebSettings_QWebSettings_SetOfflineWebApplicationCacheQuota(C.longlong(maximumSize))
}

func (ptr *QWebSettings) SetOfflineWebApplicationCacheQuota(maximumSize int64) {
	C.QWebSettings_QWebSettings_SetOfflineWebApplicationCacheQuota(C.longlong(maximumSize))
}

func QWebSettings_SetPluginSearchPaths(paths []string) {
	var pathsC = C.CString(strings.Join(paths, "|"))
	defer C.free(unsafe.Pointer(pathsC))
	C.QWebSettings_QWebSettings_SetPluginSearchPaths(C.struct_QtWebKit_PackedString{data: pathsC, len: C.longlong(len(strings.Join(paths, "|")))})
}

func (ptr *QWebSettings) SetPluginSearchPaths(paths []string) {
	var pathsC = C.CString(strings.Join(paths, "|"))
	defer C.free(unsafe.Pointer(pathsC))
	C.QWebSettings_QWebSettings_SetPluginSearchPaths(C.struct_QtWebKit_PackedString{data: pathsC, len: C.longlong(len(strings.Join(paths, "|")))})
}

func (ptr *QWebSettings) SetThirdPartyCookiePolicy(policy QWebSettings__ThirdPartyCookiePolicy) {
	if ptr.Pointer() != nil {
		C.QWebSettings_SetThirdPartyCookiePolicy(ptr.Pointer(), C.longlong(policy))
	}
}

func (ptr *QWebSettings) SetUserStyleSheetUrl(location core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSettings_SetUserStyleSheetUrl(ptr.Pointer(), core.PointerFromQUrl(location))
	}
}

func QWebSettings_SetWebGraphic(ty QWebSettings__WebGraphic, graphic gui.QPixmap_ITF) {
	C.QWebSettings_QWebSettings_SetWebGraphic(C.longlong(ty), gui.PointerFromQPixmap(graphic))
}

func (ptr *QWebSettings) SetWebGraphic(ty QWebSettings__WebGraphic, graphic gui.QPixmap_ITF) {
	C.QWebSettings_QWebSettings_SetWebGraphic(C.longlong(ty), gui.PointerFromQPixmap(graphic))
}

func (ptr *QWebSettings) CssMediaType() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSettings_CssMediaType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSettings) DefaultTextEncoding() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSettings_DefaultTextEncoding(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSettings) FontFamily(which QWebSettings__FontFamily) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSettings_FontFamily(ptr.Pointer(), C.longlong(which)))
	}
	return ""
}

func (ptr *QWebSettings) LocalStoragePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSettings_LocalStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSettings) UserStyleSheetUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebSettings_UserStyleSheetUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSettings) ThirdPartyCookiePolicy() QWebSettings__ThirdPartyCookiePolicy {
	if ptr.Pointer() != nil {
		return QWebSettings__ThirdPartyCookiePolicy(C.QWebSettings_ThirdPartyCookiePolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSettings) TestAttribute(attribute QWebSettings__WebAttribute) bool {
	if ptr.Pointer() != nil {
		return C.QWebSettings_TestAttribute(ptr.Pointer(), C.longlong(attribute)) != 0
	}
	return false
}

func (ptr *QWebSettings) FontSize(ty QWebSettings__FontSize) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebSettings_FontSize(ptr.Pointer(), C.longlong(ty))))
	}
	return 0
}

type QWebView struct {
	widgets.QWidget
}

type QWebView_ITF interface {
	widgets.QWidget_ITF
	QWebView_PTR() *QWebView
}

func (ptr *QWebView) QWebView_PTR() *QWebView {
	return ptr
}

func (ptr *QWebView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQWebView(ptr QWebView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebView_PTR().Pointer()
	}
	return nil
}

func NewQWebViewFromPointer(ptr unsafe.Pointer) *QWebView {
	var n = new(QWebView)
	n.SetPointer(ptr)
	return n
}

//export callbackQWebView_CreateWindow
func callbackQWebView_CreateWindow(ptr unsafe.Pointer, ty C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createWindow"); signal != nil {
		return PointerFromQWebView(signal.(func(QWebPage__WebWindowType) *QWebView)(QWebPage__WebWindowType(ty)))
	}

	return PointerFromQWebView(NewQWebViewFromPointer(ptr).CreateWindowDefault(QWebPage__WebWindowType(ty)))
}

func (ptr *QWebView) ConnectCreateWindow(f func(ty QWebPage__WebWindowType) *QWebView) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createWindow", func(ty QWebPage__WebWindowType) *QWebView {
				signal.(func(QWebPage__WebWindowType) *QWebView)(ty)
				return f(ty)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createWindow", f)
		}
	}
}

func (ptr *QWebView) DisconnectCreateWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createWindow")
	}
}

func (ptr *QWebView) CreateWindow(ty QWebPage__WebWindowType) *QWebView {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebViewFromPointer(C.QWebView_CreateWindow(ptr.Pointer(), C.longlong(ty)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) CreateWindowDefault(ty QWebPage__WebWindowType) *QWebView {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebViewFromPointer(C.QWebView_CreateWindowDefault(ptr.Pointer(), C.longlong(ty)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQWebView(parent widgets.QWidget_ITF) *QWebView {
	var tmpValue = NewQWebViewFromPointer(C.QWebView_NewQWebView(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebView_Event
func callbackQWebView_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebView) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebView_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebView) FindText(subString string, options QWebPage__FindFlag) bool {
	if ptr.Pointer() != nil {
		var subStringC *C.char
		if subString != "" {
			subStringC = C.CString(subString)
			defer C.free(unsafe.Pointer(subStringC))
		}
		return C.QWebView_FindText(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: subStringC, len: C.longlong(len(subString))}, C.longlong(options)) != 0
	}
	return false
}

//export callbackQWebView_FocusNextPrevChild
func callbackQWebView_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QWebView) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QWebView_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQWebView_Back
func callbackQWebView_Back(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "back"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).BackDefault()
	}
}

func (ptr *QWebView) ConnectBack(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "back"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "back", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "back", f)
		}
	}
}

func (ptr *QWebView) DisconnectBack() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "back")
	}
}

func (ptr *QWebView) Back() {
	if ptr.Pointer() != nil {
		C.QWebView_Back(ptr.Pointer())
	}
}

func (ptr *QWebView) BackDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_BackDefault(ptr.Pointer())
	}
}

//export callbackQWebView_ChangeEvent
func callbackQWebView_ChangeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQWebViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QWebView) ChangeEventDefault(e core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

//export callbackQWebView_ContextMenuEvent
func callbackQWebView_ContextMenuEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(ev))
	}
}

func (ptr *QWebView) ContextMenuEventDefault(ev gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(ev))
	}
}

//export callbackQWebView_DragEnterEvent
func callbackQWebView_DragEnterEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(ev))
	}
}

func (ptr *QWebView) DragEnterEventDefault(ev gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(ev))
	}
}

//export callbackQWebView_DragLeaveEvent
func callbackQWebView_DragLeaveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(ev))
	}
}

func (ptr *QWebView) DragLeaveEventDefault(ev gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(ev))
	}
}

//export callbackQWebView_DragMoveEvent
func callbackQWebView_DragMoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(ev))
	}
}

func (ptr *QWebView) DragMoveEventDefault(ev gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(ev))
	}
}

//export callbackQWebView_DropEvent
func callbackQWebView_DropEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(ev))
	}
}

func (ptr *QWebView) DropEventDefault(ev gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(ev))
	}
}

//export callbackQWebView_FocusInEvent
func callbackQWebView_FocusInEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QWebView) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQWebView_FocusOutEvent
func callbackQWebView_FocusOutEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QWebView) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQWebView_Forward
func callbackQWebView_Forward(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "forward"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ForwardDefault()
	}
}

func (ptr *QWebView) ConnectForward(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "forward"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "forward", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "forward", f)
		}
	}
}

func (ptr *QWebView) DisconnectForward() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "forward")
	}
}

func (ptr *QWebView) Forward() {
	if ptr.Pointer() != nil {
		C.QWebView_Forward(ptr.Pointer())
	}
}

func (ptr *QWebView) ForwardDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_ForwardDefault(ptr.Pointer())
	}
}

//export callbackQWebView_IconChanged
func callbackQWebView_IconChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebView) ConnectIconChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "iconChanged") {
			C.QWebView_ConnectIconChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "iconChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "iconChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "iconChanged", f)
		}
	}
}

func (ptr *QWebView) DisconnectIconChanged() {
	if ptr.Pointer() != nil {
		C.QWebView_DisconnectIconChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "iconChanged")
	}
}

func (ptr *QWebView) IconChanged() {
	if ptr.Pointer() != nil {
		C.QWebView_IconChanged(ptr.Pointer())
	}
}

//export callbackQWebView_InputMethodEvent
func callbackQWebView_InputMethodEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
	} else {
		NewQWebViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(e))
	}
}

func (ptr *QWebView) InputMethodEventDefault(e gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

//export callbackQWebView_KeyPressEvent
func callbackQWebView_KeyPressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QWebView) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackQWebView_KeyReleaseEvent
func callbackQWebView_KeyReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QWebView) KeyReleaseEventDefault(ev gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackQWebView_LinkClicked
func callbackQWebView_LinkClicked(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "linkClicked"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebView) ConnectLinkClicked(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "linkClicked") {
			C.QWebView_ConnectLinkClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "linkClicked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "linkClicked", func(url *core.QUrl) {
				signal.(func(*core.QUrl))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "linkClicked", f)
		}
	}
}

func (ptr *QWebView) DisconnectLinkClicked() {
	if ptr.Pointer() != nil {
		C.QWebView_DisconnectLinkClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "linkClicked")
	}
}

func (ptr *QWebView) LinkClicked(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_LinkClicked(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebView) Load2(request network.QNetworkRequest_ITF, operation network.QNetworkAccessManager__Operation, body core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_Load2(ptr.Pointer(), network.PointerFromQNetworkRequest(request), C.longlong(operation), core.PointerFromQByteArray(body))
	}
}

func (ptr *QWebView) Load(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebView_LoadFinished
func callbackQWebView_LoadFinished(ptr unsafe.Pointer, ok C.char) {
	if signal := qt.GetSignal(ptr, "loadFinished"); signal != nil {
		signal.(func(bool))(int8(ok) != 0)
	}

}

func (ptr *QWebView) ConnectLoadFinished(f func(ok bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadFinished") {
			C.QWebView_ConnectLoadFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", func(ok bool) {
				signal.(func(bool))(ok)
				f(ok)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadFinished", f)
		}
	}
}

func (ptr *QWebView) DisconnectLoadFinished() {
	if ptr.Pointer() != nil {
		C.QWebView_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadFinished")
	}
}

func (ptr *QWebView) LoadFinished(ok bool) {
	if ptr.Pointer() != nil {
		C.QWebView_LoadFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))
	}
}

//export callbackQWebView_LoadProgress
func callbackQWebView_LoadProgress(ptr unsafe.Pointer, progress C.int) {
	if signal := qt.GetSignal(ptr, "loadProgress"); signal != nil {
		signal.(func(int))(int(int32(progress)))
	}

}

func (ptr *QWebView) ConnectLoadProgress(f func(progress int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadProgress") {
			C.QWebView_ConnectLoadProgress(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadProgress"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadProgress", func(progress int) {
				signal.(func(int))(progress)
				f(progress)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadProgress", f)
		}
	}
}

func (ptr *QWebView) DisconnectLoadProgress() {
	if ptr.Pointer() != nil {
		C.QWebView_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadProgress")
	}
}

func (ptr *QWebView) LoadProgress(progress int) {
	if ptr.Pointer() != nil {
		C.QWebView_LoadProgress(ptr.Pointer(), C.int(int32(progress)))
	}
}

//export callbackQWebView_LoadStarted
func callbackQWebView_LoadStarted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebView) ConnectLoadStarted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadStarted") {
			C.QWebView_ConnectLoadStarted(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadStarted"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadStarted", f)
		}
	}
}

func (ptr *QWebView) DisconnectLoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebView_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadStarted")
	}
}

func (ptr *QWebView) LoadStarted() {
	if ptr.Pointer() != nil {
		C.QWebView_LoadStarted(ptr.Pointer())
	}
}

//export callbackQWebView_MouseDoubleClickEvent
func callbackQWebView_MouseDoubleClickEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWebView) MouseDoubleClickEventDefault(ev gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackQWebView_MouseMoveEvent
func callbackQWebView_MouseMoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWebView) MouseMoveEventDefault(ev gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackQWebView_MousePressEvent
func callbackQWebView_MousePressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWebView) MousePressEventDefault(ev gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackQWebView_MouseReleaseEvent
func callbackQWebView_MouseReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWebView) MouseReleaseEventDefault(ev gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackQWebView_PaintEvent
func callbackQWebView_PaintEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(ev))
	}
}

func (ptr *QWebView) PaintEventDefault(ev gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(ev))
	}
}

//export callbackQWebView_Reload
func callbackQWebView_Reload(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reload"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ReloadDefault()
	}
}

func (ptr *QWebView) ConnectReload(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reload"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "reload", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reload", f)
		}
	}
}

func (ptr *QWebView) DisconnectReload() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reload")
	}
}

func (ptr *QWebView) Reload() {
	if ptr.Pointer() != nil {
		C.QWebView_Reload(ptr.Pointer())
	}
}

func (ptr *QWebView) ReloadDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_ReloadDefault(ptr.Pointer())
	}
}

//export callbackQWebView_ResizeEvent
func callbackQWebView_ResizeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQWebViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QWebView) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

//export callbackQWebView_SelectionChanged
func callbackQWebView_SelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebView) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionChanged") {
			C.QWebView_ConnectSelectionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", f)
		}
	}
}

func (ptr *QWebView) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebView_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionChanged")
	}
}

func (ptr *QWebView) SelectionChanged() {
	if ptr.Pointer() != nil {
		C.QWebView_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QWebView) SetContent(data core.QByteArray_ITF, mimeType string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var mimeTypeC *C.char
		if mimeType != "" {
			mimeTypeC = C.CString(mimeType)
			defer C.free(unsafe.Pointer(mimeTypeC))
		}
		C.QWebView_SetContent(ptr.Pointer(), core.PointerFromQByteArray(data), C.struct_QtWebKit_PackedString{data: mimeTypeC, len: C.longlong(len(mimeType))}, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebView) SetHtml(html string, baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var htmlC *C.char
		if html != "" {
			htmlC = C.CString(html)
			defer C.free(unsafe.Pointer(htmlC))
		}
		C.QWebView_SetHtml(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: htmlC, len: C.longlong(len(html))}, core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebView) SetPage(page QWebPage_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_SetPage(ptr.Pointer(), PointerFromQWebPage(page))
	}
}

func (ptr *QWebView) SetRenderHint(hint gui.QPainter__RenderHint, enabled bool) {
	if ptr.Pointer() != nil {
		C.QWebView_SetRenderHint(ptr.Pointer(), C.longlong(hint), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QWebView) SetRenderHints(hints gui.QPainter__RenderHint) {
	if ptr.Pointer() != nil {
		C.QWebView_SetRenderHints(ptr.Pointer(), C.longlong(hints))
	}
}

func (ptr *QWebView) SetTextSizeMultiplier(factor float64) {
	if ptr.Pointer() != nil {
		C.QWebView_SetTextSizeMultiplier(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebView) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebView) SetZoomFactor(factor float64) {
	if ptr.Pointer() != nil {
		C.QWebView_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

//export callbackQWebView_StatusBarMessage
func callbackQWebView_StatusBarMessage(ptr unsafe.Pointer, text C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "statusBarMessage"); signal != nil {
		signal.(func(string))(cGoUnpackString(text))
	}

}

func (ptr *QWebView) ConnectStatusBarMessage(f func(text string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusBarMessage") {
			C.QWebView_ConnectStatusBarMessage(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "statusBarMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "statusBarMessage", func(text string) {
				signal.(func(string))(text)
				f(text)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusBarMessage", f)
		}
	}
}

func (ptr *QWebView) DisconnectStatusBarMessage() {
	if ptr.Pointer() != nil {
		C.QWebView_DisconnectStatusBarMessage(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusBarMessage")
	}
}

func (ptr *QWebView) StatusBarMessage(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QWebView_StatusBarMessage(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQWebView_Stop
func callbackQWebView_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).StopDefault()
	}
}

func (ptr *QWebView) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stop", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", f)
		}
	}
}

func (ptr *QWebView) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QWebView) Stop() {
	if ptr.Pointer() != nil {
		C.QWebView_Stop(ptr.Pointer())
	}
}

func (ptr *QWebView) StopDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_StopDefault(ptr.Pointer())
	}
}

//export callbackQWebView_TitleChanged
func callbackQWebView_TitleChanged(ptr unsafe.Pointer, title C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "titleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

func (ptr *QWebView) ConnectTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "titleChanged") {
			C.QWebView_ConnectTitleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "titleChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", func(title string) {
				signal.(func(string))(title)
				f(title)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "titleChanged", f)
		}
	}
}

func (ptr *QWebView) DisconnectTitleChanged() {
	if ptr.Pointer() != nil {
		C.QWebView_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "titleChanged")
	}
}

func (ptr *QWebView) TitleChanged(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QWebView_TitleChanged(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *QWebView) TriggerPageAction(action QWebPage__WebAction, checked bool) {
	if ptr.Pointer() != nil {
		C.QWebView_TriggerPageAction(ptr.Pointer(), C.longlong(action), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

//export callbackQWebView_UrlChanged
func callbackQWebView_UrlChanged(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "urlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebView) ConnectUrlChanged(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "urlChanged") {
			C.QWebView_ConnectUrlChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "urlChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "urlChanged", func(url *core.QUrl) {
				signal.(func(*core.QUrl))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "urlChanged", f)
		}
	}
}

func (ptr *QWebView) DisconnectUrlChanged() {
	if ptr.Pointer() != nil {
		C.QWebView_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "urlChanged")
	}
}

func (ptr *QWebView) UrlChanged(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebView_WheelEvent
func callbackQWebView_WheelEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(ev))
	}
}

func (ptr *QWebView) WheelEventDefault(ev gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(ev))
	}
}

func (ptr *QWebView) DestroyQWebView() {
	if ptr.Pointer() != nil {
		C.QWebView_DestroyQWebView(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebView) PageAction(action QWebPage__WebAction) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebView_PageAction(ptr.Pointer(), C.longlong(action)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQIconFromPointer(C.QWebView_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) RenderHints() gui.QPainter__RenderHint {
	if ptr.Pointer() != nil {
		return gui.QPainter__RenderHint(C.QWebView_RenderHints(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebView_SizeHint
func callbackQWebView_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebViewFromPointer(ptr).SizeHintDefault())
}

func (ptr *QWebView) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebView_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) SelectedHtml() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebView_SelectedHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebView) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebView_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebView) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebView_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebView) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebView_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQWebView_InputMethodQuery
func callbackQWebView_InputMethodQuery(ptr unsafe.Pointer, property C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(property)))
	}

	return core.PointerFromQVariant(NewQWebViewFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(property)))
}

func (ptr *QWebView) InputMethodQueryDefault(property core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QWebView_InputMethodQueryDefault(ptr.Pointer(), C.longlong(property)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) History() *QWebHistory {
	if ptr.Pointer() != nil {
		return NewQWebHistoryFromPointer(C.QWebView_History(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebView) Page() *QWebPage {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebPageFromPointer(C.QWebView_Page(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) Settings() *QWebSettings {
	if ptr.Pointer() != nil {
		return NewQWebSettingsFromPointer(C.QWebView_Settings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebView) HasSelection() bool {
	if ptr.Pointer() != nil {
		return C.QWebView_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebView) IsModified() bool {
	if ptr.Pointer() != nil {
		return C.QWebView_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebView) TextSizeMultiplier() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWebView_TextSizeMultiplier(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebView) ZoomFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWebView_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebView_Print
func callbackQWebView_Print(ptr unsafe.Pointer, printer unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "print"); signal != nil {
		signal.(func(*printsupport.QPrinter))(printsupport.NewQPrinterFromPointer(printer))
	} else {
		NewQWebViewFromPointer(ptr).PrintDefault(printsupport.NewQPrinterFromPointer(printer))
	}
}

func (ptr *QWebView) ConnectPrint(f func(printer *printsupport.QPrinter)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "print"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "print", func(printer *printsupport.QPrinter) {
				signal.(func(*printsupport.QPrinter))(printer)
				f(printer)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "print", f)
		}
	}
}

func (ptr *QWebView) DisconnectPrint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "print")
	}
}

func (ptr *QWebView) Print(printer printsupport.QPrinter_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_Print(ptr.Pointer(), printsupport.PointerFromQPrinter(printer))
	}
}

func (ptr *QWebView) PrintDefault(printer printsupport.QPrinter_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_PrintDefault(ptr.Pointer(), printsupport.PointerFromQPrinter(printer))
	}
}

func (ptr *QWebView) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebView___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QWebView) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebView___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *QWebView) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebView___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QWebView) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebView___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *QWebView) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QWebView___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QWebView) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebView___actions_newList(ptr.Pointer()))
}

func (ptr *QWebView) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QWebView___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebView) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebView___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QWebView) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebView___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebView) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QWebView___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QWebView) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebView___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebView) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QWebView___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QWebView) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebView___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebView) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebView___findChildren_newList(ptr.Pointer()))
}

func (ptr *QWebView) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebView___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebView) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebView) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QWebView___children_newList(ptr.Pointer()))
}

//export callbackQWebView_Close
func callbackQWebView_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).CloseDefault())))
}

func (ptr *QWebView) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QWebView_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebView_NativeEvent
func callbackQWebView_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QWebView) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QWebView_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQWebView_ActionEvent
func callbackQWebView_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QWebView) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQWebView_CloseEvent
func callbackQWebView_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QWebView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQWebView_CustomContextMenuRequested
func callbackQWebView_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQWebView_EnterEvent
func callbackQWebView_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebView) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebView_Hide
func callbackQWebView_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).HideDefault()
	}
}

func (ptr *QWebView) HideDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_HideDefault(ptr.Pointer())
	}
}

//export callbackQWebView_HideEvent
func callbackQWebView_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QWebView) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQWebView_LeaveEvent
func callbackQWebView_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebView) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebView_Lower
func callbackQWebView_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QWebView) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_LowerDefault(ptr.Pointer())
	}
}

//export callbackQWebView_MoveEvent
func callbackQWebView_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QWebView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQWebView_Raise
func callbackQWebView_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QWebView) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQWebView_Repaint
func callbackQWebView_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QWebView) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQWebView_SetDisabled
func callbackQWebView_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQWebViewFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QWebView) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QWebView_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQWebView_SetEnabled
func callbackQWebView_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQWebViewFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QWebView) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebView_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQWebView_SetFocus2
func callbackQWebView_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QWebView) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QWebView_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQWebView_SetHidden
func callbackQWebView_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQWebViewFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QWebView) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QWebView_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQWebView_SetStyleSheet
func callbackQWebView_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQWebViewFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QWebView) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QWebView_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQWebView_SetVisible
func callbackQWebView_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQWebViewFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QWebView) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QWebView_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQWebView_SetWindowModified
func callbackQWebView_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQWebViewFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QWebView) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWebView_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQWebView_SetWindowTitle
func callbackQWebView_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQWebViewFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QWebView) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QWebView_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtWebKit_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQWebView_Show
func callbackQWebView_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWebView) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_ShowDefault(ptr.Pointer())
	}
}

//export callbackQWebView_ShowEvent
func callbackQWebView_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QWebView) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQWebView_ShowFullScreen
func callbackQWebView_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QWebView) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQWebView_ShowMaximized
func callbackQWebView_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QWebView) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQWebView_ShowMinimized
func callbackQWebView_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QWebView) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQWebView_ShowNormal
func callbackQWebView_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QWebView) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQWebView_TabletEvent
func callbackQWebView_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QWebView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQWebView_Update
func callbackQWebView_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QWebView) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQWebView_UpdateMicroFocus
func callbackQWebView_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QWebView) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQWebView_WindowIconChanged
func callbackQWebView_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQWebView_WindowTitleChanged
func callbackQWebView_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQWebView_PaintEngine
func callbackQWebView_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQWebViewFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QWebView) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWebView_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebView_MinimumSizeHint
func callbackQWebView_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebViewFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QWebView) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QWebView_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQWebView_HasHeightForWidth
func callbackQWebView_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QWebView) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QWebView_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebView_HeightForWidth
func callbackQWebView_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQWebViewFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QWebView) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebView_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQWebView_Metric
func callbackQWebView_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQWebViewFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QWebView) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebView_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQWebView_EventFilter
func callbackQWebView_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebView_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebView_ChildEvent
func callbackQWebView_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebView) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebView_ConnectNotify
func callbackQWebView_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebViewFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebView_CustomEvent
func callbackQWebView_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebView) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebView_DeleteLater
func callbackQWebView_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebView) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebView_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebView_Destroyed
func callbackQWebView_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebView_DisconnectNotify
func callbackQWebView_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebViewFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebView_ObjectNameChanged
func callbackQWebView_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebKit_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQWebView_TimerEvent
func callbackQWebView_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebView) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebView_MetaObject
func callbackQWebView_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebViewFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebView) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebView_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
