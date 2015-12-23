package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QBuffer struct {
	QIODevice
}

type QBuffer_ITF interface {
	QIODevice_ITF
	QBuffer_PTR() *QBuffer
}

func PointerFromQBuffer(ptr QBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBuffer_PTR().Pointer()
	}
	return nil
}

func NewQBufferFromPointer(ptr unsafe.Pointer) *QBuffer {
	var n = new(QBuffer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBuffer_") {
		n.SetObjectName("QBuffer_" + qt.Identifier())
	}
	return n
}

func (ptr *QBuffer) QBuffer_PTR() *QBuffer {
	return ptr
}

func NewQBuffer2(byteArray QByteArray_ITF, parent QObject_ITF) *QBuffer {
	defer qt.Recovering("QBuffer::QBuffer")

	return NewQBufferFromPointer(C.QBuffer_NewQBuffer2(PointerFromQByteArray(byteArray), PointerFromQObject(parent)))
}

func NewQBuffer(parent QObject_ITF) *QBuffer {
	defer qt.Recovering("QBuffer::QBuffer")

	return NewQBufferFromPointer(C.QBuffer_NewQBuffer(PointerFromQObject(parent)))
}

func (ptr *QBuffer) AtEnd() bool {
	defer qt.Recovering("QBuffer::atEnd")

	if ptr.Pointer() != nil {
		return C.QBuffer_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBuffer) Buffer() *QByteArray {
	defer qt.Recovering("QBuffer::buffer")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Buffer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) Buffer2() *QByteArray {
	defer qt.Recovering("QBuffer::buffer")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Buffer2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) CanReadLine() bool {
	defer qt.Recovering("QBuffer::canReadLine")

	if ptr.Pointer() != nil {
		return C.QBuffer_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBuffer) ConnectClose(f func()) {
	defer qt.Recovering("connect QBuffer::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QBuffer) DisconnectClose() {
	defer qt.Recovering("disconnect QBuffer::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQBufferClose
func callbackQBufferClose(ptrName *C.char) bool {
	defer qt.Recovering("callback QBuffer::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QBuffer) Data() *QByteArray {
	defer qt.Recovering("QBuffer::data")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) Open(flags QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QBuffer::open")

	if ptr.Pointer() != nil {
		return C.QBuffer_Open(ptr.Pointer(), C.int(flags)) != 0
	}
	return false
}

func (ptr *QBuffer) Pos() int64 {
	defer qt.Recovering("QBuffer::pos")

	if ptr.Pointer() != nil {
		return int64(C.QBuffer_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBuffer) Seek(pos int64) bool {
	defer qt.Recovering("QBuffer::seek")

	if ptr.Pointer() != nil {
		return C.QBuffer_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QBuffer) SetBuffer(byteArray QByteArray_ITF) {
	defer qt.Recovering("QBuffer::setBuffer")

	if ptr.Pointer() != nil {
		C.QBuffer_SetBuffer(ptr.Pointer(), PointerFromQByteArray(byteArray))
	}
}

func (ptr *QBuffer) SetData(data QByteArray_ITF) {
	defer qt.Recovering("QBuffer::setData")

	if ptr.Pointer() != nil {
		C.QBuffer_SetData(ptr.Pointer(), PointerFromQByteArray(data))
	}
}

func (ptr *QBuffer) SetData2(data string, size int) {
	defer qt.Recovering("QBuffer::setData")

	if ptr.Pointer() != nil {
		C.QBuffer_SetData2(ptr.Pointer(), C.CString(data), C.int(size))
	}
}

func (ptr *QBuffer) Size() int64 {
	defer qt.Recovering("QBuffer::size")

	if ptr.Pointer() != nil {
		return int64(C.QBuffer_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBuffer) DestroyQBuffer() {
	defer qt.Recovering("QBuffer::~QBuffer")

	if ptr.Pointer() != nil {
		C.QBuffer_DestroyQBuffer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBuffer) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QBuffer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBuffer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBuffer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBufferTimerEvent
func callbackQBufferTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBuffer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QBuffer) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QBuffer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBuffer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBuffer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBufferChildEvent
func callbackQBufferChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBuffer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QBuffer) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QBuffer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBuffer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBuffer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBufferCustomEvent
func callbackQBufferCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBuffer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
