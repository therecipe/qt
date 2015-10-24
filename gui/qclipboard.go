package gui

//#include "qclipboard.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QClipboard struct {
	core.QObject
}

type QClipboardITF interface {
	core.QObjectITF
	QClipboardPTR() *QClipboard
}

func PointerFromQClipboard(ptr QClipboardITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QClipboardPTR().Pointer()
	}
	return nil
}

func QClipboardFromPointer(ptr unsafe.Pointer) *QClipboard {
	var n = new(QClipboard)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QClipboard_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QClipboard) QClipboardPTR() *QClipboard {
	return ptr
}

//QClipboard::Mode
type QClipboard__Mode int

var (
	QClipboard__Clipboard  = QClipboard__Mode(0)
	QClipboard__Selection  = QClipboard__Mode(1)
	QClipboard__FindBuffer = QClipboard__Mode(2)
	QClipboard__LastMode   = QClipboard__Mode(QClipboard__FindBuffer)
)

func (ptr *QClipboard) Clear(mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_Clear(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QClipboard) MimeData(mode QClipboard__Mode) *core.QMimeData {
	if ptr.Pointer() != nil {
		return core.QMimeDataFromPointer(unsafe.Pointer(C.QClipboard_MimeData(C.QtObjectPtr(ptr.Pointer()), C.int(mode))))
	}
	return nil
}

func (ptr *QClipboard) SetMimeData(src core.QMimeDataITF, mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_SetMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMimeData(src)), C.int(mode))
	}
}

func (ptr *QClipboard) ConnectChanged(f func(mode QClipboard__Mode)) {
	if ptr.Pointer() != nil {
		C.QClipboard_ConnectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "changed", f)
	}
}

func (ptr *QClipboard) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "changed")
	}
}

//export callbackQClipboardChanged
func callbackQClipboardChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "changed").(func(QClipboard__Mode))(QClipboard__Mode(mode))
}

func (ptr *QClipboard) ConnectDataChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QClipboard_ConnectDataChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "dataChanged", f)
	}
}

func (ptr *QClipboard) DisconnectDataChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectDataChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "dataChanged")
	}
}

//export callbackQClipboardDataChanged
func callbackQClipboardDataChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "dataChanged").(func())()
}

func (ptr *QClipboard) ConnectFindBufferChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QClipboard_ConnectFindBufferChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "findBufferChanged", f)
	}
}

func (ptr *QClipboard) DisconnectFindBufferChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectFindBufferChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "findBufferChanged")
	}
}

//export callbackQClipboardFindBufferChanged
func callbackQClipboardFindBufferChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "findBufferChanged").(func())()
}

func (ptr *QClipboard) OwnsClipboard() bool {
	if ptr.Pointer() != nil {
		return C.QClipboard_OwnsClipboard(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QClipboard) OwnsFindBuffer() bool {
	if ptr.Pointer() != nil {
		return C.QClipboard_OwnsFindBuffer(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QClipboard) OwnsSelection() bool {
	if ptr.Pointer() != nil {
		return C.QClipboard_OwnsSelection(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QClipboard) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QClipboard_ConnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QClipboard) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQClipboardSelectionChanged
func callbackQClipboardSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QClipboard) SetImage(image QImageITF, mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_SetImage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQImage(image)), C.int(mode))
	}
}

func (ptr *QClipboard) SetPixmap(pixmap QPixmapITF, mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_SetPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.int(mode))
	}
}

func (ptr *QClipboard) SetText(text string, mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.int(mode))
	}
}

func (ptr *QClipboard) SupportsFindBuffer() bool {
	if ptr.Pointer() != nil {
		return C.QClipboard_SupportsFindBuffer(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QClipboard) SupportsSelection() bool {
	if ptr.Pointer() != nil {
		return C.QClipboard_SupportsSelection(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QClipboard) Text(mode QClipboard__Mode) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QClipboard_Text(C.QtObjectPtr(ptr.Pointer()), C.int(mode)))
	}
	return ""
}
