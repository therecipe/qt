package core

//#include "qtimeline.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QTimeLine_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return int(C.QTimeLine_CurrentTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) CurveShape() QTimeLine__CurveShape {
	if ptr.Pointer() != nil {
		return QTimeLine__CurveShape(C.QTimeLine_CurveShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) Direction() QTimeLine__Direction {
	if ptr.Pointer() != nil {
		return QTimeLine__Direction(C.QTimeLine_Direction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) Duration() int {
	if ptr.Pointer() != nil {
		return int(C.QTimeLine_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) EasingCurve() *QEasingCurve {
	if ptr.Pointer() != nil {
		return NewQEasingCurveFromPointer(C.QTimeLine_EasingCurve(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTimeLine) LoopCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTimeLine_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) SetCurrentTime(msec int) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetCurrentTime(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QTimeLine) SetCurveShape(shape QTimeLine__CurveShape) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetCurveShape(ptr.Pointer(), C.int(shape))
	}
}

func (ptr *QTimeLine) SetDirection(direction QTimeLine__Direction) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QTimeLine) SetDuration(duration int) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetDuration(ptr.Pointer(), C.int(duration))
	}
}

func (ptr *QTimeLine) SetEasingCurve(curve QEasingCurve_ITF) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetEasingCurve(ptr.Pointer(), PointerFromQEasingCurve(curve))
	}
}

func (ptr *QTimeLine) SetLoopCount(count int) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetLoopCount(ptr.Pointer(), C.int(count))
	}
}

func (ptr *QTimeLine) SetUpdateInterval(interval int) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetUpdateInterval(ptr.Pointer(), C.int(interval))
	}
}

func (ptr *QTimeLine) UpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QTimeLine_UpdateInterval(ptr.Pointer()))
	}
	return 0
}

func NewQTimeLine(duration int, parent QObject_ITF) *QTimeLine {
	return NewQTimeLineFromPointer(C.QTimeLine_NewQTimeLine(C.int(duration), PointerFromQObject(parent)))
}

func (ptr *QTimeLine) CurrentFrame() int {
	if ptr.Pointer() != nil {
		return int(C.QTimeLine_CurrentFrame(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) CurrentValue() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTimeLine_CurrentValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) EndFrame() int {
	if ptr.Pointer() != nil {
		return int(C.QTimeLine_EndFrame(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QTimeLine_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QTimeLine) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QTimeLine_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQTimeLineFinished
func callbackQTimeLineFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QTimeLine) ConnectFrameChanged(f func(frame int)) {
	if ptr.Pointer() != nil {
		C.QTimeLine_ConnectFrameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frameChanged", f)
	}
}

func (ptr *QTimeLine) DisconnectFrameChanged() {
	if ptr.Pointer() != nil {
		C.QTimeLine_DisconnectFrameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frameChanged")
	}
}

//export callbackQTimeLineFrameChanged
func callbackQTimeLineFrameChanged(ptrName *C.char, frame C.int) {
	qt.GetSignal(C.GoString(ptrName), "frameChanged").(func(int))(int(frame))
}

func (ptr *QTimeLine) FrameForTime(msec int) int {
	if ptr.Pointer() != nil {
		return int(C.QTimeLine_FrameForTime(ptr.Pointer(), C.int(msec)))
	}
	return 0
}

func (ptr *QTimeLine) Resume() {
	if ptr.Pointer() != nil {
		C.QTimeLine_Resume(ptr.Pointer())
	}
}

func (ptr *QTimeLine) SetEndFrame(frame int) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetEndFrame(ptr.Pointer(), C.int(frame))
	}
}

func (ptr *QTimeLine) SetFrameRange(startFrame int, endFrame int) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetFrameRange(ptr.Pointer(), C.int(startFrame), C.int(endFrame))
	}
}

func (ptr *QTimeLine) SetPaused(paused bool) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetPaused(ptr.Pointer(), C.int(qt.GoBoolToInt(paused)))
	}
}

func (ptr *QTimeLine) SetStartFrame(frame int) {
	if ptr.Pointer() != nil {
		C.QTimeLine_SetStartFrame(ptr.Pointer(), C.int(frame))
	}
}

func (ptr *QTimeLine) Start() {
	if ptr.Pointer() != nil {
		C.QTimeLine_Start(ptr.Pointer())
	}
}

func (ptr *QTimeLine) StartFrame() int {
	if ptr.Pointer() != nil {
		return int(C.QTimeLine_StartFrame(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeLine) ConnectStateChanged(f func(newState QTimeLine__State)) {
	if ptr.Pointer() != nil {
		C.QTimeLine_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QTimeLine) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QTimeLine_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQTimeLineStateChanged
func callbackQTimeLineStateChanged(ptrName *C.char, newState C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QTimeLine__State))(QTimeLine__State(newState))
}

func (ptr *QTimeLine) Stop() {
	if ptr.Pointer() != nil {
		C.QTimeLine_Stop(ptr.Pointer())
	}
}

func (ptr *QTimeLine) ToggleDirection() {
	if ptr.Pointer() != nil {
		C.QTimeLine_ToggleDirection(ptr.Pointer())
	}
}

func (ptr *QTimeLine) ValueForTime(msec int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTimeLine_ValueForTime(ptr.Pointer(), C.int(msec)))
	}
	return 0
}

func (ptr *QTimeLine) DestroyQTimeLine() {
	if ptr.Pointer() != nil {
		C.QTimeLine_DestroyQTimeLine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
