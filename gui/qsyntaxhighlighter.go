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
func callbackQSyntaxHighlighterTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSyntaxHighlighter::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSyntaxHighlighterChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSyntaxHighlighter::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSyntaxHighlighterCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSyntaxHighlighter::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
