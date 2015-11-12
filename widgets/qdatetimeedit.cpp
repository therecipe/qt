#include "qdatetimeedit.h"
#include <QWidget>
#include <QUrl>
#include <QModelIndex>
#include <QDateTime>
#include <QTime>
#include <QCalendarWidget>
#include <QEvent>
#include <QMetaObject>
#include <QDate>
#include <QString>
#include <QVariant>
#include <QObject>
#include <QDateTimeEdit>
#include "_cgo_export.h"

class MyQDateTimeEdit: public QDateTimeEdit {
public:
void Signal_DateTimeChanged(const QDateTime & datetime){callbackQDateTimeEditDateTimeChanged(this->objectName().toUtf8().data(), new QDateTime(datetime));};
};

void* QDateTimeEdit_NewQDateTimeEdit3(void* date, void* parent){
	return new QDateTimeEdit(*static_cast<QDate*>(date), static_cast<QWidget*>(parent));
}

void* QDateTimeEdit_NewQDateTimeEdit4(void* time, void* parent){
	return new QDateTimeEdit(*static_cast<QTime*>(time), static_cast<QWidget*>(parent));
}

int QDateTimeEdit_CalendarPopup(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->calendarPopup();
}

void QDateTimeEdit_ClearMaximumDate(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMaximumDate();
}

void QDateTimeEdit_ClearMaximumDateTime(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMaximumDateTime();
}

void QDateTimeEdit_ClearMaximumTime(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMaximumTime();
}

void QDateTimeEdit_ClearMinimumDate(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMinimumDate();
}

void QDateTimeEdit_ClearMinimumDateTime(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMinimumDateTime();
}

void QDateTimeEdit_ClearMinimumTime(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMinimumTime();
}

int QDateTimeEdit_CurrentSection(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->currentSection();
}

int QDateTimeEdit_CurrentSectionIndex(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->currentSectionIndex();
}

void* QDateTimeEdit_DateTime(void* ptr){
	return new QDateTime(static_cast<QDateTimeEdit*>(ptr)->dateTime());
}

char* QDateTimeEdit_DisplayFormat(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->displayFormat().toUtf8().data();
}

int QDateTimeEdit_DisplayedSections(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->displayedSections();
}

void* QDateTimeEdit_MaximumDateTime(void* ptr){
	return new QDateTime(static_cast<QDateTimeEdit*>(ptr)->maximumDateTime());
}

void* QDateTimeEdit_MinimumDateTime(void* ptr){
	return new QDateTime(static_cast<QDateTimeEdit*>(ptr)->minimumDateTime());
}

int QDateTimeEdit_SectionCount(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->sectionCount();
}

char* QDateTimeEdit_SectionText(void* ptr, int section){
	return static_cast<QDateTimeEdit*>(ptr)->sectionText(static_cast<QDateTimeEdit::Section>(section)).toUtf8().data();
}

void QDateTimeEdit_SetCalendarPopup(void* ptr, int enable){
	static_cast<QDateTimeEdit*>(ptr)->setCalendarPopup(enable != 0);
}

void QDateTimeEdit_SetCurrentSection(void* ptr, int section){
	static_cast<QDateTimeEdit*>(ptr)->setCurrentSection(static_cast<QDateTimeEdit::Section>(section));
}

void QDateTimeEdit_SetCurrentSectionIndex(void* ptr, int index){
	static_cast<QDateTimeEdit*>(ptr)->setCurrentSectionIndex(index);
}

void QDateTimeEdit_SetDate(void* ptr, void* date){
	QMetaObject::invokeMethod(static_cast<QDateTimeEdit*>(ptr), "setDate", Q_ARG(QDate, *static_cast<QDate*>(date)));
}

void QDateTimeEdit_SetDateTime(void* ptr, void* dateTime){
	QMetaObject::invokeMethod(static_cast<QDateTimeEdit*>(ptr), "setDateTime", Q_ARG(QDateTime, *static_cast<QDateTime*>(dateTime)));
}

void QDateTimeEdit_SetDisplayFormat(void* ptr, char* format){
	static_cast<QDateTimeEdit*>(ptr)->setDisplayFormat(QString(format));
}

void QDateTimeEdit_SetMaximumDate(void* ptr, void* max){
	static_cast<QDateTimeEdit*>(ptr)->setMaximumDate(*static_cast<QDate*>(max));
}

void QDateTimeEdit_SetMaximumDateTime(void* ptr, void* dt){
	static_cast<QDateTimeEdit*>(ptr)->setMaximumDateTime(*static_cast<QDateTime*>(dt));
}

void QDateTimeEdit_SetMaximumTime(void* ptr, void* max){
	static_cast<QDateTimeEdit*>(ptr)->setMaximumTime(*static_cast<QTime*>(max));
}

void QDateTimeEdit_SetMinimumDate(void* ptr, void* min){
	static_cast<QDateTimeEdit*>(ptr)->setMinimumDate(*static_cast<QDate*>(min));
}

void QDateTimeEdit_SetMinimumDateTime(void* ptr, void* dt){
	static_cast<QDateTimeEdit*>(ptr)->setMinimumDateTime(*static_cast<QDateTime*>(dt));
}

void QDateTimeEdit_SetMinimumTime(void* ptr, void* min){
	static_cast<QDateTimeEdit*>(ptr)->setMinimumTime(*static_cast<QTime*>(min));
}

void QDateTimeEdit_SetTime(void* ptr, void* time){
	QMetaObject::invokeMethod(static_cast<QDateTimeEdit*>(ptr), "setTime", Q_ARG(QTime, *static_cast<QTime*>(time)));
}

void QDateTimeEdit_SetTimeSpec(void* ptr, int spec){
	static_cast<QDateTimeEdit*>(ptr)->setTimeSpec(static_cast<Qt::TimeSpec>(spec));
}

int QDateTimeEdit_TimeSpec(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->timeSpec();
}

void* QDateTimeEdit_NewQDateTimeEdit(void* parent){
	return new QDateTimeEdit(static_cast<QWidget*>(parent));
}

void* QDateTimeEdit_NewQDateTimeEdit2(void* datetime, void* parent){
	return new QDateTimeEdit(*static_cast<QDateTime*>(datetime), static_cast<QWidget*>(parent));
}

void* QDateTimeEdit_CalendarWidget(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->calendarWidget();
}

void QDateTimeEdit_Clear(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clear();
}

void QDateTimeEdit_ConnectDateTimeChanged(void* ptr){
	QObject::connect(static_cast<QDateTimeEdit*>(ptr), static_cast<void (QDateTimeEdit::*)(const QDateTime &)>(&QDateTimeEdit::dateTimeChanged), static_cast<MyQDateTimeEdit*>(ptr), static_cast<void (MyQDateTimeEdit::*)(const QDateTime &)>(&MyQDateTimeEdit::Signal_DateTimeChanged));;
}

void QDateTimeEdit_DisconnectDateTimeChanged(void* ptr){
	QObject::disconnect(static_cast<QDateTimeEdit*>(ptr), static_cast<void (QDateTimeEdit::*)(const QDateTime &)>(&QDateTimeEdit::dateTimeChanged), static_cast<MyQDateTimeEdit*>(ptr), static_cast<void (MyQDateTimeEdit::*)(const QDateTime &)>(&MyQDateTimeEdit::Signal_DateTimeChanged));;
}

int QDateTimeEdit_Event(void* ptr, void* event){
	return static_cast<QDateTimeEdit*>(ptr)->event(static_cast<QEvent*>(event));
}

int QDateTimeEdit_SectionAt(void* ptr, int index){
	return static_cast<QDateTimeEdit*>(ptr)->sectionAt(index);
}

void QDateTimeEdit_SetCalendarWidget(void* ptr, void* calendarWidget){
	static_cast<QDateTimeEdit*>(ptr)->setCalendarWidget(static_cast<QCalendarWidget*>(calendarWidget));
}

void QDateTimeEdit_SetDateRange(void* ptr, void* min, void* max){
	static_cast<QDateTimeEdit*>(ptr)->setDateRange(*static_cast<QDate*>(min), *static_cast<QDate*>(max));
}

void QDateTimeEdit_SetDateTimeRange(void* ptr, void* min, void* max){
	static_cast<QDateTimeEdit*>(ptr)->setDateTimeRange(*static_cast<QDateTime*>(min), *static_cast<QDateTime*>(max));
}

void QDateTimeEdit_SetSelectedSection(void* ptr, int section){
	static_cast<QDateTimeEdit*>(ptr)->setSelectedSection(static_cast<QDateTimeEdit::Section>(section));
}

void QDateTimeEdit_SetTimeRange(void* ptr, void* min, void* max){
	static_cast<QDateTimeEdit*>(ptr)->setTimeRange(*static_cast<QTime*>(min), *static_cast<QTime*>(max));
}

void QDateTimeEdit_StepBy(void* ptr, int steps){
	static_cast<QDateTimeEdit*>(ptr)->stepBy(steps);
}

void QDateTimeEdit_DestroyQDateTimeEdit(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->~QDateTimeEdit();
}

