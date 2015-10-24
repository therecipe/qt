#include "qgeoroutereply.h"
#include <QGeoRoute>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QGeoRouteReply>
#include "_cgo_export.h"

class MyQGeoRouteReply: public QGeoRouteReply {
public:
void Signal_Finished(){callbackQGeoRouteReplyFinished(this->objectName().toUtf8().data());};
};

QtObjectPtr QGeoRouteReply_NewQGeoRouteReply(int error, char* errorString, QtObjectPtr parent){
	return new QGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString(errorString), static_cast<QObject*>(parent));
}

void QGeoRouteReply_Abort(QtObjectPtr ptr){
	static_cast<QGeoRouteReply*>(ptr)->abort();
}

int QGeoRouteReply_Error(QtObjectPtr ptr){
	return static_cast<QGeoRouteReply*>(ptr)->error();
}

char* QGeoRouteReply_ErrorString(QtObjectPtr ptr){
	return static_cast<QGeoRouteReply*>(ptr)->errorString().toUtf8().data();
}

void QGeoRouteReply_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::finished), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Finished));;
}

void QGeoRouteReply_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::finished), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Finished));;
}

int QGeoRouteReply_IsFinished(QtObjectPtr ptr){
	return static_cast<QGeoRouteReply*>(ptr)->isFinished();
}

void QGeoRouteReply_DestroyQGeoRouteReply(QtObjectPtr ptr){
	static_cast<QGeoRouteReply*>(ptr)->~QGeoRouteReply();
}

