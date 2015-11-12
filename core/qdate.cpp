#include "qdate.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDate>
#include "_cgo_export.h"

class MyQDate: public QDate {
public:
};

int QDate_QDate_IsLeapYear(int year){
	return QDate::isLeapYear(year);
}

char* QDate_ToString2(void* ptr, int format){
	return static_cast<QDate*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8().data();
}

void* QDate_NewQDate(){
	return new QDate();
}

void* QDate_NewQDate3(int y, int m, int d){
	return new QDate(y, m, d);
}

int QDate_Day(void* ptr){
	return static_cast<QDate*>(ptr)->day();
}

int QDate_DayOfWeek(void* ptr){
	return static_cast<QDate*>(ptr)->dayOfWeek();
}

int QDate_DayOfYear(void* ptr){
	return static_cast<QDate*>(ptr)->dayOfYear();
}

int QDate_DaysInMonth(void* ptr){
	return static_cast<QDate*>(ptr)->daysInMonth();
}

int QDate_DaysInYear(void* ptr){
	return static_cast<QDate*>(ptr)->daysInYear();
}

void QDate_GetDate(void* ptr, int year, int month, int day){
	static_cast<QDate*>(ptr)->getDate(&year, &month, &day);
}

int QDate_IsNull(void* ptr){
	return static_cast<QDate*>(ptr)->isNull();
}

int QDate_QDate_IsValid2(int year, int month, int day){
	return QDate::isValid(year, month, day);
}

int QDate_IsValid(void* ptr){
	return static_cast<QDate*>(ptr)->isValid();
}

char* QDate_QDate_LongDayName(int weekday, int ty){
	return QDate::longDayName(weekday, static_cast<QDate::MonthNameType>(ty)).toUtf8().data();
}

char* QDate_QDate_LongMonthName(int month, int ty){
	return QDate::longMonthName(month, static_cast<QDate::MonthNameType>(ty)).toUtf8().data();
}

int QDate_Month(void* ptr){
	return static_cast<QDate*>(ptr)->month();
}

int QDate_SetDate(void* ptr, int year, int month, int day){
	return static_cast<QDate*>(ptr)->setDate(year, month, day);
}

char* QDate_QDate_ShortDayName(int weekday, int ty){
	return QDate::shortDayName(weekday, static_cast<QDate::MonthNameType>(ty)).toUtf8().data();
}

char* QDate_QDate_ShortMonthName(int month, int ty){
	return QDate::shortMonthName(month, static_cast<QDate::MonthNameType>(ty)).toUtf8().data();
}

char* QDate_ToString(void* ptr, char* format){
	return static_cast<QDate*>(ptr)->toString(QString(format)).toUtf8().data();
}

int QDate_WeekNumber(void* ptr, int yearNumber){
	return static_cast<QDate*>(ptr)->weekNumber(&yearNumber);
}

int QDate_Year(void* ptr){
	return static_cast<QDate*>(ptr)->year();
}

