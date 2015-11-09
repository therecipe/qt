#include "qtimezone.h"
#include <QDateTime>
#include <QLocale>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QDate>
#include <QUrl>
#include <QByteArray>
#include <QTime>
#include <QTimeZone>
#include "_cgo_export.h"

class MyQTimeZone: public QTimeZone {
public:
};

void* QTimeZone_NewQTimeZone(){
	return new QTimeZone();
}

void* QTimeZone_NewQTimeZone2(void* ianaId){
	return new QTimeZone(*static_cast<QByteArray*>(ianaId));
}

void* QTimeZone_NewQTimeZone4(void* ianaId, int offsetSeconds, char* name, char* abbreviation, int country, char* comment){
	return new QTimeZone(*static_cast<QByteArray*>(ianaId), offsetSeconds, QString(name), QString(abbreviation), static_cast<QLocale::Country>(country), QString(comment));
}

void* QTimeZone_NewQTimeZone5(void* other){
	return new QTimeZone(*static_cast<QTimeZone*>(other));
}

void* QTimeZone_NewQTimeZone3(int offsetSeconds){
	return new QTimeZone(offsetSeconds);
}

char* QTimeZone_Abbreviation(void* ptr, void* atDateTime){
	return static_cast<QTimeZone*>(ptr)->abbreviation(*static_cast<QDateTime*>(atDateTime)).toUtf8().data();
}

char* QTimeZone_Comment(void* ptr){
	return static_cast<QTimeZone*>(ptr)->comment().toUtf8().data();
}

int QTimeZone_Country(void* ptr){
	return static_cast<QTimeZone*>(ptr)->country();
}

int QTimeZone_DaylightTimeOffset(void* ptr, void* atDateTime){
	return static_cast<QTimeZone*>(ptr)->daylightTimeOffset(*static_cast<QDateTime*>(atDateTime));
}

char* QTimeZone_DisplayName2(void* ptr, int timeType, int nameType, void* locale){
	return static_cast<QTimeZone*>(ptr)->displayName(static_cast<QTimeZone::TimeType>(timeType), static_cast<QTimeZone::NameType>(nameType), *static_cast<QLocale*>(locale)).toUtf8().data();
}

char* QTimeZone_DisplayName(void* ptr, void* atDateTime, int nameType, void* locale){
	return static_cast<QTimeZone*>(ptr)->displayName(*static_cast<QDateTime*>(atDateTime), static_cast<QTimeZone::NameType>(nameType), *static_cast<QLocale*>(locale)).toUtf8().data();
}

int QTimeZone_HasDaylightTime(void* ptr){
	return static_cast<QTimeZone*>(ptr)->hasDaylightTime();
}

int QTimeZone_HasTransitions(void* ptr){
	return static_cast<QTimeZone*>(ptr)->hasTransitions();
}

void* QTimeZone_QTimeZone_IanaIdToWindowsId(void* ianaId){
	return new QByteArray(QTimeZone::ianaIdToWindowsId(*static_cast<QByteArray*>(ianaId)));
}

void* QTimeZone_Id(void* ptr){
	return new QByteArray(static_cast<QTimeZone*>(ptr)->id());
}

int QTimeZone_IsDaylightTime(void* ptr, void* atDateTime){
	return static_cast<QTimeZone*>(ptr)->isDaylightTime(*static_cast<QDateTime*>(atDateTime));
}

int QTimeZone_QTimeZone_IsTimeZoneIdAvailable(void* ianaId){
	return QTimeZone::isTimeZoneIdAvailable(*static_cast<QByteArray*>(ianaId));
}

int QTimeZone_IsValid(void* ptr){
	return static_cast<QTimeZone*>(ptr)->isValid();
}

int QTimeZone_OffsetFromUtc(void* ptr, void* atDateTime){
	return static_cast<QTimeZone*>(ptr)->offsetFromUtc(*static_cast<QDateTime*>(atDateTime));
}

int QTimeZone_StandardTimeOffset(void* ptr, void* atDateTime){
	return static_cast<QTimeZone*>(ptr)->standardTimeOffset(*static_cast<QDateTime*>(atDateTime));
}

void QTimeZone_Swap(void* ptr, void* other){
	static_cast<QTimeZone*>(ptr)->swap(*static_cast<QTimeZone*>(other));
}

void* QTimeZone_QTimeZone_SystemTimeZone(){
	return new QTimeZone(QTimeZone::systemTimeZone());
}

void* QTimeZone_QTimeZone_SystemTimeZoneId(){
	return new QByteArray(QTimeZone::systemTimeZoneId());
}

void* QTimeZone_QTimeZone_Utc(){
	return new QTimeZone(QTimeZone::utc());
}

void* QTimeZone_QTimeZone_WindowsIdToDefaultIanaId(void* windowsId){
	return new QByteArray(QTimeZone::windowsIdToDefaultIanaId(*static_cast<QByteArray*>(windowsId)));
}

void* QTimeZone_QTimeZone_WindowsIdToDefaultIanaId2(void* windowsId, int country){
	return new QByteArray(QTimeZone::windowsIdToDefaultIanaId(*static_cast<QByteArray*>(windowsId), static_cast<QLocale::Country>(country)));
}

void QTimeZone_DestroyQTimeZone(void* ptr){
	static_cast<QTimeZone*>(ptr)->~QTimeZone();
}

