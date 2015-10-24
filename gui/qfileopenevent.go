package gui

//#include "qfileopenevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QFileOpenEvent struct {
	core.QEvent
}

type QFileOpenEventITF interface {
	core.QEventITF
	QFileOpenEventPTR() *QFileOpenEvent
}

func PointerFromQFileOpenEvent(ptr QFileOpenEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileOpenEventPTR().Pointer()
	}
	return nil
}

func QFileOpenEventFromPointer(ptr unsafe.Pointer) *QFileOpenEvent {
	var n = new(QFileOpenEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFileOpenEvent) QFileOpenEventPTR() *QFileOpenEvent {
	return ptr
}

func (ptr *QFileOpenEvent) OpenFile(file core.QFileITF, flags core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QFileOpenEvent_OpenFile(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQFile(file)), C.int(flags)) != 0
	}
	return false
}

func (ptr *QFileOpenEvent) File() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileOpenEvent_File(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileOpenEvent) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileOpenEvent_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
