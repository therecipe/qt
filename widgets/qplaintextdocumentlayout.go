package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QPlainTextDocumentLayout struct {
	gui.QAbstractTextDocumentLayout
}

type QPlainTextDocumentLayout_ITF interface {
	gui.QAbstractTextDocumentLayout_ITF
	QPlainTextDocumentLayout_PTR() *QPlainTextDocumentLayout
}

func PointerFromQPlainTextDocumentLayout(ptr QPlainTextDocumentLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlainTextDocumentLayout_PTR().Pointer()
	}
	return nil
}

func NewQPlainTextDocumentLayoutFromPointer(ptr unsafe.Pointer) *QPlainTextDocumentLayout {
	var n = new(QPlainTextDocumentLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPlainTextDocumentLayout_") {
		n.SetObjectName("QPlainTextDocumentLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QPlainTextDocumentLayout) QPlainTextDocumentLayout_PTR() *QPlainTextDocumentLayout {
	return ptr
}

func (ptr *QPlainTextDocumentLayout) CursorWidth() int {
	defer qt.Recovering("QPlainTextDocumentLayout::cursorWidth")

	if ptr.Pointer() != nil {
		return int(C.QPlainTextDocumentLayout_CursorWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextDocumentLayout) SetCursorWidth(width int) {
	defer qt.Recovering("QPlainTextDocumentLayout::setCursorWidth")

	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_SetCursorWidth(ptr.Pointer(), C.int(width))
	}
}

func NewQPlainTextDocumentLayout(document gui.QTextDocument_ITF) *QPlainTextDocumentLayout {
	defer qt.Recovering("QPlainTextDocumentLayout::QPlainTextDocumentLayout")

	return NewQPlainTextDocumentLayoutFromPointer(C.QPlainTextDocumentLayout_NewQPlainTextDocumentLayout(gui.PointerFromQTextDocument(document)))
}

func (ptr *QPlainTextDocumentLayout) ConnectDocumentChanged(f func(from int, charsRemoved int, charsAdded int)) {
	defer qt.Recovering("connect QPlainTextDocumentLayout::documentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "documentChanged", f)
	}
}

func (ptr *QPlainTextDocumentLayout) DisconnectDocumentChanged() {
	defer qt.Recovering("disconnect QPlainTextDocumentLayout::documentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "documentChanged")
	}
}

//export callbackQPlainTextDocumentLayoutDocumentChanged
func callbackQPlainTextDocumentLayoutDocumentChanged(ptrName *C.char, from C.int, charsRemoved C.int, charsAdded C.int) bool {
	defer qt.Recovering("callback QPlainTextDocumentLayout::documentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "documentChanged"); signal != nil {
		signal.(func(int, int, int))(int(from), int(charsRemoved), int(charsAdded))
		return true
	}
	return false

}

func (ptr *QPlainTextDocumentLayout) EnsureBlockLayout(block gui.QTextBlock_ITF) {
	defer qt.Recovering("QPlainTextDocumentLayout::ensureBlockLayout")

	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_EnsureBlockLayout(ptr.Pointer(), gui.PointerFromQTextBlock(block))
	}
}

func (ptr *QPlainTextDocumentLayout) PageCount() int {
	defer qt.Recovering("QPlainTextDocumentLayout::pageCount")

	if ptr.Pointer() != nil {
		return int(C.QPlainTextDocumentLayout_PageCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextDocumentLayout) RequestUpdate() {
	defer qt.Recovering("QPlainTextDocumentLayout::requestUpdate")

	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_RequestUpdate(ptr.Pointer())
	}
}

func (ptr *QPlainTextDocumentLayout) DestroyQPlainTextDocumentLayout() {
	defer qt.Recovering("QPlainTextDocumentLayout::~QPlainTextDocumentLayout")

	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_DestroyQPlainTextDocumentLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPlainTextDocumentLayout) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPlainTextDocumentLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPlainTextDocumentLayout) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPlainTextDocumentLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPlainTextDocumentLayoutTimerEvent
func callbackQPlainTextDocumentLayoutTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextDocumentLayout::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPlainTextDocumentLayout) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPlainTextDocumentLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPlainTextDocumentLayout) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPlainTextDocumentLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPlainTextDocumentLayoutChildEvent
func callbackQPlainTextDocumentLayoutChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextDocumentLayout::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPlainTextDocumentLayout) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPlainTextDocumentLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPlainTextDocumentLayout) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPlainTextDocumentLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPlainTextDocumentLayoutCustomEvent
func callbackQPlainTextDocumentLayoutCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextDocumentLayout::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
