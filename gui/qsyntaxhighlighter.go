package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSyntaxHighlighter struct {
	core.QObject
}

type QSyntaxHighlighter_ITF interface {
	core.QObject_ITF
	QSyntaxHighlighter_PTR() *QSyntaxHighlighter
}

func PointerFromQSyntaxHighlighter(ptr QSyntaxHighlighter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSyntaxHighlighter_PTR().Pointer()
	}
	return nil
}

func NewQSyntaxHighlighterFromPointer(ptr unsafe.Pointer) *QSyntaxHighlighter {
	var n = new(QSyntaxHighlighter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSyntaxHighlighter_") {
		n.SetObjectName("QSyntaxHighlighter_" + qt.Identifier())
	}
	return n
}

func (ptr *QSyntaxHighlighter) QSyntaxHighlighter_PTR() *QSyntaxHighlighter {
	return ptr
}

func (ptr *QSyntaxHighlighter) Document() *QTextDocument {
	defer qt.Recovering("QSyntaxHighlighter::document")

	if ptr.Pointer() != nil {
		return NewQTextDocumentFromPointer(C.QSyntaxHighlighter_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSyntaxHighlighter) Rehighlight() {
	defer qt.Recovering("QSyntaxHighlighter::rehighlight")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_Rehighlight(ptr.Pointer())
	}
}

func (ptr *QSyntaxHighlighter) RehighlightBlock(block QTextBlock_ITF) {
	defer qt.Recovering("QSyntaxHighlighter::rehighlightBlock")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_RehighlightBlock(ptr.Pointer(), PointerFromQTextBlock(block))
	}
}

func (ptr *QSyntaxHighlighter) SetDocument(doc QTextDocument_ITF) {
	defer qt.Recovering("QSyntaxHighlighter::setDocument")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_SetDocument(ptr.Pointer(), PointerFromQTextDocument(doc))
	}
}

func (ptr *QSyntaxHighlighter) DestroyQSyntaxHighlighter() {
	defer qt.Recovering("QSyntaxHighlighter::~QSyntaxHighlighter")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_DestroyQSyntaxHighlighter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSyntaxHighlighter) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSyntaxHighlighter::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSyntaxHighlighter) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSyntaxHighlighter::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSyntaxHighlighterTimerEvent
func callbackQSyntaxHighlighterTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSyntaxHighlighter::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSyntaxHighlighterFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSyntaxHighlighter) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSyntaxHighlighter::timerEvent")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSyntaxHighlighter) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSyntaxHighlighter::timerEvent")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSyntaxHighlighter) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSyntaxHighlighter::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSyntaxHighlighter) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSyntaxHighlighter::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSyntaxHighlighterChildEvent
func callbackQSyntaxHighlighterChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSyntaxHighlighter::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSyntaxHighlighterFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSyntaxHighlighter) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSyntaxHighlighter::childEvent")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSyntaxHighlighter) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSyntaxHighlighter::childEvent")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSyntaxHighlighter) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSyntaxHighlighter::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSyntaxHighlighter) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSyntaxHighlighter::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSyntaxHighlighterCustomEvent
func callbackQSyntaxHighlighterCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSyntaxHighlighter::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSyntaxHighlighterFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSyntaxHighlighter) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSyntaxHighlighter::customEvent")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSyntaxHighlighter) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSyntaxHighlighter::customEvent")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
