#include "qdatetime.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTime>
#include <QTimeZone>
#include <QDate>
#include <QDateTime>
#include "_cgo_export.h"

class MyQDateTime: public QDateTime {
public:
};

char* QDateTime_ToString2(QtObjectPtr ptr, int format){
	return static_cast<QDateTime*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8().data();
}

QtObjectPtr QDateTime_NewQDateTime(){
	return new QDateTime();
}

QtObjectPtr QDateTime_NewQDateTime2(QtObjectPtr date){
	return new QDateTime(*static_cast<QDate*>(date));
}

QtObjectPtr QDateTime_NewQDateTime3(QtObjectPtr date, QtObjectPtr time, int spec){
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(time), static_cast<Qt::TimeSpec>(spec));
}

QtObjectPtr QDateTime_NewQDateTime4(QtObjectPtr date, QtObjectPtr time, int spec, int offsetSeconds){
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(time), static_cast<Qt::TimeSpec>(spec), offsetSeconds);
}

QtObjectPtr QDateTime_NewQDateTime5(QtObjectPtr date, QtObjectPtr time, QtObjectPtr timeZone){
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(time), *static_cast<QTimeZone*>(timeZone));
}

QtObjectPtr QDateTime_NewQDateTime6(QtObjectPtr other){
	return new QDateTime(*static_cast<QDateTime*>(other));
}

int QDateTime_IsDaylightTime(QtObjectPtr ptr){
	return static_cast<QDateTime*>(ptr)->isDaylightTime();
}

int QDateTime_IsNull(QtObjectPtr ptr){
	return static_cast<QDateTime*>(ptr)->isNull();
}

int QDateTime_IsValid(QtObjectPtr ptr){
	return static_cast<QDateTime*>(ptr)->isValid();
}

int QDateTime_OffsetFromUtc(QtObjectPtr ptr){
	return static_cast<QDateTime*>(ptr)->offsetFromUtc();
}

void QDateTime_SetDate(QtObjectPtr ptr, QtObjectPtr date){
	static_cast<QDateTime*>(ptr)->setDate(*static_cast<QDate*>(date));
}

void QDateTime_SetOffsetFromUtc(QtObjectPtr ptr, int offsetSeconds){
	static_cast<QDateTime*>(ptr)->setOffsetFromUtc(offsetSeconds);
}

void QDateTime_SetTime(QtObjectPtr ptr, QtObjectPtr time){
	static_cast<QDateTime*>(ptr)->setTime(*static_cast<QTime*>(time));
}

void QDateTime_SetTimeSpec(QtObjectPtr ptr, int spec){
	static_cast<QDateTime*>(ptr)->setTimeSpec(static_cast<Qt::TimeSpec>(spec));
}

void QDateTime_SetTimeZone(QtObjectPtr ptr, QtObjectPtr toZone){
	static_cast<QDateTime*>(ptr)->setTimeZone(*static_cast<QTimeZone*>(toZone));
}

void QDateTime_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QDateTime*>(ptr)->swap(*static_cast<QDateTime*>(other));
}

int QDateTime_TimeSpec(QtObjectPtr ptr){
	return static_cast<QDateTime*>(ptr)->timeSpec();
}

char* QDateTime_TimeZoneAbbreviation(QtObjectPtr ptr){
	return static_cast<QDateTime*>(ptr)->timeZoneAbbreviation().toUtf8().data();
}

char* QDateTime_ToString(QtObjectPtr ptr, char* format){
	return static_cast<QDateTime*>(ptr)->toString(QString(format)).toUtf8().data();
}

void QDateTime_DestroyQDateTime(QtObjectPtr ptr){
	static_cast<QDateTime*>(ptr)->~QDateTime();
}

