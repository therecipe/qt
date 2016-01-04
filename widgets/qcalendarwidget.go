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
func callbackQCalendarWidgetCurrentPageChanged(ptr unsafe.Pointer, ptrName *C.char, year C.int, month C.int) {
	defer qt.Recovering("callback QCalendarWidget::currentPageChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentPageChanged"); signal != nil {
		signal.(func(int, int))(int(year), int(month))
	}

}

func (ptr *QCalendarWidget) CurrentPageChanged(year int, month int) {
	defer qt.Recovering("QCalendarWidget::currentPageChanged")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_CurrentPageChanged(ptr.Pointer(), C.int(year), C.int(month))
	}
}

func (ptr *QCalendarWidget) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QCalendarWidget::event")

	if ptr.Pointer() != nil {
		return C.QCalendarWidget_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QCalendarWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QCalendarWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QCalendarWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
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
func callbackQCalendarWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QCalendarWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQCalendarWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCalendarWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQCalendarWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QCalendarWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
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
func callbackQCalendarWidgetSelectionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCalendarWidget::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCalendarWidget) SelectionChanged() {
	defer qt.Recovering("QCalendarWidget::selectionChanged")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SelectionChanged(ptr.Pointer())
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

func (ptr *QCalendarWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QCalendarWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQCalendarWidgetActionEvent
func callbackQCalendarWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QCalendarWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QCalendarWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQCalendarWidgetDragEnterEvent
func callbackQCalendarWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QCalendarWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QCalendarWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQCalendarWidgetDragLeaveEvent
func callbackQCalendarWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QCalendarWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QCalendarWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQCalendarWidgetDragMoveEvent
func callbackQCalendarWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QCalendarWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QCalendarWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQCalendarWidgetDropEvent
func callbackQCalendarWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QCalendarWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCalendarWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQCalendarWidgetEnterEvent
func callbackQCalendarWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCalendarWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCalendarWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQCalendarWidgetFocusInEvent
func callbackQCalendarWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QCalendarWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCalendarWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQCalendarWidgetFocusOutEvent
func callbackQCalendarWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QCalendarWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QCalendarWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQCalendarWidgetHideEvent
func callbackQCalendarWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QCalendarWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCalendarWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQCalendarWidgetLeaveEvent
func callbackQCalendarWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCalendarWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QCalendarWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQCalendarWidgetMoveEvent
func callbackQCalendarWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QCalendarWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QCalendarWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQCalendarWidgetPaintEvent
func callbackQCalendarWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QCalendarWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QCalendarWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QCalendarWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QCalendarWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQCalendarWidgetSetVisible
func callbackQCalendarWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QCalendarWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QCalendarWidget) SetVisible(visible bool) {
	defer qt.Recovering("QCalendarWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCalendarWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QCalendarWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCalendarWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QCalendarWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQCalendarWidgetShowEvent
func callbackQCalendarWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QCalendarWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCalendarWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQCalendarWidgetChangeEvent
func callbackQCalendarWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCalendarWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QCalendarWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQCalendarWidgetCloseEvent
func callbackQCalendarWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QCalendarWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QCalendarWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQCalendarWidgetContextMenuEvent
func callbackQCalendarWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QCalendarWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QCalendarWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QCalendarWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QCalendarWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQCalendarWidgetInitPainter
func callbackQCalendarWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQCalendarWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QCalendarWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QCalendarWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QCalendarWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QCalendarWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QCalendarWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QCalendarWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQCalendarWidgetInputMethodEvent
func callbackQCalendarWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QCalendarWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCalendarWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQCalendarWidgetKeyReleaseEvent
func callbackQCalendarWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QCalendarWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCalendarWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQCalendarWidgetMouseDoubleClickEvent
func callbackQCalendarWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCalendarWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCalendarWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQCalendarWidgetMouseMoveEvent
func callbackQCalendarWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCalendarWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCalendarWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQCalendarWidgetMouseReleaseEvent
func callbackQCalendarWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCalendarWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QCalendarWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQCalendarWidgetTabletEvent
func callbackQCalendarWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QCalendarWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QCalendarWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQCalendarWidgetWheelEvent
func callbackQCalendarWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QCalendarWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCalendarWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCalendarWidgetTimerEvent
func callbackQCalendarWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCalendarWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCalendarWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCalendarWidgetChildEvent
func callbackQCalendarWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCalendarWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCalendarWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCalendarWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCalendarWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCalendarWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCalendarWidgetCustomEvent
func callbackQCalendarWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCalendarWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCalendarWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCalendarWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCalendarWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCalendarWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QCalendarWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
