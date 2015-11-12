#include "qgeoareamonitorinfo.h"
#include <QDateTime>
#include <QGeoShape>
#include <QDate>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoAreaMonitorInfo>
#include "_cgo_export.h"

class MyQGeoAreaMonitorInfo: public QGeoAreaMonitorInfo {
public:
};

void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(void* other){
	return new QGeoAreaMonitorInfo(*static_cast<QGeoAreaMonitorInfo*>(other));
}

void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(char* name){
	return new QGeoAreaMonitorInfo(QString(name));
}

void* QGeoAreaMonitorInfo_Expiration(void* ptr){
	return new QDateTime(static_cast<QGeoAreaMonitorInfo*>(ptr)->expiration());
}

char* QGeoAreaMonitorInfo_Identifier(void* ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->identifier().toUtf8().data();
}

int QGeoAreaMonitorInfo_IsPersistent(void* ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->isPersistent();
}

int QGeoAreaMonitorInfo_IsValid(void* ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->isValid();
}

char* QGeoAreaMonitorInfo_Name(void* ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->name().toUtf8().data();
}

void QGeoAreaMonitorInfo_SetArea(void* ptr, void* newShape){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setArea(*static_cast<QGeoShape*>(newShape));
}

void QGeoAreaMonitorInfo_SetExpiration(void* ptr, void* expiry){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setExpiration(*static_cast<QDateTime*>(expiry));
}

void QGeoAreaMonitorInfo_SetName(void* ptr, char* name){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setName(QString(name));
}

void QGeoAreaMonitorInfo_SetPersistent(void* ptr, int isPersistent){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setPersistent(isPersistent != 0);
}

void QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(void* ptr){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->~QGeoAreaMonitorInfo();
}

