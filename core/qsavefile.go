package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSaveFile struct {
	QFileDevice
}

type QSaveFile_ITF interface {
	QFileDevice_ITF
	QSaveFile_PTR() *QSaveFile
}

func PointerFromQSaveFile(ptr QSaveFile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSaveFile_PTR().Pointer()
	}
	return nil
}

func NewQSaveFileFromPointer(ptr unsafe.Pointer) *QSaveFile {
	var n = new(QSaveFile)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSaveFile_") {
		n.SetObjectName("QSaveFile_" + qt.Identifier())
	}
	return n
}

func (ptr *QSaveFile) QSaveFile_PTR() *QSaveFile {
	return ptr
}

func NewQSaveFile2(parent QObject_ITF) *QSaveFile {
	defer qt.Recovering("QSaveFile::QSaveFile")

	return NewQSaveFileFromPointer(C.QSaveFile_NewQSaveFile2(PointerFromQObject(parent)))
}

func NewQSaveFile(name string) *QSaveFile {
	defer qt.Recovering("QSaveFile::QSaveFile")

	return NewQSaveFileFromPointer(C.QSaveFile_NewQSaveFile(C.CString(name)))
}

func NewQSaveFile3(name string, parent QObject_ITF) *QSaveFile {
	defer qt.Recovering("QSaveFile::QSaveFile")

	return NewQSaveFileFromPointer(C.QSaveFile_NewQSaveFile3(C.CString(name), PointerFromQObject(parent)))
}

func (ptr *QSaveFile) CancelWriting() {
	defer qt.Recovering("QSaveFile::cancelWriting")

	if ptr.Pointer() != nil {
		C.QSaveFile_CancelWriting(ptr.Pointer())
	}
}

func (ptr *QSaveFile) Commit() bool {
	defer qt.Recovering("QSaveFile::commit")

	if ptr.Pointer() != nil {
		return C.QSaveFile_Commit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSaveFile) DirectWriteFallback() bool {
	defer qt.Recovering("QSaveFile::directWriteFallback")

	if ptr.Pointer() != nil {
		return C.QSaveFile_DirectWriteFallback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSaveFile) FileName() string {
	defer qt.Recovering("QSaveFile::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSaveFile_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSaveFile) Open(mode QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QSaveFile::open")

	if ptr.Pointer() != nil {
		return C.QSaveFile_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSaveFile) SetDirectWriteFallback(enabled bool) {
	defer qt.Recovering("QSaveFile::setDirectWriteFallback")

	if ptr.Pointer() != nil {
		C.QSaveFile_SetDirectWriteFallback(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QSaveFile) SetFileName(name string) {
	defer qt.Recovering("QSaveFile::setFileName")

	if ptr.Pointer() != nil {
		C.QSaveFile_SetFileName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSaveFile) WriteData(data string, len int64) int64 {
	defer qt.Recovering("QSaveFile::writeData")

	if ptr.Pointer() != nil {
		return int64(C.QSaveFile_WriteData(ptr.Pointer(), C.CString(data), C.longlong(len)))
	}
	return 0
}

func (ptr *QSaveFile) DestroyQSaveFile() {
	defer qt.Recovering("QSaveFile::~QSaveFile")

	if ptr.Pointer() != nil {
		C.QSaveFile_DestroyQSaveFile(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSaveFile) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QSaveFile::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSaveFile) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSaveFile::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSaveFileTimerEvent
func callbackQSaveFileTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSaveFile::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQSaveFileFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSaveFile) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QSaveFile::timerEvent")

	if ptr.Pointer() != nil {
		C.QSaveFile_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSaveFile) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QSaveFile::timerEvent")

	if ptr.Pointer() != nil {
		C.QSaveFile_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSaveFile) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QSaveFile::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSaveFile) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSaveFile::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSaveFileChildEvent
func callbackQSaveFileChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSaveFile::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQSaveFileFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QSaveFile) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QSaveFile::childEvent")

	if ptr.Pointer() != nil {
		C.QSaveFile_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSaveFile) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QSaveFile::childEvent")

	if ptr.Pointer() != nil {
		C.QSaveFile_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSaveFile) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QSaveFile::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSaveFile) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSaveFile::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSaveFileCustomEvent
func callbackQSaveFileCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSaveFile::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQSaveFileFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QSaveFile) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QSaveFile::customEvent")

	if ptr.Pointer() != nil {
		C.QSaveFile_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QSaveFile) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QSaveFile::customEvent")

	if ptr.Pointer() != nil {
		C.QSaveFile_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
