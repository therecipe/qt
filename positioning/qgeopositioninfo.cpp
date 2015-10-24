#include "qgeopositioninfo.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDate>
#include <QGeoCoordinate>
#include <QDateTime>
#include <QGeoPositionInfo>
#include "_cgo_export.h"

class MyQGeoPositionInfo: public QGeoPositionInfo {
public:
};

QtObjectPtr QGeoPositionInfo_NewQGeoPositionInfo(){
	return new QGeoPositionInfo();
}

QtObjectPtr QGeoPositionInfo_NewQGeoPositionInfo2(QtObjectPtr coordinate, QtObjectPtr timestamp){
	return new QGeoPositionInfo(*static_cast<QGeoCoordinate*>(coordinate), *static_cast<QDateTime*>(timestamp));
}

QtObjectPtr QGeoPositionInfo_NewQGeoPositionInfo3(QtObjectPtr other){
	return new QGeoPositionInfo(*static_cast<QGeoPositionInfo*>(other));
}

int QGeoPositionInfo_HasAttribute(QtObjectPtr ptr, int attribute){
	return static_cast<QGeoPositionInfo*>(ptr)->hasAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute));
}

int QGeoPositionInfo_IsValid(QtObjectPtr ptr){
	return static_cast<QGeoPositionInfo*>(ptr)->isValid();
}

void QGeoPositionInfo_RemoveAttribute(QtObjectPtr ptr, int attribute){
	static_cast<QGeoPositionInfo*>(ptr)->removeAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute));
}

void QGeoPositionInfo_SetCoordinate(QtObjectPtr ptr, QtObjectPtr coordinate){
	static_cast<QGeoPositionInfo*>(ptr)->setCoordinate(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPositionInfo_SetTimestamp(QtObjectPtr ptr, QtObjectPtr timestamp){
	static_cast<QGeoPositionInfo*>(ptr)->setTimestamp(*static_cast<QDateTime*>(timestamp));
}

void QGeoPositionInfo_DestroyQGeoPositionInfo(QtObjectPtr ptr){
	static_cast<QGeoPositionInfo*>(ptr)->~QGeoPositionInfo();
}

