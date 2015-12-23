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
func callbackQSaveFileTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSaveFile::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSaveFileChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSaveFile::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSaveFileCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSaveFile::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
