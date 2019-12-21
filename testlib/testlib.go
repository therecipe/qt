// +build !minimal

package testlib

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "testlib.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtTestLib_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtTestLib_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QAbstractItemModelTester struct {
	ptr unsafe.Pointer
}

type QAbstractItemModelTester_ITF interface {
	QAbstractItemModelTester_PTR() *QAbstractItemModelTester
}

func (ptr *QAbstractItemModelTester) QAbstractItemModelTester_PTR() *QAbstractItemModelTester {
	return ptr
}

func (ptr *QAbstractItemModelTester) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractItemModelTester) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstractItemModelTester(ptr QAbstractItemModelTester_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModelTester_PTR().Pointer()
	}
	return nil
}

func NewQAbstractItemModelTesterFromPointer(ptr unsafe.Pointer) (n *QAbstractItemModelTester) {
	n = new(QAbstractItemModelTester)
	n.SetPointer(ptr)
	return
}
func (ptr *QAbstractItemModelTester) DestroyQAbstractItemModelTester() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QAbstractItemModelTester__FailureReportingMode
//QAbstractItemModelTester::FailureReportingMode
type QAbstractItemModelTester__FailureReportingMode int64

const (
	QAbstractItemModelTester__QtTest  QAbstractItemModelTester__FailureReportingMode = QAbstractItemModelTester__FailureReportingMode(0)
	QAbstractItemModelTester__Warning QAbstractItemModelTester__FailureReportingMode = QAbstractItemModelTester__FailureReportingMode(1)
	QAbstractItemModelTester__Fatal   QAbstractItemModelTester__FailureReportingMode = QAbstractItemModelTester__FailureReportingMode(2)
)

func NewQAbstractItemModelTester(model core.QAbstractItemModel_ITF, parent core.QObject_ITF) *QAbstractItemModelTester {
	tmpValue := NewQAbstractItemModelTesterFromPointer(C.QAbstractItemModelTester_NewQAbstractItemModelTester(core.PointerFromQAbstractItemModel(model), core.PointerFromQObject(parent)))
	qt.SetFinalizer(tmpValue, (*QAbstractItemModelTester).DestroyQAbstractItemModelTester)
	return tmpValue
}

func NewQAbstractItemModelTester2(model core.QAbstractItemModel_ITF, mode QAbstractItemModelTester__FailureReportingMode, parent core.QObject_ITF) *QAbstractItemModelTester {
	tmpValue := NewQAbstractItemModelTesterFromPointer(C.QAbstractItemModelTester_NewQAbstractItemModelTester2(core.PointerFromQAbstractItemModel(model), C.longlong(mode), core.PointerFromQObject(parent)))
	qt.SetFinalizer(tmpValue, (*QAbstractItemModelTester).DestroyQAbstractItemModelTester)
	return tmpValue
}

func (ptr *QAbstractItemModelTester) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QAbstractItemModelTester_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QEventSizeOfChecker struct {
	ptr unsafe.Pointer
}

type QEventSizeOfChecker_ITF interface {
	QEventSizeOfChecker_PTR() *QEventSizeOfChecker
}

func (ptr *QEventSizeOfChecker) QEventSizeOfChecker_PTR() *QEventSizeOfChecker {
	return ptr
}

func (ptr *QEventSizeOfChecker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QEventSizeOfChecker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQEventSizeOfChecker(ptr QEventSizeOfChecker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEventSizeOfChecker_PTR().Pointer()
	}
	return nil
}

func NewQEventSizeOfCheckerFromPointer(ptr unsafe.Pointer) (n *QEventSizeOfChecker) {
	n = new(QEventSizeOfChecker)
	n.SetPointer(ptr)
	return
}
func (ptr *QEventSizeOfChecker) DestroyQEventSizeOfChecker() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSignalSpy struct {
	core.QObject
}

type QSignalSpy_ITF interface {
	core.QObject_ITF
	QSignalSpy_PTR() *QSignalSpy
}

func (ptr *QSignalSpy) QSignalSpy_PTR() *QSignalSpy {
	return ptr
}

func (ptr *QSignalSpy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSignalSpy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSignalSpy(ptr QSignalSpy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalSpy_PTR().Pointer()
	}
	return nil
}

func NewQSignalSpyFromPointer(ptr unsafe.Pointer) (n *QSignalSpy) {
	n = new(QSignalSpy)
	n.SetPointer(ptr)
	return
}
func NewQSignalSpy(object core.QObject_ITF, sign string) *QSignalSpy {
	var signC *C.char
	if sign != "" {
		signC = C.CString(sign)
		defer C.free(unsafe.Pointer(signC))
	}
	tmpValue := NewQSignalSpyFromPointer(C.QSignalSpy_NewQSignalSpy(core.PointerFromQObject(object), signC))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSignalSpy) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSignalSpy_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSignalSpy) Signal() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSignalSpy_Signal(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) Wait(timeout int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSignalSpy_Wait(ptr.Pointer(), C.int(int32(timeout)))) != 0
	}
	return false
}

func (ptr *QSignalSpy) __args_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSignalSpy___args_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSignalSpy) __args_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___args_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSignalSpy) __args_newList() unsafe.Pointer {
	return C.QSignalSpy___args_newList(ptr.Pointer())
}

func (ptr *QSignalSpy) __setArgs__atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSignalSpy___setArgs__atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSignalSpy) __setArgs__setList(i int) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___setArgs__setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSignalSpy) __setArgs__newList() unsafe.Pointer {
	return C.QSignalSpy___setArgs__newList(ptr.Pointer())
}

func (ptr *QSignalSpy) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSignalSpy___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSignalSpy) __children_newList() unsafe.Pointer {
	return C.QSignalSpy___children_newList(ptr.Pointer())
}

func (ptr *QSignalSpy) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSignalSpy___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSignalSpy) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSignalSpy___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSignalSpy) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSignalSpy___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSignalSpy) __findChildren_newList() unsafe.Pointer {
	return C.QSignalSpy___findChildren_newList(ptr.Pointer())
}

func (ptr *QSignalSpy) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSignalSpy___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSignalSpy) __findChildren_newList3() unsafe.Pointer {
	return C.QSignalSpy___findChildren_newList3(ptr.Pointer())
}

//export callbackQSignalSpy_ChildEvent
func callbackQSignalSpy_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSignalSpy_ConnectNotify
func callbackQSignalSpy_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSignalSpyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSignalSpy) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSignalSpy_CustomEvent
func callbackQSignalSpy_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSignalSpy_DeleteLater
func callbackQSignalSpy_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSignalSpyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSignalSpy) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSignalSpy_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQSignalSpy_Destroyed
func callbackQSignalSpy_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSignalSpy_DisconnectNotify
func callbackQSignalSpy_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSignalSpyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSignalSpy) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSignalSpy_Event
func callbackQSignalSpy_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSignalSpyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSignalSpy) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSignalSpy_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSignalSpy_EventFilter
func callbackQSignalSpy_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSignalSpyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSignalSpy) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSignalSpy_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSignalSpy_MetaObject
func callbackQSignalSpy_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQSignalSpyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSignalSpy) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSignalSpy_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSignalSpy_ObjectNameChanged
func callbackQSignalSpy_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtTestLib_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQSignalSpy_TimerEvent
func callbackQSignalSpy_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSpontaneKeyEvent struct {
	ptr unsafe.Pointer
}

type QSpontaneKeyEvent_ITF interface {
	QSpontaneKeyEvent_PTR() *QSpontaneKeyEvent
}

func (ptr *QSpontaneKeyEvent) QSpontaneKeyEvent_PTR() *QSpontaneKeyEvent {
	return ptr
}

func (ptr *QSpontaneKeyEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSpontaneKeyEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSpontaneKeyEvent(ptr QSpontaneKeyEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSpontaneKeyEvent_PTR().Pointer()
	}
	return nil
}

func NewQSpontaneKeyEventFromPointer(ptr unsafe.Pointer) (n *QSpontaneKeyEvent) {
	n = new(QSpontaneKeyEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QSpontaneKeyEvent) DestroyQSpontaneKeyEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTest struct {
	ptr unsafe.Pointer
}

type QTest_ITF interface {
	QTest_PTR() *QTest
}

func (ptr *QTest) QTest_PTR() *QTest {
	return ptr
}

func (ptr *QTest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTest(ptr QTest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTest_PTR().Pointer()
	}
	return nil
}

func NewQTestFromPointer(ptr unsafe.Pointer) (n *QTest) {
	n = new(QTest)
	n.SetPointer(ptr)
	return
}
func (ptr *QTest) DestroyQTest() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QTest__TestFailMode
//QTest::TestFailMode
type QTest__TestFailMode int64

const (
	QTest__Abort    QTest__TestFailMode = QTest__TestFailMode(1)
	QTest__Continue QTest__TestFailMode = QTest__TestFailMode(2)
)

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
	QTest__RefCPUCycles         QTest__QBenchmarkMetric = QTest__QBenchmarkMetric(30)
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

type QTestData struct {
	ptr unsafe.Pointer
}

type QTestData_ITF interface {
	QTestData_PTR() *QTestData
}

func (ptr *QTestData) QTestData_PTR() *QTestData {
	return ptr
}

func (ptr *QTestData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTestData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTestData(ptr QTestData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestData_PTR().Pointer()
	}
	return nil
}

func NewQTestDataFromPointer(ptr unsafe.Pointer) (n *QTestData) {
	n = new(QTestData)
	n.SetPointer(ptr)
	return
}
func (ptr *QTestData) DestroyQTestData() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTestDelayEvent struct {
	QTestEvent
}

type QTestDelayEvent_ITF interface {
	QTestEvent_ITF
	QTestDelayEvent_PTR() *QTestDelayEvent
}

func (ptr *QTestDelayEvent) QTestDelayEvent_PTR() *QTestDelayEvent {
	return ptr
}

func (ptr *QTestDelayEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QTestDelayEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTestEvent_PTR().SetPointer(p)
	}
}

func PointerFromQTestDelayEvent(ptr QTestDelayEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestDelayEvent_PTR().Pointer()
	}
	return nil
}

func NewQTestDelayEventFromPointer(ptr unsafe.Pointer) (n *QTestDelayEvent) {
	n = new(QTestDelayEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QTestDelayEvent) DestroyQTestDelayEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTestEvent struct {
	ptr unsafe.Pointer
}

type QTestEvent_ITF interface {
	QTestEvent_PTR() *QTestEvent
}

func (ptr *QTestEvent) QTestEvent_PTR() *QTestEvent {
	return ptr
}

func (ptr *QTestEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTestEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTestEvent(ptr QTestEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEvent_PTR().Pointer()
	}
	return nil
}

func NewQTestEventFromPointer(ptr unsafe.Pointer) (n *QTestEvent) {
	n = new(QTestEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QTestEvent) DestroyQTestEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTestEventList struct {
	ptr unsafe.Pointer
}

type QTestEventList_ITF interface {
	QTestEventList_PTR() *QTestEventList
}

func (ptr *QTestEventList) QTestEventList_PTR() *QTestEventList {
	return ptr
}

func (ptr *QTestEventList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTestEventList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTestEventList(ptr QTestEventList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEventList_PTR().Pointer()
	}
	return nil
}

func NewQTestEventListFromPointer(ptr unsafe.Pointer) (n *QTestEventList) {
	n = new(QTestEventList)
	n.SetPointer(ptr)
	return
}
func NewQTestEventList() *QTestEventList {
	tmpValue := NewQTestEventListFromPointer(C.QTestEventList_NewQTestEventList())
	qt.SetFinalizer(tmpValue, (*QTestEventList).DestroyQTestEventList)
	return tmpValue
}

func NewQTestEventList2(other QTestEventList_ITF) *QTestEventList {
	tmpValue := NewQTestEventListFromPointer(C.QTestEventList_NewQTestEventList2(PointerFromQTestEventList(other)))
	qt.SetFinalizer(tmpValue, (*QTestEventList).DestroyQTestEventList)
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
		var asciiC *C.char
		if ascii != "" {
			asciiC = C.CString(ascii)
			defer C.free(unsafe.Pointer(asciiC))
		}
		C.QTestEventList_AddKeyClick2(ptr.Pointer(), asciiC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyClicks(keys string, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		var keysC *C.char
		if keys != "" {
			keysC = C.CString(keys)
			defer C.free(unsafe.Pointer(keysC))
		}
		C.QTestEventList_AddKeyClicks(ptr.Pointer(), C.struct_QtTestLib_PackedString{data: keysC, len: C.longlong(len(keys))}, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyPress(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddKeyPress(ptr.Pointer(), C.longlong(qtKey), C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyPress2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		var asciiC *C.char
		if ascii != "" {
			asciiC = C.CString(ascii)
			defer C.free(unsafe.Pointer(asciiC))
		}
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
		var asciiC *C.char
		if ascii != "" {
			asciiC = C.CString(ascii)
			defer C.free(unsafe.Pointer(asciiC))
		}
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

		qt.SetFinalizer(ptr, nil)
		C.QTestEventList_DestroyQTestEventList(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTestEventLoop struct {
	ptr unsafe.Pointer
}

type QTestEventLoop_ITF interface {
	QTestEventLoop_PTR() *QTestEventLoop
}

func (ptr *QTestEventLoop) QTestEventLoop_PTR() *QTestEventLoop {
	return ptr
}

func (ptr *QTestEventLoop) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTestEventLoop) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTestEventLoop(ptr QTestEventLoop_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEventLoop_PTR().Pointer()
	}
	return nil
}

func NewQTestEventLoopFromPointer(ptr unsafe.Pointer) (n *QTestEventLoop) {
	n = new(QTestEventLoop)
	n.SetPointer(ptr)
	return
}
func (ptr *QTestEventLoop) DestroyQTestEventLoop() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTestKeyClicksEvent struct {
	QTestEvent
}

type QTestKeyClicksEvent_ITF interface {
	QTestEvent_ITF
	QTestKeyClicksEvent_PTR() *QTestKeyClicksEvent
}

func (ptr *QTestKeyClicksEvent) QTestKeyClicksEvent_PTR() *QTestKeyClicksEvent {
	return ptr
}

func (ptr *QTestKeyClicksEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QTestKeyClicksEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTestEvent_PTR().SetPointer(p)
	}
}

func PointerFromQTestKeyClicksEvent(ptr QTestKeyClicksEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestKeyClicksEvent_PTR().Pointer()
	}
	return nil
}

func NewQTestKeyClicksEventFromPointer(ptr unsafe.Pointer) (n *QTestKeyClicksEvent) {
	n = new(QTestKeyClicksEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QTestKeyClicksEvent) DestroyQTestKeyClicksEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTestKeyEvent struct {
	QTestEvent
}

type QTestKeyEvent_ITF interface {
	QTestEvent_ITF
	QTestKeyEvent_PTR() *QTestKeyEvent
}

func (ptr *QTestKeyEvent) QTestKeyEvent_PTR() *QTestKeyEvent {
	return ptr
}

func (ptr *QTestKeyEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QTestKeyEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTestEvent_PTR().SetPointer(p)
	}
}

func PointerFromQTestKeyEvent(ptr QTestKeyEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestKeyEvent_PTR().Pointer()
	}
	return nil
}

func NewQTestKeyEventFromPointer(ptr unsafe.Pointer) (n *QTestKeyEvent) {
	n = new(QTestKeyEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QTestKeyEvent) DestroyQTestKeyEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTestMouseEvent struct {
	QTestEvent
}

type QTestMouseEvent_ITF interface {
	QTestEvent_ITF
	QTestMouseEvent_PTR() *QTestMouseEvent
}

func (ptr *QTestMouseEvent) QTestMouseEvent_PTR() *QTestMouseEvent {
	return ptr
}

func (ptr *QTestMouseEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QTestMouseEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTestEvent_PTR().SetPointer(p)
	}
}

func PointerFromQTestMouseEvent(ptr QTestMouseEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestMouseEvent_PTR().Pointer()
	}
	return nil
}

func NewQTestMouseEventFromPointer(ptr unsafe.Pointer) (n *QTestMouseEvent) {
	n = new(QTestMouseEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QTestMouseEvent) DestroyQTestMouseEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func init() {
	qt.ItfMap["testlib.QAbstractItemModelTester_ITF"] = QAbstractItemModelTester{}
	qt.FuncMap["testlib.NewQAbstractItemModelTester"] = NewQAbstractItemModelTester
	qt.FuncMap["testlib.NewQAbstractItemModelTester2"] = NewQAbstractItemModelTester2
	qt.EnumMap["testlib.QAbstractItemModelTester__QtTest"] = int64(QAbstractItemModelTester__QtTest)
	qt.EnumMap["testlib.QAbstractItemModelTester__Warning"] = int64(QAbstractItemModelTester__Warning)
	qt.EnumMap["testlib.QAbstractItemModelTester__Fatal"] = int64(QAbstractItemModelTester__Fatal)
	qt.ItfMap["testlib.QEventSizeOfChecker_ITF"] = QEventSizeOfChecker{}
	qt.ItfMap["testlib.QSignalSpy_ITF"] = QSignalSpy{}
	qt.FuncMap["testlib.NewQSignalSpy"] = NewQSignalSpy
	qt.ItfMap["testlib.QSpontaneKeyEvent_ITF"] = QSpontaneKeyEvent{}
	qt.EnumMap["testlib.QTest__Abort"] = int64(QTest__Abort)
	qt.EnumMap["testlib.QTest__Continue"] = int64(QTest__Continue)
	qt.EnumMap["testlib.QTest__FramesPerSecond"] = int64(QTest__FramesPerSecond)
	qt.EnumMap["testlib.QTest__BitsPerSecond"] = int64(QTest__BitsPerSecond)
	qt.EnumMap["testlib.QTest__BytesPerSecond"] = int64(QTest__BytesPerSecond)
	qt.EnumMap["testlib.QTest__WalltimeMilliseconds"] = int64(QTest__WalltimeMilliseconds)
	qt.EnumMap["testlib.QTest__CPUTicks"] = int64(QTest__CPUTicks)
	qt.EnumMap["testlib.QTest__InstructionReads"] = int64(QTest__InstructionReads)
	qt.EnumMap["testlib.QTest__Events"] = int64(QTest__Events)
	qt.EnumMap["testlib.QTest__WalltimeNanoseconds"] = int64(QTest__WalltimeNanoseconds)
	qt.EnumMap["testlib.QTest__BytesAllocated"] = int64(QTest__BytesAllocated)
	qt.EnumMap["testlib.QTest__CPUMigrations"] = int64(QTest__CPUMigrations)
	qt.EnumMap["testlib.QTest__CPUCycles"] = int64(QTest__CPUCycles)
	qt.EnumMap["testlib.QTest__BusCycles"] = int64(QTest__BusCycles)
	qt.EnumMap["testlib.QTest__StalledCycles"] = int64(QTest__StalledCycles)
	qt.EnumMap["testlib.QTest__Instructions"] = int64(QTest__Instructions)
	qt.EnumMap["testlib.QTest__BranchInstructions"] = int64(QTest__BranchInstructions)
	qt.EnumMap["testlib.QTest__BranchMisses"] = int64(QTest__BranchMisses)
	qt.EnumMap["testlib.QTest__CacheReferences"] = int64(QTest__CacheReferences)
	qt.EnumMap["testlib.QTest__CacheReads"] = int64(QTest__CacheReads)
	qt.EnumMap["testlib.QTest__CacheWrites"] = int64(QTest__CacheWrites)
	qt.EnumMap["testlib.QTest__CachePrefetches"] = int64(QTest__CachePrefetches)
	qt.EnumMap["testlib.QTest__CacheMisses"] = int64(QTest__CacheMisses)
	qt.EnumMap["testlib.QTest__CacheReadMisses"] = int64(QTest__CacheReadMisses)
	qt.EnumMap["testlib.QTest__CacheWriteMisses"] = int64(QTest__CacheWriteMisses)
	qt.EnumMap["testlib.QTest__CachePrefetchMisses"] = int64(QTest__CachePrefetchMisses)
	qt.EnumMap["testlib.QTest__ContextSwitches"] = int64(QTest__ContextSwitches)
	qt.EnumMap["testlib.QTest__PageFaults"] = int64(QTest__PageFaults)
	qt.EnumMap["testlib.QTest__MinorPageFaults"] = int64(QTest__MinorPageFaults)
	qt.EnumMap["testlib.QTest__MajorPageFaults"] = int64(QTest__MajorPageFaults)
	qt.EnumMap["testlib.QTest__AlignmentFaults"] = int64(QTest__AlignmentFaults)
	qt.EnumMap["testlib.QTest__EmulationFaults"] = int64(QTest__EmulationFaults)
	qt.EnumMap["testlib.QTest__RefCPUCycles"] = int64(QTest__RefCPUCycles)
	qt.EnumMap["testlib.QTest__Press"] = int64(QTest__Press)
	qt.EnumMap["testlib.QTest__Release"] = int64(QTest__Release)
	qt.EnumMap["testlib.QTest__Click"] = int64(QTest__Click)
	qt.EnumMap["testlib.QTest__Shortcut"] = int64(QTest__Shortcut)
	qt.EnumMap["testlib.QTest__MousePress"] = int64(QTest__MousePress)
	qt.EnumMap["testlib.QTest__MouseRelease"] = int64(QTest__MouseRelease)
	qt.EnumMap["testlib.QTest__MouseClick"] = int64(QTest__MouseClick)
	qt.EnumMap["testlib.QTest__MouseDClick"] = int64(QTest__MouseDClick)
	qt.EnumMap["testlib.QTest__MouseMove"] = int64(QTest__MouseMove)
	qt.EnumMap["testlib.QTest__AI_Undefined"] = int64(QTest__AI_Undefined)
	qt.EnumMap["testlib.QTest__AI_Name"] = int64(QTest__AI_Name)
	qt.EnumMap["testlib.QTest__AI_Result"] = int64(QTest__AI_Result)
	qt.EnumMap["testlib.QTest__AI_Tests"] = int64(QTest__AI_Tests)
	qt.EnumMap["testlib.QTest__AI_Failures"] = int64(QTest__AI_Failures)
	qt.EnumMap["testlib.QTest__AI_Errors"] = int64(QTest__AI_Errors)
	qt.EnumMap["testlib.QTest__AI_Type"] = int64(QTest__AI_Type)
	qt.EnumMap["testlib.QTest__AI_Description"] = int64(QTest__AI_Description)
	qt.EnumMap["testlib.QTest__AI_PropertyValue"] = int64(QTest__AI_PropertyValue)
	qt.EnumMap["testlib.QTest__AI_QTestVersion"] = int64(QTest__AI_QTestVersion)
	qt.EnumMap["testlib.QTest__AI_QtVersion"] = int64(QTest__AI_QtVersion)
	qt.EnumMap["testlib.QTest__AI_File"] = int64(QTest__AI_File)
	qt.EnumMap["testlib.QTest__AI_Line"] = int64(QTest__AI_Line)
	qt.EnumMap["testlib.QTest__AI_Metric"] = int64(QTest__AI_Metric)
	qt.EnumMap["testlib.QTest__AI_Tag"] = int64(QTest__AI_Tag)
	qt.EnumMap["testlib.QTest__AI_Value"] = int64(QTest__AI_Value)
	qt.EnumMap["testlib.QTest__AI_Iterations"] = int64(QTest__AI_Iterations)
	qt.EnumMap["testlib.QTest__LET_Undefined"] = int64(QTest__LET_Undefined)
	qt.EnumMap["testlib.QTest__LET_Property"] = int64(QTest__LET_Property)
	qt.EnumMap["testlib.QTest__LET_Properties"] = int64(QTest__LET_Properties)
	qt.EnumMap["testlib.QTest__LET_Failure"] = int64(QTest__LET_Failure)
	qt.EnumMap["testlib.QTest__LET_Error"] = int64(QTest__LET_Error)
	qt.EnumMap["testlib.QTest__LET_TestCase"] = int64(QTest__LET_TestCase)
	qt.EnumMap["testlib.QTest__LET_TestSuite"] = int64(QTest__LET_TestSuite)
	qt.EnumMap["testlib.QTest__LET_Benchmark"] = int64(QTest__LET_Benchmark)
	qt.EnumMap["testlib.QTest__LET_SystemError"] = int64(QTest__LET_SystemError)
	qt.ItfMap["testlib.QTestData_ITF"] = QTestData{}
	qt.ItfMap["testlib.QTestDelayEvent_ITF"] = QTestDelayEvent{}
	qt.ItfMap["testlib.QTestEvent_ITF"] = QTestEvent{}
	qt.ItfMap["testlib.QTestEventList_ITF"] = QTestEventList{}
	qt.FuncMap["testlib.NewQTestEventList"] = NewQTestEventList
	qt.FuncMap["testlib.NewQTestEventList2"] = NewQTestEventList2
	qt.ItfMap["testlib.QTestEventLoop_ITF"] = QTestEventLoop{}
	qt.ItfMap["testlib.QTestKeyClicksEvent_ITF"] = QTestKeyClicksEvent{}
	qt.ItfMap["testlib.QTestKeyEvent_ITF"] = QTestKeyEvent{}
	qt.ItfMap["testlib.QTestMouseEvent_ITF"] = QTestMouseEvent{}
}
