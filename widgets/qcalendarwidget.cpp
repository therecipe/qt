#include "qcalendarwidget.h"
#include <QString>
#include <QVariant>
#include <QMetaObject>
#include <QTextCharFormat>
#include <QWidget>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QDate>
#include <QCalendarWidget>
#include "_cgo_export.h"

class MyQCalendarWidget: public QCalendarWidget {
public:
void Signal_CurrentPageChanged(int year, int month){callbackQCalendarWidgetCurrentPageChanged(this->objectName().toUtf8().data(), year, month);};
void Signal_SelectionChanged(){callbackQCalendarWidgetSelectionChanged(this->objectName().toUtf8().data());};
};

int QCalendarWidget_DateEditAcceptDelay(QtObjectPtr ptr){
	return static_cast<QCalendarWidget*>(ptr)->dateEditAcceptDelay();
}

int QCalendarWidget_FirstDayOfWeek(QtObjectPtr ptr){
	return static_cast<QCalendarWidget*>(ptr)->firstDayOfWeek();
}

int QCalendarWidget_HorizontalHeaderFormat(QtObjectPtr ptr){
	return static_cast<QCalendarWidget*>(ptr)->horizontalHeaderFormat();
}

int QCalendarWidget_IsDateEditEnabled(QtObjectPtr ptr){
	return static_cast<QCalendarWidget*>(ptr)->isDateEditEnabled();
}

int QCalendarWidget_IsGridVisible(QtObjectPtr ptr){
	return static_cast<QCalendarWidget*>(ptr)->isGridVisible();
}

int QCalendarWidget_IsNavigationBarVisible(QtObjectPtr ptr){
	return static_cast<QCalendarWidget*>(ptr)->isNavigationBarVisible();
}

int QCalendarWidget_SelectionMode(QtObjectPtr ptr){
	return static_cast<QCalendarWidget*>(ptr)->selectionMode();
}

void QCalendarWidget_SetDateEditAcceptDelay(QtObjectPtr ptr, int delay){
	static_cast<QCalendarWidget*>(ptr)->setDateEditAcceptDelay(delay);
}

void QCalendarWidget_SetDateEditEnabled(QtObjectPtr ptr, int enable){
	static_cast<QCalendarWidget*>(ptr)->setDateEditEnabled(enable != 0);
}

void QCalendarWidget_SetFirstDayOfWeek(QtObjectPtr ptr, int dayOfWeek){
	static_cast<QCalendarWidget*>(ptr)->setFirstDayOfWeek(static_cast<Qt::DayOfWeek>(dayOfWeek));
}

void QCalendarWidget_SetGridVisible(QtObjectPtr ptr, int show){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setGridVisible", Q_ARG(bool, show != 0));
}

void QCalendarWidget_SetHorizontalHeaderFormat(QtObjectPtr ptr, int format){
	static_cast<QCalendarWidget*>(ptr)->setHorizontalHeaderFormat(static_cast<QCalendarWidget::HorizontalHeaderFormat>(format));
}

void QCalendarWidget_SetMaximumDate(QtObjectPtr ptr, QtObjectPtr date){
	static_cast<QCalendarWidget*>(ptr)->setMaximumDate(*static_cast<QDate*>(date));
}

void QCalendarWidget_SetMinimumDate(QtObjectPtr ptr, QtObjectPtr date){
	static_cast<QCalendarWidget*>(ptr)->setMinimumDate(*static_cast<QDate*>(date));
}

void QCalendarWidget_SetNavigationBarVisible(QtObjectPtr ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setNavigationBarVisible", Q_ARG(bool, visible != 0));
}

void QCalendarWidget_SetSelectedDate(QtObjectPtr ptr, QtObjectPtr date){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setSelectedDate", Q_ARG(QDate, *static_cast<QDate*>(date)));
}

void QCalendarWidget_SetSelectionMode(QtObjectPtr ptr, int mode){
	static_cast<QCalendarWidget*>(ptr)->setSelectionMode(static_cast<QCalendarWidget::SelectionMode>(mode));
}

void QCalendarWidget_SetVerticalHeaderFormat(QtObjectPtr ptr, int format){
	static_cast<QCalendarWidget*>(ptr)->setVerticalHeaderFormat(static_cast<QCalendarWidget::VerticalHeaderFormat>(format));
}

int QCalendarWidget_VerticalHeaderFormat(QtObjectPtr ptr){
	return static_cast<QCalendarWidget*>(ptr)->verticalHeaderFormat();
}

QtObjectPtr QCalendarWidget_NewQCalendarWidget(QtObjectPtr parent){
	return new QCalendarWidget(static_cast<QWidget*>(parent));
}

void QCalendarWidget_ConnectCurrentPageChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)(int, int)>(&QCalendarWidget::currentPageChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)(int, int)>(&MyQCalendarWidget::Signal_CurrentPageChanged));;
}

void QCalendarWidget_DisconnectCurrentPageChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)(int, int)>(&QCalendarWidget::currentPageChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)(int, int)>(&MyQCalendarWidget::Signal_CurrentPageChanged));;
}

int QCalendarWidget_MonthShown(QtObjectPtr ptr){
	return static_cast<QCalendarWidget*>(ptr)->monthShown();
}

void QCalendarWidget_ConnectSelectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)()>(&QCalendarWidget::selectionChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)()>(&MyQCalendarWidget::Signal_SelectionChanged));;
}

void QCalendarWidget_DisconnectSelectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)()>(&QCalendarWidget::selectionChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)()>(&MyQCalendarWidget::Signal_SelectionChanged));;
}

void QCalendarWidget_SetCurrentPage(QtObjectPtr ptr, int year, int month){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setCurrentPage", Q_ARG(int, year), Q_ARG(int, month));
}

void QCalendarWidget_SetDateRange(QtObjectPtr ptr, QtObjectPtr min, QtObjectPtr max){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setDateRange", Q_ARG(QDate, *static_cast<QDate*>(min)), Q_ARG(QDate, *static_cast<QDate*>(max)));
}

void QCalendarWidget_SetDateTextFormat(QtObjectPtr ptr, QtObjectPtr date, QtObjectPtr format){
	static_cast<QCalendarWidget*>(ptr)->setDateTextFormat(*static_cast<QDate*>(date), *static_cast<QTextCharFormat*>(format));
}

void QCalendarWidget_SetHeaderTextFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QCalendarWidget*>(ptr)->setHeaderTextFormat(*static_cast<QTextCharFormat*>(format));
}

void QCalendarWidget_SetWeekdayTextFormat(QtObjectPtr ptr, int dayOfWeek, QtObjectPtr format){
	static_cast<QCalendarWidget*>(ptr)->setWeekdayTextFormat(static_cast<Qt::DayOfWeek>(dayOfWeek), *static_cast<QTextCharFormat*>(format));
}

void QCalendarWidget_ShowNextMonth(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showNextMonth");
}

void QCalendarWidget_ShowNextYear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showNextYear");
}

void QCalendarWidget_ShowPreviousMonth(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showPreviousMonth");
}

void QCalendarWidget_ShowPreviousYear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showPreviousYear");
}

void QCalendarWidget_ShowSelectedDate(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showSelectedDate");
}

void QCalendarWidget_ShowToday(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showToday");
}

int QCalendarWidget_YearShown(QtObjectPtr ptr){
	return static_cast<QCalendarWidget*>(ptr)->yearShown();
}

void QCalendarWidget_DestroyQCalendarWidget(QtObjectPtr ptr){
	static_cast<QCalendarWidget*>(ptr)->~QCalendarWidget();
}

