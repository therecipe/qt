#include "qgeoareamonitorsource.h"
#include <QUrl>
#include <QModelIndex>
#include <QGeoPositionInfoSource>
#include <QGeoAreaMonitorInfo>
#include <QGeoPositionInfo>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QGeoAreaMonitorSource>
#include "_cgo_export.h"

class MyQGeoAreaMonitorSource: public QGeoAreaMonitorSource {
public:
};

char* QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources(){
	return QGeoAreaMonitorSource::availableSources().join("|").toUtf8().data();
}

QtObjectPtr QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(QtObjectPtr parent){
	return QGeoAreaMonitorSource::createDefaultSource(static_cast<QObject*>(parent));
}

QtObjectPtr QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(char* sourceName, QtObjectPtr parent){
	return QGeoAreaMonitorSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

int QGeoAreaMonitorSource_Error(QtObjectPtr ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->error();
}

QtObjectPtr QGeoAreaMonitorSource_PositionInfoSource(QtObjectPtr ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->positionInfoSource();
}

int QGeoAreaMonitorSource_RequestUpdate(QtObjectPtr ptr, QtObjectPtr monitor, char* signal){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->requestUpdate(*static_cast<QGeoAreaMonitorInfo*>(monitor), const_cast<const char*>(signal));
}

void QGeoAreaMonitorSource_SetPositionInfoSource(QtObjectPtr ptr, QtObjectPtr newSource){
	static_cast<QGeoAreaMonitorSource*>(ptr)->setPositionInfoSource(static_cast<QGeoPositionInfoSource*>(newSource));
}

char* QGeoAreaMonitorSource_SourceName(QtObjectPtr ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->sourceName().toUtf8().data();
}

int QGeoAreaMonitorSource_StartMonitoring(QtObjectPtr ptr, QtObjectPtr monitor){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->startMonitoring(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

int QGeoAreaMonitorSource_StopMonitoring(QtObjectPtr ptr, QtObjectPtr monitor){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->stopMonitoring(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

int QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(QtObjectPtr ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->supportedAreaMonitorFeatures();
}

void QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(QtObjectPtr ptr){
	static_cast<QGeoAreaMonitorSource*>(ptr)->~QGeoAreaMonitorSource();
}

