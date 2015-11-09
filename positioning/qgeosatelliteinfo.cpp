#include "qgeosatelliteinfo.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QGeoSatelliteInfo>
#include "_cgo_export.h"

class MyQGeoSatelliteInfo: public QGeoSatelliteInfo {
public:
};

void* QGeoSatelliteInfo_NewQGeoSatelliteInfo(){
	return new QGeoSatelliteInfo();
}

void* QGeoSatelliteInfo_NewQGeoSatelliteInfo2(void* other){
	return new QGeoSatelliteInfo(*static_cast<QGeoSatelliteInfo*>(other));
}

double QGeoSatelliteInfo_Attribute(void* ptr, int attribute){
	return static_cast<double>(static_cast<QGeoSatelliteInfo*>(ptr)->attribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute)));
}

int QGeoSatelliteInfo_HasAttribute(void* ptr, int attribute){
	return static_cast<QGeoSatelliteInfo*>(ptr)->hasAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

void QGeoSatelliteInfo_RemoveAttribute(void* ptr, int attribute){
	static_cast<QGeoSatelliteInfo*>(ptr)->removeAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

int QGeoSatelliteInfo_SatelliteIdentifier(void* ptr){
	return static_cast<QGeoSatelliteInfo*>(ptr)->satelliteIdentifier();
}

int QGeoSatelliteInfo_SatelliteSystem(void* ptr){
	return static_cast<QGeoSatelliteInfo*>(ptr)->satelliteSystem();
}

void QGeoSatelliteInfo_SetAttribute(void* ptr, int attribute, double value){
	static_cast<QGeoSatelliteInfo*>(ptr)->setAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute), static_cast<qreal>(value));
}

void QGeoSatelliteInfo_SetSatelliteIdentifier(void* ptr, int satId){
	static_cast<QGeoSatelliteInfo*>(ptr)->setSatelliteIdentifier(satId);
}

void QGeoSatelliteInfo_SetSatelliteSystem(void* ptr, int system){
	static_cast<QGeoSatelliteInfo*>(ptr)->setSatelliteSystem(static_cast<QGeoSatelliteInfo::SatelliteSystem>(system));
}

void QGeoSatelliteInfo_SetSignalStrength(void* ptr, int signalStrength){
	static_cast<QGeoSatelliteInfo*>(ptr)->setSignalStrength(signalStrength);
}

int QGeoSatelliteInfo_SignalStrength(void* ptr){
	return static_cast<QGeoSatelliteInfo*>(ptr)->signalStrength();
}

void QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(void* ptr){
	static_cast<QGeoSatelliteInfo*>(ptr)->~QGeoSatelliteInfo();
}

