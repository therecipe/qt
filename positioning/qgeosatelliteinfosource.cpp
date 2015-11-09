#include "qgeosatelliteinfosource.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoSatelliteInfo>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QGeoSatelliteInfoSource>
#include "_cgo_export.h"

class MyQGeoSatelliteInfoSource: public QGeoSatelliteInfoSource {
public:
void Signal_RequestTimeout(){callbackQGeoSatelliteInfoSourceRequestTimeout(this->objectName().toUtf8().data());};
};

void QGeoSatelliteInfoSource_SetUpdateInterval(void* ptr, int msec){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->setUpdateInterval(msec);
}

int QGeoSatelliteInfoSource_UpdateInterval(void* ptr){
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->updateInterval();
}

char* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources(){
	return QGeoSatelliteInfoSource::availableSources().join("|").toUtf8().data();
}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(void* parent){
	return QGeoSatelliteInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(char* sourceName, void* parent){
	return QGeoSatelliteInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

int QGeoSatelliteInfoSource_Error(void* ptr){
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->error();
}

int QGeoSatelliteInfoSource_MinimumUpdateInterval(void* ptr){
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->minimumUpdateInterval();
}

void QGeoSatelliteInfoSource_ConnectRequestTimeout(void* ptr){
	QObject::connect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)()>(&QGeoSatelliteInfoSource::requestTimeout), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)()>(&MyQGeoSatelliteInfoSource::Signal_RequestTimeout));;
}

void QGeoSatelliteInfoSource_DisconnectRequestTimeout(void* ptr){
	QObject::disconnect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)()>(&QGeoSatelliteInfoSource::requestTimeout), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)()>(&MyQGeoSatelliteInfoSource::Signal_RequestTimeout));;
}

void QGeoSatelliteInfoSource_RequestUpdate(void* ptr, int timeout){
	QMetaObject::invokeMethod(static_cast<QGeoSatelliteInfoSource*>(ptr), "requestUpdate", Q_ARG(int, timeout));
}

char* QGeoSatelliteInfoSource_SourceName(void* ptr){
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->sourceName().toUtf8().data();
}

void QGeoSatelliteInfoSource_StartUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGeoSatelliteInfoSource*>(ptr), "startUpdates");
}

void QGeoSatelliteInfoSource_StopUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGeoSatelliteInfoSource*>(ptr), "stopUpdates");
}

void QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(void* ptr){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->~QGeoSatelliteInfoSource();
}

