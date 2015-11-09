package multimedia

//#include "qsound.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSound struct {
	core.QObject
}

type QSound_ITF interface {
	core.QObject_ITF
	QSound_PTR() *QSound
}

func PointerFromQSound(ptr QSound_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSound_PTR().Pointer()
	}
	return nil
}

func NewQSoundFromPointer(ptr unsafe.Pointer) *QSound {
	var n = new(QSound)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QSound_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSound) QSound_PTR() *QSound {
	return ptr
}

//QSound::Loop
type QSound__Loop int64

const (
	QSound__Infinite = QSound__Loop(-1)
)

func (ptr *QSound) SetLoops(number int) {
	if ptr.Pointer() != nil {
		C.QSound_SetLoops(ptr.Pointer(), C.int(number))
	}
}

func NewQSound(filename string, parent core.QObject_ITF) *QSound {
	return NewQSoundFromPointer(C.QSound_NewQSound(C.CString(filename), core.PointerFromQObject(parent)))
}

func (ptr *QSound) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSound_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSound) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QSound_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSound) Loops() int {
	if ptr.Pointer() != nil {
		return int(C.QSound_Loops(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSound) LoopsRemaining() int {
	if ptr.Pointer() != nil {
		return int(C.QSound_LoopsRemaining(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSound) Play2() {
	if ptr.Pointer() != nil {
		C.QSound_Play2(ptr.Pointer())
	}
}

func QSound_Play(filename string) {
	C.QSound_QSound_Play(C.CString(filename))
}

func (ptr *QSound) Stop() {
	if ptr.Pointer() != nil {
		C.QSound_Stop(ptr.Pointer())
	}
}

func (ptr *QSound) DestroyQSound() {
	if ptr.Pointer() != nil {
		C.QSound_DestroyQSound(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
