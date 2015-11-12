#include "qgeopositioninfosource.h"
#include <QModelIndex>
#include <QObject>
#include <QGeoPositionInfo>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGeoPositionInfoSource>
#include "_cgo_export.h"

class MyQGeoPositionInfoSource: public QGeoPositionInfoSource {
public:
void Signal_UpdateTimeout(){callbackQGeoPositionInfoSourceUpdateTimeout(this->objectName().toUtf8().data());};
};

void QGeoPositionInfoSource_SetUpdateInterval(void* ptr, int msec){
	static_cast<QGeoPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

char* QGeoPositionInfoSource_SourceName(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->sourceName().toUtf8().data();
}

int QGeoPositionInfoSource_UpdateInterval(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->updateInterval();
}

char* QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources(){
	return QGeoPositionInfoSource::availableSources().join("|").toUtf8().data();
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(void* parent){
	return QGeoPositionInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(char* sourceName, void* parent){
	return QGeoPositionInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

int QGeoPositionInfoSource_Error(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->error();
}

int QGeoPositionInfoSource_MinimumUpdateInterval(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->minimumUpdateInterval();
}

int QGeoPositionInfoSource_PreferredPositioningMethods(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->preferredPositioningMethods();
}

void QGeoPositionInfoSource_RequestUpdate(void* ptr, int timeout){
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "requestUpdate", Q_ARG(int, timeout));
}

void QGeoPositionInfoSource_SetPreferredPositioningMethods(void* ptr, int methods){
	static_cast<QGeoPositionInfoSource*>(ptr)->setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
}

void QGeoPositionInfoSource_StartUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "startUpdates");
}

void QGeoPositionInfoSource_StopUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "stopUpdates");
}

int QGeoPositionInfoSource_SupportedPositioningMethods(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->supportedPositioningMethods();
}

void QGeoPositionInfoSource_ConnectUpdateTimeout(void* ptr){
	QObject::connect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)()>(&QGeoPositionInfoSource::updateTimeout), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)()>(&MyQGeoPositionInfoSource::Signal_UpdateTimeout));;
}

void QGeoPositionInfoSource_DisconnectUpdateTimeout(void* ptr){
	QObject::disconnect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)()>(&QGeoPositionInfoSource::updateTimeout), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)()>(&MyQGeoPositionInfoSource::Signal_UpdateTimeout));;
}

void QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(void* ptr){
	static_cast<QGeoPositionInfoSource*>(ptr)->~QGeoPositionInfoSource();
}

