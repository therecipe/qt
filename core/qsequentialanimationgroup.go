package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QSequentialAnimationGroup struct {
	QAnimationGroup
}

type QSequentialAnimationGroup_ITF interface {
	QAnimationGroup_ITF
	QSequentialAnimationGroup_PTR() *QSequentialAnimationGroup
}

func PointerFromQSequentialAnimationGroup(ptr QSequentialAnimationGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSequentialAnimationGroup_PTR().Pointer()
	}
	return nil
}

func NewQSequentialAnimationGroupFromPointer(ptr unsafe.Pointer) *QSequentialAnimationGroup {
	var n = new(QSequentialAnimationGroup)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSequentialAnimationGroup_") {
		n.SetObjectName("QSequentialAnimationGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSequentialAnimationGroup) QSequentialAnimationGroup_PTR() *QSequentialAnimationGroup {
	return ptr
}

func (ptr *QSequentialAnimationGroup) CurrentAnimation() *QAbstractAnimation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSequentialAnimationGroup::currentAnimation")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractAnimationFromPointer(C.QSequentialAnimationGroup_CurrentAnimation(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) AddPause(msecs int) *QPauseAnimation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSequentialAnimationGroup::addPause")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQPauseAnimationFromPointer(C.QSequentialAnimationGroup_AddPause(ptr.Pointer(), C.int(msecs)))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) ConnectCurrentAnimationChanged(f func(current *QAbstractAnimation)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSequentialAnimationGroup::currentAnimationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_ConnectCurrentAnimationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentAnimationChanged", f)
	}
}

func (ptr *QSequentialAnimationGroup) DisconnectCurrentAnimationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSequentialAnimationGroup::currentAnimationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_DisconnectCurrentAnimationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentAnimationChanged")
	}
}

//export callbackQSequentialAnimationGroupCurrentAnimationChanged
func callbackQSequentialAnimationGroupCurrentAnimationChanged(ptrName *C.char, current unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSequentialAnimationGroup::currentAnimationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentAnimationChanged").(func(*QAbstractAnimation))(NewQAbstractAnimationFromPointer(current))
}

func (ptr *QSequentialAnimationGroup) Duration() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSequentialAnimationGroup::duration")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSequentialAnimationGroup_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSequentialAnimationGroup) InsertPause(index int, msecs int) *QPauseAnimation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSequentialAnimationGroup::insertPause")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQPauseAnimationFromPointer(C.QSequentialAnimationGroup_InsertPause(ptr.Pointer(), C.int(index), C.int(msecs)))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) DestroyQSequentialAnimationGroup() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSequentialAnimationGroup::~QSequentialAnimationGroup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_DestroyQSequentialAnimationGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
