#include "qgeopositioninfosource.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QGeoPositionInfo>
#include <QGeoPositionInfoSource>
#include "_cgo_export.h"

class MyQGeoPositionInfoSource: public QGeoPositionInfoSource {
public:
void Signal_UpdateTimeout(){callbackQGeoPositionInfoSourceUpdateTimeout(this->objectName().toUtf8().data());};
};

void QGeoPositionInfoSource_SetUpdateInterval(QtObjectPtr ptr, int msec){
	static_cast<QGeoPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

char* QGeoPositionInfoSource_SourceName(QtObjectPtr ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->sourceName().toUtf8().data();
}

int QGeoPositionInfoSource_UpdateInterval(QtObjectPtr ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->updateInterval();
}

char* QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources(){
	return QGeoPositionInfoSource::availableSources().join("|").toUtf8().data();
}

QtObjectPtr QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(QtObjectPtr parent){
	return QGeoPositionInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

QtObjectPtr QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(char* sourceName, QtObjectPtr parent){
	return QGeoPositionInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

int QGeoPositionInfoSource_Error(QtObjectPtr ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->error();
}

int QGeoPositionInfoSource_MinimumUpdateInterval(QtObjectPtr ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->minimumUpdateInterval();
}

int QGeoPositionInfoSource_PreferredPositioningMethods(QtObjectPtr ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->preferredPositioningMethods();
}

void QGeoPositionInfoSource_RequestUpdate(QtObjectPtr ptr, int timeout){
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "requestUpdate", Q_ARG(int, timeout));
}

void QGeoPositionInfoSource_SetPreferredPositioningMethods(QtObjectPtr ptr, int methods){
	static_cast<QGeoPositionInfoSource*>(ptr)->setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
}

void QGeoPositionInfoSource_StartUpdates(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "startUpdates");
}

void QGeoPositionInfoSource_StopUpdates(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "stopUpdates");
}

int QGeoPositionInfoSource_SupportedPositioningMethods(QtObjectPtr ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->supportedPositioningMethods();
}

void QGeoPositionInfoSource_ConnectUpdateTimeout(QtObjectPtr ptr){
	QObject::connect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)()>(&QGeoPositionInfoSource::updateTimeout), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)()>(&MyQGeoPositionInfoSource::Signal_UpdateTimeout));;
}

void QGeoPositionInfoSource_DisconnectUpdateTimeout(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)()>(&QGeoPositionInfoSource::updateTimeout), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)()>(&MyQGeoPositionInfoSource::Signal_UpdateTimeout));;
}

void QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(QtObjectPtr ptr){
	static_cast<QGeoPositionInfoSource*>(ptr)->~QGeoPositionInfoSource();
}

