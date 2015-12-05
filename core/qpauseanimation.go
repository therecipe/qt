package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QPauseAnimation struct {
	QAbstractAnimation
}

type QPauseAnimation_ITF interface {
	QAbstractAnimation_ITF
	QPauseAnimation_PTR() *QPauseAnimation
}

func PointerFromQPauseAnimation(ptr QPauseAnimation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPauseAnimation_PTR().Pointer()
	}
	return nil
}

func NewQPauseAnimationFromPointer(ptr unsafe.Pointer) *QPauseAnimation {
	var n = new(QPauseAnimation)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPauseAnimation_") {
		n.SetObjectName("QPauseAnimation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPauseAnimation) QPauseAnimation_PTR() *QPauseAnimation {
	return ptr
}

func (ptr *QPauseAnimation) Duration() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPauseAnimation::duration")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QPauseAnimation_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPauseAnimation) SetDuration(msecs int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPauseAnimation::setDuration")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPauseAnimation_SetDuration(ptr.Pointer(), C.int(msecs))
	}
}

func NewQPauseAnimation(parent QObject_ITF) *QPauseAnimation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPauseAnimation::QPauseAnimation")
		}
	}()

	return NewQPauseAnimationFromPointer(C.QPauseAnimation_NewQPauseAnimation(PointerFromQObject(parent)))
}

func NewQPauseAnimation2(msecs int, parent QObject_ITF) *QPauseAnimation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPauseAnimation::QPauseAnimation")
		}
	}()

	return NewQPauseAnimationFromPointer(C.QPauseAnimation_NewQPauseAnimation2(C.int(msecs), PointerFromQObject(parent)))
}

func (ptr *QPauseAnimation) DestroyQPauseAnimation() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPauseAnimation::~QPauseAnimation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPauseAnimation_DestroyQPauseAnimation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
