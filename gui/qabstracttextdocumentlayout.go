package gui

//#include "qabstracttextdocumentlayout.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractTextDocumentLayout struct {
	core.QObject
}

type QAbstractTextDocumentLayoutITF interface {
	core.QObjectITF
	QAbstractTextDocumentLayoutPTR() *QAbstractTextDocumentLayout
}

func PointerFromQAbstractTextDocumentLayout(ptr QAbstractTextDocumentLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractTextDocumentLayoutPTR().Pointer()
	}
	return nil
}

func QAbstractTextDocumentLayoutFromPointer(ptr unsafe.Pointer) *QAbstractTextDocumentLayout {
	var n = new(QAbstractTextDocumentLayout)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractTextDocumentLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractTextDocumentLayout) QAbstractTextDocumentLayoutPTR() *QAbstractTextDocumentLayout {
	return ptr
}

func (ptr *QAbstractTextDocumentLayout) AnchorAt(position core.QPointFITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractTextDocumentLayout_AnchorAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(position))))
	}
	return ""
}

func (ptr *QAbstractTextDocumentLayout) Document() *QTextDocument {
	if ptr.Pointer() != nil {
		return QTextDocumentFromPointer(unsafe.Pointer(C.QAbstractTextDocumentLayout_Document(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractTextDocumentLayout) HandlerForObject(objectType int) *QTextObjectInterface {
	if ptr.Pointer() != nil {
		return QTextObjectInterfaceFromPointer(unsafe.Pointer(C.QAbstractTextDocumentLayout_HandlerForObject(C.QtObjectPtr(ptr.Pointer()), C.int(objectType))))
	}
	return nil
}

func (ptr *QAbstractTextDocumentLayout) PageCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractTextDocumentLayout_PageCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractTextDocumentLayout) ConnectPageCountChanged(f func(newPages int)) {
	if ptr.Pointer() != nil {
		C.QAbstractTextDocumentLayout_ConnectPageCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "pageCountChanged", f)
	}
}

func (ptr *QAbstractTextDocumentLayout) DisconnectPageCountChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractTextDocumentLayout_DisconnectPageCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "pageCountChanged")
	}
}

//export callbackQAbstractTextDocumentLayoutPageCountChanged
func callbackQAbstractTextDocumentLayoutPageCountChanged(ptrName *C.char, newPages C.int) {
	qt.GetSignal(C.GoString(ptrName), "pageCountChanged").(func(int))(int(newPages))
}

func (ptr *QAbstractTextDocumentLayout) PaintDevice() *QPaintDevice {
	if ptr.Pointer() != nil {
		return QPaintDeviceFromPointer(unsafe.Pointer(C.QAbstractTextDocumentLayout_PaintDevice(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractTextDocumentLayout) RegisterHandler(objectType int, component core.QObjectITF) {
	if ptr.Pointer() != nil {
		C.QAbstractTextDocumentLayout_RegisterHandler(C.QtObjectPtr(ptr.Pointer()), C.int(objectType), C.QtObjectPtr(core.PointerFromQObject(component)))
	}
}

func (ptr *QAbstractTextDocumentLayout) SetPaintDevice(device QPaintDeviceITF) {
	if ptr.Pointer() != nil {
		C.QAbstractTextDocumentLayout_SetPaintDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPaintDevice(device)))
	}
}

func (ptr *QAbstractTextDocumentLayout) UnregisterHandler(objectType int, component core.QObjectITF) {
	if ptr.Pointer() != nil {
		C.QAbstractTextDocumentLayout_UnregisterHandler(C.QtObjectPtr(ptr.Pointer()), C.int(objectType), C.QtObjectPtr(core.PointerFromQObject(component)))
	}
}
