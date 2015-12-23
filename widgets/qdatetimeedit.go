package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDateTimeEdit struct {
	QAbstractSpinBox
}

type QDateTimeEdit_ITF interface {
	QAbstractSpinBox_ITF
	QDateTimeEdit_PTR() *QDateTimeEdit
}

func PointerFromQDateTimeEdit(ptr QDateTimeEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateTimeEdit_PTR().Pointer()
	}
	return nil
}

func NewQDateTimeEditFromPointer(ptr unsafe.Pointer) *QDateTimeEdit {
	var n = new(QDateTimeEdit)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDateTimeEdit_") {
		n.SetObjectName("QDateTimeEdit_" + qt.Identifier())
	}
	return n
}

func (ptr *QDateTimeEdit) QDateTimeEdit_PTR() *QDateTimeEdit {
	return ptr
}

//QDateTimeEdit::Section
type QDateTimeEdit__Section int64

const (
	QDateTimeEdit__NoSection         = QDateTimeEdit__Section(0x0000)
	QDateTimeEdit__AmPmSection       = QDateTimeEdit__Section(0x0001)
	QDateTimeEdit__MSecSection       = QDateTimeEdit__Section(0x0002)
	QDateTimeEdit__SecondSection     = QDateTimeEdit__Section(0x0004)
	QDateTimeEdit__MinuteSection     = QDateTimeEdit__Section(0x0008)
	QDateTimeEdit__HourSection       = QDateTimeEdit__Section(0x0010)
	QDateTimeEdit__DaySection        = QDateTimeEdit__Section(0x0100)
	QDateTimeEdit__MonthSection      = QDateTimeEdit__Section(0x0200)
	QDateTimeEdit__YearSection       = QDateTimeEdit__Section(0x0400)
	QDateTimeEdit__TimeSections_Mask = QDateTimeEdit__Section(QDateTimeEdit__AmPmSection | QDateTimeEdit__MSecSection | QDateTimeEdit__SecondSection | QDateTimeEdit__MinuteSection | QDateTimeEdit__HourSection)
	QDateTimeEdit__DateSections_Mask = QDateTimeEdit__Section(QDateTimeEdit__DaySection | QDateTimeEdit__MonthSection | QDateTimeEdit__YearSection)
)

func NewQDateTimeEdit3(date core.QDate_ITF, parent QWidget_ITF) *QDateTimeEdit {
	defer qt.Recovering("QDateTimeEdit::QDateTimeEdit")

	return NewQDateTimeEditFromPointer(C.QDateTimeEdit_NewQDateTimeEdit3(core.PointerFromQDate(date), PointerFromQWidget(parent)))
}

func NewQDateTimeEdit4(time core.QTime_ITF, parent QWidget_ITF) *QDateTimeEdit {
	defer qt.Recovering("QDateTimeEdit::QDateTimeEdit")

	return NewQDateTimeEditFromPointer(C.QDateTimeEdit_NewQDateTimeEdit4(core.PointerFromQTime(time), PointerFromQWidget(parent)))
}

func (ptr *QDateTimeEdit) CalendarPopup() bool {
	defer qt.Recovering("QDateTimeEdit::calendarPopup")

	if ptr.Pointer() != nil {
		return C.QDateTimeEdit_CalendarPopup(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDateTimeEdit) ClearMaximumDate() {
	defer qt.Recovering("QDateTimeEdit::clearMaximumDate")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMaximumDate(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ClearMaximumDateTime() {
	defer qt.Recovering("QDateTimeEdit::clearMaximumDateTime")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMaximumDateTime(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ClearMaximumTime() {
	defer qt.Recovering("QDateTimeEdit::clearMaximumTime")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMaximumTime(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ClearMinimumDate() {
	defer qt.Recovering("QDateTimeEdit::clearMinimumDate")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMinimumDate(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ClearMinimumDateTime() {
	defer qt.Recovering("QDateTimeEdit::clearMinimumDateTime")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMinimumDateTime(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ClearMinimumTime() {
	defer qt.Recovering("QDateTimeEdit::clearMinimumTime")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMinimumTime(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) CurrentSection() QDateTimeEdit__Section {
	defer qt.Recovering("QDateTimeEdit::currentSection")

	if ptr.Pointer() != nil {
		return QDateTimeEdit__Section(C.QDateTimeEdit_CurrentSection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTimeEdit) CurrentSectionIndex() int {
	defer qt.Recovering("QDateTimeEdit::currentSectionIndex")

	if ptr.Pointer() != nil {
		return int(C.QDateTimeEdit_CurrentSectionIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTimeEdit) DateTime() *core.QDateTime {
	defer qt.Recovering("QDateTimeEdit::dateTime")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QDateTimeEdit_DateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTimeEdit) DisplayFormat() string {
	defer qt.Recovering("QDateTimeEdit::displayFormat")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTimeEdit_DisplayFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDateTimeEdit) DisplayedSections() QDateTimeEdit__Section {
	defer qt.Recovering("QDateTimeEdit::displayedSections")

	if ptr.Pointer() != nil {
		return QDateTimeEdit__Section(C.QDateTimeEdit_DisplayedSections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTimeEdit) MaximumDateTime() *core.QDateTime {
	defer qt.Recovering("QDateTimeEdit::maximumDateTime")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QDateTimeEdit_MaximumDateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTimeEdit) MinimumDateTime() *core.QDateTime {
	defer qt.Recovering("QDateTimeEdit::minimumDateTime")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QDateTimeEdit_MinimumDateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTimeEdit) SectionCount() int {
	defer qt.Recovering("QDateTimeEdit::sectionCount")

	if ptr.Pointer() != nil {
		return int(C.QDateTimeEdit_SectionCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTimeEdit) SectionText(section QDateTimeEdit__Section) string {
	defer qt.Recovering("QDateTimeEdit::sectionText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTimeEdit_SectionText(ptr.Pointer(), C.int(section)))
	}
	return ""
}

func (ptr *QDateTimeEdit) SetCalendarPopup(enable bool) {
	defer qt.Recovering("QDateTimeEdit::setCalendarPopup")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCalendarPopup(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDateTimeEdit) SetCurrentSection(section QDateTimeEdit__Section) {
	defer qt.Recovering("QDateTimeEdit::setCurrentSection")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCurrentSection(ptr.Pointer(), C.int(section))
	}
}

func (ptr *QDateTimeEdit) SetCurrentSectionIndex(index int) {
	defer qt.Recovering("QDateTimeEdit::setCurrentSectionIndex")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCurrentSectionIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QDateTimeEdit) SetDate(date core.QDate_ITF) {
	defer qt.Recovering("QDateTimeEdit::setDate")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QDateTimeEdit) SetDateTime(dateTime core.QDateTime_ITF) {
	defer qt.Recovering("QDateTimeEdit::setDateTime")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDateTime(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QDateTimeEdit) SetDisplayFormat(format string) {
	defer qt.Recovering("QDateTimeEdit::setDisplayFormat")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDisplayFormat(ptr.Pointer(), C.CString(format))
	}
}

func (ptr *QDateTimeEdit) SetMaximumDate(max core.QDate_ITF) {
	defer qt.Recovering("QDateTimeEdit::setMaximumDate")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMaximumDate(ptr.Pointer(), core.PointerFromQDate(max))
	}
}

func (ptr *QDateTimeEdit) SetMaximumDateTime(dt core.QDateTime_ITF) {
	defer qt.Recovering("QDateTimeEdit::setMaximumDateTime")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMaximumDateTime(ptr.Pointer(), core.PointerFromQDateTime(dt))
	}
}

func (ptr *QDateTimeEdit) SetMaximumTime(max core.QTime_ITF) {
	defer qt.Recovering("QDateTimeEdit::setMaximumTime")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMaximumTime(ptr.Pointer(), core.PointerFromQTime(max))
	}
}

func (ptr *QDateTimeEdit) SetMinimumDate(min core.QDate_ITF) {
	defer qt.Recovering("QDateTimeEdit::setMinimumDate")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMinimumDate(ptr.Pointer(), core.PointerFromQDate(min))
	}
}

func (ptr *QDateTimeEdit) SetMinimumDateTime(dt core.QDateTime_ITF) {
	defer qt.Recovering("QDateTimeEdit::setMinimumDateTime")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMinimumDateTime(ptr.Pointer(), core.PointerFromQDateTime(dt))
	}
}

func (ptr *QDateTimeEdit) SetMinimumTime(min core.QTime_ITF) {
	defer qt.Recovering("QDateTimeEdit::setMinimumTime")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMinimumTime(ptr.Pointer(), core.PointerFromQTime(min))
	}
}

func (ptr *QDateTimeEdit) SetTime(time core.QTime_ITF) {
	defer qt.Recovering("QDateTimeEdit::setTime")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetTime(ptr.Pointer(), core.PointerFromQTime(time))
	}
}

func (ptr *QDateTimeEdit) SetTimeSpec(spec core.Qt__TimeSpec) {
	defer qt.Recovering("QDateTimeEdit::setTimeSpec")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetTimeSpec(ptr.Pointer(), C.int(spec))
	}
}

func (ptr *QDateTimeEdit) TimeSpec() core.Qt__TimeSpec {
	defer qt.Recovering("QDateTimeEdit::timeSpec")

	if ptr.Pointer() != nil {
		return core.Qt__TimeSpec(C.QDateTimeEdit_TimeSpec(ptr.Pointer()))
	}
	return 0
}

func NewQDateTimeEdit(parent QWidget_ITF) *QDateTimeEdit {
	defer qt.Recovering("QDateTimeEdit::QDateTimeEdit")

	return NewQDateTimeEditFromPointer(C.QDateTimeEdit_NewQDateTimeEdit(PointerFromQWidget(parent)))
}

func NewQDateTimeEdit2(datetime core.QDateTime_ITF, parent QWidget_ITF) *QDateTimeEdit {
	defer qt.Recovering("QDateTimeEdit::QDateTimeEdit")

	return NewQDateTimeEditFromPointer(C.QDateTimeEdit_NewQDateTimeEdit2(core.PointerFromQDateTime(datetime), PointerFromQWidget(parent)))
}

func (ptr *QDateTimeEdit) CalendarWidget() *QCalendarWidget {
	defer qt.Recovering("QDateTimeEdit::calendarWidget")

	if ptr.Pointer() != nil {
		return NewQCalendarWidgetFromPointer(C.QDateTimeEdit_CalendarWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTimeEdit) ConnectClear(f func()) {
	defer qt.Recovering("connect QDateTimeEdit::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectClear() {
	defer qt.Recovering("disconnect QDateTimeEdit::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQDateTimeEditClear
func callbackQDateTimeEditClear(ptrName *C.char) bool {
	defer qt.Recovering("callback QDateTimeEdit::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectDateTimeChanged(f func(datetime *core.QDateTime)) {
	defer qt.Recovering("connect QDateTimeEdit::dateTimeChanged")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ConnectDateTimeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dateTimeChanged", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectDateTimeChanged() {
	defer qt.Recovering("disconnect QDateTimeEdit::dateTimeChanged")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_DisconnectDateTimeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dateTimeChanged")
	}
}

//export callbackQDateTimeEditDateTimeChanged
func callbackQDateTimeEditDateTimeChanged(ptrName *C.char, datetime unsafe.Pointer) {
	defer qt.Recovering("callback QDateTimeEdit::dateTimeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "dateTimeChanged"); signal != nil {
		signal.(func(*core.QDateTime))(core.NewQDateTimeFromPointer(datetime))
	}

}

func (ptr *QDateTimeEdit) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QDateTimeEdit::event")

	if ptr.Pointer() != nil {
		return C.QDateTimeEdit_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDateTimeEdit) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQDateTimeEditFocusInEvent
func callbackQDateTimeEditFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQDateTimeEditKeyPressEvent
func callbackQDateTimeEditKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDateTimeEditMousePressEvent
func callbackQDateTimeEditMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDateTimeEditPaintEvent
func callbackQDateTimeEditPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) SectionAt(index int) QDateTimeEdit__Section {
	defer qt.Recovering("QDateTimeEdit::sectionAt")

	if ptr.Pointer() != nil {
		return QDateTimeEdit__Section(C.QDateTimeEdit_SectionAt(ptr.Pointer(), C.int(index)))
	}
	return 0
}

func (ptr *QDateTimeEdit) SetCalendarWidget(calendarWidget QCalendarWidget_ITF) {
	defer qt.Recovering("QDateTimeEdit::setCalendarWidget")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCalendarWidget(ptr.Pointer(), PointerFromQCalendarWidget(calendarWidget))
	}
}

func (ptr *QDateTimeEdit) SetDateRange(min core.QDate_ITF, max core.QDate_ITF) {
	defer qt.Recovering("QDateTimeEdit::setDateRange")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDateRange(ptr.Pointer(), core.PointerFromQDate(min), core.PointerFromQDate(max))
	}
}

func (ptr *QDateTimeEdit) SetDateTimeRange(min core.QDateTime_ITF, max core.QDateTime_ITF) {
	defer qt.Recovering("QDateTimeEdit::setDateTimeRange")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDateTimeRange(ptr.Pointer(), core.PointerFromQDateTime(min), core.PointerFromQDateTime(max))
	}
}

func (ptr *QDateTimeEdit) SetSelectedSection(section QDateTimeEdit__Section) {
	defer qt.Recovering("QDateTimeEdit::setSelectedSection")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetSelectedSection(ptr.Pointer(), C.int(section))
	}
}

func (ptr *QDateTimeEdit) SetTimeRange(min core.QTime_ITF, max core.QTime_ITF) {
	defer qt.Recovering("QDateTimeEdit::setTimeRange")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetTimeRange(ptr.Pointer(), core.PointerFromQTime(min), core.PointerFromQTime(max))
	}
}

func (ptr *QDateTimeEdit) SizeHint() *core.QSize {
	defer qt.Recovering("QDateTimeEdit::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QDateTimeEdit_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTimeEdit) ConnectStepBy(f func(steps int)) {
	defer qt.Recovering("connect QDateTimeEdit::stepBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stepBy", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectStepBy() {
	defer qt.Recovering("disconnect QDateTimeEdit::stepBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stepBy")
	}
}

//export callbackQDateTimeEditStepBy
func callbackQDateTimeEditStepBy(ptrName *C.char, steps C.int) bool {
	defer qt.Recovering("callback QDateTimeEdit::stepBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "stepBy"); signal != nil {
		signal.(func(int))(int(steps))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQDateTimeEditWheelEvent
func callbackQDateTimeEditWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) DestroyQDateTimeEdit() {
	defer qt.Recovering("QDateTimeEdit::~QDateTimeEdit")

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_DestroyQDateTimeEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDateTimeEdit) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQDateTimeEditChangeEvent
func callbackQDateTimeEditChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQDateTimeEditCloseEvent
func callbackQDateTimeEditCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQDateTimeEditContextMenuEvent
func callbackQDateTimeEditContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQDateTimeEditFocusOutEvent
func callbackQDateTimeEditFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQDateTimeEditHideEvent
func callbackQDateTimeEditHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQDateTimeEditKeyReleaseEvent
func callbackQDateTimeEditKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDateTimeEditMouseMoveEvent
func callbackQDateTimeEditMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDateTimeEditMouseReleaseEvent
func callbackQDateTimeEditMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDateTimeEditResizeEvent
func callbackQDateTimeEditResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQDateTimeEditShowEvent
func callbackQDateTimeEditShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDateTimeEditTimerEvent
func callbackQDateTimeEditTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQDateTimeEditActionEvent
func callbackQDateTimeEditActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQDateTimeEditDragEnterEvent
func callbackQDateTimeEditDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQDateTimeEditDragLeaveEvent
func callbackQDateTimeEditDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQDateTimeEditDragMoveEvent
func callbackQDateTimeEditDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQDateTimeEditDropEvent
func callbackQDateTimeEditDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQDateTimeEditEnterEvent
func callbackQDateTimeEditEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQDateTimeEditLeaveEvent
func callbackQDateTimeEditLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQDateTimeEditMoveEvent
func callbackQDateTimeEditMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDateTimeEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDateTimeEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQDateTimeEditSetVisible
func callbackQDateTimeEditSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDateTimeEdit::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QDateTimeEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QDateTimeEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQDateTimeEditInitPainter
func callbackQDateTimeEditInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQDateTimeEditInputMethodEvent
func callbackQDateTimeEditInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQDateTimeEditMouseDoubleClickEvent
func callbackQDateTimeEditMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQDateTimeEditTabletEvent
func callbackQDateTimeEditTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDateTimeEditChildEvent
func callbackQDateTimeEditChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDateTimeEdit) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDateTimeEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDateTimeEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDateTimeEditCustomEvent
func callbackQDateTimeEditCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDateTimeEdit::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
