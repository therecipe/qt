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

type QSoundITF interface {
	core.QObjectITF
	QSoundPTR() *QSound
}

func PointerFromQSound(ptr QSoundITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSoundPTR().Pointer()
	}
	return nil
}

func QSoundFromPointer(ptr unsafe.Pointer) *QSound {
	var n = new(QSound)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSound_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSound) QSoundPTR() *QSound {
	return ptr
}

//QSound::Loop
type QSound__Loop int

var (
	QSound__Infinite = QSound__Loop(-1)
)

func (ptr *QSound) SetLoops(number int) {
	if ptr.Pointer() != nil {
		C.QSound_SetLoops(C.QtObjectPtr(ptr.Pointer()), C.int(number))
	}
}

func NewQSound(filename string, parent core.QObjectITF) *QSound {
	return QSoundFromPointer(unsafe.Pointer(C.QSound_NewQSound(C.CString(filename), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QSound) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSound_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSound) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QSound_IsFinished(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSound) Loops() int {
	if ptr.Pointer() != nil {
		return int(C.QSound_Loops(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSound) LoopsRemaining() int {
	if ptr.Pointer() != nil {
		return int(C.QSound_LoopsRemaining(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSound) Play2() {
	if ptr.Pointer() != nil {
		C.QSound_Play2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QSound_Play(filename string) {
	C.QSound_QSound_Play(C.CString(filename))
}

func (ptr *QSound) Stop() {
	if ptr.Pointer() != nil {
		C.QSound_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSound) DestroyQSound() {
	if ptr.Pointer() != nil {
		C.QSound_DestroyQSound(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
