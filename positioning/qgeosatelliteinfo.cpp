#include "qgeosatelliteinfo.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoSatelliteInfo>
#include "_cgo_export.h"

class MyQGeoSatelliteInfo: public QGeoSatelliteInfo {
public:
};

QtObjectPtr QGeoSatelliteInfo_NewQGeoSatelliteInfo(){
	return new QGeoSatelliteInfo();
}

QtObjectPtr QGeoSatelliteInfo_NewQGeoSatelliteInfo2(QtObjectPtr other){
	return new QGeoSatelliteInfo(*static_cast<QGeoSatelliteInfo*>(other));
}

int QGeoSatelliteInfo_HasAttribute(QtObjectPtr ptr, int attribute){
	return static_cast<QGeoSatelliteInfo*>(ptr)->hasAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

void QGeoSatelliteInfo_RemoveAttribute(QtObjectPtr ptr, int attribute){
	static_cast<QGeoSatelliteInfo*>(ptr)->removeAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

int QGeoSatelliteInfo_SatelliteIdentifier(QtObjectPtr ptr){
	return static_cast<QGeoSatelliteInfo*>(ptr)->satelliteIdentifier();
}

int QGeoSatelliteInfo_SatelliteSystem(QtObjectPtr ptr){
	return static_cast<QGeoSatelliteInfo*>(ptr)->satelliteSystem();
}

void QGeoSatelliteInfo_SetSatelliteIdentifier(QtObjectPtr ptr, int satId){
	static_cast<QGeoSatelliteInfo*>(ptr)->setSatelliteIdentifier(satId);
}

void QGeoSatelliteInfo_SetSatelliteSystem(QtObjectPtr ptr, int system){
	static_cast<QGeoSatelliteInfo*>(ptr)->setSatelliteSystem(static_cast<QGeoSatelliteInfo::SatelliteSystem>(system));
}

void QGeoSatelliteInfo_SetSignalStrength(QtObjectPtr ptr, int signalStrength){
	static_cast<QGeoSatelliteInfo*>(ptr)->setSignalStrength(signalStrength);
}

int QGeoSatelliteInfo_SignalStrength(QtObjectPtr ptr){
	return static_cast<QGeoSatelliteInfo*>(ptr)->signalStrength();
}

void QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(QtObjectPtr ptr){
	static_cast<QGeoSatelliteInfo*>(ptr)->~QGeoSatelliteInfo();
}

