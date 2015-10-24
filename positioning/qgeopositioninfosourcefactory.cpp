#include "qgeopositioninfosourcefactory.h"
#include <QUrl>
#include <QModelIndex>
#include <QGeoPositionInfoSource>
#include <QGeoPositionInfo>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QGeoPositionInfoSourceFactory>
#include "_cgo_export.h"

class MyQGeoPositionInfoSourceFactory: public QGeoPositionInfoSourceFactory {
public:
};

QtObjectPtr QGeoPositionInfoSourceFactory_AreaMonitor(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->areaMonitor(static_cast<QObject*>(parent));
}

QtObjectPtr QGeoPositionInfoSourceFactory_PositionInfoSource(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->positionInfoSource(static_cast<QObject*>(parent));
}

QtObjectPtr QGeoPositionInfoSourceFactory_SatelliteInfoSource(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->satelliteInfoSource(static_cast<QObject*>(parent));
}

void QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(QtObjectPtr ptr){
	static_cast<QGeoPositionInfoSourceFactory*>(ptr)->~QGeoPositionInfoSourceFactory();
}

