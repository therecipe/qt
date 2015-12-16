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

	var signal = qt.GetSignal(C.GoString(ptrName), "clear")
	if signal != nil {
		defer signal.(func())()
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

	var signal = qt.GetSignal(C.GoString(ptrName), "dateTimeChanged")
	if signal != nil {
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

	var signal = qt.GetSignal(C.GoString(ptrName), "focusInEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "stepBy")
	if signal != nil {
		defer signal.(func(int))(int(steps))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "wheelEvent")
	if signal != nil {
		defer signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
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
