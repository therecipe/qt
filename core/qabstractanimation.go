package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QAbstractAnimation struct {
	QObject
}

type QAbstractAnimation_ITF interface {
	QObject_ITF
	QAbstractAnimation_PTR() *QAbstractAnimation
}

func PointerFromQAbstractAnimation(ptr QAbstractAnimation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractAnimation_PTR().Pointer()
	}
	return nil
}

func NewQAbstractAnimationFromPointer(ptr unsafe.Pointer) *QAbstractAnimation {
	var n = new(QAbstractAnimation)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractAnimation_") {
		n.SetObjectName("QAbstractAnimation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractAnimation) QAbstractAnimation_PTR() *QAbstractAnimation {
	return ptr
}

//QAbstractAnimation::DeletionPolicy
type QAbstractAnimation__DeletionPolicy int64

const (
	QAbstractAnimation__KeepWhenStopped   = QAbstractAnimation__DeletionPolicy(0)
	QAbstractAnimation__DeleteWhenStopped = QAbstractAnimation__DeletionPolicy(1)
)

//QAbstractAnimation::Direction
type QAbstractAnimation__Direction int64

const (
	QAbstractAnimation__Forward  = QAbstractAnimation__Direction(0)
	QAbstractAnimation__Backward = QAbstractAnimation__Direction(1)
)

//QAbstractAnimation::State
type QAbstractAnimation__State int64

const (
	QAbstractAnimation__Stopped = QAbstractAnimation__State(0)
	QAbstractAnimation__Paused  = QAbstractAnimation__State(1)
	QAbstractAnimation__Running = QAbstractAnimation__State(2)
)

func (ptr *QAbstractAnimation) CurrentLoop() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::currentLoop")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_CurrentLoop(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) CurrentTime() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::currentTime")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_CurrentTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) Direction() QAbstractAnimation__Direction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::direction")
		}
	}()

	if ptr.Pointer() != nil {
		return QAbstractAnimation__Direction(C.QAbstractAnimation_Direction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) LoopCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::loopCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) SetCurrentTime(msecs int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::setCurrentTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetCurrentTime(ptr.Pointer(), C.int(msecs))
	}
}

func (ptr *QAbstractAnimation) SetDirection(direction QAbstractAnimation__Direction) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::setDirection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QAbstractAnimation) SetLoopCount(loopCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::setLoopCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetLoopCount(ptr.Pointer(), C.int(loopCount))
	}
}

func (ptr *QAbstractAnimation) ConnectCurrentLoopChanged(f func(currentLoop int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::currentLoopChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectCurrentLoopChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentLoopChanged", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectCurrentLoopChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::currentLoopChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectCurrentLoopChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentLoopChanged")
	}
}

//export callbackQAbstractAnimationCurrentLoopChanged
func callbackQAbstractAnimationCurrentLoopChanged(ptrName *C.char, currentLoop C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::currentLoopChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentLoopChanged").(func(int))(int(currentLoop))
}

func (ptr *QAbstractAnimation) CurrentLoopTime() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::currentLoopTime")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_CurrentLoopTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) ConnectDirectionChanged(f func(newDirection QAbstractAnimation__Direction)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::directionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectDirectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "directionChanged", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectDirectionChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::directionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectDirectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "directionChanged")
	}
}

//export callbackQAbstractAnimationDirectionChanged
func callbackQAbstractAnimationDirectionChanged(ptrName *C.char, newDirection C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::directionChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "directionChanged").(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(newDirection))
}

func (ptr *QAbstractAnimation) Duration() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::duration")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) ConnectFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAbstractAnimationFinished
func callbackQAbstractAnimationFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QAbstractAnimation) Group() *QAnimationGroup {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::group")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAnimationGroupFromPointer(C.QAbstractAnimation_Group(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractAnimation) Pause() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::pause")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Pause(ptr.Pointer())
	}
}

func (ptr *QAbstractAnimation) Resume() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::resume")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Resume(ptr.Pointer())
	}
}

func (ptr *QAbstractAnimation) SetPaused(paused bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::setPaused")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetPaused(ptr.Pointer(), C.int(qt.GoBoolToInt(paused)))
	}
}

func (ptr *QAbstractAnimation) Start(policy QAbstractAnimation__DeletionPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::start")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Start(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QAbstractAnimation) ConnectStateChanged(f func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAbstractAnimationStateChanged
func callbackQAbstractAnimationStateChanged(ptrName *C.char, newState C.int, oldState C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
}

func (ptr *QAbstractAnimation) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Stop(ptr.Pointer())
	}
}

func (ptr *QAbstractAnimation) TotalDuration() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::totalDuration")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_TotalDuration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) DestroyQAbstractAnimation() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractAnimation::~QAbstractAnimation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DestroyQAbstractAnimation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
