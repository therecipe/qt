package core

//#include "qpauseanimation.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPauseAnimation struct {
	QAbstractAnimation
}

type QPauseAnimationITF interface {
	QAbstractAnimationITF
	QPauseAnimationPTR() *QPauseAnimation
}

func PointerFromQPauseAnimation(ptr QPauseAnimationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPauseAnimationPTR().Pointer()
	}
	return nil
}

func QPauseAnimationFromPointer(ptr unsafe.Pointer) *QPauseAnimation {
	var n = new(QPauseAnimation)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPauseAnimation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPauseAnimation) QPauseAnimationPTR() *QPauseAnimation {
	return ptr
}

func (ptr *QPauseAnimation) Duration() int {
	if ptr.Pointer() != nil {
		return int(C.QPauseAnimation_Duration(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPauseAnimation) SetDuration(msecs int) {
	if ptr.Pointer() != nil {
		C.QPauseAnimation_SetDuration(C.QtObjectPtr(ptr.Pointer()), C.int(msecs))
	}
}

func NewQPauseAnimation(parent QObjectITF) *QPauseAnimation {
	return QPauseAnimationFromPointer(unsafe.Pointer(C.QPauseAnimation_NewQPauseAnimation(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQPauseAnimation2(msecs int, parent QObjectITF) *QPauseAnimation {
	return QPauseAnimationFromPointer(unsafe.Pointer(C.QPauseAnimation_NewQPauseAnimation2(C.int(msecs), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QPauseAnimation) DestroyQPauseAnimation() {
	if ptr.Pointer() != nil {
		C.QPauseAnimation_DestroyQPauseAnimation(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
