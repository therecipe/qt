#include "qgeopositioninfosourcefactory.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoPositionInfo>
#include <QGeoPositionInfoSource>
#include <QObject>
#include <QString>
#include <QGeoPositionInfoSourceFactory>
#include "_cgo_export.h"

class MyQGeoPositionInfoSourceFactory: public QGeoPositionInfoSourceFactory {
public:
};

void* QGeoPositionInfoSourceFactory_AreaMonitor(void* ptr, void* parent){
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->areaMonitor(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSourceFactory_PositionInfoSource(void* ptr, void* parent){
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->positionInfoSource(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSourceFactory_SatelliteInfoSource(void* ptr, void* parent){
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->satelliteInfoSource(static_cast<QObject*>(parent));
}

void QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(void* ptr){
	static_cast<QGeoPositionInfoSourceFactory*>(ptr)->~QGeoPositionInfoSourceFactory();
}

