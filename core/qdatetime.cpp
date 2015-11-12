#include "qdatetime.h"
#include <QTimeZone>
#include <QDate>
#include <QTime>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDateTime>
#include "_cgo_export.h"

class MyQDateTime: public QDateTime {
public:
};

void* QDateTime_QDateTime_CurrentDateTime(){
	return new QDateTime(QDateTime::currentDateTime());
}

void* QDateTime_QDateTime_CurrentDateTimeUtc(){
	return new QDateTime(QDateTime::currentDateTimeUtc());
}

void* QDateTime_QDateTime_FromString(char* stri, int format){
	return new QDateTime(QDateTime::fromString(QString(stri), static_cast<Qt::DateFormat>(format)));
}

void* QDateTime_QDateTime_FromString2(char* stri, char* format){
	return new QDateTime(QDateTime::fromString(QString(stri), QString(format)));
}

void* QDateTime_ToOffsetFromUtc(void* ptr, int offsetSeconds){
	return new QDateTime(static_cast<QDateTime*>(ptr)->toOffsetFromUtc(offsetSeconds));
}

char* QDateTime_ToString2(void* ptr, int format){
	return static_cast<QDateTime*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8().data();
}

void* QDateTime_ToTimeSpec(void* ptr, int spec){
	return new QDateTime(static_cast<QDateTime*>(ptr)->toTimeSpec(static_cast<Qt::TimeSpec>(spec)));
}

void* QDateTime_NewQDateTime(){
	return new QDateTime();
}

void* QDateTime_NewQDateTime2(void* date){
	return new QDateTime(*static_cast<QDate*>(date));
}

void* QDateTime_NewQDateTime3(void* date, void* time, int spec){
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(time), static_cast<Qt::TimeSpec>(spec));
}

void* QDateTime_NewQDateTime4(void* date, void* time, int spec, int offsetSeconds){
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(time), static_cast<Qt::TimeSpec>(spec), offsetSeconds);
}

void* QDateTime_NewQDateTime5(void* date, void* time, void* timeZone){
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(time), *static_cast<QTimeZone*>(timeZone));
}

void* QDateTime_NewQDateTime6(void* other){
	return new QDateTime(*static_cast<QDateTime*>(other));
}

void* QDateTime_AddMonths(void* ptr, int nmonths){
	return new QDateTime(static_cast<QDateTime*>(ptr)->addMonths(nmonths));
}

void* QDateTime_AddYears(void* ptr, int nyears){
	return new QDateTime(static_cast<QDateTime*>(ptr)->addYears(nyears));
}

int QDateTime_IsDaylightTime(void* ptr){
	return static_cast<QDateTime*>(ptr)->isDaylightTime();
}

int QDateTime_IsNull(void* ptr){
	return static_cast<QDateTime*>(ptr)->isNull();
}

int QDateTime_IsValid(void* ptr){
	return static_cast<QDateTime*>(ptr)->isValid();
}

int QDateTime_OffsetFromUtc(void* ptr){
	return static_cast<QDateTime*>(ptr)->offsetFromUtc();
}

void QDateTime_SetDate(void* ptr, void* date){
	static_cast<QDateTime*>(ptr)->setDate(*static_cast<QDate*>(date));
}

void QDateTime_SetOffsetFromUtc(void* ptr, int offsetSeconds){
	static_cast<QDateTime*>(ptr)->setOffsetFromUtc(offsetSeconds);
}

void QDateTime_SetTime(void* ptr, void* time){
	static_cast<QDateTime*>(ptr)->setTime(*static_cast<QTime*>(time));
}

void QDateTime_SetTimeSpec(void* ptr, int spec){
	static_cast<QDateTime*>(ptr)->setTimeSpec(static_cast<Qt::TimeSpec>(spec));
}

void QDateTime_SetTimeZone(void* ptr, void* toZone){
	static_cast<QDateTime*>(ptr)->setTimeZone(*static_cast<QTimeZone*>(toZone));
}

void QDateTime_Swap(void* ptr, void* other){
	static_cast<QDateTime*>(ptr)->swap(*static_cast<QDateTime*>(other));
}

int QDateTime_TimeSpec(void* ptr){
	return static_cast<QDateTime*>(ptr)->timeSpec();
}

void* QDateTime_TimeZone(void* ptr){
	return new QTimeZone(static_cast<QDateTime*>(ptr)->timeZone());
}

char* QDateTime_TimeZoneAbbreviation(void* ptr){
	return static_cast<QDateTime*>(ptr)->timeZoneAbbreviation().toUtf8().data();
}

void* QDateTime_ToLocalTime(void* ptr){
	return new QDateTime(static_cast<QDateTime*>(ptr)->toLocalTime());
}

char* QDateTime_ToString(void* ptr, char* format){
	return static_cast<QDateTime*>(ptr)->toString(QString(format)).toUtf8().data();
}

void* QDateTime_ToTimeZone(void* ptr, void* timeZone){
	return new QDateTime(static_cast<QDateTime*>(ptr)->toTimeZone(*static_cast<QTimeZone*>(timeZone)));
}

void* QDateTime_ToUTC(void* ptr){
	return new QDateTime(static_cast<QDateTime*>(ptr)->toUTC());
}

void QDateTime_DestroyQDateTime(void* ptr){
	static_cast<QDateTime*>(ptr)->~QDateTime();
}

