#include "qdatetimeedit.h"
#include <QString>
#include <QUrl>
#include <QWidget>
#include <QDate>
#include <QEvent>
#include <QMetaObject>
#include <QDateTime>
#include <QCalendarWidget>
#include <QVariant>
#include <QModelIndex>
#include <QTime>
#include <QDateTimeEdit>
#include "_cgo_export.h"

class MyQDateTimeEdit: public QDateTimeEdit {
public:
};

QtObjectPtr QDateTimeEdit_NewQDateTimeEdit3(QtObjectPtr date, QtObjectPtr parent){
	return new QDateTimeEdit(*static_cast<QDate*>(date), static_cast<QWidget*>(parent));
}

QtObjectPtr QDateTimeEdit_NewQDateTimeEdit4(QtObjectPtr time, QtObjectPtr parent){
	return new QDateTimeEdit(*static_cast<QTime*>(time), static_cast<QWidget*>(parent));
}

int QDateTimeEdit_CalendarPopup(QtObjectPtr ptr){
	return static_cast<QDateTimeEdit*>(ptr)->calendarPopup();
}

void QDateTimeEdit_ClearMaximumDate(QtObjectPtr ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMaximumDate();
}

void QDateTimeEdit_ClearMaximumDateTime(QtObjectPtr ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMaximumDateTime();
}

void QDateTimeEdit_ClearMaximumTime(QtObjectPtr ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMaximumTime();
}

void QDateTimeEdit_ClearMinimumDate(QtObjectPtr ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMinimumDate();
}

void QDateTimeEdit_ClearMinimumDateTime(QtObjectPtr ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMinimumDateTime();
}

void QDateTimeEdit_ClearMinimumTime(QtObjectPtr ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMinimumTime();
}

int QDateTimeEdit_CurrentSection(QtObjectPtr ptr){
	return static_cast<QDateTimeEdit*>(ptr)->currentSection();
}

int QDateTimeEdit_CurrentSectionIndex(QtObjectPtr ptr){
	return static_cast<QDateTimeEdit*>(ptr)->currentSectionIndex();
}

char* QDateTimeEdit_DisplayFormat(QtObjectPtr ptr){
	return static_cast<QDateTimeEdit*>(ptr)->displayFormat().toUtf8().data();
}

int QDateTimeEdit_DisplayedSections(QtObjectPtr ptr){
	return static_cast<QDateTimeEdit*>(ptr)->displayedSections();
}

int QDateTimeEdit_SectionCount(QtObjectPtr ptr){
	return static_cast<QDateTimeEdit*>(ptr)->sectionCount();
}

char* QDateTimeEdit_SectionText(QtObjectPtr ptr, int section){
	return static_cast<QDateTimeEdit*>(ptr)->sectionText(static_cast<QDateTimeEdit::Section>(section)).toUtf8().data();
}

void QDateTimeEdit_SetCalendarPopup(QtObjectPtr ptr, int enable){
	static_cast<QDateTimeEdit*>(ptr)->setCalendarPopup(enable != 0);
}

void QDateTimeEdit_SetCurrentSection(QtObjectPtr ptr, int section){
	static_cast<QDateTimeEdit*>(ptr)->setCurrentSection(static_cast<QDateTimeEdit::Section>(section));
}

void QDateTimeEdit_SetCurrentSectionIndex(QtObjectPtr ptr, int index){
	static_cast<QDateTimeEdit*>(ptr)->setCurrentSectionIndex(index);
}

void QDateTimeEdit_SetDate(QtObjectPtr ptr, QtObjectPtr date){
	QMetaObject::invokeMethod(static_cast<QDateTimeEdit*>(ptr), "setDate", Q_ARG(QDate, *static_cast<QDate*>(date)));
}

void QDateTimeEdit_SetDateTime(QtObjectPtr ptr, QtObjectPtr dateTime){
	QMetaObject::invokeMethod(static_cast<QDateTimeEdit*>(ptr), "setDateTime", Q_ARG(QDateTime, *static_cast<QDateTime*>(dateTime)));
}

void QDateTimeEdit_SetDisplayFormat(QtObjectPtr ptr, char* format){
	static_cast<QDateTimeEdit*>(ptr)->setDisplayFormat(QString(format));
}

void QDateTimeEdit_SetMaximumDate(QtObjectPtr ptr, QtObjectPtr max){
	static_cast<QDateTimeEdit*>(ptr)->setMaximumDate(*static_cast<QDate*>(max));
}

void QDateTimeEdit_SetMaximumDateTime(QtObjectPtr ptr, QtObjectPtr dt){
	static_cast<QDateTimeEdit*>(ptr)->setMaximumDateTime(*static_cast<QDateTime*>(dt));
}

void QDateTimeEdit_SetMaximumTime(QtObjectPtr ptr, QtObjectPtr max){
	static_cast<QDateTimeEdit*>(ptr)->setMaximumTime(*static_cast<QTime*>(max));
}

void QDateTimeEdit_SetMinimumDate(QtObjectPtr ptr, QtObjectPtr min){
	static_cast<QDateTimeEdit*>(ptr)->setMinimumDate(*static_cast<QDate*>(min));
}

void QDateTimeEdit_SetMinimumDateTime(QtObjectPtr ptr, QtObjectPtr dt){
	static_cast<QDateTimeEdit*>(ptr)->setMinimumDateTime(*static_cast<QDateTime*>(dt));
}

void QDateTimeEdit_SetMinimumTime(QtObjectPtr ptr, QtObjectPtr min){
	static_cast<QDateTimeEdit*>(ptr)->setMinimumTime(*static_cast<QTime*>(min));
}

void QDateTimeEdit_SetTime(QtObjectPtr ptr, QtObjectPtr time){
	QMetaObject::invokeMethod(static_cast<QDateTimeEdit*>(ptr), "setTime", Q_ARG(QTime, *static_cast<QTime*>(time)));
}

void QDateTimeEdit_SetTimeSpec(QtObjectPtr ptr, int spec){
	static_cast<QDateTimeEdit*>(ptr)->setTimeSpec(static_cast<Qt::TimeSpec>(spec));
}

int QDateTimeEdit_TimeSpec(QtObjectPtr ptr){
	return static_cast<QDateTimeEdit*>(ptr)->timeSpec();
}

QtObjectPtr QDateTimeEdit_NewQDateTimeEdit(QtObjectPtr parent){
	return new QDateTimeEdit(static_cast<QWidget*>(parent));
}

QtObjectPtr QDateTimeEdit_NewQDateTimeEdit2(QtObjectPtr datetime, QtObjectPtr parent){
	return new QDateTimeEdit(*static_cast<QDateTime*>(datetime), static_cast<QWidget*>(parent));
}

QtObjectPtr QDateTimeEdit_CalendarWidget(QtObjectPtr ptr){
	return static_cast<QDateTimeEdit*>(ptr)->calendarWidget();
}

void QDateTimeEdit_Clear(QtObjectPtr ptr){
	static_cast<QDateTimeEdit*>(ptr)->clear();
}

int QDateTimeEdit_Event(QtObjectPtr ptr, QtObjectPtr event){
	return static_cast<QDateTimeEdit*>(ptr)->event(static_cast<QEvent*>(event));
}

int QDateTimeEdit_SectionAt(QtObjectPtr ptr, int index){
	return static_cast<QDateTimeEdit*>(ptr)->sectionAt(index);
}

void QDateTimeEdit_SetCalendarWidget(QtObjectPtr ptr, QtObjectPtr calendarWidget){
	static_cast<QDateTimeEdit*>(ptr)->setCalendarWidget(static_cast<QCalendarWidget*>(calendarWidget));
}

void QDateTimeEdit_SetDateRange(QtObjectPtr ptr, QtObjectPtr min, QtObjectPtr max){
	static_cast<QDateTimeEdit*>(ptr)->setDateRange(*static_cast<QDate*>(min), *static_cast<QDate*>(max));
}

void QDateTimeEdit_SetDateTimeRange(QtObjectPtr ptr, QtObjectPtr min, QtObjectPtr max){
	static_cast<QDateTimeEdit*>(ptr)->setDateTimeRange(*static_cast<QDateTime*>(min), *static_cast<QDateTime*>(max));
}

void QDateTimeEdit_SetSelectedSection(QtObjectPtr ptr, int section){
	static_cast<QDateTimeEdit*>(ptr)->setSelectedSection(static_cast<QDateTimeEdit::Section>(section));
}

void QDateTimeEdit_SetTimeRange(QtObjectPtr ptr, QtObjectPtr min, QtObjectPtr max){
	static_cast<QDateTimeEdit*>(ptr)->setTimeRange(*static_cast<QTime*>(min), *static_cast<QTime*>(max));
}

void QDateTimeEdit_StepBy(QtObjectPtr ptr, int steps){
	static_cast<QDateTimeEdit*>(ptr)->stepBy(steps);
}

void QDateTimeEdit_DestroyQDateTimeEdit(QtObjectPtr ptr){
	static_cast<QDateTimeEdit*>(ptr)->~QDateTimeEdit();
}

