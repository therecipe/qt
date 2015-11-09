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

type QClipboard_ITF interface {
	core.QObject_ITF
	QClipboard_PTR() *QClipboard
}

func PointerFromQClipboard(ptr QClipboard_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QClipboard_PTR().Pointer()
	}
	return nil
}

func NewQClipboardFromPointer(ptr unsafe.Pointer) *QClipboard {
	var n = new(QClipboard)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QClipboard_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QClipboard) QClipboard_PTR() *QClipboard {
	return ptr
}

//QClipboard::Mode
type QClipboard__Mode int64

const (
	QClipboard__Clipboard  = QClipboard__Mode(0)
	QClipboard__Selection  = QClipboard__Mode(1)
	QClipboard__FindBuffer = QClipboard__Mode(2)
	QClipboard__LastMode   = QClipboard__Mode(QClipboard__FindBuffer)
)

func (ptr *QClipboard) Clear(mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_Clear(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QClipboard) MimeData(mode QClipboard__Mode) *core.QMimeData {
	if ptr.Pointer() != nil {
		return core.NewQMimeDataFromPointer(C.QClipboard_MimeData(ptr.Pointer(), C.int(mode)))
	}
	return nil
}

func (ptr *QClipboard) SetMimeData(src core.QMimeData_ITF, mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_SetMimeData(ptr.Pointer(), core.PointerFromQMimeData(src), C.int(mode))
	}
}

func (ptr *QClipboard) ConnectChanged(f func(mode QClipboard__Mode)) {
	if ptr.Pointer() != nil {
		C.QClipboard_ConnectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "changed", f)
	}
}

func (ptr *QClipboard) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "changed")
	}
}

//export callbackQClipboardChanged
func callbackQClipboardChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "changed").(func(QClipboard__Mode))(QClipboard__Mode(mode))
}

func (ptr *QClipboard) ConnectDataChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QClipboard_ConnectDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dataChanged", f)
	}
}

func (ptr *QClipboard) DisconnectDataChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dataChanged")
	}
}

//export callbackQClipboardDataChanged
func callbackQClipboardDataChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "dataChanged").(func())()
}

func (ptr *QClipboard) ConnectFindBufferChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QClipboard_ConnectFindBufferChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "findBufferChanged", f)
	}
}

func (ptr *QClipboard) DisconnectFindBufferChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectFindBufferChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "findBufferChanged")
	}
}

//export callbackQClipboardFindBufferChanged
func callbackQClipboardFindBufferChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "findBufferChanged").(func())()
}

func (ptr *QClipboard) OwnsClipboard() bool {
	if ptr.Pointer() != nil {
		return C.QClipboard_OwnsClipboard(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QClipboard) OwnsFindBuffer() bool {
	if ptr.Pointer() != nil {
		return C.QClipboard_OwnsFindBuffer(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QClipboard) OwnsSelection() bool {
	if ptr.Pointer() != nil {
		return C.QClipboard_OwnsSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QClipboard) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QClipboard_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QClipboard) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQClipboardSelectionChanged
func callbackQClipboardSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QClipboard) SetImage(image QImage_ITF, mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_SetImage(ptr.Pointer(), PointerFromQImage(image), C.int(mode))
	}
}

func (ptr *QClipboard) SetPixmap(pixmap QPixmap_ITF, mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_SetPixmap(ptr.Pointer(), PointerFromQPixmap(pixmap), C.int(mode))
	}
}

func (ptr *QClipboard) SetText(text string, mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_SetText(ptr.Pointer(), C.CString(text), C.int(mode))
	}
}

func (ptr *QClipboard) SupportsFindBuffer() bool {
	if ptr.Pointer() != nil {
		return C.QClipboard_SupportsFindBuffer(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QClipboard) SupportsSelection() bool {
	if ptr.Pointer() != nil {
		return C.QClipboard_SupportsSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QClipboard) Text(mode QClipboard__Mode) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QClipboard_Text(ptr.Pointer(), C.int(mode)))
	}
	return ""
}
