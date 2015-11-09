package widgets

//#include "qcalendarwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QCalendarWidget struct {
	QWidget
}

type QCalendarWidget_ITF interface {
	QWidget_ITF
	QCalendarWidget_PTR() *QCalendarWidget
}

func PointerFromQCalendarWidget(ptr QCalendarWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCalendarWidget_PTR().Pointer()
	}
	return nil
}

func NewQCalendarWidgetFromPointer(ptr unsafe.Pointer) *QCalendarWidget {
	var n = new(QCalendarWidget)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QCalendarWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCalendarWidget) QCalendarWidget_PTR() *QCalendarWidget {
	return ptr
}

//QCalendarWidget::HorizontalHeaderFormat
type QCalendarWidget__HorizontalHeaderFormat int64

const (
	QCalendarWidget__NoHorizontalHeader   = QCalendarWidget__HorizontalHeaderFormat(0)
	QCalendarWidget__SingleLetterDayNames = QCalendarWidget__HorizontalHeaderFormat(1)
	QCalendarWidget__ShortDayNames        = QCalendarWidget__HorizontalHeaderFormat(2)
	QCalendarWidget__LongDayNames         = QCalendarWidget__HorizontalHeaderFormat(3)
)

//QCalendarWidget::SelectionMode
type QCalendarWidget__SelectionMode int64

const (
	QCalendarWidget__NoSelection     = QCalendarWidget__SelectionMode(0)
	QCalendarWidget__SingleSelection = QCalendarWidget__SelectionMode(1)
)

//QCalendarWidget::VerticalHeaderFormat
type QCalendarWidget__VerticalHeaderFormat int64

const (
	QCalendarWidget__NoVerticalHeader = QCalendarWidget__VerticalHeaderFormat(0)
	QCalendarWidget__ISOWeekNumbers   = QCalendarWidget__VerticalHeaderFormat(1)
)

func (ptr *QCalendarWidget) DateEditAcceptDelay() int {
	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_DateEditAcceptDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) FirstDayOfWeek() core.Qt__DayOfWeek {
	if ptr.Pointer() != nil {
		return core.Qt__DayOfWeek(C.QCalendarWidget_FirstDayOfWeek(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) HorizontalHeaderFormat() QCalendarWidget__HorizontalHeaderFormat {
	if ptr.Pointer() != nil {
		return QCalendarWidget__HorizontalHeaderFormat(C.QCalendarWidget_HorizontalHeaderFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) IsDateEditEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsDateEditEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCalendarWidget) IsGridVisible() bool {
	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsGridVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCalendarWidget) IsNavigationBarVisible() bool {
	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsNavigationBarVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCalendarWidget) SelectionMode() QCalendarWidget__SelectionMode {
	if ptr.Pointer() != nil {
		return QCalendarWidget__SelectionMode(C.QCalendarWidget_SelectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) SetDateEditAcceptDelay(delay int) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateEditAcceptDelay(ptr.Pointer(), C.int(delay))
	}
}

func (ptr *QCalendarWidget) SetDateEditEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateEditEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QCalendarWidget) SetFirstDayOfWeek(dayOfWeek core.Qt__DayOfWeek) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetFirstDayOfWeek(ptr.Pointer(), C.int(dayOfWeek))
	}
}

func (ptr *QCalendarWidget) SetGridVisible(show bool) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetGridVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QCalendarWidget) SetHorizontalHeaderFormat(format QCalendarWidget__HorizontalHeaderFormat) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetHorizontalHeaderFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCalendarWidget) SetMaximumDate(date core.QDate_ITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetMaximumDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QCalendarWidget) SetMinimumDate(date core.QDate_ITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetMinimumDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QCalendarWidget) SetNavigationBarVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetNavigationBarVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCalendarWidget) SetSelectedDate(date core.QDate_ITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetSelectedDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QCalendarWidget) SetSelectionMode(mode QCalendarWidget__SelectionMode) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetSelectionMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCalendarWidget) SetVerticalHeaderFormat(format QCalendarWidget__VerticalHeaderFormat) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetVerticalHeaderFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCalendarWidget) VerticalHeaderFormat() QCalendarWidget__VerticalHeaderFormat {
	if ptr.Pointer() != nil {
		return QCalendarWidget__VerticalHeaderFormat(C.QCalendarWidget_VerticalHeaderFormat(ptr.Pointer()))
	}
	return 0
}

func NewQCalendarWidget(parent QWidget_ITF) *QCalendarWidget {
	return NewQCalendarWidgetFromPointer(C.QCalendarWidget_NewQCalendarWidget(PointerFromQWidget(parent)))
}

func (ptr *QCalendarWidget) ConnectCurrentPageChanged(f func(year int, month int)) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ConnectCurrentPageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentPageChanged", f)
	}
}

func (ptr *QCalendarWidget) DisconnectCurrentPageChanged() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_DisconnectCurrentPageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentPageChanged")
	}
}

//export callbackQCalendarWidgetCurrentPageChanged
func callbackQCalendarWidgetCurrentPageChanged(ptrName *C.char, year C.int, month C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentPageChanged").(func(int, int))(int(year), int(month))
}

func (ptr *QCalendarWidget) MonthShown() int {
	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_MonthShown(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QCalendarWidget) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQCalendarWidgetSelectionChanged
func callbackQCalendarWidgetSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QCalendarWidget) SetCurrentPage(year int, month int) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetCurrentPage(ptr.Pointer(), C.int(year), C.int(month))
	}
}

func (ptr *QCalendarWidget) SetDateRange(min core.QDate_ITF, max core.QDate_ITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateRange(ptr.Pointer(), core.PointerFromQDate(min), core.PointerFromQDate(max))
	}
}

func (ptr *QCalendarWidget) SetDateTextFormat(date core.QDate_ITF, format gui.QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateTextFormat(ptr.Pointer(), core.PointerFromQDate(date), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QCalendarWidget) SetHeaderTextFormat(format gui.QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetHeaderTextFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QCalendarWidget) SetWeekdayTextFormat(dayOfWeek core.Qt__DayOfWeek, format gui.QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetWeekdayTextFormat(ptr.Pointer(), C.int(dayOfWeek), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QCalendarWidget) ShowNextMonth() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowNextMonth(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowNextYear() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowNextYear(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowPreviousMonth() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowPreviousMonth(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowPreviousYear() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowPreviousYear(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowSelectedDate() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowSelectedDate(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowToday() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowToday(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) YearShown() int {
	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_YearShown(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) DestroyQCalendarWidget() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_DestroyQCalendarWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
