package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QFileOpenEvent struct {
	core.QEvent
}

type QFileOpenEvent_ITF interface {
	core.QEvent_ITF
	QFileOpenEvent_PTR() *QFileOpenEvent
}

func PointerFromQFileOpenEvent(ptr QFileOpenEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileOpenEvent_PTR().Pointer()
	}
	return nil
}

func NewQFileOpenEventFromPointer(ptr unsafe.Pointer) *QFileOpenEvent {
	var n = new(QFileOpenEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFileOpenEvent) QFileOpenEvent_PTR() *QFileOpenEvent {
	return ptr
}

func (ptr *QFileOpenEvent) OpenFile(file core.QFile_ITF, flags core.QIODevice__OpenModeFlag) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileOpenEvent::openFile")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileOpenEvent_OpenFile(ptr.Pointer(), core.PointerFromQFile(file), C.int(flags)) != 0
	}
	return false
}

func (ptr *QFileOpenEvent) File() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileOpenEvent::file")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileOpenEvent_File(ptr.Pointer()))
	}
	return ""
}
