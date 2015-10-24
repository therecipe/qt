#include "qnmeapositioninfosource.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QMetaObject>
#include <QObject>
#include <QNmeaPositionInfoSource>
#include "_cgo_export.h"

class MyQNmeaPositionInfoSource: public QNmeaPositionInfoSource {
public:
};

QtObjectPtr QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(int updateMode, QtObjectPtr parent){
	return new QNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QObject*>(parent));
}

QtObjectPtr QNmeaPositionInfoSource_Device(QtObjectPtr ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->device();
}

int QNmeaPositionInfoSource_Error(QtObjectPtr ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->error();
}

int QNmeaPositionInfoSource_MinimumUpdateInterval(QtObjectPtr ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->minimumUpdateInterval();
}

void QNmeaPositionInfoSource_RequestUpdate(QtObjectPtr ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "requestUpdate", Q_ARG(int, msec));
}

void QNmeaPositionInfoSource_SetDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QNmeaPositionInfoSource*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QNmeaPositionInfoSource_SetUpdateInterval(QtObjectPtr ptr, int msec){
	static_cast<QNmeaPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QNmeaPositionInfoSource_StartUpdates(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "startUpdates");
}

void QNmeaPositionInfoSource_StopUpdates(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "stopUpdates");
}

int QNmeaPositionInfoSource_SupportedPositioningMethods(QtObjectPtr ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->supportedPositioningMethods();
}

int QNmeaPositionInfoSource_UpdateMode(QtObjectPtr ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->updateMode();
}

void QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(QtObjectPtr ptr){
	static_cast<QNmeaPositionInfoSource*>(ptr)->~QNmeaPositionInfoSource();
}

