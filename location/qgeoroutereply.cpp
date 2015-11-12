#include "qgeoroutereply.h"
#include <QModelIndex>
#include <QGeoRoute>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGeoRouteReply>
#include "_cgo_export.h"

class MyQGeoRouteReply: public QGeoRouteReply {
public:
void Signal_Finished(){callbackQGeoRouteReplyFinished(this->objectName().toUtf8().data());};
};

void* QGeoRouteReply_NewQGeoRouteReply(int error, char* errorString, void* parent){
	return new QGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString(errorString), static_cast<QObject*>(parent));
}

void QGeoRouteReply_Abort(void* ptr){
	static_cast<QGeoRouteReply*>(ptr)->abort();
}

int QGeoRouteReply_Error(void* ptr){
	return static_cast<QGeoRouteReply*>(ptr)->error();
}

char* QGeoRouteReply_ErrorString(void* ptr){
	return static_cast<QGeoRouteReply*>(ptr)->errorString().toUtf8().data();
}

void QGeoRouteReply_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::finished), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Finished));;
}

void QGeoRouteReply_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::finished), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Finished));;
}

int QGeoRouteReply_IsFinished(void* ptr){
	return static_cast<QGeoRouteReply*>(ptr)->isFinished();
}

void QGeoRouteReply_DestroyQGeoRouteReply(void* ptr){
	static_cast<QGeoRouteReply*>(ptr)->~QGeoRouteReply();
}

