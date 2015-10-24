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

type QCalendarWidgetITF interface {
	QWidgetITF
	QCalendarWidgetPTR() *QCalendarWidget
}

func PointerFromQCalendarWidget(ptr QCalendarWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCalendarWidgetPTR().Pointer()
	}
	return nil
}

func QCalendarWidgetFromPointer(ptr unsafe.Pointer) *QCalendarWidget {
	var n = new(QCalendarWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCalendarWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCalendarWidget) QCalendarWidgetPTR() *QCalendarWidget {
	return ptr
}

//QCalendarWidget::HorizontalHeaderFormat
type QCalendarWidget__HorizontalHeaderFormat int

var (
	QCalendarWidget__NoHorizontalHeader   = QCalendarWidget__HorizontalHeaderFormat(0)
	QCalendarWidget__SingleLetterDayNames = QCalendarWidget__HorizontalHeaderFormat(1)
	QCalendarWidget__ShortDayNames        = QCalendarWidget__HorizontalHeaderFormat(2)
	QCalendarWidget__LongDayNames         = QCalendarWidget__HorizontalHeaderFormat(3)
)

//QCalendarWidget::SelectionMode
type QCalendarWidget__SelectionMode int

var (
	QCalendarWidget__NoSelection     = QCalendarWidget__SelectionMode(0)
	QCalendarWidget__SingleSelection = QCalendarWidget__SelectionMode(1)
)

//QCalendarWidget::VerticalHeaderFormat
type QCalendarWidget__VerticalHeaderFormat int

var (
	QCalendarWidget__NoVerticalHeader = QCalendarWidget__VerticalHeaderFormat(0)
	QCalendarWidget__ISOWeekNumbers   = QCalendarWidget__VerticalHeaderFormat(1)
)

func (ptr *QCalendarWidget) DateEditAcceptDelay() int {
	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_DateEditAcceptDelay(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCalendarWidget) FirstDayOfWeek() core.Qt__DayOfWeek {
	if ptr.Pointer() != nil {
		return core.Qt__DayOfWeek(C.QCalendarWidget_FirstDayOfWeek(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCalendarWidget) HorizontalHeaderFormat() QCalendarWidget__HorizontalHeaderFormat {
	if ptr.Pointer() != nil {
		return QCalendarWidget__HorizontalHeaderFormat(C.QCalendarWidget_HorizontalHeaderFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCalendarWidget) IsDateEditEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsDateEditEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCalendarWidget) IsGridVisible() bool {
	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsGridVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCalendarWidget) IsNavigationBarVisible() bool {
	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsNavigationBarVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCalendarWidget) SelectionMode() QCalendarWidget__SelectionMode {
	if ptr.Pointer() != nil {
		return QCalendarWidget__SelectionMode(C.QCalendarWidget_SelectionMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCalendarWidget) SetDateEditAcceptDelay(delay int) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateEditAcceptDelay(C.QtObjectPtr(ptr.Pointer()), C.int(delay))
	}
}

func (ptr *QCalendarWidget) SetDateEditEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateEditEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QCalendarWidget) SetFirstDayOfWeek(dayOfWeek core.Qt__DayOfWeek) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetFirstDayOfWeek(C.QtObjectPtr(ptr.Pointer()), C.int(dayOfWeek))
	}
}

func (ptr *QCalendarWidget) SetGridVisible(show bool) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetGridVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QCalendarWidget) SetHorizontalHeaderFormat(format QCalendarWidget__HorizontalHeaderFormat) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetHorizontalHeaderFormat(C.QtObjectPtr(ptr.Pointer()), C.int(format))
	}
}

func (ptr *QCalendarWidget) SetMaximumDate(date core.QDateITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetMaximumDate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDate(date)))
	}
}

func (ptr *QCalendarWidget) SetMinimumDate(date core.QDateITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetMinimumDate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDate(date)))
	}
}

func (ptr *QCalendarWidget) SetNavigationBarVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetNavigationBarVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCalendarWidget) SetSelectedDate(date core.QDateITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetSelectedDate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDate(date)))
	}
}

func (ptr *QCalendarWidget) SetSelectionMode(mode QCalendarWidget__SelectionMode) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetSelectionMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QCalendarWidget) SetVerticalHeaderFormat(format QCalendarWidget__VerticalHeaderFormat) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetVerticalHeaderFormat(C.QtObjectPtr(ptr.Pointer()), C.int(format))
	}
}

func (ptr *QCalendarWidget) VerticalHeaderFormat() QCalendarWidget__VerticalHeaderFormat {
	if ptr.Pointer() != nil {
		return QCalendarWidget__VerticalHeaderFormat(C.QCalendarWidget_VerticalHeaderFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQCalendarWidget(parent QWidgetITF) *QCalendarWidget {
	return QCalendarWidgetFromPointer(unsafe.Pointer(C.QCalendarWidget_NewQCalendarWidget(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QCalendarWidget) ConnectCurrentPageChanged(f func(year int, month int)) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ConnectCurrentPageChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentPageChanged", f)
	}
}

func (ptr *QCalendarWidget) DisconnectCurrentPageChanged() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_DisconnectCurrentPageChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentPageChanged")
	}
}

//export callbackQCalendarWidgetCurrentPageChanged
func callbackQCalendarWidgetCurrentPageChanged(ptrName *C.char, year C.int, month C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentPageChanged").(func(int, int))(int(year), int(month))
}

func (ptr *QCalendarWidget) MonthShown() int {
	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_MonthShown(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCalendarWidget) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ConnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QCalendarWidget) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_DisconnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQCalendarWidgetSelectionChanged
func callbackQCalendarWidgetSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QCalendarWidget) SetCurrentPage(year int, month int) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetCurrentPage(C.QtObjectPtr(ptr.Pointer()), C.int(year), C.int(month))
	}
}

func (ptr *QCalendarWidget) SetDateRange(min core.QDateITF, max core.QDateITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateRange(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDate(min)), C.QtObjectPtr(core.PointerFromQDate(max)))
	}
}

func (ptr *QCalendarWidget) SetDateTextFormat(date core.QDateITF, format gui.QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateTextFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDate(date)), C.QtObjectPtr(gui.PointerFromQTextCharFormat(format)))
	}
}

func (ptr *QCalendarWidget) SetHeaderTextFormat(format gui.QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetHeaderTextFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextCharFormat(format)))
	}
}

func (ptr *QCalendarWidget) SetWeekdayTextFormat(dayOfWeek core.Qt__DayOfWeek, format gui.QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetWeekdayTextFormat(C.QtObjectPtr(ptr.Pointer()), C.int(dayOfWeek), C.QtObjectPtr(gui.PointerFromQTextCharFormat(format)))
	}
}

func (ptr *QCalendarWidget) ShowNextMonth() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowNextMonth(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCalendarWidget) ShowNextYear() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowNextYear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCalendarWidget) ShowPreviousMonth() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowPreviousMonth(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCalendarWidget) ShowPreviousYear() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowPreviousYear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCalendarWidget) ShowSelectedDate() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowSelectedDate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCalendarWidget) ShowToday() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowToday(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCalendarWidget) YearShown() int {
	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_YearShown(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCalendarWidget) DestroyQCalendarWidget() {
	if ptr.Pointer() != nil {
		C.QCalendarWidget_DestroyQCalendarWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
