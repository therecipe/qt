package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QCalendarWidget_") {
		n.SetObjectName("QCalendarWidget_" + qt.Identifier())
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
	defer qt.Recovering("QCalendarWidget::dateEditAcceptDelay")

	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_DateEditAcceptDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) FirstDayOfWeek() core.Qt__DayOfWeek {
	defer qt.Recovering("QCalendarWidget::firstDayOfWeek")

	if ptr.Pointer() != nil {
		return core.Qt__DayOfWeek(C.QCalendarWidget_FirstDayOfWeek(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) HorizontalHeaderFormat() QCalendarWidget__HorizontalHeaderFormat {
	defer qt.Recovering("QCalendarWidget::horizontalHeaderFormat")

	if ptr.Pointer() != nil {
		return QCalendarWidget__HorizontalHeaderFormat(C.QCalendarWidget_HorizontalHeaderFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) IsDateEditEnabled() bool {
	defer qt.Recovering("QCalendarWidget::isDateEditEnabled")

	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsDateEditEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCalendarWidget) IsGridVisible() bool {
	defer qt.Recovering("QCalendarWidget::isGridVisible")

	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsGridVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCalendarWidget) IsNavigationBarVisible() bool {
	defer qt.Recovering("QCalendarWidget::isNavigationBarVisible")

	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsNavigationBarVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCalendarWidget) SelectionMode() QCalendarWidget__SelectionMode {
	defer qt.Recovering("QCalendarWidget::selectionMode")

	if ptr.Pointer() != nil {
		return QCalendarWidget__SelectionMode(C.QCalendarWidget_SelectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) SetDateEditAcceptDelay(delay int) {
	defer qt.Recovering("QCalendarWidget::setDateEditAcceptDelay")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateEditAcceptDelay(ptr.Pointer(), C.int(delay))
	}
}

func (ptr *QCalendarWidget) SetDateEditEnabled(enable bool) {
	defer qt.Recovering("QCalendarWidget::setDateEditEnabled")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateEditEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QCalendarWidget) SetFirstDayOfWeek(dayOfWeek core.Qt__DayOfWeek) {
	defer qt.Recovering("QCalendarWidget::setFirstDayOfWeek")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetFirstDayOfWeek(ptr.Pointer(), C.int(dayOfWeek))
	}
}

func (ptr *QCalendarWidget) SetGridVisible(show bool) {
	defer qt.Recovering("QCalendarWidget::setGridVisible")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetGridVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QCalendarWidget) SetHorizontalHeaderFormat(format QCalendarWidget__HorizontalHeaderFormat) {
	defer qt.Recovering("QCalendarWidget::setHorizontalHeaderFormat")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetHorizontalHeaderFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCalendarWidget) SetMaximumDate(date core.QDate_ITF) {
	defer qt.Recovering("QCalendarWidget::setMaximumDate")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetMaximumDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QCalendarWidget) SetMinimumDate(date core.QDate_ITF) {
	defer qt.Recovering("QCalendarWidget::setMinimumDate")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetMinimumDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QCalendarWidget) SetNavigationBarVisible(visible bool) {
	defer qt.Recovering("QCalendarWidget::setNavigationBarVisible")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetNavigationBarVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCalendarWidget) SetSelectedDate(date core.QDate_ITF) {
	defer qt.Recovering("QCalendarWidget::setSelectedDate")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetSelectedDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QCalendarWidget) SetSelectionMode(mode QCalendarWidget__SelectionMode) {
	defer qt.Recovering("QCalendarWidget::setSelectionMode")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetSelectionMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCalendarWidget) SetVerticalHeaderFormat(format QCalendarWidget__VerticalHeaderFormat) {
	defer qt.Recovering("QCalendarWidget::setVerticalHeaderFormat")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetVerticalHeaderFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCalendarWidget) VerticalHeaderFormat() QCalendarWidget__VerticalHeaderFormat {
	defer qt.Recovering("QCalendarWidget::verticalHeaderFormat")

	if ptr.Pointer() != nil {
		return QCalendarWidget__VerticalHeaderFormat(C.QCalendarWidget_VerticalHeaderFormat(ptr.Pointer()))
	}
	return 0
}

func NewQCalendarWidget(parent QWidget_ITF) *QCalendarWidget {
	defer qt.Recovering("QCalendarWidget::QCalendarWidget")

	return NewQCalendarWidgetFromPointer(C.QCalendarWidget_NewQCalendarWidget(PointerFromQWidget(parent)))
}

func (ptr *QCalendarWidget) ConnectCurrentPageChanged(f func(year int, month int)) {
	defer qt.Recovering("connect QCalendarWidget::currentPageChanged")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ConnectCurrentPageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentPageChanged", f)
	}
}

func (ptr *QCalendarWidget) DisconnectCurrentPageChanged() {
	defer qt.Recovering("disconnect QCalendarWidget::currentPageChanged")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DisconnectCurrentPageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentPageChanged")
	}
}

//export callbackQCalendarWidgetCurrentPageChanged
func callbackQCalendarWidgetCurrentPageChanged(ptrName *C.char, year C.int, month C.int) {
	defer qt.Recovering("callback QCalendarWidget::currentPageChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentPageChanged")
	if signal != nil {
		signal.(func(int, int))(int(year), int(month))
	}

}

func (ptr *QCalendarWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCalendarWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQCalendarWidgetKeyPressEvent
func callbackQCalendarWidgetKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCalendarWidget::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCalendarWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QCalendarWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCalendarWidget_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCalendarWidget) MonthShown() int {
	defer qt.Recovering("QCalendarWidget::monthShown")

	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_MonthShown(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCalendarWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQCalendarWidgetMousePressEvent
func callbackQCalendarWidgetMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCalendarWidget::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCalendarWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QCalendarWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQCalendarWidgetResizeEvent
func callbackQCalendarWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCalendarWidget::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCalendarWidget) ConnectSelectionChanged(f func()) {
	defer qt.Recovering("connect QCalendarWidget::selectionChanged")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QCalendarWidget) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QCalendarWidget::selectionChanged")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQCalendarWidgetSelectionChanged
func callbackQCalendarWidgetSelectionChanged(ptrName *C.char) {
	defer qt.Recovering("callback QCalendarWidget::selectionChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "selectionChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QCalendarWidget) SetCurrentPage(year int, month int) {
	defer qt.Recovering("QCalendarWidget::setCurrentPage")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetCurrentPage(ptr.Pointer(), C.int(year), C.int(month))
	}
}

func (ptr *QCalendarWidget) SetDateRange(min core.QDate_ITF, max core.QDate_ITF) {
	defer qt.Recovering("QCalendarWidget::setDateRange")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateRange(ptr.Pointer(), core.PointerFromQDate(min), core.PointerFromQDate(max))
	}
}

func (ptr *QCalendarWidget) SetDateTextFormat(date core.QDate_ITF, format gui.QTextCharFormat_ITF) {
	defer qt.Recovering("QCalendarWidget::setDateTextFormat")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateTextFormat(ptr.Pointer(), core.PointerFromQDate(date), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QCalendarWidget) SetHeaderTextFormat(format gui.QTextCharFormat_ITF) {
	defer qt.Recovering("QCalendarWidget::setHeaderTextFormat")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetHeaderTextFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QCalendarWidget) SetWeekdayTextFormat(dayOfWeek core.Qt__DayOfWeek, format gui.QTextCharFormat_ITF) {
	defer qt.Recovering("QCalendarWidget::setWeekdayTextFormat")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetWeekdayTextFormat(ptr.Pointer(), C.int(dayOfWeek), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QCalendarWidget) ShowNextMonth() {
	defer qt.Recovering("QCalendarWidget::showNextMonth")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowNextMonth(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowNextYear() {
	defer qt.Recovering("QCalendarWidget::showNextYear")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowNextYear(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowPreviousMonth() {
	defer qt.Recovering("QCalendarWidget::showPreviousMonth")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowPreviousMonth(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowPreviousYear() {
	defer qt.Recovering("QCalendarWidget::showPreviousYear")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowPreviousYear(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowSelectedDate() {
	defer qt.Recovering("QCalendarWidget::showSelectedDate")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowSelectedDate(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowToday() {
	defer qt.Recovering("QCalendarWidget::showToday")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowToday(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QCalendarWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCalendarWidget_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCalendarWidget) YearShown() int {
	defer qt.Recovering("QCalendarWidget::yearShown")

	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_YearShown(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) DestroyQCalendarWidget() {
	defer qt.Recovering("QCalendarWidget::~QCalendarWidget")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DestroyQCalendarWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
