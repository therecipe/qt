package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QQuickTextDocument struct {
	core.QObject
}

type QQuickTextDocument_ITF interface {
	core.QObject_ITF
	QQuickTextDocument_PTR() *QQuickTextDocument
}

func PointerFromQQuickTextDocument(ptr QQuickTextDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextDocument_PTR().Pointer()
	}
	return nil
}

func NewQQuickTextDocumentFromPointer(ptr unsafe.Pointer) *QQuickTextDocument {
	var n = new(QQuickTextDocument)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickTextDocument_") {
		n.SetObjectName("QQuickTextDocument_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickTextDocument) QQuickTextDocument_PTR() *QQuickTextDocument {
	return ptr
}

func NewQQuickTextDocument(parent QQuickItem_ITF) *QQuickTextDocument {
	defer qt.Recovering("QQuickTextDocument::QQuickTextDocument")

	return NewQQuickTextDocumentFromPointer(C.QQuickTextDocument_NewQQuickTextDocument(PointerFromQQuickItem(parent)))
}

func (ptr *QQuickTextDocument) TextDocument() *gui.QTextDocument {
	defer qt.Recovering("QQuickTextDocument::textDocument")

	if ptr.Pointer() != nil {
		return gui.NewQTextDocumentFromPointer(C.QQuickTextDocument_TextDocument(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextDocument) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickTextDocument::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickTextDocument::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickTextDocumentTimerEvent
func callbackQQuickTextDocumentTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickTextDocument) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickTextDocument) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickTextDocument::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickTextDocument::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickTextDocumentChildEvent
func callbackQQuickTextDocumentChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickTextDocument) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickTextDocument) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickTextDocument::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickTextDocument::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickTextDocumentCustomEvent
func callbackQQuickTextDocumentCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickTextDocument) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
