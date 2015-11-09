#include "qnmeapositioninfosource.h"
#include <QModelIndex>
#include <QIODevice>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QNmeaPositionInfoSource>
#include "_cgo_export.h"

class MyQNmeaPositionInfoSource: public QNmeaPositionInfoSource {
public:
};

void* QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(int updateMode, void* parent){
	return new QNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QObject*>(parent));
}

void* QNmeaPositionInfoSource_Device(void* ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->device();
}

int QNmeaPositionInfoSource_Error(void* ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->error();
}

int QNmeaPositionInfoSource_MinimumUpdateInterval(void* ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->minimumUpdateInterval();
}

void QNmeaPositionInfoSource_RequestUpdate(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "requestUpdate", Q_ARG(int, msec));
}

void QNmeaPositionInfoSource_SetDevice(void* ptr, void* device){
	static_cast<QNmeaPositionInfoSource*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QNmeaPositionInfoSource_SetUpdateInterval(void* ptr, int msec){
	static_cast<QNmeaPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QNmeaPositionInfoSource_StartUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "startUpdates");
}

void QNmeaPositionInfoSource_StopUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "stopUpdates");
}

int QNmeaPositionInfoSource_SupportedPositioningMethods(void* ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->supportedPositioningMethods();
}

int QNmeaPositionInfoSource_UpdateMode(void* ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->updateMode();
}

void QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(void* ptr){
	static_cast<QNmeaPositionInfoSource*>(ptr)->~QNmeaPositionInfoSource();
}

