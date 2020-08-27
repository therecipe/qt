// +build !minimal

package testlib

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QAbstractItemModelTester struct {
	internal.Internal
}

type QAbstractItemModelTester_ITF interface {
	QAbstractItemModelTester_PTR() *QAbstractItemModelTester
}

func (ptr *QAbstractItemModelTester) QAbstractItemModelTester_PTR() *QAbstractItemModelTester {
	return ptr
}

func (ptr *QAbstractItemModelTester) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAbstractItemModelTester) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAbstractItemModelTester(ptr QAbstractItemModelTester_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModelTester_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractItemModelTester) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAbstractItemModelTesterFromPointer(ptr unsafe.Pointer) (n *QAbstractItemModelTester) {
	n = new(QAbstractItemModelTester)
	n.InitFromInternal(uintptr(ptr), "testlib.QAbstractItemModelTester")
	return
}

func (ptr *QAbstractItemModelTester) DestroyQAbstractItemModelTester() {
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

	return internal.CallLocalFunction([]interface{}{"", "", "testlib.NewQAbstractItemModelTester", "", model, parent}).(*QAbstractItemModelTester)
}

func NewQAbstractItemModelTester2(model core.QAbstractItemModel_ITF, mode QAbstractItemModelTester__FailureReportingMode, parent core.QObject_ITF) *QAbstractItemModelTester {

	return internal.CallLocalFunction([]interface{}{"", "", "testlib.NewQAbstractItemModelTester2", "", model, mode, parent}).(*QAbstractItemModelTester)
}

func (ptr *QAbstractItemModelTester) Model() *core.QAbstractItemModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Model"}).(*core.QAbstractItemModel)
}

type QEventSizeOfChecker struct {
	internal.Internal
}

type QEventSizeOfChecker_ITF interface {
	QEventSizeOfChecker_PTR() *QEventSizeOfChecker
}

func (ptr *QEventSizeOfChecker) QEventSizeOfChecker_PTR() *QEventSizeOfChecker {
	return ptr
}

func (ptr *QEventSizeOfChecker) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QEventSizeOfChecker) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQEventSizeOfChecker(ptr QEventSizeOfChecker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEventSizeOfChecker_PTR().Pointer()
	}
	return nil
}

func (n *QEventSizeOfChecker) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQEventSizeOfCheckerFromPointer(ptr unsafe.Pointer) (n *QEventSizeOfChecker) {
	n = new(QEventSizeOfChecker)
	n.InitFromInternal(uintptr(ptr), "testlib.QEventSizeOfChecker")
	return
}

func (ptr *QEventSizeOfChecker) DestroyQEventSizeOfChecker() {
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

func (n *QSignalSpy) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSignalSpy) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSignalSpyFromPointer(ptr unsafe.Pointer) (n *QSignalSpy) {
	n = new(QSignalSpy)
	n.InitFromInternal(uintptr(ptr), "testlib.QSignalSpy")
	return
}

func (ptr *QSignalSpy) DestroyQSignalSpy() {
}

func NewQSignalSpy(object core.QObject_ITF, sign string) *QSignalSpy {

	return internal.CallLocalFunction([]interface{}{"", "", "testlib.NewQSignalSpy", "", object, sign}).(*QSignalSpy)
}

func (ptr *QSignalSpy) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QSignalSpy) Signal() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Signal"}).(*core.QByteArray)
}

func (ptr *QSignalSpy) Wait(timeout int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Wait", timeout}).(bool)
}

func (ptr *QSignalSpy) __args_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__args_atList", i}).(float64))
}

func (ptr *QSignalSpy) __args_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__args_setList", i})
}

func (ptr *QSignalSpy) __args_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__args_newList"}).(unsafe.Pointer)
}

func (ptr *QSignalSpy) __setArgs__atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setArgs__atList", i}).(float64))
}

func (ptr *QSignalSpy) __setArgs__setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setArgs__setList", i})
}

func (ptr *QSignalSpy) __setArgs__newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setArgs__newList"}).(unsafe.Pointer)
}

func (ptr *QSignalSpy) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSignalSpy) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSignalSpy) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSignalSpy) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSignalSpy) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSignalSpy) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSignalSpy) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSignalSpy) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSignalSpy) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSignalSpy) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSignalSpy) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSignalSpy) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSignalSpy) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSignalSpy) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSignalSpy) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSignalSpy) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSignalSpy) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSignalSpy) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSignalSpy) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSignalSpy) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSignalSpy) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSpontaneKeyEvent struct {
	internal.Internal
}

type QSpontaneKeyEvent_ITF interface {
	QSpontaneKeyEvent_PTR() *QSpontaneKeyEvent
}

func (ptr *QSpontaneKeyEvent) QSpontaneKeyEvent_PTR() *QSpontaneKeyEvent {
	return ptr
}

func (ptr *QSpontaneKeyEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSpontaneKeyEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSpontaneKeyEvent(ptr QSpontaneKeyEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSpontaneKeyEvent_PTR().Pointer()
	}
	return nil
}

func (n *QSpontaneKeyEvent) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSpontaneKeyEventFromPointer(ptr unsafe.Pointer) (n *QSpontaneKeyEvent) {
	n = new(QSpontaneKeyEvent)
	n.InitFromInternal(uintptr(ptr), "testlib.QSpontaneKeyEvent")
	return
}

func (ptr *QSpontaneKeyEvent) DestroyQSpontaneKeyEvent() {
}

type QTest struct {
	internal.Internal
}

type QTest_ITF interface {
	QTest_PTR() *QTest
}

func (ptr *QTest) QTest_PTR() *QTest {
	return ptr
}

func (ptr *QTest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QTest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQTest(ptr QTest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTest_PTR().Pointer()
	}
	return nil
}

func (n *QTest) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQTestFromPointer(ptr unsafe.Pointer) (n *QTest) {
	n = new(QTest)
	n.InitFromInternal(uintptr(ptr), "testlib.QTest")
	return
}

func (ptr *QTest) DestroyQTest() {
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
	internal.Internal
}

type QTestData_ITF interface {
	QTestData_PTR() *QTestData
}

func (ptr *QTestData) QTestData_PTR() *QTestData {
	return ptr
}

func (ptr *QTestData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QTestData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQTestData(ptr QTestData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestData_PTR().Pointer()
	}
	return nil
}

func (n *QTestData) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQTestDataFromPointer(ptr unsafe.Pointer) (n *QTestData) {
	n = new(QTestData)
	n.InitFromInternal(uintptr(ptr), "testlib.QTestData")
	return
}

func (ptr *QTestData) DestroyQTestData() {
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

func (n *QTestDelayEvent) InitFromInternal(ptr uintptr, name string) {
	n.QTestEvent_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTestDelayEvent) ClassNameInternalF() string {
	return n.QTestEvent_PTR().ClassNameInternalF()
}

func NewQTestDelayEventFromPointer(ptr unsafe.Pointer) (n *QTestDelayEvent) {
	n = new(QTestDelayEvent)
	n.InitFromInternal(uintptr(ptr), "testlib.QTestDelayEvent")
	return
}

func (ptr *QTestDelayEvent) DestroyQTestDelayEvent() {
}

type QTestEvent struct {
	internal.Internal
}

type QTestEvent_ITF interface {
	QTestEvent_PTR() *QTestEvent
}

func (ptr *QTestEvent) QTestEvent_PTR() *QTestEvent {
	return ptr
}

func (ptr *QTestEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QTestEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQTestEvent(ptr QTestEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEvent_PTR().Pointer()
	}
	return nil
}

func (n *QTestEvent) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQTestEventFromPointer(ptr unsafe.Pointer) (n *QTestEvent) {
	n = new(QTestEvent)
	n.InitFromInternal(uintptr(ptr), "testlib.QTestEvent")
	return
}

func (ptr *QTestEvent) DestroyQTestEvent() {
}

type QTestEventList struct {
	internal.Internal
}

type QTestEventList_ITF interface {
	QTestEventList_PTR() *QTestEventList
}

func (ptr *QTestEventList) QTestEventList_PTR() *QTestEventList {
	return ptr
}

func (ptr *QTestEventList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QTestEventList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQTestEventList(ptr QTestEventList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEventList_PTR().Pointer()
	}
	return nil
}

func (n *QTestEventList) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQTestEventListFromPointer(ptr unsafe.Pointer) (n *QTestEventList) {
	n = new(QTestEventList)
	n.InitFromInternal(uintptr(ptr), "testlib.QTestEventList")
	return
}
func NewQTestEventList() *QTestEventList {

	return internal.CallLocalFunction([]interface{}{"", "", "testlib.NewQTestEventList", ""}).(*QTestEventList)
}

func NewQTestEventList2(other QTestEventList_ITF) *QTestEventList {

	return internal.CallLocalFunction([]interface{}{"", "", "testlib.NewQTestEventList2", "", other}).(*QTestEventList)
}

func (ptr *QTestEventList) AddDelay(msecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddDelay", msecs})
}

func (ptr *QTestEventList) AddKeyClick(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddKeyClick", qtKey, modifiers, msecs})
}

func (ptr *QTestEventList) AddKeyClick2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddKeyClick2", ascii, modifiers, msecs})
}

func (ptr *QTestEventList) AddKeyClicks(keys string, modifiers core.Qt__KeyboardModifier, msecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddKeyClicks", keys, modifiers, msecs})
}

func (ptr *QTestEventList) AddKeyPress(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddKeyPress", qtKey, modifiers, msecs})
}

func (ptr *QTestEventList) AddKeyPress2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddKeyPress2", ascii, modifiers, msecs})
}

func (ptr *QTestEventList) AddKeyRelease(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddKeyRelease", qtKey, modifiers, msecs})
}

func (ptr *QTestEventList) AddKeyRelease2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddKeyRelease2", ascii, modifiers, msecs})
}

func (ptr *QTestEventList) AddMouseClick(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddMouseClick", button, modifiers, pos, delay})
}

func (ptr *QTestEventList) AddMouseDClick(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddMouseDClick", button, modifiers, pos, delay})
}

func (ptr *QTestEventList) AddMouseMove(pos core.QPoint_ITF, delay int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddMouseMove", pos, delay})
}

func (ptr *QTestEventList) AddMousePress(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddMousePress", button, modifiers, pos, delay})
}

func (ptr *QTestEventList) AddMouseRelease(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddMouseRelease", button, modifiers, pos, delay})
}

func (ptr *QTestEventList) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QTestEventList) Simulate(w widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Simulate", w})
}

func (ptr *QTestEventList) DestroyQTestEventList() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTestEventList"})
}

type QTestEventLoop struct {
	internal.Internal
}

type QTestEventLoop_ITF interface {
	QTestEventLoop_PTR() *QTestEventLoop
}

func (ptr *QTestEventLoop) QTestEventLoop_PTR() *QTestEventLoop {
	return ptr
}

func (ptr *QTestEventLoop) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QTestEventLoop) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQTestEventLoop(ptr QTestEventLoop_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEventLoop_PTR().Pointer()
	}
	return nil
}

func (n *QTestEventLoop) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQTestEventLoopFromPointer(ptr unsafe.Pointer) (n *QTestEventLoop) {
	n = new(QTestEventLoop)
	n.InitFromInternal(uintptr(ptr), "testlib.QTestEventLoop")
	return
}

func (ptr *QTestEventLoop) DestroyQTestEventLoop() {
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

func (n *QTestKeyClicksEvent) InitFromInternal(ptr uintptr, name string) {
	n.QTestEvent_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTestKeyClicksEvent) ClassNameInternalF() string {
	return n.QTestEvent_PTR().ClassNameInternalF()
}

func NewQTestKeyClicksEventFromPointer(ptr unsafe.Pointer) (n *QTestKeyClicksEvent) {
	n = new(QTestKeyClicksEvent)
	n.InitFromInternal(uintptr(ptr), "testlib.QTestKeyClicksEvent")
	return
}

func (ptr *QTestKeyClicksEvent) DestroyQTestKeyClicksEvent() {
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

func (n *QTestKeyEvent) InitFromInternal(ptr uintptr, name string) {
	n.QTestEvent_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTestKeyEvent) ClassNameInternalF() string {
	return n.QTestEvent_PTR().ClassNameInternalF()
}

func NewQTestKeyEventFromPointer(ptr unsafe.Pointer) (n *QTestKeyEvent) {
	n = new(QTestKeyEvent)
	n.InitFromInternal(uintptr(ptr), "testlib.QTestKeyEvent")
	return
}

func (ptr *QTestKeyEvent) DestroyQTestKeyEvent() {
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

func (n *QTestMouseEvent) InitFromInternal(ptr uintptr, name string) {
	n.QTestEvent_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTestMouseEvent) ClassNameInternalF() string {
	return n.QTestEvent_PTR().ClassNameInternalF()
}

func NewQTestMouseEventFromPointer(ptr unsafe.Pointer) (n *QTestMouseEvent) {
	n = new(QTestMouseEvent)
	n.InitFromInternal(uintptr(ptr), "testlib.QTestMouseEvent")
	return
}

func (ptr *QTestMouseEvent) DestroyQTestMouseEvent() {
}

func init() {
	internal.ConstructorTable["testlib.QAbstractItemModelTester"] = NewQAbstractItemModelTesterFromPointer
	internal.ConstructorTable["testlib.QEventSizeOfChecker"] = NewQEventSizeOfCheckerFromPointer
	internal.ConstructorTable["testlib.QSignalSpy"] = NewQSignalSpyFromPointer
	internal.ConstructorTable["testlib.QSpontaneKeyEvent"] = NewQSpontaneKeyEventFromPointer
	internal.ConstructorTable["testlib.QTestData"] = NewQTestDataFromPointer
	internal.ConstructorTable["testlib.QTestDelayEvent"] = NewQTestDelayEventFromPointer
	internal.ConstructorTable["testlib.QTestEvent"] = NewQTestEventFromPointer
	internal.ConstructorTable["testlib.QTestEventList"] = NewQTestEventListFromPointer
	internal.ConstructorTable["testlib.QTestEventLoop"] = NewQTestEventLoopFromPointer
	internal.ConstructorTable["testlib.QTestKeyClicksEvent"] = NewQTestKeyClicksEventFromPointer
	internal.ConstructorTable["testlib.QTestKeyEvent"] = NewQTestKeyEventFromPointer
	internal.ConstructorTable["testlib.QTestMouseEvent"] = NewQTestMouseEventFromPointer
}
