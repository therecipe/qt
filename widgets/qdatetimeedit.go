package widgets

//#include "qdatetimeedit.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDateTimeEdit struct {
	QAbstractSpinBox
}

type QDateTimeEditITF interface {
	QAbstractSpinBoxITF
	QDateTimeEditPTR() *QDateTimeEdit
}

func PointerFromQDateTimeEdit(ptr QDateTimeEditITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateTimeEditPTR().Pointer()
	}
	return nil
}

func QDateTimeEditFromPointer(ptr unsafe.Pointer) *QDateTimeEdit {
	var n = new(QDateTimeEdit)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDateTimeEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDateTimeEdit) QDateTimeEditPTR() *QDateTimeEdit {
	return ptr
}

//QDateTimeEdit::Section
type QDateTimeEdit__Section int

var (
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

func NewQDateTimeEdit3(date core.QDateITF, parent QWidgetITF) *QDateTimeEdit {
	return QDateTimeEditFromPointer(unsafe.Pointer(C.QDateTimeEdit_NewQDateTimeEdit3(C.QtObjectPtr(core.PointerFromQDate(date)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQDateTimeEdit4(time core.QTimeITF, parent QWidgetITF) *QDateTimeEdit {
	return QDateTimeEditFromPointer(unsafe.Pointer(C.QDateTimeEdit_NewQDateTimeEdit4(C.QtObjectPtr(core.PointerFromQTime(time)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QDateTimeEdit) CalendarPopup() bool {
	if ptr.Pointer() != nil {
		return C.QDateTimeEdit_CalendarPopup(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDateTimeEdit) ClearMaximumDate() {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMaximumDate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDateTimeEdit) ClearMaximumDateTime() {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMaximumDateTime(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDateTimeEdit) ClearMaximumTime() {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMaximumTime(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDateTimeEdit) ClearMinimumDate() {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMinimumDate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDateTimeEdit) ClearMinimumDateTime() {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMinimumDateTime(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDateTimeEdit) ClearMinimumTime() {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMinimumTime(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDateTimeEdit) CurrentSection() QDateTimeEdit__Section {
	if ptr.Pointer() != nil {
		return QDateTimeEdit__Section(C.QDateTimeEdit_CurrentSection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDateTimeEdit) CurrentSectionIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QDateTimeEdit_CurrentSectionIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDateTimeEdit) DisplayFormat() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTimeEdit_DisplayFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDateTimeEdit) DisplayedSections() QDateTimeEdit__Section {
	if ptr.Pointer() != nil {
		return QDateTimeEdit__Section(C.QDateTimeEdit_DisplayedSections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDateTimeEdit) SectionCount() int {
	if ptr.Pointer() != nil {
		return int(C.QDateTimeEdit_SectionCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDateTimeEdit) SectionText(section QDateTimeEdit__Section) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTimeEdit_SectionText(C.QtObjectPtr(ptr.Pointer()), C.int(section)))
	}
	return ""
}

func (ptr *QDateTimeEdit) SetCalendarPopup(enable bool) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCalendarPopup(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDateTimeEdit) SetCurrentSection(section QDateTimeEdit__Section) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCurrentSection(C.QtObjectPtr(ptr.Pointer()), C.int(section))
	}
}

func (ptr *QDateTimeEdit) SetCurrentSectionIndex(index int) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCurrentSectionIndex(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QDateTimeEdit) SetDate(date core.QDateITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDate(date)))
	}
}

func (ptr *QDateTimeEdit) SetDateTime(dateTime core.QDateTimeITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDateTime(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDateTime(dateTime)))
	}
}

func (ptr *QDateTimeEdit) SetDisplayFormat(format string) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDisplayFormat(C.QtObjectPtr(ptr.Pointer()), C.CString(format))
	}
}

func (ptr *QDateTimeEdit) SetMaximumDate(max core.QDateITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMaximumDate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDate(max)))
	}
}

func (ptr *QDateTimeEdit) SetMaximumDateTime(dt core.QDateTimeITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMaximumDateTime(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDateTime(dt)))
	}
}

func (ptr *QDateTimeEdit) SetMaximumTime(max core.QTimeITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMaximumTime(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQTime(max)))
	}
}

func (ptr *QDateTimeEdit) SetMinimumDate(min core.QDateITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMinimumDate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDate(min)))
	}
}

func (ptr *QDateTimeEdit) SetMinimumDateTime(dt core.QDateTimeITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMinimumDateTime(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDateTime(dt)))
	}
}

func (ptr *QDateTimeEdit) SetMinimumTime(min core.QTimeITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMinimumTime(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQTime(min)))
	}
}

func (ptr *QDateTimeEdit) SetTime(time core.QTimeITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetTime(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQTime(time)))
	}
}

func (ptr *QDateTimeEdit) SetTimeSpec(spec core.Qt__TimeSpec) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetTimeSpec(C.QtObjectPtr(ptr.Pointer()), C.int(spec))
	}
}

func (ptr *QDateTimeEdit) TimeSpec() core.Qt__TimeSpec {
	if ptr.Pointer() != nil {
		return core.Qt__TimeSpec(C.QDateTimeEdit_TimeSpec(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQDateTimeEdit(parent QWidgetITF) *QDateTimeEdit {
	return QDateTimeEditFromPointer(unsafe.Pointer(C.QDateTimeEdit_NewQDateTimeEdit(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQDateTimeEdit2(datetime core.QDateTimeITF, parent QWidgetITF) *QDateTimeEdit {
	return QDateTimeEditFromPointer(unsafe.Pointer(C.QDateTimeEdit_NewQDateTimeEdit2(C.QtObjectPtr(core.PointerFromQDateTime(datetime)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QDateTimeEdit) CalendarWidget() *QCalendarWidget {
	if ptr.Pointer() != nil {
		return QCalendarWidgetFromPointer(unsafe.Pointer(C.QDateTimeEdit_CalendarWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDateTimeEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDateTimeEdit) Event(event core.QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QDateTimeEdit_Event(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QDateTimeEdit) SectionAt(index int) QDateTimeEdit__Section {
	if ptr.Pointer() != nil {
		return QDateTimeEdit__Section(C.QDateTimeEdit_SectionAt(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return 0
}

func (ptr *QDateTimeEdit) SetCalendarWidget(calendarWidget QCalendarWidgetITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCalendarWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCalendarWidget(calendarWidget)))
	}
}

func (ptr *QDateTimeEdit) SetDateRange(min core.QDateITF, max core.QDateITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDateRange(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDate(min)), C.QtObjectPtr(core.PointerFromQDate(max)))
	}
}

func (ptr *QDateTimeEdit) SetDateTimeRange(min core.QDateTimeITF, max core.QDateTimeITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDateTimeRange(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDateTime(min)), C.QtObjectPtr(core.PointerFromQDateTime(max)))
	}
}

func (ptr *QDateTimeEdit) SetSelectedSection(section QDateTimeEdit__Section) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetSelectedSection(C.QtObjectPtr(ptr.Pointer()), C.int(section))
	}
}

func (ptr *QDateTimeEdit) SetTimeRange(min core.QTimeITF, max core.QTimeITF) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetTimeRange(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQTime(min)), C.QtObjectPtr(core.PointerFromQTime(max)))
	}
}

func (ptr *QDateTimeEdit) StepBy(steps int) {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_StepBy(C.QtObjectPtr(ptr.Pointer()), C.int(steps))
	}
}

func (ptr *QDateTimeEdit) DestroyQDateTimeEdit() {
	if ptr.Pointer() != nil {
		C.QDateTimeEdit_DestroyQDateTimeEdit(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
