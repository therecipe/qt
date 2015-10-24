#include "qgeoareamonitorinfo.h"
#include <QModelIndex>
#include <QGeoShape>
#include <QDate>
#include <QDateTime>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGeoAreaMonitorInfo>
#include "_cgo_export.h"

class MyQGeoAreaMonitorInfo: public QGeoAreaMonitorInfo {
public:
};

QtObjectPtr QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(QtObjectPtr other){
	return new QGeoAreaMonitorInfo(*static_cast<QGeoAreaMonitorInfo*>(other));
}

QtObjectPtr QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(char* name){
	return new QGeoAreaMonitorInfo(QString(name));
}

char* QGeoAreaMonitorInfo_Identifier(QtObjectPtr ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->identifier().toUtf8().data();
}

int QGeoAreaMonitorInfo_IsPersistent(QtObjectPtr ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->isPersistent();
}

int QGeoAreaMonitorInfo_IsValid(QtObjectPtr ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->isValid();
}

char* QGeoAreaMonitorInfo_Name(QtObjectPtr ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->name().toUtf8().data();
}

void QGeoAreaMonitorInfo_SetArea(QtObjectPtr ptr, QtObjectPtr newShape){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setArea(*static_cast<QGeoShape*>(newShape));
}

void QGeoAreaMonitorInfo_SetExpiration(QtObjectPtr ptr, QtObjectPtr expiry){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setExpiration(*static_cast<QDateTime*>(expiry));
}

void QGeoAreaMonitorInfo_SetName(QtObjectPtr ptr, char* name){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setName(QString(name));
}

void QGeoAreaMonitorInfo_SetPersistent(QtObjectPtr ptr, int isPersistent){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setPersistent(isPersistent != 0);
}

void QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(QtObjectPtr ptr){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->~QGeoAreaMonitorInfo();
}

