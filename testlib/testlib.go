// +build !minimal

package testlib

//#include <stdint.h>
//#include <stdlib.h>
//#include "testlib.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtTestLib_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QSignalSpy struct {
	core.QObject
	core.QList
}

type QSignalSpy_ITF interface {
	core.QObject_ITF
	core.QList_ITF
	QSignalSpy_PTR() *QSignalSpy
}

func (p *QSignalSpy) QSignalSpy_PTR() *QSignalSpy {
	return p
}

func (p *QSignalSpy) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSignalSpy) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
		p.QList_PTR().SetPointer(ptr)
	}
}

func PointerFromQSignalSpy(ptr QSignalSpy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalSpy_PTR().Pointer()
	}
	return nil
}

func NewQSignalSpyFromPointer(ptr unsafe.Pointer) *QSignalSpy {
	var n = new(QSignalSpy)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSignalSpy) DestroyQSignalSpy() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func NewQSignalSpy(object core.QObject_ITF, sign string) *QSignalSpy {
	var signC = C.CString(sign)
	defer C.free(unsafe.Pointer(signC))
	var tmpValue = NewQSignalSpyFromPointer(C.QSignalSpy_NewQSignalSpy(core.PointerFromQObject(object), signC))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSignalSpy) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSignalSpy_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSignalSpy) Signal() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSignalSpy_Signal(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) Wait(timeout int) bool {
	if ptr.Pointer() != nil {
		return C.QSignalSpy_Wait(ptr.Pointer(), C.int(int32(timeout))) != 0
	}
	return false
}

//export callbackQSignalSpy_TimerEvent
func callbackQSignalSpy_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSignalSpy::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::timerEvent", f)
	}
}

func (ptr *QSignalSpy) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::timerEvent")
	}
}

func (ptr *QSignalSpy) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSignalSpy) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSignalSpy_ChildEvent
func callbackQSignalSpy_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSignalSpy::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::childEvent", f)
	}
}

func (ptr *QSignalSpy) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::childEvent")
	}
}

func (ptr *QSignalSpy) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSignalSpy) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSignalSpy_ConnectNotify
func callbackQSignalSpy_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSignalSpy::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSignalSpyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSignalSpy) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::connectNotify", f)
	}
}

func (ptr *QSignalSpy) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::connectNotify")
	}
}

func (ptr *QSignalSpy) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSignalSpy) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSignalSpy_CustomEvent
func callbackQSignalSpy_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSignalSpy::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::customEvent", f)
	}
}

func (ptr *QSignalSpy) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::customEvent")
	}
}

func (ptr *QSignalSpy) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSignalSpy) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSignalSpy_DeleteLater
func callbackQSignalSpy_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSignalSpy::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSignalSpyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSignalSpy) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::deleteLater", f)
	}
}

func (ptr *QSignalSpy) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::deleteLater")
	}
}

func (ptr *QSignalSpy) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSignalSpy_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSignalSpy) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSignalSpy_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSignalSpy_DisconnectNotify
func callbackQSignalSpy_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSignalSpy::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSignalSpyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSignalSpy) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::disconnectNotify", f)
	}
}

func (ptr *QSignalSpy) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::disconnectNotify")
	}
}

func (ptr *QSignalSpy) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSignalSpy) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSignalSpy_Event
func callbackQSignalSpy_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSignalSpy::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSignalSpyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSignalSpy) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::event", f)
	}
}

func (ptr *QSignalSpy) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::event")
	}
}

func (ptr *QSignalSpy) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSignalSpy_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSignalSpy) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSignalSpy_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSignalSpy_EventFilter
func callbackQSignalSpy_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSignalSpy::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSignalSpyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSignalSpy) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::eventFilter", f)
	}
}

func (ptr *QSignalSpy) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::eventFilter")
	}
}

func (ptr *QSignalSpy) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSignalSpy_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSignalSpy) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSignalSpy_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSignalSpy_MetaObject
func callbackQSignalSpy_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSignalSpy::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSignalSpyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSignalSpy) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::metaObject", f)
	}
}

func (ptr *QSignalSpy) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSignalSpy::metaObject")
	}
}

func (ptr *QSignalSpy) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSignalSpy_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSignalSpy) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSignalSpy_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//go:generate stringer -type=QTest__QBenchmarkMetric
//QTest::QBenchmarkMetric
type QTest__QBenchmarkMetric int64

const (
	QTest__FramesPerSecond      QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(0)
	QTest__BitsPerSecond        QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(1)
	QTest__BytesPerSecond       QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(2)
	QTest__WalltimeMilliseconds QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(3)
	QTest__CPUTicks             QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(4)
	QTest__InstructionReads     QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(5)
	QTest__Events               QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(6)
	QTest__WalltimeNanoseconds  QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(7)
	QTest__BytesAllocated       QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(8)
	QTest__CPUMigrations        QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(9)
	QTest__CPUCycles            QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(10)
	QTest__BusCycles            QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(11)
	QTest__StalledCycles        QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(12)
	QTest__Instructions         QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(13)
	QTest__BranchInstructions   QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(14)
	QTest__BranchMisses         QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(15)
	QTest__CacheReferences      QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(16)
	QTest__CacheReads           QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(17)
	QTest__CacheWrites          QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(18)
	QTest__CachePrefetches      QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(19)
	QTest__CacheMisses          QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(20)
	QTest__CacheReadMisses      QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(21)
	QTest__CacheWriteMisses     QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(22)
	QTest__CachePrefetchMisses  QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(23)
	QTest__ContextSwitches      QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(24)
	QTest__PageFaults           QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(25)
	QTest__MinorPageFaults      QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(26)
	QTest__MajorPageFaults      QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(27)
	QTest__AlignmentFaults      QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(28)
	QTest__EmulationFaults      QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(29)
)

//go:generate stringer -type=QTest__TestFailMode
//QTest::TestFailMode
type QTest__TestFailMode int64

const (
	QTest__Abort    QTest__TestFailMode = QTest__TestFailMode(1)
	QTest__Continue QTest__TestFailMode = QTest__TestFailMode(2)
)

//go:generate stringer -type=QTest__AttributeIndex
//QTest::AttributeIndex
type QTest__AttributeIndex int64

const (
	QTest__AI_Undefined     QTest__AttributeIndex = QTest__AttributeIndex(-1)
	QTest__AI_Name          QTest__AttributeIndex = QTest__AttributeIndex(0)
	QTest__AI_Result        QTest__AttributeIndex = QTest__AttributeIndex(1)
	QTest__AI_Tests         QTest__AttributeIndex = QTest__AttributeIndex(2)
	QTest__AI_Failures      QTest__AttributeIndex = QTest__AttributeIndex(3)
	QTest__AI_Errors        QTest__AttributeIndex = QTest__AttributeIndex(4)
	QTest__AI_Type          QTest__AttributeIndex = QTest__AttributeIndex(5)
	QTest__AI_Description   QTest__AttributeIndex = QTest__AttributeIndex(6)
	QTest__AI_PropertyValue QTest__AttributeIndex = QTest__AttributeIndex(7)
	QTest__AI_QTestVersion  QTest__AttributeIndex = QTest__AttributeIndex(8)
	QTest__AI_QtVersion     QTest__AttributeIndex = QTest__AttributeIndex(9)
	QTest__AI_File          QTest__AttributeIndex = QTest__AttributeIndex(10)
	QTest__AI_Line          QTest__AttributeIndex = QTest__AttributeIndex(11)
	QTest__AI_Metric        QTest__AttributeIndex = QTest__AttributeIndex(12)
	QTest__AI_Tag           QTest__AttributeIndex = QTest__AttributeIndex(13)
	QTest__AI_Value         QTest__AttributeIndex = QTest__AttributeIndex(14)
	QTest__AI_Iterations    QTest__AttributeIndex = QTest__AttributeIndex(15)
)

//go:generate stringer -type=QTest__LogElementType
//QTest::LogElementType
type QTest__LogElementType int64

const (
	QTest__LET_Undefined   QTest__LogElementType = QTest__LogElementType(-1)
	QTest__LET_Property    QTest__LogElementType = QTest__LogElementType(0)
	QTest__LET_Properties  QTest__LogElementType = QTest__LogElementType(1)
	QTest__LET_Failure     QTest__LogElementType = QTest__LogElementType(2)
	QTest__LET_Error       QTest__LogElementType = QTest__LogElementType(3)
	QTest__LET_TestCase    QTest__LogElementType = QTest__LogElementType(4)
	QTest__LET_TestSuite   QTest__LogElementType = QTest__LogElementType(5)
	QTest__LET_Benchmark   QTest__LogElementType = QTest__LogElementType(6)
	QTest__LET_SystemError QTest__LogElementType = QTest__LogElementType(7)
)

//go:generate stringer -type=QTest__KeyAction
//QTest::KeyAction
type QTest__KeyAction int64

const (
	QTest__Press    QTest__KeyAction = QTest__KeyAction(0)
	QTest__Release  QTest__KeyAction = QTest__KeyAction(1)
	QTest__Click    QTest__KeyAction = QTest__KeyAction(2)
	QTest__Shortcut QTest__KeyAction = QTest__KeyAction(3)
)

//go:generate stringer -type=QTest__MouseAction
//QTest::MouseAction
type QTest__MouseAction int64

const (
	QTest__MousePress   QTest__MouseAction = QTest__MouseAction(0)
	QTest__MouseRelease QTest__MouseAction = QTest__MouseAction(1)
	QTest__MouseClick   QTest__MouseAction = QTest__MouseAction(2)
	QTest__MouseDClick  QTest__MouseAction = QTest__MouseAction(3)
	QTest__MouseMove    QTest__MouseAction = QTest__MouseAction(4)
)

type QTest struct {
	ptr unsafe.Pointer
}

type QTest_ITF interface {
	QTest_PTR() *QTest
}

func (p *QTest) QTest_PTR() *QTest {
	return p
}

func (p *QTest) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QTest) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQTest(ptr QTest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTest_PTR().Pointer()
	}
	return nil
}

func NewQTestFromPointer(ptr unsafe.Pointer) *QTest {
	var n = new(QTest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTest) DestroyQTest() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

type QTestEventList struct {
	core.QList
}

type QTestEventList_ITF interface {
	core.QList_ITF
	QTestEventList_PTR() *QTestEventList
}

func (p *QTestEventList) QTestEventList_PTR() *QTestEventList {
	return p
}

func (p *QTestEventList) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QList_PTR().Pointer()
	}
	return nil
}

func (p *QTestEventList) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QList_PTR().SetPointer(ptr)
	}
}

func PointerFromQTestEventList(ptr QTestEventList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEventList_PTR().Pointer()
	}
	return nil
}

func NewQTestEventListFromPointer(ptr unsafe.Pointer) *QTestEventList {
	var n = new(QTestEventList)
	n.SetPointer(ptr)
	return n
}
func NewQTestEventList() *QTestEventList {
	var tmpValue = NewQTestEventListFromPointer(C.QTestEventList_NewQTestEventList())
	runtime.SetFinalizer(tmpValue, (*QTestEventList).DestroyQTestEventList)
	return tmpValue
}

func NewQTestEventList2(other QTestEventList_ITF) *QTestEventList {
	var tmpValue = NewQTestEventListFromPointer(C.QTestEventList_NewQTestEventList2(PointerFromQTestEventList(other)))
	runtime.SetFinalizer(tmpValue, (*QTestEventList).DestroyQTestEventList)
	return tmpValue
}

func (ptr *QTestEventList) AddDelay(msecs int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddDelay(ptr.Pointer(), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyClick(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddKeyClick(ptr.Pointer(), C.longlong(qtKey), C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyClick2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		var asciiC = C.CString(ascii)
		defer C.free(unsafe.Pointer(asciiC))
		C.QTestEventList_AddKeyClick2(ptr.Pointer(), asciiC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyClicks(keys string, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		var keysC = C.CString(keys)
		defer C.free(unsafe.Pointer(keysC))
		C.QTestEventList_AddKeyClicks(ptr.Pointer(), keysC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyPress(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddKeyPress(ptr.Pointer(), C.longlong(qtKey), C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyPress2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		var asciiC = C.CString(ascii)
		defer C.free(unsafe.Pointer(asciiC))
		C.QTestEventList_AddKeyPress2(ptr.Pointer(), asciiC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyRelease(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddKeyRelease(ptr.Pointer(), C.longlong(qtKey), C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyRelease2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		var asciiC = C.CString(ascii)
		defer C.free(unsafe.Pointer(asciiC))
		C.QTestEventList_AddKeyRelease2(ptr.Pointer(), asciiC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddMouseClick(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseClick(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMouseDClick(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseDClick(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMouseMove(pos core.QPoint_ITF, delay int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseMove(ptr.Pointer(), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMousePress(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddMousePress(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMouseRelease(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseRelease(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) Clear() {
	if ptr.Pointer() != nil {
		C.QTestEventList_Clear(ptr.Pointer())
	}
}

func (ptr *QTestEventList) Simulate(w widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QTestEventList_Simulate(ptr.Pointer(), widgets.PointerFromQWidget(w))
	}
}

func (ptr *QTestEventList) DestroyQTestEventList() {
	if ptr.Pointer() != nil {
		C.QTestEventList_DestroyQTestEventList(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
