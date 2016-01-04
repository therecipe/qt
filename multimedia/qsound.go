package multimedia

//#include "multimedia.h"
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
	for len(n.ObjectName()) < len("QSound_") {
		n.SetObjectName("QSound_" + qt.Identifier())
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
	defer qt.Recovering("QSound::setLoops")

	if ptr.Pointer() != nil {
		C.QSound_SetLoops(ptr.Pointer(), C.int(number))
	}
}

func NewQSound(filename string, parent core.QObject_ITF) *QSound {
	defer qt.Recovering("QSound::QSound")

	return NewQSoundFromPointer(C.QSound_NewQSound(C.CString(filename), core.PointerFromQObject(parent)))
}

func (ptr *QSound) FileName() string {
	defer qt.Recovering("QSound::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSound_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSound) IsFinished() bool {
	defer qt.Recovering("QSound::isFinished")

	if ptr.Pointer() != nil {
		return C.QSound_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSound) Loops() int {
	defer qt.Recovering("QSound::loops")

	if ptr.Pointer() != nil {
		return int(C.QSound_Loops(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSound) LoopsRemaining() int {
	defer qt.Recovering("QSound::loopsRemaining")

	if ptr.Pointer() != nil {
		return int(C.QSound_LoopsRemaining(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSound) Play2() {
	defer qt.Recovering("QSound::play")

	if ptr.Pointer() != nil {
		C.QSound_Play2(ptr.Pointer())
	}
}

func QSound_Play(filename string) {
	defer qt.Recovering("QSound::play")

	C.QSound_QSound_Play(C.CString(filename))
}

func (ptr *QSound) Stop() {
	defer qt.Recovering("QSound::stop")

	if ptr.Pointer() != nil {
		C.QSound_Stop(ptr.Pointer())
	}
}

func (ptr *QSound) DestroyQSound() {
	defer qt.Recovering("QSound::~QSound")

	if ptr.Pointer() != nil {
		C.QSound_DestroyQSound(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSound) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSound::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSound) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSound::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSoundTimerEvent
func callbackQSoundTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSound::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSoundFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSound) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSound::timerEvent")

	if ptr.Pointer() != nil {
		C.QSound_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSound) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSound::timerEvent")

	if ptr.Pointer() != nil {
		C.QSound_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSound) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSound::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSound) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSound::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSoundChildEvent
func callbackQSoundChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSound::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSoundFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSound) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSound::childEvent")

	if ptr.Pointer() != nil {
		C.QSound_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSound) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSound::childEvent")

	if ptr.Pointer() != nil {
		C.QSound_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSound) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSound::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSound) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSound::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSoundCustomEvent
func callbackQSoundCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSound::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSoundFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSound) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSound::customEvent")

	if ptr.Pointer() != nil {
		C.QSound_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSound) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSound::customEvent")

	if ptr.Pointer() != nil {
		C.QSound_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
