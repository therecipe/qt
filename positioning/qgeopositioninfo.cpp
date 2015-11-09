#include "qgeopositioninfo.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDate>
#include <QDateTime>
#include <QGeoCoordinate>
#include <QString>
#include <QGeoPositionInfo>
#include "_cgo_export.h"

class MyQGeoPositionInfo: public QGeoPositionInfo {
public:
};

void* QGeoPositionInfo_NewQGeoPositionInfo(){
	return new QGeoPositionInfo();
}

void* QGeoPositionInfo_NewQGeoPositionInfo2(void* coordinate, void* timestamp){
	return new QGeoPositionInfo(*static_cast<QGeoCoordinate*>(coordinate), *static_cast<QDateTime*>(timestamp));
}

void* QGeoPositionInfo_NewQGeoPositionInfo3(void* other){
	return new QGeoPositionInfo(*static_cast<QGeoPositionInfo*>(other));
}

double QGeoPositionInfo_Attribute(void* ptr, int attribute){
	return static_cast<double>(static_cast<QGeoPositionInfo*>(ptr)->attribute(static_cast<QGeoPositionInfo::Attribute>(attribute)));
}

int QGeoPositionInfo_HasAttribute(void* ptr, int attribute){
	return static_cast<QGeoPositionInfo*>(ptr)->hasAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute));
}

int QGeoPositionInfo_IsValid(void* ptr){
	return static_cast<QGeoPositionInfo*>(ptr)->isValid();
}

void QGeoPositionInfo_RemoveAttribute(void* ptr, int attribute){
	static_cast<QGeoPositionInfo*>(ptr)->removeAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute));
}

void QGeoPositionInfo_SetAttribute(void* ptr, int attribute, double value){
	static_cast<QGeoPositionInfo*>(ptr)->setAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute), static_cast<qreal>(value));
}

void QGeoPositionInfo_SetCoordinate(void* ptr, void* coordinate){
	static_cast<QGeoPositionInfo*>(ptr)->setCoordinate(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPositionInfo_SetTimestamp(void* ptr, void* timestamp){
	static_cast<QGeoPositionInfo*>(ptr)->setTimestamp(*static_cast<QDateTime*>(timestamp));
}

void* QGeoPositionInfo_Timestamp(void* ptr){
	return new QDateTime(static_cast<QGeoPositionInfo*>(ptr)->timestamp());
}

void QGeoPositionInfo_DestroyQGeoPositionInfo(void* ptr){
	static_cast<QGeoPositionInfo*>(ptr)->~QGeoPositionInfo();
}

