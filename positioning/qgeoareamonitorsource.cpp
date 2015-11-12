#include "qgeoareamonitorsource.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoAreaMonitorInfo>
#include <QGeoPositionInfoSource>
#include <QObject>
#include <QGeoPositionInfo>
#include <QGeoAreaMonitorSource>
#include "_cgo_export.h"

class MyQGeoAreaMonitorSource: public QGeoAreaMonitorSource {
public:
};

char* QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources(){
	return QGeoAreaMonitorSource::availableSources().join("|").toUtf8().data();
}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(void* parent){
	return QGeoAreaMonitorSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(char* sourceName, void* parent){
	return QGeoAreaMonitorSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

int QGeoAreaMonitorSource_Error(void* ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->error();
}

void* QGeoAreaMonitorSource_PositionInfoSource(void* ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->positionInfoSource();
}

int QGeoAreaMonitorSource_RequestUpdate(void* ptr, void* monitor, char* signal){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->requestUpdate(*static_cast<QGeoAreaMonitorInfo*>(monitor), const_cast<const char*>(signal));
}

void QGeoAreaMonitorSource_SetPositionInfoSource(void* ptr, void* newSource){
	static_cast<QGeoAreaMonitorSource*>(ptr)->setPositionInfoSource(static_cast<QGeoPositionInfoSource*>(newSource));
}

char* QGeoAreaMonitorSource_SourceName(void* ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->sourceName().toUtf8().data();
}

int QGeoAreaMonitorSource_StartMonitoring(void* ptr, void* monitor){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->startMonitoring(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

int QGeoAreaMonitorSource_StopMonitoring(void* ptr, void* monitor){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->stopMonitoring(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

int QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(void* ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->supportedAreaMonitorFeatures();
}

void QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(void* ptr){
	static_cast<QGeoAreaMonitorSource*>(ptr)->~QGeoAreaMonitorSource();
}

