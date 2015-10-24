#include "qtimezone.h"
#include <QDateTime>
#include <QString>
#include <QModelIndex>
#include <QTime>
#include <QDate>
#include <QByteArray>
#include <QVariant>
#include <QUrl>
#include <QLocale>
#include <QTimeZone>
#include "_cgo_export.h"

class MyQTimeZone: public QTimeZone {
public:
};

QtObjectPtr QTimeZone_NewQTimeZone(){
	return new QTimeZone();
}

QtObjectPtr QTimeZone_NewQTimeZone2(QtObjectPtr ianaId){
	return new QTimeZone(*static_cast<QByteArray*>(ianaId));
}

QtObjectPtr QTimeZone_NewQTimeZone4(QtObjectPtr ianaId, int offsetSeconds, char* name, char* abbreviation, int country, char* comment){
	return new QTimeZone(*static_cast<QByteArray*>(ianaId), offsetSeconds, QString(name), QString(abbreviation), static_cast<QLocale::Country>(country), QString(comment));
}

QtObjectPtr QTimeZone_NewQTimeZone5(QtObjectPtr other){
	return new QTimeZone(*static_cast<QTimeZone*>(other));
}

QtObjectPtr QTimeZone_NewQTimeZone3(int offsetSeconds){
	return new QTimeZone(offsetSeconds);
}

char* QTimeZone_Abbreviation(QtObjectPtr ptr, QtObjectPtr atDateTime){
	return static_cast<QTimeZone*>(ptr)->abbreviation(*static_cast<QDateTime*>(atDateTime)).toUtf8().data();
}

char* QTimeZone_Comment(QtObjectPtr ptr){
	return static_cast<QTimeZone*>(ptr)->comment().toUtf8().data();
}

int QTimeZone_Country(QtObjectPtr ptr){
	return static_cast<QTimeZone*>(ptr)->country();
}

int QTimeZone_DaylightTimeOffset(QtObjectPtr ptr, QtObjectPtr atDateTime){
	return static_cast<QTimeZone*>(ptr)->daylightTimeOffset(*static_cast<QDateTime*>(atDateTime));
}

char* QTimeZone_DisplayName2(QtObjectPtr ptr, int timeType, int nameType, QtObjectPtr locale){
	return static_cast<QTimeZone*>(ptr)->displayName(static_cast<QTimeZone::TimeType>(timeType), static_cast<QTimeZone::NameType>(nameType), *static_cast<QLocale*>(locale)).toUtf8().data();
}

char* QTimeZone_DisplayName(QtObjectPtr ptr, QtObjectPtr atDateTime, int nameType, QtObjectPtr locale){
	return static_cast<QTimeZone*>(ptr)->displayName(*static_cast<QDateTime*>(atDateTime), static_cast<QTimeZone::NameType>(nameType), *static_cast<QLocale*>(locale)).toUtf8().data();
}

int QTimeZone_HasDaylightTime(QtObjectPtr ptr){
	return static_cast<QTimeZone*>(ptr)->hasDaylightTime();
}

int QTimeZone_HasTransitions(QtObjectPtr ptr){
	return static_cast<QTimeZone*>(ptr)->hasTransitions();
}

int QTimeZone_IsDaylightTime(QtObjectPtr ptr, QtObjectPtr atDateTime){
	return static_cast<QTimeZone*>(ptr)->isDaylightTime(*static_cast<QDateTime*>(atDateTime));
}

int QTimeZone_QTimeZone_IsTimeZoneIdAvailable(QtObjectPtr ianaId){
	return QTimeZone::isTimeZoneIdAvailable(*static_cast<QByteArray*>(ianaId));
}

int QTimeZone_IsValid(QtObjectPtr ptr){
	return static_cast<QTimeZone*>(ptr)->isValid();
}

int QTimeZone_OffsetFromUtc(QtObjectPtr ptr, QtObjectPtr atDateTime){
	return static_cast<QTimeZone*>(ptr)->offsetFromUtc(*static_cast<QDateTime*>(atDateTime));
}

int QTimeZone_StandardTimeOffset(QtObjectPtr ptr, QtObjectPtr atDateTime){
	return static_cast<QTimeZone*>(ptr)->standardTimeOffset(*static_cast<QDateTime*>(atDateTime));
}

void QTimeZone_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QTimeZone*>(ptr)->swap(*static_cast<QTimeZone*>(other));
}

void QTimeZone_DestroyQTimeZone(QtObjectPtr ptr){
	static_cast<QTimeZone*>(ptr)->~QTimeZone();
}

