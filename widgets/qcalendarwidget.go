package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::dateEditAcceptDelay")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_DateEditAcceptDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) FirstDayOfWeek() core.Qt__DayOfWeek {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::firstDayOfWeek")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__DayOfWeek(C.QCalendarWidget_FirstDayOfWeek(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) HorizontalHeaderFormat() QCalendarWidget__HorizontalHeaderFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::horizontalHeaderFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return QCalendarWidget__HorizontalHeaderFormat(C.QCalendarWidget_HorizontalHeaderFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) IsDateEditEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::isDateEditEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsDateEditEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCalendarWidget) IsGridVisible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::isGridVisible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsGridVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCalendarWidget) IsNavigationBarVisible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::isNavigationBarVisible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCalendarWidget_IsNavigationBarVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCalendarWidget) SelectionMode() QCalendarWidget__SelectionMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::selectionMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QCalendarWidget__SelectionMode(C.QCalendarWidget_SelectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) SetDateEditAcceptDelay(delay int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setDateEditAcceptDelay")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateEditAcceptDelay(ptr.Pointer(), C.int(delay))
	}
}

func (ptr *QCalendarWidget) SetDateEditEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setDateEditEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateEditEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QCalendarWidget) SetFirstDayOfWeek(dayOfWeek core.Qt__DayOfWeek) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setFirstDayOfWeek")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetFirstDayOfWeek(ptr.Pointer(), C.int(dayOfWeek))
	}
}

func (ptr *QCalendarWidget) SetGridVisible(show bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setGridVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetGridVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QCalendarWidget) SetHorizontalHeaderFormat(format QCalendarWidget__HorizontalHeaderFormat) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setHorizontalHeaderFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetHorizontalHeaderFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCalendarWidget) SetMaximumDate(date core.QDate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setMaximumDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetMaximumDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QCalendarWidget) SetMinimumDate(date core.QDate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setMinimumDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetMinimumDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QCalendarWidget) SetNavigationBarVisible(visible bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setNavigationBarVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetNavigationBarVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCalendarWidget) SetSelectedDate(date core.QDate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setSelectedDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetSelectedDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QCalendarWidget) SetSelectionMode(mode QCalendarWidget__SelectionMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setSelectionMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetSelectionMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCalendarWidget) SetVerticalHeaderFormat(format QCalendarWidget__VerticalHeaderFormat) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setVerticalHeaderFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetVerticalHeaderFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCalendarWidget) VerticalHeaderFormat() QCalendarWidget__VerticalHeaderFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::verticalHeaderFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return QCalendarWidget__VerticalHeaderFormat(C.QCalendarWidget_VerticalHeaderFormat(ptr.Pointer()))
	}
	return 0
}

func NewQCalendarWidget(parent QWidget_ITF) *QCalendarWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::QCalendarWidget")
		}
	}()

	return NewQCalendarWidgetFromPointer(C.QCalendarWidget_NewQCalendarWidget(PointerFromQWidget(parent)))
}

func (ptr *QCalendarWidget) ConnectCurrentPageChanged(f func(year int, month int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::currentPageChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ConnectCurrentPageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentPageChanged", f)
	}
}

func (ptr *QCalendarWidget) DisconnectCurrentPageChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::currentPageChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DisconnectCurrentPageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentPageChanged")
	}
}

//export callbackQCalendarWidgetCurrentPageChanged
func callbackQCalendarWidgetCurrentPageChanged(ptrName *C.char, year C.int, month C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::currentPageChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentPageChanged").(func(int, int))(int(year), int(month))
}

func (ptr *QCalendarWidget) MonthShown() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::monthShown")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_MonthShown(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) ConnectSelectionChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::selectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QCalendarWidget) DisconnectSelectionChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::selectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQCalendarWidgetSelectionChanged
func callbackQCalendarWidgetSelectionChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::selectionChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QCalendarWidget) SetCurrentPage(year int, month int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setCurrentPage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetCurrentPage(ptr.Pointer(), C.int(year), C.int(month))
	}
}

func (ptr *QCalendarWidget) SetDateRange(min core.QDate_ITF, max core.QDate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setDateRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateRange(ptr.Pointer(), core.PointerFromQDate(min), core.PointerFromQDate(max))
	}
}

func (ptr *QCalendarWidget) SetDateTextFormat(date core.QDate_ITF, format gui.QTextCharFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setDateTextFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetDateTextFormat(ptr.Pointer(), core.PointerFromQDate(date), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QCalendarWidget) SetHeaderTextFormat(format gui.QTextCharFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setHeaderTextFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetHeaderTextFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QCalendarWidget) SetWeekdayTextFormat(dayOfWeek core.Qt__DayOfWeek, format gui.QTextCharFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::setWeekdayTextFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_SetWeekdayTextFormat(ptr.Pointer(), C.int(dayOfWeek), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QCalendarWidget) ShowNextMonth() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::showNextMonth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowNextMonth(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowNextYear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::showNextYear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowNextYear(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowPreviousMonth() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::showPreviousMonth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowPreviousMonth(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowPreviousYear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::showPreviousYear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowPreviousYear(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowSelectedDate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::showSelectedDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowSelectedDate(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) ShowToday() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::showToday")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_ShowToday(ptr.Pointer())
	}
}

func (ptr *QCalendarWidget) YearShown() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::yearShown")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCalendarWidget_YearShown(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCalendarWidget) DestroyQCalendarWidget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCalendarWidget::~QCalendarWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCalendarWidget_DestroyQCalendarWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
