#include "qcalendarwidget.h"
#include <QUrl>
#include <QObject>
#include <QDate>
#include <QWidget>
#include <QTextCharFormat>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QMetaObject>
#include <QCalendarWidget>
#include "_cgo_export.h"

class MyQCalendarWidget: public QCalendarWidget {
public:
void Signal_CurrentPageChanged(int year, int month){callbackQCalendarWidgetCurrentPageChanged(this->objectName().toUtf8().data(), year, month);};
void Signal_SelectionChanged(){callbackQCalendarWidgetSelectionChanged(this->objectName().toUtf8().data());};
};

int QCalendarWidget_DateEditAcceptDelay(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->dateEditAcceptDelay();
}

int QCalendarWidget_FirstDayOfWeek(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->firstDayOfWeek();
}

int QCalendarWidget_HorizontalHeaderFormat(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->horizontalHeaderFormat();
}

int QCalendarWidget_IsDateEditEnabled(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->isDateEditEnabled();
}

int QCalendarWidget_IsGridVisible(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->isGridVisible();
}

int QCalendarWidget_IsNavigationBarVisible(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->isNavigationBarVisible();
}

int QCalendarWidget_SelectionMode(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->selectionMode();
}

void QCalendarWidget_SetDateEditAcceptDelay(void* ptr, int delay){
	static_cast<QCalendarWidget*>(ptr)->setDateEditAcceptDelay(delay);
}

void QCalendarWidget_SetDateEditEnabled(void* ptr, int enable){
	static_cast<QCalendarWidget*>(ptr)->setDateEditEnabled(enable != 0);
}

void QCalendarWidget_SetFirstDayOfWeek(void* ptr, int dayOfWeek){
	static_cast<QCalendarWidget*>(ptr)->setFirstDayOfWeek(static_cast<Qt::DayOfWeek>(dayOfWeek));
}

void QCalendarWidget_SetGridVisible(void* ptr, int show){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setGridVisible", Q_ARG(bool, show != 0));
}

void QCalendarWidget_SetHorizontalHeaderFormat(void* ptr, int format){
	static_cast<QCalendarWidget*>(ptr)->setHorizontalHeaderFormat(static_cast<QCalendarWidget::HorizontalHeaderFormat>(format));
}

void QCalendarWidget_SetMaximumDate(void* ptr, void* date){
	static_cast<QCalendarWidget*>(ptr)->setMaximumDate(*static_cast<QDate*>(date));
}

void QCalendarWidget_SetMinimumDate(void* ptr, void* date){
	static_cast<QCalendarWidget*>(ptr)->setMinimumDate(*static_cast<QDate*>(date));
}

void QCalendarWidget_SetNavigationBarVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setNavigationBarVisible", Q_ARG(bool, visible != 0));
}

void QCalendarWidget_SetSelectedDate(void* ptr, void* date){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setSelectedDate", Q_ARG(QDate, *static_cast<QDate*>(date)));
}

void QCalendarWidget_SetSelectionMode(void* ptr, int mode){
	static_cast<QCalendarWidget*>(ptr)->setSelectionMode(static_cast<QCalendarWidget::SelectionMode>(mode));
}

void QCalendarWidget_SetVerticalHeaderFormat(void* ptr, int format){
	static_cast<QCalendarWidget*>(ptr)->setVerticalHeaderFormat(static_cast<QCalendarWidget::VerticalHeaderFormat>(format));
}

int QCalendarWidget_VerticalHeaderFormat(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->verticalHeaderFormat();
}

void* QCalendarWidget_NewQCalendarWidget(void* parent){
	return new QCalendarWidget(static_cast<QWidget*>(parent));
}

void QCalendarWidget_ConnectCurrentPageChanged(void* ptr){
	QObject::connect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)(int, int)>(&QCalendarWidget::currentPageChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)(int, int)>(&MyQCalendarWidget::Signal_CurrentPageChanged));;
}

void QCalendarWidget_DisconnectCurrentPageChanged(void* ptr){
	QObject::disconnect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)(int, int)>(&QCalendarWidget::currentPageChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)(int, int)>(&MyQCalendarWidget::Signal_CurrentPageChanged));;
}

int QCalendarWidget_MonthShown(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->monthShown();
}

void QCalendarWidget_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)()>(&QCalendarWidget::selectionChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)()>(&MyQCalendarWidget::Signal_SelectionChanged));;
}

void QCalendarWidget_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)()>(&QCalendarWidget::selectionChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)()>(&MyQCalendarWidget::Signal_SelectionChanged));;
}

void QCalendarWidget_SetCurrentPage(void* ptr, int year, int month){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setCurrentPage", Q_ARG(int, year), Q_ARG(int, month));
}

void QCalendarWidget_SetDateRange(void* ptr, void* min, void* max){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setDateRange", Q_ARG(QDate, *static_cast<QDate*>(min)), Q_ARG(QDate, *static_cast<QDate*>(max)));
}

void QCalendarWidget_SetDateTextFormat(void* ptr, void* date, void* format){
	static_cast<QCalendarWidget*>(ptr)->setDateTextFormat(*static_cast<QDate*>(date), *static_cast<QTextCharFormat*>(format));
}

void QCalendarWidget_SetHeaderTextFormat(void* ptr, void* format){
	static_cast<QCalendarWidget*>(ptr)->setHeaderTextFormat(*static_cast<QTextCharFormat*>(format));
}

void QCalendarWidget_SetWeekdayTextFormat(void* ptr, int dayOfWeek, void* format){
	static_cast<QCalendarWidget*>(ptr)->setWeekdayTextFormat(static_cast<Qt::DayOfWeek>(dayOfWeek), *static_cast<QTextCharFormat*>(format));
}

void QCalendarWidget_ShowNextMonth(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showNextMonth");
}

void QCalendarWidget_ShowNextYear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showNextYear");
}

void QCalendarWidget_ShowPreviousMonth(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showPreviousMonth");
}

void QCalendarWidget_ShowPreviousYear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showPreviousYear");
}

void QCalendarWidget_ShowSelectedDate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showSelectedDate");
}

void QCalendarWidget_ShowToday(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showToday");
}

int QCalendarWidget_YearShown(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->yearShown();
}

void QCalendarWidget_DestroyQCalendarWidget(void* ptr){
	static_cast<QCalendarWidget*>(ptr)->~QCalendarWidget();
}

