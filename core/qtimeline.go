package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTimeLine struct {
	QObject
}

type QTimeLine_ITF interface {
	QObject_ITF
	QTimeLine_PTR() *QTimeLine
}

func PointerFromQTimeLine(ptr QTimeLine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimeLine_PTR().Pointer()
	}
	return nil
}

func NewQTimeLineFromPointer(ptr unsafe.Pointer) *QTimeLine {
	var n = new(QTimeLine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTimeLine_") {
		n.SetObjectName("QTimeLine_" + qt.Identifier())
	}
	return n
}

func (ptr *QTimeLine) QTimeLine_PTR() *QTimeLine {
	return ptr
}

//QTimeLine::CurveShape
type QTimeLine__CurveShape int64

const (
	QTimeLine__EaseInCurve    = QTimeLine__CurveShape(0)
	QTimeLine__EaseOutCurve   = QTimeLine__CurveShape(1)
	QTimeLine__EaseInOutCurve = QTimeLine__CurveShape(2)
	QTimeLine__LinearCurve    = QTimeLine__CurveShape(3)
	QTimeLine__SineCurve      = QTimeLine__CurveShape(4)
	QTimeLine__CosineCurve    = QTimeLine__CurveShape(5)
)

//QTimeLine::Direction
type QTimeLine__Direction int64

const (
	QTimeLine__Forward  = QTimeLine__Direction(0)
	QTimeLine__Backward = QTimeLine__Direction(1)
)

//QTimeLine::State
type QTimeLine__State int64

const (
	QTimeLine__NotRunning = QTimeLine__State(0)
	QTimeLine__Paused     = QTimeLine__State(1)
	QTimeLine__Running    = QTimeLine__State(2)
)

func (ptr *QTimeLine) CurrentTime() int {
	defer qt.Recovering("QTimeLine::currentTime")

	if ptr.Pointer() != nil {
		return int(C.QTimeLine_CurrentTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) CurveShape() QTimeLine__CurveShape {
	defer qt.Recovering("QTimeLine::curveShape")

	if ptr.Pointer() != nil {
		return QTimeLine__CurveShape(C.QTimeLine_CurveShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) Direction() QTimeLine__Direction {
	defer qt.Recovering("QTimeLine::direction")

	if ptr.Pointer() != nil {
		return QTimeLine__Direction(C.QTimeLine_Direction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) Duration() int {
	defer qt.Recovering("QTimeLine::duration")

	if ptr.Pointer() != nil {
		return int(C.QTimeLine_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) EasingCurve() *QEasingCurve {
	defer qt.Recovering("QTimeLine::easingCurve")

	if ptr.Pointer() != nil {
		return NewQEasingCurveFromPointer(C.QTimeLine_EasingCurve(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTimeLine) LoopCount() int {
	defer qt.Recovering("QTimeLine::loopCount")

	if ptr.Pointer() != nil {
		return int(C.QTimeLine_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) SetCurrentTime(msec int) {
	defer qt.Recovering("QTimeLine::setCurrentTime")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetCurrentTime(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QTimeLine) SetCurveShape(shape QTimeLine__CurveShape) {
	defer qt.Recovering("QTimeLine::setCurveShape")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetCurveShape(ptr.Pointer(), C.int(shape))
	}
}

func (ptr *QTimeLine) SetDirection(direction QTimeLine__Direction) {
	defer qt.Recovering("QTimeLine::setDirection")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QTimeLine) SetDuration(duration int) {
	defer qt.Recovering("QTimeLine::setDuration")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetDuration(ptr.Pointer(), C.int(duration))
	}
}

func (ptr *QTimeLine) SetEasingCurve(curve QEasingCurve_ITF) {
	defer qt.Recovering("QTimeLine::setEasingCurve")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetEasingCurve(ptr.Pointer(), PointerFromQEasingCurve(curve))
	}
}

func (ptr *QTimeLine) SetLoopCount(count int) {
	defer qt.Recovering("QTimeLine::setLoopCount")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetLoopCount(ptr.Pointer(), C.int(count))
	}
}

func (ptr *QTimeLine) SetUpdateInterval(interval int) {
	defer qt.Recovering("QTimeLine::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetUpdateInterval(ptr.Pointer(), C.int(interval))
	}
}

func (ptr *QTimeLine) UpdateInterval() int {
	defer qt.Recovering("QTimeLine::updateInterval")

	if ptr.Pointer() != nil {
		return int(C.QTimeLine_UpdateInterval(ptr.Pointer()))
	}
	return 0
}

func NewQTimeLine(duration int, parent QObject_ITF) *QTimeLine {
	defer qt.Recovering("QTimeLine::QTimeLine")

	return NewQTimeLineFromPointer(C.QTimeLine_NewQTimeLine(C.int(duration), PointerFromQObject(parent)))
}

func (ptr *QTimeLine) CurrentFrame() int {
	defer qt.Recovering("QTimeLine::currentFrame")

	if ptr.Pointer() != nil {
		return int(C.QTimeLine_CurrentFrame(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) CurrentValue() float64 {
	defer qt.Recovering("QTimeLine::currentValue")

	if ptr.Pointer() != nil {
		return float64(C.QTimeLine_CurrentValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) EndFrame() int {
	defer qt.Recovering("QTimeLine::endFrame")

	if ptr.Pointer() != nil {
		return int(C.QTimeLine_EndFrame(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) ConnectFinished(f func()) {
	defer qt.Recovering("connect QTimeLine::finished")

	if ptr.Pointer() != nil {
		C.QTimeLine_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QTimeLine) DisconnectFinished() {
	defer qt.Recovering("disconnect QTimeLine::finished")

	if ptr.Pointer() != nil {
		C.QTimeLine_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQTimeLineFinished
func callbackQTimeLineFinished(ptrName *C.char) {
	defer qt.Recovering("callback QTimeLine::finished")

	var signal = qt.GetSignal(C.GoString(ptrName), "finished")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QTimeLine) ConnectFrameChanged(f func(frame int)) {
	defer qt.Recovering("connect QTimeLine::frameChanged")

	if ptr.Pointer() != nil {
		C.QTimeLine_ConnectFrameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frameChanged", f)
	}
}

func (ptr *QTimeLine) DisconnectFrameChanged() {
	defer qt.Recovering("disconnect QTimeLine::frameChanged")

	if ptr.Pointer() != nil {
		C.QTimeLine_DisconnectFrameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frameChanged")
	}
}

//export callbackQTimeLineFrameChanged
func callbackQTimeLineFrameChanged(ptrName *C.char, frame C.int) {
	defer qt.Recovering("callback QTimeLine::frameChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "frameChanged")
	if signal != nil {
		signal.(func(int))(int(frame))
	}

}

func (ptr *QTimeLine) FrameForTime(msec int) int {
	defer qt.Recovering("QTimeLine::frameForTime")

	if ptr.Pointer() != nil {
		return int(C.QTimeLine_FrameForTime(ptr.Pointer(), C.int(msec)))
	}
	return 0
}

func (ptr *QTimeLine) Resume() {
	defer qt.Recovering("QTimeLine::resume")

	if ptr.Pointer() != nil {
		C.QTimeLine_Resume(ptr.Pointer())
	}
}

func (ptr *QTimeLine) SetEndFrame(frame int) {
	defer qt.Recovering("QTimeLine::setEndFrame")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetEndFrame(ptr.Pointer(), C.int(frame))
	}
}

func (ptr *QTimeLine) SetFrameRange(startFrame int, endFrame int) {
	defer qt.Recovering("QTimeLine::setFrameRange")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetFrameRange(ptr.Pointer(), C.int(startFrame), C.int(endFrame))
	}
}

func (ptr *QTimeLine) SetPaused(paused bool) {
	defer qt.Recovering("QTimeLine::setPaused")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetPaused(ptr.Pointer(), C.int(qt.GoBoolToInt(paused)))
	}
}

func (ptr *QTimeLine) SetStartFrame(frame int) {
	defer qt.Recovering("QTimeLine::setStartFrame")

	if ptr.Pointer() != nil {
		C.QTimeLine_SetStartFrame(ptr.Pointer(), C.int(frame))
	}
}

func (ptr *QTimeLine) Start() {
	defer qt.Recovering("QTimeLine::start")

	if ptr.Pointer() != nil {
		C.QTimeLine_Start(ptr.Pointer())
	}
}

func (ptr *QTimeLine) StartFrame() int {
	defer qt.Recovering("QTimeLine::startFrame")

	if ptr.Pointer() != nil {
		return int(C.QTimeLine_StartFrame(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) ConnectStateChanged(f func(newState QTimeLine__State)) {
	defer qt.Recovering("connect QTimeLine::stateChanged")

	if ptr.Pointer() != nil {
		C.QTimeLine_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QTimeLine) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QTimeLine::stateChanged")

	if ptr.Pointer() != nil {
		C.QTimeLine_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQTimeLineStateChanged
func callbackQTimeLineStateChanged(ptrName *C.char, newState C.int) {
	defer qt.Recovering("callback QTimeLine::stateChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stateChanged")
	if signal != nil {
		signal.(func(QTimeLine__State))(QTimeLine__State(newState))
	}

}

func (ptr *QTimeLine) Stop() {
	defer qt.Recovering("QTimeLine::stop")

	if ptr.Pointer() != nil {
		C.QTimeLine_Stop(ptr.Pointer())
	}
}

func (ptr *QTimeLine) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QTimeLine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTimeLine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTimeLine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTimeLineTimerEvent
func callbackQTimeLineTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTimeLine::timerEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "timerEvent")
	if signal != nil {
		defer signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTimeLine) ToggleDirection() {
	defer qt.Recovering("QTimeLine::toggleDirection")

	if ptr.Pointer() != nil {
		C.QTimeLine_ToggleDirection(ptr.Pointer())
	}
}

func (ptr *QTimeLine) ValueForTime(msec int) float64 {
	defer qt.Recovering("QTimeLine::valueForTime")

	if ptr.Pointer() != nil {
		return float64(C.QTimeLine_ValueForTime(ptr.Pointer(), C.int(msec)))
	}
	return 0
}

func (ptr *QTimeLine) DestroyQTimeLine() {
	defer qt.Recovering("QTimeLine::~QTimeLine")

	if ptr.Pointer() != nil {
		C.QTimeLine_DestroyQTimeLine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
