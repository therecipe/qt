package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractTextDocumentLayout struct {
	core.QObject
}

type QAbstractTextDocumentLayout_ITF interface {
	core.QObject_ITF
	QAbstractTextDocumentLayout_PTR() *QAbstractTextDocumentLayout
}

func PointerFromQAbstractTextDocumentLayout(ptr QAbstractTextDocumentLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractTextDocumentLayout_PTR().Pointer()
	}
	return nil
}

func NewQAbstractTextDocumentLayoutFromPointer(ptr unsafe.Pointer) *QAbstractTextDocumentLayout {
	var n = new(QAbstractTextDocumentLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractTextDocumentLayout_") {
		n.SetObjectName("QAbstractTextDocumentLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractTextDocumentLayout) QAbstractTextDocumentLayout_PTR() *QAbstractTextDocumentLayout {
	return ptr
}

func (ptr *QAbstractTextDocumentLayout) AnchorAt(position core.QPointF_ITF) string {
	defer qt.Recovering("QAbstractTextDocumentLayout::anchorAt")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractTextDocumentLayout_AnchorAt(ptr.Pointer(), core.PointerFromQPointF(position)))
	}
	return ""
}

func (ptr *QAbstractTextDocumentLayout) Document() *QTextDocument {
	defer qt.Recovering("QAbstractTextDocumentLayout::document")

	if ptr.Pointer() != nil {
		return NewQTextDocumentFromPointer(C.QAbstractTextDocumentLayout_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractTextDocumentLayout) HandlerForObject(objectType int) *QTextObjectInterface {
	defer qt.Recovering("QAbstractTextDocumentLayout::handlerForObject")

	if ptr.Pointer() != nil {
		return NewQTextObjectInterfaceFromPointer(C.QAbstractTextDocumentLayout_HandlerForObject(ptr.Pointer(), C.int(objectType)))
	}
	return nil
}

func (ptr *QAbstractTextDocumentLayout) PageCount() int {
	defer qt.Recovering("QAbstractTextDocumentLayout::pageCount")

	if ptr.Pointer() != nil {
		return int(C.QAbstractTextDocumentLayout_PageCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractTextDocumentLayout) ConnectPageCountChanged(f func(newPages int)) {
	defer qt.Recovering("connect QAbstractTextDocumentLayout::pageCountChanged")

	if ptr.Pointer() != nil {
		C.QAbstractTextDocumentLayout_ConnectPageCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pageCountChanged", f)
	}
}

func (ptr *QAbstractTextDocumentLayout) DisconnectPageCountChanged() {
	defer qt.Recovering("disconnect QAbstractTextDocumentLayout::pageCountChanged")

	if ptr.Pointer() != nil {
		C.QAbstractTextDocumentLayout_DisconnectPageCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pageCountChanged")
	}
}

//export callbackQAbstractTextDocumentLayoutPageCountChanged
func callbackQAbstractTextDocumentLayoutPageCountChanged(ptrName *C.char, newPages C.int) {
	defer qt.Recovering("callback QAbstractTextDocumentLayout::pageCountChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "pageCountChanged")
	if signal != nil {
		signal.(func(int))(int(newPages))
	}

}

func (ptr *QAbstractTextDocumentLayout) PaintDevice() *QPaintDevice {
	defer qt.Recovering("QAbstractTextDocumentLayout::paintDevice")

	if ptr.Pointer() != nil {
		return NewQPaintDeviceFromPointer(C.QAbstractTextDocumentLayout_PaintDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractTextDocumentLayout) RegisterHandler(objectType int, component core.QObject_ITF) {
	defer qt.Recovering("QAbstractTextDocumentLayout::registerHandler")

	if ptr.Pointer() != nil {
		C.QAbstractTextDocumentLayout_RegisterHandler(ptr.Pointer(), C.int(objectType), core.PointerFromQObject(component))
	}
}

func (ptr *QAbstractTextDocumentLayout) SetPaintDevice(device QPaintDevice_ITF) {
	defer qt.Recovering("QAbstractTextDocumentLayout::setPaintDevice")

	if ptr.Pointer() != nil {
		C.QAbstractTextDocumentLayout_SetPaintDevice(ptr.Pointer(), PointerFromQPaintDevice(device))
	}
}

func (ptr *QAbstractTextDocumentLayout) UnregisterHandler(objectType int, component core.QObject_ITF) {
	defer qt.Recovering("QAbstractTextDocumentLayout::unregisterHandler")

	if ptr.Pointer() != nil {
		C.QAbstractTextDocumentLayout_UnregisterHandler(ptr.Pointer(), C.int(objectType), core.PointerFromQObject(component))
	}
}
